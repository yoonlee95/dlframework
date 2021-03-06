// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DlframeworkExecutionOptions dlframework execution options
// swagger:model dlframeworkExecutionOptions

type DlframeworkExecutionOptions struct {

	// Options that apply to all CPUs.
	CPUOptions DlframeworkCPUOptions `json:"cpu_options,omitempty"`

	// Map from device type name (e.g., "CPU" or "GPU" ) to maximum
	// number of devices of that type to use.  If a particular device
	// type is not found in the map, the system picks an appropriate
	// number.
	DeviceCount map[string]int32 `json:"device_count,omitempty"`

	// Options that apply to all GPUs.
	GpuOptions *DlframeworkGPUOptions `json:"gpu_options,omitempty"`

	// Time to wait for operation to complete in milliseconds.
	TimeoutInMs int64 `json:"timeout_in_ms,omitempty"`

	// trace level
	TraceLevel ExecutionOptionsTraceLevel `json:"trace_level,omitempty"`
}

/* polymorph dlframeworkExecutionOptions cpu_options false */

/* polymorph dlframeworkExecutionOptions device_count false */

/* polymorph dlframeworkExecutionOptions gpu_options false */

/* polymorph dlframeworkExecutionOptions timeout_in_ms false */

/* polymorph dlframeworkExecutionOptions trace_level false */

// Validate validates this dlframework execution options
func (m *DlframeworkExecutionOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGpuOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTraceLevel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DlframeworkExecutionOptions) validateGpuOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.GpuOptions) { // not required
		return nil
	}

	if m.GpuOptions != nil {

		if err := m.GpuOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gpu_options")
			}
			return err
		}
	}

	return nil
}

func (m *DlframeworkExecutionOptions) validateTraceLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.TraceLevel) { // not required
		return nil
	}

	if err := m.TraceLevel.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trace_level")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkExecutionOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkExecutionOptions) UnmarshalBinary(b []byte) error {
	var res DlframeworkExecutionOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
