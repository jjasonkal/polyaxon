import logging

from docker.errors import DockerException

from django.db.models import Count

from experiments.models import Experiment, ExperimentMetric
from experiments.paths import create_experiment_outputs_path
from experiments.restart import handle_restarted_experiment
from experiments.statuses import ExperimentLifeCycle
from experiments.utils import get_valid_experiment
from polyaxon.celery_api import app as celery_app
from polyaxon.settings import CeleryTasks, Intervals
from repos.models import Repo
from runner.dockerizer.builders import experiments as experiments_builder
from runner.schedulers import experiment_scheduler

logger = logging.getLogger('polyaxon.tasks.experiments')


@celery_app.task(name=CeleryTasks.EXPERIMENTS_BUILD, bind=True, max_retries=3)
def build_experiment(self, experiment_id):
    experiment = get_valid_experiment(experiment_id=experiment_id)
    if not experiment:
        if self.request.retries < 2:
            logger.info('Trying again for Experiment `%s`.', experiment_id)
            self.retry(countdown=Intervals.EXPERIMENTS_SCHEDULER)

        logger.info('Something went wrong, '
                    'the Experiment `%s` does not exist anymore.', experiment_id)
        return

    # No need to build the image, start the experiment directly
    if not experiment.compiled_spec.run_exec:
        start_experiment.delay(experiment_id=experiment_id)
        return

    if not ExperimentLifeCycle.can_transition(status_from=experiment.last_status,
                                              status_to=ExperimentLifeCycle.BUILDING):
        logger.info('Experiment id `%s` cannot transition from `%s` to `%s`.',
                    experiment_id, experiment.last_status, ExperimentLifeCycle.BUILDING)
        return None

    # Update experiment status to show that its building
    experiment.set_status(ExperimentLifeCycle.BUILDING)

    # Building the docker image
    try:
        status = experiments_builder.build_experiment(experiment)
    except DockerException as e:
        logger.warning('Failed to build experiment %s', e)
        experiment.set_status(ExperimentLifeCycle.FAILED,
                              message='Failed to build image for experiment.')
        return
    except Repo.DoesNotExist:
        logger.warning('No code was found for this project')
        experiment.set_status(ExperimentLifeCycle.FAILED,
                              message='No code was found for to build this experiment.')
        return

    if not status:
        return

    # Now we can start the experiment
    start_experiment.delay(experiment_id=experiment_id)


@celery_app.task(name=CeleryTasks.EXPERIMENTS_START, ignore_result=True)
def start_experiment(experiment_id):
    experiment = get_valid_experiment(experiment_id=experiment_id)
    if not experiment:
        logger.info('Something went wrong, '
                    'the Experiment `%s` does not exist anymore.', experiment_id)
        return

    if not ExperimentLifeCycle.can_transition(status_from=experiment.last_status,
                                              status_to=ExperimentLifeCycle.SCHEDULED):
        logger.info('Experiment id `%s` cannot transition from `%s` to `%s`.',
                    experiment_id, experiment.last_status, ExperimentLifeCycle.BUILDING)
        return None

    # Check if we need to restart an experiment
    if experiment.is_clone:
        handle_restarted_experiment(experiment)
    else:
        create_experiment_outputs_path(experiment.unique_name)

    experiment_scheduler.start_experiment(experiment)


@celery_app.task(name=CeleryTasks.EXPERIMENTS_STOP, ignore_result=True)
def stop_experiment(experiment_id):
    experiment = get_valid_experiment(experiment_id=experiment_id)
    if not experiment:
        logger.info('Something went wrong, '
                    'the Experiment `%s` does not exist anymore.', experiment_id)
        return

    experiment_scheduler.stop_experiment(experiment, update_status=True)


@celery_app.task(name=CeleryTasks.EXPERIMENTS_CHECK_STATUS, ignore_result=True)
def check_experiment_status(experiment_uuid):
    experiment = Experiment.objects.get(uuid=experiment_uuid)
    experiment.update_status()


@celery_app.task(name=CeleryTasks.EXPERIMENTS_SET_METRICS, ignore_result=True)
def set_metrics(experiment_uuid, metrics, created_at=None):
    try:
        experiment = Experiment.objects.get(uuid=experiment_uuid)
    except Experiment.DoesNotExist:
        logger.info('Experiment uuid `%s` does not exist', experiment_uuid)
        return None

    kwargs = {}
    if created_at:
        kwargs['created_at'] = created_at
    ExperimentMetric.objects.create(experiment=experiment, values=metrics, **kwargs)


@celery_app.task(name=CeleryTasks.EXPERIMENTS_SYNC_JOBS_STATUSES, ignore_result=True)
def sync_experiments_and_jobs_statuses():
    experiments = Experiment.objects.exclude(
        experiment_status__status__in=ExperimentLifeCycle.DONE_STATUS)
    experiments = experiments.annotate(num_jobs=Count('jobs')).filter(num_jobs__gt=0)
    for experiment in experiments:
        check_experiment_status.delay(experiment_uuid=experiment.uuid.hex)
