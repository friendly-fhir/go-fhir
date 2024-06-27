// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package organizationaffiliation

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Defines an affiliation/assotiation/relationship between 2 distinct
// oganizations, that is not a part-of relationship/sub-division relationship.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/OrganizationAffiliation
//   - Source File: StructureDefinition-OrganizationAffiliation.json
type OrganizationAffiliation struct {

	// Whether this organization affiliation record is in active use.
	Active *fhir.Boolean `fhirpath:"active"`

	// Definition of the role the participatingOrganization plays in the
	// association.
	Code []*fhir.CodeableConcept `fhirpath:"code"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// Technical endpoints providing access to services operated for this role.
	Endpoint []*fhir.Reference `fhirpath:"endpoint"`

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Healthcare services provided through the role.
	HealthcareService []*fhir.Reference `fhirpath:"healthcareService"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// Business identifiers that are specific to this role.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// The location(s) at which the role occurs.
	Location []*fhir.Reference `fhirpath:"location"`

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

	// Health insurance provider network in which the participatingOrganization
	// provides the role's services (if defined) at the indicated locations (if
	// defined).
	Network []*fhir.Reference `fhirpath:"network"`

	// Organization where the role is available (primary organization/has members).
	Organization *fhir.Reference `fhirpath:"organization"`

	// The Participating Organization provides/performs the role(s) defined by the
	// code to the Primary Organization (e.g. providing services or is a member
	// of).
	ParticipatingOrganization *fhir.Reference `fhirpath:"participatingOrganization"`

	// The period during which the participatingOrganization is affiliated with the
	// primary organization.
	Period *fhir.Period `fhirpath:"period"`

	// Specific specialty of the participatingOrganization in the context of the
	// role.
	Specialty []*fhir.CodeableConcept `fhirpath:"specialty"`

	// Contact details at the participatingOrganization relevant to this
	// Affiliation.
	Telecom []*fhir.ContactPoint `fhirpath:"telecom"`

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

// GetActive returns the value of the field Active.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetActive() *fhir.Boolean {
	if oa == nil {
		return nil
	}
	return oa.Active
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetCode() []*fhir.CodeableConcept {
	if oa == nil {
		return nil
	}
	return oa.Code
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetContained() []fhir.Resource {
	if oa == nil {
		return nil
	}
	return oa.Contained
}

// GetEndpoint returns the value of the field Endpoint.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetEndpoint() []*fhir.Reference {
	if oa == nil {
		return nil
	}
	return oa.Endpoint
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetExtension() []*fhir.Extension {
	if oa == nil {
		return nil
	}
	return oa.Extension
}

// GetHealthcareService returns the value of the field HealthcareService.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetHealthcareService() []*fhir.Reference {
	if oa == nil {
		return nil
	}
	return oa.HealthcareService
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetID() string {
	if oa == nil {
		return ""
	}
	return oa.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetIdentifier() []*fhir.Identifier {
	if oa == nil {
		return nil
	}
	return oa.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetImplicitRules() *fhir.URI {
	if oa == nil {
		return nil
	}
	return oa.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetLanguage() *fhir.Code {
	if oa == nil {
		return nil
	}
	return oa.Language
}

// GetLocation returns the value of the field Location.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetLocation() []*fhir.Reference {
	if oa == nil {
		return nil
	}
	return oa.Location
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetMeta() *fhir.Meta {
	if oa == nil {
		return nil
	}
	return oa.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetModifierExtension() []*fhir.Extension {
	if oa == nil {
		return nil
	}
	return oa.ModifierExtension
}

// GetNetwork returns the value of the field Network.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetNetwork() []*fhir.Reference {
	if oa == nil {
		return nil
	}
	return oa.Network
}

// GetOrganization returns the value of the field Organization.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetOrganization() *fhir.Reference {
	if oa == nil {
		return nil
	}
	return oa.Organization
}

// GetParticipatingOrganization returns the value of the field ParticipatingOrganization.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetParticipatingOrganization() *fhir.Reference {
	if oa == nil {
		return nil
	}
	return oa.ParticipatingOrganization
}

// GetPeriod returns the value of the field Period.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetPeriod() *fhir.Period {
	if oa == nil {
		return nil
	}
	return oa.Period
}

// GetSpecialty returns the value of the field Specialty.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetSpecialty() []*fhir.CodeableConcept {
	if oa == nil {
		return nil
	}
	return oa.Specialty
}

// GetTelecom returns the value of the field Telecom.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetTelecom() []*fhir.ContactPoint {
	if oa == nil {
		return nil
	}
	return oa.Telecom
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (oa *OrganizationAffiliation) GetText() *fhir.Narrative {
	if oa == nil {
		return nil
	}
	return oa.Text
}

func (oa *OrganizationAffiliation) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (oa *OrganizationAffiliation) UnmarshalJSON(data []byte) error {
	var raw struct {
		Active            *fhir.Boolean           `json:"active"`
		Code              []*fhir.CodeableConcept `json:"code"`
		Contained         []fhir.Resource         `json:"contained"`
		Endpoint          []*fhir.Reference       `json:"endpoint"`
		Extension         []*fhir.Extension       `json:"extension"`
		HealthcareService []*fhir.Reference       `json:"healthcareService"`

		ID                        string                  `json:"id"`
		Identifier                []*fhir.Identifier      `json:"identifier"`
		ImplicitRules             *fhir.URI               `json:"implicitRules"`
		Language                  *fhir.Code              `json:"language"`
		Location                  []*fhir.Reference       `json:"location"`
		Meta                      *fhir.Meta              `json:"meta"`
		ModifierExtension         []*fhir.Extension       `json:"modifierExtension"`
		Network                   []*fhir.Reference       `json:"network"`
		Organization              *fhir.Reference         `json:"organization"`
		ParticipatingOrganization *fhir.Reference         `json:"participatingOrganization"`
		Period                    *fhir.Period            `json:"period"`
		Specialty                 []*fhir.CodeableConcept `json:"specialty"`
		Telecom                   []*fhir.ContactPoint    `json:"telecom"`
		Text                      *fhir.Narrative         `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	oa.Active = raw.Active
	oa.Code = raw.Code
	oa.Contained = raw.Contained
	oa.Endpoint = raw.Endpoint
	oa.Extension = raw.Extension
	oa.HealthcareService = raw.HealthcareService
	oa.ID = raw.ID
	oa.Identifier = raw.Identifier
	oa.ImplicitRules = raw.ImplicitRules
	oa.Language = raw.Language
	oa.Location = raw.Location
	oa.Meta = raw.Meta
	oa.ModifierExtension = raw.ModifierExtension
	oa.Network = raw.Network
	oa.Organization = raw.Organization
	oa.ParticipatingOrganization = raw.ParticipatingOrganization
	oa.Period = raw.Period
	oa.Specialty = raw.Specialty
	oa.Telecom = raw.Telecom
	oa.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*OrganizationAffiliation)(nil)
var _ json.Unmarshaler = (*OrganizationAffiliation)(nil)
