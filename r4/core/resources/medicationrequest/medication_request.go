// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package medicationrequest

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// An order or request for both supply of the medication and the instructions
// for administration of the medication to a patient. The resource is called
// "MedicationRequest" rather than "MedicationPrescription" or
// "MedicationOrder" to generalize the use across inpatient and outpatient
// settings, including care plans, etc., and to harmonize with workflow
// patterns.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/MedicationRequest
//   - Source File: StructureDefinition-MedicationRequest.json
type MedicationRequest struct {

	// The date (and perhaps time) when the prescription was initially written or
	// authored on.
	AuthoredOn *fhir.DateTime `fhirpath:"authoredOn"`

	// A plan or request that is fulfilled in whole or in part by this medication
	// request.
	BasedOn []*fhir.Reference `fhirpath:"basedOn"`

	// Indicates the type of medication request (for example, where the medication
	// is expected to be consumed or administered (i.e. inpatient or outpatient)).
	Category []*fhir.CodeableConcept `fhirpath:"category"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The description of the overall patte3rn of the administration of the
	// medication to the patient.
	CourseOfTherapyType *fhir.CodeableConcept `fhirpath:"courseOfTherapyType"`

	// Indicates an actual or potential clinical issue with or between one or more
	// active or proposed clinical actions for a patient; e.g. Drug-drug
	// interaction, duplicate therapy, dosage alert etc.
	DetectedIssue []*fhir.Reference `fhirpath:"detectedIssue"`

	// Indicates the specific details for the dispense or medication supply part of
	// a medication request (also known as a Medication Prescription or Medication
	// Order). Note that this information is not always sent with the order. There
	// may be in some settings (e.g. hospitals) institutional or system support for
	// completing the dispense details in the pharmacy department.
	DispenseRequest *MedicationRequestDispenseRequest `fhirpath:"dispenseRequest"`

	// If true indicates that the provider is asking for the medication request not
	// to occur.
	DoNotPerform *fhir.Boolean `fhirpath:"doNotPerform"`

	// Indicates how the medication is to be used by the patient.
	DosageInstruction []*fhir.Dosage `fhirpath:"dosageInstruction"`

	// The Encounter during which this [x] was created or to which the creation of
	// this record is tightly associated.
	Encounter *fhir.Reference `fhirpath:"encounter"`

	// Links to Provenance records for past versions of this resource or fulfilling
	// request or event resources that identify key state transitions or updates
	// that are likely to be relevant to a user looking at the current version of
	// the resource.
	EventHistory []*fhir.Reference `fhirpath:"eventHistory"`

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// A shared identifier common to all requests that were authorized more or less
	// simultaneously by a single author, representing the identifier of the
	// requisition or prescription.
	GroupIdentifier *fhir.Identifier `fhirpath:"groupIdentifier"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// Identifiers associated with this medication request that are defined by
	// business processes and/or used to refer to it when a direct URL reference to
	// the resource itself is not appropriate. They are business identifiers
	// assigned to this resource by the performer or other systems and remain
	// constant as the resource is updated and propagates from server to server.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The URL pointing to a protocol, guideline, orderset, or other definition
	// that is adhered to in whole or in part by this MedicationRequest.
	InstantiatesCanonical []*fhir.Canonical `fhirpath:"instantiatesCanonical"`

	// The URL pointing to an externally maintained protocol, guideline, orderset
	// or other definition that is adhered to in whole or in part by this
	// MedicationRequest.
	InstantiatesURI []*fhir.URI `fhirpath:"instantiatesUri"`

	// Insurance plans, coverage extensions, pre-authorizations and/or
	// pre-determinations that may be required for delivering the requested
	// service.
	Insurance []*fhir.Reference `fhirpath:"insurance"`

	// Whether the request is a proposal, plan, or an original order.
	Intent *fhir.Code `fhirpath:"intent"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// Identifies the medication being requested. This is a link to a resource that
	// represents the medication which may be the details of the medication or
	// simply an attribute carrying a code that identifies the medication from a
	// known list of medications.
	Medication fhir.Element `fhirpath:"medication"`

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

	// Extra information about the prescription that could not be conveyed by the
	// other attributes.
	Note []*fhir.Annotation `fhirpath:"note"`

	// The specified desired performer of the medication treatment (e.g. the
	// performer of the medication administration).
	Performer *fhir.Reference `fhirpath:"performer"`

	// Indicates the type of performer of the administration of the medication.
	PerformerType *fhir.CodeableConcept `fhirpath:"performerType"`

	// A link to a resource representing an earlier order related order or
	// prescription.
	PriorPrescription *fhir.Reference `fhirpath:"priorPrescription"`

	// Indicates how quickly the Medication Request should be addressed with
	// respect to other requests.
	Priority *fhir.Code `fhirpath:"priority"`

	// The reason or the indication for ordering or not ordering the medication.
	ReasonCode []*fhir.CodeableConcept `fhirpath:"reasonCode"`

	// Condition or observation that supports why the medication was ordered.
	ReasonReference []*fhir.Reference `fhirpath:"reasonReference"`

	// The person who entered the order on behalf of another individual for example
	// in the case of a verbal or a telephone order.
	Recorder *fhir.Reference `fhirpath:"recorder"`

	// Indicates if this record was captured as a secondary 'reported' record
	// rather than as an original primary source-of-truth record. It may also
	// indicate the source of the report.
	Reported fhir.Element `fhirpath:"reported"`

	// The individual, organization, or device that initiated the request and has
	// responsibility for its activation.
	Requester *fhir.Reference `fhirpath:"requester"`

	// A code specifying the current state of the order. Generally, this will be
	// active or completed state.
	Status *fhir.Code `fhirpath:"status"`

	// Captures the reason for the current state of the MedicationRequest.
	StatusReason *fhir.CodeableConcept `fhirpath:"statusReason"`

	// A link to a resource representing the person or set of individuals to whom
	// the medication will be given.
	Subject *fhir.Reference `fhirpath:"subject"`

	// Indicates whether or not substitution can or should be part of the dispense.
	// In some cases, substitution must happen, in other cases substitution must
	// not happen. This block explains the prescriber's intent. If nothing is
	// specified substitution may be done.
	Substitution *MedicationRequestSubstitution `fhirpath:"substitution"`

	// Include additional information (for example, patient height and weight) that
	// supports the ordering of the medication.
	SupportingInformation []*fhir.Reference `fhirpath:"supportingInformation"`

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

// GetAuthoredOn returns the value of the field AuthoredOn.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetAuthoredOn() *fhir.DateTime {
	if mr == nil {
		return nil
	}
	return mr.AuthoredOn
}

// GetBasedOn returns the value of the field BasedOn.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetBasedOn() []*fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.BasedOn
}

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetCategory() []*fhir.CodeableConcept {
	if mr == nil {
		return nil
	}
	return mr.Category
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetContained() []fhir.Resource {
	if mr == nil {
		return nil
	}
	return mr.Contained
}

// GetCourseOfTherapyType returns the value of the field CourseOfTherapyType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetCourseOfTherapyType() *fhir.CodeableConcept {
	if mr == nil {
		return nil
	}
	return mr.CourseOfTherapyType
}

// GetDetectedIssue returns the value of the field DetectedIssue.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetDetectedIssue() []*fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.DetectedIssue
}

// GetDispenseRequest returns the value of the field DispenseRequest.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetDispenseRequest() *MedicationRequestDispenseRequest {
	if mr == nil {
		return nil
	}
	return mr.DispenseRequest
}

// GetDoNotPerform returns the value of the field DoNotPerform.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetDoNotPerform() *fhir.Boolean {
	if mr == nil {
		return nil
	}
	return mr.DoNotPerform
}

// GetDosageInstruction returns the value of the field DosageInstruction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetDosageInstruction() []*fhir.Dosage {
	if mr == nil {
		return nil
	}
	return mr.DosageInstruction
}

// GetEncounter returns the value of the field Encounter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetEncounter() *fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.Encounter
}

// GetEventHistory returns the value of the field EventHistory.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetEventHistory() []*fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.EventHistory
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetExtension() []*fhir.Extension {
	if mr == nil {
		return nil
	}
	return mr.Extension
}

// GetGroupIdentifier returns the value of the field GroupIdentifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetGroupIdentifier() *fhir.Identifier {
	if mr == nil {
		return nil
	}
	return mr.GroupIdentifier
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetID() string {
	if mr == nil {
		return ""
	}
	return mr.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetIdentifier() []*fhir.Identifier {
	if mr == nil {
		return nil
	}
	return mr.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetImplicitRules() *fhir.URI {
	if mr == nil {
		return nil
	}
	return mr.ImplicitRules
}

// GetInstantiatesCanonical returns the value of the field InstantiatesCanonical.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetInstantiatesCanonical() []*fhir.Canonical {
	if mr == nil {
		return nil
	}
	return mr.InstantiatesCanonical
}

// GetInstantiatesURI returns the value of the field InstantiatesURI.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetInstantiatesURI() []*fhir.URI {
	if mr == nil {
		return nil
	}
	return mr.InstantiatesURI
}

// GetInsurance returns the value of the field Insurance.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetInsurance() []*fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.Insurance
}

// GetIntent returns the value of the field Intent.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetIntent() *fhir.Code {
	if mr == nil {
		return nil
	}
	return mr.Intent
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetLanguage() *fhir.Code {
	if mr == nil {
		return nil
	}
	return mr.Language
}

// GetMedication returns the value of the field Medication.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetMedication() fhir.Element {
	if mr == nil {
		return nil
	}
	return mr.Medication
}

// GetMedicationCodeableConcept returns the value of the field Medication.
func (mr *MedicationRequest) GetMedicationCodeableConcept() *fhir.CodeableConcept {
	if mr == nil {
		return nil
	}
	val, ok := mr.Medication.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetMedicationReference returns the value of the field Medication.
func (mr *MedicationRequest) GetMedicationReference() *fhir.Reference {
	if mr == nil {
		return nil
	}
	val, ok := mr.Medication.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
} // GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetMeta() *fhir.Meta {
	if mr == nil {
		return nil
	}
	return mr.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetModifierExtension() []*fhir.Extension {
	if mr == nil {
		return nil
	}
	return mr.ModifierExtension
}

// GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetNote() []*fhir.Annotation {
	if mr == nil {
		return nil
	}
	return mr.Note
}

// GetPerformer returns the value of the field Performer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetPerformer() *fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.Performer
}

// GetPerformerType returns the value of the field PerformerType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetPerformerType() *fhir.CodeableConcept {
	if mr == nil {
		return nil
	}
	return mr.PerformerType
}

// GetPriorPrescription returns the value of the field PriorPrescription.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetPriorPrescription() *fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.PriorPrescription
}

// GetPriority returns the value of the field Priority.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetPriority() *fhir.Code {
	if mr == nil {
		return nil
	}
	return mr.Priority
}

// GetReasonCode returns the value of the field ReasonCode.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetReasonCode() []*fhir.CodeableConcept {
	if mr == nil {
		return nil
	}
	return mr.ReasonCode
}

// GetReasonReference returns the value of the field ReasonReference.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetReasonReference() []*fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.ReasonReference
}

// GetRecorder returns the value of the field Recorder.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetRecorder() *fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.Recorder
}

// GetReported returns the value of the field Reported.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetReported() fhir.Element {
	if mr == nil {
		return nil
	}
	return mr.Reported
}

// GetReportedBoolean returns the value of the field Reported.
func (mr *MedicationRequest) GetReportedBoolean() *fhir.Boolean {
	if mr == nil {
		return nil
	}
	val, ok := mr.Reported.(*fhir.Boolean)
	if !ok {
		return nil
	}
	return val
}

// GetReportedReference returns the value of the field Reported.
func (mr *MedicationRequest) GetReportedReference() *fhir.Reference {
	if mr == nil {
		return nil
	}
	val, ok := mr.Reported.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
} // GetRequester returns the value of the field Requester.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetRequester() *fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.Requester
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetStatus() *fhir.Code {
	if mr == nil {
		return nil
	}
	return mr.Status
}

// GetStatusReason returns the value of the field StatusReason.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetStatusReason() *fhir.CodeableConcept {
	if mr == nil {
		return nil
	}
	return mr.StatusReason
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetSubject() *fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.Subject
}

// GetSubstitution returns the value of the field Substitution.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetSubstitution() *MedicationRequestSubstitution {
	if mr == nil {
		return nil
	}
	return mr.Substitution
}

// GetSupportingInformation returns the value of the field SupportingInformation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetSupportingInformation() []*fhir.Reference {
	if mr == nil {
		return nil
	}
	return mr.SupportingInformation
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mr *MedicationRequest) GetText() *fhir.Narrative {
	if mr == nil {
		return nil
	}
	return mr.Text
}

// Medication supply authorization// Indicates the specific details for the dispense or medication supply part of
// a medication request (also known as a Medication Prescription or Medication
// Order). Note that this information is not always sent with the order. There
// may be in some settings (e.g. hospitals) institutional or system support for
// completing the dispense details in the pharmacy department.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MedicationRequest.json
type MedicationRequestDispenseRequest struct {

	// The minimum period of time that must occur between dispenses of the
	// medication.
	DispenseInterval *fhir.Duration `fhirpath:"dispenseInterval"`

	// Identifies the period time over which the supplied product is expected to be
	// used, or the length of time the dispense is expected to last.
	ExpectedSupplyDuration *fhir.Duration `fhirpath:"expectedSupplyDuration"`

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

	// Indicates the quantity or duration for the first dispense of the medication.
	InitialFill *MedicationRequestDispenseRequestInitialFill `fhirpath:"initialFill"`

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

	// An integer indicating the number of times, in addition to the original
	// dispense, (aka refills or repeats) that the patient can receive the
	// prescribed medication. Usage Notes: This integer does not include the
	// original order dispense. This means that if an order indicates dispense 30
	// tablets plus "3 repeats", then the order can be dispensed a total of 4 times
	// and the patient can receive a total of 120 tablets. A prescriber may
	// explicitly say that zero refills are permitted after the initial dispense.
	NumberOfRepeatsAllowed *fhir.UnsignedInt `fhirpath:"numberOfRepeatsAllowed"`

	// Indicates the intended dispensing Organization specified by the prescriber.
	Performer *fhir.Reference `fhirpath:"performer"`

	// The amount that is to be dispensed for one fill.
	Quantity *fhir.Quantity `fhirpath:"quantity"`

	// This indicates the validity period of a prescription (stale dating the
	// Prescription).
	ValidityPeriod *fhir.Period `fhirpath:"validityPeriod"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetDispenseInterval returns the value of the field DispenseInterval.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdr *MedicationRequestDispenseRequest) GetDispenseInterval() *fhir.Duration {
	if mrdr == nil {
		return nil
	}
	return mrdr.DispenseInterval
}

