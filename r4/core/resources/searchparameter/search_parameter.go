// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package searchparameter

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// A search parameter that defines a named search item that can be used to
// search/filter on a resource.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/SearchParameter
//   - Source File: StructureDefinition-SearchParameter.json
type SearchParameter struct {

	// The base resource type(s) that this search parameter can be used against.
	Base []*fhir.Code `fhirpath:"base"`

	// Contains the names of any search parameters which may be chained to the
	// containing search parameter. Chained parameters may be added to search
	// parameters of type reference and specify that resources will only be
	// returned if they contain a reference to a resource which matches the chained
	// parameter value. Values for this field should be drawn from
	// SearchParameter.code for a parameter on the target resource type.
	Chain []*fhir.String `fhirpath:"chain"`

	// The code used in the URL or the parameter name in a parameters resource for
	// this search parameter.
	Code *fhir.Code `fhirpath:"code"`

	// Comparators supported for the search parameter.
	Comparator []*fhir.Code `fhirpath:"comparator"`

	// Used to define the parts of a composite search parameter.
	Component []*SearchParameterComponent `fhirpath:"component"`

	// Contact details to assist a user in finding and communicating with the
	// publisher.
	Contact []*fhir.ContactDetail `fhirpath:"contact"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The date (and optionally time) when the search parameter was published. The
	// date must change when the business version changes and it must change if the
	// status code changes. In addition, it should change when the substantive
	// content of the search parameter changes.
	Date *fhir.DateTime `fhirpath:"date"`

	// Where this search parameter is originally defined. If a derivedFrom is
	// provided, then the details in the search parameter must be consistent with
	// the definition from which it is defined. i.e. the parameter should have the
	// same meaning, and (usually) the functionality should be a proper subset of
	// the underlying search parameter.
	DerivedFrom *fhir.Canonical `fhirpath:"derivedFrom"`

	// And how it used.
	Description *fhir.Markdown `fhirpath:"description"`

	// A Boolean value to indicate that this search parameter is authored for
	// testing purposes (or education/evaluation/marketing) and is not intended to
	// be used for genuine usage.
	Experimental *fhir.Boolean `fhirpath:"experimental"`

	// A FHIRPath expression that returns a set of elements for the search
	// parameter.
	Expression *fhir.String `fhirpath:"expression"`

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

	// A legal or geographic region in which the search parameter is intended to be
	// used.
	Jurisdiction []*fhir.CodeableConcept `fhirpath:"jurisdiction"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// The metadata about the resource. This is content that is maintained by the
	// infrastructure. Changes to the content might not always be associated with
	// version changes to the resource.
	Meta *fhir.Meta `fhirpath:"meta"`

	// A modifier supported for the search parameter.
	Modifier []*fhir.Code `fhirpath:"modifier"`

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

	// Whether multiple parameters are allowed - e.g. more than one parameter with
	// the same name. The search matches if all the parameters match.
	MultipleAnd *fhir.Boolean `fhirpath:"multipleAnd"`

	// Whether multiple values are allowed for each time the parameter exists.
	// Values are separated by commas, and the parameter matches if any of the
	// values match.
	MultipleOr *fhir.Boolean `fhirpath:"multipleOr"`

	// A natural language name identifying the search parameter. This name should
	// be usable as an identifier for the module by machine processing applications
	// such as code generation.
	Name *fhir.String `fhirpath:"name"`

	// The name of the organization or individual that published the search
	// parameter.
	Publisher *fhir.String `fhirpath:"publisher"`

	// Explanation of why this search parameter is needed and why it has been
	// designed as it has.
	Purpose *fhir.Markdown `fhirpath:"purpose"`

	// The status of this search parameter. Enables tracking the life-cycle of the
	// content.
	Status *fhir.Code `fhirpath:"status"`

	// Types of resource (if a resource is referenced).
	Target []*fhir.Code `fhirpath:"target"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// The type of value that a search parameter may contain, and how the content
	// is interpreted.
	Type *fhir.Code `fhirpath:"type"`

	// An absolute URI that is used to identify this search parameter when it is
	// referenced in a specification, model, design or an instance; also called its
	// canonical identifier. This SHOULD be globally unique and SHOULD be a literal
	// address at which at which an authoritative instance of this search parameter
	// is (or will be) published. This URL can be the target of a canonical
	// reference. It SHALL remain the same when the search parameter is stored on
	// different servers.
	URL *fhir.URI `fhirpath:"url"`

	// The content was developed with a focus and intent of supporting the contexts
	// that are listed. These contexts may be general categories (gender, age, ...)
	// or may be references to specific programs (insurance plans, studies, ...)
	// and may be used to assist with indexing and searching for appropriate search
	// parameter instances.
	UseContext []*fhir.UsageContext `fhirpath:"useContext"`

	// The identifier that is used to identify this version of the search parameter
	// when it is referenced in a specification, model, design or instance. This is
	// an arbitrary value managed by the search parameter author and is not
	// expected to be globally unique. For example, it might be a timestamp (e.g.
	// yyyymmdd) if a managed version is not available. There is also no
	// expectation that versions can be placed in a lexicographical sequence.
	Version *fhir.String `fhirpath:"version"`

	// An XPath expression that returns a set of elements for the search parameter.
	Xpath *fhir.String `fhirpath:"xpath"`

	// How the search parameter relates to the set of elements returned by
	// evaluating the xpath query.
	XpathUsage *fhir.Code `fhirpath:"xpathUsage"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetBase returns the value of the field Base.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetBase() []*fhir.Code {
	if sp == nil {
		return nil
	}
	return sp.Base
}

// GetChain returns the value of the field Chain.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetChain() []*fhir.String {
	if sp == nil {
		return nil
	}
	return sp.Chain
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetCode() *fhir.Code {
	if sp == nil {
		return nil
	}
	return sp.Code
}

// GetComparator returns the value of the field Comparator.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetComparator() []*fhir.Code {
	if sp == nil {
		return nil
	}
	return sp.Comparator
}

