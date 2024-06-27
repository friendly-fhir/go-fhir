// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package researchstudy

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// A process where a researcher or organization plans and then executes a
// series of steps intended to increase the field of healthcare-related
// knowledge. This includes studies of safety, efficacy, comparative
// effectiveness and other information about medications, devices, therapies
// and other interventional and investigative techniques. A ResearchStudy
// involves the gathering of information about human or animal subjects.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/ResearchStudy
//   - Source File: StructureDefinition-ResearchStudy.json
type ResearchStudy struct {

	// Describes an expected sequence of events for one of the participants of a
	// study. E.g. Exposure to drug A, wash-out, exposure to drug B, wash-out,
	// follow-up.
	Arm []*ResearchStudyArm `fhirpath:"arm"`

	// Codes categorizing the type of study such as investigational vs.
	// observational, type of blinding, type of randomization, safety vs. efficacy,
	// etc.
	Category []*fhir.CodeableConcept `fhirpath:"category"`

	// The condition that is the focus of the study. For example, In a study to
	// examine risk factors for Lupus, might have as an inclusion criterion
	// "healthy volunteer", but the target condition code would be a Lupus SNOMED
	// code.
	Condition []*fhir.CodeableConcept `fhirpath:"condition"`

	// Contact details to assist a user in learning more about or engaging with the
	// study.
	Contact []*fhir.ContactDetail `fhirpath:"contact"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// A full description of how the study is being conducted.
	Description *fhir.Markdown `fhirpath:"description"`

	// Reference to a Group that defines the criteria for and quantity of subjects
	// participating in the study. E.g. " 200 female Europeans between the ages of
	// 20 and 45 with early onset diabetes".
	Enrollment []*fhir.Reference `fhirpath:"enrollment"`

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// The medication(s), food(s), therapy(ies), device(s) or other concerns or
	// interventions that the study is seeking to gain more information about.
	Focus []*fhir.CodeableConcept `fhirpath:"focus"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// Identifiers assigned to this research study by the sponsor or other systems.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// Key terms to aid in searching for or filtering the study.
	Keyword []*fhir.CodeableConcept `fhirpath:"keyword"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// Indicates a country, state or other region where the study is taking place.
	Location []*fhir.CodeableConcept `fhirpath:"location"`

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

	// Comments made about the study by the performer, subject or other
	// participants.
	Note []*fhir.Annotation `fhirpath:"note"`

	// A goal that the study is aiming to achieve in terms of a scientific question
	// to be answered by the analysis of data collected during the study.
	Objective []*ResearchStudyObjective `fhirpath:"objective"`

	// A larger research study of which this particular study is a component or
	// step.
	PartOf []*fhir.Reference `fhirpath:"partOf"`

	// Identifies the start date and the expected (or actual, depending on status)
	// end date for the study.
	Period *fhir.Period `fhirpath:"period"`

	// The stage in the progression of a therapy from initial experimental use in
	// humans in clinical trials to post-market evaluation.
	Phase *fhir.CodeableConcept `fhirpath:"phase"`

	// The type of study based upon the intent of the study's activities. A
	// classification of the intent of the study.
	PrimaryPurposeType *fhir.CodeableConcept `fhirpath:"primaryPurposeType"`

	// A researcher in a study who oversees multiple aspects of the study, such as
	// concept development, protocol writing, protocol submission for IRB approval,
	// participant recruitment, informed consent, data collection, analysis,
	// interpretation and presentation.
	PrincipalInvestigator *fhir.Reference `fhirpath:"principalInvestigator"`

	// The set of steps expected to be performed as part of the execution of the
	// study.
	Protocol []*fhir.Reference `fhirpath:"protocol"`

	// A description and/or code explaining the premature termination of the study.
	ReasonStopped *fhir.CodeableConcept `fhirpath:"reasonStopped"`

	// Citations, references and other related documents.
	RelatedArtifact []*fhir.RelatedArtifact `fhirpath:"relatedArtifact"`

	// A facility in which study activities are conducted.
	Site []*fhir.Reference `fhirpath:"site"`

	// An organization that initiates the investigation and is legally responsible
	// for the study.
	Sponsor *fhir.Reference `fhirpath:"sponsor"`

	// The current state of the study.
	Status *fhir.Code `fhirpath:"status"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// A short, descriptive user-friendly label for the study.
	Title *fhir.String `fhirpath:"title"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetArm returns the value of the field Arm.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetArm() []*ResearchStudyArm {
	if rs == nil {
		return nil
	}
	return rs.Arm
}

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetCategory() []*fhir.CodeableConcept {
	if rs == nil {
		return nil
	}
	return rs.Category
}

// GetCondition returns the value of the field Condition.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetCondition() []*fhir.CodeableConcept {
	if rs == nil {
		return nil
	}
	return rs.Condition
}

// GetContact returns the value of the field Contact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetContact() []*fhir.ContactDetail {
	if rs == nil {
		return nil
	}
	return rs.Contact
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetContained() []fhir.Resource {
	if rs == nil {
		return nil
	}
	return rs.Contained
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetDescription() *fhir.Markdown {
	if rs == nil {
		return nil
	}
	return rs.Description
}

// GetEnrollment returns the value of the field Enrollment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetEnrollment() []*fhir.Reference {
	if rs == nil {
		return nil
	}
	return rs.Enrollment
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetExtension() []*fhir.Extension {
	if rs == nil {
		return nil
	}
	return rs.Extension
}

// GetFocus returns the value of the field Focus.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetFocus() []*fhir.CodeableConcept {
	if rs == nil {
		return nil
	}
	return rs.Focus
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetID() string {
	if rs == nil {
		return ""
	}
	return rs.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetIdentifier() []*fhir.Identifier {
	if rs == nil {
		return nil
	}
	return rs.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetImplicitRules() *fhir.URI {
	if rs == nil {
		return nil
	}
	return rs.ImplicitRules
}

// GetKeyword returns the value of the field Keyword.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetKeyword() []*fhir.CodeableConcept {
	if rs == nil {
		return nil
	}
	return rs.Keyword
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetLanguage() *fhir.Code {
	if rs == nil {
		return nil
	}
	return rs.Language
}

// GetLocation returns the value of the field Location.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetLocation() []*fhir.CodeableConcept {
	if rs == nil {
		return nil
	}
	return rs.Location
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetMeta() *fhir.Meta {
	if rs == nil {
		return nil
	}
	return rs.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetModifierExtension() []*fhir.Extension {
	if rs == nil {
		return nil
	}
	return rs.ModifierExtension
}

// GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetNote() []*fhir.Annotation {
	if rs == nil {
		return nil
	}
	return rs.Note
}

// GetObjective returns the value of the field Objective.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetObjective() []*ResearchStudyObjective {
	if rs == nil {
		return nil
	}
	return rs.Objective
}

// GetPartOf returns the value of the field PartOf.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetPartOf() []*fhir.Reference {
	if rs == nil {
		return nil
	}
	return rs.PartOf
}

// GetPeriod returns the value of the field Period.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetPeriod() *fhir.Period {
	if rs == nil {
		return nil
	}
	return rs.Period
}

// GetPhase returns the value of the field Phase.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetPhase() *fhir.CodeableConcept {
	if rs == nil {
		return nil
	}
	return rs.Phase
}

// GetPrimaryPurposeType returns the value of the field PrimaryPurposeType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetPrimaryPurposeType() *fhir.CodeableConcept {
	if rs == nil {
		return nil
	}
	return rs.PrimaryPurposeType
}

// GetPrincipalInvestigator returns the value of the field PrincipalInvestigator.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetPrincipalInvestigator() *fhir.Reference {
	if rs == nil {
		return nil
	}
	return rs.PrincipalInvestigator
}

// GetProtocol returns the value of the field Protocol.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetProtocol() []*fhir.Reference {
	if rs == nil {
		return nil
	}
	return rs.Protocol
}

// GetReasonStopped returns the value of the field ReasonStopped.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetReasonStopped() *fhir.CodeableConcept {
	if rs == nil {
		return nil
	}
	return rs.ReasonStopped
}

// GetRelatedArtifact returns the value of the field RelatedArtifact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetRelatedArtifact() []*fhir.RelatedArtifact {
	if rs == nil {
		return nil
	}
	return rs.RelatedArtifact
}

// GetSite returns the value of the field Site.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetSite() []*fhir.Reference {
	if rs == nil {
		return nil
	}
	return rs.Site
}

// GetSponsor returns the value of the field Sponsor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetSponsor() *fhir.Reference {
	if rs == nil {
		return nil
	}
	return rs.Sponsor
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetStatus() *fhir.Code {
	if rs == nil {
		return nil
	}
	return rs.Status
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetText() *fhir.Narrative {
	if rs == nil {
		return nil
	}
	return rs.Text
}

// GetTitle returns the value of the field Title.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rs *ResearchStudy) GetTitle() *fhir.String {
	if rs == nil {
		return nil
	}
	return rs.Title
}

// Defined path through the study for a subject// Describes an expected sequence of events for one of the participants of a
// study. E.g. Exposure to drug A, wash-out, exposure to drug B, wash-out,
// follow-up.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-ResearchStudy.json
type ResearchStudyArm struct {

	// A succinct description of the path through the study that would be followed
	// by a subject adhering to this arm.
	Description *fhir.String `fhirpath:"description"`

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

	// Unique, human-readable label for this arm of the study.
	Name *fhir.String `fhirpath:"name"`

	// Categorization of study arm, e.g. experimental, active comparator, placebo
	// comparater.
	Type *fhir.CodeableConcept `fhirpath:"type"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rsa *ResearchStudyArm) GetDescription() *fhir.String {
	if rsa == nil {
		return nil
	}
	return rsa.Description
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rsa *ResearchStudyArm) GetExtension() []*fhir.Extension {
	if rsa == nil {
		return nil
	}
	return rsa.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rsa *ResearchStudyArm) GetID() string {
	if rsa == nil {
		return ""
	}
	return rsa.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rsa *ResearchStudyArm) GetModifierExtension() []*fhir.Extension {
	if rsa == nil {
		return nil
	}
	return rsa.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rsa *ResearchStudyArm) GetName() *fhir.String {
	if rsa == nil {
		return nil
	}
	return rsa.Name
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rsa *ResearchStudyArm) GetType() *fhir.CodeableConcept {
	if rsa == nil {
		return nil
	}
	return rsa.Type
}