// GetExpectedSupplyDuration returns the value of the field ExpectedSupplyDuration.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdr *MedicationRequestDispenseRequest) GetExpectedSupplyDuration() *fhir.Duration {
	if mrdr == nil {
		return nil
	}
	return mrdr.ExpectedSupplyDuration
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdr *MedicationRequestDispenseRequest) GetExtension() []*fhir.Extension {
	if mrdr == nil {
		return nil
	}
	return mrdr.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdr *MedicationRequestDispenseRequest) GetID() string {
	if mrdr == nil {
		return ""
	}
	return mrdr.ID
}

// GetInitialFill returns the value of the field InitialFill.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdr *MedicationRequestDispenseRequest) GetInitialFill() *MedicationRequestDispenseRequestInitialFill {
	if mrdr == nil {
		return nil
	}
	return mrdr.InitialFill
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdr *MedicationRequestDispenseRequest) GetModifierExtension() []*fhir.Extension {
	if mrdr == nil {
		return nil
	}
	return mrdr.ModifierExtension
}

// GetNumberOfRepeatsAllowed returns the value of the field NumberOfRepeatsAllowed.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdr *MedicationRequestDispenseRequest) GetNumberOfRepeatsAllowed() *fhir.UnsignedInt {
	if mrdr == nil {
		return nil
	}
	return mrdr.NumberOfRepeatsAllowed
}

