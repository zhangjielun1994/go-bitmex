// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewOrderCancelAllAfterParams creates a new OrderCancelAllAfterParams object
// with the default values initialized.
func NewOrderCancelAllAfterParams() *OrderCancelAllAfterParams {
	var ()
	return &OrderCancelAllAfterParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewOrderCancelAllAfterParamsWithTimeout creates a new OrderCancelAllAfterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrderCancelAllAfterParamsWithTimeout(timeout time.Duration) *OrderCancelAllAfterParams {
	var ()
	return &OrderCancelAllAfterParams{

		requestTimeout: timeout,
	}
}

// NewOrderCancelAllAfterParamsWithContext creates a new OrderCancelAllAfterParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrderCancelAllAfterParamsWithContext(ctx context.Context) *OrderCancelAllAfterParams {
	var ()
	return &OrderCancelAllAfterParams{

		Context: ctx,
	}
}

// NewOrderCancelAllAfterParamsWithHTTPClient creates a new OrderCancelAllAfterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrderCancelAllAfterParamsWithHTTPClient(client *http.Client) *OrderCancelAllAfterParams {
	var ()
	return &OrderCancelAllAfterParams{
		HTTPClient: client,
	}
}

/*OrderCancelAllAfterParams contains all the parameters to send to the API endpoint
for the order cancel all after operation typically these are written to a http.Request
*/
type OrderCancelAllAfterParams struct {

	/*Timeout
	  Timeout in ms. Set to 0 to cancel this timer.

	*/
	Timeout float64

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the order cancel all after params
func (o *OrderCancelAllAfterParams) WithRequestTimeout(timeout time.Duration) *OrderCancelAllAfterParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the order cancel all after params
func (o *OrderCancelAllAfterParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the order cancel all after params
func (o *OrderCancelAllAfterParams) WithContext(ctx context.Context) *OrderCancelAllAfterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order cancel all after params
func (o *OrderCancelAllAfterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order cancel all after params
func (o *OrderCancelAllAfterParams) WithHTTPClient(client *http.Client) *OrderCancelAllAfterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order cancel all after params
func (o *OrderCancelAllAfterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimeout adds the timeout to the order cancel all after params
func (o *OrderCancelAllAfterParams) WithTimeout(timeout float64) *OrderCancelAllAfterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order cancel all after params
func (o *OrderCancelAllAfterParams) SetTimeout(timeout float64) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *OrderCancelAllAfterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	// form param timeout
	frTimeout := o.Timeout
	fTimeout := swag.FormatFloat64(frTimeout)
	if fTimeout != "" {
		if err := r.SetFormParam("timeout", fTimeout); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
