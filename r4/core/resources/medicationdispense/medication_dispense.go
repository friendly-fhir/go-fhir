// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package medicationdispense

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Indicates that a medication product is to be or has been dispensed for a
// named person/patient. This includes a description of the medication product
// (supply) provided and the instructions for administering the medication. The
// medication dispense is the result of a pharmacy system responding to a
// medication order.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/MedicationDispense
//   - Source File: StructureDefinition-MedicationDispense.json
type MedicationDispense struct {

	// Indicates the medication order that is being dispensed against.
	AuthorizingPrescription []*fhir.Reference `fhirpath:"authorizingPrescription"`

	// Indicates the type of medication dispense (for example, where the medication
	// is expected to be consumed or administered (i.e. inpatient or outpatient)).
	Category *fhir.CodeableConcept `fhirpath:"category"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The encounter or episode of care that establishes the context for this
	// event.
	Context *fhir.Reference `fhirpath:"context"`

	// The amount of medication expressed as a timing amount.
	DaysSupply *fhir.Quantity `fhirpath:"daysSupply"`

	// Identification of the facility/location where the medication was shipped to,
	// as part of the dispense event.
	Destination *fhir.Reference `fhirpath:"destination"`

	// Indicates an actual or potential clinical issue with or between one or more
	// active or proposed clinical actions for a patient; e.g. drug-drug
	// interaction, duplicate therapy, dosage alert etc.
	DetectedIssue []*fhir.Reference `fhirpath:"detectedIssue"`

	// Indicates how the medication is to be used by the patient.
	DosageInstruction []*fhir.Dosage `fhirpath:"dosageInstruction"`

	// A summary of the events of interest that have occurred, such as when the
	// dispense was verified.
	EventHistory []*fhir.Reference `fhirpath:"eventHistory"`

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

	// Identifiers associated with this Medication Dispense that are defined by
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

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// The principal physical location where the dispense was performed.
	Location *fhir.Reference `fhirpath:"location"`

	// Identifies the medication being administered. This is either a link to a
	// resource representing the details of the medication or a simple attribute
	// carrying a code that identifies the medication from a known list of
	// medications.
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

	// Extra information about the dispense that could not be conveyed in the other
	// attributes.
	Note []*fhir.Annotation `fhirpath:"note"`

	// The procedure that trigger the dispense.
	PartOf []*fhir.Reference `fhirpath:"partOf"`

	// Indicates who or what performed the event.
	Performer []*MedicationDispensePerformer `fhirpath:"performer"`

	// The amount of medication that has been dispensed. Includes unit of measure.
	Quantity *fhir.Quantity `fhirpath:"quantity"`

	// Identifies the person who picked up the medication. This will usually be a
	// patient or their caregiver, but some cases exist where it can be a
	// healthcare professional.
	Receiver []*fhir.Reference `fhirpath:"receiver"`

	// A code specifying the state of the set of dispense events.
	Status *fhir.Code `fhirpath:"status"`

	// Indicates the reason why a dispense was not performed.
	StatusReason fhir.Element `fhirpath:"statusReason"`

	// A link to a resource representing the person or the group to whom the
	// medication will be given.
	Subject *fhir.Reference `fhirpath:"subject"`

	// Indicates whether or not substitution was made as part of the dispense. In
	// some cases, substitution will be expected but does not happen, in other
	// cases substitution is not expected but does happen. This block explains what
	// substitution did or did not happen and why. If nothing is specified,
	// substitution was not done.
	Substitution *MedicationDispenseSubstitution `fhirpath:"substitution"`

	// Additional information that supports the medication being dispensed.
	SupportingInformation []*fhir.Reference `fhirpath:"supportingInformation"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// Indicates the type of dispensing event that is performed. For example, Trial
	// Fill, Completion of Trial, Partial Fill, Emergency Fill, Samples, etc.
	Type *fhir.CodeableConcept `fhirpath:"type"`

	// The time the dispensed product was provided to the patient or their
	// representative.
	WhenHandedOver *fhir.DateTime `fhirpath:"whenHandedOver"`

	// The time when the dispensed product was packaged and reviewed.
	WhenPrepared *fhir.DateTime `fhirpath:"whenPrepared"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetAuthorizingPrescription returns the value of the field AuthorizingPrescription.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetAuthorizingPrescription() []*fhir.Reference {
	if md == nil {
		return nil
	}
	return md.AuthorizingPrescription
}

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetCategory() *fhir.CodeableConcept {
	if md == nil {
		return nil
	}
	return md.Category
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetContained() []fhir.Resource {
	if md == nil {
		return nil
	}
	return md.Contained
}

// GetContext returns the value of the field Context.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetContext() *fhir.Reference {
	if md == nil {
		return nil
	}
	return md.Context
}

// GetDaysSupply returns the value of the field DaysSupply.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetDaysSupply() *fhir.Quantity {
	if md == nil {
		return nil
	}
	return md.DaysSupply
}

// GetDestination returns the value of the field Destination.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetDestination() *fhir.Reference {
	if md == nil {
		return nil
	}
	return md.Destination
}

// GetDetectedIssue returns the value of the field DetectedIssue.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetDetectedIssue() []*fhir.Reference {
	if md == nil {
		return nil
	}
	return md.DetectedIssue
}

// GetDosageInstruction returns the value of the field DosageInstruction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetDosageInstruction() []*fhir.Dosage {
	if md == nil {
		return nil
	}
	return md.DosageInstruction
}

// GetEventHistory returns the value of the field EventHistory.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetEventHistory() []*fhir.Reference {
	if md == nil {
		return nil
	}
	return md.EventHistory
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetExtension() []*fhir.Extension {
	if md == nil {
		return nil
	}
	return md.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetID() string {
	if md == nil {
		return ""
	}
	return md.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetIdentifier() []*fhir.Identifier {
	if md == nil {
		return nil
	}
	return md.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetImplicitRules() *fhir.URI {
	if md == nil {
		return nil
	}
	return md.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetLanguage() *fhir.Code {
	if md == nil {
		return nil
	}
	return md.Language
}

// GetLocation returns the value of the field Location.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetLocation() *fhir.Reference {
	if md == nil {
		return nil
	}
	return md.Location
}

// GetMedication returns the value of the field Medication.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetMedication() fhir.Element {
	if md == nil {
		return nil
	}
	return md.Medication
}

// GetMedicationCodeableConcept returns the value of the field Medication.
func (md *MedicationDispense) GetMedicationCodeableConcept() *fhir.CodeableConcept {
	if md == nil {
		return nil
	}
	val, ok := md.Medication.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetMedicationReference returns the value of the field Medication.
func (md *MedicationDispense) GetMedicationReference() *fhir.Reference {
	if md == nil {
		return nil
	}
	val, ok := md.Medication.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
} // GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetMeta() *fhir.Meta {
	if md == nil {
		return nil
	}
	return md.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetModifierExtension() []*fhir.Extension {
	if md == nil {
		return nil
	}
	return md.ModifierExtension
}

// GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetNote() []*fhir.Annotation {
	if md == nil {
		return nil
	}
	return md.Note
}

// GetPartOf returns the value of the field PartOf.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetPartOf() []*fhir.Reference {
	if md == nil {
		return nil
	}
	return md.PartOf
}

// GetPerformer returns the value of the field Performer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetPerformer() []*MedicationDispensePerformer {
	if md == nil {
		return nil
	}
	return md.Performer
}

// GetQuantity returns the value of the field Quantity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetQuantity() *fhir.Quantity {
	if md == nil {
		return nil
	}
	return md.Quantity
}

// GetReceiver returns the value of the field Receiver.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetReceiver() []*fhir.Reference {
	if md == nil {
		return nil
	}
	return md.Receiver
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetStatus() *fhir.Code {
	if md == nil {
		return nil
	}
	return md.Status
}

// GetStatusReason returns the value of the field StatusReason.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetStatusReason() fhir.Element {
	if md == nil {
		return nil
	}
	return md.StatusReason
}

// GetStatusReasonCodeableConcept returns the value of the field StatusReason.
func (md *MedicationDispense) GetStatusReasonCodeableConcept() *fhir.CodeableConcept {
	if md == nil {
		return nil
	}
	val, ok := md.StatusReason.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetStatusReasonReference returns the value of the field StatusReason.
func (md *MedicationDispense) GetStatusReasonReference() *fhir.Reference {
	if md == nil {
		return nil
	}
	val, ok := md.StatusReason.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
} // GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetSubject() *fhir.Reference {
	if md == nil {
		return nil
	}
	return md.Subject
}

// GetSubstitution returns the value of the field Substitution.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetSubstitution() *MedicationDispenseSubstitution {
	if md == nil {
		return nil
	}
	return md.Substitution
}

// GetSupportingInformation returns the value of the field SupportingInformation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetSupportingInformation() []*fhir.Reference {
	if md == nil {
		return nil
	}
	return md.SupportingInformation
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetText() *fhir.Narrative {
	if md == nil {
		return nil
	}
	return md.Text
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetType() *fhir.CodeableConcept {
	if md == nil {
		return nil
	}
	return md.Type
}

// GetWhenHandedOver returns the value of the field WhenHandedOver.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetWhenHandedOver() *fhir.DateTime {
	if md == nil {
		return nil
	}
	return md.WhenHandedOver
}

// GetWhenPrepared returns the value of the field WhenPrepared.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (md *MedicationDispense) GetWhenPrepared() *fhir.DateTime {
	if md == nil {
		return nil
	}
	return md.WhenPrepared
}

// Who performed event// Indicates who or what performed the event.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MedicationDispense.json
type MedicationDispensePerformer struct {

	// The device, practitioner, etc. who performed the action. It should be
	// assumed that the actor is the dispenser of the medication.
	Actor *fhir.Reference `fhirpath:"actor"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Distinguishes the type of performer in the dispense. For example, date
	// enterer, packager, final checker.
	Function *fhir.CodeableConcept `fhirpath:"function"`

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

// GetActor returns the value of the field Actor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mdp *MedicationDispensePerformer) GetActor() *fhir.Reference {
	if mdp == nil {
		return nil
	}
	return mdp.Actor
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mdp *MedicationDispensePerformer) GetExtension() []*fhir.Extension {
	if mdp == nil {
		return nil
	}
	return mdp.Extension
}

