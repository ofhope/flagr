// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EvalContext eval context
// swagger:model evalContext
type EvalContext struct {

	// enable debug
	EnableDebug bool `json:"enableDebug,omitempty"`

	// entity context
	EntityContext interface{} `json:"entityContext,omitempty"`

	// entityID is used to deterministically at random to evaluate the flag result. If it's empty, flagr will randomly generate one.
	EntityID string `json:"entityID,omitempty"`

	// entity type
	// Required: true
	// Min Length: 1
	EntityType *string `json:"entityType"`

	// flag ID
	// Required: true
	// Minimum: 1
	FlagID *int64 `json:"flagID"`
}

// Validate validates this eval context
func (m *EvalContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlagID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EvalContext) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entityType", "body", m.EntityType); err != nil {
		return err
	}

	if err := validate.MinLength("entityType", "body", string(*m.EntityType), 1); err != nil {
		return err
	}

	return nil
}

func (m *EvalContext) validateFlagID(formats strfmt.Registry) error {

	if err := validate.Required("flagID", "body", m.FlagID); err != nil {
		return err
	}

	if err := validate.MinimumInt("flagID", "body", int64(*m.FlagID), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EvalContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvalContext) UnmarshalBinary(b []byte) error {
	var res EvalContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
