// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.
package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Base StructureDefinition for Coding Type: A reference to a code defined by a
// terminology system.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Coding
//   - Source File: StructureDefinition-Coding.json
type Coding struct {

	// A symbol in syntax defined by the system. The symbol may be a predefined
	// code or an expression in a syntax defined by the coding system (e.g.
	// post-coordination).
	Code *Code `json:"code"`

	// A representation of the meaning of the code in the system, following the
	// rules of the system.
	Display *String `json:"display"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*Extension `json:"extension"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	Id string `json:"id"`

	// The identification of the code system that defines the meaning of the symbol
	// in the code.
	System *Uri `json:"system"`

	// Indicates that this coding was chosen by a user directly - e.g. off a pick
	// list of available items (codes or displays).
	UserSelected *Boolean `json:"userSelected"`

	// The version of the code system which was used when choosing this code. Note
	// that a well-maintained code system does not need the version reported,
	// because the meaning of codes is consistent across versions. However this
	// cannot consistently be assured, and when the meaning is not guaranteed to be
	// consistent, the version SHOULD be exchanged.
	Version *String `json:"version"`

	profileimpl.BaseElement
}

func (c *Coding) GetCode() *Code {
	if c == nil {
		return nil
	}
	return c.Code
}

func (c *Coding) GetDisplay() *String {
	if c == nil {
		return nil
	}
	return c.Display
}

func (c *Coding) GetExtension() []*Extension {
	if c == nil {
		return nil
	}
	return c.Extension
}

func (c *Coding) GetId() string {
	if c == nil {
		return ""
	}
	return c.Id
}

func (c *Coding) GetSystem() *Uri {
	if c == nil {
		return nil
	}
	return c.System
}

func (c *Coding) GetUserSelected() *Boolean {
	if c == nil {
		return nil
	}
	return c.UserSelected
}

func (c *Coding) GetVersion() *String {
	if c == nil {
		return nil
	}
	return c.Version
}
