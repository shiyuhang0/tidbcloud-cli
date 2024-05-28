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

// ServerlessServiceCreateClusterReader is a Reader for the ServerlessServiceCreateCluster structure.
type ServerlessServiceCreateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServerlessServiceCreateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServerlessServiceCreateClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServerlessServiceCreateClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServerlessServiceCreateClusterOK creates a ServerlessServiceCreateClusterOK with default headers values
func NewServerlessServiceCreateClusterOK() *ServerlessServiceCreateClusterOK {
	return &ServerlessServiceCreateClusterOK{}
}

/*
ServerlessServiceCreateClusterOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServerlessServiceCreateClusterOK struct {
	Payload *models.TidbCloudOpenApiserverlessv1beta1Cluster
}

// IsSuccess returns true when this serverless service create cluster o k response has a 2xx status code
func (o *ServerlessServiceCreateClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this serverless service create cluster o k response has a 3xx status code
func (o *ServerlessServiceCreateClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this serverless service create cluster o k response has a 4xx status code
func (o *ServerlessServiceCreateClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this serverless service create cluster o k response has a 5xx status code
func (o *ServerlessServiceCreateClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this serverless service create cluster o k response a status code equal to that given
func (o *ServerlessServiceCreateClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the serverless service create cluster o k response
func (o *ServerlessServiceCreateClusterOK) Code() int {
	return 200
}

func (o *ServerlessServiceCreateClusterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /clusters][%d] serverlessServiceCreateClusterOK %s", 200, payload)
}

func (o *ServerlessServiceCreateClusterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /clusters][%d] serverlessServiceCreateClusterOK %s", 200, payload)
}

func (o *ServerlessServiceCreateClusterOK) GetPayload() *models.TidbCloudOpenApiserverlessv1beta1Cluster {
	return o.Payload
}

func (o *ServerlessServiceCreateClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TidbCloudOpenApiserverlessv1beta1Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServerlessServiceCreateClusterDefault creates a ServerlessServiceCreateClusterDefault with default headers values
func NewServerlessServiceCreateClusterDefault(code int) *ServerlessServiceCreateClusterDefault {
	return &ServerlessServiceCreateClusterDefault{
		_statusCode: code,
	}
}

/*
ServerlessServiceCreateClusterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServerlessServiceCreateClusterDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this serverless service create cluster default response has a 2xx status code
func (o *ServerlessServiceCreateClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this serverless service create cluster default response has a 3xx status code
func (o *ServerlessServiceCreateClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this serverless service create cluster default response has a 4xx status code
func (o *ServerlessServiceCreateClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this serverless service create cluster default response has a 5xx status code
func (o *ServerlessServiceCreateClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this serverless service create cluster default response a status code equal to that given
func (o *ServerlessServiceCreateClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the serverless service create cluster default response
func (o *ServerlessServiceCreateClusterDefault) Code() int {
	return o._statusCode
}

func (o *ServerlessServiceCreateClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /clusters][%d] ServerlessService_CreateCluster default %s", o._statusCode, payload)
}

func (o *ServerlessServiceCreateClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /clusters][%d] ServerlessService_CreateCluster default %s", o._statusCode, payload)
}

func (o *ServerlessServiceCreateClusterDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ServerlessServiceCreateClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