// GetFunction returns the value of the field Function.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mdp *MedicationDispensePerformer) GetFunction() *fhir.CodeableConcept {
	if mdp == nil {
		return nil
	}
	return mdp.Function
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mdp *MedicationDispensePerformer) GetID() string {
	if mdp == nil {
		return ""
	}
	return mdp.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mdp *MedicationDispensePerformer) GetModifierExtension() []*fhir.Extension {
	if mdp == nil {
		return nil
	}
	return mdp.ModifierExtension
}

// Whether a substitution was performed on the dispense// Indicates whether or not substitution was made as part of the dispense. In
// some cases, substitution will be expected but does not happen, in other
// cases substitution is not expected but does happen. This block explains what
// substitution did or did not happen and why. If nothing is specified,
// substitution was not done.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MedicationDispense.json
type MedicationDispenseSubstitution struct {

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

	// Indicates the reason for the substitution (or lack of substitution) from
	// what was prescribed.
	Reason []*fhir.CodeableConcept `fhirpath:"reason"`

	// The person or organization that has primary responsibility for the
	// substitution.
	ResponsibleParty []*fhir.Reference `fhirpath:"responsibleParty"`

	// A code signifying whether a different drug was dispensed from what was
	// prescribed.
	Type *fhir.CodeableConcept `fhirpath:"type"`

	// True if the dispenser dispensed a different drug or product from what was
	// prescribed.
	WasSubstituted *fhir.Boolean `fhirpath:"wasSubstituted"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mds *MedicationDispenseSubstitution) GetExtension() []*fhir.Extension {
	if mds == nil {
		return nil
	}
	return mds.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mds *MedicationDispenseSubstitution) GetID() string {
	if mds == nil {
		return ""
	}
	return mds.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mds *MedicationDispenseSubstitution) GetModifierExtension() []*fhir.Extension {
	if mds == nil {
		return nil
	}
	return mds.ModifierExtension
}

// GetReason returns the value of the field Reason.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mds *MedicationDispenseSubstitution) GetReason() []*fhir.CodeableConcept {
	if mds == nil {
		return nil
	}
	return mds.Reason
}

// GetResponsibleParty returns the value of the field ResponsibleParty.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mds *MedicationDispenseSubstitution) GetResponsibleParty() []*fhir.Reference {
	if mds == nil {
		return nil
	}
	return mds.ResponsibleParty
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mds *MedicationDispenseSubstitution) GetType() *fhir.CodeableConcept {
	if mds == nil {
		return nil
	}
	return mds.Type
}

// GetWasSubstituted returns the value of the field WasSubstituted.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mds *MedicationDispenseSubstitution) GetWasSubstituted() *fhir.Boolean {
	if mds == nil {
		return nil
	}
	return mds.WasSubstituted
}

func (md *MedicationDispense) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (md *MedicationDispense) UnmarshalJSON(data []byte) error {
	var raw struct {
		AuthorizingPrescription []*fhir.Reference     `json:"authorizingPrescription"`
		Category                *fhir.CodeableConcept `json:"category"`
		Contained               []fhir.Resource       `json:"contained"`
		Context                 *fhir.Reference       `json:"context"`
		DaysSupply              *fhir.Quantity        `json:"daysSupply"`
		Destination             *fhir.Reference       `json:"destination"`
		DetectedIssue           []*fhir.Reference     `json:"detectedIssue"`
		DosageInstruction       []*fhir.Dosage        `json:"dosageInstruction"`
		EventHistory            []*fhir.Reference     `json:"eventHistory"`
		Extension               []*fhir.Extension     `json:"extension"`

		ID                          string                          `json:"id"`
		Identifier                  []*fhir.Identifier              `json:"identifier"`
		ImplicitRules               *fhir.URI                       `json:"implicitRules"`
		Language                    *fhir.Code                      `json:"language"`
		Location                    *fhir.Reference                 `json:"location"`
		MedicationCodeableConcept   *fhir.CodeableConcept           `json:"medicationCodeableConcept"`
		MedicationReference         *fhir.Reference                 `json:"medicationReference"`
		Meta                        *fhir.Meta                      `json:"meta"`
		ModifierExtension           []*fhir.Extension               `json:"modifierExtension"`
		Note                        []*fhir.Annotation              `json:"note"`
		PartOf                      []*fhir.Reference               `json:"partOf"`
		Performer                   []*MedicationDispensePerformer  `json:"performer"`
		Quantity                    *fhir.Quantity                  `json:"quantity"`
		Receiver                    []*fhir.Reference               `json:"receiver"`
		Status                      *fhir.Code                      `json:"status"`
		StatusReasonCodeableConcept *fhir.CodeableConcept           `json:"statusReasonCodeableConcept"`
		StatusReasonReference       *fhir.Reference                 `json:"statusReasonReference"`
		Subject                     *fhir.Reference                 `json:"subject"`
		Substitution                *MedicationDispenseSubstitution `json:"substitution"`
		SupportingInformation       []*fhir.Reference               `json:"supportingInformation"`
		Text                        *fhir.Narrative                 `json:"text"`
		Type                        *fhir.CodeableConcept           `json:"type"`
		WhenHandedOver              *fhir.DateTime                  `json:"whenHandedOver"`
		WhenPrepared                *fhir.DateTime                  `json:"whenPrepared"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	md.AuthorizingPrescription = raw.AuthorizingPrescription
	md.Category = raw.Category
	md.Contained = raw.Contained
	md.Context = raw.Context
	md.DaysSupply = raw.DaysSupply
	md.Destination = raw.Destination
	md.DetectedIssue = raw.DetectedIssue
	md.DosageInstruction = raw.DosageInstruction
	md.EventHistory = raw.EventHistory
	md.Extension = raw.Extension
	md.ID = raw.ID
	md.Identifier = raw.Identifier
	md.ImplicitRules = raw.ImplicitRules
	md.Language = raw.Language
	md.Location = raw.Location
	md.Medication, err = validate.SelectOneOf[fhir.Element]("MedicationDispense.medication",
		raw.MedicationCodeableConcept,
		raw.MedicationReference)
	if err != nil {
		return err
	}
	md.Meta = raw.Meta
	md.ModifierExtension = raw.ModifierExtension
	md.Note = raw.Note
	md.PartOf = raw.PartOf
	md.Performer = raw.Performer
	md.Quantity = raw.Quantity
	md.Receiver = raw.Receiver
	md.Status = raw.Status
	md.StatusReason, err = validate.SelectOneOf[fhir.Element]("MedicationDispense.statusReason",
		raw.StatusReasonCodeableConcept,
		raw.StatusReasonReference)
	if err != nil {
		return err
	}
	md.Subject = raw.Subject
	md.Substitution = raw.Substitution
	md.SupportingInformation = raw.SupportingInformation
	md.Text = raw.Text
	md.Type = raw.Type
	md.WhenHandedOver = raw.WhenHandedOver
	md.WhenPrepared = raw.WhenPrepared
	return nil
}

