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

// EvalResult eval result
// swagger:model evalResult
type EvalResult struct {

	// eval context
	// Required: true
	EvalContext *EvalContext `json:"evalContext"`

	// eval debug log
	EvalDebugLog *EvalDebugLog `json:"evalDebugLog,omitempty"`

	// flag ID
	// Required: true
	// Minimum: 1
	FlagID *int64 `json:"flagID"`

	// flag snapshot ID
	FlagSnapshotID int64 `json:"flagSnapshotID,omitempty"`

	// segment ID
	// Required: true
	// Minimum: 1
	SegmentID *int64 `json:"segmentID"`

	// timestamp
	// Required: true
	// Min Length: 1
	Timestamp *string `json:"timestamp"`

	// variant attachment
	// Required: true
	VariantAttachment interface{} `json:"variantAttachment"`

	// variant ID
	// Required: true
	// Minimum: 1
	VariantID *int64 `json:"variantID"`

	// variant key
	// Required: true
	// Min Length: 1
	VariantKey *string `json:"variantKey"`
}

// Validate validates this eval result
func (m *EvalResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvalContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvalDebugLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlagID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariantAttachment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariantKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EvalResult) validateEvalContext(formats strfmt.Registry) error {

	if err := validate.Required("evalContext", "body", m.EvalContext); err != nil {
		return err
	}

	if m.EvalContext != nil {
		if err := m.EvalContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evalContext")
			}
			return err
		}
	}

	return nil
}

func (m *EvalResult) validateEvalDebugLog(formats strfmt.Registry) error {

	if swag.IsZero(m.EvalDebugLog) { // not required
		return nil
	}

	if m.EvalDebugLog != nil {
		if err := m.EvalDebugLog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evalDebugLog")
			}
			return err
		}
	}

	return nil
}

func (m *EvalResult) validateFlagID(formats strfmt.Registry) error {

	if err := validate.Required("flagID", "body", m.FlagID); err != nil {
		return err
	}

	if err := validate.MinimumInt("flagID", "body", int64(*m.FlagID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *EvalResult) validateSegmentID(formats strfmt.Registry) error {

	if err := validate.Required("segmentID", "body", m.SegmentID); err != nil {
		return err
	}

	if err := validate.MinimumInt("segmentID", "body", int64(*m.SegmentID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *EvalResult) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.MinLength("timestamp", "body", string(*m.Timestamp), 1); err != nil {
		return err
	}

	return nil
}

func (m *EvalResult) validateVariantAttachment(formats strfmt.Registry) error {

	if err := validate.Required("variantAttachment", "body", m.VariantAttachment); err != nil {
		return err
	}

	return nil
}

func (m *EvalResult) validateVariantID(formats strfmt.Registry) error {

	if err := validate.Required("variantID", "body", m.VariantID); err != nil {
		return err
	}

	if err := validate.MinimumInt("variantID", "body", int64(*m.VariantID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *EvalResult) validateVariantKey(formats strfmt.Registry) error {

	if err := validate.Required("variantKey", "body", m.VariantKey); err != nil {
		return err
	}

	if err := validate.MinLength("variantKey", "body", string(*m.VariantKey), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EvalResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvalResult) UnmarshalBinary(b []byte) error {
	var res EvalResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
