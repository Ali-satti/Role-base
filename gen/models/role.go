// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Role role
//
// swagger:model role
type Role struct {

	// api name
	// Required: true
	APIName *string `json:"apiName"`

	// designation
	// Required: true
	Designation *string `json:"designation"`
}

// Validate validates this role
func (m *Role) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesignation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Role) validateAPIName(formats strfmt.Registry) error {

	if err := validate.Required("apiName", "body", m.APIName); err != nil {
		return err
	}

	return nil
}

func (m *Role) validateDesignation(formats strfmt.Registry) error {

	if err := validate.Required("designation", "body", m.Designation); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this role based on context it is used
func (m *Role) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Role) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Role) UnmarshalBinary(b []byte) error {
	var res Role
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
