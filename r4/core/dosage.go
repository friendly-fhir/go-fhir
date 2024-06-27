// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
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
	AdditionalInstruction []*CodeableConcept `fhirpath:"additionalInstruction"`

	// Indicates whether the Medication is only taken when needed within a specific
	// dosing schedule (Boolean option), or it indicates the precondition for
	// taking the Medication (CodeableConcept).
	AsNeeded Element `fhirpath:"asNeeded"`

	// The amount of medication administered.
	DoseAndRate []*DosageDoseAndRate `fhirpath:"doseAndRate"`

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

	// Upper limit on medication per administration.
	MaxDosePerAdministration *Quantity `fhirpath:"maxDosePerAdministration"`

	// Upper limit on medication per lifetime of the patient.
	MaxDosePerLifetime *Quantity `fhirpath:"maxDosePerLifetime"`

	// Upper limit on medication per unit of time.
	MaxDosePerPeriod *Ratio `fhirpath:"maxDosePerPeriod"`

	// Technique for administering medication.
	Method *CodeableConcept `fhirpath:"method"`

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
	ModifierExtension []*Extension `fhirpath:"modifierExtension"`

	// Instructions in terms that are understood by the patient or consumer.
	PatientInstruction *String `fhirpath:"patientInstruction"`

	// How drug should enter body.
	Route *CodeableConcept `fhirpath:"route"`

	// Indicates the order in which the dosage instructions should be applied or
	// interpreted.
	Sequence *Integer `fhirpath:"sequence"`

	// Body site to administer to.
	Site *CodeableConcept `fhirpath:"site"`

	// Free text dosage instructions e.g. SIG.
	Text *String `fhirpath:"text"`

	// When medication should be administered.
	Timing *Timing `fhirpath:"timing"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAdditionalInstruction returns the value of the field AdditionalInstruction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetAdditionalInstruction() []*CodeableConcept {
	if d == nil {
		return nil
	}
	return d.AdditionalInstruction
}

// GetAsNeeded returns the value of the field AsNeeded.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetAsNeeded() Element {
	if d == nil {
		return nil
	}
	return d.AsNeeded
}

// GetAsNeededBoolean returns the value of the field AsNeeded.
func (d *Dosage) GetAsNeededBoolean() *Boolean {
	if d == nil {
		return nil
	}
	val, ok := d.AsNeeded.(*Boolean)
	if !ok {
		return nil
	}
	return val
}

// GetAsNeededCodeableConcept returns the value of the field AsNeeded.
func (d *Dosage) GetAsNeededCodeableConcept() *CodeableConcept {
	if d == nil {
		return nil
	}
	val, ok := d.AsNeeded.(*CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetDoseAndRate returns the value of the field DoseAndRate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetDoseAndRate() []*DosageDoseAndRate {
	if d == nil {
		return nil
	}
	return d.DoseAndRate
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetExtension() []*Extension {
	if d == nil {
		return nil
	}
	return d.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetID() string {
	if d == nil {
		return ""
	}
	return d.ID
}

// GetMaxDosePerAdministration returns the value of the field MaxDosePerAdministration.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetMaxDosePerAdministration() *Quantity {
	if d == nil {
		return nil
	}
	return d.MaxDosePerAdministration
}

// GetMaxDosePerLifetime returns the value of the field MaxDosePerLifetime.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetMaxDosePerLifetime() *Quantity {
	if d == nil {
		return nil
	}
	return d.MaxDosePerLifetime
}

// GetMaxDosePerPeriod returns the value of the field MaxDosePerPeriod.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetMaxDosePerPeriod() *Ratio {
	if d == nil {
		return nil
	}
	return d.MaxDosePerPeriod
}

// GetMethod returns the value of the field Method.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetMethod() *CodeableConcept {
	if d == nil {
		return nil
	}
	return d.Method
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetModifierExtension() []*Extension {
	if d == nil {
		return nil
	}
	return d.ModifierExtension
}

// GetPatientInstruction returns the value of the field PatientInstruction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetPatientInstruction() *String {
	if d == nil {
		return nil
	}
	return d.PatientInstruction
}

// GetRoute returns the value of the field Route.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetRoute() *CodeableConcept {
	if d == nil {
		return nil
	}
	return d.Route
}

// GetSequence returns the value of the field Sequence.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetSequence() *Integer {
	if d == nil {
		return nil
	}
	return d.Sequence
}

// GetSite returns the value of the field Site.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetSite() *CodeableConcept {
	if d == nil {
		return nil
	}
	return d.Site
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (d *Dosage) GetText() *String {
	if d == nil {
		return nil
	}
	return d.Text
}

// GetTiming returns the value of the field Timing.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
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
	Dose Element `fhirpath:"dose"`

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

	// Amount of medication per unit of time.
	Rate Element `fhirpath:"rate"`

	// The kind of dose or rate specified, for example, ordered or calculated.
	Type *CodeableConcept `fhirpath:"type"`

	profileimpl.BaseElement
}

// GetDose returns the value of the field Dose.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ddar *DosageDoseAndRate) GetDose() Element {
	if ddar == nil {
		return nil
	}
	return ddar.Dose
}

// GetDoseRange returns the value of the field Dose.
func (ddar *DosageDoseAndRate) GetDoseRange() *Range {
	if ddar == nil {
		return nil
	}
	val, ok := ddar.Dose.(*Range)
	if !ok {
		return nil
	}
	return val
}

// GetDoseQuantity returns the value of the field Dose.
func (ddar *DosageDoseAndRate) GetDoseQuantity() *Quantity {
	if ddar == nil {
		return nil
	}
	val, ok := ddar.Dose.(*Quantity)
	if !ok {
		return nil
	}
	return val
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ddar *DosageDoseAndRate) GetExtension() []*Extension {
	if ddar == nil {
		return nil
	}
	return ddar.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ddar *DosageDoseAndRate) GetID() string {
	if ddar == nil {
		return ""
	}
	return ddar.ID
}

// GetRate returns the value of the field Rate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ddar *DosageDoseAndRate) GetRate() Element {
	if ddar == nil {
		return nil
	}
	return ddar.Rate
}

// GetRateRatio returns the value of the field Rate.
func (ddar *DosageDoseAndRate) GetRateRatio() *Ratio {
	if ddar == nil {
		return nil
	}
	val, ok := ddar.Rate.(*Ratio)
	if !ok {
		return nil
	}
	return val
}

// GetRateRange returns the value of the field Rate.
func (ddar *DosageDoseAndRate) GetRateRange() *Range {
	if ddar == nil {
		return nil
	}
	val, ok := ddar.Rate.(*Range)
	if !ok {
		return nil
	}
	return val
}

// GetRateQuantity returns the value of the field Rate.
func (ddar *DosageDoseAndRate) GetRateQuantity() *Quantity {
	if ddar == nil {
		return nil
	}
	val, ok := ddar.Rate.(*Quantity)
	if !ok {
		return nil
	}
	return val
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ddar *DosageDoseAndRate) GetType() *CodeableConcept {
	if ddar == nil {
		return nil
	}
	return ddar.Type
}

func (d *Dosage) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (d *Dosage) UnmarshalJSON(data []byte) error {
	var raw struct {
		AdditionalInstruction   []*CodeableConcept   `json:"additionalInstruction"`
		AsNeededBoolean         *Boolean             `json:"asNeededBoolean"`
		AsNeededCodeableConcept *CodeableConcept     `json:"asNeededCodeableConcept"`
		DoseAndRate             []*DosageDoseAndRate `json:"doseAndRate"`
		Extension               []*Extension         `json:"extension"`

		ID                       string           `json:"id"`
		MaxDosePerAdministration *Quantity        `json:"maxDosePerAdministration"`
		MaxDosePerLifetime       *Quantity        `json:"maxDosePerLifetime"`
		MaxDosePerPeriod         *Ratio           `json:"maxDosePerPeriod"`
		Method                   *CodeableConcept `json:"method"`
		ModifierExtension        []*Extension     `json:"modifierExtension"`
		PatientInstruction       *String          `json:"patientInstruction"`
		Route                    *CodeableConcept `json:"route"`
		Sequence                 *Integer         `json:"sequence"`
		Site                     *CodeableConcept `json:"site"`
		Text                     *String          `json:"text"`
		Timing                   *Timing          `json:"timing"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	d.AdditionalInstruction = raw.AdditionalInstruction
	d.AsNeeded, err = validate.SelectOneOf[Element]("Dosage.asNeeded",
		raw.AsNeededBoolean,
		raw.AsNeededCodeableConcept)
	if err != nil {
		return err
	}
	d.DoseAndRate = raw.DoseAndRate
	d.Extension = raw.Extension
	d.ID = raw.ID
	d.MaxDosePerAdministration = raw.MaxDosePerAdministration
	d.MaxDosePerLifetime = raw.MaxDosePerLifetime
	d.MaxDosePerPeriod = raw.MaxDosePerPeriod
	d.Method = raw.Method
	d.ModifierExtension = raw.ModifierExtension
	d.PatientInstruction = raw.PatientInstruction
	d.Route = raw.Route
	d.Sequence = raw.Sequence
	d.Site = raw.Site
	d.Text = raw.Text
	d.Timing = raw.Timing
	return nil
}