// GetComponent returns the value of the field Component.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetComponent() []*SearchParameterComponent {
	if sp == nil {
		return nil
	}
	return sp.Component
}

// GetContact returns the value of the field Contact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetContact() []*fhir.ContactDetail {
	if sp == nil {
		return nil
	}
	return sp.Contact
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetContained() []fhir.Resource {
	if sp == nil {
		return nil
	}
	return sp.Contained
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetDate() *fhir.DateTime {
	if sp == nil {
		return nil
	}
	return sp.Date
}

// GetDerivedFrom returns the value of the field DerivedFrom.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetDerivedFrom() *fhir.Canonical {
	if sp == nil {
		return nil
	}
	return sp.DerivedFrom
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetDescription() *fhir.Markdown {
	if sp == nil {
		return nil
	}
	return sp.Description
}

// GetExperimental returns the value of the field Experimental.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetExperimental() *fhir.Boolean {
	if sp == nil {
		return nil
	}
	return sp.Experimental
}

// GetExpression returns the value of the field Expression.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetExpression() *fhir.String {
	if sp == nil {
		return nil
	}
	return sp.Expression
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetExtension() []*fhir.Extension {
	if sp == nil {
		return nil
	}
	return sp.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetID() string {
	if sp == nil {
		return ""
	}
	return sp.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetImplicitRules() *fhir.URI {
	if sp == nil {
		return nil
	}
	return sp.ImplicitRules
}

// GetJurisdiction returns the value of the field Jurisdiction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetJurisdiction() []*fhir.CodeableConcept {
	if sp == nil {
		return nil
	}
	return sp.Jurisdiction
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetLanguage() *fhir.Code {
	if sp == nil {
		return nil
	}
	return sp.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetMeta() *fhir.Meta {
	if sp == nil {
		return nil
	}
	return sp.Meta
}

// GetModifier returns the value of the field Modifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetModifier() []*fhir.Code {
	if sp == nil {
		return nil
	}
	return sp.Modifier
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetModifierExtension() []*fhir.Extension {
	if sp == nil {
		return nil
	}
	return sp.ModifierExtension
}

// GetMultipleAnd returns the value of the field MultipleAnd.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetMultipleAnd() *fhir.Boolean {
	if sp == nil {
		return nil
	}
	return sp.MultipleAnd
}

// GetMultipleOr returns the value of the field MultipleOr.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetMultipleOr() *fhir.Boolean {
	if sp == nil {
		return nil
	}
	return sp.MultipleOr
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetName() *fhir.String {
	if sp == nil {
		return nil
	}
	return sp.Name
}

// GetPublisher returns the value of the field Publisher.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetPublisher() *fhir.String {
	if sp == nil {
		return nil
	}
	return sp.Publisher
}

// GetPurpose returns the value of the field Purpose.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetPurpose() *fhir.Markdown {
	if sp == nil {
		return nil
	}
	return sp.Purpose
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetStatus() *fhir.Code {
	if sp == nil {
		return nil
	}
	return sp.Status
}

// GetTarget returns the value of the field Target.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetTarget() []*fhir.Code {
	if sp == nil {
		return nil
	}
	return sp.Target
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetText() *fhir.Narrative {
	if sp == nil {
		return nil
	}
	return sp.Text
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetType() *fhir.Code {
	if sp == nil {
		return nil
	}
	return sp.Type
}

// GetURL returns the value of the field URL.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetURL() *fhir.URI {
	if sp == nil {
		return nil
	}
	return sp.URL
}

// GetUseContext returns the value of the field UseContext.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetUseContext() []*fhir.UsageContext {
	if sp == nil {
		return nil
	}
	return sp.UseContext
}

// GetVersion returns the value of the field Version.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetVersion() *fhir.String {
	if sp == nil {
		return nil
	}
	return sp.Version
}

// GetXpath returns the value of the field Xpath.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetXpath() *fhir.String {
	if sp == nil {
		return nil
	}
	return sp.Xpath
}

// GetXpathUsage returns the value of the field XpathUsage.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SearchParameter) GetXpathUsage() *fhir.Code {
	if sp == nil {
		return nil
	}
	return sp.XpathUsage
}

// For Composite resources to define the parts// Used to define the parts of a composite search parameter.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-SearchParameter.json
type SearchParameterComponent struct {

	// The definition of the search parameter that describes this part.
	Definition *fhir.Canonical `fhirpath:"definition"`

	// A sub-expression that defines how to extract values for this component from
	// the output of the main SearchParameter.expression.
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

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetDefinition returns the value of the field Definition.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (spc *SearchParameterComponent) GetDefinition() *fhir.Canonical {
	if spc == nil {
		return nil
	}
	return spc.Definition
}

// GetExpression returns the value of the field Expression.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (spc *SearchParameterComponent) GetExpression() *fhir.String {
	if spc == nil {
		return nil
	}
	return spc.Expression
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (spc *SearchParameterComponent) GetExtension() []*fhir.Extension {
	if spc == nil {
		return nil
	}
	return spc.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (spc *SearchParameterComponent) GetID() string {
	if spc == nil {
		return ""
	}
	return spc.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (spc *SearchParameterComponent) GetModifierExtension() []*fhir.Extension {
	if spc == nil {
		return nil
	}
	return spc.ModifierExtension
}
