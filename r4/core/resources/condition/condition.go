// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package condition

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// A clinical condition, problem, diagnosis, or other event, situation, issue,
// or clinical concept that has risen to a level of concern.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Condition
//   - Source File: StructureDefinition-Condition.json
type Condition struct {

	// The date or estimated date that the condition resolved or went into
	// remission. This is called "abatement" because of the many overloaded
	// connotations associated with "remission" or "resolution" - Conditions are
	// never really resolved, but they can abate.
	Abatement fhir.Element `fhirpath:"abatement"`

	// Individual who is making the condition statement.
	Asserter *fhir.Reference `fhirpath:"asserter"`

	// The anatomical location where this condition manifests itself.
	BodySite []*fhir.CodeableConcept `fhirpath:"bodySite"`

	// A category assigned to the condition.
	Category []*fhir.CodeableConcept `fhirpath:"category"`

	// The clinical status of the condition.
	ClinicalStatus *fhir.CodeableConcept `fhirpath:"clinicalStatus"`

	// Identification of the condition, problem or diagnosis.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The Encounter during which this Condition was created or to which the
	// creation of this record is tightly associated.
	Encounter *fhir.Reference `fhirpath:"encounter"`

	// Supporting evidence / manifestations that are the basis of the Condition's
	// verification status, such as evidence that confirmed or refuted the
	// condition.
	Evidence []*ConditionEvidence `fhirpath:"evidence"`

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// Business identifiers assigned to this condition by the performer or other
	// systems which remain constant as the resource is updated and propagates from
	// server to server.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

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

	// Additional information about the Condition. This is a general notes/comments
	// entry for description of the Condition, its diagnosis and prognosis.
	Note []*fhir.Annotation `fhirpath:"note"`

	// Estimated or actual date or date-time the condition began, in the opinion of
	// the clinician.
	Onset fhir.Element `fhirpath:"onset"`

	// The recordedDate represents when this particular Condition record was
	// created in the system, which is often a system-generated date.
	RecordedDate *fhir.DateTime `fhirpath:"recordedDate"`

	// Individual who recorded the record and takes responsibility for its content.
	Recorder *fhir.Reference `fhirpath:"recorder"`

	// A subjective assessment of the severity of the condition as evaluated by the
	// clinician.
	Severity *fhir.CodeableConcept `fhirpath:"severity"`

	// Clinical stage or grade of a condition. May include formal severity
	// assessments.
	Stage []*ConditionStage `fhirpath:"stage"`

	// Indicates the patient or group who the condition record is associated with.
	Subject *fhir.Reference `fhirpath:"subject"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// The verification status to support the clinical status of the condition.
	VerificationStatus *fhir.CodeableConcept `fhirpath:"verificationStatus"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetAbatement returns the value of the field Abatement.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetAbatement() fhir.Element {
	if c == nil {
		return nil
	}
	return c.Abatement
}

// GetAbatementDateTime returns the value of the field Abatement.
func (c *Condition) GetAbatementDateTime() *fhir.DateTime {
	if c == nil {
		return nil
	}
	val, ok := c.Abatement.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
}

// GetAbatementAge returns the value of the field Abatement.
func (c *Condition) GetAbatementAge() *fhir.Age {
	if c == nil {
		return nil
	}
	val, ok := c.Abatement.(*fhir.Age)
	if !ok {
		return nil
	}
	return val
}

// GetAbatementPeriod returns the value of the field Abatement.
func (c *Condition) GetAbatementPeriod() *fhir.Period {
	if c == nil {
		return nil
	}
	val, ok := c.Abatement.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
}

// GetAbatementRange returns the value of the field Abatement.
func (c *Condition) GetAbatementRange() *fhir.Range {
	if c == nil {
		return nil
	}
	val, ok := c.Abatement.(*fhir.Range)
	if !ok {
		return nil
	}
	return val
}

// GetAbatementString returns the value of the field Abatement.
func (c *Condition) GetAbatementString() *fhir.String {
	if c == nil {
		return nil
	}
	val, ok := c.Abatement.(*fhir.String)
	if !ok {
		return nil
	}
	return val
} // GetAsserter returns the value of the field Asserter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetAsserter() *fhir.Reference {
	if c == nil {
		return nil
	}
	return c.Asserter
}

