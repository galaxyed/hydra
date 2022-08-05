// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// DeleteTrustedJwtGrantIssuerReader is a Reader for the DeleteTrustedJwtGrantIssuer structure.
type DeleteTrustedJwtGrantIssuerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTrustedJwtGrantIssuerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteTrustedJwtGrantIssuerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTrustedJwtGrantIssuerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTrustedJwtGrantIssuerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTrustedJwtGrantIssuerNoContent creates a DeleteTrustedJwtGrantIssuerNoContent with default headers values
func NewDeleteTrustedJwtGrantIssuerNoContent() *DeleteTrustedJwtGrantIssuerNoContent {
	return &DeleteTrustedJwtGrantIssuerNoContent{}
}

/* DeleteTrustedJwtGrantIssuerNoContent describes a response with status code 204, with default header values.

 Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type DeleteTrustedJwtGrantIssuerNoContent struct {
}

func (o *DeleteTrustedJwtGrantIssuerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /trust/grants/jwt-bearer/issuers/{id}][%d] deleteTrustedJwtGrantIssuerNoContent ", 204)
}

func (o *DeleteTrustedJwtGrantIssuerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTrustedJwtGrantIssuerNotFound creates a DeleteTrustedJwtGrantIssuerNotFound with default headers values
func NewDeleteTrustedJwtGrantIssuerNotFound() *DeleteTrustedJwtGrantIssuerNotFound {
	return &DeleteTrustedJwtGrantIssuerNotFound{}
}

/* DeleteTrustedJwtGrantIssuerNotFound describes a response with status code 404, with default header values.

genericError
*/
type DeleteTrustedJwtGrantIssuerNotFound struct {
	Payload *models.GenericError
}

func (o *DeleteTrustedJwtGrantIssuerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /trust/grants/jwt-bearer/issuers/{id}][%d] deleteTrustedJwtGrantIssuerNotFound  %+v", 404, o.Payload)
}
func (o *DeleteTrustedJwtGrantIssuerNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteTrustedJwtGrantIssuerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTrustedJwtGrantIssuerInternalServerError creates a DeleteTrustedJwtGrantIssuerInternalServerError with default headers values
func NewDeleteTrustedJwtGrantIssuerInternalServerError() *DeleteTrustedJwtGrantIssuerInternalServerError {
	return &DeleteTrustedJwtGrantIssuerInternalServerError{}
}

/* DeleteTrustedJwtGrantIssuerInternalServerError describes a response with status code 500, with default header values.

genericError
*/
type DeleteTrustedJwtGrantIssuerInternalServerError struct {
	Payload *models.GenericError
}

func (o *DeleteTrustedJwtGrantIssuerInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /trust/grants/jwt-bearer/issuers/{id}][%d] deleteTrustedJwtGrantIssuerInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteTrustedJwtGrantIssuerInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteTrustedJwtGrantIssuerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}