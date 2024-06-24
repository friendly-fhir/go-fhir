// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package graphdefinition

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// A formal computable definition of a graph of resources - that is, a coherent
// set of resources that form a graph by following references. The Graph
// Definition resource defines a set and makes rules about the set.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/GraphDefinition
//   - Source File: StructureDefinition-GraphDefinition.json
type GraphDefinition struct {

	// Contact details to assist a user in finding and communicating with the
	// publisher.
	Contact []*fhir.ContactDetail `fhirpath:"contact"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The date (and optionally time) when the graph definition was published. The
	// date must change when the business version changes and it must change if the
	// status code changes. In addition, it should change when the substantive
	// content of the graph definition changes.
	Date *fhir.DateTime `fhirpath:"date"`

	// A free text natural language description of the graph definition from a
	// consumer's perspective.
	Description *fhir.Markdown `fhirpath:"description"`

	// A Boolean value to indicate that this graph definition is authored for
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

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// A legal or geographic region in which the graph definition is intended to be
	// used.
	Jurisdiction []*fhir.CodeableConcept `fhirpath:"jurisdiction"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// Links this graph makes rules about.
	Link []*GraphDefinitionLink `fhirpath:"link"`

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

	// A natural language name identifying the graph definition. This name should
	// be usable as an identifier for the module by machine processing applications
	// such as code generation.
	Name *fhir.String `fhirpath:"name"`

	// The profile that describes the use of the base resource.
	Profile *fhir.Canonical `fhirpath:"profile"`

	// The name of the organization or individual that published the graph
	// definition.
	Publisher *fhir.String `fhirpath:"publisher"`

	// Explanation of why this graph definition is needed and why it has been
	// designed as it has.
	Purpose *fhir.Markdown `fhirpath:"purpose"`

	// The type of FHIR resource at which instances of this graph start.
	Start *fhir.Code `fhirpath:"start"`

	// The status of this graph definition. Enables tracking the life-cycle of the
	// content.
	Status *fhir.Code `fhirpath:"status"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// An absolute URI that is used to identify this graph definition when it is
	// referenced in a specification, model, design or an instance; also called its
	// canonical identifier. This SHOULD be globally unique and SHOULD be a literal
	// address at which at which an authoritative instance of this graph definition
	// is (or will be) published. This URL can be the target of a canonical
	// reference. It SHALL remain the same when the graph definition is stored on
	// different servers.
	URL *fhir.URI `fhirpath:"url"`

	// The content was developed with a focus and intent of supporting the contexts
	// that are listed. These contexts may be general categories (gender, age, ...)
	// or may be references to specific programs (insurance plans, studies, ...)
	// and may be used to assist with indexing and searching for appropriate graph
	// definition instances.
	UseContext []*fhir.UsageContext `fhirpath:"useContext"`

	// The identifier that is used to identify this version of the graph definition
	// when it is referenced in a specification, model, design or instance. This is
	// an arbitrary value managed by the graph definition author and is not
	// expected to be globally unique. For example, it might be a timestamp (e.g.
	// yyyymmdd) if a managed version is not available. There is also no
	// expectation that versions can be placed in a lexicographical sequence.
	Version *fhir.String `fhirpath:"version"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetContact returns the value of the field Contact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetContact() []*fhir.ContactDetail {
	if gd == nil {
		return nil
	}
	return gd.Contact
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetContained() []fhir.Resource {
	if gd == nil {
		return nil
	}
	return gd.Contained
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetDate() *fhir.DateTime {
	if gd == nil {
		return nil
	}
	return gd.Date
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetDescription() *fhir.Markdown {
	if gd == nil {
		return nil
	}
	return gd.Description
}

// GetExperimental returns the value of the field Experimental.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetExperimental() *fhir.Boolean {
	if gd == nil {
		return nil
	}
	return gd.Experimental
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetExtension() []*fhir.Extension {
	if gd == nil {
		return nil
	}
	return gd.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetID() string {
	if gd == nil {
		return ""
	}
	return gd.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetImplicitRules() *fhir.URI {
	if gd == nil {
		return nil
	}
	return gd.ImplicitRules
}

// GetJurisdiction returns the value of the field Jurisdiction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetJurisdiction() []*fhir.CodeableConcept {
	if gd == nil {
		return nil
	}
	return gd.Jurisdiction
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetLanguage() *fhir.Code {
	if gd == nil {
		return nil
	}
	return gd.Language
}

// GetLink returns the value of the field Link.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetLink() []*GraphDefinitionLink {
	if gd == nil {
		return nil
	}
	return gd.Link
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetMeta() *fhir.Meta {
	if gd == nil {
		return nil
	}
	return gd.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetModifierExtension() []*fhir.Extension {
	if gd == nil {
		return nil
	}
	return gd.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetName() *fhir.String {
	if gd == nil {
		return nil
	}
	return gd.Name
}

// GetProfile returns the value of the field Profile.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetProfile() *fhir.Canonical {
	if gd == nil {
		return nil
	}
	return gd.Profile
}

// GetPublisher returns the value of the field Publisher.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetPublisher() *fhir.String {
	if gd == nil {
		return nil
	}
	return gd.Publisher
}

// GetPurpose returns the value of the field Purpose.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetPurpose() *fhir.Markdown {
	if gd == nil {
		return nil
	}
	return gd.Purpose
}

// GetStart returns the value of the field Start.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetStart() *fhir.Code {
	if gd == nil {
		return nil
	}
	return gd.Start
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetStatus() *fhir.Code {
	if gd == nil {
		return nil
	}
	return gd.Status
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetText() *fhir.Narrative {
	if gd == nil {
		return nil
	}
	return gd.Text
}

// GetURL returns the value of the field URL.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetURL() *fhir.URI {
	if gd == nil {
		return nil
	}
	return gd.URL
}

// GetUseContext returns the value of the field UseContext.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetUseContext() []*fhir.UsageContext {
	if gd == nil {
		return nil
	}
	return gd.UseContext
}

// GetVersion returns the value of the field Version.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gd *GraphDefinition) GetVersion() *fhir.String {
	if gd == nil {
		return nil
	}
	return gd.Version
}

// Links this graph makes rules about// Links this graph makes rules about.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-GraphDefinition.json
type GraphDefinitionLink struct {

	// Information about why this link is of interest in this graph definition.
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

	// Maximum occurrences for this link.
	Max *fhir.String `fhirpath:"max"`

	// Minimum occurrences for this link.
	Min *fhir.Integer `fhirpath:"min"`

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

	// A FHIR expression that identifies one of FHIR References to other resources.
	Path *fhir.String `fhirpath:"path"`

	// Which slice (if profiled).
	SliceName *fhir.String `fhirpath:"sliceName"`

	// Potential target for the link.
	Target []*GraphDefinitionLinkTarget `fhirpath:"target"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdl *GraphDefinitionLink) GetDescription() *fhir.String {
	if gdl == nil {
		return nil
	}
	return gdl.Description
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdl *GraphDefinitionLink) GetExtension() []*fhir.Extension {
	if gdl == nil {
		return nil
	}
	return gdl.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdl *GraphDefinitionLink) GetID() string {
	if gdl == nil {
		return ""
	}
	return gdl.ID
}

// GetMax returns the value of the field Max.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdl *GraphDefinitionLink) GetMax() *fhir.String {
	if gdl == nil {
		return nil
	}
	return gdl.Max
}

// GetMin returns the value of the field Min.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdl *GraphDefinitionLink) GetMin() *fhir.Integer {
	if gdl == nil {
		return nil
	}
	return gdl.Min
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdl *GraphDefinitionLink) GetModifierExtension() []*fhir.Extension {
	if gdl == nil {
		return nil
	}
	return gdl.ModifierExtension
}

// GetPath returns the value of the field Path.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdl *GraphDefinitionLink) GetPath() *fhir.String {
	if gdl == nil {
		return nil
	}
	return gdl.Path
}

// GetSliceName returns the value of the field SliceName.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdl *GraphDefinitionLink) GetSliceName() *fhir.String {
	if gdl == nil {
		return nil
	}
	return gdl.SliceName
}

// GetTarget returns the value of the field Target.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdl *GraphDefinitionLink) GetTarget() []*GraphDefinitionLinkTarget {
	if gdl == nil {
		return nil
	}
	return gdl.Target
}

// Potential target for the link// Potential target for the link.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-GraphDefinition.json
type GraphDefinitionLinkTarget struct {

	// Compartment Consistency Rules.
	Compartment []*GraphDefinitionLinkTargetCompartment `fhirpath:"compartment"`

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

	// A set of parameters to look up.
	Params *fhir.String `fhirpath:"params"`

	// Profile for the target resource.
	Profile *fhir.Canonical `fhirpath:"profile"`

	// Type of resource this link refers to.
	Type *fhir.Code `fhirpath:"type"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCompartment returns the value of the field Compartment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdlt *GraphDefinitionLinkTarget) GetCompartment() []*GraphDefinitionLinkTargetCompartment {
	if gdlt == nil {
		return nil
	}
	return gdlt.Compartment
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdlt *GraphDefinitionLinkTarget) GetExtension() []*fhir.Extension {
	if gdlt == nil {
		return nil
	}
	return gdlt.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdlt *GraphDefinitionLinkTarget) GetID() string {
	if gdlt == nil {
		return ""
	}
	return gdlt.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdlt *GraphDefinitionLinkTarget) GetModifierExtension() []*fhir.Extension {
	if gdlt == nil {
		return nil
	}
	return gdlt.ModifierExtension
}

// GetParams returns the value of the field Params.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdlt *GraphDefinitionLinkTarget) GetParams() *fhir.String {
	if gdlt == nil {
		return nil
	}
	return gdlt.Params
}

// GetProfile returns the value of the field Profile.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdlt *GraphDefinitionLinkTarget) GetProfile() *fhir.Canonical {
	if gdlt == nil {
		return nil
	}
	return gdlt.Profile
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdlt *GraphDefinitionLinkTarget) GetType() *fhir.Code {
	if gdlt == nil {
		return nil
	}
	return gdlt.Type
}

// Compartment Consistency Rules// Compartment Consistency Rules.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-GraphDefinition.json
type GraphDefinitionLinkTargetCompartment struct {

	// Identifies the compartment.
	Code *fhir.Code `fhirpath:"code"`

	// Documentation for FHIRPath expression.
	Description *fhir.String `fhirpath:"description"`

	// Custom rule, as a FHIRPath expression.
	Expression *fhir.String `fhirpath:"expression"`

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

	// identical | matching | different | no-rule | custom.
	Rule *fhir.Code `fhirpath:"rule"`

	// Defines how the compartment rule is used - whether it it is used to test
	// whether resources are subject to the rule, or whether it is a rule that must
	// be followed.
	Use *fhir.Code `fhirpath:"use"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdltc *GraphDefinitionLinkTargetCompartment) GetCode() *fhir.Code {
	if gdltc == nil {
		return nil
	}
	return gdltc.Code
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdltc *GraphDefinitionLinkTargetCompartment) GetDescription() *fhir.String {
	if gdltc == nil {
		return nil
	}
	return gdltc.Description
}

// GetExpression returns the value of the field Expression.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdltc *GraphDefinitionLinkTargetCompartment) GetExpression() *fhir.String {
	if gdltc == nil {
		return nil
	}
	return gdltc.Expression
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdltc *GraphDefinitionLinkTargetCompartment) GetExtension() []*fhir.Extension {
	if gdltc == nil {
		return nil
	}
	return gdltc.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdltc *GraphDefinitionLinkTargetCompartment) GetID() string {
	if gdltc == nil {
		return ""
	}
	return gdltc.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdltc *GraphDefinitionLinkTargetCompartment) GetModifierExtension() []*fhir.Extension {
	if gdltc == nil {
		return nil
	}
	return gdltc.ModifierExtension
}

// GetRule returns the value of the field Rule.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdltc *GraphDefinitionLinkTargetCompartment) GetRule() *fhir.Code {
	if gdltc == nil {
		return nil
	}
	return gdltc.Rule
}

// GetUse returns the value of the field Use.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gdltc *GraphDefinitionLinkTargetCompartment) GetUse() *fhir.Code {
	if gdltc == nil {
		return nil
	}
	return gdltc.Use
}
