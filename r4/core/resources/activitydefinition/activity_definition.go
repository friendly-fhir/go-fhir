// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package activitydefinition

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// This resource allows for the definition of some activity to be performed,
// independent of a particular patient, practitioner, or other performance
// context.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/ActivityDefinition
//   - Source File: StructureDefinition-ActivityDefinition.json
type ActivityDefinition struct {

	// The date on which the resource content was approved by the publisher.
	// Approval happens once when the content is officially approved for usage.
	ApprovalDate *fhir.Date `fhirpath:"approvalDate"`

	// An individiual or organization primarily involved in the creation and
	// maintenance of the content.
	Author []*fhir.ContactDetail `fhirpath:"author"`

	// Indicates the sites on the subject's body where the procedure should be
	// performed (I.e. the target sites).
	BodySite []*fhir.CodeableConcept `fhirpath:"bodySite"`

	// Detailed description of the type of activity; e.g. What lab test, what
	// procedure, what kind of encounter.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// Contact details to assist a user in finding and communicating with the
	// publisher.
	Contact []*fhir.ContactDetail `fhirpath:"contact"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// A copyright statement relating to the activity definition and/or its
	// contents. Copyright statements are generally legal restrictions on the use
	// and publishing of the activity definition.
	Copyright *fhir.Markdown `fhirpath:"copyright"`

	// The date (and optionally time) when the activity definition was published.
	// The date must change when the business version changes and it must change if
	// the status code changes. In addition, it should change when the substantive
	// content of the activity definition changes.
	Date *fhir.DateTime `fhirpath:"date"`

	// A free text natural language description of the activity definition from a
	// consumer's perspective.
	Description *fhir.Markdown `fhirpath:"description"`

	// Set this to true if the definition is to indicate that a particular activity
	// should NOT be performed. If true, this element should be interpreted to
	// reinforce a negative coding. For example NPO as a code with a doNotPerform
	// of true would still indicate to NOT perform the action.
	DoNotPerform *fhir.Boolean `fhirpath:"doNotPerform"`

	// Provides detailed dosage instructions in the same way that they are
	// described for MedicationRequest resources.
	Dosage []*fhir.Dosage `fhirpath:"dosage"`

	// Dynamic values that will be evaluated to produce values for elements of the
	// resulting resource. For example, if the dosage of a medication must be
	// computed based on the patient's weight, a dynamic value would be used to
	// specify an expression that calculated the weight, and the path on the
	// request resource that would contain the result.
	DynamicValue []*ActivityDefinitionDynamicValue `fhirpath:"dynamicValue"`

	// An individual or organization primarily responsible for internal coherence
	// of the content.
	Editor []*fhir.ContactDetail `fhirpath:"editor"`

	// The period during which the activity definition content was or is planned to
	// be in active use.
	EffectivePeriod *fhir.Period `fhirpath:"effectivePeriod"`

	// An individual or organization responsible for officially endorsing the
	// content for use in some setting.
	Endorser []*fhir.ContactDetail `fhirpath:"endorser"`

	// A Boolean value to indicate that this activity definition is authored for
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

	// A formal identifier that is used to identify this activity definition when
	// it is represented in other formats, or referenced in a specification, model,
	// design or an instance.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// Indicates the level of authority/intentionality associated with the activity
	// and where the request should fit into the workflow chain.
	Intent *fhir.Code `fhirpath:"intent"`

	// A legal or geographic region in which the activity definition is intended to
	// be used.
	Jurisdiction []*fhir.CodeableConcept `fhirpath:"jurisdiction"`

	// A description of the kind of resource the activity definition is
	// representing. For example, a MedicationRequest, a ServiceRequest, or a
	// CommunicationRequest. Typically, but not always, this is a Request resource.
	Kind *fhir.Code `fhirpath:"kind"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// The date on which the resource content was last reviewed. Review happens
	// periodically after approval but does not change the original approval date.
	LastReviewDate *fhir.Date `fhirpath:"lastReviewDate"`

	// A reference to a Library resource containing any formal logic used by the
	// activity definition.
	Library []*fhir.Canonical `fhirpath:"library"`

	// Identifies the facility where the activity will occur; e.g. home, hospital,
	// specific clinic, etc.
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

	// A natural language name identifying the activity definition. This name
	// should be usable as an identifier for the module by machine processing
	// applications such as code generation.
	Name *fhir.String `fhirpath:"name"`

	// Defines observation requirements for the action to be performed, such as
	// body weight or surface area.
	ObservationRequirement []*fhir.Reference `fhirpath:"observationRequirement"`

	// Defines the observations that are expected to be produced by the action.
	ObservationResultRequirement []*fhir.Reference `fhirpath:"observationResultRequirement"`

	// Indicates who should participate in performing the action described.
	Participant []*ActivityDefinitionParticipant `fhirpath:"participant"`

	// Indicates how quickly the activity should be addressed with respect to other
	// requests.
	Priority *fhir.Code `fhirpath:"priority"`

	// Identifies the food, drug or other product being consumed or supplied in the
	// activity.
	Product fhir.Element `fhirpath:"product"`

	// A profile to which the target of the activity definition is expected to
	// conform.
	Profile *fhir.Canonical `fhirpath:"profile"`

	// The name of the organization or individual that published the activity
	// definition.
	Publisher *fhir.String `fhirpath:"publisher"`

	// Explanation of why this activity definition is needed and why it has been
	// designed as it has.
	Purpose *fhir.Markdown `fhirpath:"purpose"`

	// Identifies the quantity expected to be consumed at once (per dose, per meal,
	// etc.).
	Quantity *fhir.Quantity `fhirpath:"quantity"`

	// Related artifacts such as additional documentation, justification, or
	// bibliographic references.
	RelatedArtifact []*fhir.RelatedArtifact `fhirpath:"relatedArtifact"`

	// An individual or organization primarily responsible for review of some
	// aspect of the content.
	Reviewer []*fhir.ContactDetail `fhirpath:"reviewer"`

	// Defines specimen requirements for the action to be performed, such as
	// required specimens for a lab test.
	SpecimenRequirement []*fhir.Reference `fhirpath:"specimenRequirement"`

	// The status of this activity definition. Enables tracking the life-cycle of
	// the content.
	Status *fhir.Code `fhirpath:"status"`

	// A code or group definition that describes the intended subject of the
	// activity being defined.
	Subject fhir.Element `fhirpath:"subject"`

	// An explanatory or alternate title for the activity definition giving
	// additional information about its content.
	Subtitle *fhir.String `fhirpath:"subtitle"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// The period, timing or frequency upon which the described activity is to
	// occur.
	Timing fhir.Element `fhirpath:"timing"`

	// A short, descriptive, user-friendly title for the activity definition.
	Title *fhir.String `fhirpath:"title"`

	// Descriptive topics related to the content of the activity. Topics provide a
	// high-level categorization of the activity that can be useful for filtering
	// and searching.
	Topic []*fhir.CodeableConcept `fhirpath:"topic"`

	// A reference to a StructureMap resource that defines a transform that can be
	// executed to produce the intent resource using the ActivityDefinition
	// instance as the input.
	Transform *fhir.Canonical `fhirpath:"transform"`

	// An absolute URI that is used to identify this activity definition when it is
	// referenced in a specification, model, design or an instance; also called its
	// canonical identifier. This SHOULD be globally unique and SHOULD be a literal
	// address at which at which an authoritative instance of this activity
	// definition is (or will be) published. This URL can be the target of a
	// canonical reference. It SHALL remain the same when the activity definition
	// is stored on different servers.
	URL *fhir.URI `fhirpath:"url"`

	// A detailed description of how the activity definition is used from a
	// clinical perspective.
	Usage *fhir.String `fhirpath:"usage"`

	// The content was developed with a focus and intent of supporting the contexts
	// that are listed. These contexts may be general categories (gender, age, ...)
	// or may be references to specific programs (insurance plans, studies, ...)
	// and may be used to assist with indexing and searching for appropriate
	// activity definition instances.
	UseContext []*fhir.UsageContext `fhirpath:"useContext"`

	// The identifier that is used to identify this version of the activity
	// definition when it is referenced in a specification, model, design or
	// instance. This is an arbitrary value managed by the activity definition
	// author and is not expected to be globally unique. For example, it might be a
	// timestamp (e.g. yyyymmdd) if a managed version is not available. There is
	// also no expectation that versions can be placed in a lexicographical
	// sequence. To provide a version consistent with the Decision Support Service
	// specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more
	// information on versioning knowledge assets, refer to the Decision Support
	// Service specification. Note that a version is required for non-experimental
	// active assets.
	Version *fhir.String `fhirpath:"version"`

	profileimpl.BaseActivityDefinition
	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetApprovalDate returns the value of the field ApprovalDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetApprovalDate() *fhir.Date {
	if ad == nil {
		return nil
	}
	return ad.ApprovalDate
}

// GetAuthor returns the value of the field Author.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetAuthor() []*fhir.ContactDetail {
	if ad == nil {
		return nil
	}
	return ad.Author
}

// GetBodySite returns the value of the field BodySite.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetBodySite() []*fhir.CodeableConcept {
	if ad == nil {
		return nil
	}
	return ad.BodySite
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetCode() *fhir.CodeableConcept {
	if ad == nil {
		return nil
	}
	return ad.Code
}

// GetContact returns the value of the field Contact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetContact() []*fhir.ContactDetail {
	if ad == nil {
		return nil
	}
	return ad.Contact
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetContained() []fhir.Resource {
	if ad == nil {
		return nil
	}
	return ad.Contained
}

// GetCopyright returns the value of the field Copyright.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetCopyright() *fhir.Markdown {
	if ad == nil {
		return nil
	}
	return ad.Copyright
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetDate() *fhir.DateTime {
	if ad == nil {
		return nil
	}
	return ad.Date
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetDescription() *fhir.Markdown {
	if ad == nil {
		return nil
	}
	return ad.Description
}

// GetDoNotPerform returns the value of the field DoNotPerform.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetDoNotPerform() *fhir.Boolean {
	if ad == nil {
		return nil
	}
	return ad.DoNotPerform
}

// GetDosage returns the value of the field Dosage.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetDosage() []*fhir.Dosage {
	if ad == nil {
		return nil
	}
	return ad.Dosage
}

// GetDynamicValue returns the value of the field DynamicValue.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetDynamicValue() []*ActivityDefinitionDynamicValue {
	if ad == nil {
		return nil
	}
	return ad.DynamicValue
}

// GetEditor returns the value of the field Editor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetEditor() []*fhir.ContactDetail {
	if ad == nil {
		return nil
	}
	return ad.Editor
}

// GetEffectivePeriod returns the value of the field EffectivePeriod.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetEffectivePeriod() *fhir.Period {
	if ad == nil {
		return nil
	}
	return ad.EffectivePeriod
}

// GetEndorser returns the value of the field Endorser.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetEndorser() []*fhir.ContactDetail {
	if ad == nil {
		return nil
	}
	return ad.Endorser
}

// GetExperimental returns the value of the field Experimental.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetExperimental() *fhir.Boolean {
	if ad == nil {
		return nil
	}
	return ad.Experimental
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetExtension() []*fhir.Extension {
	if ad == nil {
		return nil
	}
	return ad.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetID() string {
	if ad == nil {
		return ""
	}
	return ad.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetIdentifier() []*fhir.Identifier {
	if ad == nil {
		return nil
	}
	return ad.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetImplicitRules() *fhir.URI {
	if ad == nil {
		return nil
	}
	return ad.ImplicitRules
}

// GetIntent returns the value of the field Intent.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetIntent() *fhir.Code {
	if ad == nil {
		return nil
	}
	return ad.Intent
}

// GetJurisdiction returns the value of the field Jurisdiction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetJurisdiction() []*fhir.CodeableConcept {
	if ad == nil {
		return nil
	}
	return ad.Jurisdiction
}

// GetKind returns the value of the field Kind.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetKind() *fhir.Code {
	if ad == nil {
		return nil
	}
	return ad.Kind
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetLanguage() *fhir.Code {
	if ad == nil {
		return nil
	}
	return ad.Language
}

// GetLastReviewDate returns the value of the field LastReviewDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetLastReviewDate() *fhir.Date {
	if ad == nil {
		return nil
	}
	return ad.LastReviewDate
}

// GetLibrary returns the value of the field Library.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetLibrary() []*fhir.Canonical {
	if ad == nil {
		return nil
	}
	return ad.Library
}

// GetLocation returns the value of the field Location.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetLocation() *fhir.Reference {
	if ad == nil {
		return nil
	}
	return ad.Location
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetMeta() *fhir.Meta {
	if ad == nil {
		return nil
	}
	return ad.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetModifierExtension() []*fhir.Extension {
	if ad == nil {
		return nil
	}
	return ad.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetName() *fhir.String {
	if ad == nil {
		return nil
	}
	return ad.Name
}

// GetObservationRequirement returns the value of the field ObservationRequirement.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetObservationRequirement() []*fhir.Reference {
	if ad == nil {
		return nil
	}
	return ad.ObservationRequirement
}

// GetObservationResultRequirement returns the value of the field ObservationResultRequirement.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetObservationResultRequirement() []*fhir.Reference {
	if ad == nil {
		return nil
	}
	return ad.ObservationResultRequirement
}

// GetParticipant returns the value of the field Participant.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetParticipant() []*ActivityDefinitionParticipant {
	if ad == nil {
		return nil
	}
	return ad.Participant
}

// GetPriority returns the value of the field Priority.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetPriority() *fhir.Code {
	if ad == nil {
		return nil
	}
	return ad.Priority
}

// GetProduct returns the value of the field Product.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetProduct() fhir.Element {
	if ad == nil {
		return nil
	}
	return ad.Product
}

// GetProductReference returns the value of the field Product.
func (ad *ActivityDefinition) GetProductReference() *fhir.Reference {
	if ad == nil {
		return nil
	}
	val, ok := ad.Product.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
}

// GetProductCodeableConcept returns the value of the field Product.
func (ad *ActivityDefinition) GetProductCodeableConcept() *fhir.CodeableConcept {
	if ad == nil {
		return nil
	}
	val, ok := ad.Product.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
} // GetProfile returns the value of the field Profile.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetProfile() *fhir.Canonical {
	if ad == nil {
		return nil
	}
	return ad.Profile
}

// GetPublisher returns the value of the field Publisher.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetPublisher() *fhir.String {
	if ad == nil {
		return nil
	}
	return ad.Publisher
}

// GetPurpose returns the value of the field Purpose.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetPurpose() *fhir.Markdown {
	if ad == nil {
		return nil
	}
	return ad.Purpose
}

// GetQuantity returns the value of the field Quantity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetQuantity() *fhir.Quantity {
	if ad == nil {
		return nil
	}
	return ad.Quantity
}

// GetRelatedArtifact returns the value of the field RelatedArtifact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetRelatedArtifact() []*fhir.RelatedArtifact {
	if ad == nil {
		return nil
	}
	return ad.RelatedArtifact
}

// GetReviewer returns the value of the field Reviewer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetReviewer() []*fhir.ContactDetail {
	if ad == nil {
		return nil
	}
	return ad.Reviewer
}

// GetSpecimenRequirement returns the value of the field SpecimenRequirement.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetSpecimenRequirement() []*fhir.Reference {
	if ad == nil {
		return nil
	}
	return ad.SpecimenRequirement
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetStatus() *fhir.Code {
	if ad == nil {
		return nil
	}
	return ad.Status
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetSubject() fhir.Element {
	if ad == nil {
		return nil
	}
	return ad.Subject
}

// GetSubjectCodeableConcept returns the value of the field Subject.
func (ad *ActivityDefinition) GetSubjectCodeableConcept() *fhir.CodeableConcept {
	if ad == nil {
		return nil
	}
	val, ok := ad.Subject.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetSubjectReference returns the value of the field Subject.
func (ad *ActivityDefinition) GetSubjectReference() *fhir.Reference {
	if ad == nil {
		return nil
	}
	val, ok := ad.Subject.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
} // GetSubtitle returns the value of the field Subtitle.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetSubtitle() *fhir.String {
	if ad == nil {
		return nil
	}
	return ad.Subtitle
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetText() *fhir.Narrative {
	if ad == nil {
		return nil
	}
	return ad.Text
}

// GetTiming returns the value of the field Timing.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetTiming() fhir.Element {
	if ad == nil {
		return nil
	}
	return ad.Timing
}

// GetTimingTiming returns the value of the field Timing.
func (ad *ActivityDefinition) GetTimingTiming() *fhir.Timing {
	if ad == nil {
		return nil
	}
	val, ok := ad.Timing.(*fhir.Timing)
	if !ok {
		return nil
	}
	return val
}

// GetTimingDateTime returns the value of the field Timing.
func (ad *ActivityDefinition) GetTimingDateTime() *fhir.DateTime {
	if ad == nil {
		return nil
	}
	val, ok := ad.Timing.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
}

// GetTimingAge returns the value of the field Timing.
func (ad *ActivityDefinition) GetTimingAge() *fhir.Age {
	if ad == nil {
		return nil
	}
	val, ok := ad.Timing.(*fhir.Age)
	if !ok {
		return nil
	}
	return val
}

// GetTimingPeriod returns the value of the field Timing.
func (ad *ActivityDefinition) GetTimingPeriod() *fhir.Period {
	if ad == nil {
		return nil
	}
	val, ok := ad.Timing.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
}

// GetTimingRange returns the value of the field Timing.
func (ad *ActivityDefinition) GetTimingRange() *fhir.Range {
	if ad == nil {
		return nil
	}
	val, ok := ad.Timing.(*fhir.Range)
	if !ok {
		return nil
	}
	return val
}

// GetTimingDuration returns the value of the field Timing.
func (ad *ActivityDefinition) GetTimingDuration() *fhir.Duration {
	if ad == nil {
		return nil
	}
	val, ok := ad.Timing.(*fhir.Duration)
	if !ok {
		return nil
	}
	return val
} // GetTitle returns the value of the field Title.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetTitle() *fhir.String {
	if ad == nil {
		return nil
	}
	return ad.Title
}

// GetTopic returns the value of the field Topic.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetTopic() []*fhir.CodeableConcept {
	if ad == nil {
		return nil
	}
	return ad.Topic
}

// GetTransform returns the value of the field Transform.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetTransform() *fhir.Canonical {
	if ad == nil {
		return nil
	}
	return ad.Transform
}

// GetURL returns the value of the field URL.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetURL() *fhir.URI {
	if ad == nil {
		return nil
	}
	return ad.URL
}

// GetUsage returns the value of the field Usage.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetUsage() *fhir.String {
	if ad == nil {
		return nil
	}
	return ad.Usage
}

// GetUseContext returns the value of the field UseContext.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetUseContext() []*fhir.UsageContext {
	if ad == nil {
		return nil
	}
	return ad.UseContext
}

// GetVersion returns the value of the field Version.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ad *ActivityDefinition) GetVersion() *fhir.String {
	if ad == nil {
		return nil
	}
	return ad.Version
}

// Dynamic aspects of the definition// Dynamic values that will be evaluated to produce values for elements of the
// resulting resource. For example, if the dosage of a medication must be
// computed based on the patient's weight, a dynamic value would be used to
// specify an expression that calculated the weight, and the path on the
// request resource that would contain the result.// Dynamic values are applied in the order in which they are defined in the
// ActivityDefinition. Note that if both a transform and dynamic values are
// specified, the dynamic values will be applied to the result of the
// transform.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-ActivityDefinition.json
type ActivityDefinitionDynamicValue struct {

	// An expression specifying the value of the customized element.
	Expression *fhir.Expression `fhirpath:"expression"`

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

	// The path to the element to be customized. This is the path on the resource
	// that will hold the result of the calculation defined by the expression. The
	// specified path SHALL be a FHIRPath resolveable on the specified target type
	// of the ActivityDefinition, and SHALL consist only of identifiers, constant
	// indexers, and a restricted subset of functions. The path is allowed to
	// contain qualifiers (.) to traverse sub-elements, as well as indexers ([x])
	// to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath
	// Profile](fhirpath.html#simple) for full details).
	Path *fhir.String `fhirpath:"path"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExpression returns the value of the field Expression.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (addv *ActivityDefinitionDynamicValue) GetExpression() *fhir.Expression {
	if addv == nil {
		return nil
	}
	return addv.Expression
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (addv *ActivityDefinitionDynamicValue) GetExtension() []*fhir.Extension {
	if addv == nil {
		return nil
	}
	return addv.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (addv *ActivityDefinitionDynamicValue) GetID() string {
	if addv == nil {
		return ""
	}
	return addv.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (addv *ActivityDefinitionDynamicValue) GetModifierExtension() []*fhir.Extension {
	if addv == nil {
		return nil
	}
	return addv.ModifierExtension
}

// GetPath returns the value of the field Path.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (addv *ActivityDefinitionDynamicValue) GetPath() *fhir.String {
	if addv == nil {
		return nil
	}
	return addv.Path
}

// Who should participate in the action// Indicates who should participate in performing the action described.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-ActivityDefinition.json
type ActivityDefinitionParticipant struct {

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

	// The role the participant should play in performing the described action.
	Role *fhir.CodeableConcept `fhirpath:"role"`

	// The type of participant in the action.
	Type *fhir.Code `fhirpath:"type"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (adp *ActivityDefinitionParticipant) GetExtension() []*fhir.Extension {
	if adp == nil {
		return nil
	}
	return adp.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (adp *ActivityDefinitionParticipant) GetID() string {
	if adp == nil {
		return ""
	}
	return adp.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (adp *ActivityDefinitionParticipant) GetModifierExtension() []*fhir.Extension {
	if adp == nil {
		return nil
	}
	return adp.ModifierExtension
}

// GetRole returns the value of the field Role.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (adp *ActivityDefinitionParticipant) GetRole() *fhir.CodeableConcept {
	if adp == nil {
		return nil
	}
	return adp.Role
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (adp *ActivityDefinitionParticipant) GetType() *fhir.Code {
	if adp == nil {
		return nil
	}
	return adp.Type
}
