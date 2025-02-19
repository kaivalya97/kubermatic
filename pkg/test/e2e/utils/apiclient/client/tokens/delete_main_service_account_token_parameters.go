// Code generated by go-swagger; DO NOT EDIT.

package tokens

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
)

// NewDeleteMainServiceAccountTokenParams creates a new DeleteMainServiceAccountTokenParams object
// with the default values initialized.
func NewDeleteMainServiceAccountTokenParams() *DeleteMainServiceAccountTokenParams {
	var ()
	return &DeleteMainServiceAccountTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMainServiceAccountTokenParamsWithTimeout creates a new DeleteMainServiceAccountTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMainServiceAccountTokenParamsWithTimeout(timeout time.Duration) *DeleteMainServiceAccountTokenParams {
	var ()
	return &DeleteMainServiceAccountTokenParams{

		timeout: timeout,
	}
}

// NewDeleteMainServiceAccountTokenParamsWithContext creates a new DeleteMainServiceAccountTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMainServiceAccountTokenParamsWithContext(ctx context.Context) *DeleteMainServiceAccountTokenParams {
	var ()
	return &DeleteMainServiceAccountTokenParams{

		Context: ctx,
	}
}

// NewDeleteMainServiceAccountTokenParamsWithHTTPClient creates a new DeleteMainServiceAccountTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMainServiceAccountTokenParamsWithHTTPClient(client *http.Client) *DeleteMainServiceAccountTokenParams {
	var ()
	return &DeleteMainServiceAccountTokenParams{
		HTTPClient: client,
	}
}

/*DeleteMainServiceAccountTokenParams contains all the parameters to send to the API endpoint
for the delete main service account token operation typically these are written to a http.Request
*/
type DeleteMainServiceAccountTokenParams struct {

	/*ServiceaccountID*/
	ServiceAccountID string
	/*TokenID*/
	TokenID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete main service account token params
func (o *DeleteMainServiceAccountTokenParams) WithTimeout(timeout time.Duration) *DeleteMainServiceAccountTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete main service account token params
func (o *DeleteMainServiceAccountTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete main service account token params
func (o *DeleteMainServiceAccountTokenParams) WithContext(ctx context.Context) *DeleteMainServiceAccountTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete main service account token params
func (o *DeleteMainServiceAccountTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete main service account token params
func (o *DeleteMainServiceAccountTokenParams) WithHTTPClient(client *http.Client) *DeleteMainServiceAccountTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete main service account token params
func (o *DeleteMainServiceAccountTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceAccountID adds the serviceaccountID to the delete main service account token params
func (o *DeleteMainServiceAccountTokenParams) WithServiceAccountID(serviceaccountID string) *DeleteMainServiceAccountTokenParams {
	o.SetServiceAccountID(serviceaccountID)
	return o
}

// SetServiceAccountID adds the serviceaccountId to the delete main service account token params
func (o *DeleteMainServiceAccountTokenParams) SetServiceAccountID(serviceaccountID string) {
	o.ServiceAccountID = serviceaccountID
}

// WithTokenID adds the tokenID to the delete main service account token params
func (o *DeleteMainServiceAccountTokenParams) WithTokenID(tokenID string) *DeleteMainServiceAccountTokenParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the delete main service account token params
func (o *DeleteMainServiceAccountTokenParams) SetTokenID(tokenID string) {
	o.TokenID = tokenID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMainServiceAccountTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceaccount_id
	if err := r.SetPathParam("serviceaccount_id", o.ServiceAccountID); err != nil {
		return err
	}

	// path param token_id
	if err := r.SetPathParam("token_id", o.TokenID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
