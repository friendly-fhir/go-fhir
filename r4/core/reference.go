// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Base StructureDefinition for Reference Type: A reference from one resource
// to another.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Reference
//   - Source File: StructureDefinition-Reference.json
type Reference struct {

	// Plain text narrative that identifies the resource in addition to the
	// resource reference.
	Display *String `fhirpath:"display"`

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

	// An identifier for the target resource. This is used when there is no way to
	// reference the other resource directly, either because the entity it
	// represents is not available through a FHIR server, or because there is no
	// way for the author of the resource to convert a known identifier to an
	// actual location. There is no requirement that a Reference.identifier point
	// to something that is actually exposed as a FHIR instance, but it SHALL point
	// to a business concept that would be expected to be exposed as a FHIR
	// instance, and that instance would need to be of a FHIR resource type allowed
	// by the reference.
	Identifier *Identifier `fhirpath:"identifier"`

	// A reference to a location at which the other resource is found. The
	// reference may be a relative reference, in which case it is relative to the
	// service base URL, or an absolute URL that resolves to the location where the
	// resource is found. The reference may be version specific or not. If the
	// reference is not to a FHIR RESTful server, then it should be assumed to be
	// version specific. Internal fragment references (start with '#') refer to
	// contained resources.
	Reference *String `fhirpath:"reference"`

	// The expected type of the target of the reference. If both Reference.type and
	// Reference.reference are populated and Reference.reference is a FHIR URL,
	// both SHALL be consistent.
	// The type is the Canonical URL of Resource Definition that is the type this
	// reference refers to. References are URLs that are relative to
	// http://hl7.org/fhir/StructureDefinition/ e.g. "Patient" is a reference to
	// http://hl7.org/fhir/StructureDefinition/Patient. Absolute URLs are only
	// allowed for logical models (and can only be used in references in logical
	// models, not resources).
	Type *URI `fhirpath:"type"`

	profileimpl.BaseElement
}

// GetDisplay returns the value of the field Display.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (r *Reference) GetDisplay() *String {
	if r == nil {
		return nil
	}
	return r.Display
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (r *Reference) GetExtension() []*Extension {
	if r == nil {
		return nil
	}
	return r.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (r *Reference) GetID() string {
	if r == nil {
		return ""
	}
	return r.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (r *Reference) GetIdentifier() *Identifier {
	if r == nil {
		return nil
	}
	return r.Identifier
}

// GetReference returns the value of the field Reference.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (r *Reference) GetReference() *String {
	if r == nil {
		return nil
	}
	return r.Reference
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (r *Reference) GetType() *URI {
	if r == nil {
		return nil
	}
	return r.Type
}

func (r *Reference) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (r *Reference) UnmarshalJSON(data []byte) error {
	var raw struct {
		Display   *String      `json:"display"`
		Extension []*Extension `json:"extension"`

		ID         string      `json:"id"`
		Identifier *Identifier `json:"identifier"`
		Reference  *String     `json:"reference"`
		Type       *URI        `json:"type"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	r.Display = raw.Display
	r.Extension = raw.Extension
	r.ID = raw.ID
	r.Identifier = raw.Identifier
	r.Reference = raw.Reference
	r.Type = raw.Type
	return nil
}

var _ json.Marshaler = (*Reference)(nil)
var _ json.Unmarshaler = (*Reference)(nil)
