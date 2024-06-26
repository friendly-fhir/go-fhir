// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package servicerequest

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// A record of a request for service such as diagnostic investigations,
// treatments, or operations to be performed.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/ServiceRequest
//   - Source File: StructureDefinition-ServiceRequest.json
type ServiceRequest struct {

	// If a CodeableConcept is present, it indicates the pre-condition for
	// performing the service. For example "pain", "on flare-up", etc.
	AsNeeded fhir.Element `fhirpath:"asNeeded"`

	// When the request transitioned to being actionable.
	AuthoredOn *fhir.DateTime `fhirpath:"authoredOn"`

	// Plan/proposal/order fulfilled by this request.
	BasedOn []*fhir.Reference `fhirpath:"basedOn"`

	// Anatomic location where the procedure should be performed. This is the
	// target site.
	BodySite []*fhir.CodeableConcept `fhirpath:"bodySite"`

	// A code that classifies the service for searching, sorting and display
	// purposes (e.g. "Surgical Procedure").
	Category []*fhir.CodeableConcept `fhirpath:"category"`

	// A code that identifies a particular service (i.e., procedure, diagnostic
	// investigation, or panel of investigations) that have been requested.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// Set this to true if the record is saying that the service/procedure should
	// NOT be performed.
	DoNotPerform *fhir.Boolean `fhirpath:"doNotPerform"`

	// An encounter that provides additional information about the healthcare
	// context in which this request is made.
	Encounter *fhir.Reference `fhirpath:"encounter"`

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

	// Identifiers assigned to this order instance by the orderer and/or the
	// receiver and/or order fulfiller.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other
	// definition that is adhered to in whole or in part by this ServiceRequest.
	InstantiatesCanonical []*fhir.Canonical `fhirpath:"instantiatesCanonical"`

	// The URL pointing to an externally maintained protocol, guideline, orderset
	// or other definition that is adhered to in whole or in part by this
	// ServiceRequest.
	InstantiatesURI []*fhir.URI `fhirpath:"instantiatesUri"`

	// Insurance plans, coverage extensions, pre-authorizations and/or
	// pre-determinations that may be needed for delivering the requested service.
	Insurance []*fhir.Reference `fhirpath:"insurance"`

	// Whether the request is a proposal, plan, an original order or a reflex
	// order.
	Intent *fhir.Code `fhirpath:"intent"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// The preferred location(s) where the procedure should actually happen in
	// coded or free text form. E.g. at home or nursing day care center.
	LocationCode []*fhir.CodeableConcept `fhirpath:"locationCode"`

	// A reference to the the preferred location(s) where the procedure should
	// actually happen. E.g. at home or nursing day care center.
	LocationReference []*fhir.Reference `fhirpath:"locationReference"`

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

	// Any other notes and comments made about the service request. For example,
	// internal billing notes.
	Note []*fhir.Annotation `fhirpath:"note"`

	// The date/time at which the requested service should occur.
	Occurrence fhir.Element `fhirpath:"occurrence"`

	// Additional details and instructions about the how the services are to be
	// delivered. For example, and order for a urinary catheter may have an order
	// detail for an external or indwelling catheter, or an order for a bandage may
	// require additional instructions specifying how the bandage should be
	// applied.
	OrderDetail []*fhir.CodeableConcept `fhirpath:"orderDetail"`

	// Instructions in terms that are understood by the patient or consumer.
	PatientInstruction *fhir.String `fhirpath:"patientInstruction"`

	// The desired performer for doing the requested service. For example, the
	// surgeon, dermatopathologist, endoscopist, etc.
	Performer []*fhir.Reference `fhirpath:"performer"`

	// Desired type of performer for doing the requested service.
	PerformerType *fhir.CodeableConcept `fhirpath:"performerType"`

	// Indicates how quickly the ServiceRequest should be addressed with respect to
	// other requests.
	Priority *fhir.Code `fhirpath:"priority"`

	// An amount of service being requested which can be a quantity ( for example
	// $1,500 home modification), a ratio ( for example, 20 half day visits per
	// month), or a range (2.0 to 1.8 Gy per fraction).
	Quantity fhir.Element `fhirpath:"quantity"`

	// An explanation or justification for why this service is being requested in
	// coded or textual form. This is often for billing purposes. May relate to the
	// resources referred to in `supportingInfo`.
	ReasonCode []*fhir.CodeableConcept `fhirpath:"reasonCode"`

	// Indicates another resource that provides a justification for why this
	// service is being requested. May relate to the resources referred to in
	// `supportingInfo`.
	ReasonReference []*fhir.Reference `fhirpath:"reasonReference"`

	// Key events in the history of the request.
	RelevantHistory []*fhir.Reference `fhirpath:"relevantHistory"`

	// The request takes the place of the referenced completed or terminated
	// request(s).
	Replaces []*fhir.Reference `fhirpath:"replaces"`

	// The individual who initiated the request and has responsibility for its
	// activation.
	Requester *fhir.Reference `fhirpath:"requester"`

	// A shared identifier common to all service requests that were authorized more
	// or less simultaneously by a single author, representing the composite or
	// group identifier.
	Requisition *fhir.Identifier `fhirpath:"requisition"`

	// One or more specimens that the laboratory procedure will use.
	Specimen []*fhir.Reference `fhirpath:"specimen"`

	// The status of the order.
	Status *fhir.Code `fhirpath:"status"`

	// On whom or what the service is to be performed. This is usually a human
	// patient, but can also be requested on animals, groups of humans or animals,
	// devices such as dialysis machines, or even locations (typically for
	// environmental scans).
	Subject *fhir.Reference `fhirpath:"subject"`

	// Additional clinical information about the patient or specimen that may
	// influence the services or their interpretations. This information includes
	// diagnosis, clinical findings and other observations. In laboratory ordering
	// these are typically referred to as "ask at order entry questions (AOEs)".
	// This includes observations explicitly requested by the producer (filler) to
	// provide context or supporting information needed to complete the order. For
	// example, reporting the amount of inspired oxygen for blood gas measurements.
	SupportingInfo []*fhir.Reference `fhirpath:"supportingInfo"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	profileimpl.BaseServiceRequest
	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetAsNeeded returns the value of the field AsNeeded.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetAsNeeded() fhir.Element {
	if sr == nil {
		return nil
	}
	return sr.AsNeeded
}

// GetAsNeededBoolean returns the value of the field AsNeeded.
func (sr *ServiceRequest) GetAsNeededBoolean() *fhir.Boolean {
	if sr == nil {
		return nil
	}
	val, ok := sr.AsNeeded.(*fhir.Boolean)
	if !ok {
		return nil
	}
	return val
}

// GetAsNeededCodeableConcept returns the value of the field AsNeeded.
func (sr *ServiceRequest) GetAsNeededCodeableConcept() *fhir.CodeableConcept {
	if sr == nil {
		return nil
	}
	val, ok := sr.AsNeeded.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
} // GetAuthoredOn returns the value of the field AuthoredOn.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetAuthoredOn() *fhir.DateTime {
	if sr == nil {
		return nil
	}
	return sr.AuthoredOn
}

