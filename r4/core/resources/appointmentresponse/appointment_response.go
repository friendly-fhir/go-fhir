// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package appointmentresponse

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// A reply to an appointment request for a patient and/or practitioner(s), such
// as a confirmation or rejection.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/AppointmentResponse
//   - Source File: StructureDefinition-AppointmentResponse.json
type AppointmentResponse struct {

	// A Person, Location, HealthcareService, or Device that is participating in
	// the appointment.
	Actor *fhir.Reference `fhirpath:"actor"`

	// Appointment that this response is replying to.
	Appointment *fhir.Reference `fhirpath:"appointment"`

	// Additional comments about the appointment.
	Comment *fhir.String `fhirpath:"comment"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// This may be either the same as the appointment request to confirm the
	// details of the appointment, or alternately a new time to request a
	// re-negotiation of the end time.
	End *fhir.Instant `fhirpath:"end"`

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

	// This records identifiers associated with this appointment response concern
	// that are defined by business processes and/ or used to refer to it when a
	// direct URL reference to the resource itself is not appropriate.
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

	// Participation status of the participant. When the status is declined or
	// tentative if the start/end times are different to the appointment, then
	// these times should be interpreted as a requested time change. When the
	// status is accepted, the times can either be the time of the appointment (as
	// a confirmation of the time) or can be empty.
	ParticipantStatus *fhir.Code `fhirpath:"participantStatus"`

	// Role of participant in the appointment.
	ParticipantType []*fhir.CodeableConcept `fhirpath:"participantType"`

	// Date/Time that the appointment is to take place, or requested new start
	// time.
	Start *fhir.Instant `fhirpath:"start"`

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

// GetActor returns the value of the field Actor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetActor() *fhir.Reference {
	if ar == nil {
		return nil
	}
	return ar.Actor
}

// GetAppointment returns the value of the field Appointment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetAppointment() *fhir.Reference {
	if ar == nil {
		return nil
	}
	return ar.Appointment
}

// GetComment returns the value of the field Comment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetComment() *fhir.String {
	if ar == nil {
		return nil
	}
	return ar.Comment
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetContained() []fhir.Resource {
	if ar == nil {
		return nil
	}
	return ar.Contained
}

// GetEnd returns the value of the field End.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetEnd() *fhir.Instant {
	if ar == nil {
		return nil
	}
	return ar.End
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetExtension() []*fhir.Extension {
	if ar == nil {
		return nil
	}
	return ar.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetID() string {
	if ar == nil {
		return ""
	}
	return ar.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetIdentifier() []*fhir.Identifier {
	if ar == nil {
		return nil
	}
	return ar.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetImplicitRules() *fhir.URI {
	if ar == nil {
		return nil
	}
	return ar.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetLanguage() *fhir.Code {
	if ar == nil {
		return nil
	}
	return ar.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetMeta() *fhir.Meta {
	if ar == nil {
		return nil
	}
	return ar.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetModifierExtension() []*fhir.Extension {
	if ar == nil {
		return nil
	}
	return ar.ModifierExtension
}

// GetParticipantStatus returns the value of the field ParticipantStatus.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetParticipantStatus() *fhir.Code {
	if ar == nil {
		return nil
	}
	return ar.ParticipantStatus
}

// GetParticipantType returns the value of the field ParticipantType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetParticipantType() []*fhir.CodeableConcept {
	if ar == nil {
		return nil
	}
	return ar.ParticipantType
}

// GetStart returns the value of the field Start.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetStart() *fhir.Instant {
	if ar == nil {
		return nil
	}
	return ar.Start
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ar *AppointmentResponse) GetText() *fhir.Narrative {
	if ar == nil {
		return nil
	}
	return ar.Text
}
