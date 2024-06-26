// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package healthcareservice

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// The details of a healthcare service available at a location.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/HealthcareService
//   - Source File: StructureDefinition-HealthcareService.json
type HealthcareService struct {

	// This flag is used to mark the record to not be used. This is not used when a
	// center is closed for maintenance, or for holidays, the notAvailable period
	// is to be used for this.
	Active *fhir.Boolean `fhirpath:"active"`

	// Indicates whether or not a prospective consumer will require an appointment
	// for a particular service at a site to be provided by the Organization.
	// Indicates if an appointment is required for access to this service.
	AppointmentRequired *fhir.Boolean `fhirpath:"appointmentRequired"`

	// A description of site availability exceptions, e.g. public holiday
	// availability. Succinctly describing all possible exceptions to normal site
	// availability as details in the available Times and not available Times.
	AvailabilityExceptions *fhir.String `fhirpath:"availabilityExceptions"`

	// A collection of times that the Service Site is available.
	AvailableTime []*HealthcareServiceAvailableTime `fhirpath:"availableTime"`

	// Identifies the broad category of service being performed or delivered.
	Category []*fhir.CodeableConcept `fhirpath:"category"`

	// Collection of characteristics (attributes).
	Characteristic []*fhir.CodeableConcept `fhirpath:"characteristic"`

	// Any additional description of the service and/or any specific issues not
	// covered by the other attributes, which can be displayed as further detail
	// under the serviceName.
	Comment *fhir.String `fhirpath:"comment"`

	// Some services are specifically made available in multiple languages, this
	// property permits a directory to declare the languages this is offered in.
	// Typically this is only provided where a service operates in communities with
	// mixed languages used.
	Communication []*fhir.CodeableConcept `fhirpath:"communication"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The location(s) that this service is available to (not where the service is
	// provided).
	CoverageArea []*fhir.Reference `fhirpath:"coverageArea"`

	// Does this service have specific eligibility requirements that need to be met
	// in order to use the service?
	Eligibility []*HealthcareServiceEligibility `fhirpath:"eligibility"`

	// Technical endpoints providing access to services operated for the specific
	// healthcare services defined at this resource.
	Endpoint []*fhir.Reference `fhirpath:"endpoint"`

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Extra details about the service that can't be placed in the other fields.
	ExtraDetails *fhir.Markdown `fhirpath:"extraDetails"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// External identifiers for this item.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// The location(s) where this healthcare service may be provided.
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

	// Further description of the service as it would be presented to a consumer
	// while searching.
	Name *fhir.String `fhirpath:"name"`

	// The HealthcareService is not available during this period of time due to the
	// provided reason.
	NotAvailable []*HealthcareServiceNotAvailable `fhirpath:"notAvailable"`

	// If there is a photo/symbol associated with this HealthcareService, it may be
	// included here to facilitate quick identification of the service in a list.
	Photo *fhir.Attachment `fhirpath:"photo"`

	// Programs that this service is applicable to.
	Program []*fhir.CodeableConcept `fhirpath:"program"`

	// The organization that provides this healthcare service.
	ProvidedBy *fhir.Reference `fhirpath:"providedBy"`

	// Ways that the service accepts referrals, if this is not provided then it is
	// implied that no referral is required.
	ReferralMethod []*fhir.CodeableConcept `fhirpath:"referralMethod"`

	// The code(s) that detail the conditions under which the healthcare service is
	// available/offered.
	ServiceProvisionCode []*fhir.CodeableConcept `fhirpath:"serviceProvisionCode"`

	// Collection of specialties handled by the service site. This is more of a
	// medical term.
	Specialty []*fhir.CodeableConcept `fhirpath:"specialty"`

	// List of contacts related to this specific healthcare service.
	Telecom []*fhir.ContactPoint `fhirpath:"telecom"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// The specific type of service that may be delivered or performed.
	Type []*fhir.CodeableConcept `fhirpath:"type"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetActive returns the value of the field Active.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetActive() *fhir.Boolean {
	if hs == nil {
		return nil
	}
	return hs.Active
}

// GetAppointmentRequired returns the value of the field AppointmentRequired.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetAppointmentRequired() *fhir.Boolean {
	if hs == nil {
		return nil
	}
	return hs.AppointmentRequired
}