// GetBasedOn returns the value of the field BasedOn.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetBasedOn() []*fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.BasedOn
}

// GetBodySite returns the value of the field BodySite.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetBodySite() []*fhir.CodeableConcept {
	if sr == nil {
		return nil
	}
	return sr.BodySite
}

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetCategory() []*fhir.CodeableConcept {
	if sr == nil {
		return nil
	}
	return sr.Category
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetCode() *fhir.CodeableConcept {
	if sr == nil {
		return nil
	}
	return sr.Code
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetContained() []fhir.Resource {
	if sr == nil {
		return nil
	}
	return sr.Contained
}

// GetDoNotPerform returns the value of the field DoNotPerform.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetDoNotPerform() *fhir.Boolean {
	if sr == nil {
		return nil
	}
	return sr.DoNotPerform
}

// GetEncounter returns the value of the field Encounter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetEncounter() *fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.Encounter
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetExtension() []*fhir.Extension {
	if sr == nil {
		return nil
	}
	return sr.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetID() string {
	if sr == nil {
		return ""
	}
	return sr.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetIdentifier() []*fhir.Identifier {
	if sr == nil {
		return nil
	}
	return sr.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetImplicitRules() *fhir.URI {
	if sr == nil {
		return nil
	}
	return sr.ImplicitRules
}

// GetInstantiatesCanonical returns the value of the field InstantiatesCanonical.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetInstantiatesCanonical() []*fhir.Canonical {
	if sr == nil {
		return nil
	}
	return sr.InstantiatesCanonical
}

// GetInstantiatesURI returns the value of the field InstantiatesURI.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetInstantiatesURI() []*fhir.URI {
	if sr == nil {
		return nil
	}
	return sr.InstantiatesURI
}

// GetInsurance returns the value of the field Insurance.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetInsurance() []*fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.Insurance
}

// GetIntent returns the value of the field Intent.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetIntent() *fhir.Code {
	if sr == nil {
		return nil
	}
	return sr.Intent
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetLanguage() *fhir.Code {
	if sr == nil {
		return nil
	}
	return sr.Language
}

// GetLocationCode returns the value of the field LocationCode.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetLocationCode() []*fhir.CodeableConcept {
	if sr == nil {
		return nil
	}
	return sr.LocationCode
}

