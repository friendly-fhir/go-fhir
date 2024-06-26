// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package paymentreconciliation

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// This resource provides the details including amount of a payment and
// allocates the payment items being paid.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/PaymentReconciliation
//   - Source File: StructureDefinition-PaymentReconciliation.json
type PaymentReconciliation struct {

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The date when the resource was created.
	Created *fhir.DateTime `fhirpath:"created"`

	// Distribution of the payment amount for a previously acknowledged payable.
	Detail []*PaymentReconciliationDetail `fhirpath:"detail"`

	// A human readable description of the status of the request for the
	// reconciliation.
	Disposition *fhir.String `fhirpath:"disposition"`

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// A code for the form to be used for printing the content.
	FormCode *fhir.CodeableConcept `fhirpath:"formCode"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// A unique identifier assigned to this payment reconciliation.
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

	// The outcome of a request for a reconciliation.
	Outcome *fhir.Code `fhirpath:"outcome"`

	// Total payment amount as indicated on the financial instrument.
	PaymentAmount *fhir.Money `fhirpath:"paymentAmount"`

	// The date of payment as indicated on the financial instrument.
	PaymentDate *fhir.Date `fhirpath:"paymentDate"`

	// Issuer's unique identifier for the payment instrument.
	PaymentIdentifier *fhir.Identifier `fhirpath:"paymentIdentifier"`

	// The party who generated the payment.
	PaymentIssuer *fhir.Reference `fhirpath:"paymentIssuer"`

	// The period of time for which payments have been gathered into this bulk
	// payment for settlement.
	Period *fhir.Period `fhirpath:"period"`

	// A note that describes or explains the processing in a human readable form.
	ProcessNote []*PaymentReconciliationProcessNote `fhirpath:"processNote"`

	// Original request resource reference.
	Request *fhir.Reference `fhirpath:"request"`

	// The practitioner who is responsible for the services rendered to the
	// patient.
	Requestor *fhir.Reference `fhirpath:"requestor"`

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
func (pr *PaymentReconciliation) GetContained() []fhir.Resource {
	if pr == nil {
		return nil
	}
	return pr.Contained
}

// GetCreated returns the value of the field Created.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetCreated() *fhir.DateTime {
	if pr == nil {
		return nil
	}
	return pr.Created
}

// GetDetail returns the value of the field Detail.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetDetail() []*PaymentReconciliationDetail {
	if pr == nil {
		return nil
	}
	return pr.Detail
}

// GetDisposition returns the value of the field Disposition.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetDisposition() *fhir.String {
	if pr == nil {
		return nil
	}
	return pr.Disposition
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetExtension() []*fhir.Extension {
	if pr == nil {
		return nil
	}
	return pr.Extension
}

// GetFormCode returns the value of the field FormCode.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetFormCode() *fhir.CodeableConcept {
	if pr == nil {
		return nil
	}
	return pr.FormCode
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetID() string {
	if pr == nil {
		return ""
	}
	return pr.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetIdentifier() []*fhir.Identifier {
	if pr == nil {
		return nil
	}
	return pr.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetImplicitRules() *fhir.URI {
	if pr == nil {
		return nil
	}
	return pr.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetLanguage() *fhir.Code {
	if pr == nil {
		return nil
	}
	return pr.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetMeta() *fhir.Meta {
	if pr == nil {
		return nil
	}
	return pr.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetModifierExtension() []*fhir.Extension {
	if pr == nil {
		return nil
	}
	return pr.ModifierExtension
}

// GetOutcome returns the value of the field Outcome.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetOutcome() *fhir.Code {
	if pr == nil {
		return nil
	}
	return pr.Outcome
}

// GetPaymentAmount returns the value of the field PaymentAmount.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetPaymentAmount() *fhir.Money {
	if pr == nil {
		return nil
	}
	return pr.PaymentAmount
}

// GetPaymentDate returns the value of the field PaymentDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetPaymentDate() *fhir.Date {
	if pr == nil {
		return nil
	}
	return pr.PaymentDate
}

// GetPaymentIdentifier returns the value of the field PaymentIdentifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetPaymentIdentifier() *fhir.Identifier {
	if pr == nil {
		return nil
	}
	return pr.PaymentIdentifier
}

// GetPaymentIssuer returns the value of the field PaymentIssuer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetPaymentIssuer() *fhir.Reference {
	if pr == nil {
		return nil
	}
	return pr.PaymentIssuer
}

// GetPeriod returns the value of the field Period.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetPeriod() *fhir.Period {
	if pr == nil {
		return nil
	}
	return pr.Period
}

// GetProcessNote returns the value of the field ProcessNote.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetProcessNote() []*PaymentReconciliationProcessNote {
	if pr == nil {
		return nil
	}
	return pr.ProcessNote
}

// GetRequest returns the value of the field Request.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetRequest() *fhir.Reference {
	if pr == nil {
		return nil
	}
	return pr.Request
}

// GetRequestor returns the value of the field Requestor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetRequestor() *fhir.Reference {
	if pr == nil {
		return nil
	}
	return pr.Requestor
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetStatus() *fhir.Code {
	if pr == nil {
		return nil
	}
	return pr.Status
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pr *PaymentReconciliation) GetText() *fhir.Narrative {
	if pr == nil {
		return nil
	}
	return pr.Text
}

// Settlement particulars// Distribution of the payment amount for a previously acknowledged payable.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-PaymentReconciliation.json
type PaymentReconciliationDetail struct {

	// The monetary amount allocated from the total payment to the payable.
	Amount *fhir.Money `fhirpath:"amount"`

	// The date from the response resource containing a commitment to pay.
	Date *fhir.Date `fhirpath:"date"`

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

	// Unique identifier for the current payment item for the referenced payable.
	Identifier *fhir.Identifier `fhirpath:"identifier"`

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

	// The party which is receiving the payment.
	Payee *fhir.Reference `fhirpath:"payee"`

	// Unique identifier for the prior payment item for the referenced payable.
	Predecessor *fhir.Identifier `fhirpath:"predecessor"`

	// A resource, such as a Claim, the evaluation of which could lead to payment.
	Request *fhir.Reference `fhirpath:"request"`

	// A resource, such as a ClaimResponse, which contains a commitment to payment.
	Response *fhir.Reference `fhirpath:"response"`

	// A reference to the individual who is responsible for inquiries regarding the
	// response and its payment.
	Responsible *fhir.Reference `fhirpath:"responsible"`

	// The party which submitted the claim or financial transaction.
	Submitter *fhir.Reference `fhirpath:"submitter"`

	// Code to indicate the nature of the payment.
	Type *fhir.CodeableConcept `fhirpath:"type"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAmount returns the value of the field Amount.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetAmount() *fhir.Money {
	if prd == nil {
		return nil
	}
	return prd.Amount
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetDate() *fhir.Date {
	if prd == nil {
		return nil
	}
	return prd.Date
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetExtension() []*fhir.Extension {
	if prd == nil {
		return nil
	}
	return prd.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetID() string {
	if prd == nil {
		return ""
	}
	return prd.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetIdentifier() *fhir.Identifier {
	if prd == nil {
		return nil
	}
	return prd.Identifier
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetModifierExtension() []*fhir.Extension {
	if prd == nil {
		return nil
	}
	return prd.ModifierExtension
}

// GetPayee returns the value of the field Payee.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetPayee() *fhir.Reference {
	if prd == nil {
		return nil
	}
	return prd.Payee
}

// GetPredecessor returns the value of the field Predecessor.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetPredecessor() *fhir.Identifier {
	if prd == nil {
		return nil
	}
	return prd.Predecessor
}

// GetRequest returns the value of the field Request.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetRequest() *fhir.Reference {
	if prd == nil {
		return nil
	}
	return prd.Request
}

// GetResponse returns the value of the field Response.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetResponse() *fhir.Reference {
	if prd == nil {
		return nil
	}
	return prd.Response
}

// GetResponsible returns the value of the field Responsible.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetResponsible() *fhir.Reference {
	if prd == nil {
		return nil
	}
	return prd.Responsible
}

// GetSubmitter returns the value of the field Submitter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetSubmitter() *fhir.Reference {
	if prd == nil {
		return nil
	}
	return prd.Submitter
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prd *PaymentReconciliationDetail) GetType() *fhir.CodeableConcept {
	if prd == nil {
		return nil
	}
	return prd.Type
}