// GetAvailabilityExceptions returns the value of the field AvailabilityExceptions.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetAvailabilityExceptions() *fhir.String {
	if hs == nil {
		return nil
	}
	return hs.AvailabilityExceptions
}

// GetAvailableTime returns the value of the field AvailableTime.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetAvailableTime() []*HealthcareServiceAvailableTime {
	if hs == nil {
		return nil
	}
	return hs.AvailableTime
}

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetCategory() []*fhir.CodeableConcept {
	if hs == nil {
		return nil
	}
	return hs.Category
}

// GetCharacteristic returns the value of the field Characteristic.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetCharacteristic() []*fhir.CodeableConcept {
	if hs == nil {
		return nil
	}
	return hs.Characteristic
}

// GetComment returns the value of the field Comment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetComment() *fhir.String {
	if hs == nil {
		return nil
	}
	return hs.Comment
}

// GetCommunication returns the value of the field Communication.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetCommunication() []*fhir.CodeableConcept {
	if hs == nil {
		return nil
	}
	return hs.Communication
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetContained() []fhir.Resource {
	if hs == nil {
		return nil
	}
	return hs.Contained
}

// GetCoverageArea returns the value of the field CoverageArea.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetCoverageArea() []*fhir.Reference {
	if hs == nil {
		return nil
	}
	return hs.CoverageArea
}

// GetEligibility returns the value of the field Eligibility.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetEligibility() []*HealthcareServiceEligibility {
	if hs == nil {
		return nil
	}
	return hs.Eligibility
}

// GetEndpoint returns the value of the field Endpoint.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetEndpoint() []*fhir.Reference {
	if hs == nil {
		return nil
	}
	return hs.Endpoint
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetExtension() []*fhir.Extension {
	if hs == nil {
		return nil
	}
	return hs.Extension
}

// GetExtraDetails returns the value of the field ExtraDetails.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetExtraDetails() *fhir.Markdown {
	if hs == nil {
		return nil
	}
	return hs.ExtraDetails
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetID() string {
	if hs == nil {
		return ""
	}
	return hs.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetIdentifier() []*fhir.Identifier {
	if hs == nil {
		return nil
	}
	return hs.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetImplicitRules() *fhir.URI {
	if hs == nil {
		return nil
	}
	return hs.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetLanguage() *fhir.Code {
	if hs == nil {
		return nil
	}
	return hs.Language
}

// GetLocation returns the value of the field Location.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetLocation() []*fhir.Reference {
	if hs == nil {
		return nil
	}
	return hs.Location
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetMeta() *fhir.Meta {
	if hs == nil {
		return nil
	}
	return hs.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetModifierExtension() []*fhir.Extension {
	if hs == nil {
		return nil
	}
	return hs.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetName() *fhir.String {
	if hs == nil {
		return nil
	}
	return hs.Name
}

// GetNotAvailable returns the value of the field NotAvailable.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetNotAvailable() []*HealthcareServiceNotAvailable {
	if hs == nil {
		return nil
	}
	return hs.NotAvailable
}

// GetPhoto returns the value of the field Photo.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetPhoto() *fhir.Attachment {
	if hs == nil {
		return nil
	}
	return hs.Photo
}

// GetProgram returns the value of the field Program.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetProgram() []*fhir.CodeableConcept {
	if hs == nil {
		return nil
	}
	return hs.Program
}

// GetProvidedBy returns the value of the field ProvidedBy.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetProvidedBy() *fhir.Reference {
	if hs == nil {
		return nil
	}
	return hs.ProvidedBy
}

// GetReferralMethod returns the value of the field ReferralMethod.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetReferralMethod() []*fhir.CodeableConcept {
	if hs == nil {
		return nil
	}
	return hs.ReferralMethod
}

