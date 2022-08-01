// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/client/v1"
)

// NewAdminDeleteOAuth2ClientParams creates a new AdminDeleteOAuth2ClientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminDeleteOAuth2ClientParams() *AdminDeleteOAuth2ClientParams {
	return &AdminDeleteOAuth2ClientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminDeleteOAuth2ClientParamsWithTimeout creates a new AdminDeleteOAuth2ClientParams object
// with the ability to set a timeout on a request.
func NewAdminDeleteOAuth2ClientParamsWithTimeout(timeout time.Duration) *AdminDeleteOAuth2ClientParams {
	return &AdminDeleteOAuth2ClientParams{
		timeout: timeout,
	}
}

// NewAdminDeleteOAuth2ClientParamsWithContext creates a new AdminDeleteOAuth2ClientParams object
// with the ability to set a context for a request.
func NewAdminDeleteOAuth2ClientParamsWithContext(ctx context.Context) *AdminDeleteOAuth2ClientParams {
	return &AdminDeleteOAuth2ClientParams{
		Context: ctx,
	}
}

// NewAdminDeleteOAuth2ClientParamsWithHTTPClient creates a new AdminDeleteOAuth2ClientParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminDeleteOAuth2ClientParamsWithHTTPClient(client *http.Client) *AdminDeleteOAuth2ClientParams {
	return &AdminDeleteOAuth2ClientParams{
		HTTPClient: client,
	}
}

/* AdminDeleteOAuth2ClientParams contains all the parameters to send to the API endpoint
   for the admin delete o auth2 client operation.

   Typically these are written to a http.Request.
*/
type AdminDeleteOAuth2ClientParams struct {

	/* ID.

	   The id of the OAuth 2.0 Client.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin delete o auth2 client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminDeleteOAuth2ClientParams) WithDefaults() *AdminDeleteOAuth2ClientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin delete o auth2 client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminDeleteOAuth2ClientParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin delete o auth2 client params
func (o *AdminDeleteOAuth2ClientParams) WithTimeout(timeout time.Duration) *AdminDeleteOAuth2ClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin delete o auth2 client params
func (o *AdminDeleteOAuth2ClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin delete o auth2 client params
func (o *AdminDeleteOAuth2ClientParams) WithContext(ctx context.Context) *AdminDeleteOAuth2ClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin delete o auth2 client params
func (o *AdminDeleteOAuth2ClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin delete o auth2 client params
func (o *AdminDeleteOAuth2ClientParams) WithHTTPClient(client *http.Client) *AdminDeleteOAuth2ClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin delete o auth2 client params
func (o *AdminDeleteOAuth2ClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the admin delete o auth2 client params
func (o *AdminDeleteOAuth2ClientParams) WithID(id string) *AdminDeleteOAuth2ClientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the admin delete o auth2 client params
func (o *AdminDeleteOAuth2ClientParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AdminDeleteOAuth2ClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