// Note concerning processing// A note that describes or explains the processing in a human readable form.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-PaymentReconciliation.json
type PaymentReconciliationProcessNote struct {

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

	// The explanation or description associated with the processing.
	Text *fhir.String `fhirpath:"text"`

	// The business purpose of the note text.
	Type *fhir.Code `fhirpath:"type"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prpn *PaymentReconciliationProcessNote) GetExtension() []*fhir.Extension {
	if prpn == nil {
		return nil
	}
	return prpn.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prpn *PaymentReconciliationProcessNote) GetID() string {
	if prpn == nil {
		return ""
	}
	return prpn.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prpn *PaymentReconciliationProcessNote) GetModifierExtension() []*fhir.Extension {
	if prpn == nil {
		return nil
	}
	return prpn.ModifierExtension
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prpn *PaymentReconciliationProcessNote) GetText() *fhir.String {
	if prpn == nil {
		return nil
	}
	return prpn.Text
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (prpn *PaymentReconciliationProcessNote) GetType() *fhir.Code {
	if prpn == nil {
		return nil
	}
	return prpn.Type
}

func (pr *PaymentReconciliation) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (pr *PaymentReconciliation) UnmarshalJSON(data []byte) error {
	var raw struct {
		Contained   []fhir.Resource                `json:"contained"`
		Created     *fhir.DateTime                 `json:"created"`
		Detail      []*PaymentReconciliationDetail `json:"detail"`
		Disposition *fhir.String                   `json:"disposition"`
		Extension   []*fhir.Extension              `json:"extension"`
		FormCode    *fhir.CodeableConcept          `json:"formCode"`

		ID                string                              `json:"id"`
		Identifier        []*fhir.Identifier                  `json:"identifier"`
		ImplicitRules     *fhir.URI                           `json:"implicitRules"`
		Language          *fhir.Code                          `json:"language"`
		Meta              *fhir.Meta                          `json:"meta"`
		ModifierExtension []*fhir.Extension                   `json:"modifierExtension"`
		Outcome           *fhir.Code                          `json:"outcome"`
		PaymentAmount     *fhir.Money                         `json:"paymentAmount"`
		PaymentDate       *fhir.Date                          `json:"paymentDate"`
		PaymentIdentifier *fhir.Identifier                    `json:"paymentIdentifier"`
		PaymentIssuer     *fhir.Reference                     `json:"paymentIssuer"`
		Period            *fhir.Period                        `json:"period"`
		ProcessNote       []*PaymentReconciliationProcessNote `json:"processNote"`
		Request           *fhir.Reference                     `json:"request"`
		Requestor         *fhir.Reference                     `json:"requestor"`
		Status            *fhir.Code                          `json:"status"`
		Text              *fhir.Narrative                     `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	pr.Contained = raw.Contained
	pr.Created = raw.Created
	pr.Detail = raw.Detail
	pr.Disposition = raw.Disposition
	pr.Extension = raw.Extension
	pr.FormCode = raw.FormCode
	pr.ID = raw.ID
	pr.Identifier = raw.Identifier
	pr.ImplicitRules = raw.ImplicitRules
	pr.Language = raw.Language
	pr.Meta = raw.Meta
	pr.ModifierExtension = raw.ModifierExtension
	pr.Outcome = raw.Outcome
	pr.PaymentAmount = raw.PaymentAmount
	pr.PaymentDate = raw.PaymentDate
	pr.PaymentIdentifier = raw.PaymentIdentifier
	pr.PaymentIssuer = raw.PaymentIssuer
	pr.Period = raw.Period
	pr.ProcessNote = raw.ProcessNote
	pr.Request = raw.Request
	pr.Requestor = raw.Requestor
	pr.Status = raw.Status
	pr.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*PaymentReconciliation)(nil)