var _ json.Marshaler = (*MedicationDispense)(nil)
var _ json.Unmarshaler = (*MedicationDispense)(nil)

func (mdp *MedicationDispensePerformer) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (mdp *MedicationDispensePerformer) UnmarshalJSON(data []byte) error {
	var raw struct {
		Actor     *fhir.Reference       `json:"actor"`
		Extension []*fhir.Extension     `json:"extension"`
		Function  *fhir.CodeableConcept `json:"function"`

		ID                string            `json:"id"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mdp.Actor = raw.Actor
	mdp.Extension = raw.Extension
	mdp.Function = raw.Function
	mdp.ID = raw.ID
	mdp.ModifierExtension = raw.ModifierExtension
	return nil
}

var _ json.Marshaler = (*MedicationDispensePerformer)(nil)
var _ json.Unmarshaler = (*MedicationDispensePerformer)(nil)

func (mds *MedicationDispenseSubstitution) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (mds *MedicationDispenseSubstitution) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                string                  `json:"id"`
		ModifierExtension []*fhir.Extension       `json:"modifierExtension"`
		Reason            []*fhir.CodeableConcept `json:"reason"`
		ResponsibleParty  []*fhir.Reference       `json:"responsibleParty"`
		Type              *fhir.CodeableConcept   `json:"type"`
		WasSubstituted    *fhir.Boolean           `json:"wasSubstituted"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mds.Extension = raw.Extension
	mds.ID = raw.ID
	mds.ModifierExtension = raw.ModifierExtension
	mds.Reason = raw.Reason
	mds.ResponsibleParty = raw.ResponsibleParty
	mds.Type = raw.Type
	mds.WasSubstituted = raw.WasSubstituted
	return nil
}

var _ json.Marshaler = (*MedicationDispenseSubstitution)(nil)
var _ json.Unmarshaler = (*MedicationDispenseSubstitution)(nil)
