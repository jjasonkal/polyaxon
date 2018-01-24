export class JobModel {
	public uuid: string;
	public unique_name: string;
	public sequence: number;
	public role: string;
	public last_status: string;
	public experiment_name: string;
	public experiment: string;
	public definition: string;
	public deleted?: boolean;
	public project?: string;
	public status?: string;
	public created_at: string;
	public updated_at: string;
	public started_at: string;
	public finished_at: string;
}

export class JobStateSchema {
	byUuids: {[uuid: string]: JobModel};
	uuids: string[];
}

export const JobsEmptyState = {byUuids: {}, uuids: []};
