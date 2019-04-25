// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDeviceGroupClusterAlertConfListParams creates a new GetDeviceGroupClusterAlertConfListParams object
// with the default values initialized.
func NewGetDeviceGroupClusterAlertConfListParams() *GetDeviceGroupClusterAlertConfListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceGroupClusterAlertConfListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceGroupClusterAlertConfListParamsWithTimeout creates a new GetDeviceGroupClusterAlertConfListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceGroupClusterAlertConfListParamsWithTimeout(timeout time.Duration) *GetDeviceGroupClusterAlertConfListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceGroupClusterAlertConfListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: timeout,
	}
}

// NewGetDeviceGroupClusterAlertConfListParamsWithContext creates a new GetDeviceGroupClusterAlertConfListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceGroupClusterAlertConfListParamsWithContext(ctx context.Context) *GetDeviceGroupClusterAlertConfListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceGroupClusterAlertConfListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		Context: ctx,
	}
}

// NewGetDeviceGroupClusterAlertConfListParamsWithHTTPClient creates a new GetDeviceGroupClusterAlertConfListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceGroupClusterAlertConfListParamsWithHTTPClient(client *http.Client) *GetDeviceGroupClusterAlertConfListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceGroupClusterAlertConfListParams{
		Offset:     &offsetDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*GetDeviceGroupClusterAlertConfListParams contains all the parameters to send to the API endpoint
for the get device group cluster alert conf list operation typically these are written to a http.Request
*/
type GetDeviceGroupClusterAlertConfListParams struct {

	/*DeviceGroupID*/
	DeviceGroupID int32
	/*Fields*/
	Fields *string
	/*Filter*/
	Filter *string
	/*Offset*/
	Offset *int32
	/*Size*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) WithTimeout(timeout time.Duration) *GetDeviceGroupClusterAlertConfListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) WithContext(ctx context.Context) *GetDeviceGroupClusterAlertConfListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) WithHTTPClient(client *http.Client) *GetDeviceGroupClusterAlertConfListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceGroupID adds the deviceGroupID to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) WithDeviceGroupID(deviceGroupID int32) *GetDeviceGroupClusterAlertConfListParams {
	o.SetDeviceGroupID(deviceGroupID)
	return o
}

// SetDeviceGroupID adds the deviceGroupId to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) SetDeviceGroupID(deviceGroupID int32) {
	o.DeviceGroupID = deviceGroupID
}

// WithFields adds the fields to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) WithFields(fields *string) *GetDeviceGroupClusterAlertConfListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) WithFilter(filter *string) *GetDeviceGroupClusterAlertConfListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) WithOffset(offset *int32) *GetDeviceGroupClusterAlertConfListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) WithSize(size *int32) *GetDeviceGroupClusterAlertConfListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get device group cluster alert conf list params
func (o *GetDeviceGroupClusterAlertConfListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceGroupClusterAlertConfListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceGroupId
	if err := r.SetPathParam("deviceGroupId", swag.FormatInt32(o.DeviceGroupID)); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int32
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}