// GetBodySite returns the value of the field BodySite.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetBodySite() []*fhir.CodeableConcept {
	if c == nil {
		return nil
	}
	return c.BodySite
}

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetCategory() []*fhir.CodeableConcept {
	if c == nil {
		return nil
	}
	return c.Category
}

// GetClinicalStatus returns the value of the field ClinicalStatus.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetClinicalStatus() *fhir.CodeableConcept {
	if c == nil {
		return nil
	}
	return c.ClinicalStatus
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetCode() *fhir.CodeableConcept {
	if c == nil {
		return nil
	}
	return c.Code
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetContained() []fhir.Resource {
	if c == nil {
		return nil
	}
	return c.Contained
}

// GetEncounter returns the value of the field Encounter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetEncounter() *fhir.Reference {
	if c == nil {
		return nil
	}
	return c.Encounter
}

// GetEvidence returns the value of the field Evidence.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetEvidence() []*ConditionEvidence {
	if c == nil {
		return nil
	}
	return c.Evidence
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetExtension() []*fhir.Extension {
	if c == nil {
		return nil
	}
	return c.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetIdentifier() []*fhir.Identifier {
	if c == nil {
		return nil
	}
	return c.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetImplicitRules() *fhir.URI {
	if c == nil {
		return nil
	}
	return c.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetLanguage() *fhir.Code {
	if c == nil {
		return nil
	}
	return c.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetMeta() *fhir.Meta {
	if c == nil {
		return nil
	}
	return c.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetModifierExtension() []*fhir.Extension {
	if c == nil {
		return nil
	}
	return c.ModifierExtension
}

// GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetNote() []*fhir.Annotation {
	if c == nil {
		return nil
	}
	return c.Note
}

// GetOnset returns the value of the field Onset.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetOnset() fhir.Element {
	if c == nil {
		return nil
	}
	return c.Onset
}

// GetOnsetDateTime returns the value of the field Onset.
func (c *Condition) GetOnsetDateTime() *fhir.DateTime {
	if c == nil {
		return nil
	}
	val, ok := c.Onset.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
}

// GetOnsetAge returns the value of the field Onset.
func (c *Condition) GetOnsetAge() *fhir.Age {
	if c == nil {
		return nil
	}
	val, ok := c.Onset.(*fhir.Age)
	if !ok {
		return nil
	}
	return val
}

// GetOnsetPeriod returns the value of the field Onset.
func (c *Condition) GetOnsetPeriod() *fhir.Period {
	if c == nil {
		return nil
	}
	val, ok := c.Onset.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
}

// GetOnsetRange returns the value of the field Onset.
func (c *Condition) GetOnsetRange() *fhir.Range {
	if c == nil {
		return nil
	}
	val, ok := c.Onset.(*fhir.Range)
	if !ok {
		return nil
	}
	return val
}

// GetOnsetString returns the value of the field Onset.
func (c *Condition) GetOnsetString() *fhir.String {
	if c == nil {
		return nil
	}
	val, ok := c.Onset.(*fhir.String)
	if !ok {
		return nil
	}
	return val
} // GetRecordedDate returns the value of the field RecordedDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetRecordedDate() *fhir.DateTime {
	if c == nil {
		return nil
	}
	return c.RecordedDate
}

// GetRecorder returns the value of the field Recorder.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetRecorder() *fhir.Reference {
	if c == nil {
		return nil
	}
	return c.Recorder
}

// GetSeverity returns the value of the field Severity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetSeverity() *fhir.CodeableConcept {
	if c == nil {
		return nil
	}
	return c.Severity
}

// GetStage returns the value of the field Stage.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetStage() []*ConditionStage {
	if c == nil {
		return nil
	}
	return c.Stage
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetSubject() *fhir.Reference {
	if c == nil {
		return nil
	}
	return c.Subject
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetText() *fhir.Narrative {
	if c == nil {
		return nil
	}
	return c.Text
}

// GetVerificationStatus returns the value of the field VerificationStatus.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Condition) GetVerificationStatus() *fhir.CodeableConcept {
	if c == nil {
		return nil
	}
	return c.VerificationStatus
}

