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

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// operation
	// Required: true
	// Enum: [deposit debit]
	Operation *string `json:"operation"`

	// sevice
	// Required: true
	Sevice *string `json:"sevice"`
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

		// id
		// Required: true
		// Format: uuid
		ID *strfmt.UUID `json:"id"`

		// operation
		// Required: true
		// Enum: [deposit debit]
		Operation *string `json:"operation"`

		// sevice
		// Required: true
		Sevice *string `json:"sevice"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Amount = props.Amount
	m.CreatedAt = props.CreatedAt
	m.ID = props.ID
	m.Operation = props.Operation
	m.Sevice = props.Sevice
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

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSevice(formats); err != nil {
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

func (m *Transaction) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

var transactionTypeOperationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["deposit","debit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionTypeOperationPropEnum = append(transactionTypeOperationPropEnum, v)
	}
}

const (

	// TransactionOperationDeposit captures enum value "deposit"
	TransactionOperationDeposit string = "deposit"

	// TransactionOperationDebit captures enum value "debit"
	TransactionOperationDebit string = "debit"
)

// prop value enum
func (m *Transaction) validateOperationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transactionTypeOperationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Transaction) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	// value enum
	if err := m.validateOperationEnum("operation", "body", *m.Operation); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateSevice(formats strfmt.Registry) error {

	if err := validate.Required("sevice", "body", m.Sevice); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this transaction based on context it is used
func (m *Transaction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