// A goal for the study// A goal that the study is aiming to achieve in terms of a scientific question
// to be answered by the analysis of data collected during the study.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-ResearchStudy.json
type ResearchStudyObjective struct {

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

	// Unique, human-readable label for this objective of the study.
	Name *fhir.String `fhirpath:"name"`

	// The kind of study objective.
	Type *fhir.CodeableConcept `fhirpath:"type"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rso *ResearchStudyObjective) GetExtension() []*fhir.Extension {
	if rso == nil {
		return nil
	}
	return rso.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rso *ResearchStudyObjective) GetID() string {
	if rso == nil {
		return ""
	}
	return rso.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rso *ResearchStudyObjective) GetModifierExtension() []*fhir.Extension {
	if rso == nil {
		return nil
	}
	return rso.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rso *ResearchStudyObjective) GetName() *fhir.String {
	if rso == nil {
		return nil
	}
	return rso.Name
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rso *ResearchStudyObjective) GetType() *fhir.CodeableConcept {
	if rso == nil {
		return nil
	}
	return rso.Type
}

func (rs *ResearchStudy) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (rs *ResearchStudy) UnmarshalJSON(data []byte) error {
	var raw struct {
		Arm         []*ResearchStudyArm     `json:"arm"`
		Category    []*fhir.CodeableConcept `json:"category"`
		Condition   []*fhir.CodeableConcept `json:"condition"`
		Contact     []*fhir.ContactDetail   `json:"contact"`
		Contained   []fhir.Resource         `json:"contained"`
		Description *fhir.Markdown          `json:"description"`
		Enrollment  []*fhir.Reference       `json:"enrollment"`
		Extension   []*fhir.Extension       `json:"extension"`
		Focus       []*fhir.CodeableConcept `json:"focus"`

		ID                    string                    `json:"id"`
		Identifier            []*fhir.Identifier        `json:"identifier"`
		ImplicitRules         *fhir.URI                 `json:"implicitRules"`
		Keyword               []*fhir.CodeableConcept   `json:"keyword"`
		Language              *fhir.Code                `json:"language"`
		Location              []*fhir.CodeableConcept   `json:"location"`
		Meta                  *fhir.Meta                `json:"meta"`
		ModifierExtension     []*fhir.Extension         `json:"modifierExtension"`
		Note                  []*fhir.Annotation        `json:"note"`
		Objective             []*ResearchStudyObjective `json:"objective"`
		PartOf                []*fhir.Reference         `json:"partOf"`
		Period                *fhir.Period              `json:"period"`
		Phase                 *fhir.CodeableConcept     `json:"phase"`
		PrimaryPurposeType    *fhir.CodeableConcept     `json:"primaryPurposeType"`
		PrincipalInvestigator *fhir.Reference           `json:"principalInvestigator"`
		Protocol              []*fhir.Reference         `json:"protocol"`
		ReasonStopped         *fhir.CodeableConcept     `json:"reasonStopped"`
		RelatedArtifact       []*fhir.RelatedArtifact   `json:"relatedArtifact"`
		Site                  []*fhir.Reference         `json:"site"`
		Sponsor               *fhir.Reference           `json:"sponsor"`
		Status                *fhir.Code                `json:"status"`
		Text                  *fhir.Narrative           `json:"text"`
		Title                 *fhir.String              `json:"title"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	rs.Arm = raw.Arm
	rs.Category = raw.Category
	rs.Condition = raw.Condition
	rs.Contact = raw.Contact
	rs.Contained = raw.Contained
	rs.Description = raw.Description
	rs.Enrollment = raw.Enrollment
	rs.Extension = raw.Extension
	rs.Focus = raw.Focus
	rs.ID = raw.ID
	rs.Identifier = raw.Identifier
	rs.ImplicitRules = raw.ImplicitRules
	rs.Keyword = raw.Keyword
	rs.Language = raw.Language
	rs.Location = raw.Location
	rs.Meta = raw.Meta
	rs.ModifierExtension = raw.ModifierExtension
	rs.Note = raw.Note
	rs.Objective = raw.Objective
	rs.PartOf = raw.PartOf
	rs.Period = raw.Period
	rs.Phase = raw.Phase
	rs.PrimaryPurposeType = raw.PrimaryPurposeType
	rs.PrincipalInvestigator = raw.PrincipalInvestigator
	rs.Protocol = raw.Protocol
	rs.ReasonStopped = raw.ReasonStopped
	rs.RelatedArtifact = raw.RelatedArtifact
	rs.Site = raw.Site
	rs.Sponsor = raw.Sponsor
	rs.Status = raw.Status
	rs.Text = raw.Text
	rs.Title = raw.Title
	return nil
}

