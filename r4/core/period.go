// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Base StructureDefinition for Period Type: A time period defined by a start
// and end date and optionally time.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Period
//   - Source File: StructureDefinition-Period.json
type Period struct {

	// The end of the period. If the end of the period is missing, it means no end
	// was known or planned at the time the instance was created. The start may be
	// in the past, and the end date in the future, which means that period is
	// expected/planned to end at that time.
	End *DateTime `fhirpath:"end"`

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

	// The start of the period. The boundary is inclusive.
	Start *DateTime `fhirpath:"start"`

	profileimpl.BaseElement
}

// GetEnd returns the value of the field End.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Period) GetEnd() *DateTime {
	if p == nil {
		return nil
	}
	return p.End
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Period) GetExtension() []*Extension {
	if p == nil {
		return nil
	}
	return p.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Period) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}

// GetStart returns the value of the field Start.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Period) GetStart() *DateTime {
	if p == nil {
		return nil
	}
	return p.Start
}

func (p *Period) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (p *Period) UnmarshalJSON(data []byte) error {
	var raw struct {
		End       *DateTime    `json:"end"`
		Extension []*Extension `json:"extension"`

		ID    string    `json:"id"`
		Start *DateTime `json:"start"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	p.End = raw.End
	p.Extension = raw.Extension
	p.ID = raw.ID
	p.Start = raw.Start
	return nil
}

var _ json.Marshaler = (*Period)(nil)
var _ json.Unmarshaler = (*Period)(nil)
