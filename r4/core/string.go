// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Base StructureDefinition for string Type: A sequence of Unicode characters
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/string
//   - Source File: StructureDefinition-string.json
type String struct {

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

	profileimpl.BaseString
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *String) GetExtension() []*Extension {
	if s == nil {
		return nil
	}
	return s.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *String) GetID() string {
	if s == nil {
		return ""
	}
	return s.ID
}

// GetValue returns the value of the field Value.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *String) GetValue() string {
	if s == nil {
		return ""
	}
	return s.Value
}

func (s *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Value)
}

func (s *String) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.Value)
}

var _ json.Marshaler = (*String)(nil)
var _ json.Unmarshaler = (*String)(nil)
