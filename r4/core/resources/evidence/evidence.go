// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package evidence

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// The Evidence resource describes the conditional state (population and any
// exposures being compared within the population) and outcome (if specified)
// that the knowledge (evidence, assertion, recommendation) is about.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Evidence
//   - Source File: StructureDefinition-Evidence.json
type Evidence struct {

	// The date on which the resource content was approved by the publisher.
	// Approval happens once when the content is officially approved for usage.
	ApprovalDate *fhir.Date `fhirpath:"approvalDate"`

	// An individiual or organization primarily involved in the creation and
	// maintenance of the content.
	Author []*fhir.ContactDetail `fhirpath:"author"`

	// Contact details to assist a user in finding and communicating with the
	// publisher.
	Contact []*fhir.ContactDetail `fhirpath:"contact"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// A copyright statement relating to the evidence and/or its contents.
	// Copyright statements are generally legal restrictions on the use and
	// publishing of the evidence.
	Copyright *fhir.Markdown `fhirpath:"copyright"`

	// The date (and optionally time) when the evidence was published. The date
	// must change when the business version changes and it must change if the
	// status code changes. In addition, it should change when the substantive
	// content of the evidence changes.
	Date *fhir.DateTime `fhirpath:"date"`

	// A free text natural language description of the evidence from a consumer's
	// perspective.
	Description *fhir.Markdown `fhirpath:"description"`

	// An individual or organization primarily responsible for internal coherence
	// of the content.
	Editor []*fhir.ContactDetail `fhirpath:"editor"`

	// The period during which the evidence content was or is planned to be in
	// active use.
	EffectivePeriod *fhir.Period `fhirpath:"effectivePeriod"`

	// An individual or organization responsible for officially endorsing the
	// content for use in some setting.
	Endorser []*fhir.ContactDetail `fhirpath:"endorser"`

	// A reference to a EvidenceVariable resource that defines the population for
	// the research.
	ExposureBackground *fhir.Reference `fhirpath:"exposureBackground"`

	// A reference to a EvidenceVariable resource that defines the exposure for the
	// research.
	ExposureVariant []*fhir.Reference `fhirpath:"exposureVariant"`

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

	// A formal identifier that is used to identify this evidence when it is
	// represented in other formats, or referenced in a specification, model,
	// design or an instance.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// A legal or geographic region in which the evidence is intended to be used.
	Jurisdiction []*fhir.CodeableConcept `fhirpath:"jurisdiction"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// The date on which the resource content was last reviewed. Review happens
	// periodically after approval but does not change the original approval date.
	LastReviewDate *fhir.Date `fhirpath:"lastReviewDate"`

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

	// A natural language name identifying the evidence. This name should be usable
	// as an identifier for the module by machine processing applications such as
	// code generation.
	Name *fhir.String `fhirpath:"name"`

	// A human-readable string to clarify or explain concepts about the resource.
	Note []*fhir.Annotation `fhirpath:"note"`

	// A reference to a EvidenceVariable resomece that defines the outcome for the
	// research.
	Outcome []*fhir.Reference `fhirpath:"outcome"`

	// The name of the organization or individual that published the evidence.
	Publisher *fhir.String `fhirpath:"publisher"`

	// Related artifacts such as additional documentation, justification, or
	// bibliographic references.
	RelatedArtifact []*fhir.RelatedArtifact `fhirpath:"relatedArtifact"`

	// An individual or organization primarily responsible for review of some
	// aspect of the content.
	Reviewer []*fhir.ContactDetail `fhirpath:"reviewer"`

	// The short title provides an alternate title for use in informal descriptive
	// contexts where the full, formal title is not necessary.
	ShortTitle *fhir.String `fhirpath:"shortTitle"`

	// The status of this evidence. Enables tracking the life-cycle of the content.
	Status *fhir.Code `fhirpath:"status"`

	// An explanatory or alternate title for the Evidence giving additional
	// information about its content.
	Subtitle *fhir.String `fhirpath:"subtitle"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// A short, descriptive, user-friendly title for the evidence.
	Title *fhir.String `fhirpath:"title"`

	// Descriptive topics related to the content of the Evidence. Topics provide a
	// high-level categorization grouping types of Evidences that can be useful for
	// filtering and searching.
	Topic []*fhir.CodeableConcept `fhirpath:"topic"`

	// An absolute URI that is used to identify this evidence when it is referenced
	// in a specification, model, design or an instance; also called its canonical
	// identifier. This SHOULD be globally unique and SHOULD be a literal address
	// at which at which an authoritative instance of this evidence is (or will be)
	// published. This URL can be the target of a canonical reference. It SHALL
	// remain the same when the evidence is stored on different servers.
	URL *fhir.URI `fhirpath:"url"`

	// The content was developed with a focus and intent of supporting the contexts
	// that are listed. These contexts may be general categories (gender, age, ...)
	// or may be references to specific programs (insurance plans, studies, ...)
	// and may be used to assist with indexing and searching for appropriate
	// evidence instances.
	UseContext []*fhir.UsageContext `fhirpath:"useContext"`

	// The identifier that is used to identify this version of the evidence when it
	// is referenced in a specification, model, design or instance. This is an
	// arbitrary value managed by the evidence author and is not expected to be
	// globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a
	// managed version is not available. There is also no expectation that versions
	// can be placed in a lexicographical sequence. To provide a version consistent
	// with the Decision Support Service specification, use the format
	// Major.Minor.Revision (e.g. 1.0.0). For more information on versioning
	// knowledge assets, refer to the Decision Support Service specification. Note
	// that a version is required for non-experimental active artifacts.
	Version *fhir.String `fhirpath:"version"`

	profileimpl.BaseEvidence
	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetApprovalDate returns the value of the field ApprovalDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetApprovalDate() *fhir.Date {
	if e == nil {
		return nil
	}
	return e.ApprovalDate
}

// GetAuthor returns the value of the field Author.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetAuthor() []*fhir.ContactDetail {
	if e == nil {
		return nil
	}
	return e.Author
}

// GetContact returns the value of the field Contact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetContact() []*fhir.ContactDetail {
	if e == nil {
		return nil
	}
	return e.Contact
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetContained() []fhir.Resource {
	if e == nil {
		return nil
	}
	return e.Contained
}

// GetCopyright returns the value of the field Copyright.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetCopyright() *fhir.Markdown {
	if e == nil {
		return nil
	}
	return e.Copyright
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetDate() *fhir.DateTime {
	if e == nil {
		return nil
	}
	return e.Date
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetDescription() *fhir.Markdown {
	if e == nil {
		return nil
	}
	return e.Description
}

// GetEditor returns the value of the field Editor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetEditor() []*fhir.ContactDetail {
	if e == nil {
		return nil
	}
	return e.Editor
}

// GetEffectivePeriod returns the value of the field EffectivePeriod.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetEffectivePeriod() *fhir.Period {
	if e == nil {
		return nil
	}
	return e.EffectivePeriod
}

// GetEndorser returns the value of the field Endorser.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetEndorser() []*fhir.ContactDetail {
	if e == nil {
		return nil
	}
	return e.Endorser
}

// GetExposureBackground returns the value of the field ExposureBackground.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetExposureBackground() *fhir.Reference {
	if e == nil {
		return nil
	}
	return e.ExposureBackground
}

// GetExposureVariant returns the value of the field ExposureVariant.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetExposureVariant() []*fhir.Reference {
	if e == nil {
		return nil
	}
	return e.ExposureVariant
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetExtension() []*fhir.Extension {
	if e == nil {
		return nil
	}
	return e.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetID() string {
	if e == nil {
		return ""
	}
	return e.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetIdentifier() []*fhir.Identifier {
	if e == nil {
		return nil
	}
	return e.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetImplicitRules() *fhir.URI {
	if e == nil {
		return nil
	}
	return e.ImplicitRules
}

// GetJurisdiction returns the value of the field Jurisdiction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetJurisdiction() []*fhir.CodeableConcept {
	if e == nil {
		return nil
	}
	return e.Jurisdiction
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetLanguage() *fhir.Code {
	if e == nil {
		return nil
	}
	return e.Language
}

// GetLastReviewDate returns the value of the field LastReviewDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetLastReviewDate() *fhir.Date {
	if e == nil {
		return nil
	}
	return e.LastReviewDate
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetMeta() *fhir.Meta {
	if e == nil {
		return nil
	}
	return e.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetModifierExtension() []*fhir.Extension {
	if e == nil {
		return nil
	}
	return e.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetName() *fhir.String {
	if e == nil {
		return nil
	}
	return e.Name
}

// GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetNote() []*fhir.Annotation {
	if e == nil {
		return nil
	}
	return e.Note
}

// GetOutcome returns the value of the field Outcome.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetOutcome() []*fhir.Reference {
	if e == nil {
		return nil
	}
	return e.Outcome
}

// GetPublisher returns the value of the field Publisher.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetPublisher() *fhir.String {
	if e == nil {
		return nil
	}
	return e.Publisher
}

// GetRelatedArtifact returns the value of the field RelatedArtifact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetRelatedArtifact() []*fhir.RelatedArtifact {
	if e == nil {
		return nil
	}
	return e.RelatedArtifact
}

// GetReviewer returns the value of the field Reviewer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetReviewer() []*fhir.ContactDetail {
	if e == nil {
		return nil
	}
	return e.Reviewer
}

// GetShortTitle returns the value of the field ShortTitle.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetShortTitle() *fhir.String {
	if e == nil {
		return nil
	}
	return e.ShortTitle
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetStatus() *fhir.Code {
	if e == nil {
		return nil
	}
	return e.Status
}

// GetSubtitle returns the value of the field Subtitle.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetSubtitle() *fhir.String {
	if e == nil {
		return nil
	}
	return e.Subtitle
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetText() *fhir.Narrative {
	if e == nil {
		return nil
	}
	return e.Text
}

// GetTitle returns the value of the field Title.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetTitle() *fhir.String {
	if e == nil {
		return nil
	}
	return e.Title
}

// GetTopic returns the value of the field Topic.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetTopic() []*fhir.CodeableConcept {
	if e == nil {
		return nil
	}
	return e.Topic
}

// GetURL returns the value of the field URL.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetURL() *fhir.URI {
	if e == nil {
		return nil
	}
	return e.URL
}

// GetUseContext returns the value of the field UseContext.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetUseContext() []*fhir.UsageContext {
	if e == nil {
		return nil
	}
	return e.UseContext
}

// GetVersion returns the value of the field Version.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (e *Evidence) GetVersion() *fhir.String {
	if e == nil {
		return nil
	}
	return e.Version
}

func (e *Evidence) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (e *Evidence) UnmarshalJSON(data []byte) error {
	var raw struct {
		ApprovalDate       *fhir.Date            `json:"approvalDate"`
		Author             []*fhir.ContactDetail `json:"author"`
		Contact            []*fhir.ContactDetail `json:"contact"`
		Contained          []fhir.Resource       `json:"contained"`
		Copyright          *fhir.Markdown        `json:"copyright"`
		Date               *fhir.DateTime        `json:"date"`
		Description        *fhir.Markdown        `json:"description"`
		Editor             []*fhir.ContactDetail `json:"editor"`
		EffectivePeriod    *fhir.Period          `json:"effectivePeriod"`
		Endorser           []*fhir.ContactDetail `json:"endorser"`
		ExposureBackground *fhir.Reference       `json:"exposureBackground"`
		ExposureVariant    []*fhir.Reference     `json:"exposureVariant"`
		Extension          []*fhir.Extension     `json:"extension"`

		ID                string                  `json:"id"`
		Identifier        []*fhir.Identifier      `json:"identifier"`
		ImplicitRules     *fhir.URI               `json:"implicitRules"`
		Jurisdiction      []*fhir.CodeableConcept `json:"jurisdiction"`
		Language          *fhir.Code              `json:"language"`
		LastReviewDate    *fhir.Date              `json:"lastReviewDate"`
		Meta              *fhir.Meta              `json:"meta"`
		ModifierExtension []*fhir.Extension       `json:"modifierExtension"`
		Name              *fhir.String            `json:"name"`
		Note              []*fhir.Annotation      `json:"note"`
		Outcome           []*fhir.Reference       `json:"outcome"`
		Publisher         *fhir.String            `json:"publisher"`
		RelatedArtifact   []*fhir.RelatedArtifact `json:"relatedArtifact"`
		Reviewer          []*fhir.ContactDetail   `json:"reviewer"`
		ShortTitle        *fhir.String            `json:"shortTitle"`
		Status            *fhir.Code              `json:"status"`
		Subtitle          *fhir.String            `json:"subtitle"`
		Text              *fhir.Narrative         `json:"text"`
		Title             *fhir.String            `json:"title"`
		Topic             []*fhir.CodeableConcept `json:"topic"`
		URL               *fhir.URI               `json:"url"`
		UseContext        []*fhir.UsageContext    `json:"useContext"`
		Version           *fhir.String            `json:"version"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	e.ApprovalDate = raw.ApprovalDate
	e.Author = raw.Author
	e.Contact = raw.Contact
	e.Contained = raw.Contained
	e.Copyright = raw.Copyright
	e.Date = raw.Date
	e.Description = raw.Description
	e.Editor = raw.Editor
	e.EffectivePeriod = raw.EffectivePeriod
	e.Endorser = raw.Endorser
	e.ExposureBackground = raw.ExposureBackground
	e.ExposureVariant = raw.ExposureVariant
	e.Extension = raw.Extension
	e.ID = raw.ID
	e.Identifier = raw.Identifier
	e.ImplicitRules = raw.ImplicitRules
	e.Jurisdiction = raw.Jurisdiction
	e.Language = raw.Language
	e.LastReviewDate = raw.LastReviewDate
	e.Meta = raw.Meta
	e.ModifierExtension = raw.ModifierExtension
	e.Name = raw.Name
	e.Note = raw.Note
	e.Outcome = raw.Outcome
	e.Publisher = raw.Publisher
	e.RelatedArtifact = raw.RelatedArtifact
	e.Reviewer = raw.Reviewer
	e.ShortTitle = raw.ShortTitle
	e.Status = raw.Status
	e.Subtitle = raw.Subtitle
	e.Text = raw.Text
	e.Title = raw.Title
	e.Topic = raw.Topic
	e.URL = raw.URL
	e.UseContext = raw.UseContext
	e.Version = raw.Version
	return nil
}

var _ json.Marshaler = (*Evidence)(nil)
var _ json.Unmarshaler = (*Evidence)(nil)
