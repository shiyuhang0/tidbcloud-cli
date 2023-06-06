// Code generated by go-swagger; DO NOT EDIT.

package branch_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"tidbcloud-cli/pkg/tidbcloud/branch/models"
)

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUserOK creates a DeleteUserOK with default headers values
func NewDeleteUserOK() *DeleteUserOK {
	return &DeleteUserOK{}
}

/*
DeleteUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteUserOK struct {
	Payload interface{}
}

// IsSuccess returns true when this delete user o k response has a 2xx status code
func (o *DeleteUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user o k response has a 3xx status code
func (o *DeleteUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user o k response has a 4xx status code
func (o *DeleteUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user o k response has a 5xx status code
func (o *DeleteUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user o k response a status code equal to that given
func (o *DeleteUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete user o k response
func (o *DeleteUserOK) Code() int {
	return 200
}

func (o *DeleteUserOK) Error() string {
	return fmt.Sprintf("[DELETE /api/internal/projects/{project_id}/clusters/{cluster_id}/branches/{branch_name}/users/{username}][%d] deleteUserOK  %+v", 200, o.Payload)
}

func (o *DeleteUserOK) String() string {
	return fmt.Sprintf("[DELETE /api/internal/projects/{project_id}/clusters/{cluster_id}/branches/{branch_name}/users/{username}][%d] deleteUserOK  %+v", 200, o.Payload)
}

func (o *DeleteUserOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserDefault creates a DeleteUserDefault with default headers values
func NewDeleteUserDefault(code int) *DeleteUserDefault {
	return &DeleteUserDefault{
		_statusCode: code,
	}
}

/*
DeleteUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteUserDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this delete user default response has a 2xx status code
func (o *DeleteUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete user default response has a 3xx status code
func (o *DeleteUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete user default response has a 4xx status code
func (o *DeleteUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete user default response has a 5xx status code
func (o *DeleteUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete user default response a status code equal to that given
func (o *DeleteUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete user default response
func (o *DeleteUserDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/internal/projects/{project_id}/clusters/{cluster_id}/branches/{branch_name}/users/{username}][%d] DeleteUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteUserDefault) String() string {
	return fmt.Sprintf("[DELETE /api/internal/projects/{project_id}/clusters/{cluster_id}/branches/{branch_name}/users/{username}][%d] DeleteUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteUserDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *DeleteUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