// GetLocationReference returns the value of the field LocationReference.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetLocationReference() []*fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.LocationReference
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetMeta() *fhir.Meta {
	if sr == nil {
		return nil
	}
	return sr.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetModifierExtension() []*fhir.Extension {
	if sr == nil {
		return nil
	}
	return sr.ModifierExtension
}

// GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetNote() []*fhir.Annotation {
	if sr == nil {
		return nil
	}
	return sr.Note
}

// GetOccurrence returns the value of the field Occurrence.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetOccurrence() fhir.Element {
	if sr == nil {
		return nil
	}
	return sr.Occurrence
}

// GetOccurrenceDateTime returns the value of the field Occurrence.
func (sr *ServiceRequest) GetOccurrenceDateTime() *fhir.DateTime {
	if sr == nil {
		return nil
	}
	val, ok := sr.Occurrence.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
}

// GetOccurrencePeriod returns the value of the field Occurrence.
func (sr *ServiceRequest) GetOccurrencePeriod() *fhir.Period {
	if sr == nil {
		return nil
	}
	val, ok := sr.Occurrence.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
}

// GetOccurrenceTiming returns the value of the field Occurrence.
func (sr *ServiceRequest) GetOccurrenceTiming() *fhir.Timing {
	if sr == nil {
		return nil
	}
	val, ok := sr.Occurrence.(*fhir.Timing)
	if !ok {
		return nil
	}
	return val
} // GetOrderDetail returns the value of the field OrderDetail.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetOrderDetail() []*fhir.CodeableConcept {
	if sr == nil {
		return nil
	}
	return sr.OrderDetail
}

// GetPatientInstruction returns the value of the field PatientInstruction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetPatientInstruction() *fhir.String {
	if sr == nil {
		return nil
	}
	return sr.PatientInstruction
}

// GetPerformer returns the value of the field Performer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetPerformer() []*fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.Performer
}

// GetPerformerType returns the value of the field PerformerType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetPerformerType() *fhir.CodeableConcept {
	if sr == nil {
		return nil
	}
	return sr.PerformerType
}

// GetPriority returns the value of the field Priority.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetPriority() *fhir.Code {
	if sr == nil {
		return nil
	}
	return sr.Priority
}

// GetQuantity returns the value of the field Quantity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetQuantity() fhir.Element {
	if sr == nil {
		return nil
	}
	return sr.Quantity
}

