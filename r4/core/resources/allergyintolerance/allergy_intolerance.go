// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package allergyintolerance

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Risk of harmful or undesirable, physiological response which is unique to an
// individual and associated with exposure to a substance.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/AllergyIntolerance
//   - Source File: StructureDefinition-AllergyIntolerance.json
type AllergyIntolerance struct {

	// The source of the information about the allergy that is recorded.
	Asserter *fhir.Reference `fhirpath:"asserter"`

	// Category of the identified substance.
	Category []*fhir.Code `fhirpath:"category"`

	// The clinical status of the allergy or intolerance.
	ClinicalStatus *fhir.CodeableConcept `fhirpath:"clinicalStatus"`

	// Code for an allergy or intolerance statement (either a positive or a
	// negated/excluded statement). This may be a code for a substance or
	// pharmaceutical product that is considered to be responsible for the adverse
	// reaction risk (e.g., "Latex"), an allergy or intolerance condition (e.g.,
	// "Latex allergy"), or a negated/excluded code for a specific substance or
	// class (e.g., "No latex allergy") or a general or categorical negated
	// statement (e.g., "No known allergy", "No known drug allergies"). Note: the
	// substance for a specific reaction may be different from the substance
	// identified as the cause of the risk, but it must be consistent with it. For
	// instance, it may be a more specific substance (e.g. a brand medication) or a
	// composite product that includes the identified substance. It must be
	// clinically safe to only process the 'code' and ignore the
	// 'reaction.substance'. If a receiving system is unable to confirm that
	// AllergyIntolerance.reaction.substance falls within the semantic scope of
	// AllergyIntolerance.code, then the receiving system should ignore
	// AllergyIntolerance.reaction.substance.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// Estimate of the potential clinical harm, or seriousness, of the reaction to
	// the identified substance.
	Criticality *fhir.Code `fhirpath:"criticality"`

	// The encounter when the allergy or intolerance was asserted.
	Encounter *fhir.Reference `fhirpath:"encounter"`

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

	// Business identifiers assigned to this AllergyIntolerance by the performer or
	// other systems which remain constant as the resource is updated and
	// propagates from server to server.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// Represents the date and/or time of the last known occurrence of a reaction
	// event.
	LastOccurrence *fhir.DateTime `fhirpath:"lastOccurrence"`

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

	// Additional narrative about the propensity for the Adverse Reaction, not
	// captured in other fields.
	Note []*fhir.Annotation `fhirpath:"note"`

	// Estimated or actual date, date-time, or age when allergy or intolerance was
	// identified.
	Onset fhir.Element `fhirpath:"onset"`

	// The patient who has the allergy or intolerance.
	Patient *fhir.Reference `fhirpath:"patient"`

	// Details about each adverse reaction event linked to exposure to the
	// identified substance.
	Reaction []*AllergyIntoleranceReaction `fhirpath:"reaction"`

	// The recordedDate represents when this particular AllergyIntolerance record
	// was created in the system, which is often a system-generated date.
	RecordedDate *fhir.DateTime `fhirpath:"recordedDate"`

	// Individual who recorded the record and takes responsibility for its content.
	Recorder *fhir.Reference `fhirpath:"recorder"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// Identification of the underlying physiological mechanism for the reaction
	// risk.
	Type *fhir.Code `fhirpath:"type"`

	// Assertion about certainty associated with the propensity, or potential risk,
	// of a reaction to the identified substance (including pharmaceutical
	// product).
	VerificationStatus *fhir.CodeableConcept `fhirpath:"verificationStatus"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetAsserter returns the value of the field Asserter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetAsserter() *fhir.Reference {
	if ai == nil {
		return nil
	}
	return ai.Asserter
}

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetCategory() []*fhir.Code {
	if ai == nil {
		return nil
	}
	return ai.Category
}

// GetClinicalStatus returns the value of the field ClinicalStatus.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetClinicalStatus() *fhir.CodeableConcept {
	if ai == nil {
		return nil
	}
	return ai.ClinicalStatus
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetCode() *fhir.CodeableConcept {
	if ai == nil {
		return nil
	}
	return ai.Code
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetContained() []fhir.Resource {
	if ai == nil {
		return nil
	}
	return ai.Contained
}

// GetCriticality returns the value of the field Criticality.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetCriticality() *fhir.Code {
	if ai == nil {
		return nil
	}
	return ai.Criticality
}

// GetEncounter returns the value of the field Encounter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetEncounter() *fhir.Reference {
	if ai == nil {
		return nil
	}
	return ai.Encounter
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetExtension() []*fhir.Extension {
	if ai == nil {
		return nil
	}
	return ai.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetID() string {
	if ai == nil {
		return ""
	}
	return ai.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetIdentifier() []*fhir.Identifier {
	if ai == nil {
		return nil
	}
	return ai.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetImplicitRules() *fhir.URI {
	if ai == nil {
		return nil
	}
	return ai.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetLanguage() *fhir.Code {
	if ai == nil {
		return nil
	}
	return ai.Language
}

// GetLastOccurrence returns the value of the field LastOccurrence.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetLastOccurrence() *fhir.DateTime {
	if ai == nil {
		return nil
	}
	return ai.LastOccurrence
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetMeta() *fhir.Meta {
	if ai == nil {
		return nil
	}
	return ai.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetModifierExtension() []*fhir.Extension {
	if ai == nil {
		return nil
	}
	return ai.ModifierExtension
}

// GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetNote() []*fhir.Annotation {
	if ai == nil {
		return nil
	}
	return ai.Note
}

// GetOnset returns the value of the field Onset.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetOnset() fhir.Element {
	if ai == nil {
		return nil
	}
	return ai.Onset
}

// GetOnsetDateTime returns the value of the field Onset.
func (ai *AllergyIntolerance) GetOnsetDateTime() *fhir.DateTime {
	if ai == nil {
		return nil
	}
	val, ok := ai.Onset.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
}

// GetOnsetAge returns the value of the field Onset.
func (ai *AllergyIntolerance) GetOnsetAge() *fhir.Age {
	if ai == nil {
		return nil
	}
	val, ok := ai.Onset.(*fhir.Age)
	if !ok {
		return nil
	}
	return val
}

// GetOnsetPeriod returns the value of the field Onset.
func (ai *AllergyIntolerance) GetOnsetPeriod() *fhir.Period {
	if ai == nil {
		return nil
	}
	val, ok := ai.Onset.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
}

// GetOnsetRange returns the value of the field Onset.
func (ai *AllergyIntolerance) GetOnsetRange() *fhir.Range {
	if ai == nil {
		return nil
	}
	val, ok := ai.Onset.(*fhir.Range)
	if !ok {
		return nil
	}
	return val
}

// GetOnsetString returns the value of the field Onset.
func (ai *AllergyIntolerance) GetOnsetString() *fhir.String {
	if ai == nil {
		return nil
	}
	val, ok := ai.Onset.(*fhir.String)
	if !ok {
		return nil
	}
	return val
} // GetPatient returns the value of the field Patient.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetPatient() *fhir.Reference {
	if ai == nil {
		return nil
	}
	return ai.Patient
}

// GetReaction returns the value of the field Reaction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetReaction() []*AllergyIntoleranceReaction {
	if ai == nil {
		return nil
	}
	return ai.Reaction
}

// GetRecordedDate returns the value of the field RecordedDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetRecordedDate() *fhir.DateTime {
	if ai == nil {
		return nil
	}
	return ai.RecordedDate
}

// GetRecorder returns the value of the field Recorder.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetRecorder() *fhir.Reference {
	if ai == nil {
		return nil
	}
	return ai.Recorder
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetText() *fhir.Narrative {
	if ai == nil {
		return nil
	}
	return ai.Text
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetType() *fhir.Code {
	if ai == nil {
		return nil
	}
	return ai.Type
}

// GetVerificationStatus returns the value of the field VerificationStatus.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ai *AllergyIntolerance) GetVerificationStatus() *fhir.CodeableConcept {
	if ai == nil {
		return nil
	}
	return ai.VerificationStatus
}

// Adverse Reaction Events linked to exposure to substance// Details about each adverse reaction event linked to exposure to the
// identified substance.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-AllergyIntolerance.json
type AllergyIntoleranceReaction struct {

	// Text description about the reaction as a whole, including details of the
	// manifestation if required.
	Description *fhir.String `fhirpath:"description"`

	// Identification of the route by which the subject was exposed to the
	// substance.
	ExposureRoute *fhir.CodeableConcept `fhirpath:"exposureRoute"`

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

	// Clinical symptoms and/or signs that are observed or associated with the
	// adverse reaction event.
	Manifestation []*fhir.CodeableConcept `fhirpath:"manifestation"`

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

	// Additional text about the adverse reaction event not captured in other
	// fields.
	Note []*fhir.Annotation `fhirpath:"note"`

	// Record of the date and/or time of the onset of the Reaction.
	Onset *fhir.DateTime `fhirpath:"onset"`

	// Clinical assessment of the severity of the reaction event as a whole,
	// potentially considering multiple different manifestations.
	Severity *fhir.Code `fhirpath:"severity"`

	// Identification of the specific substance (or pharmaceutical product)
	// considered to be responsible for the Adverse Reaction event. Note: the
	// substance for a specific reaction may be different from the substance
	// identified as the cause of the risk, but it must be consistent with it. For
	// instance, it may be a more specific substance (e.g. a brand medication) or a
	// composite product that includes the identified substance. It must be
	// clinically safe to only process the 'code' and ignore the
	// 'reaction.substance'. If a receiving system is unable to confirm that
	// AllergyIntolerance.reaction.substance falls within the semantic scope of
	// AllergyIntolerance.code, then the receiving system should ignore
	// AllergyIntolerance.reaction.substance.
	Substance *fhir.CodeableConcept `fhirpath:"substance"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (air *AllergyIntoleranceReaction) GetDescription() *fhir.String {
	if air == nil {
		return nil
	}
	return air.Description
}

// GetExposureRoute returns the value of the field ExposureRoute.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (air *AllergyIntoleranceReaction) GetExposureRoute() *fhir.CodeableConcept {
	if air == nil {
		return nil
	}
	return air.ExposureRoute
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (air *AllergyIntoleranceReaction) GetExtension() []*fhir.Extension {
	if air == nil {
		return nil
	}
	return air.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (air *AllergyIntoleranceReaction) GetID() string {
	if air == nil {
		return ""
	}
	return air.ID
}

// GetManifestation returns the value of the field Manifestation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (air *AllergyIntoleranceReaction) GetManifestation() []*fhir.CodeableConcept {
	if air == nil {
		return nil
	}
	return air.Manifestation
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (air *AllergyIntoleranceReaction) GetModifierExtension() []*fhir.Extension {
	if air == nil {
		return nil
	}
	return air.ModifierExtension
}

// GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (air *AllergyIntoleranceReaction) GetNote() []*fhir.Annotation {
	if air == nil {
		return nil
	}
	return air.Note
}

// GetOnset returns the value of the field Onset.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (air *AllergyIntoleranceReaction) GetOnset() *fhir.DateTime {
	if air == nil {
		return nil
	}
	return air.Onset
}

// GetSeverity returns the value of the field Severity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (air *AllergyIntoleranceReaction) GetSeverity() *fhir.Code {
	if air == nil {
		return nil
	}
	return air.Severity
}

// GetSubstance returns the value of the field Substance.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (air *AllergyIntoleranceReaction) GetSubstance() *fhir.CodeableConcept {
	if air == nil {
		return nil
	}
	return air.Substance
}

func (ai *AllergyIntolerance) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (ai *AllergyIntolerance) UnmarshalJSON(data []byte) error {
	var raw struct {
		Asserter       *fhir.Reference       `json:"asserter"`
		Category       []*fhir.Code          `json:"category"`
		ClinicalStatus *fhir.CodeableConcept `json:"clinicalStatus"`
		Code           *fhir.CodeableConcept `json:"code"`
		Contained      []fhir.Resource       `json:"contained"`
		Criticality    *fhir.Code            `json:"criticality"`
		Encounter      *fhir.Reference       `json:"encounter"`
		Extension      []*fhir.Extension     `json:"extension"`

		ID                 string                        `json:"id"`
		Identifier         []*fhir.Identifier            `json:"identifier"`
		ImplicitRules      *fhir.URI                     `json:"implicitRules"`
		Language           *fhir.Code                    `json:"language"`
		LastOccurrence     *fhir.DateTime                `json:"lastOccurrence"`
		Meta               *fhir.Meta                    `json:"meta"`
		ModifierExtension  []*fhir.Extension             `json:"modifierExtension"`
		Note               []*fhir.Annotation            `json:"note"`
		OnsetDateTime      *fhir.DateTime                `json:"onsetDateTime"`
		OnsetAge           *fhir.Age                     `json:"onsetAge"`
		OnsetPeriod        *fhir.Period                  `json:"onsetPeriod"`
		OnsetRange         *fhir.Range                   `json:"onsetRange"`
		OnsetString        *fhir.String                  `json:"onsetString"`
		Patient            *fhir.Reference               `json:"patient"`
		Reaction           []*AllergyIntoleranceReaction `json:"reaction"`
		RecordedDate       *fhir.DateTime                `json:"recordedDate"`
		Recorder           *fhir.Reference               `json:"recorder"`
		Text               *fhir.Narrative               `json:"text"`
		Type               *fhir.Code                    `json:"type"`
		VerificationStatus *fhir.CodeableConcept         `json:"verificationStatus"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	ai.Asserter = raw.Asserter
	ai.Category = raw.Category
	ai.ClinicalStatus = raw.ClinicalStatus
	ai.Code = raw.Code
	ai.Contained = raw.Contained
	ai.Criticality = raw.Criticality
	ai.Encounter = raw.Encounter
	ai.Extension = raw.Extension
	ai.ID = raw.ID
	ai.Identifier = raw.Identifier
	ai.ImplicitRules = raw.ImplicitRules
	ai.Language = raw.Language
	ai.LastOccurrence = raw.LastOccurrence
	ai.Meta = raw.Meta
	ai.ModifierExtension = raw.ModifierExtension
	ai.Note = raw.Note
	ai.Onset, err = validate.SelectOneOf[fhir.Element]("AllergyIntolerance.onset",
		raw.OnsetDateTime,
		raw.OnsetAge,
		raw.OnsetPeriod,
		raw.OnsetRange,
		raw.OnsetString)
	if err != nil {
		return err
	}
	ai.Patient = raw.Patient
	ai.Reaction = raw.Reaction
	ai.RecordedDate = raw.RecordedDate
	ai.Recorder = raw.Recorder
	ai.Text = raw.Text
	ai.Type = raw.Type
	ai.VerificationStatus = raw.VerificationStatus
	return nil
}

var _ json.Marshaler = (*AllergyIntolerance)(nil)
var _ json.Unmarshaler = (*AllergyIntolerance)(nil)

func (air *AllergyIntoleranceReaction) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (air *AllergyIntoleranceReaction) UnmarshalJSON(data []byte) error {
	var raw struct {
		Description   *fhir.String          `json:"description"`
		ExposureRoute *fhir.CodeableConcept `json:"exposureRoute"`
		Extension     []*fhir.Extension     `json:"extension"`

		ID                string                  `json:"id"`
		Manifestation     []*fhir.CodeableConcept `json:"manifestation"`
		ModifierExtension []*fhir.Extension       `json:"modifierExtension"`
		Note              []*fhir.Annotation      `json:"note"`
		Onset             *fhir.DateTime          `json:"onset"`
		Severity          *fhir.Code              `json:"severity"`
		Substance         *fhir.CodeableConcept   `json:"substance"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	air.Description = raw.Description
	air.ExposureRoute = raw.ExposureRoute
	air.Extension = raw.Extension
	air.ID = raw.ID
	air.Manifestation = raw.Manifestation
	air.ModifierExtension = raw.ModifierExtension
	air.Note = raw.Note
	air.Onset = raw.Onset
	air.Severity = raw.Severity
	air.Substance = raw.Substance
	return nil
}

var _ json.Marshaler = (*AllergyIntoleranceReaction)(nil)
var _ json.Unmarshaler = (*AllergyIntoleranceReaction)(nil)
