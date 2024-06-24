// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package measurereport

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// The MeasureReport resource contains the results of the calculation of a
// measure; and optionally a reference to the resources involved in that
// calculation.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/MeasureReport
//   - Source File: StructureDefinition-MeasureReport.json
type MeasureReport struct {

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The date this measure report was generated.
	Date *fhir.DateTime `fhirpath:"date"`

	// A reference to a Bundle containing the Resources that were used in the
	// calculation of this measure.
	EvaluatedResource []*fhir.Reference `fhirpath:"evaluatedResource"`

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// The results of the calculation, one for each population group in the
	// measure.
	Group []*MeasureReportGroup `fhirpath:"group"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// A formal identifier that is used to identify this MeasureReport when it is
	// represented in other formats or referenced in a specification, model, design
	// or an instance.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// Whether improvement in the measure is noted by an increase or decrease in
	// the measure score.
	ImprovementNotation *fhir.CodeableConcept `fhirpath:"improvementNotation"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// A reference to the Measure that was calculated to produce this report.
	Measure *fhir.Canonical `fhirpath:"measure"`

	// The metadata about the resource. This is content that is maintained by the
	// infrastructure. Changes to the content might not always be associated with
	// version changes to the resource.
	Meta *fhir.Meta `fhirpath:"meta"`

	// May be used to represent additional information that is not part of the
	// basic definition of the resource and that modifies the understanding of the
	// element that contains it and/or the understanding of the containing
	// element's descendants. Usually modifier elements provide negation or
	// qualification. To make the use of extensions safe and manageable, there is a
	// strict set of governance applied to the definition and use of extensions.
	// Though any implementer is allowed to define an extension, there is a set of
	// requirements that SHALL be met as part of the definition of the extension.
	// Applications processing a resource are required to check for modifier
	// extensions.
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource
	// or DomainResource (including cannot change the meaning of modifierExtension
	// itself).
	ModifierExtension []*fhir.Extension `fhirpath:"modifierExtension"`

	// The reporting period for which the report was calculated.
	Period *fhir.Period `fhirpath:"period"`

	// The individual, location, or organization that is reporting the data.
	Reporter *fhir.Reference `fhirpath:"reporter"`

	// The MeasureReport status. No data will be available until the MeasureReport
	// status is complete.
	Status *fhir.Code `fhirpath:"status"`

	// Optional subject identifying the individual or individuals the report is
	// for.
	Subject *fhir.Reference `fhirpath:"subject"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// The type of measure report. This may be an individual report, which provides
	// the score for the measure for an individual member of the population; a
	// subject-listing, which returns the list of members that meet the various
	// criteria in the measure; a summary report, which returns a population count
	// for each of the criteria in the measure; or a data-collection, which enables
	// the MeasureReport to be used to exchange the data-of-interest for a quality
	// measure.
	Type *fhir.Code `fhirpath:"type"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetContained() []fhir.Resource {
	if mr == nil {
		return nil
	}
	return mr.Contained
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetDate() *fhir.DateTime {
	if mr == nil {
		return nil
	}
	return mr.Date
}

// GetEvaluatedResource returns the value of the field EvaluatedResource.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetEvaluatedResource() []*fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.EvaluatedResource
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetExtension() []*fhir.Extension {
	if mr == nil {
		return nil
	}
	return mr.Extension
}

// GetGroup returns the value of the field Group.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetGroup() []*MeasureReportGroup {
	if mr == nil {
		return nil
	}
	return mr.Group
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetID() string {
	if mr == nil {
		return ""
	}
	return mr.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetIdentifier() []*fhir.Identifier {
	if mr == nil {
		return nil
	}
	return mr.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetImplicitRules() *fhir.URI {
	if mr == nil {
		return nil
	}
	return mr.ImplicitRules
}

// GetImprovementNotation returns the value of the field ImprovementNotation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetImprovementNotation() *fhir.CodeableConcept {
	if mr == nil {
		return nil
	}
	return mr.ImprovementNotation
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetLanguage() *fhir.Code {
	if mr == nil {
		return nil
	}
	return mr.Language
}

// GetMeasure returns the value of the field Measure.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetMeasure() *fhir.Canonical {
	if mr == nil {
		return nil
	}
	return mr.Measure
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetMeta() *fhir.Meta {
	if mr == nil {
		return nil
	}
	return mr.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetModifierExtension() []*fhir.Extension {
	if mr == nil {
		return nil
	}
	return mr.ModifierExtension
}

// GetPeriod returns the value of the field Period.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetPeriod() *fhir.Period {
	if mr == nil {
		return nil
	}
	return mr.Period
}

// GetReporter returns the value of the field Reporter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetReporter() *fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.Reporter
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetStatus() *fhir.Code {
	if mr == nil {
		return nil
	}
	return mr.Status
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetSubject() *fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.Subject
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetText() *fhir.Narrative {
	if mr == nil {
		return nil
	}
	return mr.Text
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MeasureReport) GetType() *fhir.Code {
	if mr == nil {
		return nil
	}
	return mr.Type
}

// Measure results for each group// The results of the calculation, one for each population group in the
// measure.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MeasureReport.json
type MeasureReportGroup struct {

	// The meaning of the population group as defined in the measure definition.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID string `fhirpath:"id"`

	// The measure score for this population group, calculated as appropriate for
	// the measure type and scoring method, and based on the contents of the
	// populations defined in the group.
	MeasureScore *fhir.Quantity `fhirpath:"measureScore"`

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
	ModifierExtension []*fhir.Extension `fhirpath:"modifierExtension"`

	// The populations that make up the population group, one for each type of
	// population appropriate for the measure.
	Population []*MeasureReportGroupPopulation `fhirpath:"population"`

	// When a measure includes multiple stratifiers, there will be a stratifier
	// group for each stratifier defined by the measure.
	Stratifier []*MeasureReportGroupStratifier `fhirpath:"stratifier"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrg *MeasureReportGroup) GetCode() *fhir.CodeableConcept {
	if mrg == nil {
		return nil
	}
	return mrg.Code
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrg *MeasureReportGroup) GetExtension() []*fhir.Extension {
	if mrg == nil {
		return nil
	}
	return mrg.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrg *MeasureReportGroup) GetID() string {
	if mrg == nil {
		return ""
	}
	return mrg.ID
}

// GetMeasureScore returns the value of the field MeasureScore.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrg *MeasureReportGroup) GetMeasureScore() *fhir.Quantity {
	if mrg == nil {
		return nil
	}
	return mrg.MeasureScore
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrg *MeasureReportGroup) GetModifierExtension() []*fhir.Extension {
	if mrg == nil {
		return nil
	}
	return mrg.ModifierExtension
}

// GetPopulation returns the value of the field Population.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrg *MeasureReportGroup) GetPopulation() []*MeasureReportGroupPopulation {
	if mrg == nil {
		return nil
	}
	return mrg.Population
}

// GetStratifier returns the value of the field Stratifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrg *MeasureReportGroup) GetStratifier() []*MeasureReportGroupStratifier {
	if mrg == nil {
		return nil
	}
	return mrg.Stratifier
}

// The populations in the group// The populations that make up the population group, one for each type of
// population appropriate for the measure.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MeasureReport.json
type MeasureReportGroupPopulation struct {

	// The type of the population.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// The number of members of the population.
	Count *fhir.Integer `fhirpath:"count"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID string `fhirpath:"id"`

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
	ModifierExtension []*fhir.Extension `fhirpath:"modifierExtension"`

	// This element refers to a List of subject level MeasureReport resources, one
	// for each subject in this population.
	SubjectResults *fhir.Reference `fhirpath:"subjectResults"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgp *MeasureReportGroupPopulation) GetCode() *fhir.CodeableConcept {
	if mrgp == nil {
		return nil
	}
	return mrgp.Code
}

// GetCount returns the value of the field Count.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgp *MeasureReportGroupPopulation) GetCount() *fhir.Integer {
	if mrgp == nil {
		return nil
	}
	return mrgp.Count
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgp *MeasureReportGroupPopulation) GetExtension() []*fhir.Extension {
	if mrgp == nil {
		return nil
	}
	return mrgp.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgp *MeasureReportGroupPopulation) GetID() string {
	if mrgp == nil {
		return ""
	}
	return mrgp.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgp *MeasureReportGroupPopulation) GetModifierExtension() []*fhir.Extension {
	if mrgp == nil {
		return nil
	}
	return mrgp.ModifierExtension
}

// GetSubjectResults returns the value of the field SubjectResults.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgp *MeasureReportGroupPopulation) GetSubjectResults() *fhir.Reference {
	if mrgp == nil {
		return nil
	}
	return mrgp.SubjectResults
}

// Stratification results// When a measure includes multiple stratifiers, there will be a stratifier
// group for each stratifier defined by the measure.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MeasureReport.json
type MeasureReportGroupStratifier struct {

	// The meaning of this stratifier, as defined in the measure definition.
	Code []*fhir.CodeableConcept `fhirpath:"code"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID string `fhirpath:"id"`

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
	ModifierExtension []*fhir.Extension `fhirpath:"modifierExtension"`

	// This element contains the results for a single stratum within the
	// stratifier. For example, when stratifying on administrative gender, there
	// will be four strata, one for each possible gender value.
	Stratum []*MeasureReportGroupStratifierStratum `fhirpath:"stratum"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgs *MeasureReportGroupStratifier) GetCode() []*fhir.CodeableConcept {
	if mrgs == nil {
		return nil
	}
	return mrgs.Code
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgs *MeasureReportGroupStratifier) GetExtension() []*fhir.Extension {
	if mrgs == nil {
		return nil
	}
	return mrgs.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgs *MeasureReportGroupStratifier) GetID() string {
	if mrgs == nil {
		return ""
	}
	return mrgs.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgs *MeasureReportGroupStratifier) GetModifierExtension() []*fhir.Extension {
	if mrgs == nil {
		return nil
	}
	return mrgs.ModifierExtension
}

// GetStratum returns the value of the field Stratum.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgs *MeasureReportGroupStratifier) GetStratum() []*MeasureReportGroupStratifierStratum {
	if mrgs == nil {
		return nil
	}
	return mrgs.Stratum
}

// Stratum results, one for each unique value, or set of values, in the
// stratifier, or stratifier components// This element contains the results for a single stratum within the
// stratifier. For example, when stratifying on administrative gender, there
// will be four strata, one for each possible gender value.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MeasureReport.json
type MeasureReportGroupStratifierStratum struct {

	// A stratifier component value.
	Component []*MeasureReportGroupStratifierStratumComponent `fhirpath:"component"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID string `fhirpath:"id"`

	// The measure score for this stratum, calculated as appropriate for the
	// measure type and scoring method, and based on only the members of this
	// stratum.
	MeasureScore *fhir.Quantity `fhirpath:"measureScore"`

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
	ModifierExtension []*fhir.Extension `fhirpath:"modifierExtension"`

	// The populations that make up the stratum, one for each type of population
	// appropriate to the measure.
	Population []*MeasureReportGroupStratifierStratumPopulation `fhirpath:"population"`

	// The value for this stratum, expressed as a CodeableConcept. When defining
	// stratifiers on complex values, the value must be rendered such that the
	// value for each stratum within the stratifier is unique.
	Value *fhir.CodeableConcept `fhirpath:"value"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetComponent returns the value of the field Component.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgss *MeasureReportGroupStratifierStratum) GetComponent() []*MeasureReportGroupStratifierStratumComponent {
	if mrgss == nil {
		return nil
	}
	return mrgss.Component
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgss *MeasureReportGroupStratifierStratum) GetExtension() []*fhir.Extension {
	if mrgss == nil {
		return nil
	}
	return mrgss.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgss *MeasureReportGroupStratifierStratum) GetID() string {
	if mrgss == nil {
		return ""
	}
	return mrgss.ID
}

// GetMeasureScore returns the value of the field MeasureScore.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgss *MeasureReportGroupStratifierStratum) GetMeasureScore() *fhir.Quantity {
	if mrgss == nil {
		return nil
	}
	return mrgss.MeasureScore
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgss *MeasureReportGroupStratifierStratum) GetModifierExtension() []*fhir.Extension {
	if mrgss == nil {
		return nil
	}
	return mrgss.ModifierExtension
}

// GetPopulation returns the value of the field Population.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgss *MeasureReportGroupStratifierStratum) GetPopulation() []*MeasureReportGroupStratifierStratumPopulation {
	if mrgss == nil {
		return nil
	}
	return mrgss.Population
}

// GetValue returns the value of the field Value.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgss *MeasureReportGroupStratifierStratum) GetValue() *fhir.CodeableConcept {
	if mrgss == nil {
		return nil
	}
	return mrgss.Value
}

// Stratifier component values// A stratifier component value.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MeasureReport.json
type MeasureReportGroupStratifierStratumComponent struct {

	// The code for the stratum component value.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID string `fhirpath:"id"`

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
	ModifierExtension []*fhir.Extension `fhirpath:"modifierExtension"`

	// The stratum component value.
	Value *fhir.CodeableConcept `fhirpath:"value"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgssc *MeasureReportGroupStratifierStratumComponent) GetCode() *fhir.CodeableConcept {
	if mrgssc == nil {
		return nil
	}
	return mrgssc.Code
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgssc *MeasureReportGroupStratifierStratumComponent) GetExtension() []*fhir.Extension {
	if mrgssc == nil {
		return nil
	}
	return mrgssc.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgssc *MeasureReportGroupStratifierStratumComponent) GetID() string {
	if mrgssc == nil {
		return ""
	}
	return mrgssc.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgssc *MeasureReportGroupStratifierStratumComponent) GetModifierExtension() []*fhir.Extension {
	if mrgssc == nil {
		return nil
	}
	return mrgssc.ModifierExtension
}

// GetValue returns the value of the field Value.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgssc *MeasureReportGroupStratifierStratumComponent) GetValue() *fhir.CodeableConcept {
	if mrgssc == nil {
		return nil
	}
	return mrgssc.Value
}

// Population results in this stratum// The populations that make up the stratum, one for each type of population
// appropriate to the measure.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MeasureReport.json
type MeasureReportGroupStratifierStratumPopulation struct {

	// The type of the population.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// The number of members of the population in this stratum.
	Count *fhir.Integer `fhirpath:"count"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID string `fhirpath:"id"`

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
	ModifierExtension []*fhir.Extension `fhirpath:"modifierExtension"`

	// This element refers to a List of subject level MeasureReport resources, one
	// for each subject in this population in this stratum.
	SubjectResults *fhir.Reference `fhirpath:"subjectResults"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgssp *MeasureReportGroupStratifierStratumPopulation) GetCode() *fhir.CodeableConcept {
	if mrgssp == nil {
		return nil
	}
	return mrgssp.Code
}

// GetCount returns the value of the field Count.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgssp *MeasureReportGroupStratifierStratumPopulation) GetCount() *fhir.Integer {
	if mrgssp == nil {
		return nil
	}
	return mrgssp.Count
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgssp *MeasureReportGroupStratifierStratumPopulation) GetExtension() []*fhir.Extension {
	if mrgssp == nil {
		return nil
	}
	return mrgssp.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgssp *MeasureReportGroupStratifierStratumPopulation) GetID() string {
	if mrgssp == nil {
		return ""
	}
	return mrgssp.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgssp *MeasureReportGroupStratifierStratumPopulation) GetModifierExtension() []*fhir.Extension {
	if mrgssp == nil {
		return nil
	}
	return mrgssp.ModifierExtension
}

// GetSubjectResults returns the value of the field SubjectResults.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrgssp *MeasureReportGroupStratifierStratumPopulation) GetSubjectResults() *fhir.Reference {
	if mrgssp == nil {
		return nil
	}
	return mrgssp.SubjectResults
}
