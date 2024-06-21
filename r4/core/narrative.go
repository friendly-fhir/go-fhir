// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.
package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Base StructureDefinition for Narrative Type: A human-readable summary of the
// resource conveying the essential clinical and business information for the
// resource.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Narrative
//   - Source File: StructureDefinition-Narrative.json
type Narrative struct {

	// The actual narrative content, a stripped down version of XHTML.
	Div *Xhtml `json:"div"`

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

	// The status of the narrative - whether it's entirely generated (from just the
	// defined data or the extensions too), or whether a human authored it and it
	// may contain additional data.
	Status *Code `json:"status"`

	profileimpl.BaseElement
}

func (n *Narrative) GetDiv() *Xhtml {
	if n == nil {
		return nil
	}
	return n.Div
}

func (n *Narrative) GetExtension() []*Extension {
	if n == nil {
		return nil
	}
	return n.Extension
}

func (n *Narrative) GetId() string {
	if n == nil {
		return ""
	}
	return n.Id
}

func (n *Narrative) GetStatus() *Code {
	if n == nil {
		return nil
	}
	return n.Status
}