// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"tidbcloud-cli/pkg/tidbcloud/v1beta1/iam/models"
)

// PostV1beta1SqluserReader is a Reader for the PostV1beta1Sqluser structure.
type PostV1beta1SqluserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1beta1SqluserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1beta1SqluserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1beta1SqluserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1beta1/sqluser] PostV1beta1Sqluser", response, response.Code())
	}
}

// NewPostV1beta1SqluserOK creates a PostV1beta1SqluserOK with default headers values
func NewPostV1beta1SqluserOK() *PostV1beta1SqluserOK {
	return &PostV1beta1SqluserOK{}
}

/*
PostV1beta1SqluserOK describes a response with status code 200, with default header values.

OK
*/
type PostV1beta1SqluserOK struct {
	Payload *models.CentralSQLUser
}

// IsSuccess returns true when this post v1beta1 sqluser o k response has a 2xx status code
func (o *PostV1beta1SqluserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1beta1 sqluser o k response has a 3xx status code
func (o *PostV1beta1SqluserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1beta1 sqluser o k response has a 4xx status code
func (o *PostV1beta1SqluserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1beta1 sqluser o k response has a 5xx status code
func (o *PostV1beta1SqluserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1beta1 sqluser o k response a status code equal to that given
func (o *PostV1beta1SqluserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1beta1 sqluser o k response
func (o *PostV1beta1SqluserOK) Code() int {
	return 200
}

func (o *PostV1beta1SqluserOK) Error() string {
	return fmt.Sprintf("[POST /v1beta1/sqluser][%d] postV1beta1SqluserOK  %+v", 200, o.Payload)
}

func (o *PostV1beta1SqluserOK) String() string {
	return fmt.Sprintf("[POST /v1beta1/sqluser][%d] postV1beta1SqluserOK  %+v", 200, o.Payload)
}

func (o *PostV1beta1SqluserOK) GetPayload() *models.CentralSQLUser {
	return o.Payload
}

func (o *PostV1beta1SqluserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CentralSQLUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1beta1SqluserBadRequest creates a PostV1beta1SqluserBadRequest with default headers values
func NewPostV1beta1SqluserBadRequest() *PostV1beta1SqluserBadRequest {
	return &PostV1beta1SqluserBadRequest{}
}

/*
PostV1beta1SqluserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1beta1SqluserBadRequest struct {
	Payload *models.APIOpenAPIError
}

// IsSuccess returns true when this post v1beta1 sqluser bad request response has a 2xx status code
func (o *PostV1beta1SqluserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1beta1 sqluser bad request response has a 3xx status code
func (o *PostV1beta1SqluserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1beta1 sqluser bad request response has a 4xx status code
func (o *PostV1beta1SqluserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1beta1 sqluser bad request response has a 5xx status code
func (o *PostV1beta1SqluserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1beta1 sqluser bad request response a status code equal to that given
func (o *PostV1beta1SqluserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post v1beta1 sqluser bad request response
func (o *PostV1beta1SqluserBadRequest) Code() int {
	return 400
}

func (o *PostV1beta1SqluserBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1beta1/sqluser][%d] postV1beta1SqluserBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1beta1SqluserBadRequest) String() string {
	return fmt.Sprintf("[POST /v1beta1/sqluser][%d] postV1beta1SqluserBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1beta1SqluserBadRequest) GetPayload() *models.APIOpenAPIError {
	return o.Payload
}

func (o *PostV1beta1SqluserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIOpenAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
