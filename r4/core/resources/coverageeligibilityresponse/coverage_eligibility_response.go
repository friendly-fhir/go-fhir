// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package coverageeligibilityresponse

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// This resource provides eligibility and plan details from the processing of
// an CoverageEligibilityRequest resource.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/CoverageEligibilityResponse
//   - Source File: StructureDefinition-CoverageEligibilityResponse.json
type CoverageEligibilityResponse struct {

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The date this resource was created.
	Created *fhir.DateTime `fhirpath:"created"`

	// A human readable description of the status of the adjudication.
	Disposition *fhir.String `fhirpath:"disposition"`

	// Errors encountered during the processing of the request.
	Error []*CoverageEligibilityResponseError `fhirpath:"error"`

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// A code for the form to be used for printing the content.
	Form *fhir.CodeableConcept `fhirpath:"form"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// A unique identifier assigned to this coverage eligiblity request.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// Financial instruments for reimbursement for the health care products and
	// services.
	Insurance []*CoverageEligibilityResponseInsurance `fhirpath:"insurance"`

	// The Insurer who issued the coverage in question and is the author of the
	// response.
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

	// The outcome of the request processing.
	Outcome *fhir.Code `fhirpath:"outcome"`

	// The party who is the beneficiary of the supplied coverage and for whom
	// eligibility is sought.
	Patient *fhir.Reference `fhirpath:"patient"`

	// A reference from the Insurer to which these services pertain to be used on
	// further communication and as proof that the request occurred.
	PreAuthRef *fhir.String `fhirpath:"preAuthRef"`

	// Code to specify whether requesting: prior authorization requirements for
	// some service categories or billing codes; benefits for coverages specified
	// or discovered; discovery and return of coverages for the patient; and/or
	// validation that the specified coverage is in-force at the date/period
	// specified or 'now' if not specified.
	Purpose []*fhir.Code `fhirpath:"purpose"`

	// Reference to the original request resource.
	Request *fhir.Reference `fhirpath:"request"`

	// The provider which is responsible for the request.
	Requestor *fhir.Reference `fhirpath:"requestor"`

	// The date or dates when the enclosed suite of services were performed or
	// completed.
	Serviced fhir.Element `fhirpath:"serviced"`

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
func (cer *CoverageEligibilityResponse) GetContained() []fhir.Resource {
	if cer == nil {
		return nil
	}
	return cer.Contained
}

// GetCreated returns the value of the field Created.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetCreated() *fhir.DateTime {
	if cer == nil {
		return nil
	}
	return cer.Created
}

// GetDisposition returns the value of the field Disposition.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetDisposition() *fhir.String {
	if cer == nil {
		return nil
	}
	return cer.Disposition
}

// GetError returns the value of the field Error.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetError() []*CoverageEligibilityResponseError {
	if cer == nil {
		return nil
	}
	return cer.Error
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetExtension() []*fhir.Extension {
	if cer == nil {
		return nil
	}
	return cer.Extension
}

// GetForm returns the value of the field Form.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetForm() *fhir.CodeableConcept {
	if cer == nil {
		return nil
	}
	return cer.Form
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetID() string {
	if cer == nil {
		return ""
	}
	return cer.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetIdentifier() []*fhir.Identifier {
	if cer == nil {
		return nil
	}
	return cer.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetImplicitRules() *fhir.URI {
	if cer == nil {
		return nil
	}
	return cer.ImplicitRules
}

// GetInsurance returns the value of the field Insurance.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetInsurance() []*CoverageEligibilityResponseInsurance {
	if cer == nil {
		return nil
	}
	return cer.Insurance
}

// GetInsurer returns the value of the field Insurer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetInsurer() *fhir.Reference {
	if cer == nil {
		return nil
	}
	return cer.Insurer
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetLanguage() *fhir.Code {
	if cer == nil {
		return nil
	}
	return cer.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetMeta() *fhir.Meta {
	if cer == nil {
		return nil
	}
	return cer.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetModifierExtension() []*fhir.Extension {
	if cer == nil {
		return nil
	}
	return cer.ModifierExtension
}

// GetOutcome returns the value of the field Outcome.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetOutcome() *fhir.Code {
	if cer == nil {
		return nil
	}
	return cer.Outcome
}

// GetPatient returns the value of the field Patient.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetPatient() *fhir.Reference {
	if cer == nil {
		return nil
	}
	return cer.Patient
}

// GetPreAuthRef returns the value of the field PreAuthRef.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetPreAuthRef() *fhir.String {
	if cer == nil {
		return nil
	}
	return cer.PreAuthRef
}

// GetPurpose returns the value of the field Purpose.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetPurpose() []*fhir.Code {
	if cer == nil {
		return nil
	}
	return cer.Purpose
}

// GetRequest returns the value of the field Request.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetRequest() *fhir.Reference {
	if cer == nil {
		return nil
	}
	return cer.Request
}

// GetRequestor returns the value of the field Requestor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetRequestor() *fhir.Reference {
	if cer == nil {
		return nil
	}
	return cer.Requestor
}

// GetServiced returns the value of the field Serviced.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetServiced() fhir.Element {
	if cer == nil {
		return nil
	}
	return cer.Serviced
}

// GetServicedDate returns the value of the field Serviced.
func (cer *CoverageEligibilityResponse) GetServicedDate() *fhir.Date {
	if cer == nil {
		return nil
	}
	val, ok := cer.Serviced.(*fhir.Date)
	if !ok {
		return nil
	}
	return val
}

// GetServicedPeriod returns the value of the field Serviced.
func (cer *CoverageEligibilityResponse) GetServicedPeriod() *fhir.Period {
	if cer == nil {
		return nil
	}
	val, ok := cer.Serviced.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
} // GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetStatus() *fhir.Code {
	if cer == nil {
		return nil
	}
	return cer.Status
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cer *CoverageEligibilityResponse) GetText() *fhir.Narrative {
	if cer == nil {
		return nil
	}
	return cer.Text
}

// Processing errors// Errors encountered during the processing of the request.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-CoverageEligibilityResponse.json
type CoverageEligibilityResponseError struct {

	// An error code,from a specified code system, which details why the
	// eligibility check could not be performed.
	Code *fhir.CodeableConcept `fhirpath:"code"`

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
func (cere *CoverageEligibilityResponseError) GetCode() *fhir.CodeableConcept {
	if cere == nil {
		return nil
	}
	return cere.Code
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cere *CoverageEligibilityResponseError) GetExtension() []*fhir.Extension {
	if cere == nil {
		return nil
	}
	return cere.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cere *CoverageEligibilityResponseError) GetID() string {
	if cere == nil {
		return ""
	}
	return cere.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cere *CoverageEligibilityResponseError) GetModifierExtension() []*fhir.Extension {
	if cere == nil {
		return nil
	}
	return cere.ModifierExtension
}

// Patient insurance information// Financial instruments for reimbursement for the health care products and
// services.// All insurance coverages for the patient which may be applicable for
// reimbursement, of the products and services listed in the claim, are
// typically provided in the claim to allow insurers to confirm the ordering of
// the insurance coverages relative to local 'coordination of benefit' rules.
// One coverage (and only one) with 'focal=true' is to be used in the
// adjudication of this claim. Coverages appearing before the focal Coverage in
// the list, and where 'subrogation=false', should provide a reference to the
// ClaimResponse containing the adjudication results of the prior claim.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-CoverageEligibilityResponse.json
type CoverageEligibilityResponseInsurance struct {

	// The term of the benefits documented in this response.
	BenefitPeriod *fhir.Period `fhirpath:"benefitPeriod"`

	// Reference to the insurance card level information contained in the Coverage
	// resource. The coverage issuing insurer will use these details to locate the
	// patient's actual coverage within the insurer's information system.
	Coverage *fhir.Reference `fhirpath:"coverage"`

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

	// Flag indicating if the coverage provided is inforce currently if no service
	// date(s) specified or for the whole duration of the service dates.
	Inforce *fhir.Boolean `fhirpath:"inforce"`

	// Benefits and optionally current balances, and authorization details by
	// category or service.
	Item []*CoverageEligibilityResponseInsuranceItem `fhirpath:"item"`

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

// GetBenefitPeriod returns the value of the field BenefitPeriod.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceri *CoverageEligibilityResponseInsurance) GetBenefitPeriod() *fhir.Period {
	if ceri == nil {
		return nil
	}
	return ceri.BenefitPeriod
}

// GetCoverage returns the value of the field Coverage.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceri *CoverageEligibilityResponseInsurance) GetCoverage() *fhir.Reference {
	if ceri == nil {
		return nil
	}
	return ceri.Coverage
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceri *CoverageEligibilityResponseInsurance) GetExtension() []*fhir.Extension {
	if ceri == nil {
		return nil
	}
	return ceri.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceri *CoverageEligibilityResponseInsurance) GetID() string {
	if ceri == nil {
		return ""
	}
	return ceri.ID
}

// GetInforce returns the value of the field Inforce.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceri *CoverageEligibilityResponseInsurance) GetInforce() *fhir.Boolean {
	if ceri == nil {
		return nil
	}
	return ceri.Inforce
}

// GetItem returns the value of the field Item.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceri *CoverageEligibilityResponseInsurance) GetItem() []*CoverageEligibilityResponseInsuranceItem {
	if ceri == nil {
		return nil
	}
	return ceri.Item
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceri *CoverageEligibilityResponseInsurance) GetModifierExtension() []*fhir.Extension {
	if ceri == nil {
		return nil
	}
	return ceri.ModifierExtension
}

// Benefits and authorization details// Benefits and optionally current balances, and authorization details by
// category or service.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-CoverageEligibilityResponse.json
type CoverageEligibilityResponseInsuranceItem struct {

	// A boolean flag indicating whether a preauthorization is required prior to
	// actual service delivery.
	AuthorizationRequired *fhir.Boolean `fhirpath:"authorizationRequired"`

	// Codes or comments regarding information or actions associated with the
	// preauthorization.
	AuthorizationSupporting []*fhir.CodeableConcept `fhirpath:"authorizationSupporting"`

	// A web location for obtaining requirements or descriptive information
	// regarding the preauthorization.
	AuthorizationURL *fhir.URI `fhirpath:"authorizationUrl"`

	// Benefits used to date.
	Benefit []*CoverageEligibilityResponseInsuranceItemBenefit `fhirpath:"benefit"`

	// Code to identify the general type of benefits under which products and
	// services are provided.
	Category *fhir.CodeableConcept `fhirpath:"category"`

	// A richer description of the benefit or services covered.
	Description *fhir.String `fhirpath:"description"`

	// True if the indicated class of service is excluded from the plan, missing or
	// False indicates the product or service is included in the coverage.
	Excluded *fhir.Boolean `fhirpath:"excluded"`

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

	// Item typification or modifiers codes to convey additional context for the
	// product or service.
	Modifier []*fhir.CodeableConcept `fhirpath:"modifier"`

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

	// A short name or tag for the benefit.
	Name *fhir.String `fhirpath:"name"`

	// Is a flag to indicate whether the benefits refer to in-network providers or
	// out-of-network providers.
	Network *fhir.CodeableConcept `fhirpath:"network"`

	// This contains the product, service, drug or other billing code for the item.
	ProductOrService *fhir.CodeableConcept `fhirpath:"productOrService"`

	// The practitioner who is eligible for the provision of the product or
	// service.
	Provider *fhir.Reference `fhirpath:"provider"`

	// The term or period of the values such as 'maximum lifetime benefit' or
	// 'maximum annual visits'.
	Term *fhir.CodeableConcept `fhirpath:"term"`

	// Indicates if the benefits apply to an individual or to the family.
	Unit *fhir.CodeableConcept `fhirpath:"unit"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAuthorizationRequired returns the value of the field AuthorizationRequired.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetAuthorizationRequired() *fhir.Boolean {
	if cerii == nil {
		return nil
	}
	return cerii.AuthorizationRequired
}

// GetAuthorizationSupporting returns the value of the field AuthorizationSupporting.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetAuthorizationSupporting() []*fhir.CodeableConcept {
	if cerii == nil {
		return nil
	}
	return cerii.AuthorizationSupporting
}

