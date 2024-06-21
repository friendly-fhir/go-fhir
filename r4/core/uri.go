// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.
package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Base StructureDefinition for uri Type: String of characters used to identify
// a name or a resource
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/uri
//   - Source File: StructureDefinition-uri.json
type Uri struct {

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*Extension `json:"extension"`

	// unique id for the element within a resource (for internal references)
	Id string `json:"id"`

	// The actual value
	Value string `json:"value"`
	profileimpl.BaseUri

	profileimpl.BaseElement
}

func (u *Uri) GetExtension() []*Extension {
	if u == nil {
		return nil
	}
	return u.Extension
}

func (u *Uri) GetId() string {
	if u == nil {
		return ""
	}
	return u.Id
}

func (u *Uri) GetValue() string {
	if u == nil {
		return ""
	}
	return u.Value
}

var _ Element = (*Uri)(nil)