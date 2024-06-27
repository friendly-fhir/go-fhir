// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package chargeitem

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// The resource ChargeItem describes the provision of healthcare provider
// products for a certain patient, therefore referring not only to the product,
// but containing in addition details of the provision, like date, time,
// amounts and participating organizations and persons. Main Usage of the
// ChargeItem is to enable the billing process and internal cost allocation.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/ChargeItem
//   - Source File: StructureDefinition-ChargeItem.json
type ChargeItem struct {

	// Account into which this ChargeItems belongs.
	Account []*fhir.Reference `fhirpath:"account"`

	// The anatomical location where the related service has been applied.
	Bodysite []*fhir.CodeableConcept `fhirpath:"bodysite"`

	// A code that identifies the charge, like a billing code.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The encounter or episode of care that establishes the context for this
	// event.
	Context *fhir.Reference `fhirpath:"context"`

	// The financial cost center permits the tracking of charge attribution.
	CostCenter *fhir.Reference `fhirpath:"costCenter"`

	// References the source of pricing information, rules of application for the
	// code this ChargeItem uses.
	DefinitionCanonical []*fhir.Canonical `fhirpath:"definitionCanonical"`

	// References the (external) source of pricing information, rules of
	// application for the code this ChargeItem uses.
	DefinitionURI []*fhir.URI `fhirpath:"definitionUri"`

	// Date the charge item was entered.
	EnteredDate *fhir.DateTime `fhirpath:"enteredDate"`

	// The device, practitioner, etc. who entered the charge item.
	Enterer *fhir.Reference `fhirpath:"enterer"`

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Factor overriding the factor determined by the rules associated with the
	// code.
	FactorOverride *fhir.Decimal `fhirpath:"factorOverride"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// Identifiers assigned to this event performer or other systems.
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

	// Comments made about the event by the performer, subject or other
	// participants.
	Note []*fhir.Annotation `fhirpath:"note"`

	// Date/time(s) or duration when the charged service was applied.
	Occurrence fhir.Element `fhirpath:"occurrence"`

	// If the list price or the rule-based factor associated with the code is
	// overridden, this attribute can capture a text to indicate the reason for
	// this action.
	OverrideReason *fhir.String `fhirpath:"overrideReason"`

	// ChargeItems can be grouped to larger ChargeItems covering the whole set.
	PartOf []*fhir.Reference `fhirpath:"partOf"`

	// Indicates who or what performed or participated in the charged service.
	Performer []*ChargeItemPerformer `fhirpath:"performer"`

	// The organization requesting the service.
	PerformingOrganization *fhir.Reference `fhirpath:"performingOrganization"`

	// Total price of the charge overriding the list price associated with the
	// code.
	PriceOverride *fhir.Money `fhirpath:"priceOverride"`

	// Identifies the device, food, drug or other product being charged either by
	// type code or reference to an instance.
	Product fhir.Element `fhirpath:"product"`

	// Quantity of which the charge item has been serviced.
	Quantity *fhir.Quantity `fhirpath:"quantity"`

	// Describes why the event occurred in coded or textual form.
	Reason []*fhir.CodeableConcept `fhirpath:"reason"`

	// The organization performing the service.
	RequestingOrganization *fhir.Reference `fhirpath:"requestingOrganization"`

	// Indicated the rendered service that caused this charge.
	Service []*fhir.Reference `fhirpath:"service"`

	// The current state of the ChargeItem.
	Status *fhir.Code `fhirpath:"status"`

	// The individual or set of individuals the action is being or was performed
	// on.
	Subject *fhir.Reference `fhirpath:"subject"`

	// Further information supporting this charge.
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

// GetAccount returns the value of the field Account.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetAccount() []*fhir.Reference {
	if ci == nil {
		return nil
	}
	return ci.Account
}

// GetBodysite returns the value of the field Bodysite.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetBodysite() []*fhir.CodeableConcept {
	if ci == nil {
		return nil
	}
	return ci.Bodysite
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetCode() *fhir.CodeableConcept {
	if ci == nil {
		return nil
	}
	return ci.Code
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetContained() []fhir.Resource {
	if ci == nil {
		return nil
	}
	return ci.Contained
}

// GetContext returns the value of the field Context.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetContext() *fhir.Reference {
	if ci == nil {
		return nil
	}
	return ci.Context
}

// GetCostCenter returns the value of the field CostCenter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetCostCenter() *fhir.Reference {
	if ci == nil {
		return nil
	}
	return ci.CostCenter
}

// GetDefinitionCanonical returns the value of the field DefinitionCanonical.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetDefinitionCanonical() []*fhir.Canonical {
	if ci == nil {
		return nil
	}
	return ci.DefinitionCanonical
}

// GetDefinitionURI returns the value of the field DefinitionURI.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetDefinitionURI() []*fhir.URI {
	if ci == nil {
		return nil
	}
	return ci.DefinitionURI
}

// GetEnteredDate returns the value of the field EnteredDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetEnteredDate() *fhir.DateTime {
	if ci == nil {
		return nil
	}
	return ci.EnteredDate
}

// GetEnterer returns the value of the field Enterer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetEnterer() *fhir.Reference {
	if ci == nil {
		return nil
	}
	return ci.Enterer
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetExtension() []*fhir.Extension {
	if ci == nil {
		return nil
	}
	return ci.Extension
}

// GetFactorOverride returns the value of the field FactorOverride.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetFactorOverride() *fhir.Decimal {
	if ci == nil {
		return nil
	}
	return ci.FactorOverride
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetID() string {
	if ci == nil {
		return ""
	}
	return ci.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetIdentifier() []*fhir.Identifier {
	if ci == nil {
		return nil
	}
	return ci.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetImplicitRules() *fhir.URI {
	if ci == nil {
		return nil
	}
	return ci.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetLanguage() *fhir.Code {
	if ci == nil {
		return nil
	}
	return ci.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetMeta() *fhir.Meta {
	if ci == nil {
		return nil
	}
	return ci.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetModifierExtension() []*fhir.Extension {
	if ci == nil {
		return nil
	}
	return ci.ModifierExtension
}

// GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetNote() []*fhir.Annotation {
	if ci == nil {
		return nil
	}
	return ci.Note
}

// GetOccurrence returns the value of the field Occurrence.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetOccurrence() fhir.Element {
	if ci == nil {
		return nil
	}
	return ci.Occurrence
}

// GetOccurrenceDateTime returns the value of the field Occurrence.
func (ci *ChargeItem) GetOccurrenceDateTime() *fhir.DateTime {
	if ci == nil {
		return nil
	}
	val, ok := ci.Occurrence.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
}

// GetOccurrencePeriod returns the value of the field Occurrence.
func (ci *ChargeItem) GetOccurrencePeriod() *fhir.Period {
	if ci == nil {
		return nil
	}
	val, ok := ci.Occurrence.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
}

// GetOccurrenceTiming returns the value of the field Occurrence.
func (ci *ChargeItem) GetOccurrenceTiming() *fhir.Timing {
	if ci == nil {
		return nil
	}
	val, ok := ci.Occurrence.(*fhir.Timing)
	if !ok {
		return nil
	}
	return val
} // GetOverrideReason returns the value of the field OverrideReason.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetOverrideReason() *fhir.String {
	if ci == nil {
		return nil
	}
	return ci.OverrideReason
}

// GetPartOf returns the value of the field PartOf.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetPartOf() []*fhir.Reference {
	if ci == nil {
		return nil
	}
	return ci.PartOf
}

// GetPerformer returns the value of the field Performer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetPerformer() []*ChargeItemPerformer {
	if ci == nil {
		return nil
	}
	return ci.Performer
}

// GetPerformingOrganization returns the value of the field PerformingOrganization.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetPerformingOrganization() *fhir.Reference {
	if ci == nil {
		return nil
	}
	return ci.PerformingOrganization
}

// GetPriceOverride returns the value of the field PriceOverride.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetPriceOverride() *fhir.Money {
	if ci == nil {
		return nil
	}
	return ci.PriceOverride
}

// GetProduct returns the value of the field Product.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetProduct() fhir.Element {
	if ci == nil {
		return nil
	}
	return ci.Product
}

// GetProductReference returns the value of the field Product.
func (ci *ChargeItem) GetProductReference() *fhir.Reference {
	if ci == nil {
		return nil
	}
	val, ok := ci.Product.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
}

// GetProductCodeableConcept returns the value of the field Product.
func (ci *ChargeItem) GetProductCodeableConcept() *fhir.CodeableConcept {
	if ci == nil {
		return nil
	}
	val, ok := ci.Product.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
} // GetQuantity returns the value of the field Quantity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetQuantity() *fhir.Quantity {
	if ci == nil {
		return nil
	}
	return ci.Quantity
}

// GetReason returns the value of the field Reason.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetReason() []*fhir.CodeableConcept {
	if ci == nil {
		return nil
	}
	return ci.Reason
}

// GetRequestingOrganization returns the value of the field RequestingOrganization.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetRequestingOrganization() *fhir.Reference {
	if ci == nil {
		return nil
	}
	return ci.RequestingOrganization
}

// GetService returns the value of the field Service.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetService() []*fhir.Reference {
	if ci == nil {
		return nil
	}
	return ci.Service
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetStatus() *fhir.Code {
	if ci == nil {
		return nil
	}
	return ci.Status
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetSubject() *fhir.Reference {
	if ci == nil {
		return nil
	}
	return ci.Subject
}

// GetSupportingInformation returns the value of the field SupportingInformation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetSupportingInformation() []*fhir.Reference {
	if ci == nil {
		return nil
	}
	return ci.SupportingInformation
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ci *ChargeItem) GetText() *fhir.Narrative {
	if ci == nil {
		return nil
	}
	return ci.Text
}

// Who performed charged service// Indicates who or what performed or participated in the charged service.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-ChargeItem.json
type ChargeItemPerformer struct {

	// The device, practitioner, etc. who performed or participated in the service.
	Actor *fhir.Reference `fhirpath:"actor"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Describes the type of performance or participation(e.g. primary surgeon,
	// anesthesiologiest, etc.).
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
func (cip *ChargeItemPerformer) GetActor() *fhir.Reference {
	if cip == nil {
		return nil
	}
	return cip.Actor
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cip *ChargeItemPerformer) GetExtension() []*fhir.Extension {
	if cip == nil {
		return nil
	}
	return cip.Extension
}

// GetFunction returns the value of the field Function.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cip *ChargeItemPerformer) GetFunction() *fhir.CodeableConcept {
	if cip == nil {
		return nil
	}
	return cip.Function
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cip *ChargeItemPerformer) GetID() string {
	if cip == nil {
		return ""
	}
	return cip.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cip *ChargeItemPerformer) GetModifierExtension() []*fhir.Extension {
	if cip == nil {
		return nil
	}
	return cip.ModifierExtension
}

func (ci *ChargeItem) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (ci *ChargeItem) UnmarshalJSON(data []byte) error {
	var raw struct {
		Account             []*fhir.Reference       `json:"account"`
		Bodysite            []*fhir.CodeableConcept `json:"bodysite"`
		Code                *fhir.CodeableConcept   `json:"code"`
		Contained           []fhir.Resource         `json:"contained"`
		Context             *fhir.Reference         `json:"context"`
		CostCenter          *fhir.Reference         `json:"costCenter"`
		DefinitionCanonical []*fhir.Canonical       `json:"definitionCanonical"`
		DefinitionURI       []*fhir.URI             `json:"definitionUri"`
		EnteredDate         *fhir.DateTime          `json:"enteredDate"`
		Enterer             *fhir.Reference         `json:"enterer"`
		Extension           []*fhir.Extension       `json:"extension"`
		FactorOverride      *fhir.Decimal           `json:"factorOverride"`

		ID                     string                  `json:"id"`
		Identifier             []*fhir.Identifier      `json:"identifier"`
		ImplicitRules          *fhir.URI               `json:"implicitRules"`
		Language               *fhir.Code              `json:"language"`
		Meta                   *fhir.Meta              `json:"meta"`
		ModifierExtension      []*fhir.Extension       `json:"modifierExtension"`
		Note                   []*fhir.Annotation      `json:"note"`
		OccurrenceDateTime     *fhir.DateTime          `json:"occurrenceDateTime"`
		OccurrencePeriod       *fhir.Period            `json:"occurrencePeriod"`
		OccurrenceTiming       *fhir.Timing            `json:"occurrenceTiming"`
		OverrideReason         *fhir.String            `json:"overrideReason"`
		PartOf                 []*fhir.Reference       `json:"partOf"`
		Performer              []*ChargeItemPerformer  `json:"performer"`
		PerformingOrganization *fhir.Reference         `json:"performingOrganization"`
		PriceOverride          *fhir.Money             `json:"priceOverride"`
		ProductReference       *fhir.Reference         `json:"productReference"`
		ProductCodeableConcept *fhir.CodeableConcept   `json:"productCodeableConcept"`
		Quantity               *fhir.Quantity          `json:"quantity"`
		Reason                 []*fhir.CodeableConcept `json:"reason"`
		RequestingOrganization *fhir.Reference         `json:"requestingOrganization"`
		Service                []*fhir.Reference       `json:"service"`
		Status                 *fhir.Code              `json:"status"`
		Subject                *fhir.Reference         `json:"subject"`
		SupportingInformation  []*fhir.Reference       `json:"supportingInformation"`
		Text                   *fhir.Narrative         `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	ci.Account = raw.Account
	ci.Bodysite = raw.Bodysite
	ci.Code = raw.Code
	ci.Contained = raw.Contained
	ci.Context = raw.Context
	ci.CostCenter = raw.CostCenter
	ci.DefinitionCanonical = raw.DefinitionCanonical
	ci.DefinitionURI = raw.DefinitionURI
	ci.EnteredDate = raw.EnteredDate
	ci.Enterer = raw.Enterer
	ci.Extension = raw.Extension
	ci.FactorOverride = raw.FactorOverride
	ci.ID = raw.ID
	ci.Identifier = raw.Identifier
	ci.ImplicitRules = raw.ImplicitRules
	ci.Language = raw.Language
	ci.Meta = raw.Meta
	ci.ModifierExtension = raw.ModifierExtension
	ci.Note = raw.Note
	ci.Occurrence, err = validate.SelectOneOf[fhir.Element]("ChargeItem.occurrence",
		raw.OccurrenceDateTime,
		raw.OccurrencePeriod,
		raw.OccurrenceTiming)
	if err != nil {
		return err
	}
	ci.OverrideReason = raw.OverrideReason
	ci.PartOf = raw.PartOf
	ci.Performer = raw.Performer
	ci.PerformingOrganization = raw.PerformingOrganization
	ci.PriceOverride = raw.PriceOverride
	ci.Product, err = validate.SelectOneOf[fhir.Element]("ChargeItem.product",
		raw.ProductReference,
		raw.ProductCodeableConcept)
	if err != nil {
		return err
	}
	ci.Quantity = raw.Quantity
	ci.Reason = raw.Reason
	ci.RequestingOrganization = raw.RequestingOrganization
	ci.Service = raw.Service
	ci.Status = raw.Status
	ci.Subject = raw.Subject
	ci.SupportingInformation = raw.SupportingInformation
	ci.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*ChargeItem)(nil)
var _ json.Unmarshaler = (*ChargeItem)(nil)

func (cip *ChargeItemPerformer) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (cip *ChargeItemPerformer) UnmarshalJSON(data []byte) error {
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

	cip.Actor = raw.Actor
	cip.Extension = raw.Extension
	cip.Function = raw.Function
	cip.ID = raw.ID
	cip.ModifierExtension = raw.ModifierExtension
	return nil
}

var _ json.Marshaler = (*ChargeItemPerformer)(nil)
var _ json.Unmarshaler = (*ChargeItemPerformer)(nil)
