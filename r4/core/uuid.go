// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Base StructureDefinition for uuid type: A UUID, represented as a URI
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/uuid
//   - Source File: StructureDefinition-uuid.json
type UUID struct {

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*Extension `fhirpath:"extension"`

	// unique id for the element within a resource (for internal references)
	ID string `fhirpath:"id"`

	// Primitive value for uuid
	Value string `fhirpath:"value"`

	profileimpl.BaseURI
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (u *UUID) GetExtension() []*Extension {
	if u == nil {
		return nil
	}
	return u.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (u *UUID) GetID() string {
	if u == nil {
		return ""
	}
	return u.ID
}

// GetValue returns the value of the field Value.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (u *UUID) GetValue() string {
	if u == nil {
		return ""
	}
	return u.Value
}

func (u *UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Value)
}

func (u *UUID) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &u.Value)
}

var _ json.Marshaler = (*UUID)(nil)
var _ json.Unmarshaler = (*UUID)(nil)
