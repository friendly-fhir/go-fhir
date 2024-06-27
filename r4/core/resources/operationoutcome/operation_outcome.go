// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package operationoutcome

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// A collection of error, warning, or information messages that result from a
// system action.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/OperationOutcome
//   - Source File: StructureDefinition-OperationOutcome.json
type OperationOutcome struct {

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

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

	// An error, warning, or information message that results from a system action.
	Issue []*OperationOutcomeIssue `fhirpath:"issue"`

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

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oo *OperationOutcome) GetContained() []fhir.Resource {
	if oo == nil {
		return nil
	}
	return oo.Contained
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oo *OperationOutcome) GetExtension() []*fhir.Extension {
	if oo == nil {
		return nil
	}
	return oo.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oo *OperationOutcome) GetID() string {
	if oo == nil {
		return ""
	}
	return oo.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oo *OperationOutcome) GetImplicitRules() *fhir.URI {
	if oo == nil {
		return nil
	}
	return oo.ImplicitRules
}

// GetIssue returns the value of the field Issue.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oo *OperationOutcome) GetIssue() []*OperationOutcomeIssue {
	if oo == nil {
		return nil
	}
	return oo.Issue
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oo *OperationOutcome) GetLanguage() *fhir.Code {
	if oo == nil {
		return nil
	}
	return oo.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oo *OperationOutcome) GetMeta() *fhir.Meta {
	if oo == nil {
		return nil
	}
	return oo.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oo *OperationOutcome) GetModifierExtension() []*fhir.Extension {
	if oo == nil {
		return nil
	}
	return oo.ModifierExtension
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oo *OperationOutcome) GetText() *fhir.Narrative {
	if oo == nil {
		return nil
	}
	return oo.Text
}

// A single issue associated with the action// An error, warning, or information message that results from a system action.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-OperationOutcome.json
type OperationOutcomeIssue struct {

	// Describes the type of the issue. The system that creates an OperationOutcome
	// SHALL choose the most applicable code from the IssueType value set, and may
	// additional provide its own code for the error in the details element.
	Code *fhir.Code `fhirpath:"code"`

	// Additional details about the error. This may be a text description of the
	// error or a system code that identifies the error.
	Details *fhir.CodeableConcept `fhirpath:"details"`

	// Additional diagnostic information about the issue.
	Diagnostics *fhir.String `fhirpath:"diagnostics"`

	// A [simple subset of FHIRPath](fhirpath.html#simple) limited to element
	// names, repetition indicators and the default child accessor that identifies
	// one of the elements in the resource that caused this issue to be raised.
	Expression []*fhir.String `fhirpath:"expression"`

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

	// This element is deprecated because it is XML specific. It is replaced by
	// issue.expression, which is format independent, and simpler to parse.
	// For resource issues, this will be a simple XPath limited to element names,
	// repetition indicators and the default child accessor that identifies one of
	// the elements in the resource that caused this issue to be raised. For HTTP
	// errors, will be "http." + the parameter name.
	Location []*fhir.String `fhirpath:"location"`

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

	// Indicates whether the issue indicates a variation from successful
	// processing.
	Severity *fhir.Code `fhirpath:"severity"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ooi *OperationOutcomeIssue) GetCode() *fhir.Code {
	if ooi == nil {
		return nil
	}
	return ooi.Code
}

// GetDetails returns the value of the field Details.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ooi *OperationOutcomeIssue) GetDetails() *fhir.CodeableConcept {
	if ooi == nil {
		return nil
	}
	return ooi.Details
}

// GetDiagnostics returns the value of the field Diagnostics.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ooi *OperationOutcomeIssue) GetDiagnostics() *fhir.String {
	if ooi == nil {
		return nil
	}
	return ooi.Diagnostics
}

// GetExpression returns the value of the field Expression.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ooi *OperationOutcomeIssue) GetExpression() []*fhir.String {
	if ooi == nil {
		return nil
	}
	return ooi.Expression
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ooi *OperationOutcomeIssue) GetExtension() []*fhir.Extension {
	if ooi == nil {
		return nil
	}
	return ooi.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ooi *OperationOutcomeIssue) GetID() string {
	if ooi == nil {
		return ""
	}
	return ooi.ID
}

// GetLocation returns the value of the field Location.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ooi *OperationOutcomeIssue) GetLocation() []*fhir.String {
	if ooi == nil {
		return nil
	}
	return ooi.Location
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ooi *OperationOutcomeIssue) GetModifierExtension() []*fhir.Extension {
	if ooi == nil {
		return nil
	}
	return ooi.ModifierExtension
}

// GetSeverity returns the value of the field Severity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ooi *OperationOutcomeIssue) GetSeverity() *fhir.Code {
	if ooi == nil {
		return nil
	}
	return ooi.Severity
}

func (oo *OperationOutcome) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (oo *OperationOutcome) UnmarshalJSON(data []byte) error {
	var raw struct {
		Contained []fhir.Resource   `json:"contained"`
		Extension []*fhir.Extension `json:"extension"`

		ID                string                   `json:"id"`
		ImplicitRules     *fhir.URI                `json:"implicitRules"`
		Issue             []*OperationOutcomeIssue `json:"issue"`
		Language          *fhir.Code               `json:"language"`
		Meta              *fhir.Meta               `json:"meta"`
		ModifierExtension []*fhir.Extension        `json:"modifierExtension"`
		Text              *fhir.Narrative          `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	oo.Contained = raw.Contained
	oo.Extension = raw.Extension
	oo.ID = raw.ID
	oo.ImplicitRules = raw.ImplicitRules
	oo.Issue = raw.Issue
	oo.Language = raw.Language
	oo.Meta = raw.Meta
	oo.ModifierExtension = raw.ModifierExtension
	oo.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*OperationOutcome)(nil)
var _ json.Unmarshaler = (*OperationOutcome)(nil)

func (ooi *OperationOutcomeIssue) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (ooi *OperationOutcomeIssue) UnmarshalJSON(data []byte) error {
	var raw struct {
		Code        *fhir.Code            `json:"code"`
		Details     *fhir.CodeableConcept `json:"details"`
		Diagnostics *fhir.String          `json:"diagnostics"`
		Expression  []*fhir.String        `json:"expression"`
		Extension   []*fhir.Extension     `json:"extension"`

		ID                string            `json:"id"`
		Location          []*fhir.String    `json:"location"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		Severity          *fhir.Code        `json:"severity"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	ooi.Code = raw.Code
	ooi.Details = raw.Details
	ooi.Diagnostics = raw.Diagnostics
	ooi.Expression = raw.Expression
	ooi.Extension = raw.Extension
	ooi.ID = raw.ID
	ooi.Location = raw.Location
	ooi.ModifierExtension = raw.ModifierExtension
	ooi.Severity = raw.Severity
	return nil
}

var _ json.Marshaler = (*OperationOutcomeIssue)(nil)
var _ json.Unmarshaler = (*OperationOutcomeIssue)(nil)
