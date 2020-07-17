// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetPatientParams creates a new GetPatientParams object
// with the default values initialized.
func NewGetPatientParams() *GetPatientParams {
	var ()
	return &GetPatientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPatientParamsWithTimeout creates a new GetPatientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPatientParamsWithTimeout(timeout time.Duration) *GetPatientParams {
	var ()
	return &GetPatientParams{

		timeout: timeout,
	}
}

// NewGetPatientParamsWithContext creates a new GetPatientParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPatientParamsWithContext(ctx context.Context) *GetPatientParams {
	var ()
	return &GetPatientParams{

		Context: ctx,
	}
}

// NewGetPatientParamsWithHTTPClient creates a new GetPatientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPatientParamsWithHTTPClient(client *http.Client) *GetPatientParams {
	var ()
	return &GetPatientParams{
		HTTPClient: client,
	}
}

/*GetPatientParams contains all the parameters to send to the API endpoint
for the get patient operation typically these are written to a http.Request
*/
type GetPatientParams struct {

	/*ID
	  UUID of the patient

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get patient params
func (o *GetPatientParams) WithTimeout(timeout time.Duration) *GetPatientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get patient params
func (o *GetPatientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get patient params
func (o *GetPatientParams) WithContext(ctx context.Context) *GetPatientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get patient params
func (o *GetPatientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get patient params
func (o *GetPatientParams) WithHTTPClient(client *http.Client) *GetPatientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get patient params
func (o *GetPatientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get patient params
func (o *GetPatientParams) WithID(id string) *GetPatientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get patient params
func (o *GetPatientParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPatientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ID
	if err := r.SetPathParam("ID", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
