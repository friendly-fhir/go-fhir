// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package researchelementdefinition

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// The ResearchElementDefinition resource describes a "PICO" element that
// knowledge (evidence, assertion, recommendation) is about.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/ResearchElementDefinition
//   - Source File: StructureDefinition-ResearchElementDefinition.json
type ResearchElementDefinition struct {

	// The date on which the resource content was approved by the publisher.
	// Approval happens once when the content is officially approved for usage.
	ApprovalDate *fhir.Date `fhirpath:"approvalDate"`

	// An individiual or organization primarily involved in the creation and
	// maintenance of the content.
	Author []*fhir.ContactDetail `fhirpath:"author"`

	// A characteristic that defines the members of the research element. Multiple
	// characteristics are applied with "and" semantics.
	Characteristic []*ResearchElementDefinitionCharacteristic `fhirpath:"characteristic"`

	// A human-readable string to clarify or explain concepts about the resource.
	Comment []*fhir.String `fhirpath:"comment"`

	// Contact details to assist a user in finding and communicating with the
	// publisher.
	Contact []*fhir.ContactDetail `fhirpath:"contact"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// A copyright statement relating to the research element definition and/or its
	// contents. Copyright statements are generally legal restrictions on the use
	// and publishing of the research element definition.
	Copyright *fhir.Markdown `fhirpath:"copyright"`

	// The date (and optionally time) when the research element definition was
	// published. The date must change when the business version changes and it
	// must change if the status code changes. In addition, it should change when
	// the substantive content of the research element definition changes.
	Date *fhir.DateTime `fhirpath:"date"`

	// A free text natural language description of the research element definition
	// from a consumer's perspective.
	Description *fhir.Markdown `fhirpath:"description"`

	// An individual or organization primarily responsible for internal coherence
	// of the content.
	Editor []*fhir.ContactDetail `fhirpath:"editor"`

	// The period during which the research element definition content was or is
	// planned to be in active use.
	EffectivePeriod *fhir.Period `fhirpath:"effectivePeriod"`

	// An individual or organization responsible for officially endorsing the
	// content for use in some setting.
	Endorser []*fhir.ContactDetail `fhirpath:"endorser"`

	// A Boolean value to indicate that this research element definition is
	// authored for testing purposes (or education/evaluation/marketing) and is not
	// intended to be used for genuine usage.
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

	// A formal identifier that is used to identify this research element
	// definition when it is represented in other formats, or referenced in a
	// specification, model, design or an instance.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// A legal or geographic region in which the research element definition is
	// intended to be used.
	Jurisdiction []*fhir.CodeableConcept `fhirpath:"jurisdiction"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// The date on which the resource content was last reviewed. Review happens
	// periodically after approval but does not change the original approval date.
	LastReviewDate *fhir.Date `fhirpath:"lastReviewDate"`

	// A reference to a Library resource containing the formal logic used by the
	// ResearchElementDefinition.
	Library []*fhir.Canonical `fhirpath:"library"`

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

	// A natural language name identifying the research element definition. This
	// name should be usable as an identifier for the module by machine processing
	// applications such as code generation.
	Name *fhir.String `fhirpath:"name"`

	// The name of the organization or individual that published the research
	// element definition.
	Publisher *fhir.String `fhirpath:"publisher"`

	// Explanation of why this research element definition is needed and why it has
	// been designed as it has.
	Purpose *fhir.Markdown `fhirpath:"purpose"`

	// Related artifacts such as additional documentation, justification, or
	// bibliographic references.
	RelatedArtifact []*fhir.RelatedArtifact `fhirpath:"relatedArtifact"`

	// An individual or organization primarily responsible for review of some
	// aspect of the content.
	Reviewer []*fhir.ContactDetail `fhirpath:"reviewer"`

	// The short title provides an alternate title for use in informal descriptive
	// contexts where the full, formal title is not necessary.
	ShortTitle *fhir.String `fhirpath:"shortTitle"`

	// The status of this research element definition. Enables tracking the
	// life-cycle of the content.
	Status *fhir.Code `fhirpath:"status"`

	// The intended subjects for the ResearchElementDefinition. If this element is
	// not provided, a Patient subject is assumed, but the subject of the
	// ResearchElementDefinition can be anything.
	Subject fhir.Element `fhirpath:"subject"`

	// An explanatory or alternate title for the ResearchElementDefinition giving
	// additional information about its content.
	Subtitle *fhir.String `fhirpath:"subtitle"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// A short, descriptive, user-friendly title for the research element
	// definition.
	Title *fhir.String `fhirpath:"title"`

	// Descriptive topics related to the content of the ResearchElementDefinition.
	// Topics provide a high-level categorization grouping types of
	// ResearchElementDefinitions that can be useful for filtering and searching.
	Topic []*fhir.CodeableConcept `fhirpath:"topic"`

	// The type of research element, a population, an exposure, or an outcome.
	Type *fhir.Code `fhirpath:"type"`

	// An absolute URI that is used to identify this research element definition
	// when it is referenced in a specification, model, design or an instance; also
	// called its canonical identifier. This SHOULD be globally unique and SHOULD
	// be a literal address at which at which an authoritative instance of this
	// research element definition is (or will be) published. This URL can be the
	// target of a canonical reference. It SHALL remain the same when the research
	// element definition is stored on different servers.
	URL *fhir.URI `fhirpath:"url"`

	// A detailed description, from a clinical perspective, of how the
	// ResearchElementDefinition is used.
	Usage *fhir.String `fhirpath:"usage"`

	// The content was developed with a focus and intent of supporting the contexts
	// that are listed. These contexts may be general categories (gender, age, ...)
	// or may be references to specific programs (insurance plans, studies, ...)
	// and may be used to assist with indexing and searching for appropriate
	// research element definition instances.
	UseContext []*fhir.UsageContext `fhirpath:"useContext"`

	// The type of the outcome (e.g. Dichotomous, Continuous, or Descriptive).
	VariableType *fhir.Code `fhirpath:"variableType"`

	// The identifier that is used to identify this version of the research element
	// definition when it is referenced in a specification, model, design or
	// instance. This is an arbitrary value managed by the research element
	// definition author and is not expected to be globally unique. For example, it
	// might be a timestamp (e.g. yyyymmdd) if a managed version is not available.
	// There is also no expectation that versions can be placed in a
	// lexicographical sequence. To provide a version consistent with the Decision
	// Support Service specification, use the format Major.Minor.Revision (e.g.
	// 1.0.0). For more information on versioning knowledge assets, refer to the
	// Decision Support Service specification. Note that a version is required for
	// non-experimental active artifacts.
	Version *fhir.String `fhirpath:"version"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetApprovalDate returns the value of the field ApprovalDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetApprovalDate() *fhir.Date {
	if red == nil {
		return nil
	}
	return red.ApprovalDate
}

// GetAuthor returns the value of the field Author.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetAuthor() []*fhir.ContactDetail {
	if red == nil {
		return nil
	}
	return red.Author
}