// GetServiceProvisionCode returns the value of the field ServiceProvisionCode.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetServiceProvisionCode() []*fhir.CodeableConcept {
	if hs == nil {
		return nil
	}
	return hs.ServiceProvisionCode
}

// GetSpecialty returns the value of the field Specialty.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetSpecialty() []*fhir.CodeableConcept {
	if hs == nil {
		return nil
	}
	return hs.Specialty
}

// GetTelecom returns the value of the field Telecom.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetTelecom() []*fhir.ContactPoint {
	if hs == nil {
		return nil
	}
	return hs.Telecom
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetText() *fhir.Narrative {
	if hs == nil {
		return nil
	}
	return hs.Text
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hs *HealthcareService) GetType() []*fhir.CodeableConcept {
	if hs == nil {
		return nil
	}
	return hs.Type
}

// Times the Service Site is available// A collection of times that the Service Site is available.// More detailed availability information may be provided in associated
// Schedule/Slot resources.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-HealthcareService.json
type HealthcareServiceAvailableTime struct {

	// Is this always available? (hence times are irrelevant) e.g. 24 hour service.
	AllDay *fhir.Boolean `fhirpath:"allDay"`

	// The closing time of day. Note: If the AllDay flag is set, then this time is
	// ignored.
	AvailableEndTime *fhir.Time `fhirpath:"availableEndTime"`

	// The opening time of day. Note: If the AllDay flag is set, then this time is
	// ignored.
	AvailableStartTime *fhir.Time `fhirpath:"availableStartTime"`

	// Indicates which days of the week are available between the start and end
	// Times.
	DaysOfWeek []*fhir.Code `fhirpath:"daysOfWeek"`

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

// GetAllDay returns the value of the field AllDay.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hsat *HealthcareServiceAvailableTime) GetAllDay() *fhir.Boolean {
	if hsat == nil {
		return nil
	}
	return hsat.AllDay
}

// GetAvailableEndTime returns the value of the field AvailableEndTime.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hsat *HealthcareServiceAvailableTime) GetAvailableEndTime() *fhir.Time {
	if hsat == nil {
		return nil
	}
	return hsat.AvailableEndTime
}

// GetAvailableStartTime returns the value of the field AvailableStartTime.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hsat *HealthcareServiceAvailableTime) GetAvailableStartTime() *fhir.Time {
	if hsat == nil {
		return nil
	}
	return hsat.AvailableStartTime
}

// GetDaysOfWeek returns the value of the field DaysOfWeek.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hsat *HealthcareServiceAvailableTime) GetDaysOfWeek() []*fhir.Code {
	if hsat == nil {
		return nil
	}
	return hsat.DaysOfWeek
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hsat *HealthcareServiceAvailableTime) GetExtension() []*fhir.Extension {
	if hsat == nil {
		return nil
	}
	return hsat.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hsat *HealthcareServiceAvailableTime) GetID() string {
	if hsat == nil {
		return ""
	}
	return hsat.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hsat *HealthcareServiceAvailableTime) GetModifierExtension() []*fhir.Extension {
	if hsat == nil {
		return nil
	}
	return hsat.ModifierExtension
}

// Specific eligibility requirements required to use the service// Does this service have specific eligibility requirements that need to be met
// in order to use the service?
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-HealthcareService.json
type HealthcareServiceEligibility struct {

	// Coded value for the eligibility.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// Describes the eligibility conditions for the service.
	Comment *fhir.Markdown `fhirpath:"comment"`

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
func (hse *HealthcareServiceEligibility) GetCode() *fhir.CodeableConcept {
	if hse == nil {
		return nil
	}
	return hse.Code
}

// GetComment returns the value of the field Comment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hse *HealthcareServiceEligibility) GetComment() *fhir.Markdown {
	if hse == nil {
		return nil
	}
	return hse.Comment
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hse *HealthcareServiceEligibility) GetExtension() []*fhir.Extension {
	if hse == nil {
		return nil
	}
	return hse.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hse *HealthcareServiceEligibility) GetID() string {
	if hse == nil {
		return ""
	}
	return hse.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hse *HealthcareServiceEligibility) GetModifierExtension() []*fhir.Extension {
	if hse == nil {
		return nil
	}
	return hse.ModifierExtension
}

