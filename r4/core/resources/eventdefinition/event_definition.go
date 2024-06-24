// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package eventdefinition

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// The EventDefinition resource provides a reusable description of when a
// particular event can occur.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/EventDefinition
//   - Source File: StructureDefinition-EventDefinition.json
type EventDefinition struct {

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

	// A copyright statement relating to the event definition and/or its contents.
	// Copyright statements are generally legal restrictions on the use and
	// publishing of the event definition.
	Copyright *fhir.Markdown `fhirpath:"copyright"`

	// The date (and optionally time) when the event definition was published. The
	// date must change when the business version changes and it must change if the
	// status code changes. In addition, it should change when the substantive
	// content of the event definition changes.
	Date *fhir.DateTime `fhirpath:"date"`

	// A free text natural language description of the event definition from a
	// consumer's perspective.
	Description *fhir.Markdown `fhirpath:"description"`

	// An individual or organization primarily responsible for internal coherence
	// of the content.
	Editor []*fhir.ContactDetail `fhirpath:"editor"`

	// The period during which the event definition content was or is planned to be
	// in active use.
	EffectivePeriod *fhir.Period `fhirpath:"effectivePeriod"`

	// An individual or organization responsible for officially endorsing the
	// content for use in some setting.
	Endorser []*fhir.ContactDetail `fhirpath:"endorser"`

	// A Boolean value to indicate that this event definition is authored for
	// testing purposes (or education/evaluation/marketing) and is not intended to
	// be used for genuine usage.
	Experimental *fhir.Boolean `fhirpath:"experimental"`

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

	// A formal identifier that is used to identify this event definition when it
	// is represented in other formats, or referenced in a specification, model,
	// design or an instance.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// A legal or geographic region in which the event definition is intended to be
	// used.
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

	// A natural language name identifying the event definition. This name should
	// be usable as an identifier for the module by machine processing applications
	// such as code generation.
	Name *fhir.String `fhirpath:"name"`

	// The name of the organization or individual that published the event
	// definition.
	Publisher *fhir.String `fhirpath:"publisher"`

	// Explanation of why this event definition is needed and why it has been
	// designed as it has.
	Purpose *fhir.Markdown `fhirpath:"purpose"`

	// Related resources such as additional documentation, justification, or
	// bibliographic references.
	RelatedArtifact []*fhir.RelatedArtifact `fhirpath:"relatedArtifact"`

	// An individual or organization primarily responsible for review of some
	// aspect of the content.
	Reviewer []*fhir.ContactDetail `fhirpath:"reviewer"`

	// The status of this event definition. Enables tracking the life-cycle of the
	// content.
	Status *fhir.Code `fhirpath:"status"`

	// A code or group definition that describes the intended subject of the event
	// definition.
	Subject fhir.Element `fhirpath:"subject"`

	// An explanatory or alternate title for the event definition giving additional
	// information about its content.
	Subtitle *fhir.String `fhirpath:"subtitle"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// A short, descriptive, user-friendly title for the event definition.
	Title *fhir.String `fhirpath:"title"`

	// Descriptive topics related to the module. Topics provide a high-level
	// categorization of the module that can be useful for filtering and searching.
	Topic []*fhir.CodeableConcept `fhirpath:"topic"`

	// The trigger element defines when the event occurs. If more than one trigger
	// condition is specified, the event fires whenever any one of the trigger
	// conditions is met.
	Trigger []*fhir.TriggerDefinition `fhirpath:"trigger"`

	// An absolute URI that is used to identify this event definition when it is
	// referenced in a specification, model, design or an instance; also called its
	// canonical identifier. This SHOULD be globally unique and SHOULD be a literal
	// address at which at which an authoritative instance of this event definition
	// is (or will be) published. This URL can be the target of a canonical
	// reference. It SHALL remain the same when the event definition is stored on
	// different servers.
	URL *fhir.URI `fhirpath:"url"`

	// A detailed description of how the event definition is used from a clinical
	// perspective.
	Usage *fhir.String `fhirpath:"usage"`

	// The content was developed with a focus and intent of supporting the contexts
	// that are listed. These contexts may be general categories (gender, age, ...)
	// or may be references to specific programs (insurance plans, studies, ...)
	// and may be used to assist with indexing and searching for appropriate event
	// definition instances.
	UseContext []*fhir.UsageContext `fhirpath:"useContext"`

	// The identifier that is used to identify this version of the event definition
	// when it is referenced in a specification, model, design or instance. This is
	// an arbitrary value managed by the event definition author and is not
	// expected to be globally unique. For example, it might be a timestamp (e.g.
	// yyyymmdd) if a managed version is not available. There is also no
	// expectation that versions can be placed in a lexicographical sequence.
	Version *fhir.String `fhirpath:"version"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetApprovalDate returns the value of the field ApprovalDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetApprovalDate() *fhir.Date {
	if ed == nil {
		return nil
	}
	return ed.ApprovalDate
}

// GetAuthor returns the value of the field Author.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetAuthor() []*fhir.ContactDetail {
	if ed == nil {
		return nil
	}
	return ed.Author
}