// GetPerformer returns the value of the field Performer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdr *MedicationRequestDispenseRequest) GetPerformer() *fhir.Reference {
	if mrdr == nil {
		return nil
	}
	return mrdr.Performer
}

// GetQuantity returns the value of the field Quantity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdr *MedicationRequestDispenseRequest) GetQuantity() *fhir.Quantity {
	if mrdr == nil {
		return nil
	}
	return mrdr.Quantity
}

// GetValidityPeriod returns the value of the field ValidityPeriod.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdr *MedicationRequestDispenseRequest) GetValidityPeriod() *fhir.Period {
	if mrdr == nil {
		return nil
	}
	return mrdr.ValidityPeriod
}

// First fill details// Indicates the quantity or duration for the first dispense of the medication.// If populating this element, either the quantity or the duration must be
// included.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MedicationRequest.json
type MedicationRequestDispenseRequestInitialFill struct {

	// The length of time that the first dispense is expected to last.
	Duration *fhir.Duration `fhirpath:"duration"`

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

	// The amount or quantity to provide as part of the first dispense.
	Quantity *fhir.Quantity `fhirpath:"quantity"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetDuration returns the value of the field Duration.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdrif *MedicationRequestDispenseRequestInitialFill) GetDuration() *fhir.Duration {
	if mrdrif == nil {
		return nil
	}
	return mrdrif.Duration
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdrif *MedicationRequestDispenseRequestInitialFill) GetExtension() []*fhir.Extension {
	if mrdrif == nil {
		return nil
	}
	return mrdrif.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdrif *MedicationRequestDispenseRequestInitialFill) GetID() string {
	if mrdrif == nil {
		return ""
	}
	return mrdrif.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdrif *MedicationRequestDispenseRequestInitialFill) GetModifierExtension() []*fhir.Extension {
	if mrdrif == nil {
		return nil
	}
	return mrdrif.ModifierExtension
}

// GetQuantity returns the value of the field Quantity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrdrif *MedicationRequestDispenseRequestInitialFill) GetQuantity() *fhir.Quantity {
	if mrdrif == nil {
		return nil
	}
	return mrdrif.Quantity
}

// Any restrictions on medication substitution// Indicates whether or not substitution can or should be part of the dispense.
// In some cases, substitution must happen, in other cases substitution must
// not happen. This block explains the prescriber's intent. If nothing is
// specified substitution may be done.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MedicationRequest.json
type MedicationRequestSubstitution struct {

	// True if the prescriber allows a different drug to be dispensed from what was
	// prescribed.
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

	// Indicates the reason for the substitution, or why substitution must or must
	// not be performed.
	Reason *fhir.CodeableConcept `fhirpath:"reason"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAllowed returns the value of the field Allowed.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrs *MedicationRequestSubstitution) GetAllowed() fhir.Element {
	if mrs == nil {
		return nil
	}
	return mrs.Allowed
}

// GetAllowedBoolean returns the value of the field Allowed.
func (mrs *MedicationRequestSubstitution) GetAllowedBoolean() *fhir.Boolean {
	if mrs == nil {
		return nil
	}
	val, ok := mrs.Allowed.(*fhir.Boolean)
	if !ok {
		return nil
	}
	return val
}

// GetAllowedCodeableConcept returns the value of the field Allowed.
func (mrs *MedicationRequestSubstitution) GetAllowedCodeableConcept() *fhir.CodeableConcept {
	if mrs == nil {
		return nil
	}
	val, ok := mrs.Allowed.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
} // GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrs *MedicationRequestSubstitution) GetExtension() []*fhir.Extension {
	if mrs == nil {
		return nil
	}
	return mrs.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrs *MedicationRequestSubstitution) GetID() string {
	if mrs == nil {
		return ""
	}
	return mrs.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrs *MedicationRequestSubstitution) GetModifierExtension() []*fhir.Extension {
	if mrs == nil {
		return nil
	}
	return mrs.ModifierExtension
}

// GetReason returns the value of the field Reason.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mrs *MedicationRequestSubstitution) GetReason() *fhir.CodeableConcept {
	if mrs == nil {
		return nil
	}
	return mrs.Reason
}