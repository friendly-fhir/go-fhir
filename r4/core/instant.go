// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Base StructureDefinition for instant Type: An instant in time - known at
// least to the second
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/instant
//   - Source File: StructureDefinition-instant.json
type Instant struct {

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
	Value string `fhirpath:"value"`

	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Instant) GetExtension() []*Extension {
	if i == nil {
		return nil
	}
	return i.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Instant) GetID() string {
	if i == nil {
		return ""
	}
	return i.ID
}

// GetValue returns the value of the field Value.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Instant) GetValue() string {
	if i == nil {
		return ""
	}
	return i.Value
}

func (i *Instant) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}

func (i *Instant) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &i.Value)
}

var _ json.Marshaler = (*Instant)(nil)
var _ json.Unmarshaler = (*Instant)(nil)
