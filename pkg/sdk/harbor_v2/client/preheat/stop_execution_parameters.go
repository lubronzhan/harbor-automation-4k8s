// Code generated by go-swagger; DO NOT EDIT.

package preheat

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
	"github.com/go-openapi/swag"

	"github.com/szlabs/harbor-automation-4k8s/pkg/sdk/harbor_v2/models"
)

// NewStopExecutionParams creates a new StopExecutionParams object
// with the default values initialized.
func NewStopExecutionParams() *StopExecutionParams {
	var ()
	return &StopExecutionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopExecutionParamsWithTimeout creates a new StopExecutionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopExecutionParamsWithTimeout(timeout time.Duration) *StopExecutionParams {
	var ()
	return &StopExecutionParams{

		timeout: timeout,
	}
}

// NewStopExecutionParamsWithContext creates a new StopExecutionParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopExecutionParamsWithContext(ctx context.Context) *StopExecutionParams {
	var ()
	return &StopExecutionParams{

		Context: ctx,
	}
}

// NewStopExecutionParamsWithHTTPClient creates a new StopExecutionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopExecutionParamsWithHTTPClient(client *http.Client) *StopExecutionParams {
	var ()
	return &StopExecutionParams{
		HTTPClient: client,
	}
}

/*StopExecutionParams contains all the parameters to send to the API endpoint
for the stop execution operation typically these are written to a http.Request
*/
type StopExecutionParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Execution
	  The data of execution

	*/
	Execution *models.Execution
	/*ExecutionID
	  Execution ID

	*/
	ExecutionID int64
	/*PreheatPolicyName
	  Preheat Policy Name

	*/
	PreheatPolicyName string
	/*ProjectName
	  The name of the project

	*/
	ProjectName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop execution params
func (o *StopExecutionParams) WithTimeout(timeout time.Duration) *StopExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop execution params
func (o *StopExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop execution params
func (o *StopExecutionParams) WithContext(ctx context.Context) *StopExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop execution params
func (o *StopExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop execution params
func (o *StopExecutionParams) WithHTTPClient(client *http.Client) *StopExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop execution params
func (o *StopExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the stop execution params
func (o *StopExecutionParams) WithXRequestID(xRequestID *string) *StopExecutionParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the stop execution params
func (o *StopExecutionParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithExecution adds the execution to the stop execution params
func (o *StopExecutionParams) WithExecution(execution *models.Execution) *StopExecutionParams {
	o.SetExecution(execution)
	return o
}

// SetExecution adds the execution to the stop execution params
func (o *StopExecutionParams) SetExecution(execution *models.Execution) {
	o.Execution = execution
}

// WithExecutionID adds the executionID to the stop execution params
func (o *StopExecutionParams) WithExecutionID(executionID int64) *StopExecutionParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the stop execution params
func (o *StopExecutionParams) SetExecutionID(executionID int64) {
	o.ExecutionID = executionID
}

// WithPreheatPolicyName adds the preheatPolicyName to the stop execution params
func (o *StopExecutionParams) WithPreheatPolicyName(preheatPolicyName string) *StopExecutionParams {
	o.SetPreheatPolicyName(preheatPolicyName)
	return o
}

// SetPreheatPolicyName adds the preheatPolicyName to the stop execution params
func (o *StopExecutionParams) SetPreheatPolicyName(preheatPolicyName string) {
	o.PreheatPolicyName = preheatPolicyName
}

// WithProjectName adds the projectName to the stop execution params
func (o *StopExecutionParams) WithProjectName(projectName string) *StopExecutionParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the stop execution params
func (o *StopExecutionParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WriteToRequest writes these params to a swagger request
func (o *StopExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if o.Execution != nil {
		if err := r.SetBodyParam(o.Execution); err != nil {
			return err
		}
	}

	// path param execution_id
	if err := r.SetPathParam("execution_id", swag.FormatInt64(o.ExecutionID)); err != nil {
		return err
	}

	// path param preheat_policy_name
	if err := r.SetPathParam("preheat_policy_name", o.PreheatPolicyName); err != nil {
		return err
	}

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
