// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/client/v1"
	"github.com/ory/hydra/internal/httpclient/models"
)

// DiscoverOpenIDConfigurationReader is a Reader for the DiscoverOpenIDConfiguration structure.
type DiscoverOpenIDConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiscoverOpenIDConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiscoverOpenIDConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDiscoverOpenIDConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDiscoverOpenIDConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDiscoverOpenIDConfigurationOK creates a DiscoverOpenIDConfigurationOK with default headers values
func NewDiscoverOpenIDConfigurationOK() *DiscoverOpenIDConfigurationOK {
	return &DiscoverOpenIDConfigurationOK{}
}

/* DiscoverOpenIDConfigurationOK describes a response with status code 200, with default header values.

wellKnown
*/
type DiscoverOpenIDConfigurationOK struct {
	Payload *models.WellKnown
}

func (o *DiscoverOpenIDConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /.well-known/openid-configuration][%d] discoverOpenIdConfigurationOK  %+v", 200, o.Payload)
}
func (o *DiscoverOpenIDConfigurationOK) GetPayload() *models.WellKnown {
	return o.Payload
}

func (o *DiscoverOpenIDConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WellKnown)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoverOpenIDConfigurationUnauthorized creates a DiscoverOpenIDConfigurationUnauthorized with default headers values
func NewDiscoverOpenIDConfigurationUnauthorized() *DiscoverOpenIDConfigurationUnauthorized {
	return &DiscoverOpenIDConfigurationUnauthorized{}
}

/* DiscoverOpenIDConfigurationUnauthorized describes a response with status code 401, with default header values.

oAuth2ApiError
*/
type DiscoverOpenIDConfigurationUnauthorized struct {
	Payload *models.OAuth2APIError
}

func (o *DiscoverOpenIDConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /.well-known/openid-configuration][%d] discoverOpenIdConfigurationUnauthorized  %+v", 401, o.Payload)
}
func (o *DiscoverOpenIDConfigurationUnauthorized) GetPayload() *models.OAuth2APIError {
	return o.Payload
}

func (o *DiscoverOpenIDConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoverOpenIDConfigurationDefault creates a DiscoverOpenIDConfigurationDefault with default headers values
func NewDiscoverOpenIDConfigurationDefault(code int) *DiscoverOpenIDConfigurationDefault {
	return &DiscoverOpenIDConfigurationDefault{
		_statusCode: code,
	}
}

/* DiscoverOpenIDConfigurationDefault describes a response with status code -1, with default header values.

oAuth2ApiError
*/
type DiscoverOpenIDConfigurationDefault struct {
	_statusCode int

	Payload *models.OAuth2APIError
}

// Code gets the status code for the discover open ID configuration default response
func (o *DiscoverOpenIDConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *DiscoverOpenIDConfigurationDefault) Error() string {
	return fmt.Sprintf("[GET /.well-known/openid-configuration][%d] discoverOpenIDConfiguration default  %+v", o._statusCode, o.Payload)
}
func (o *DiscoverOpenIDConfigurationDefault) GetPayload() *models.OAuth2APIError {
	return o.Payload
}

func (o *DiscoverOpenIDConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