var _ json.Marshaler = (*ResearchStudy)(nil)
var _ json.Unmarshaler = (*ResearchStudy)(nil)

func (rsa *ResearchStudyArm) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (rsa *ResearchStudyArm) UnmarshalJSON(data []byte) error {
	var raw struct {
		Description *fhir.String      `json:"description"`
		Extension   []*fhir.Extension `json:"extension"`

		ID                string                `json:"id"`
		ModifierExtension []*fhir.Extension     `json:"modifierExtension"`
		Name              *fhir.String          `json:"name"`
		Type              *fhir.CodeableConcept `json:"type"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	rsa.Description = raw.Description
	rsa.Extension = raw.Extension
	rsa.ID = raw.ID
	rsa.ModifierExtension = raw.ModifierExtension
	rsa.Name = raw.Name
	rsa.Type = raw.Type
	return nil
}

var _ json.Marshaler = (*ResearchStudyArm)(nil)
var _ json.Unmarshaler = (*ResearchStudyArm)(nil)

func (rso *ResearchStudyObjective) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (rso *ResearchStudyObjective) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                string                `json:"id"`
		ModifierExtension []*fhir.Extension     `json:"modifierExtension"`
		Name              *fhir.String          `json:"name"`
		Type              *fhir.CodeableConcept `json:"type"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	rso.Extension = raw.Extension
	rso.ID = raw.ID
	rso.ModifierExtension = raw.ModifierExtension
	rso.Name = raw.Name
	rso.Type = raw.Type
	return nil
}

var _ json.Marshaler = (*ResearchStudyObjective)(nil)
var _ json.Unmarshaler = (*ResearchStudyObjective)(nil)
