// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package adverseevent

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Actual or potential/avoided event causing unintended physical injury
// resulting from or contributed to by medical care, a research study or other
// healthcare setting factors that requires additional monitoring, treatment,
// or hospitalization, or that results in death.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/AdverseEvent
//   - Source File: StructureDefinition-AdverseEvent.json
type AdverseEvent struct {

	// Whether the event actually happened, or just had the potential to. Note that
	// this is independent of whether anyone was affected or harmed or how
	// severely.
	Actuality *fhir.Code `fhirpath:"actuality"`

	// The overall type of event, intended for search and filtering purposes.
	Category []*fhir.CodeableConcept `fhirpath:"category"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// Parties that may or should contribute or have contributed information to the
	// adverse event, which can consist of one or more activities. Such information
	// includes information leading to the decision to perform the activity and how
	// to perform the activity (e.g. consultant), information that the activity
	// itself seeks to reveal (e.g. informant of clinical history), or information
	// about what activity was performed (e.g. informant witness).
	Contributor []*fhir.Reference `fhirpath:"contributor"`

	// The date (and perhaps time) when the adverse event occurred.
	Date *fhir.DateTime `fhirpath:"date"`

	// Estimated or actual date the AdverseEvent began, in the opinion of the
	// reporter.
	Detected *fhir.DateTime `fhirpath:"detected"`

	// The Encounter during which AdverseEvent was created or to which the creation
	// of this record is tightly associated.
	Encounter *fhir.Reference `fhirpath:"encounter"`

	// This element defines the specific type of event that occurred or that was
	// prevented from occurring.
	Event *fhir.CodeableConcept `fhirpath:"event"`

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

	// Business identifiers assigned to this adverse event by the performer or
	// other systems which remain constant as the resource is updated and
	// propagates from server to server.
	Identifier *fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// The information about where the adverse event occurred.
	Location *fhir.Reference `fhirpath:"location"`

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

	// Describes the type of outcome from the adverse event.
	Outcome *fhir.CodeableConcept `fhirpath:"outcome"`

	// The date on which the existence of the AdverseEvent was first recorded.
	RecordedDate *fhir.DateTime `fhirpath:"recordedDate"`

	// Information on who recorded the adverse event. May be the patient or a
	// practitioner.
	Recorder *fhir.Reference `fhirpath:"recorder"`

	// AdverseEvent.referenceDocument.
	ReferenceDocument []*fhir.Reference `fhirpath:"referenceDocument"`

	// Includes information about the reaction that occurred as a result of
	// exposure to a substance (for example, a drug or a chemical).
	ResultingCondition []*fhir.Reference `fhirpath:"resultingCondition"`

	// Assessment whether this event was of real importance.
	Seriousness *fhir.CodeableConcept `fhirpath:"seriousness"`

	// Describes the severity of the adverse event, in relation to the subject.
	// Contrast to AdverseEvent.seriousness - a severe rash might not be serious,
	// but a mild heart problem is.
	Severity *fhir.CodeableConcept `fhirpath:"severity"`

	// AdverseEvent.study.
	Study []*fhir.Reference `fhirpath:"study"`

	// This subject or group impacted by the event.
	Subject *fhir.Reference `fhirpath:"subject"`

	// AdverseEvent.subjectMedicalHistory.
	SubjectMedicalHistory []*fhir.Reference `fhirpath:"subjectMedicalHistory"`

	// Describes the entity that is suspected to have caused the adverse event.
	SuspectEntity []*AdverseEventSuspectEntity `fhirpath:"suspectEntity"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetActuality returns the value of the field Actuality.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetActuality() *fhir.Code {
	if ae == nil {
		return nil
	}
	return ae.Actuality
}

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetCategory() []*fhir.CodeableConcept {
	if ae == nil {
		return nil
	}
	return ae.Category
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetContained() []fhir.Resource {
	if ae == nil {
		return nil
	}
	return ae.Contained
}

// GetContributor returns the value of the field Contributor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetContributor() []*fhir.Reference {
	if ae == nil {
		return nil
	}
	return ae.Contributor
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetDate() *fhir.DateTime {
	if ae == nil {
		return nil
	}
	return ae.Date
}

// GetDetected returns the value of the field Detected.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetDetected() *fhir.DateTime {
	if ae == nil {
		return nil
	}
	return ae.Detected
}

// GetEncounter returns the value of the field Encounter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetEncounter() *fhir.Reference {
	if ae == nil {
		return nil
	}
	return ae.Encounter
}