// GetAuthorizationURL returns the value of the field AuthorizationURL.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetAuthorizationURL() *fhir.URI {
	if cerii == nil {
		return nil
	}
	return cerii.AuthorizationURL
}

// GetBenefit returns the value of the field Benefit.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetBenefit() []*CoverageEligibilityResponseInsuranceItemBenefit {
	if cerii == nil {
		return nil
	}
	return cerii.Benefit
}

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetCategory() *fhir.CodeableConcept {
	if cerii == nil {
		return nil
	}
	return cerii.Category
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetDescription() *fhir.String {
	if cerii == nil {
		return nil
	}
	return cerii.Description
}

// GetExcluded returns the value of the field Excluded.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetExcluded() *fhir.Boolean {
	if cerii == nil {
		return nil
	}
	return cerii.Excluded
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetExtension() []*fhir.Extension {
	if cerii == nil {
		return nil
	}
	return cerii.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetID() string {
	if cerii == nil {
		return ""
	}
	return cerii.ID
}

// GetModifier returns the value of the field Modifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetModifier() []*fhir.CodeableConcept {
	if cerii == nil {
		return nil
	}
	return cerii.Modifier
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetModifierExtension() []*fhir.Extension {
	if cerii == nil {
		return nil
	}
	return cerii.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetName() *fhir.String {
	if cerii == nil {
		return nil
	}
	return cerii.Name
}

// GetNetwork returns the value of the field Network.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetNetwork() *fhir.CodeableConcept {
	if cerii == nil {
		return nil
	}
	return cerii.Network
}

// GetProductOrService returns the value of the field ProductOrService.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetProductOrService() *fhir.CodeableConcept {
	if cerii == nil {
		return nil
	}
	return cerii.ProductOrService
}

// GetProvider returns the value of the field Provider.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetProvider() *fhir.Reference {
	if cerii == nil {
		return nil
	}
	return cerii.Provider
}