// GetQuantityQuantity returns the value of the field Quantity.
func (sr *ServiceRequest) GetQuantityQuantity() *fhir.Quantity {
	if sr == nil {
		return nil
	}
	val, ok := sr.Quantity.(*fhir.Quantity)
	if !ok {
		return nil
	}
	return val
}

// GetQuantityRatio returns the value of the field Quantity.
func (sr *ServiceRequest) GetQuantityRatio() *fhir.Ratio {
	if sr == nil {
		return nil
	}
	val, ok := sr.Quantity.(*fhir.Ratio)
	if !ok {
		return nil
	}
	return val
}

// GetQuantityRange returns the value of the field Quantity.
func (sr *ServiceRequest) GetQuantityRange() *fhir.Range {
	if sr == nil {
		return nil
	}
	val, ok := sr.Quantity.(*fhir.Range)
	if !ok {
		return nil
	}
	return val
} // GetReasonCode returns the value of the field ReasonCode.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetReasonCode() []*fhir.CodeableConcept {
	if sr == nil {
		return nil
	}
	return sr.ReasonCode
}

// GetReasonReference returns the value of the field ReasonReference.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetReasonReference() []*fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.ReasonReference
}

// GetRelevantHistory returns the value of the field RelevantHistory.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetRelevantHistory() []*fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.RelevantHistory
}

// GetReplaces returns the value of the field Replaces.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetReplaces() []*fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.Replaces
}

// GetRequester returns the value of the field Requester.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetRequester() *fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.Requester
}

// GetRequisition returns the value of the field Requisition.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetRequisition() *fhir.Identifier {
	if sr == nil {
		return nil
	}
	return sr.Requisition
}

// GetSpecimen returns the value of the field Specimen.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetSpecimen() []*fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.Specimen
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetStatus() *fhir.Code {
	if sr == nil {
		return nil
	}
	return sr.Status
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetSubject() *fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.Subject
}

// GetSupportingInfo returns the value of the field SupportingInfo.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetSupportingInfo() []*fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.SupportingInfo
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *ServiceRequest) GetText() *fhir.Narrative {
	if sr == nil {
		return nil
	}
	return sr.Text
}

