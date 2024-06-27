// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package schedule

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// A container for slots of time that may be available for booking
// appointments.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Schedule
//   - Source File: StructureDefinition-Schedule.json
type Schedule struct {

	// Whether this schedule record is in active use or should not be used (such as
	// was entered in error).
	Active *fhir.Boolean `fhirpath:"active"`

	// Slots that reference this schedule resource provide the availability details
	// to these referenced resource(s).
	Actor []*fhir.Reference `fhirpath:"actor"`

	// Comments on the availability to describe any extended information. Such as
	// custom constraints on the slots that may be associated.
	Comment *fhir.String `fhirpath:"comment"`

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

	// External Ids for this item.
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

	// The period of time that the slots that reference this Schedule resource
	// cover (even if none exist). These cover the amount of time that an
	// organization's planning horizon; the interval for which they are currently
	// accepting appointments. This does not define a "template" for planning
	// outside these dates.
	PlanningHorizon *fhir.Period `fhirpath:"planningHorizon"`

	// A broad categorization of the service that is to be performed during this
	// appointment.
	ServiceCategory []*fhir.CodeableConcept `fhirpath:"serviceCategory"`

	// The specific service that is to be performed during this appointment.
	ServiceType []*fhir.CodeableConcept `fhirpath:"serviceType"`

	// The specialty of a practitioner that would be required to perform the
	// service requested in this appointment.
	Specialty []*fhir.CodeableConcept `fhirpath:"specialty"`

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
func (s *Schedule) GetActive() *fhir.Boolean {
	if s == nil {
		return nil
	}
	return s.Active
}

// GetActor returns the value of the field Actor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetActor() []*fhir.Reference {
	if s == nil {
		return nil
	}
	return s.Actor
}

// GetComment returns the value of the field Comment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetComment() *fhir.String {
	if s == nil {
		return nil
	}
	return s.Comment
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetContained() []fhir.Resource {
	if s == nil {
		return nil
	}
	return s.Contained
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetExtension() []*fhir.Extension {
	if s == nil {
		return nil
	}
	return s.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetID() string {
	if s == nil {
		return ""
	}
	return s.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetIdentifier() []*fhir.Identifier {
	if s == nil {
		return nil
	}
	return s.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetImplicitRules() *fhir.URI {
	if s == nil {
		return nil
	}
	return s.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetLanguage() *fhir.Code {
	if s == nil {
		return nil
	}
	return s.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetMeta() *fhir.Meta {
	if s == nil {
		return nil
	}
	return s.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetModifierExtension() []*fhir.Extension {
	if s == nil {
		return nil
	}
	return s.ModifierExtension
}

// GetPlanningHorizon returns the value of the field PlanningHorizon.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetPlanningHorizon() *fhir.Period {
	if s == nil {
		return nil
	}
	return s.PlanningHorizon
}

// GetServiceCategory returns the value of the field ServiceCategory.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetServiceCategory() []*fhir.CodeableConcept {
	if s == nil {
		return nil
	}
	return s.ServiceCategory
}

// GetServiceType returns the value of the field ServiceType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetServiceType() []*fhir.CodeableConcept {
	if s == nil {
		return nil
	}
	return s.ServiceType
}

// GetSpecialty returns the value of the field Specialty.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetSpecialty() []*fhir.CodeableConcept {
	if s == nil {
		return nil
	}
	return s.Specialty
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Schedule) GetText() *fhir.Narrative {
	if s == nil {
		return nil
	}
	return s.Text
}

func (s *Schedule) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (s *Schedule) UnmarshalJSON(data []byte) error {
	var raw struct {
		Active    *fhir.Boolean     `json:"active"`
		Actor     []*fhir.Reference `json:"actor"`
		Comment   *fhir.String      `json:"comment"`
		Contained []fhir.Resource   `json:"contained"`
		Extension []*fhir.Extension `json:"extension"`

		ID                string                  `json:"id"`
		Identifier        []*fhir.Identifier      `json:"identifier"`
		ImplicitRules     *fhir.URI               `json:"implicitRules"`
		Language          *fhir.Code              `json:"language"`
		Meta              *fhir.Meta              `json:"meta"`
		ModifierExtension []*fhir.Extension       `json:"modifierExtension"`
		PlanningHorizon   *fhir.Period            `json:"planningHorizon"`
		ServiceCategory   []*fhir.CodeableConcept `json:"serviceCategory"`
		ServiceType       []*fhir.CodeableConcept `json:"serviceType"`
		Specialty         []*fhir.CodeableConcept `json:"specialty"`
		Text              *fhir.Narrative         `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	s.Active = raw.Active
	s.Actor = raw.Actor
	s.Comment = raw.Comment
	s.Contained = raw.Contained
	s.Extension = raw.Extension
	s.ID = raw.ID
	s.Identifier = raw.Identifier
	s.ImplicitRules = raw.ImplicitRules
	s.Language = raw.Language
	s.Meta = raw.Meta
	s.ModifierExtension = raw.ModifierExtension
	s.PlanningHorizon = raw.PlanningHorizon
	s.ServiceCategory = raw.ServiceCategory
	s.ServiceType = raw.ServiceType
	s.Specialty = raw.Specialty
	s.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*Schedule)(nil)
var _ json.Unmarshaler = (*Schedule)(nil)
