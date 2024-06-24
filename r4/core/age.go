// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Base StructureDefinition for Age Type: A duration of time during which an
// organism (or a process) has existed.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Age
//   - Source File: StructureDefinition-Age.json
type Age struct {

	// A computer processable form of the unit in some unit representation system.
	Code *Code `fhirpath:"code"`

	// How the value should be understood and represented - whether the actual
	// value is greater or less than the stated value due to measurement issues;
	// e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *Code `fhirpath:"comparator"`

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

	// The identification of the system that provides the coded form of the unit.
	System *URI `fhirpath:"system"`

	// A human-readable form of the unit.
	Unit *String `fhirpath:"unit"`

	// The value of the measured amount. The value includes an implicit precision
	// in the presentation of the value.
	Value *Decimal `fhirpath:"value"`

	profileimpl.BaseQuantity
	profileimpl.BaseElement
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Age) GetCode() *Code {
	if a == nil {
		return nil
	}
	return a.Code
}

// GetComparator returns the value of the field Comparator.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Age) GetComparator() *Code {
	if a == nil {
		return nil
	}
	return a.Comparator
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Age) GetExtension() []*Extension {
	if a == nil {
		return nil
	}
	return a.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Age) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

// GetSystem returns the value of the field System.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Age) GetSystem() *URI {
	if a == nil {
		return nil
	}
	return a.System
}

// GetUnit returns the value of the field Unit.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Age) GetUnit() *String {
	if a == nil {
		return nil
	}
	return a.Unit
}

// GetValue returns the value of the field Value.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Age) GetValue() *Decimal {
	if a == nil {
		return nil
	}
	return a.Value
}
