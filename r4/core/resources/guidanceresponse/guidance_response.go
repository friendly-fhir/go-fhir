// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package guidanceresponse

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// A guidance response is the formal response to a guidance request, including
// any output parameters returned by the evaluation, as well as the description
// of any proposed actions to be taken.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/GuidanceResponse
//   - Source File: StructureDefinition-GuidanceResponse.json
type GuidanceResponse struct {

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// If the evaluation could not be completed due to lack of information, or
	// additional information would potentially result in a more accurate response,
	// this element will a description of the data required in order to proceed
	// with the evaluation. A subsequent request to the service should include this
	// data.
	DataRequirement []*fhir.DataRequirement `fhirpath:"dataRequirement"`

	// The encounter during which this response was created or to which the
	// creation of this record is tightly associated.
	Encounter *fhir.Reference `fhirpath:"encounter"`

	// Messages resulting from the evaluation of the artifact or artifacts. As part
	// of evaluating the request, the engine may produce informational or warning
	// messages. These messages will be provided by this element.
	EvaluationMessage []*fhir.Reference `fhirpath:"evaluationMessage"`

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

	// Allows a service to provide unique, business identifiers for the response.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

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

	// An identifier, CodeableConcept or canonical reference to the guidance that
	// was requested.
	Module fhir.Element `fhirpath:"module"`

	// Provides a mechanism to communicate additional information about the
	// response.
	Note []*fhir.Annotation `fhirpath:"note"`

	// Indicates when the guidance response was processed.
	OccurrenceDateTime *fhir.DateTime `fhirpath:"occurrenceDateTime"`

	// The output parameters of the evaluation, if any. Many modules will result in
	// the return of specific resources such as procedure or communication requests
	// that are returned as part of the operation result. However, modules may
	// define specific outputs that would be returned as the result of the
	// evaluation, and these would be returned in this element.
	OutputParameters *fhir.Reference `fhirpath:"outputParameters"`

	// Provides a reference to the device that performed the guidance.
	Performer *fhir.Reference `fhirpath:"performer"`

	// Describes the reason for the guidance response in coded or textual form.
	ReasonCode []*fhir.CodeableConcept `fhirpath:"reasonCode"`

	// Indicates the reason the request was initiated. This is typically provided
	// as a parameter to the evaluation and echoed by the service, although for
	// some use cases, such as subscription- or event-based scenarios, it may
	// provide an indication of the cause for the response.
	ReasonReference []*fhir.Reference `fhirpath:"reasonReference"`

	// The identifier of the request associated with this response. If an
	// identifier was given as part of the request, it will be reproduced here to
	// enable the requester to more easily identify the response in a multi-request
	// scenario.
	RequestIdentifier *fhir.Identifier `fhirpath:"requestIdentifier"`

	// The actions, if any, produced by the evaluation of the artifact.
	Result *fhir.Reference `fhirpath:"result"`

	// The status of the response. If the evaluation is completed successfully, the
	// status will indicate success. However, in order to complete the evaluation,
	// the engine may require more information. In this case, the status will be
	// data-required, and the response will contain a description of the additional
	// required information. If the evaluation completed successfully, but the
	// engine determines that a potentially more accurate response could be
	// provided if more data was available, the status will be data-requested, and
	// the response will contain a description of the additional requested
	// information.
	Status *fhir.Code `fhirpath:"status"`

	// The patient for which the request was processed.
	Subject *fhir.Reference `fhirpath:"subject"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	profileimpl.BaseGuidanceResponse
	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetContained() []fhir.Resource {
	if gr == nil {
		return nil
	}
	return gr.Contained
}

// GetDataRequirement returns the value of the field DataRequirement.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetDataRequirement() []*fhir.DataRequirement {
	if gr == nil {
		return nil
	}
	return gr.DataRequirement
}

// GetEncounter returns the value of the field Encounter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetEncounter() *fhir.Reference {
	if gr == nil {
		return nil
	}
	return gr.Encounter
}

// GetEvaluationMessage returns the value of the field EvaluationMessage.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetEvaluationMessage() []*fhir.Reference {
	if gr == nil {
		return nil
	}
	return gr.EvaluationMessage
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetExtension() []*fhir.Extension {
	if gr == nil {
		return nil
	}
	return gr.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetID() string {
	if gr == nil {
		return ""
	}
	return gr.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetIdentifier() []*fhir.Identifier {
	if gr == nil {
		return nil
	}
	return gr.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetImplicitRules() *fhir.URI {
	if gr == nil {
		return nil
	}
	return gr.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetLanguage() *fhir.Code {
	if gr == nil {
		return nil
	}
	return gr.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetMeta() *fhir.Meta {
	if gr == nil {
		return nil
	}
	return gr.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetModifierExtension() []*fhir.Extension {
	if gr == nil {
		return nil
	}
	return gr.ModifierExtension
}

// GetModule returns the value of the field Module.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetModule() fhir.Element {
	if gr == nil {
		return nil
	}
	return gr.Module
}

// GetModuleURI returns the value of the field Module.
func (gr *GuidanceResponse) GetModuleURI() *fhir.URI {
	if gr == nil {
		return nil
	}
	val, ok := gr.Module.(*fhir.URI)
	if !ok {
		return nil
	}
	return val
}

// GetModuleCanonical returns the value of the field Module.
func (gr *GuidanceResponse) GetModuleCanonical() *fhir.Canonical {
	if gr == nil {
		return nil
	}
	val, ok := gr.Module.(*fhir.Canonical)
	if !ok {
		return nil
	}
	return val
}

// GetModuleCodeableConcept returns the value of the field Module.
func (gr *GuidanceResponse) GetModuleCodeableConcept() *fhir.CodeableConcept {
	if gr == nil {
		return nil
	}
	val, ok := gr.Module.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
} // GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetNote() []*fhir.Annotation {
	if gr == nil {
		return nil
	}
	return gr.Note
}

// GetOccurrenceDateTime returns the value of the field OccurrenceDateTime.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetOccurrenceDateTime() *fhir.DateTime {
	if gr == nil {
		return nil
	}
	return gr.OccurrenceDateTime
}

// GetOutputParameters returns the value of the field OutputParameters.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetOutputParameters() *fhir.Reference {
	if gr == nil {
		return nil
	}
	return gr.OutputParameters
}

// GetPerformer returns the value of the field Performer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetPerformer() *fhir.Reference {
	if gr == nil {
		return nil
	}
	return gr.Performer
}

// GetReasonCode returns the value of the field ReasonCode.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetReasonCode() []*fhir.CodeableConcept {
	if gr == nil {
		return nil
	}
	return gr.ReasonCode
}

// GetReasonReference returns the value of the field ReasonReference.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetReasonReference() []*fhir.Reference {
	if gr == nil {
		return nil
	}
	return gr.ReasonReference
}

// GetRequestIdentifier returns the value of the field RequestIdentifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetRequestIdentifier() *fhir.Identifier {
	if gr == nil {
		return nil
	}
	return gr.RequestIdentifier
}

// GetResult returns the value of the field Result.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetResult() *fhir.Reference {
	if gr == nil {
		return nil
	}
	return gr.Result
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetStatus() *fhir.Code {
	if gr == nil {
		return nil
	}
	return gr.Status
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetSubject() *fhir.Reference {
	if gr == nil {
		return nil
	}
	return gr.Subject
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gr *GuidanceResponse) GetText() *fhir.Narrative {
	if gr == nil {
		return nil
	}
	return gr.Text
}

func (gr *GuidanceResponse) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (gr *GuidanceResponse) UnmarshalJSON(data []byte) error {
	var raw struct {
		Contained         []fhir.Resource         `json:"contained"`
		DataRequirement   []*fhir.DataRequirement `json:"dataRequirement"`
		Encounter         *fhir.Reference         `json:"encounter"`
		EvaluationMessage []*fhir.Reference       `json:"evaluationMessage"`
		Extension         []*fhir.Extension       `json:"extension"`

		ID                    string                  `json:"id"`
		Identifier            []*fhir.Identifier      `json:"identifier"`
		ImplicitRules         *fhir.URI               `json:"implicitRules"`
		Language              *fhir.Code              `json:"language"`
		Meta                  *fhir.Meta              `json:"meta"`
		ModifierExtension     []*fhir.Extension       `json:"modifierExtension"`
		ModuleURI             *fhir.URI               `json:"moduleURI"`
		ModuleCanonical       *fhir.Canonical         `json:"moduleCanonical"`
		ModuleCodeableConcept *fhir.CodeableConcept   `json:"moduleCodeableConcept"`
		Note                  []*fhir.Annotation      `json:"note"`
		OccurrenceDateTime    *fhir.DateTime          `json:"occurrenceDateTime"`
		OutputParameters      *fhir.Reference         `json:"outputParameters"`
		Performer             *fhir.Reference         `json:"performer"`
		ReasonCode            []*fhir.CodeableConcept `json:"reasonCode"`
		ReasonReference       []*fhir.Reference       `json:"reasonReference"`
		RequestIdentifier     *fhir.Identifier        `json:"requestIdentifier"`
		Result                *fhir.Reference         `json:"result"`
		Status                *fhir.Code              `json:"status"`
		Subject               *fhir.Reference         `json:"subject"`
		Text                  *fhir.Narrative         `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	gr.Contained = raw.Contained
	gr.DataRequirement = raw.DataRequirement
	gr.Encounter = raw.Encounter
	gr.EvaluationMessage = raw.EvaluationMessage
	gr.Extension = raw.Extension
	gr.ID = raw.ID
	gr.Identifier = raw.Identifier
	gr.ImplicitRules = raw.ImplicitRules
	gr.Language = raw.Language
	gr.Meta = raw.Meta
	gr.ModifierExtension = raw.ModifierExtension
	gr.Module, err = validate.SelectOneOf[fhir.Element]("GuidanceResponse.module",
		raw.ModuleURI,
		raw.ModuleCanonical,
		raw.ModuleCodeableConcept)
	if err != nil {
		return err
	}
	gr.Note = raw.Note
	gr.OccurrenceDateTime = raw.OccurrenceDateTime
	gr.OutputParameters = raw.OutputParameters
	gr.Performer = raw.Performer
	gr.ReasonCode = raw.ReasonCode
	gr.ReasonReference = raw.ReasonReference
	gr.RequestIdentifier = raw.RequestIdentifier
	gr.Result = raw.Result
	gr.Status = raw.Status
	gr.Subject = raw.Subject
	gr.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*GuidanceResponse)(nil)
var _ json.Unmarshaler = (*GuidanceResponse)(nil)
