// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dcim

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

	"github.com/digitalocean/go-netbox/netbox/models"
)

// NewDcimDeviceBayTemplatesUpdateParams creates a new DcimDeviceBayTemplatesUpdateParams object
// with the default values initialized.
func NewDcimDeviceBayTemplatesUpdateParams() *DcimDeviceBayTemplatesUpdateParams {
	var ()
	return &DcimDeviceBayTemplatesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBayTemplatesUpdateParamsWithTimeout creates a new DcimDeviceBayTemplatesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceBayTemplatesUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesUpdateParams {
	var ()
	return &DcimDeviceBayTemplatesUpdateParams{

		timeout: timeout,
	}
}

// NewDcimDeviceBayTemplatesUpdateParamsWithContext creates a new DcimDeviceBayTemplatesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceBayTemplatesUpdateParamsWithContext(ctx context.Context) *DcimDeviceBayTemplatesUpdateParams {
	var ()
	return &DcimDeviceBayTemplatesUpdateParams{

		Context: ctx,
	}
}

// NewDcimDeviceBayTemplatesUpdateParamsWithHTTPClient creates a new DcimDeviceBayTemplatesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceBayTemplatesUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesUpdateParams {
	var ()
	return &DcimDeviceBayTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/*DcimDeviceBayTemplatesUpdateParams contains all the parameters to send to the API endpoint
for the dcim device bay templates update operation typically these are written to a http.Request
*/
type DcimDeviceBayTemplatesUpdateParams struct {

	/*Data*/
	Data *models.WritableDeviceBayTemplate
	/*ID
	  A unique integer value identifying this device bay template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) WithContext(ctx context.Context) *DcimDeviceBayTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) WithData(data *models.WritableDeviceBayTemplate) *DcimDeviceBayTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) SetData(data *models.WritableDeviceBayTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) WithID(id int64) *DcimDeviceBayTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bay templates update params
func (o *DcimDeviceBayTemplatesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBayTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
