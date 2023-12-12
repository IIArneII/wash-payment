// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Deposit deposit
//
// swagger:model Deposit
type Deposit struct {

	// Amount in kopecks (RUB * 10^2)
	// Required: true
	// Minimum: 1
	Amount *int64 `json:"amount"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Deposit) UnmarshalJSON(data []byte) error {
	var props struct {

		// Amount in kopecks (RUB * 10^2)
		// Required: true
		// Minimum: 1
		Amount *int64 `json:"amount"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Amount = props.Amount
	return nil
}

// Validate validates this deposit
func (m *Deposit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Deposit) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.MinimumInt("amount", "body", *m.Amount, 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this deposit based on context it is used
func (m *Deposit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Deposit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Deposit) UnmarshalBinary(b []byte) error {
	var res Deposit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}