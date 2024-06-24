// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package invoice

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Invoice containing collected ChargeItems from an Account with calculated
// individual and total price for Billing purpose.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Invoice
//   - Source File: StructureDefinition-Invoice.json
type Invoice struct {

	// Account which is supposed to be balanced with this Invoice.
	Account *fhir.Reference `fhirpath:"account"`

	// In case of Invoice cancellation a reason must be given (entered in error,
	// superseded by corrected invoice etc.).
	CancelledReason *fhir.String `fhirpath:"cancelledReason"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// Date/time(s) of when this Invoice was posted.
	Date *fhir.DateTime `fhirpath:"date"`

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

	// Identifier of this Invoice, often used for reference in correspondence about
	// this invoice or for tracking of payments.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The organizationissuing the Invoice.
	Issuer *fhir.Reference `fhirpath:"issuer"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// Each line item represents one charge for goods and services rendered.
	// Details such as date, code and amount are found in the referenced ChargeItem
	// resource.
	LineItem []*InvoiceLineItem `fhirpath:"lineItem"`

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

	// Comments made about the invoice by the issuer, subject, or other
	// participants.
	Note []*fhir.Annotation `fhirpath:"note"`

	// Indicates who or what performed or participated in the charged service.
	Participant []*InvoiceParticipant `fhirpath:"participant"`

	// Payment details such as banking details, period of payment, deductibles,
	// methods of payment.
	PaymentTerms *fhir.Markdown `fhirpath:"paymentTerms"`

	// The individual or Organization responsible for balancing of this invoice.
	Recipient *fhir.Reference `fhirpath:"recipient"`

	// The current state of the Invoice.
	Status *fhir.Code `fhirpath:"status"`

	// The individual or set of individuals receiving the goods and services billed
	// in this invoice.
	Subject *fhir.Reference `fhirpath:"subject"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// Invoice total, tax included.
	TotalGross *fhir.Money `fhirpath:"totalGross"`

	// Invoice total , taxes excluded.
	TotalNet *fhir.Money `fhirpath:"totalNet"`

	// Type of Invoice depending on domain, realm an usage (e.g. internal/external,
	// dental, preliminary).
	Type *fhir.CodeableConcept `fhirpath:"type"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetAccount returns the value of the field Account.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetAccount() *fhir.Reference {
	if i == nil {
		return nil
	}
	return i.Account
}

// GetCancelledReason returns the value of the field CancelledReason.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetCancelledReason() *fhir.String {
	if i == nil {
		return nil
	}
	return i.CancelledReason
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetContained() []fhir.Resource {
	if i == nil {
		return nil
	}
	return i.Contained
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetDate() *fhir.DateTime {
	if i == nil {
		return nil
	}
	return i.Date
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetExtension() []*fhir.Extension {
	if i == nil {
		return nil
	}
	return i.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetID() string {
	if i == nil {
		return ""
	}
	return i.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetIdentifier() []*fhir.Identifier {
	if i == nil {
		return nil
	}
	return i.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetImplicitRules() *fhir.URI {
	if i == nil {
		return nil
	}
	return i.ImplicitRules
}

// GetIssuer returns the value of the field Issuer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetIssuer() *fhir.Reference {
	if i == nil {
		return nil
	}
	return i.Issuer
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetLanguage() *fhir.Code {
	if i == nil {
		return nil
	}
	return i.Language
}

// GetLineItem returns the value of the field LineItem.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetLineItem() []*InvoiceLineItem {
	if i == nil {
		return nil
	}
	return i.LineItem
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetMeta() *fhir.Meta {
	if i == nil {
		return nil
	}
	return i.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetModifierExtension() []*fhir.Extension {
	if i == nil {
		return nil
	}
	return i.ModifierExtension
}

// GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetNote() []*fhir.Annotation {
	if i == nil {
		return nil
	}
	return i.Note
}

// GetParticipant returns the value of the field Participant.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetParticipant() []*InvoiceParticipant {
	if i == nil {
		return nil
	}
	return i.Participant
}

// GetPaymentTerms returns the value of the field PaymentTerms.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetPaymentTerms() *fhir.Markdown {
	if i == nil {
		return nil
	}
	return i.PaymentTerms
}

// GetRecipient returns the value of the field Recipient.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetRecipient() *fhir.Reference {
	if i == nil {
		return nil
	}
	return i.Recipient
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetStatus() *fhir.Code {
	if i == nil {
		return nil
	}
	return i.Status
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetSubject() *fhir.Reference {
	if i == nil {
		return nil
	}
	return i.Subject
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetText() *fhir.Narrative {
	if i == nil {
		return nil
	}
	return i.Text
}

// GetTotalGross returns the value of the field TotalGross.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetTotalGross() *fhir.Money {
	if i == nil {
		return nil
	}
	return i.TotalGross
}

// GetTotalNet returns the value of the field TotalNet.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetTotalNet() *fhir.Money {
	if i == nil {
		return nil
	}
	return i.TotalNet
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (i *Invoice) GetType() *fhir.CodeableConcept {
	if i == nil {
		return nil
	}
	return i.Type
}

// Line items of this Invoice// Each line item represents one charge for goods and services rendered.
// Details such as date, code and amount are found in the referenced ChargeItem
// resource.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Invoice.json
type InvoiceLineItem struct {

	// The ChargeItem contains information such as the billing code, date, amount
	// etc. If no further details are required for the lineItem, inline billing
	// codes can be added using the CodeableConcept data type instead of the
	// Reference.
	ChargeItem fhir.Element `fhirpath:"chargeItem"`

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

	// The price for a ChargeItem may be calculated as a base price with
	// surcharges/deductions that apply in certain conditions. A
	// ChargeItemDefinition resource that defines the prices, factors and
	// conditions that apply to a billing code is currently under development. The
	// priceComponent element can be used to offer transparency to the recipient of
	// the Invoice as to how the prices have been calculated.
	PriceComponent []*InvoiceLineItemPriceComponent `fhirpath:"priceComponent"`

	// Sequence in which the items appear on the invoice.
	Sequence *fhir.PositiveInt `fhirpath:"sequence"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetChargeItem returns the value of the field ChargeItem.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ili *InvoiceLineItem) GetChargeItem() fhir.Element {
	if ili == nil {
		return nil
	}
	return ili.ChargeItem
}

// GetChargeItemReference returns the value of the field ChargeItem.
func (ili *InvoiceLineItem) GetChargeItemReference() *fhir.Reference {
	if ili == nil {
		return nil
	}
	val, ok := ili.ChargeItem.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
}

// GetChargeItemCodeableConcept returns the value of the field ChargeItem.
func (ili *InvoiceLineItem) GetChargeItemCodeableConcept() *fhir.CodeableConcept {
	if ili == nil {
		return nil
	}
	val, ok := ili.ChargeItem.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
} // GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ili *InvoiceLineItem) GetExtension() []*fhir.Extension {
	if ili == nil {
		return nil
	}
	return ili.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ili *InvoiceLineItem) GetID() string {
	if ili == nil {
		return ""
	}
	return ili.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ili *InvoiceLineItem) GetModifierExtension() []*fhir.Extension {
	if ili == nil {
		return nil
	}
	return ili.ModifierExtension
}

// GetPriceComponent returns the value of the field PriceComponent.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ili *InvoiceLineItem) GetPriceComponent() []*InvoiceLineItemPriceComponent {
	if ili == nil {
		return nil
	}
	return ili.PriceComponent
}

// GetSequence returns the value of the field Sequence.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ili *InvoiceLineItem) GetSequence() *fhir.PositiveInt {
	if ili == nil {
		return nil
	}
	return ili.Sequence
}

// Components of total line item price// The price for a ChargeItem may be calculated as a base price with
// surcharges/deductions that apply in certain conditions. A
// ChargeItemDefinition resource that defines the prices, factors and
// conditions that apply to a billing code is currently under development. The
// priceComponent element can be used to offer transparency to the recipient of
// the Invoice as to how the prices have been calculated.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Invoice.json
type InvoiceLineItemPriceComponent struct {

	// The amount calculated for this component.
	Amount *fhir.Money `fhirpath:"amount"`

	// A code that identifies the component. Codes may be used to differentiate
	// between kinds of taxes, surcharges, discounts etc.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// The factor that has been applied on the base price for calculating this
	// component.
	Factor *fhir.Decimal `fhirpath:"factor"`

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

	// This code identifies the type of the component.
	Type *fhir.Code `fhirpath:"type"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAmount returns the value of the field Amount.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ilipc *InvoiceLineItemPriceComponent) GetAmount() *fhir.Money {
	if ilipc == nil {
		return nil
	}
	return ilipc.Amount
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ilipc *InvoiceLineItemPriceComponent) GetCode() *fhir.CodeableConcept {
	if ilipc == nil {
		return nil
	}
	return ilipc.Code
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ilipc *InvoiceLineItemPriceComponent) GetExtension() []*fhir.Extension {
	if ilipc == nil {
		return nil
	}
	return ilipc.Extension
}

// GetFactor returns the value of the field Factor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ilipc *InvoiceLineItemPriceComponent) GetFactor() *fhir.Decimal {
	if ilipc == nil {
		return nil
	}
	return ilipc.Factor
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ilipc *InvoiceLineItemPriceComponent) GetID() string {
	if ilipc == nil {
		return ""
	}
	return ilipc.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ilipc *InvoiceLineItemPriceComponent) GetModifierExtension() []*fhir.Extension {
	if ilipc == nil {
		return nil
	}
	return ilipc.ModifierExtension
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ilipc *InvoiceLineItemPriceComponent) GetType() *fhir.Code {
	if ilipc == nil {
		return nil
	}
	return ilipc.Type
}

// Participant in creation of this Invoice// Indicates who or what performed or participated in the charged service.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Invoice.json
type InvoiceParticipant struct {

	// The device, practitioner, etc. who performed or participated in the service.
	Actor *fhir.Reference `fhirpath:"actor"`

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

	// Describes the type of involvement (e.g. transcriptionist, creator etc.). If
	// the invoice has been created automatically, the Participant may be a billing
	// engine or another kind of device.
	Role *fhir.CodeableConcept `fhirpath:"role"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetActor returns the value of the field Actor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ip *InvoiceParticipant) GetActor() *fhir.Reference {
	if ip == nil {
		return nil
	}
	return ip.Actor
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ip *InvoiceParticipant) GetExtension() []*fhir.Extension {
	if ip == nil {
		return nil
	}
	return ip.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ip *InvoiceParticipant) GetID() string {
	if ip == nil {
		return ""
	}
	return ip.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ip *InvoiceParticipant) GetModifierExtension() []*fhir.Extension {
	if ip == nil {
		return nil
	}
	return ip.ModifierExtension
}

// GetRole returns the value of the field Role.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ip *InvoiceParticipant) GetRole() *fhir.CodeableConcept {
	if ip == nil {
		return nil
	}
	return ip.Role
}