// GetCharacteristic returns the value of the field Characteristic.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetCharacteristic() []*ResearchElementDefinitionCharacteristic {
	if red == nil {
		return nil
	}
	return red.Characteristic
}

// GetComment returns the value of the field Comment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetComment() []*fhir.String {
	if red == nil {
		return nil
	}
	return red.Comment
}

// GetContact returns the value of the field Contact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetContact() []*fhir.ContactDetail {
	if red == nil {
		return nil
	}
	return red.Contact
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetContained() []fhir.Resource {
	if red == nil {
		return nil
	}
	return red.Contained
}

// GetCopyright returns the value of the field Copyright.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetCopyright() *fhir.Markdown {
	if red == nil {
		return nil
	}
	return red.Copyright
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetDate() *fhir.DateTime {
	if red == nil {
		return nil
	}
	return red.Date
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetDescription() *fhir.Markdown {
	if red == nil {
		return nil
	}
	return red.Description
}

// GetEditor returns the value of the field Editor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetEditor() []*fhir.ContactDetail {
	if red == nil {
		return nil
	}
	return red.Editor
}

// GetEffectivePeriod returns the value of the field EffectivePeriod.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetEffectivePeriod() *fhir.Period {
	if red == nil {
		return nil
	}
	return red.EffectivePeriod
}

// GetEndorser returns the value of the field Endorser.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetEndorser() []*fhir.ContactDetail {
	if red == nil {
		return nil
	}
	return red.Endorser
}

