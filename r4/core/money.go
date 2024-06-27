// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Base StructureDefinition for Money Type: An amount of economic utility in
// some recognized currency.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Money
//   - Source File: StructureDefinition-Money.json
type Money struct {

	// ISO 4217 Currency Code.
	Currency *Code `fhirpath:"currency"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*Extension `fhirpath:"extension"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID string `fhirpath:"id"`

	// Numerical value (with implicit precision).
	Value *Decimal `fhirpath:"value"`

	profileimpl.BaseElement
}

// GetCurrency returns the value of the field Currency.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (m *Money) GetCurrency() *Code {
	if m == nil {
		return nil
	}
	return m.Currency
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (m *Money) GetExtension() []*Extension {
	if m == nil {
		return nil
	}
	return m.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (m *Money) GetID() string {
	if m == nil {
		return ""
	}
	return m.ID
}

// GetValue returns the value of the field Value.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (m *Money) GetValue() *Decimal {
	if m == nil {
		return nil
	}
	return m.Value
}

func (m *Money) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (m *Money) UnmarshalJSON(data []byte) error {
	var raw struct {
		Currency  *Code        `json:"currency"`
		Extension []*Extension `json:"extension"`

		ID    string   `json:"id"`
		Value *Decimal `json:"value"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	m.Currency = raw.Currency
	m.Extension = raw.Extension
	m.ID = raw.ID
	m.Value = raw.Value
	return nil
}

var _ json.Marshaler = (*Money)(nil)
var _ json.Unmarshaler = (*Money)(nil)
