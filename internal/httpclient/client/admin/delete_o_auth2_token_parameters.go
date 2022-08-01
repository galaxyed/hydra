// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewDeleteOAuth2TokenParams creates a new DeleteOAuth2TokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOAuth2TokenParams() *DeleteOAuth2TokenParams {
	return &DeleteOAuth2TokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOAuth2TokenParamsWithTimeout creates a new DeleteOAuth2TokenParams object
// with the ability to set a timeout on a request.
func NewDeleteOAuth2TokenParamsWithTimeout(timeout time.Duration) *DeleteOAuth2TokenParams {
	return &DeleteOAuth2TokenParams{
		timeout: timeout,
	}
}

// NewDeleteOAuth2TokenParamsWithContext creates a new DeleteOAuth2TokenParams object
// with the ability to set a context for a request.
func NewDeleteOAuth2TokenParamsWithContext(ctx context.Context) *DeleteOAuth2TokenParams {
	return &DeleteOAuth2TokenParams{
		Context: ctx,
	}
}

// NewDeleteOAuth2TokenParamsWithHTTPClient creates a new DeleteOAuth2TokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOAuth2TokenParamsWithHTTPClient(client *http.Client) *DeleteOAuth2TokenParams {
	return &DeleteOAuth2TokenParams{
		HTTPClient: client,
	}
}

/* DeleteOAuth2TokenParams contains all the parameters to send to the API endpoint
   for the delete o auth2 token operation.

   Typically these are written to a http.Request.
*/
type DeleteOAuth2TokenParams struct {

	// ClientID.
	ClientID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete o auth2 token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOAuth2TokenParams) WithDefaults() *DeleteOAuth2TokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete o auth2 token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOAuth2TokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete o auth2 token params
func (o *DeleteOAuth2TokenParams) WithTimeout(timeout time.Duration) *DeleteOAuth2TokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete o auth2 token params
func (o *DeleteOAuth2TokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete o auth2 token params
func (o *DeleteOAuth2TokenParams) WithContext(ctx context.Context) *DeleteOAuth2TokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete o auth2 token params
func (o *DeleteOAuth2TokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete o auth2 token params
func (o *DeleteOAuth2TokenParams) WithHTTPClient(client *http.Client) *DeleteOAuth2TokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete o auth2 token params
func (o *DeleteOAuth2TokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the delete o auth2 token params
func (o *DeleteOAuth2TokenParams) WithClientID(clientID string) *DeleteOAuth2TokenParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the delete o auth2 token params
func (o *DeleteOAuth2TokenParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOAuth2TokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param client_id
	qrClientID := o.ClientID
	qClientID := qrClientID
	if qClientID != "" {

		if err := r.SetQueryParam("client_id", qClientID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