// GetEvent returns the value of the field Event.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetEvent() *fhir.CodeableConcept {
	if ae == nil {
		return nil
	}
	return ae.Event
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetExtension() []*fhir.Extension {
	if ae == nil {
		return nil
	}
	return ae.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetID() string {
	if ae == nil {
		return ""
	}
	return ae.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetIdentifier() *fhir.Identifier {
	if ae == nil {
		return nil
	}
	return ae.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetImplicitRules() *fhir.URI {
	if ae == nil {
		return nil
	}
	return ae.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetLanguage() *fhir.Code {
	if ae == nil {
		return nil
	}
	return ae.Language
}

// GetLocation returns the value of the field Location.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetLocation() *fhir.Reference {
	if ae == nil {
		return nil
	}
	return ae.Location
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetMeta() *fhir.Meta {
	if ae == nil {
		return nil
	}
	return ae.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetModifierExtension() []*fhir.Extension {
	if ae == nil {
		return nil
	}
	return ae.ModifierExtension
}

// GetOutcome returns the value of the field Outcome.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetOutcome() *fhir.CodeableConcept {
	if ae == nil {
		return nil
	}
	return ae.Outcome
}

// GetRecordedDate returns the value of the field RecordedDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetRecordedDate() *fhir.DateTime {
	if ae == nil {
		return nil
	}
	return ae.RecordedDate
}

// GetRecorder returns the value of the field Recorder.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetRecorder() *fhir.Reference {
	if ae == nil {
		return nil
	}
	return ae.Recorder
}

// GetReferenceDocument returns the value of the field ReferenceDocument.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetReferenceDocument() []*fhir.Reference {
	if ae == nil {
		return nil
	}
	return ae.ReferenceDocument
}

// GetResultingCondition returns the value of the field ResultingCondition.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetResultingCondition() []*fhir.Reference {
	if ae == nil {
		return nil
	}
	return ae.ResultingCondition
}

// GetSeriousness returns the value of the field Seriousness.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetSeriousness() *fhir.CodeableConcept {
	if ae == nil {
		return nil
	}
	return ae.Seriousness
}

// GetSeverity returns the value of the field Severity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetSeverity() *fhir.CodeableConcept {
	if ae == nil {
		return nil
	}
	return ae.Severity
}

// GetStudy returns the value of the field Study.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetStudy() []*fhir.Reference {
	if ae == nil {
		return nil
	}
	return ae.Study
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetSubject() *fhir.Reference {
	if ae == nil {
		return nil
	}
	return ae.Subject
}

// GetSubjectMedicalHistory returns the value of the field SubjectMedicalHistory.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetSubjectMedicalHistory() []*fhir.Reference {
	if ae == nil {
		return nil
	}
	return ae.SubjectMedicalHistory
}

// GetSuspectEntity returns the value of the field SuspectEntity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetSuspectEntity() []*AdverseEventSuspectEntity {
	if ae == nil {
		return nil
	}
	return ae.SuspectEntity
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ae *AdverseEvent) GetText() *fhir.Narrative {
	if ae == nil {
		return nil
	}
	return ae.Text
}

