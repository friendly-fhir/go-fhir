// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package communication

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// An occurrence of information being transmitted; e.g. an alert that was sent
// to a responsible provider, a public health agency that was notified about a
// reportable condition.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Communication
//   - Source File: StructureDefinition-Communication.json
type Communication struct {

	// Other resources that pertain to this communication and to which this
	// communication should be associated.
	About []*fhir.Reference `fhirpath:"about"`

	// An order, proposal or plan fulfilled in whole or in part by this
	// Communication.
	BasedOn []*fhir.Reference `fhirpath:"basedOn"`

	// The type of message conveyed such as alert, notification, reminder,
	// instruction, etc.
	Category []*fhir.CodeableConcept `fhirpath:"category"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The Encounter during which this Communication was created or to which the
	// creation of this record is tightly associated.
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

	// Business identifiers assigned to this communication by the performer or
	// other systems which remain constant as the resource is updated and
	// propagates from server to server.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// Prior communication that this communication is in response to.
	InResponseTo []*fhir.Reference `fhirpath:"inResponseTo"`

	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other
	// definition that is adhered to in whole or in part by this Communication.
	InstantiatesCanonical []*fhir.Canonical `fhirpath:"instantiatesCanonical"`

	// The URL pointing to an externally maintained protocol, guideline, orderset
	// or other definition that is adhered to in whole or in part by this
	// Communication.
	InstantiatesURI []*fhir.URI `fhirpath:"instantiatesUri"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// A channel that was used for this communication (e.g. email, fax).
	Medium []*fhir.CodeableConcept `fhirpath:"medium"`

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

	// Additional notes or commentary about the communication by the sender,
	// receiver or other interested parties.
	Note []*fhir.Annotation `fhirpath:"note"`

	// Part of this action.
	PartOf []*fhir.Reference `fhirpath:"partOf"`

	// Text, attachment(s), or resource(s) that was communicated to the recipient.
	Payload []*CommunicationPayload `fhirpath:"payload"`

	// Characterizes how quickly the planned or in progress communication must be
	// addressed. Includes concepts such as stat, urgent, routine.
	Priority *fhir.Code `fhirpath:"priority"`

	// The reason or justification for the communication.
	ReasonCode []*fhir.CodeableConcept `fhirpath:"reasonCode"`

	// Indicates another resource whose existence justifies this communication.
	ReasonReference []*fhir.Reference `fhirpath:"reasonReference"`

	// The time when this communication arrived at the destination.
	Received *fhir.DateTime `fhirpath:"received"`

	// The entity (e.g. person, organization, clinical information system, care
	// team or device) which was the target of the communication. If receipts need
	// to be tracked by an individual, a separate resource instance will need to be
	// created for each recipient. Multiple recipient communications are intended
	// where either receipts are not tracked (e.g. a mass mail-out) or a receipt is
	// captured in aggregate (all emails confirmed received by a particular time).
	Recipient []*fhir.Reference `fhirpath:"recipient"`

	// The entity (e.g. person, organization, clinical information system, or
	// device) which was the source of the communication.
	Sender *fhir.Reference `fhirpath:"sender"`

	// The time when this communication was sent.
	Sent *fhir.DateTime `fhirpath:"sent"`

	// The status of the transmission.
	Status *fhir.Code `fhirpath:"status"`

	// Captures the reason for the current state of the Communication.
	StatusReason *fhir.CodeableConcept `fhirpath:"statusReason"`

	// The patient or group that was the focus of this communication.
	Subject *fhir.Reference `fhirpath:"subject"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// Description of the purpose/content, similar to a subject line in an email.
	Topic *fhir.CodeableConcept `fhirpath:"topic"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetAbout returns the value of the field About.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetAbout() []*fhir.Reference {
	if c == nil {
		return nil
	}
	return c.About
}

// GetBasedOn returns the value of the field BasedOn.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetBasedOn() []*fhir.Reference {
	if c == nil {
		return nil
	}
	return c.BasedOn
}

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetCategory() []*fhir.CodeableConcept {
	if c == nil {
		return nil
	}
	return c.Category
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetContained() []fhir.Resource {
	if c == nil {
		return nil
	}
	return c.Contained
}

// GetEncounter returns the value of the field Encounter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetEncounter() *fhir.Reference {
	if c == nil {
		return nil
	}
	return c.Encounter
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetExtension() []*fhir.Extension {
	if c == nil {
		return nil
	}
	return c.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetIdentifier() []*fhir.Identifier {
	if c == nil {
		return nil
	}
	return c.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetImplicitRules() *fhir.URI {
	if c == nil {
		return nil
	}
	return c.ImplicitRules
}

// GetInResponseTo returns the value of the field InResponseTo.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetInResponseTo() []*fhir.Reference {
	if c == nil {
		return nil
	}
	return c.InResponseTo
}

// GetInstantiatesCanonical returns the value of the field InstantiatesCanonical.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetInstantiatesCanonical() []*fhir.Canonical {
	if c == nil {
		return nil
	}
	return c.InstantiatesCanonical
}

// GetInstantiatesURI returns the value of the field InstantiatesURI.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetInstantiatesURI() []*fhir.URI {
	if c == nil {
		return nil
	}
	return c.InstantiatesURI
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetLanguage() *fhir.Code {
	if c == nil {
		return nil
	}
	return c.Language
}

// GetMedium returns the value of the field Medium.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetMedium() []*fhir.CodeableConcept {
	if c == nil {
		return nil
	}
	return c.Medium
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetMeta() *fhir.Meta {
	if c == nil {
		return nil
	}
	return c.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetModifierExtension() []*fhir.Extension {
	if c == nil {
		return nil
	}
	return c.ModifierExtension
}

// GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetNote() []*fhir.Annotation {
	if c == nil {
		return nil
	}
	return c.Note
}

// GetPartOf returns the value of the field PartOf.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetPartOf() []*fhir.Reference {
	if c == nil {
		return nil
	}
	return c.PartOf
}