var _ json.Marshaler = (*Dosage)(nil)
var _ json.Unmarshaler = (*Dosage)(nil)

func (ddar *DosageDoseAndRate) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (ddar *DosageDoseAndRate) UnmarshalJSON(data []byte) error {
	var raw struct {
		DoseRange    *Range       `json:"doseRange"`
		DoseQuantity *Quantity    `json:"doseQuantity"`
		Extension    []*Extension `json:"extension"`

		ID           string           `json:"id"`
		RateRatio    *Ratio           `json:"rateRatio"`
		RateRange    *Range           `json:"rateRange"`
		RateQuantity *Quantity        `json:"rateQuantity"`
		Type         *CodeableConcept `json:"type"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	ddar.Dose, err = validate.SelectOneOf[Element]("Dosage.doseAndRate.dose",
		raw.DoseRange,
		raw.DoseQuantity)
	if err != nil {
		return err
	}
	ddar.Extension = raw.Extension
	ddar.ID = raw.ID
	ddar.Rate, err = validate.SelectOneOf[Element]("Dosage.doseAndRate.rate",
		raw.RateRatio,
		raw.RateRange,
		raw.RateQuantity)
	if err != nil {
		return err
	}
	ddar.Type = raw.Type
	return nil
}

var _ json.Marshaler = (*DosageDoseAndRate)(nil)
var _ json.Unmarshaler = (*DosageDoseAndRate)(nil)
