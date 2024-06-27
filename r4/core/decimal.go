// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Base StructureDefinition for decimal Type: A rational number with implicit
// precision
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/decimal
//   - Source File: StructureDefinition-decimal.json
type Decimal struct {

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*Extension `fhirpath:"extension"`

	// unique id for the element within a resource (for internal references)
	ID string `fhirpath:"id"`

	// The actual value
	Value float64 `fhirpath:"value"`

	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Decimal) GetExtension() []*Extension {
	if d == nil {
		return nil
	}
	return d.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Decimal) GetID() string {
	if d == nil {
		return ""
	}
	return d.ID
}

// GetValue returns the value of the field Value.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Decimal) GetValue() float64 {
	if d == nil {
		return 0.0
	}
	return d.Value
}

func (d *Decimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Value)
}

func (d *Decimal) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &d.Value)
}

var _ json.Marshaler = (*Decimal)(nil)
var _ json.Unmarshaler = (*Decimal)(nil)
