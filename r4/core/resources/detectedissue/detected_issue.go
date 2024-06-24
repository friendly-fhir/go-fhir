// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package detectedissue

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Indicates an actual or potential clinical issue with or between one or more
// active or proposed clinical actions for a patient; e.g. Drug-drug
// interaction, Ineffective treatment frequency, Procedure-condition conflict,
// etc.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/DetectedIssue
//   - Source File: StructureDefinition-DetectedIssue.json
type DetectedIssue struct {

	// Individual or device responsible for the issue being raised. For example, a
	// decision support application or a pharmacist conducting a medication review.
	Author *fhir.Reference `fhirpath:"author"`

	// Identifies the general type of issue identified.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// A textual explanation of the detected issue.
	Detail *fhir.String `fhirpath:"detail"`

	// Supporting evidence or manifestations that provide the basis for identifying
	// the detected issue such as a GuidanceResponse or MeasureReport.
	Evidence []*DetectedIssueEvidence `fhirpath:"evidence"`

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

	// The date or period when the detected issue was initially identified.
	Identified fhir.Element `fhirpath:"identified"`

	// Business identifier associated with the detected issue record.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// Indicates the resource representing the current activity or proposed
	// activity that is potentially problematic.
	Implicated []*fhir.Reference `fhirpath:"implicated"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// The metadata about the resource. This is content that is maintained by the
	// infrastructure. Changes to the content might not always be associated with
	// version changes to the resource.
	Meta *fhir.Meta `fhirpath:"meta"`

	// Indicates an action that has been taken or is committed to reduce or
	// eliminate the likelihood of the risk identified by the detected issue from
	// manifesting. Can also reflect an observation of known mitigating factors
	// that may reduce/eliminate the need for any action.
	Mitigation []*DetectedIssueMitigation `fhirpath:"mitigation"`

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

	// Indicates the patient whose record the detected issue is associated with.
	Patient *fhir.Reference `fhirpath:"patient"`

	// The literature, knowledge-base or similar reference that describes the
	// propensity for the detected issue identified.
	Reference *fhir.URI `fhirpath:"reference"`

	// Indicates the degree of importance associated with the identified issue
	// based on the potential impact on the patient.
	Severity *fhir.Code `fhirpath:"severity"`

	// Indicates the status of the detected issue.
	Status *fhir.Code `fhirpath:"status"`

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

// GetAuthor returns the value of the field Author.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetAuthor() *fhir.Reference {
	if di == nil {
		return nil
	}
	return di.Author
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetCode() *fhir.CodeableConcept {
	if di == nil {
		return nil
	}
	return di.Code
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetContained() []fhir.Resource {
	if di == nil {
		return nil
	}
	return di.Contained
}

// GetDetail returns the value of the field Detail.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetDetail() *fhir.String {
	if di == nil {
		return nil
	}
	return di.Detail
}

// GetEvidence returns the value of the field Evidence.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetEvidence() []*DetectedIssueEvidence {
	if di == nil {
		return nil
	}
	return di.Evidence
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetExtension() []*fhir.Extension {
	if di == nil {
		return nil
	}
	return di.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetID() string {
	if di == nil {
		return ""
	}
	return di.ID
}

// GetIdentified returns the value of the field Identified.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetIdentified() fhir.Element {
	if di == nil {
		return nil
	}
	return di.Identified
}

// GetIdentifiedDateTime returns the value of the field Identified.
func (di *DetectedIssue) GetIdentifiedDateTime() *fhir.DateTime {
	if di == nil {
		return nil
	}
	val, ok := di.Identified.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
}

// GetIdentifiedPeriod returns the value of the field Identified.
func (di *DetectedIssue) GetIdentifiedPeriod() *fhir.Period {
	if di == nil {
		return nil
	}
	val, ok := di.Identified.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
} // GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetIdentifier() []*fhir.Identifier {
	if di == nil {
		return nil
	}
	return di.Identifier
}

// GetImplicated returns the value of the field Implicated.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetImplicated() []*fhir.Reference {
	if di == nil {
		return nil
	}
	return di.Implicated
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetImplicitRules() *fhir.URI {
	if di == nil {
		return nil
	}
	return di.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetLanguage() *fhir.Code {
	if di == nil {
		return nil
	}
	return di.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetMeta() *fhir.Meta {
	if di == nil {
		return nil
	}
	return di.Meta
}

// GetMitigation returns the value of the field Mitigation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetMitigation() []*DetectedIssueMitigation {
	if di == nil {
		return nil
	}
	return di.Mitigation
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetModifierExtension() []*fhir.Extension {
	if di == nil {
		return nil
	}
	return di.ModifierExtension
}

// GetPatient returns the value of the field Patient.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetPatient() *fhir.Reference {
	if di == nil {
		return nil
	}
	return di.Patient
}

// GetReference returns the value of the field Reference.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetReference() *fhir.URI {
	if di == nil {
		return nil
	}
	return di.Reference
}

// GetSeverity returns the value of the field Severity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetSeverity() *fhir.Code {
	if di == nil {
		return nil
	}
	return di.Severity
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetStatus() *fhir.Code {
	if di == nil {
		return nil
	}
	return di.Status
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (di *DetectedIssue) GetText() *fhir.Narrative {
	if di == nil {
		return nil
	}
	return di.Text
}

// Supporting evidence// Supporting evidence or manifestations that provide the basis for identifying
// the detected issue such as a GuidanceResponse or MeasureReport.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-DetectedIssue.json
type DetectedIssueEvidence struct {

	// A manifestation that led to the recording of this detected issue.
	Code []*fhir.CodeableConcept `fhirpath:"code"`

	// Links to resources that constitute evidence for the detected issue such as a
	// GuidanceResponse or MeasureReport.
	Detail []*fhir.Reference `fhirpath:"detail"`

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

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (die *DetectedIssueEvidence) GetCode() []*fhir.CodeableConcept {
	if die == nil {
		return nil
	}
	return die.Code
}

// GetDetail returns the value of the field Detail.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (die *DetectedIssueEvidence) GetDetail() []*fhir.Reference {
	if die == nil {
		return nil
	}
	return die.Detail
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (die *DetectedIssueEvidence) GetExtension() []*fhir.Extension {
	if die == nil {
		return nil
	}
	return die.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (die *DetectedIssueEvidence) GetID() string {
	if die == nil {
		return ""
	}
	return die.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (die *DetectedIssueEvidence) GetModifierExtension() []*fhir.Extension {
	if die == nil {
		return nil
	}
	return die.ModifierExtension
}

// Step taken to address// Indicates an action that has been taken or is committed to reduce or
// eliminate the likelihood of the risk identified by the detected issue from
// manifesting. Can also reflect an observation of known mitigating factors
// that may reduce/eliminate the need for any action.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-DetectedIssue.json
type DetectedIssueMitigation struct {

	// Describes the action that was taken or the observation that was made that
	// reduces/eliminates the risk associated with the identified issue.
	Action *fhir.CodeableConcept `fhirpath:"action"`

	// Identifies the practitioner who determined the mitigation and takes
	// responsibility for the mitigation step occurring.
	Author *fhir.Reference `fhirpath:"author"`

	// Indicates when the mitigating action was documented.
	Date *fhir.DateTime `fhirpath:"date"`

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

// GetAction returns the value of the field Action.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dim *DetectedIssueMitigation) GetAction() *fhir.CodeableConcept {
	if dim == nil {
		return nil
	}
	return dim.Action
}

// GetAuthor returns the value of the field Author.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dim *DetectedIssueMitigation) GetAuthor() *fhir.Reference {
	if dim == nil {
		return nil
	}
	return dim.Author
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dim *DetectedIssueMitigation) GetDate() *fhir.DateTime {
	if dim == nil {
		return nil
	}
	return dim.Date
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dim *DetectedIssueMitigation) GetExtension() []*fhir.Extension {
	if dim == nil {
		return nil
	}
	return dim.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dim *DetectedIssueMitigation) GetID() string {
	if dim == nil {
		return ""
	}
	return dim.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dim *DetectedIssueMitigation) GetModifierExtension() []*fhir.Extension {
	if dim == nil {
		return nil
	}
	return dim.ModifierExtension
}
