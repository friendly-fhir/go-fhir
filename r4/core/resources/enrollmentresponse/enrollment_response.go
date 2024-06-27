// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package enrollmentresponse

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// This resource provides enrollment and plan details from the processing of an
// EnrollmentRequest resource.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/EnrollmentResponse
//   - Source File: StructureDefinition-EnrollmentResponse.json
type EnrollmentResponse struct {

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The date when the enclosed suite of services were performed or completed.
	Created *fhir.DateTime `fhirpath:"created"`

	// A description of the status of the adjudication.
	Disposition *fhir.String `fhirpath:"disposition"`

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

	// The Insurer who produced this adjudicated response.
	Organization *fhir.Reference `fhirpath:"organization"`

	// Processing status: error, complete.
	Outcome *fhir.Code `fhirpath:"outcome"`

	// Original request resource reference.
	Request *fhir.Reference `fhirpath:"request"`

	// The practitioner who is responsible for the services rendered to the
	// patient.
	RequestProvider *fhir.Reference `fhirpath:"requestProvider"`

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

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetContained() []fhir.Resource {
	if er == nil {
		return nil
	}
	return er.Contained
}

// GetCreated returns the value of the field Created.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetCreated() *fhir.DateTime {
	if er == nil {
		return nil
	}
	return er.Created
}

// GetDisposition returns the value of the field Disposition.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetDisposition() *fhir.String {
	if er == nil {
		return nil
	}
	return er.Disposition
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetExtension() []*fhir.Extension {
	if er == nil {
		return nil
	}
	return er.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetID() string {
	if er == nil {
		return ""
	}
	return er.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetIdentifier() []*fhir.Identifier {
	if er == nil {
		return nil
	}
	return er.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetImplicitRules() *fhir.URI {
	if er == nil {
		return nil
	}
	return er.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetLanguage() *fhir.Code {
	if er == nil {
		return nil
	}
	return er.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetMeta() *fhir.Meta {
	if er == nil {
		return nil
	}
	return er.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetModifierExtension() []*fhir.Extension {
	if er == nil {
		return nil
	}
	return er.ModifierExtension
}

// GetOrganization returns the value of the field Organization.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetOrganization() *fhir.Reference {
	if er == nil {
		return nil
	}
	return er.Organization
}

// GetOutcome returns the value of the field Outcome.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetOutcome() *fhir.Code {
	if er == nil {
		return nil
	}
	return er.Outcome
}

// GetRequest returns the value of the field Request.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetRequest() *fhir.Reference {
	if er == nil {
		return nil
	}
	return er.Request
}

// GetRequestProvider returns the value of the field RequestProvider.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetRequestProvider() *fhir.Reference {
	if er == nil {
		return nil
	}
	return er.RequestProvider
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetStatus() *fhir.Code {
	if er == nil {
		return nil
	}
	return er.Status
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (er *EnrollmentResponse) GetText() *fhir.Narrative {
	if er == nil {
		return nil
	}
	return er.Text
}

func (er *EnrollmentResponse) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (er *EnrollmentResponse) UnmarshalJSON(data []byte) error {
	var raw struct {
		Contained   []fhir.Resource   `json:"contained"`
		Created     *fhir.DateTime    `json:"created"`
		Disposition *fhir.String      `json:"disposition"`
		Extension   []*fhir.Extension `json:"extension"`

		ID                string             `json:"id"`
		Identifier        []*fhir.Identifier `json:"identifier"`
		ImplicitRules     *fhir.URI          `json:"implicitRules"`
		Language          *fhir.Code         `json:"language"`
		Meta              *fhir.Meta         `json:"meta"`
		ModifierExtension []*fhir.Extension  `json:"modifierExtension"`
		Organization      *fhir.Reference    `json:"organization"`
		Outcome           *fhir.Code         `json:"outcome"`
		Request           *fhir.Reference    `json:"request"`
		RequestProvider   *fhir.Reference    `json:"requestProvider"`
		Status            *fhir.Code         `json:"status"`
		Text              *fhir.Narrative    `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	er.Contained = raw.Contained
	er.Created = raw.Created
	er.Disposition = raw.Disposition
	er.Extension = raw.Extension
	er.ID = raw.ID
	er.Identifier = raw.Identifier
	er.ImplicitRules = raw.ImplicitRules
	er.Language = raw.Language
	er.Meta = raw.Meta
	er.ModifierExtension = raw.ModifierExtension
	er.Organization = raw.Organization
	er.Outcome = raw.Outcome
	er.Request = raw.Request
	er.RequestProvider = raw.RequestProvider
	er.Status = raw.Status
	er.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*EnrollmentResponse)(nil)
var _ json.Unmarshaler = (*EnrollmentResponse)(nil)
