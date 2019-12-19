// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package config_resources_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateConfigResourceReader is a Reader for the UpdateConfigResource structure.
type UpdateConfigResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConfigResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateConfigResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateConfigResourceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateConfigResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConfigResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateConfigResourceOK creates a UpdateConfigResourceOK with default headers values
func NewUpdateConfigResourceOK() *UpdateConfigResourceOK {
	return &UpdateConfigResourceOK{}
}

/*UpdateConfigResourceOK handles this case with default header values.

A successful response.
*/
type UpdateConfigResourceOK struct {
	Payload *service_model.V1ConfigResource
}

func (o *UpdateConfigResourceOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/config_resources/{config_resource.uuid}][%d] updateConfigResourceOK  %+v", 200, o.Payload)
}

func (o *UpdateConfigResourceOK) GetPayload() *service_model.V1ConfigResource {
	return o.Payload
}

func (o *UpdateConfigResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ConfigResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConfigResourceNoContent creates a UpdateConfigResourceNoContent with default headers values
func NewUpdateConfigResourceNoContent() *UpdateConfigResourceNoContent {
	return &UpdateConfigResourceNoContent{}
}

/*UpdateConfigResourceNoContent handles this case with default header values.

No content.
*/
type UpdateConfigResourceNoContent struct {
	Payload interface{}
}

func (o *UpdateConfigResourceNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/config_resources/{config_resource.uuid}][%d] updateConfigResourceNoContent  %+v", 204, o.Payload)
}

func (o *UpdateConfigResourceNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateConfigResourceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConfigResourceForbidden creates a UpdateConfigResourceForbidden with default headers values
func NewUpdateConfigResourceForbidden() *UpdateConfigResourceForbidden {
	return &UpdateConfigResourceForbidden{}
}

/*UpdateConfigResourceForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UpdateConfigResourceForbidden struct {
	Payload interface{}
}

func (o *UpdateConfigResourceForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/config_resources/{config_resource.uuid}][%d] updateConfigResourceForbidden  %+v", 403, o.Payload)
}

func (o *UpdateConfigResourceForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateConfigResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConfigResourceNotFound creates a UpdateConfigResourceNotFound with default headers values
func NewUpdateConfigResourceNotFound() *UpdateConfigResourceNotFound {
	return &UpdateConfigResourceNotFound{}
}

/*UpdateConfigResourceNotFound handles this case with default header values.

Resource does not exist.
*/
type UpdateConfigResourceNotFound struct {
	Payload interface{}
}

func (o *UpdateConfigResourceNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/config_resources/{config_resource.uuid}][%d] updateConfigResourceNotFound  %+v", 404, o.Payload)
}

func (o *UpdateConfigResourceNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateConfigResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}