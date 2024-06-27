// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package enrollmentrequest

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// This resource provides the insurance enrollment details to the insurer
// regarding a specified coverage.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/EnrollmentRequest
//   - Source File: StructureDefinition-EnrollmentRequest.json
type EnrollmentRequest struct {

	// Patient Resource.
	Candidate *fhir.Reference `fhirpath:"candidate"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// Reference to the program or plan identification, underwriter or payor.
	Coverage *fhir.Reference `fhirpath:"coverage"`

	// The date when this resource was created.
	Created *fhir.DateTime `fhirpath:"created"`

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

	// The Response business identifier.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The Insurer who is target of the request.
	Insurer *fhir.Reference `fhirpath:"insurer"`

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

	// The practitioner who is responsible for the services rendered to the
	// patient.
	Provider *fhir.Reference `fhirpath:"provider"`

	// The status of the resource instance.
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

// GetCandidate returns the value of the field Candidate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetCandidate() *fhir.Reference {
	if er == nil {
		return nil
	}
	return er.Candidate
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetContained() []fhir.Resource {
	if er == nil {
		return nil
	}
	return er.Contained
}

// GetCoverage returns the value of the field Coverage.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetCoverage() *fhir.Reference {
	if er == nil {
		return nil
	}
	return er.Coverage
}

// GetCreated returns the value of the field Created.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetCreated() *fhir.DateTime {
	if er == nil {
		return nil
	}
	return er.Created
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetExtension() []*fhir.Extension {
	if er == nil {
		return nil
	}
	return er.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetID() string {
	if er == nil {
		return ""
	}
	return er.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetIdentifier() []*fhir.Identifier {
	if er == nil {
		return nil
	}
	return er.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetImplicitRules() *fhir.URI {
	if er == nil {
		return nil
	}
	return er.ImplicitRules
}

// GetInsurer returns the value of the field Insurer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetInsurer() *fhir.Reference {
	if er == nil {
		return nil
	}
	return er.Insurer
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetLanguage() *fhir.Code {
	if er == nil {
		return nil
	}
	return er.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetMeta() *fhir.Meta {
	if er == nil {
		return nil
	}
	return er.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetModifierExtension() []*fhir.Extension {
	if er == nil {
		return nil
	}
	return er.ModifierExtension
}

// GetProvider returns the value of the field Provider.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetProvider() *fhir.Reference {
	if er == nil {
		return nil
	}
	return er.Provider
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetStatus() *fhir.Code {
	if er == nil {
		return nil
	}
	return er.Status
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentRequest) GetText() *fhir.Narrative {
	if er == nil {
		return nil
	}
	return er.Text
}

func (er *EnrollmentRequest) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (er *EnrollmentRequest) UnmarshalJSON(data []byte) error {
	var raw struct {
		Candidate *fhir.Reference   `json:"candidate"`
		Contained []fhir.Resource   `json:"contained"`
		Coverage  *fhir.Reference   `json:"coverage"`
		Created   *fhir.DateTime    `json:"created"`
		Extension []*fhir.Extension `json:"extension"`

		ID                string             `json:"id"`
		Identifier        []*fhir.Identifier `json:"identifier"`
		ImplicitRules     *fhir.URI          `json:"implicitRules"`
		Insurer           *fhir.Reference    `json:"insurer"`
		Language          *fhir.Code         `json:"language"`
		Meta              *fhir.Meta         `json:"meta"`
		ModifierExtension []*fhir.Extension  `json:"modifierExtension"`
		Provider          *fhir.Reference    `json:"provider"`
		Status            *fhir.Code         `json:"status"`
		Text              *fhir.Narrative    `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	er.Candidate = raw.Candidate
	er.Contained = raw.Contained
	er.Coverage = raw.Coverage
	er.Created = raw.Created
	er.Extension = raw.Extension
	er.ID = raw.ID
	er.Identifier = raw.Identifier
	er.ImplicitRules = raw.ImplicitRules
	er.Insurer = raw.Insurer
	er.Language = raw.Language
	er.Meta = raw.Meta
	er.ModifierExtension = raw.ModifierExtension
	er.Provider = raw.Provider
	er.Status = raw.Status
	er.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*EnrollmentRequest)(nil)
var _ json.Unmarshaler = (*EnrollmentRequest)(nil)
