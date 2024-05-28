// Code generated by go-swagger; DO NOT EDIT.

package serverless_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/models"
)

// ServerlessServiceListClustersReader is a Reader for the ServerlessServiceListClusters structure.
type ServerlessServiceListClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServerlessServiceListClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServerlessServiceListClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServerlessServiceListClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServerlessServiceListClustersOK creates a ServerlessServiceListClustersOK with default headers values
func NewServerlessServiceListClustersOK() *ServerlessServiceListClustersOK {
	return &ServerlessServiceListClustersOK{}
}

/*
ServerlessServiceListClustersOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServerlessServiceListClustersOK struct {
	Payload *models.TidbCloudOpenApiserverlessv1beta1ListClustersResponse
}

// IsSuccess returns true when this serverless service list clusters o k response has a 2xx status code
func (o *ServerlessServiceListClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this serverless service list clusters o k response has a 3xx status code
func (o *ServerlessServiceListClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this serverless service list clusters o k response has a 4xx status code
func (o *ServerlessServiceListClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this serverless service list clusters o k response has a 5xx status code
func (o *ServerlessServiceListClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this serverless service list clusters o k response a status code equal to that given
func (o *ServerlessServiceListClustersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the serverless service list clusters o k response
func (o *ServerlessServiceListClustersOK) Code() int {
	return 200
}

func (o *ServerlessServiceListClustersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters][%d] serverlessServiceListClustersOK %s", 200, payload)
}

func (o *ServerlessServiceListClustersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters][%d] serverlessServiceListClustersOK %s", 200, payload)
}

func (o *ServerlessServiceListClustersOK) GetPayload() *models.TidbCloudOpenApiserverlessv1beta1ListClustersResponse {
	return o.Payload
}

func (o *ServerlessServiceListClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TidbCloudOpenApiserverlessv1beta1ListClustersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServerlessServiceListClustersDefault creates a ServerlessServiceListClustersDefault with default headers values
func NewServerlessServiceListClustersDefault(code int) *ServerlessServiceListClustersDefault {
	return &ServerlessServiceListClustersDefault{
		_statusCode: code,
	}
}

/*
ServerlessServiceListClustersDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServerlessServiceListClustersDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this serverless service list clusters default response has a 2xx status code
func (o *ServerlessServiceListClustersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this serverless service list clusters default response has a 3xx status code
func (o *ServerlessServiceListClustersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this serverless service list clusters default response has a 4xx status code
func (o *ServerlessServiceListClustersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this serverless service list clusters default response has a 5xx status code
func (o *ServerlessServiceListClustersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this serverless service list clusters default response a status code equal to that given
func (o *ServerlessServiceListClustersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the serverless service list clusters default response
func (o *ServerlessServiceListClustersDefault) Code() int {
	return o._statusCode
}

func (o *ServerlessServiceListClustersDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters][%d] ServerlessService_ListClusters default %s", o._statusCode, payload)
}

func (o *ServerlessServiceListClustersDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters][%d] ServerlessService_ListClusters default %s", o._statusCode, payload)
}

func (o *ServerlessServiceListClustersDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ServerlessServiceListClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
