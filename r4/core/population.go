// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.
package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Base StructureDefinition for Population Type: A populatioof people with some
// set of grouping criteria.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Population
//   - Source File: StructureDefinition-Population.json
type Population struct {

	// The age of the specific population.
	Age Element `json:"age"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*Extension `json:"extension"`

	// The gender of the specific population.
	Gender *CodeableConcept `json:"gender"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	Id string `json:"id"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element and that modifies the understanding of the
	// element in which it is contained and/or the understanding of the containing
	// element's descendants. Usually modifier elements provide negation or
	// qualification. To make the use of extensions safe and manageable, there is a
	// strict set of governance applied to the definition and use of extensions.
	// Though any implementer can define an extension, there is a set of
	// requirements that SHALL be met as part of the definition of the extension.
	// Applications processing a resource are required to check for modifier
	// extensions.
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource
	// or DomainResource (including cannot change the meaning of modifierExtension
	// itself).
	ModifierExtension []*Extension `json:"modifierExtension"`

	// The existing physiological conditions of the specific population to which
	// this applies.
	PhysiologicalCondition *CodeableConcept `json:"physiologicalCondition"`

	// Race of the specific population.
	Race *CodeableConcept `json:"race"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

func (v *Population) GetAgeRange() *Range {
	if v == nil {
		return nil
	}
	result, ok := v.Age.(*Range)
	if ok {
		return result
	}
	return nil
}
func (v *Population) GetAgeCodeableConcept() *CodeableConcept {
	if v == nil {
		return nil
	}
	result, ok := v.Age.(*CodeableConcept)
	if ok {
		return result
	}
	return nil
}
func (p *Population) GetAge() Element {
	if p == nil {
		return nil
	}
	return p.Age
}

func (p *Population) GetExtension() []*Extension {
	if p == nil {
		return nil
	}
	return p.Extension
}

func (p *Population) GetGender() *CodeableConcept {
	if p == nil {
		return nil
	}
	return p.Gender
}

func (p *Population) GetId() string {
	if p == nil {
		return ""
	}
	return p.Id
}

func (p *Population) GetModifierExtension() []*Extension {
	if p == nil {
		return nil
	}
	return p.ModifierExtension
}

func (p *Population) GetPhysiologicalCondition() *CodeableConcept {
	if p == nil {
		return nil
	}
	return p.PhysiologicalCondition
}

func (p *Population) GetRace() *CodeableConcept {
	if p == nil {
		return nil
	}
	return p.Race
}