// The suspected agent causing the adverse event// Describes the entity that is suspected to have caused the adverse event.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-AdverseEvent.json
type AdverseEventSuspectEntity struct {

	// Information on the possible cause of the event.
	Causality []*AdverseEventSuspectEntityCausality `fhirpath:"causality"`

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

	// Identifies the actual instance of what caused the adverse event. May be a
	// substance, medication, medication administration, medication statement or a
	// device.
	Instance *fhir.Reference `fhirpath:"instance"`

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

// GetCausality returns the value of the field Causality.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (aese *AdverseEventSuspectEntity) GetCausality() []*AdverseEventSuspectEntityCausality {
	if aese == nil {
		return nil
	}
	return aese.Causality
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (aese *AdverseEventSuspectEntity) GetExtension() []*fhir.Extension {
	if aese == nil {
		return nil
	}
	return aese.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (aese *AdverseEventSuspectEntity) GetID() string {
	if aese == nil {
		return ""
	}
	return aese.ID
}

// GetInstance returns the value of the field Instance.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (aese *AdverseEventSuspectEntity) GetInstance() *fhir.Reference {
	if aese == nil {
		return nil
	}
	return aese.Instance
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (aese *AdverseEventSuspectEntity) GetModifierExtension() []*fhir.Extension {
	if aese == nil {
		return nil
	}
	return aese.ModifierExtension
}

// Information on the possible cause of the event// Information on the possible cause of the event.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-AdverseEvent.json
type AdverseEventSuspectEntityCausality struct {

	// Assessment of if the entity caused the event.
	Assessment *fhir.CodeableConcept `fhirpath:"assessment"`

	// AdverseEvent.suspectEntity.causalityAuthor.
	Author *fhir.Reference `fhirpath:"author"`

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

	// ProbabilityScale | Bayesian | Checklist.
	Method *fhir.CodeableConcept `fhirpath:"method"`

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

	// AdverseEvent.suspectEntity.causalityProductRelatedness.
	ProductRelatedness *fhir.String `fhirpath:"productRelatedness"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAssessment returns the value of the field Assessment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (aesec *AdverseEventSuspectEntityCausality) GetAssessment() *fhir.CodeableConcept {
	if aesec == nil {
		return nil
	}
	return aesec.Assessment
}

// GetAuthor returns the value of the field Author.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (aesec *AdverseEventSuspectEntityCausality) GetAuthor() *fhir.Reference {
	if aesec == nil {
		return nil
	}
	return aesec.Author
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (aesec *AdverseEventSuspectEntityCausality) GetExtension() []*fhir.Extension {
	if aesec == nil {
		return nil
	}
	return aesec.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (aesec *AdverseEventSuspectEntityCausality) GetID() string {
	if aesec == nil {
		return ""
	}
	return aesec.ID
}

// GetMethod returns the value of the field Method.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (aesec *AdverseEventSuspectEntityCausality) GetMethod() *fhir.CodeableConcept {
	if aesec == nil {
		return nil
	}
	return aesec.Method
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (aesec *AdverseEventSuspectEntityCausality) GetModifierExtension() []*fhir.Extension {
	if aesec == nil {
		return nil
	}
	return aesec.ModifierExtension
}

// GetProductRelatedness returns the value of the field ProductRelatedness.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (aesec *AdverseEventSuspectEntityCausality) GetProductRelatedness() *fhir.String {
	if aesec == nil {
		return nil
	}
	return aesec.ProductRelatedness
}

func (ae *AdverseEvent) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (ae *AdverseEvent) UnmarshalJSON(data []byte) error {
	var raw struct {
		Actuality   *fhir.Code              `json:"actuality"`
		Category    []*fhir.CodeableConcept `json:"category"`
		Contained   []fhir.Resource         `json:"contained"`
		Contributor []*fhir.Reference       `json:"contributor"`
		Date        *fhir.DateTime          `json:"date"`
		Detected    *fhir.DateTime          `json:"detected"`
		Encounter   *fhir.Reference         `json:"encounter"`
		Event       *fhir.CodeableConcept   `json:"event"`
		Extension   []*fhir.Extension       `json:"extension"`

		ID                    string                       `json:"id"`
		Identifier            *fhir.Identifier             `json:"identifier"`
		ImplicitRules         *fhir.URI                    `json:"implicitRules"`
		Language              *fhir.Code                   `json:"language"`
		Location              *fhir.Reference              `json:"location"`
		Meta                  *fhir.Meta                   `json:"meta"`
		ModifierExtension     []*fhir.Extension            `json:"modifierExtension"`
		Outcome               *fhir.CodeableConcept        `json:"outcome"`
		RecordedDate          *fhir.DateTime               `json:"recordedDate"`
		Recorder              *fhir.Reference              `json:"recorder"`
		ReferenceDocument     []*fhir.Reference            `json:"referenceDocument"`
		ResultingCondition    []*fhir.Reference            `json:"resultingCondition"`
		Seriousness           *fhir.CodeableConcept        `json:"seriousness"`
		Severity              *fhir.CodeableConcept        `json:"severity"`
		Study                 []*fhir.Reference            `json:"study"`
		Subject               *fhir.Reference              `json:"subject"`
		SubjectMedicalHistory []*fhir.Reference            `json:"subjectMedicalHistory"`
		SuspectEntity         []*AdverseEventSuspectEntity `json:"suspectEntity"`
		Text                  *fhir.Narrative              `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	ae.Actuality = raw.Actuality
	ae.Category = raw.Category
	ae.Contained = raw.Contained
	ae.Contributor = raw.Contributor
	ae.Date = raw.Date
	ae.Detected = raw.Detected
	ae.Encounter = raw.Encounter
	ae.Event = raw.Event
	ae.Extension = raw.Extension
	ae.ID = raw.ID
	ae.Identifier = raw.Identifier
	ae.ImplicitRules = raw.ImplicitRules
	ae.Language = raw.Language
	ae.Location = raw.Location
	ae.Meta = raw.Meta
	ae.ModifierExtension = raw.ModifierExtension
	ae.Outcome = raw.Outcome
	ae.RecordedDate = raw.RecordedDate
	ae.Recorder = raw.Recorder
	ae.ReferenceDocument = raw.ReferenceDocument
	ae.ResultingCondition = raw.ResultingCondition
	ae.Seriousness = raw.Seriousness
	ae.Severity = raw.Severity
	ae.Study = raw.Study
	ae.Subject = raw.Subject
	ae.SubjectMedicalHistory = raw.SubjectMedicalHistory
	ae.SuspectEntity = raw.SuspectEntity
	ae.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*AdverseEvent)(nil)
var _ json.Unmarshaler = (*AdverseEvent)(nil)

func (aese *AdverseEventSuspectEntity) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (aese *AdverseEventSuspectEntity) UnmarshalJSON(data []byte) error {
	var raw struct {
		Causality []*AdverseEventSuspectEntityCausality `json:"causality"`
		Extension []*fhir.Extension                     `json:"extension"`

		ID                string            `json:"id"`
		Instance          *fhir.Reference   `json:"instance"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	aese.Causality = raw.Causality
	aese.Extension = raw.Extension
	aese.ID = raw.ID
	aese.Instance = raw.Instance
	aese.ModifierExtension = raw.ModifierExtension
	return nil
}

var _ json.Marshaler = (*AdverseEventSuspectEntity)(nil)
var _ json.Unmarshaler = (*AdverseEventSuspectEntity)(nil)

func (aesec *AdverseEventSuspectEntityCausality) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (aesec *AdverseEventSuspectEntityCausality) UnmarshalJSON(data []byte) error {
	var raw struct {
		Assessment *fhir.CodeableConcept `json:"assessment"`
		Author     *fhir.Reference       `json:"author"`
		Extension  []*fhir.Extension     `json:"extension"`

		ID                 string                `json:"id"`
		Method             *fhir.CodeableConcept `json:"method"`
		ModifierExtension  []*fhir.Extension     `json:"modifierExtension"`
		ProductRelatedness *fhir.String          `json:"productRelatedness"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	aesec.Assessment = raw.Assessment
	aesec.Author = raw.Author
	aesec.Extension = raw.Extension
	aesec.ID = raw.ID
	aesec.Method = raw.Method
	aesec.ModifierExtension = raw.ModifierExtension
	aesec.ProductRelatedness = raw.ProductRelatedness
	return nil
}

var _ json.Marshaler = (*AdverseEventSuspectEntityCausality)(nil)
var _ json.Unmarshaler = (*AdverseEventSuspectEntityCausality)(nil)
