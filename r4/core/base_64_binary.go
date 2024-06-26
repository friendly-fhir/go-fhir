// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Base StructureDefinition for base64Binary Type: A stream of bytes
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/base64Binary
//   - Source File: StructureDefinition-base64Binary.json
type Base64Binary struct {

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
func (b6b *Base64Binary) GetExtension() []*Extension {
	if b6b == nil {
		return nil
	}
	return b6b.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b6b *Base64Binary) GetID() string {
	if b6b == nil {
		return ""
	}
	return b6b.ID
}

// GetValue returns the value of the field Value.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b6b *Base64Binary) GetValue() string {
	if b6b == nil {
		return ""
	}
	return b6b.Value
}

func (b6b *Base64Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(b6b.Value)
}

func (b6b *Base64Binary) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &b6b.Value)
}

var _ json.Marshaler = (*Base64Binary)(nil)
var _ json.Unmarshaler = (*Base64Binary)(nil)
