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

	models "github.com/logicmonitor/lm-sdk-go/models"
)

// NewPatchDeviceGroupClusterAlertConfByIDParams creates a new PatchDeviceGroupClusterAlertConfByIDParams object
// with the default values initialized.
func NewPatchDeviceGroupClusterAlertConfByIDParams() *PatchDeviceGroupClusterAlertConfByIDParams {
	var ()
	return &PatchDeviceGroupClusterAlertConfByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDeviceGroupClusterAlertConfByIDParamsWithTimeout creates a new PatchDeviceGroupClusterAlertConfByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchDeviceGroupClusterAlertConfByIDParamsWithTimeout(timeout time.Duration) *PatchDeviceGroupClusterAlertConfByIDParams {
	var ()
	return &PatchDeviceGroupClusterAlertConfByIDParams{

		timeout: timeout,
	}
}

// NewPatchDeviceGroupClusterAlertConfByIDParamsWithContext creates a new PatchDeviceGroupClusterAlertConfByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchDeviceGroupClusterAlertConfByIDParamsWithContext(ctx context.Context) *PatchDeviceGroupClusterAlertConfByIDParams {
	var ()
	return &PatchDeviceGroupClusterAlertConfByIDParams{

		Context: ctx,
	}
}

// NewPatchDeviceGroupClusterAlertConfByIDParamsWithHTTPClient creates a new PatchDeviceGroupClusterAlertConfByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchDeviceGroupClusterAlertConfByIDParamsWithHTTPClient(client *http.Client) *PatchDeviceGroupClusterAlertConfByIDParams {
	var ()
	return &PatchDeviceGroupClusterAlertConfByIDParams{
		HTTPClient: client,
	}
}

/*PatchDeviceGroupClusterAlertConfByIDParams contains all the parameters to send to the API endpoint
for the patch device group cluster alert conf by Id operation typically these are written to a http.Request
*/
type PatchDeviceGroupClusterAlertConfByIDParams struct {

	/*Body*/
	Body *models.DeviceClusterAlertConfig
	/*DeviceGroupID*/
	DeviceGroupID int32
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch device group cluster alert conf by Id params
func (o *PatchDeviceGroupClusterAlertConfByIDParams) WithTimeout(timeout time.Duration) *PatchDeviceGroupClusterAlertConfByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch device group cluster alert conf by Id params
func (o *PatchDeviceGroupClusterAlertConfByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch device group cluster alert conf by Id params
func (o *PatchDeviceGroupClusterAlertConfByIDParams) WithContext(ctx context.Context) *PatchDeviceGroupClusterAlertConfByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch device group cluster alert conf by Id params
func (o *PatchDeviceGroupClusterAlertConfByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch device group cluster alert conf by Id params
func (o *PatchDeviceGroupClusterAlertConfByIDParams) WithHTTPClient(client *http.Client) *PatchDeviceGroupClusterAlertConfByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch device group cluster alert conf by Id params
func (o *PatchDeviceGroupClusterAlertConfByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch device group cluster alert conf by Id params
func (o *PatchDeviceGroupClusterAlertConfByIDParams) WithBody(body *models.DeviceClusterAlertConfig) *PatchDeviceGroupClusterAlertConfByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch device group cluster alert conf by Id params
func (o *PatchDeviceGroupClusterAlertConfByIDParams) SetBody(body *models.DeviceClusterAlertConfig) {
	o.Body = body
}

// WithDeviceGroupID adds the deviceGroupID to the patch device group cluster alert conf by Id params
func (o *PatchDeviceGroupClusterAlertConfByIDParams) WithDeviceGroupID(deviceGroupID int32) *PatchDeviceGroupClusterAlertConfByIDParams {
	o.SetDeviceGroupID(deviceGroupID)
	return o
}

// SetDeviceGroupID adds the deviceGroupId to the patch device group cluster alert conf by Id params
func (o *PatchDeviceGroupClusterAlertConfByIDParams) SetDeviceGroupID(deviceGroupID int32) {
	o.DeviceGroupID = deviceGroupID
}

// WithID adds the id to the patch device group cluster alert conf by Id params
func (o *PatchDeviceGroupClusterAlertConfByIDParams) WithID(id int32) *PatchDeviceGroupClusterAlertConfByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch device group cluster alert conf by Id params
func (o *PatchDeviceGroupClusterAlertConfByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDeviceGroupClusterAlertConfByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deviceGroupId
	if err := r.SetPathParam("deviceGroupId", swag.FormatInt32(o.DeviceGroupID)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}