// GetPayload returns the value of the field Payload.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetPayload() []*CommunicationPayload {
	if c == nil {
		return nil
	}
	return c.Payload
}

// GetPriority returns the value of the field Priority.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetPriority() *fhir.Code {
	if c == nil {
		return nil
	}
	return c.Priority
}

// GetReasonCode returns the value of the field ReasonCode.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetReasonCode() []*fhir.CodeableConcept {
	if c == nil {
		return nil
	}
	return c.ReasonCode
}

// GetReasonReference returns the value of the field ReasonReference.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetReasonReference() []*fhir.Reference {
	if c == nil {
		return nil
	}
	return c.ReasonReference
}

// GetReceived returns the value of the field Received.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetReceived() *fhir.DateTime {
	if c == nil {
		return nil
	}
	return c.Received
}

// GetRecipient returns the value of the field Recipient.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetRecipient() []*fhir.Reference {
	if c == nil {
		return nil
	}
	return c.Recipient
}

// GetSender returns the value of the field Sender.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetSender() *fhir.Reference {
	if c == nil {
		return nil
	}
	return c.Sender
}

// GetSent returns the value of the field Sent.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetSent() *fhir.DateTime {
	if c == nil {
		return nil
	}
	return c.Sent
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetStatus() *fhir.Code {
	if c == nil {
		return nil
	}
	return c.Status
}

// GetStatusReason returns the value of the field StatusReason.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetStatusReason() *fhir.CodeableConcept {
	if c == nil {
		return nil
	}
	return c.StatusReason
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetSubject() *fhir.Reference {
	if c == nil {
		return nil
	}
	return c.Subject
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetText() *fhir.Narrative {
	if c == nil {
		return nil
	}
	return c.Text
}

// GetTopic returns the value of the field Topic.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (c *Communication) GetTopic() *fhir.CodeableConcept {
	if c == nil {
		return nil
	}
	return c.Topic
}

// Message payload// Text, attachment(s), or resource(s) that was communicated to the recipient.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Communication.json
type CommunicationPayload struct {

	// A communicated content (or for multi-part communications, one portion of the
	// communication).
	Content fhir.Element `fhirpath:"content"`

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

// GetContent returns the value of the field Content.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cp *CommunicationPayload) GetContent() fhir.Element {
	if cp == nil {
		return nil
	}
	return cp.Content
}

// GetContentString returns the value of the field Content.
func (cp *CommunicationPayload) GetContentString() *fhir.String {
	if cp == nil {
		return nil
	}
	val, ok := cp.Content.(*fhir.String)
	if !ok {
		return nil
	}
	return val
}

// GetContentAttachment returns the value of the field Content.
func (cp *CommunicationPayload) GetContentAttachment() *fhir.Attachment {
	if cp == nil {
		return nil
	}
	val, ok := cp.Content.(*fhir.Attachment)
	if !ok {
		return nil
	}
	return val
}

