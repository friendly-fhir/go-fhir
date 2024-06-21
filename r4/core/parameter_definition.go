// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.
package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Base StructureDefinition for ParameterDefinition Type: The parameters to the
// module. This collection specifies both the input and output parameters.
// Input parameters are provided by the caller as part of the $evaluate
// operation. Output parameters are included in the GuidanceResponse.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/ParameterDefinition
//   - Source File: StructureDefinition-ParameterDefinition.json
type ParameterDefinition struct {

	// A brief discussion of what the parameter is for and how it is used by the
	// module.
	Documentation *String `json:"documentation"`

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

	// The maximum number of times this element is permitted to appear in the
	// request or response.
	Max *String `json:"max"`

	// The minimum number of times this parameter SHALL appear in the request or
	// response.
	Min *Integer `json:"min"`

	// The name of the parameter used to allow access to the value of the parameter
	// in evaluation contexts.
	Name *Code `json:"name"`

	// If specified, this indicates a profile that the input data must conform to,
	// or that the output data will conform to.
	Profile *Canonical `json:"profile"`

	// The type of the parameter.
	Type *Code `json:"type"`

	// Whether the parameter is input or output for the module.
	Use *Code `json:"use"`

	profileimpl.BaseElement
}

func (p *ParameterDefinition) GetDocumentation() *String {
	if p == nil {
		return nil
	}
	return p.Documentation
}

func (p *ParameterDefinition) GetExtension() []*Extension {
	if p == nil {
		return nil
	}
	return p.Extension
}

func (p *ParameterDefinition) GetId() string {
	if p == nil {
		return ""
	}
	return p.Id
}

func (p *ParameterDefinition) GetMax() *String {
	if p == nil {
		return nil
	}
	return p.Max
}

func (p *ParameterDefinition) GetMin() *Integer {
	if p == nil {
		return nil
	}
	return p.Min
}

func (p *ParameterDefinition) GetName() *Code {
	if p == nil {
		return nil
	}
	return p.Name
}

func (p *ParameterDefinition) GetProfile() *Canonical {
	if p == nil {
		return nil
	}
	return p.Profile
}

func (p *ParameterDefinition) GetType() *Code {
	if p == nil {
		return nil
	}
	return p.Type
}

func (p *ParameterDefinition) GetUse() *Code {
	if p == nil {
		return nil
	}
	return p.Use
}