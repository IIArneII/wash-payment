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

// Transaction transaction
//
// swagger:model Transaction
type Transaction struct {

	// Amount in kopecks (RUB * 10^2)
	// Required: true
	// Minimum: 1
	Amount *int64 `json:"amount"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// Group that requested payment for using the service
	// Format: uuid
	GroupID *strfmt.UUID `json:"groupId,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// operation
	// Required: true
	Operation *Operation `json:"operation"`

	// organization Id
	// Required: true
	// Format: uuid
	OrganizationID *strfmt.UUID `json:"organizationId"`

	// sevice
	// Required: true
	Sevice *Service `json:"sevice"`

	// Number of stations in the car wash that requested payment for using of the service
	// Minimum: 1
	StationsСount *int64 `json:"stationsСount,omitempty"`

	// The user who credited the organisation's account
	UserID *string `json:"userId,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Transaction) UnmarshalJSON(data []byte) error {
	var props struct {

		// Amount in kopecks (RUB * 10^2)
		// Required: true
		// Minimum: 1
		Amount *int64 `json:"amount"`

		// created at
		// Required: true
		// Format: date-time
		CreatedAt *strfmt.DateTime `json:"createdAt"`

		// Group that requested payment for using the service
		// Format: uuid
		GroupID *strfmt.UUID `json:"groupId,omitempty"`

		// id
		// Required: true
		// Format: uuid
		ID *strfmt.UUID `json:"id"`

		// operation
		// Required: true
		Operation *Operation `json:"operation"`

		// organization Id
		// Required: true
		// Format: uuid
		OrganizationID *strfmt.UUID `json:"organizationId"`

		// sevice
		// Required: true
		Sevice *Service `json:"sevice"`

		// Number of stations in the car wash that requested payment for using of the service
		// Minimum: 1
		StationsСount *int64 `json:"stationsСount,omitempty"`

		// The user who credited the organisation's account
		UserID *string `json:"userId,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Amount = props.Amount
	m.CreatedAt = props.CreatedAt
	m.GroupID = props.GroupID
	m.ID = props.ID
	m.Operation = props.Operation
	m.OrganizationID = props.OrganizationID
	m.Sevice = props.Sevice
	m.StationsСount = props.StationsСount
	m.UserID = props.UserID
	return nil
}

// Validate validates this transaction
func (m *Transaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationsСount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.MinimumInt("amount", "body", *m.Amount, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateGroupID(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupID) { // not required
		return nil
	}

	if err := validate.FormatOf("groupId", "body", "uuid", m.GroupID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (m *Transaction) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationId", "body", m.OrganizationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organizationId", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateSevice(formats strfmt.Registry) error {

	if err := validate.Required("sevice", "body", m.Sevice); err != nil {
		return err
	}

	if err := validate.Required("sevice", "body", m.Sevice); err != nil {
		return err
	}

	if m.Sevice != nil {
		if err := m.Sevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sevice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sevice")
			}
			return err
		}
	}

	return nil
}

func (m *Transaction) validateStationsСount(formats strfmt.Registry) error {
	if swag.IsZero(m.StationsСount) { // not required
		return nil
	}

	if err := validate.MinimumInt("stationsСount", "body", *m.StationsСount, 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this transaction based on the context it is used
func (m *Transaction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction) contextValidateOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.Operation != nil {

		if err := m.Operation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (m *Transaction) contextValidateSevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Sevice != nil {

		if err := m.Sevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sevice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sevice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Transaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transaction) UnmarshalBinary(b []byte) error {
	var res Transaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