// GetTerm returns the value of the field Term.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetTerm() *fhir.CodeableConcept {
	if cerii == nil {
		return nil
	}
	return cerii.Term
}

// GetUnit returns the value of the field Unit.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cerii *CoverageEligibilityResponseInsuranceItem) GetUnit() *fhir.CodeableConcept {
	if cerii == nil {
		return nil
	}
	return cerii.Unit
}

// Benefit Summary// Benefits used to date.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-CoverageEligibilityResponse.json
type CoverageEligibilityResponseInsuranceItemBenefit struct {

	// The quantity of the benefit which is permitted under the coverage.
	Allowed fhir.Element `fhirpath:"allowed"`

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

	// Classification of benefit being provided.
	Type *fhir.CodeableConcept `fhirpath:"type"`

	// The quantity of the benefit which have been consumed to date.
	Used fhir.Element `fhirpath:"used"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAllowed returns the value of the field Allowed.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceriib *CoverageEligibilityResponseInsuranceItemBenefit) GetAllowed() fhir.Element {
	if ceriib == nil {
		return nil
	}
	return ceriib.Allowed
}

// GetAllowedUnsignedInt returns the value of the field Allowed.
func (ceriib *CoverageEligibilityResponseInsuranceItemBenefit) GetAllowedUnsignedInt() *fhir.UnsignedInt {
	if ceriib == nil {
		return nil
	}
	val, ok := ceriib.Allowed.(*fhir.UnsignedInt)
	if !ok {
		return nil
	}
	return val
}

// GetAllowedString returns the value of the field Allowed.
func (ceriib *CoverageEligibilityResponseInsuranceItemBenefit) GetAllowedString() *fhir.String {
	if ceriib == nil {
		return nil
	}
	val, ok := ceriib.Allowed.(*fhir.String)
	if !ok {
		return nil
	}
	return val
}

// GetAllowedMoney returns the value of the field Allowed.
func (ceriib *CoverageEligibilityResponseInsuranceItemBenefit) GetAllowedMoney() *fhir.Money {
	if ceriib == nil {
		return nil
	}
	val, ok := ceriib.Allowed.(*fhir.Money)
	if !ok {
		return nil
	}
	return val
} // GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceriib *CoverageEligibilityResponseInsuranceItemBenefit) GetExtension() []*fhir.Extension {
	if ceriib == nil {
		return nil
	}
	return ceriib.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceriib *CoverageEligibilityResponseInsuranceItemBenefit) GetID() string {
	if ceriib == nil {
		return ""
	}
	return ceriib.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceriib *CoverageEligibilityResponseInsuranceItemBenefit) GetModifierExtension() []*fhir.Extension {
	if ceriib == nil {
		return nil
	}
	return ceriib.ModifierExtension
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceriib *CoverageEligibilityResponseInsuranceItemBenefit) GetType() *fhir.CodeableConcept {
	if ceriib == nil {
		return nil
	}
	return ceriib.Type
}

// GetUsed returns the value of the field Used.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ceriib *CoverageEligibilityResponseInsuranceItemBenefit) GetUsed() fhir.Element {
	if ceriib == nil {
		return nil
	}
	return ceriib.Used
}

// GetUsedUnsignedInt returns the value of the field Used.
func (ceriib *CoverageEligibilityResponseInsuranceItemBenefit) GetUsedUnsignedInt() *fhir.UnsignedInt {
	if ceriib == nil {
		return nil
	}
	val, ok := ceriib.Used.(*fhir.UnsignedInt)
	if !ok {
		return nil
	}
	return val
}

// GetUsedString returns the value of the field Used.
func (ceriib *CoverageEligibilityResponseInsuranceItemBenefit) GetUsedString() *fhir.String {
	if ceriib == nil {
		return nil
	}
	val, ok := ceriib.Used.(*fhir.String)
	if !ok {
		return nil
	}
	return val
}

// GetUsedMoney returns the value of the field Used.
func (ceriib *CoverageEligibilityResponseInsuranceItemBenefit) GetUsedMoney() *fhir.Money {
	if ceriib == nil {
		return nil
	}
	val, ok := ceriib.Used.(*fhir.Money)
	if !ok {
		return nil
	}
	return val
}