// Supporting evidence// Supporting evidence / manifestations that are the basis of the Condition's
// verification status, such as evidence that confirmed or refuted the
// condition.// The evidence may be a simple list of coded symptoms/manifestations, or
// references to observations or formal assessments, or both.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Condition.json
type ConditionEvidence struct {

	// A manifestation or symptom that led to the recording of this condition.
	Code []*fhir.CodeableConcept `fhirpath:"code"`

	// Links to other relevant information, including pathology reports.
	Detail []*fhir.Reference `fhirpath:"detail"`

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

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ce *ConditionEvidence) GetCode() []*fhir.CodeableConcept {
	if ce == nil {
		return nil
	}
	return ce.Code
}

// GetDetail returns the value of the field Detail.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ce *ConditionEvidence) GetDetail() []*fhir.Reference {
	if ce == nil {
		return nil
	}
	return ce.Detail
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ce *ConditionEvidence) GetExtension() []*fhir.Extension {
	if ce == nil {
		return nil
	}
	return ce.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ce *ConditionEvidence) GetID() string {
	if ce == nil {
		return ""
	}
	return ce.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ce *ConditionEvidence) GetModifierExtension() []*fhir.Extension {
	if ce == nil {
		return nil
	}
	return ce.ModifierExtension
}

// Stage/grade, usually assessed formally// Clinical stage or grade of a condition. May include formal severity
// assessments.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Condition.json
type ConditionStage struct {

	// Reference to a formal record of the evidence on which the staging assessment
	// is based.
	Assessment []*fhir.Reference `fhirpath:"assessment"`

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

	// A simple summary of the stage such as "Stage 3". The determination of the
	// stage is disease-specific.
	Summary *fhir.CodeableConcept `fhirpath:"summary"`

	// The kind of staging, such as pathological or clinical staging.
	Type *fhir.CodeableConcept `fhirpath:"type"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAssessment returns the value of the field Assessment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cs *ConditionStage) GetAssessment() []*fhir.Reference {
	if cs == nil {
		return nil
	}
	return cs.Assessment
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cs *ConditionStage) GetExtension() []*fhir.Extension {
	if cs == nil {
		return nil
	}
	return cs.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cs *ConditionStage) GetID() string {
	if cs == nil {
		return ""
	}
	return cs.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cs *ConditionStage) GetModifierExtension() []*fhir.Extension {
	if cs == nil {
		return nil
	}
	return cs.ModifierExtension
}

// GetSummary returns the value of the field Summary.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cs *ConditionStage) GetSummary() *fhir.CodeableConcept {
	if cs == nil {
		return nil
	}
	return cs.Summary
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cs *ConditionStage) GetType() *fhir.CodeableConcept {
	if cs == nil {
		return nil
	}
	return cs.Type
}

func (c *Condition) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (c *Condition) UnmarshalJSON(data []byte) error {
	var raw struct {
		AbatementDateTime *fhir.DateTime          `json:"abatementDateTime"`
		AbatementAge      *fhir.Age               `json:"abatementAge"`
		AbatementPeriod   *fhir.Period            `json:"abatementPeriod"`
		AbatementRange    *fhir.Range             `json:"abatementRange"`
		AbatementString   *fhir.String            `json:"abatementString"`
		Asserter          *fhir.Reference         `json:"asserter"`
		BodySite          []*fhir.CodeableConcept `json:"bodySite"`
		Category          []*fhir.CodeableConcept `json:"category"`
		ClinicalStatus    *fhir.CodeableConcept   `json:"clinicalStatus"`
		Code              *fhir.CodeableConcept   `json:"code"`
		Contained         []fhir.Resource         `json:"contained"`
		Encounter         *fhir.Reference         `json:"encounter"`
		Evidence          []*ConditionEvidence    `json:"evidence"`
		Extension         []*fhir.Extension       `json:"extension"`

		ID                 string                `json:"id"`
		Identifier         []*fhir.Identifier    `json:"identifier"`
		ImplicitRules      *fhir.URI             `json:"implicitRules"`
		Language           *fhir.Code            `json:"language"`
		Meta               *fhir.Meta            `json:"meta"`
		ModifierExtension  []*fhir.Extension     `json:"modifierExtension"`
		Note               []*fhir.Annotation    `json:"note"`
		OnsetDateTime      *fhir.DateTime        `json:"onsetDateTime"`
		OnsetAge           *fhir.Age             `json:"onsetAge"`
		OnsetPeriod        *fhir.Period          `json:"onsetPeriod"`
		OnsetRange         *fhir.Range           `json:"onsetRange"`
		OnsetString        *fhir.String          `json:"onsetString"`
		RecordedDate       *fhir.DateTime        `json:"recordedDate"`
		Recorder           *fhir.Reference       `json:"recorder"`
		Severity           *fhir.CodeableConcept `json:"severity"`
		Stage              []*ConditionStage     `json:"stage"`
		Subject            *fhir.Reference       `json:"subject"`
		Text               *fhir.Narrative       `json:"text"`
		VerificationStatus *fhir.CodeableConcept `json:"verificationStatus"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c.Abatement, err = validate.SelectOneOf[fhir.Element]("Condition.abatement",
		raw.AbatementDateTime,
		raw.AbatementAge,
		raw.AbatementPeriod,
		raw.AbatementRange,
		raw.AbatementString)
	if err != nil {
		return err
	}
	c.Asserter = raw.Asserter
	c.BodySite = raw.BodySite
	c.Category = raw.Category
	c.ClinicalStatus = raw.ClinicalStatus
	c.Code = raw.Code
	c.Contained = raw.Contained
	c.Encounter = raw.Encounter
	c.Evidence = raw.Evidence
	c.Extension = raw.Extension
	c.ID = raw.ID
	c.Identifier = raw.Identifier
	c.ImplicitRules = raw.ImplicitRules
	c.Language = raw.Language
	c.Meta = raw.Meta
	c.ModifierExtension = raw.ModifierExtension
	c.Note = raw.Note
	c.Onset, err = validate.SelectOneOf[fhir.Element]("Condition.onset",
		raw.OnsetDateTime,
		raw.OnsetAge,
		raw.OnsetPeriod,
		raw.OnsetRange,
		raw.OnsetString)
	if err != nil {
		return err
	}
	c.RecordedDate = raw.RecordedDate
	c.Recorder = raw.Recorder
	c.Severity = raw.Severity
	c.Stage = raw.Stage
	c.Subject = raw.Subject
	c.Text = raw.Text
	c.VerificationStatus = raw.VerificationStatus
	return nil
}