// GetContentReference returns the value of the field Content.
func (cp *CommunicationPayload) GetContentReference() *fhir.Reference {
	if cp == nil {
		return nil
	}
	val, ok := cp.Content.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
} // GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cp *CommunicationPayload) GetExtension() []*fhir.Extension {
	if cp == nil {
		return nil
	}
	return cp.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cp *CommunicationPayload) GetID() string {
	if cp == nil {
		return ""
	}
	return cp.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (cp *CommunicationPayload) GetModifierExtension() []*fhir.Extension {
	if cp == nil {
		return nil
	}
	return cp.ModifierExtension
}

func (c *Communication) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (c *Communication) UnmarshalJSON(data []byte) error {
	var raw struct {
		About     []*fhir.Reference       `json:"about"`
		BasedOn   []*fhir.Reference       `json:"basedOn"`
		Category  []*fhir.CodeableConcept `json:"category"`
		Contained []fhir.Resource         `json:"contained"`
		Encounter *fhir.Reference         `json:"encounter"`
		Extension []*fhir.Extension       `json:"extension"`

		ID                    string                  `json:"id"`
		Identifier            []*fhir.Identifier      `json:"identifier"`
		ImplicitRules         *fhir.URI               `json:"implicitRules"`
		InResponseTo          []*fhir.Reference       `json:"inResponseTo"`
		InstantiatesCanonical []*fhir.Canonical       `json:"instantiatesCanonical"`
		InstantiatesURI       []*fhir.URI             `json:"instantiatesUri"`
		Language              *fhir.Code              `json:"language"`
		Medium                []*fhir.CodeableConcept `json:"medium"`
		Meta                  *fhir.Meta              `json:"meta"`
		ModifierExtension     []*fhir.Extension       `json:"modifierExtension"`
		Note                  []*fhir.Annotation      `json:"note"`
		PartOf                []*fhir.Reference       `json:"partOf"`
		Payload               []*CommunicationPayload `json:"payload"`
		Priority              *fhir.Code              `json:"priority"`
		ReasonCode            []*fhir.CodeableConcept `json:"reasonCode"`
		ReasonReference       []*fhir.Reference       `json:"reasonReference"`
		Received              *fhir.DateTime          `json:"received"`
		Recipient             []*fhir.Reference       `json:"recipient"`
		Sender                *fhir.Reference         `json:"sender"`
		Sent                  *fhir.DateTime          `json:"sent"`
		Status                *fhir.Code              `json:"status"`
		StatusReason          *fhir.CodeableConcept   `json:"statusReason"`
		Subject               *fhir.Reference         `json:"subject"`
		Text                  *fhir.Narrative         `json:"text"`
		Topic                 *fhir.CodeableConcept   `json:"topic"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c.About = raw.About
	c.BasedOn = raw.BasedOn
	c.Category = raw.Category
	c.Contained = raw.Contained
	c.Encounter = raw.Encounter
	c.Extension = raw.Extension
	c.ID = raw.ID
	c.Identifier = raw.Identifier
	c.ImplicitRules = raw.ImplicitRules
	c.InResponseTo = raw.InResponseTo
	c.InstantiatesCanonical = raw.InstantiatesCanonical
	c.InstantiatesURI = raw.InstantiatesURI
	c.Language = raw.Language
	c.Medium = raw.Medium
	c.Meta = raw.Meta
	c.ModifierExtension = raw.ModifierExtension
	c.Note = raw.Note
	c.PartOf = raw.PartOf
	c.Payload = raw.Payload
	c.Priority = raw.Priority
	c.ReasonCode = raw.ReasonCode
	c.ReasonReference = raw.ReasonReference
	c.Received = raw.Received
	c.Recipient = raw.Recipient
	c.Sender = raw.Sender
	c.Sent = raw.Sent
	c.Status = raw.Status
	c.StatusReason = raw.StatusReason
	c.Subject = raw.Subject
	c.Text = raw.Text
	c.Topic = raw.Topic
	return nil
}

var _ json.Marshaler = (*Communication)(nil)
var _ json.Unmarshaler = (*Communication)(nil)

func (cp *CommunicationPayload) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (cp *CommunicationPayload) UnmarshalJSON(data []byte) error {
	var raw struct {
		ContentString     *fhir.String      `json:"contentString"`
		ContentAttachment *fhir.Attachment  `json:"contentAttachment"`
		ContentReference  *fhir.Reference   `json:"contentReference"`
		Extension         []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	cp.Content, err = validate.SelectOneOf[fhir.Element]("Communication.payload.content",
		raw.ContentString,
		raw.ContentAttachment,
		raw.ContentReference)
	if err != nil {
		return err
	}
	cp.Extension = raw.Extension
	cp.ID = raw.ID
	cp.ModifierExtension = raw.ModifierExtension
	return nil
}

var _ json.Marshaler = (*CommunicationPayload)(nil)
var _ json.Unmarshaler = (*CommunicationPayload)(nil)