// GetContact returns the value of the field Contact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetContact() []*fhir.ContactDetail {
	if ed == nil {
		return nil
	}
	return ed.Contact
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetContained() []fhir.Resource {
	if ed == nil {
		return nil
	}
	return ed.Contained
}

// GetCopyright returns the value of the field Copyright.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetCopyright() *fhir.Markdown {
	if ed == nil {
		return nil
	}
	return ed.Copyright
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetDate() *fhir.DateTime {
	if ed == nil {
		return nil
	}
	return ed.Date
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetDescription() *fhir.Markdown {
	if ed == nil {
		return nil
	}
	return ed.Description
}

// GetEditor returns the value of the field Editor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetEditor() []*fhir.ContactDetail {
	if ed == nil {
		return nil
	}
	return ed.Editor
}

// GetEffectivePeriod returns the value of the field EffectivePeriod.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetEffectivePeriod() *fhir.Period {
	if ed == nil {
		return nil
	}
	return ed.EffectivePeriod
}

// GetEndorser returns the value of the field Endorser.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetEndorser() []*fhir.ContactDetail {
	if ed == nil {
		return nil
	}
	return ed.Endorser
}

// GetExperimental returns the value of the field Experimental.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetExperimental() *fhir.Boolean {
	if ed == nil {
		return nil
	}
	return ed.Experimental
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetExtension() []*fhir.Extension {
	if ed == nil {
		return nil
	}
	return ed.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetID() string {
	if ed == nil {
		return ""
	}
	return ed.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetIdentifier() []*fhir.Identifier {
	if ed == nil {
		return nil
	}
	return ed.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetImplicitRules() *fhir.URI {
	if ed == nil {
		return nil
	}
	return ed.ImplicitRules
}

// GetJurisdiction returns the value of the field Jurisdiction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetJurisdiction() []*fhir.CodeableConcept {
	if ed == nil {
		return nil
	}
	return ed.Jurisdiction
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetLanguage() *fhir.Code {
	if ed == nil {
		return nil
	}
	return ed.Language
}

// GetLastReviewDate returns the value of the field LastReviewDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetLastReviewDate() *fhir.Date {
	if ed == nil {
		return nil
	}
	return ed.LastReviewDate
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetMeta() *fhir.Meta {
	if ed == nil {
		return nil
	}
	return ed.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetModifierExtension() []*fhir.Extension {
	if ed == nil {
		return nil
	}
	return ed.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetName() *fhir.String {
	if ed == nil {
		return nil
	}
	return ed.Name
}

// GetPublisher returns the value of the field Publisher.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetPublisher() *fhir.String {
	if ed == nil {
		return nil
	}
	return ed.Publisher
}

// GetPurpose returns the value of the field Purpose.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetPurpose() *fhir.Markdown {
	if ed == nil {
		return nil
	}
	return ed.Purpose
}

// GetRelatedArtifact returns the value of the field RelatedArtifact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetRelatedArtifact() []*fhir.RelatedArtifact {
	if ed == nil {
		return nil
	}
	return ed.RelatedArtifact
}

// GetReviewer returns the value of the field Reviewer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetReviewer() []*fhir.ContactDetail {
	if ed == nil {
		return nil
	}
	return ed.Reviewer
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetStatus() *fhir.Code {
	if ed == nil {
		return nil
	}
	return ed.Status
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetSubject() fhir.Element {
	if ed == nil {
		return nil
	}
	return ed.Subject
}

// GetSubjectCodeableConcept returns the value of the field Subject.
func (ed *EventDefinition) GetSubjectCodeableConcept() *fhir.CodeableConcept {
	if ed == nil {
		return nil
	}
	val, ok := ed.Subject.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetSubjectReference returns the value of the field Subject.
func (ed *EventDefinition) GetSubjectReference() *fhir.Reference {
	if ed == nil {
		return nil
	}
	val, ok := ed.Subject.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
} // GetSubtitle returns the value of the field Subtitle.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetSubtitle() *fhir.String {
	if ed == nil {
		return nil
	}
	return ed.Subtitle
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetText() *fhir.Narrative {
	if ed == nil {
		return nil
	}
	return ed.Text
}

// GetTitle returns the value of the field Title.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetTitle() *fhir.String {
	if ed == nil {
		return nil
	}
	return ed.Title
}

// GetTopic returns the value of the field Topic.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetTopic() []*fhir.CodeableConcept {
	if ed == nil {
		return nil
	}
	return ed.Topic
}

// GetTrigger returns the value of the field Trigger.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetTrigger() []*fhir.TriggerDefinition {
	if ed == nil {
		return nil
	}
	return ed.Trigger
}

// GetURL returns the value of the field URL.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetURL() *fhir.URI {
	if ed == nil {
		return nil
	}
	return ed.URL
}

// GetUsage returns the value of the field Usage.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetUsage() *fhir.String {
	if ed == nil {
		return nil
	}
	return ed.Usage
}

// GetUseContext returns the value of the field UseContext.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetUseContext() []*fhir.UsageContext {
	if ed == nil {
		return nil
	}
	return ed.UseContext
}

// GetVersion returns the value of the field Version.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ed *EventDefinition) GetVersion() *fhir.String {
	if ed == nil {
		return nil
	}
	return ed.Version
}