var _ json.Marshaler = (*Condition)(nil)
var _ json.Unmarshaler = (*Condition)(nil)

func (ce *ConditionEvidence) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (ce *ConditionEvidence) UnmarshalJSON(data []byte) error {
	var raw struct {
		Code      []*fhir.CodeableConcept `json:"code"`
		Detail    []*fhir.Reference       `json:"detail"`
		Extension []*fhir.Extension       `json:"extension"`

		ID                string            `json:"id"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	ce.Code = raw.Code
	ce.Detail = raw.Detail
	ce.Extension = raw.Extension
	ce.ID = raw.ID
	ce.ModifierExtension = raw.ModifierExtension
	return nil
}

var _ json.Marshaler = (*ConditionEvidence)(nil)
var _ json.Unmarshaler = (*ConditionEvidence)(nil)

func (cs *ConditionStage) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (cs *ConditionStage) UnmarshalJSON(data []byte) error {
	var raw struct {
		Assessment []*fhir.Reference `json:"assessment"`
		Extension  []*fhir.Extension `json:"extension"`

		ID                string                `json:"id"`
		ModifierExtension []*fhir.Extension     `json:"modifierExtension"`
		Summary           *fhir.CodeableConcept `json:"summary"`
		Type              *fhir.CodeableConcept `json:"type"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	cs.Assessment = raw.Assessment
	cs.Extension = raw.Extension
	cs.ID = raw.ID
	cs.ModifierExtension = raw.ModifierExtension
	cs.Summary = raw.Summary
	cs.Type = raw.Type
	return nil
}

var _ json.Marshaler = (*ConditionStage)(nil)
var _ json.Unmarshaler = (*ConditionStage)(nil)