var _ json.Unmarshaler = (*PaymentReconciliation)(nil)

func (prd *PaymentReconciliationDetail) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (prd *PaymentReconciliationDetail) UnmarshalJSON(data []byte) error {
	var raw struct {
		Amount    *fhir.Money       `json:"amount"`
		Date      *fhir.Date        `json:"date"`
		Extension []*fhir.Extension `json:"extension"`

		ID                string                `json:"id"`
		Identifier        *fhir.Identifier      `json:"identifier"`
		ModifierExtension []*fhir.Extension     `json:"modifierExtension"`
		Payee             *fhir.Reference       `json:"payee"`
		Predecessor       *fhir.Identifier      `json:"predecessor"`
		Request           *fhir.Reference       `json:"request"`
		Response          *fhir.Reference       `json:"response"`
		Responsible       *fhir.Reference       `json:"responsible"`
		Submitter         *fhir.Reference       `json:"submitter"`
		Type              *fhir.CodeableConcept `json:"type"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	prd.Amount = raw.Amount
	prd.Date = raw.Date
	prd.Extension = raw.Extension
	prd.ID = raw.ID
	prd.Identifier = raw.Identifier
	prd.ModifierExtension = raw.ModifierExtension
	prd.Payee = raw.Payee
	prd.Predecessor = raw.Predecessor
	prd.Request = raw.Request
	prd.Response = raw.Response
	prd.Responsible = raw.Responsible
	prd.Submitter = raw.Submitter
	prd.Type = raw.Type
	return nil
}

var _ json.Marshaler = (*PaymentReconciliationDetail)(nil)
var _ json.Unmarshaler = (*PaymentReconciliationDetail)(nil)

func (prpn *PaymentReconciliationProcessNote) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (prpn *PaymentReconciliationProcessNote) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		Text              *fhir.String      `json:"text"`
		Type              *fhir.Code        `json:"type"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	prpn.Extension = raw.Extension
	prpn.ID = raw.ID
	prpn.ModifierExtension = raw.ModifierExtension
	prpn.Text = raw.Text
	prpn.Type = raw.Type
	return nil
}

var _ json.Marshaler = (*PaymentReconciliationProcessNote)(nil)
var _ json.Unmarshaler = (*PaymentReconciliationProcessNote)(nil)
