// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package immunizationevaluation

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Describes a comparison of an immunization event against published
// recommendations to determine if the administration is "valid" in relation to
// those recommendations.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/ImmunizationEvaluation
//   - Source File: StructureDefinition-ImmunizationEvaluation.json
type ImmunizationEvaluation struct {

	// Indicates the authority who published the protocol (e.g. ACIP).
	Authority *fhir.Reference `fhirpath:"authority"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The date the evaluation of the vaccine administration event was performed.
	Date *fhir.DateTime `fhirpath:"date"`

	// Additional information about the evaluation.
	Description *fhir.String `fhirpath:"description"`

	// Nominal position in a series.
	DoseNumber fhir.Element `fhirpath:"doseNumber"`

	// Indicates if the dose is valid or not valid with respect to the published
	// recommendations.
	DoseStatus *fhir.CodeableConcept `fhirpath:"doseStatus"`

	// Provides an explanation as to why the vaccine administration event is valid
	// or not relative to the published recommendations.
	DoseStatusReason []*fhir.CodeableConcept `fhirpath:"doseStatusReason"`

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

	// A unique identifier assigned to this immunization evaluation record.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// The vaccine administration event being evaluated.
	ImmunizationEvent *fhir.Reference `fhirpath:"immunizationEvent"`

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

	// The individual for whom the evaluation is being done.
	Patient *fhir.Reference `fhirpath:"patient"`

	// One possible path to achieve presumed immunity against a disease - within
	// the context of an authority.
	Series *fhir.String `fhirpath:"series"`

	// The recommended number of doses to achieve immunity.
	SeriesDoses fhir.Element `fhirpath:"seriesDoses"`

	// Indicates the current status of the evaluation of the vaccination
	// administration event.
	Status *fhir.Code `fhirpath:"status"`

	// The vaccine preventable disease the dose is being evaluated against.
	TargetDisease *fhir.CodeableConcept `fhirpath:"targetDisease"`

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

// GetAuthority returns the value of the field Authority.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetAuthority() *fhir.Reference {
	if ie == nil {
		return nil
	}
	return ie.Authority
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetContained() []fhir.Resource {
	if ie == nil {
		return nil
	}
	return ie.Contained
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetDate() *fhir.DateTime {
	if ie == nil {
		return nil
	}
	return ie.Date
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetDescription() *fhir.String {
	if ie == nil {
		return nil
	}
	return ie.Description
}

// GetDoseNumber returns the value of the field DoseNumber.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetDoseNumber() fhir.Element {
	if ie == nil {
		return nil
	}
	return ie.DoseNumber
}

// GetDoseNumberPositiveInt returns the value of the field DoseNumber.
func (ie *ImmunizationEvaluation) GetDoseNumberPositiveInt() *fhir.PositiveInt {
	if ie == nil {
		return nil
	}
	val, ok := ie.DoseNumber.(*fhir.PositiveInt)
	if !ok {
		return nil
	}
	return val
}

// GetDoseNumberString returns the value of the field DoseNumber.
func (ie *ImmunizationEvaluation) GetDoseNumberString() *fhir.String {
	if ie == nil {
		return nil
	}
	val, ok := ie.DoseNumber.(*fhir.String)
	if !ok {
		return nil
	}
	return val
} // GetDoseStatus returns the value of the field DoseStatus.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetDoseStatus() *fhir.CodeableConcept {
	if ie == nil {
		return nil
	}
	return ie.DoseStatus
}

// GetDoseStatusReason returns the value of the field DoseStatusReason.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetDoseStatusReason() []*fhir.CodeableConcept {
	if ie == nil {
		return nil
	}
	return ie.DoseStatusReason
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetExtension() []*fhir.Extension {
	if ie == nil {
		return nil
	}
	return ie.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetID() string {
	if ie == nil {
		return ""
	}
	return ie.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetIdentifier() []*fhir.Identifier {
	if ie == nil {
		return nil
	}
	return ie.Identifier
}

// GetImmunizationEvent returns the value of the field ImmunizationEvent.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetImmunizationEvent() *fhir.Reference {
	if ie == nil {
		return nil
	}
	return ie.ImmunizationEvent
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetImplicitRules() *fhir.URI {
	if ie == nil {
		return nil
	}
	return ie.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetLanguage() *fhir.Code {
	if ie == nil {
		return nil
	}
	return ie.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetMeta() *fhir.Meta {
	if ie == nil {
		return nil
	}
	return ie.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetModifierExtension() []*fhir.Extension {
	if ie == nil {
		return nil
	}
	return ie.ModifierExtension
}

// GetPatient returns the value of the field Patient.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetPatient() *fhir.Reference {
	if ie == nil {
		return nil
	}
	return ie.Patient
}

// GetSeries returns the value of the field Series.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetSeries() *fhir.String {
	if ie == nil {
		return nil
	}
	return ie.Series
}

// GetSeriesDoses returns the value of the field SeriesDoses.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetSeriesDoses() fhir.Element {
	if ie == nil {
		return nil
	}
	return ie.SeriesDoses
}

// GetSeriesDosesPositiveInt returns the value of the field SeriesDoses.
func (ie *ImmunizationEvaluation) GetSeriesDosesPositiveInt() *fhir.PositiveInt {
	if ie == nil {
		return nil
	}
	val, ok := ie.SeriesDoses.(*fhir.PositiveInt)
	if !ok {
		return nil
	}
	return val
}

// GetSeriesDosesString returns the value of the field SeriesDoses.
func (ie *ImmunizationEvaluation) GetSeriesDosesString() *fhir.String {
	if ie == nil {
		return nil
	}
	val, ok := ie.SeriesDoses.(*fhir.String)
	if !ok {
		return nil
	}
	return val
} // GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetStatus() *fhir.Code {
	if ie == nil {
		return nil
	}
	return ie.Status
}

// GetTargetDisease returns the value of the field TargetDisease.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetTargetDisease() *fhir.CodeableConcept {
	if ie == nil {
		return nil
	}
	return ie.TargetDisease
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ie *ImmunizationEvaluation) GetText() *fhir.Narrative {
	if ie == nil {
		return nil
	}
	return ie.Text
}

func (ie *ImmunizationEvaluation) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (ie *ImmunizationEvaluation) UnmarshalJSON(data []byte) error {
	var raw struct {
		Authority             *fhir.Reference         `json:"authority"`
		Contained             []fhir.Resource         `json:"contained"`
		Date                  *fhir.DateTime          `json:"date"`
		Description           *fhir.String            `json:"description"`
		DoseNumberPositiveInt *fhir.PositiveInt       `json:"doseNumberPositiveInt"`
		DoseNumberString      *fhir.String            `json:"doseNumberString"`
		DoseStatus            *fhir.CodeableConcept   `json:"doseStatus"`
		DoseStatusReason      []*fhir.CodeableConcept `json:"doseStatusReason"`
		Extension             []*fhir.Extension       `json:"extension"`

		ID                     string                `json:"id"`
		Identifier             []*fhir.Identifier    `json:"identifier"`
		ImmunizationEvent      *fhir.Reference       `json:"immunizationEvent"`
		ImplicitRules          *fhir.URI             `json:"implicitRules"`
		Language               *fhir.Code            `json:"language"`
		Meta                   *fhir.Meta            `json:"meta"`
		ModifierExtension      []*fhir.Extension     `json:"modifierExtension"`
		Patient                *fhir.Reference       `json:"patient"`
		Series                 *fhir.String          `json:"series"`
		SeriesDosesPositiveInt *fhir.PositiveInt     `json:"seriesDosesPositiveInt"`
		SeriesDosesString      *fhir.String          `json:"seriesDosesString"`
		Status                 *fhir.Code            `json:"status"`
		TargetDisease          *fhir.CodeableConcept `json:"targetDisease"`
		Text                   *fhir.Narrative       `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	ie.Authority = raw.Authority
	ie.Contained = raw.Contained
	ie.Date = raw.Date
	ie.Description = raw.Description
	ie.DoseNumber, err = validate.SelectOneOf[fhir.Element]("ImmunizationEvaluation.doseNumber",
		raw.DoseNumberPositiveInt,
		raw.DoseNumberString)
	if err != nil {
		return err
	}
	ie.DoseStatus = raw.DoseStatus
	ie.DoseStatusReason = raw.DoseStatusReason
	ie.Extension = raw.Extension
	ie.ID = raw.ID
	ie.Identifier = raw.Identifier
	ie.ImmunizationEvent = raw.ImmunizationEvent
	ie.ImplicitRules = raw.ImplicitRules
	ie.Language = raw.Language
	ie.Meta = raw.Meta
	ie.ModifierExtension = raw.ModifierExtension
	ie.Patient = raw.Patient
	ie.Series = raw.Series
	ie.SeriesDoses, err = validate.SelectOneOf[fhir.Element]("ImmunizationEvaluation.seriesDoses",
		raw.SeriesDosesPositiveInt,
		raw.SeriesDosesString)
	if err != nil {
		return err
	}
	ie.Status = raw.Status
	ie.TargetDisease = raw.TargetDisease
	ie.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*ImmunizationEvaluation)(nil)
var _ json.Unmarshaler = (*ImmunizationEvaluation)(nil)
