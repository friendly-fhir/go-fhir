// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.
package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Base StructureDefinition for dateTime Type: A date, date-time or partial
// date (e.g. just year or year + month). If hours and minutes are specified, a
// time zone SHALL be populated. The format is a union of the schema types
// gYear, gYearMonth, date and dateTime. Seconds must be provided due to schema
// type constraints but may be zero-filled and may be ignored. Dates SHALL be
// valid dates.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/dateTime
//   - Source File: StructureDefinition-dateTime.json
type DateTime struct {

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

	profileimpl.BaseElement
}

func (d *DateTime) GetExtension() []*Extension {
	if d == nil {
		return nil
	}
	return d.Extension
}

func (d *DateTime) GetId() string {
	if d == nil {
		return ""
	}
	return d.Id
}

func (d *DateTime) GetValue() string {
	if d == nil {
		return ""
	}
	return d.Value
}

var _ Element = (*DateTime)(nil)