// GetExperimental returns the value of the field Experimental.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetExperimental() *fhir.Boolean {
	if red == nil {
		return nil
	}
	return red.Experimental
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetExtension() []*fhir.Extension {
	if red == nil {
		return nil
	}
	return red.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetID() string {
	if red == nil {
		return ""
	}
	return red.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetIdentifier() []*fhir.Identifier {
	if red == nil {
		return nil
	}
	return red.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetImplicitRules() *fhir.URI {
	if red == nil {
		return nil
	}
	return red.ImplicitRules
}

// GetJurisdiction returns the value of the field Jurisdiction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetJurisdiction() []*fhir.CodeableConcept {
	if red == nil {
		return nil
	}
	return red.Jurisdiction
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetLanguage() *fhir.Code {
	if red == nil {
		return nil
	}
	return red.Language
}

// GetLastReviewDate returns the value of the field LastReviewDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetLastReviewDate() *fhir.Date {
	if red == nil {
		return nil
	}
	return red.LastReviewDate
}

// GetLibrary returns the value of the field Library.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetLibrary() []*fhir.Canonical {
	if red == nil {
		return nil
	}
	return red.Library
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetMeta() *fhir.Meta {
	if red == nil {
		return nil
	}
	return red.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetModifierExtension() []*fhir.Extension {
	if red == nil {
		return nil
	}
	return red.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetName() *fhir.String {
	if red == nil {
		return nil
	}
	return red.Name
}

// GetPublisher returns the value of the field Publisher.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetPublisher() *fhir.String {
	if red == nil {
		return nil
	}
	return red.Publisher
}

// GetPurpose returns the value of the field Purpose.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetPurpose() *fhir.Markdown {
	if red == nil {
		return nil
	}
	return red.Purpose
}

// GetRelatedArtifact returns the value of the field RelatedArtifact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetRelatedArtifact() []*fhir.RelatedArtifact {
	if red == nil {
		return nil
	}
	return red.RelatedArtifact
}

// GetReviewer returns the value of the field Reviewer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetReviewer() []*fhir.ContactDetail {
	if red == nil {
		return nil
	}
	return red.Reviewer
}

// GetShortTitle returns the value of the field ShortTitle.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetShortTitle() *fhir.String {
	if red == nil {
		return nil
	}
	return red.ShortTitle
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetStatus() *fhir.Code {
	if red == nil {
		return nil
	}
	return red.Status
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetSubject() fhir.Element {
	if red == nil {
		return nil
	}
	return red.Subject
}

// GetSubjectCodeableConcept returns the value of the field Subject.
func (red *ResearchElementDefinition) GetSubjectCodeableConcept() *fhir.CodeableConcept {
	if red == nil {
		return nil
	}
	val, ok := red.Subject.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetSubjectReference returns the value of the field Subject.
func (red *ResearchElementDefinition) GetSubjectReference() *fhir.Reference {
	if red == nil {
		return nil
	}
	val, ok := red.Subject.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
} // GetSubtitle returns the value of the field Subtitle.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetSubtitle() *fhir.String {
	if red == nil {
		return nil
	}
	return red.Subtitle
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetText() *fhir.Narrative {
	if red == nil {
		return nil
	}
	return red.Text
}

// GetTitle returns the value of the field Title.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetTitle() *fhir.String {
	if red == nil {
		return nil
	}
	return red.Title
}

// GetTopic returns the value of the field Topic.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetTopic() []*fhir.CodeableConcept {
	if red == nil {
		return nil
	}
	return red.Topic
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetType() *fhir.Code {
	if red == nil {
		return nil
	}
	return red.Type
}

// GetURL returns the value of the field URL.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetURL() *fhir.URI {
	if red == nil {
		return nil
	}
	return red.URL
}

// GetUsage returns the value of the field Usage.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetUsage() *fhir.String {
	if red == nil {
		return nil
	}
	return red.Usage
}

// GetUseContext returns the value of the field UseContext.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetUseContext() []*fhir.UsageContext {
	if red == nil {
		return nil
	}
	return red.UseContext
}

// GetVariableType returns the value of the field VariableType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetVariableType() *fhir.Code {
	if red == nil {
		return nil
	}
	return red.VariableType
}

// GetVersion returns the value of the field Version.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (red *ResearchElementDefinition) GetVersion() *fhir.String {
	if red == nil {
		return nil
	}
	return red.Version
}