// Not available during this time due to provided reason// The HealthcareService is not available during this period of time due to the
// provided reason.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-HealthcareService.json
type HealthcareServiceNotAvailable struct {

	// The reason that can be presented to the user as to why this time is not
	// available.
	Description *fhir.String `fhirpath:"description"`

	// Service is not available (seasonally or for a public holiday) from this
	// date.
	During *fhir.Period `fhirpath:"during"`

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

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hsna *HealthcareServiceNotAvailable) GetDescription() *fhir.String {
	if hsna == nil {
		return nil
	}
	return hsna.Description
}

// GetDuring returns the value of the field During.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hsna *HealthcareServiceNotAvailable) GetDuring() *fhir.Period {
	if hsna == nil {
		return nil
	}
	return hsna.During
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hsna *HealthcareServiceNotAvailable) GetExtension() []*fhir.Extension {
	if hsna == nil {
		return nil
	}
	return hsna.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hsna *HealthcareServiceNotAvailable) GetID() string {
	if hsna == nil {
		return ""
	}
	return hsna.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (hsna *HealthcareServiceNotAvailable) GetModifierExtension() []*fhir.Extension {
	if hsna == nil {
		return nil
	}
	return hsna.ModifierExtension
}

func (hs *HealthcareService) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (hs *HealthcareService) UnmarshalJSON(data []byte) error {
	var raw struct {
		Active                 *fhir.Boolean                     `json:"active"`
		AppointmentRequired    *fhir.Boolean                     `json:"appointmentRequired"`
		AvailabilityExceptions *fhir.String                      `json:"availabilityExceptions"`
		AvailableTime          []*HealthcareServiceAvailableTime `json:"availableTime"`
		Category               []*fhir.CodeableConcept           `json:"category"`
		Characteristic         []*fhir.CodeableConcept           `json:"characteristic"`
		Comment                *fhir.String                      `json:"comment"`
		Communication          []*fhir.CodeableConcept           `json:"communication"`
		Contained              []fhir.Resource                   `json:"contained"`
		CoverageArea           []*fhir.Reference                 `json:"coverageArea"`
		Eligibility            []*HealthcareServiceEligibility   `json:"eligibility"`
		Endpoint               []*fhir.Reference                 `json:"endpoint"`
		Extension              []*fhir.Extension                 `json:"extension"`
		ExtraDetails           *fhir.Markdown                    `json:"extraDetails"`

		ID                   string                           `json:"id"`
		Identifier           []*fhir.Identifier               `json:"identifier"`
		ImplicitRules        *fhir.URI                        `json:"implicitRules"`
		Language             *fhir.Code                       `json:"language"`
		Location             []*fhir.Reference                `json:"location"`
		Meta                 *fhir.Meta                       `json:"meta"`
		ModifierExtension    []*fhir.Extension                `json:"modifierExtension"`
		Name                 *fhir.String                     `json:"name"`
		NotAvailable         []*HealthcareServiceNotAvailable `json:"notAvailable"`
		Photo                *fhir.Attachment                 `json:"photo"`
		Program              []*fhir.CodeableConcept          `json:"program"`
		ProvidedBy           *fhir.Reference                  `json:"providedBy"`
		ReferralMethod       []*fhir.CodeableConcept          `json:"referralMethod"`
		ServiceProvisionCode []*fhir.CodeableConcept          `json:"serviceProvisionCode"`
		Specialty            []*fhir.CodeableConcept          `json:"specialty"`
		Telecom              []*fhir.ContactPoint             `json:"telecom"`
		Text                 *fhir.Narrative                  `json:"text"`
		Type                 []*fhir.CodeableConcept          `json:"type"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	hs.Active = raw.Active
	hs.AppointmentRequired = raw.AppointmentRequired
	hs.AvailabilityExceptions = raw.AvailabilityExceptions
	hs.AvailableTime = raw.AvailableTime
	hs.Category = raw.Category
	hs.Characteristic = raw.Characteristic
	hs.Comment = raw.Comment
	hs.Communication = raw.Communication
	hs.Contained = raw.Contained
	hs.CoverageArea = raw.CoverageArea
	hs.Eligibility = raw.Eligibility
	hs.Endpoint = raw.Endpoint
	hs.Extension = raw.Extension
	hs.ExtraDetails = raw.ExtraDetails
	hs.ID = raw.ID
	hs.Identifier = raw.Identifier
	hs.ImplicitRules = raw.ImplicitRules
	hs.Language = raw.Language
	hs.Location = raw.Location
	hs.Meta = raw.Meta
	hs.ModifierExtension = raw.ModifierExtension
	hs.Name = raw.Name
	hs.NotAvailable = raw.NotAvailable
	hs.Photo = raw.Photo
	hs.Program = raw.Program
	hs.ProvidedBy = raw.ProvidedBy
	hs.ReferralMethod = raw.ReferralMethod
	hs.ServiceProvisionCode = raw.ServiceProvisionCode
	hs.Specialty = raw.Specialty
	hs.Telecom = raw.Telecom
	hs.Text = raw.Text
	hs.Type = raw.Type
	return nil
}

var _ json.Marshaler = (*HealthcareService)(nil)
var _ json.Unmarshaler = (*HealthcareService)(nil)

func (hsat *HealthcareServiceAvailableTime) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (hsat *HealthcareServiceAvailableTime) UnmarshalJSON(data []byte) error {
	var raw struct {
		AllDay             *fhir.Boolean     `json:"allDay"`
		AvailableEndTime   *fhir.Time        `json:"availableEndTime"`
		AvailableStartTime *fhir.Time        `json:"availableStartTime"`
		DaysOfWeek         []*fhir.Code      `json:"daysOfWeek"`
		Extension          []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	hsat.AllDay = raw.AllDay
	hsat.AvailableEndTime = raw.AvailableEndTime
	hsat.AvailableStartTime = raw.AvailableStartTime
	hsat.DaysOfWeek = raw.DaysOfWeek
	hsat.Extension = raw.Extension
	hsat.ID = raw.ID
	hsat.ModifierExtension = raw.ModifierExtension
	return nil
}

var _ json.Marshaler = (*HealthcareServiceAvailableTime)(nil)
var _ json.Unmarshaler = (*HealthcareServiceAvailableTime)(nil)

func (hse *HealthcareServiceEligibility) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (hse *HealthcareServiceEligibility) UnmarshalJSON(data []byte) error {
	var raw struct {
		Code      *fhir.CodeableConcept `json:"code"`
		Comment   *fhir.Markdown        `json:"comment"`
		Extension []*fhir.Extension     `json:"extension"`

		ID                string            `json:"id"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	hse.Code = raw.Code
	hse.Comment = raw.Comment
	hse.Extension = raw.Extension
	hse.ID = raw.ID
	hse.ModifierExtension = raw.ModifierExtension
	return nil
}

var _ json.Marshaler = (*HealthcareServiceEligibility)(nil)
var _ json.Unmarshaler = (*HealthcareServiceEligibility)(nil)

func (hsna *HealthcareServiceNotAvailable) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (hsna *HealthcareServiceNotAvailable) UnmarshalJSON(data []byte) error {
	var raw struct {
		Description *fhir.String      `json:"description"`
		During      *fhir.Period      `json:"during"`
		Extension   []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	hsna.Description = raw.Description
	hsna.During = raw.During
	hsna.Extension = raw.Extension
	hsna.ID = raw.ID
	hsna.ModifierExtension = raw.ModifierExtension
	return nil
}

var _ json.Marshaler = (*HealthcareServiceNotAvailable)(nil)
var _ json.Unmarshaler = (*HealthcareServiceNotAvailable)(nil)