func (sr *ServiceRequest) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (sr *ServiceRequest) UnmarshalJSON(data []byte) error {
	var raw struct {
		AsNeededBoolean         *fhir.Boolean           `json:"asNeededBoolean"`
		AsNeededCodeableConcept *fhir.CodeableConcept   `json:"asNeededCodeableConcept"`
		AuthoredOn              *fhir.DateTime          `json:"authoredOn"`
		BasedOn                 []*fhir.Reference       `json:"basedOn"`
		BodySite                []*fhir.CodeableConcept `json:"bodySite"`
		Category                []*fhir.CodeableConcept `json:"category"`
		Code                    *fhir.CodeableConcept   `json:"code"`
		Contained               []fhir.Resource         `json:"contained"`
		DoNotPerform            *fhir.Boolean           `json:"doNotPerform"`
		Encounter               *fhir.Reference         `json:"encounter"`
		Extension               []*fhir.Extension       `json:"extension"`

		ID                    string                  `json:"id"`
		Identifier            []*fhir.Identifier      `json:"identifier"`
		ImplicitRules         *fhir.URI               `json:"implicitRules"`
		InstantiatesCanonical []*fhir.Canonical       `json:"instantiatesCanonical"`
		InstantiatesURI       []*fhir.URI             `json:"instantiatesUri"`
		Insurance             []*fhir.Reference       `json:"insurance"`
		Intent                *fhir.Code              `json:"intent"`
		Language              *fhir.Code              `json:"language"`
		LocationCode          []*fhir.CodeableConcept `json:"locationCode"`
		LocationReference     []*fhir.Reference       `json:"locationReference"`
		Meta                  *fhir.Meta              `json:"meta"`
		ModifierExtension     []*fhir.Extension       `json:"modifierExtension"`
		Note                  []*fhir.Annotation      `json:"note"`
		OccurrenceDateTime    *fhir.DateTime          `json:"occurrenceDateTime"`
		OccurrencePeriod      *fhir.Period            `json:"occurrencePeriod"`
		OccurrenceTiming      *fhir.Timing            `json:"occurrenceTiming"`
		OrderDetail           []*fhir.CodeableConcept `json:"orderDetail"`
		PatientInstruction    *fhir.String            `json:"patientInstruction"`
		Performer             []*fhir.Reference       `json:"performer"`
		PerformerType         *fhir.CodeableConcept   `json:"performerType"`
		Priority              *fhir.Code              `json:"priority"`
		QuantityQuantity      *fhir.Quantity          `json:"quantityQuantity"`
		QuantityRatio         *fhir.Ratio             `json:"quantityRatio"`
		QuantityRange         *fhir.Range             `json:"quantityRange"`
		ReasonCode            []*fhir.CodeableConcept `json:"reasonCode"`
		ReasonReference       []*fhir.Reference       `json:"reasonReference"`
		RelevantHistory       []*fhir.Reference       `json:"relevantHistory"`
		Replaces              []*fhir.Reference       `json:"replaces"`
		Requester             *fhir.Reference         `json:"requester"`
		Requisition           *fhir.Identifier        `json:"requisition"`
		Specimen              []*fhir.Reference       `json:"specimen"`
		Status                *fhir.Code              `json:"status"`
		Subject               *fhir.Reference         `json:"subject"`
		SupportingInfo        []*fhir.Reference       `json:"supportingInfo"`
		Text                  *fhir.Narrative         `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	sr.AsNeeded, err = validate.SelectOneOf[fhir.Element]("ServiceRequest.asNeeded",
		raw.AsNeededBoolean,
		raw.AsNeededCodeableConcept)
	if err != nil {
		return err
	}
	sr.AuthoredOn = raw.AuthoredOn
	sr.BasedOn = raw.BasedOn
	sr.BodySite = raw.BodySite
	sr.Category = raw.Category
	sr.Code = raw.Code
	sr.Contained = raw.Contained
	sr.DoNotPerform = raw.DoNotPerform
	sr.Encounter = raw.Encounter
	sr.Extension = raw.Extension
	sr.ID = raw.ID
	sr.Identifier = raw.Identifier
	sr.ImplicitRules = raw.ImplicitRules
	sr.InstantiatesCanonical = raw.InstantiatesCanonical
	sr.InstantiatesURI = raw.InstantiatesURI
	sr.Insurance = raw.Insurance
	sr.Intent = raw.Intent
	sr.Language = raw.Language
	sr.LocationCode = raw.LocationCode
	sr.LocationReference = raw.LocationReference
	sr.Meta = raw.Meta
	sr.ModifierExtension = raw.ModifierExtension
	sr.Note = raw.Note
	sr.Occurrence, err = validate.SelectOneOf[fhir.Element]("ServiceRequest.occurrence",
		raw.OccurrenceDateTime,
		raw.OccurrencePeriod,
		raw.OccurrenceTiming)
	if err != nil {
		return err
	}
	sr.OrderDetail = raw.OrderDetail
	sr.PatientInstruction = raw.PatientInstruction
	sr.Performer = raw.Performer
	sr.PerformerType = raw.PerformerType
	sr.Priority = raw.Priority
	sr.Quantity, err = validate.SelectOneOf[fhir.Element]("ServiceRequest.quantity",
		raw.QuantityQuantity,
		raw.QuantityRatio,
		raw.QuantityRange)
	if err != nil {
		return err
	}
	sr.ReasonCode = raw.ReasonCode
	sr.ReasonReference = raw.ReasonReference
	sr.RelevantHistory = raw.RelevantHistory
	sr.Replaces = raw.Replaces
	sr.Requester = raw.Requester
	sr.Requisition = raw.Requisition
	sr.Specimen = raw.Specimen
	sr.Status = raw.Status
	sr.Subject = raw.Subject
	sr.SupportingInfo = raw.SupportingInfo
	sr.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*ServiceRequest)(nil)
var _ json.Unmarshaler = (*ServiceRequest)(nil)