// What defines the members of the research element// A characteristic that defines the members of the research element. Multiple
// characteristics are applied with "and" semantics.// Characteristics can be defined flexibly to accommodate different use cases
// for membership criteria, ranging from simple codes, all the way to using an
// expression language to express the criteria.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-ResearchElementDefinition.json
type ResearchElementDefinitionCharacteristic struct {

	// Define members of the research element using Codes (such as condition,
	// medication, or observation), Expressions ( using an expression language such
	// as FHIRPath or CQL) or DataRequirements (such as Diabetes diagnosis onset in
	// the last year).
	Definition fhir.Element `fhirpath:"definition"`

	// When true, members with this characteristic are excluded from the element.
	Exclude *fhir.Boolean `fhirpath:"exclude"`

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

	// A narrative description of the time period the study covers.
	ParticipantEffectiveDescription *fhir.String `fhirpath:"participantEffectiveDescription"`

	// Indicates how elements are aggregated within the study effective period.
	ParticipantEffectiveGroupMeasure *fhir.Code `fhirpath:"participantEffectiveGroupMeasure"`

	// Indicates duration from the participant's study entry.
	ParticipantEffectiveTimeFromStart *fhir.Duration `fhirpath:"participantEffectiveTimeFromStart"`

	// Indicates what effective period the study covers.
	ParticipantEffective fhir.Element `fhirpath:"participantEffective"`

	// A narrative description of the time period the study covers.
	StudyEffectiveDescription *fhir.String `fhirpath:"studyEffectiveDescription"`

	// Indicates how elements are aggregated within the study effective period.
	StudyEffectiveGroupMeasure *fhir.Code `fhirpath:"studyEffectiveGroupMeasure"`

	// Indicates duration from the study initiation.
	StudyEffectiveTimeFromStart *fhir.Duration `fhirpath:"studyEffectiveTimeFromStart"`

	// Indicates what effective period the study covers.
	StudyEffective fhir.Element `fhirpath:"studyEffective"`

	// Specifies the UCUM unit for the outcome.
	UnitOfMeasure *fhir.CodeableConcept `fhirpath:"unitOfMeasure"`

	// Use UsageContext to define the members of the population, such as Age
	// Ranges, Genders, Settings.
	UsageContext []*fhir.UsageContext `fhirpath:"usageContext"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetDefinition returns the value of the field Definition.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetDefinition() fhir.Element {
	if redc == nil {
		return nil
	}
	return redc.Definition
}

// GetDefinitionCodeableConcept returns the value of the field Definition.
func (redc *ResearchElementDefinitionCharacteristic) GetDefinitionCodeableConcept() *fhir.CodeableConcept {
	if redc == nil {
		return nil
	}
	val, ok := redc.Definition.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetDefinitionCanonical returns the value of the field Definition.
func (redc *ResearchElementDefinitionCharacteristic) GetDefinitionCanonical() *fhir.Canonical {
	if redc == nil {
		return nil
	}
	val, ok := redc.Definition.(*fhir.Canonical)
	if !ok {
		return nil
	}
	return val
}

// GetDefinitionExpression returns the value of the field Definition.
func (redc *ResearchElementDefinitionCharacteristic) GetDefinitionExpression() *fhir.Expression {
	if redc == nil {
		return nil
	}
	val, ok := redc.Definition.(*fhir.Expression)
	if !ok {
		return nil
	}
	return val
}

// GetDefinitionDataRequirement returns the value of the field Definition.
func (redc *ResearchElementDefinitionCharacteristic) GetDefinitionDataRequirement() *fhir.DataRequirement {
	if redc == nil {
		return nil
	}
	val, ok := redc.Definition.(*fhir.DataRequirement)
	if !ok {
		return nil
	}
	return val
} // GetExclude returns the value of the field Exclude.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetExclude() *fhir.Boolean {
	if redc == nil {
		return nil
	}
	return redc.Exclude
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetExtension() []*fhir.Extension {
	if redc == nil {
		return nil
	}
	return redc.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetID() string {
	if redc == nil {
		return ""
	}
	return redc.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetModifierExtension() []*fhir.Extension {
	if redc == nil {
		return nil
	}
	return redc.ModifierExtension
}

// GetParticipantEffectiveDescription returns the value of the field ParticipantEffectiveDescription.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetParticipantEffectiveDescription() *fhir.String {
	if redc == nil {
		return nil
	}
	return redc.ParticipantEffectiveDescription
}

// GetParticipantEffectiveGroupMeasure returns the value of the field ParticipantEffectiveGroupMeasure.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetParticipantEffectiveGroupMeasure() *fhir.Code {
	if redc == nil {
		return nil
	}
	return redc.ParticipantEffectiveGroupMeasure
}

// GetParticipantEffectiveTimeFromStart returns the value of the field ParticipantEffectiveTimeFromStart.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetParticipantEffectiveTimeFromStart() *fhir.Duration {
	if redc == nil {
		return nil
	}
	return redc.ParticipantEffectiveTimeFromStart
}

// GetParticipantEffective returns the value of the field ParticipantEffective.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetParticipantEffective() fhir.Element {
	if redc == nil {
		return nil
	}
	return redc.ParticipantEffective
}

// GetParticipantEffectiveDateTime returns the value of the field ParticipantEffective.
func (redc *ResearchElementDefinitionCharacteristic) GetParticipantEffectiveDateTime() *fhir.DateTime {
	if redc == nil {
		return nil
	}
	val, ok := redc.ParticipantEffective.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
}

// GetParticipantEffectivePeriod returns the value of the field ParticipantEffective.
func (redc *ResearchElementDefinitionCharacteristic) GetParticipantEffectivePeriod() *fhir.Period {
	if redc == nil {
		return nil
	}
	val, ok := redc.ParticipantEffective.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
}

// GetParticipantEffectiveDuration returns the value of the field ParticipantEffective.
func (redc *ResearchElementDefinitionCharacteristic) GetParticipantEffectiveDuration() *fhir.Duration {
	if redc == nil {
		return nil
	}
	val, ok := redc.ParticipantEffective.(*fhir.Duration)
	if !ok {
		return nil
	}
	return val
}

// GetParticipantEffectiveTiming returns the value of the field ParticipantEffective.
func (redc *ResearchElementDefinitionCharacteristic) GetParticipantEffectiveTiming() *fhir.Timing {
	if redc == nil {
		return nil
	}
	val, ok := redc.ParticipantEffective.(*fhir.Timing)
	if !ok {
		return nil
	}
	return val
} // GetStudyEffectiveDescription returns the value of the field StudyEffectiveDescription.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetStudyEffectiveDescription() *fhir.String {
	if redc == nil {
		return nil
	}
	return redc.StudyEffectiveDescription
}

// GetStudyEffectiveGroupMeasure returns the value of the field StudyEffectiveGroupMeasure.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetStudyEffectiveGroupMeasure() *fhir.Code {
	if redc == nil {
		return nil
	}
	return redc.StudyEffectiveGroupMeasure
}

// GetStudyEffectiveTimeFromStart returns the value of the field StudyEffectiveTimeFromStart.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetStudyEffectiveTimeFromStart() *fhir.Duration {
	if redc == nil {
		return nil
	}
	return redc.StudyEffectiveTimeFromStart
}

// GetStudyEffective returns the value of the field StudyEffective.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetStudyEffective() fhir.Element {
	if redc == nil {
		return nil
	}
	return redc.StudyEffective
}

// GetStudyEffectiveDateTime returns the value of the field StudyEffective.
func (redc *ResearchElementDefinitionCharacteristic) GetStudyEffectiveDateTime() *fhir.DateTime {
	if redc == nil {
		return nil
	}
	val, ok := redc.StudyEffective.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
}

// GetStudyEffectivePeriod returns the value of the field StudyEffective.
func (redc *ResearchElementDefinitionCharacteristic) GetStudyEffectivePeriod() *fhir.Period {
	if redc == nil {
		return nil
	}
	val, ok := redc.StudyEffective.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
}

// GetStudyEffectiveDuration returns the value of the field StudyEffective.
func (redc *ResearchElementDefinitionCharacteristic) GetStudyEffectiveDuration() *fhir.Duration {
	if redc == nil {
		return nil
	}
	val, ok := redc.StudyEffective.(*fhir.Duration)
	if !ok {
		return nil
	}
	return val
}

// GetStudyEffectiveTiming returns the value of the field StudyEffective.
func (redc *ResearchElementDefinitionCharacteristic) GetStudyEffectiveTiming() *fhir.Timing {
	if redc == nil {
		return nil
	}
	val, ok := redc.StudyEffective.(*fhir.Timing)
	if !ok {
		return nil
	}
	return val
} // GetUnitOfMeasure returns the value of the field UnitOfMeasure.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetUnitOfMeasure() *fhir.CodeableConcept {
	if redc == nil {
		return nil
	}
	return redc.UnitOfMeasure
}

// GetUsageContext returns the value of the field UsageContext.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (redc *ResearchElementDefinitionCharacteristic) GetUsageContext() []*fhir.UsageContext {
	if redc == nil {
		return nil
	}
	return redc.UsageContext
}