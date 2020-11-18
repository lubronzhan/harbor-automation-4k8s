// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Metadata metadata
//
// swagger:model Metadata
type Metadata struct {

	// icon
	Icon string `json:"icon,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// maintainers
	Maintainers []string `json:"maintainers"`

	// name
	Name string `json:"name,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this metadata
func (m *Metadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Metadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metadata) UnmarshalBinary(b []byte) error {
	var res Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
