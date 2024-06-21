// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.
package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Base StructureDefinition for Dosage Type: Indicates how the medication
// is/was taken or should be taken by the patient.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Dosage
//   - Source File: StructureDefinition-Dosage.json
type Dosage struct {

	// Supplemental instructions to the patient on how to take the medication (e.g.
	// "with meals" or"take half to one hour before food") or warnings for the
	// patient about the medication (e.g. "may cause drowsiness" or "avoid exposure
	// of skin to direct sunlight or sunlamps").
	AdditionalInstruction []*CodeableConcept `json:"additionalInstruction"`

	// Indicates whether the Medication is only taken when needed within a specific
	// dosing schedule (Boolean option), or it indicates the precondition for
	// taking the Medication (CodeableConcept).
	AsNeeded Element `json:"asNeeded"`

	// The amount of medication administered.
	DoseAndRate []Element `json:"doseAndRate"`

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

	// Upper limit on medication per administration.
	MaxDosePerAdministration *Quantity `json:"maxDosePerAdministration"`

	// Upper limit on medication per lifetime of the patient.
	MaxDosePerLifetime *Quantity `json:"maxDosePerLifetime"`

	// Upper limit on medication per unit of time.
	MaxDosePerPeriod *Ratio `json:"maxDosePerPeriod"`

	// Technique for administering medication.
	Method *CodeableConcept `json:"method"`

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

	// Instructions in terms that are understood by the patient or consumer.
	PatientInstruction *String `json:"patientInstruction"`

	// How drug should enter body.
	Route *CodeableConcept `json:"route"`

	// Indicates the order in which the dosage instructions should be applied or
	// interpreted.
	Sequence *Integer `json:"sequence"`

	// Body site to administer to.
	Site *CodeableConcept `json:"site"`

	// Free text dosage instructions e.g. SIG.
	Text *String `json:"text"`

	// When medication should be administered.
	Timing *Timing `json:"timing"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

func (v *Dosage) GetAsNeededBoolean() *Boolean {
	if v == nil {
		return nil
	}
	result, ok := v.AsNeeded.(*Boolean)
	if ok {
		return result
	}
	return nil
}
func (v *Dosage) GetAsNeededCodeableConcept() *CodeableConcept {
	if v == nil {
		return nil
	}
	result, ok := v.AsNeeded.(*CodeableConcept)
	if ok {
		return result
	}
	return nil
}
func (d *Dosage) GetAdditionalInstruction() []*CodeableConcept {
	if d == nil {
		return nil
	}
	return d.AdditionalInstruction
}

func (d *Dosage) GetAsNeeded() Element {
	if d == nil {
		return nil
	}
	return d.AsNeeded
}

func (d *Dosage) GetDoseAndRate() []Element {
	if d == nil {
		return nil
	}
	return d.DoseAndRate
}

func (d *Dosage) GetExtension() []*Extension {
	if d == nil {
		return nil
	}
	return d.Extension
}

func (d *Dosage) GetId() string {
	if d == nil {
		return ""
	}
	return d.Id
}

func (d *Dosage) GetMaxDosePerAdministration() *Quantity {
	if d == nil {
		return nil
	}
	return d.MaxDosePerAdministration
}

func (d *Dosage) GetMaxDosePerLifetime() *Quantity {
	if d == nil {
		return nil
	}
	return d.MaxDosePerLifetime
}

func (d *Dosage) GetMaxDosePerPeriod() *Ratio {
	if d == nil {
		return nil
	}
	return d.MaxDosePerPeriod
}

func (d *Dosage) GetMethod() *CodeableConcept {
	if d == nil {
		return nil
	}
	return d.Method
}

func (d *Dosage) GetModifierExtension() []*Extension {
	if d == nil {
		return nil
	}
	return d.ModifierExtension
}

func (d *Dosage) GetPatientInstruction() *String {
	if d == nil {
		return nil
	}
	return d.PatientInstruction
}

func (d *Dosage) GetRoute() *CodeableConcept {
	if d == nil {
		return nil
	}
	return d.Route
}

func (d *Dosage) GetSequence() *Integer {
	if d == nil {
		return nil
	}
	return d.Sequence
}

func (d *Dosage) GetSite() *CodeableConcept {
	if d == nil {
		return nil
	}
	return d.Site
}

func (d *Dosage) GetText() *String {
	if d == nil {
		return nil
	}
	return d.Text
}

func (d *Dosage) GetTiming() *Timing {
	if d == nil {
		return nil
	}
	return d.Timing
}

// Amount of medication administered// The amount of medication administered.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Dosage.json
type DosageDoseAndRate struct {

	// Amount of medication per dose.
	Dose Element `json:"dose"`

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

	// Amount of medication per unit of time.
	Rate Element `json:"rate"`

	// The kind of dose or rate specified, for example, ordered or calculated.
	Type *CodeableConcept `json:"type"`

	profileimpl.BaseElement
}

func (v *DosageDoseAndRate) GetDoseRange() *Range {
	if v == nil {
		return nil
	}
	result, ok := v.Dose.(*Range)
	if ok {
		return result
	}
	return nil
}
func (v *DosageDoseAndRate) GetDoseQuantity() *Quantity {
	if v == nil {
		return nil
	}
	result, ok := v.Dose.(*Quantity)
	if ok {
		return result
	}
	return nil
}
func (v *DosageDoseAndRate) GetRateRatio() *Ratio {
	if v == nil {
		return nil
	}
	result, ok := v.Rate.(*Ratio)
	if ok {
		return result
	}
	return nil
}
func (v *DosageDoseAndRate) GetRateRange() *Range {
	if v == nil {
		return nil
	}
	result, ok := v.Rate.(*Range)
	if ok {
		return result
	}
	return nil
}
func (v *DosageDoseAndRate) GetRateQuantity() *Quantity {
	if v == nil {
		return nil
	}
	result, ok := v.Rate.(*Quantity)
	if ok {
		return result
	}
	return nil
}
func (d *DosageDoseAndRate) GetDose() Element {
	if d == nil {
		return nil
	}
	return d.Dose
}

func (d *DosageDoseAndRate) GetExtension() []*Extension {
	if d == nil {
		return nil
	}
	return d.Extension
}

func (d *DosageDoseAndRate) GetId() string {
	if d == nil {
		return ""
	}
	return d.Id
}

func (d *DosageDoseAndRate) GetRate() Element {
	if d == nil {
		return nil
	}
	return d.Rate
}

func (d *DosageDoseAndRate) GetType() *CodeableConcept {
	if d == nil {
		return nil
	}
	return d.Type
}