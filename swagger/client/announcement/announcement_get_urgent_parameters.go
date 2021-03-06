// Code generated by go-swagger; DO NOT EDIT.

package announcement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAnnouncementGetUrgentParams creates a new AnnouncementGetUrgentParams object
// with the default values initialized.
func NewAnnouncementGetUrgentParams() *AnnouncementGetUrgentParams {

	return &AnnouncementGetUrgentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAnnouncementGetUrgentParamsWithTimeout creates a new AnnouncementGetUrgentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAnnouncementGetUrgentParamsWithTimeout(timeout time.Duration) *AnnouncementGetUrgentParams {

	return &AnnouncementGetUrgentParams{

		timeout: timeout,
	}
}

// NewAnnouncementGetUrgentParamsWithContext creates a new AnnouncementGetUrgentParams object
// with the default values initialized, and the ability to set a context for a request
func NewAnnouncementGetUrgentParamsWithContext(ctx context.Context) *AnnouncementGetUrgentParams {

	return &AnnouncementGetUrgentParams{

		Context: ctx,
	}
}

// NewAnnouncementGetUrgentParamsWithHTTPClient creates a new AnnouncementGetUrgentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAnnouncementGetUrgentParamsWithHTTPClient(client *http.Client) *AnnouncementGetUrgentParams {

	return &AnnouncementGetUrgentParams{
		HTTPClient: client,
	}
}

/*AnnouncementGetUrgentParams contains all the parameters to send to the API endpoint
for the announcement get urgent operation typically these are written to a http.Request
*/
type AnnouncementGetUrgentParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the announcement get urgent params
func (o *AnnouncementGetUrgentParams) WithTimeout(timeout time.Duration) *AnnouncementGetUrgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the announcement get urgent params
func (o *AnnouncementGetUrgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the announcement get urgent params
func (o *AnnouncementGetUrgentParams) WithContext(ctx context.Context) *AnnouncementGetUrgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the announcement get urgent params
func (o *AnnouncementGetUrgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the announcement get urgent params
func (o *AnnouncementGetUrgentParams) WithHTTPClient(client *http.Client) *AnnouncementGetUrgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the announcement get urgent params
func (o *AnnouncementGetUrgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AnnouncementGetUrgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
