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

// NewGetOidcUserInfoParams creates a new GetOidcUserInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOidcUserInfoParams() *GetOidcUserInfoParams {
	return &GetOidcUserInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOidcUserInfoParamsWithTimeout creates a new GetOidcUserInfoParams object
// with the ability to set a timeout on a request.
func NewGetOidcUserInfoParamsWithTimeout(timeout time.Duration) *GetOidcUserInfoParams {
	return &GetOidcUserInfoParams{
		timeout: timeout,
	}
}

// NewGetOidcUserInfoParamsWithContext creates a new GetOidcUserInfoParams object
// with the ability to set a context for a request.
func NewGetOidcUserInfoParamsWithContext(ctx context.Context) *GetOidcUserInfoParams {
	return &GetOidcUserInfoParams{
		Context: ctx,
	}
}

// NewGetOidcUserInfoParamsWithHTTPClient creates a new GetOidcUserInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOidcUserInfoParamsWithHTTPClient(client *http.Client) *GetOidcUserInfoParams {
	return &GetOidcUserInfoParams{
		HTTPClient: client,
	}
}

/* GetOidcUserInfoParams contains all the parameters to send to the API endpoint
   for the get oidc user info operation.

   Typically these are written to a http.Request.
*/
type GetOidcUserInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get oidc user info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOidcUserInfoParams) WithDefaults() *GetOidcUserInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get oidc user info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOidcUserInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get oidc user info params
func (o *GetOidcUserInfoParams) WithTimeout(timeout time.Duration) *GetOidcUserInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oidc user info params
func (o *GetOidcUserInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oidc user info params
func (o *GetOidcUserInfoParams) WithContext(ctx context.Context) *GetOidcUserInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oidc user info params
func (o *GetOidcUserInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get oidc user info params
func (o *GetOidcUserInfoParams) WithHTTPClient(client *http.Client) *GetOidcUserInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get oidc user info params
func (o *GetOidcUserInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOidcUserInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
