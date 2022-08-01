// Code generated by go-swagger; DO NOT EDIT.

package v1

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

// AdminDeleteJSONWebKeyReader is a Reader for the AdminDeleteJSONWebKey structure.
type AdminDeleteJSONWebKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminDeleteJSONWebKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAdminDeleteJSONWebKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAdminDeleteJSONWebKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAdminDeleteJSONWebKeyNoContent creates a AdminDeleteJSONWebKeyNoContent with default headers values
func NewAdminDeleteJSONWebKeyNoContent() *AdminDeleteJSONWebKeyNoContent {
	return &AdminDeleteJSONWebKeyNoContent{}
}

/* AdminDeleteJSONWebKeyNoContent describes a response with status code 204, with default header values.

 Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type AdminDeleteJSONWebKeyNoContent struct {
}

func (o *AdminDeleteJSONWebKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /admin/keys/{set}/{kid}][%d] adminDeleteJsonWebKeyNoContent ", 204)
}

func (o *AdminDeleteJSONWebKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeleteJSONWebKeyDefault creates a AdminDeleteJSONWebKeyDefault with default headers values
func NewAdminDeleteJSONWebKeyDefault(code int) *AdminDeleteJSONWebKeyDefault {
	return &AdminDeleteJSONWebKeyDefault{
		_statusCode: code,
	}
}

/* AdminDeleteJSONWebKeyDefault describes a response with status code -1, with default header values.

oAuth2ApiError
*/
type AdminDeleteJSONWebKeyDefault struct {
	_statusCode int

	Payload *models.OAuth2APIError
}

// Code gets the status code for the admin delete Json web key default response
func (o *AdminDeleteJSONWebKeyDefault) Code() int {
	return o._statusCode
}

func (o *AdminDeleteJSONWebKeyDefault) Error() string {
	return fmt.Sprintf("[DELETE /admin/keys/{set}/{kid}][%d] adminDeleteJsonWebKey default  %+v", o._statusCode, o.Payload)
}
func (o *AdminDeleteJSONWebKeyDefault) GetPayload() *models.OAuth2APIError {
	return o.Payload
}

func (o *AdminDeleteJSONWebKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
