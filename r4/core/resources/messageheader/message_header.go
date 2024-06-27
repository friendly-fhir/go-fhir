// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package messageheader

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// The header for a message exchange that is either requesting or responding to
// an action. The reference(s) that are the subject of the action as well as
// other information related to the action are typically transmitted in a
// bundle in which the MessageHeader resource instance is the first resource in
// the bundle.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/MessageHeader
//   - Source File: StructureDefinition-MessageHeader.json
type MessageHeader struct {

	// The logical author of the message - the person or device that decided the
	// described event should happen. When there is more than one candidate, pick
	// the most proximal to the MessageHeader. Can provide other authors in
	// extensions.
	Author *fhir.Reference `fhirpath:"author"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// Permanent link to the MessageDefinition for this message.
	Definition *fhir.Canonical `fhirpath:"definition"`

	// The destination application which the message is intended for.
	Destination []*MessageHeaderDestination `fhirpath:"destination"`

	// The person or device that performed the data entry leading to this message.
	// When there is more than one candidate, pick the most proximal to the
	// message. Can provide other enterers in extensions.
	Enterer *fhir.Reference `fhirpath:"enterer"`

	// Code that identifies the event this message represents and connects it with
	// its definition. Events defined as part of the FHIR specification have the
	// system value "http://terminology.hl7.org/CodeSystem/message-events".
	// Alternatively uri to the EventDefinition.
	Event fhir.Element `fhirpath:"event"`

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// The actual data of the message - a reference to the root/focus class of the
	// event.
	Focus []*fhir.Reference `fhirpath:"focus"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

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

	// Coded indication of the cause for the event - indicates a reason for the
	// occurrence of the event that is a focus of this message.
	Reason *fhir.CodeableConcept `fhirpath:"reason"`

	// Information about the message that this message is a response to. Only
	// present if this message is a response.
	Response *MessageHeaderResponse `fhirpath:"response"`

	// The person or organization that accepts overall responsibility for the
	// contents of the message. The implication is that the message event happened
	// under the policies of the responsible party.
	Responsible *fhir.Reference `fhirpath:"responsible"`

	// Identifies the sending system to allow the use of a trust relationship.
	Sender *fhir.Reference `fhirpath:"sender"`

	// The source application from which this message originated.
	Source *MessageHeaderSource `fhirpath:"source"`

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

// GetAuthor returns the value of the field Author.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetAuthor() *fhir.Reference {
	if mh == nil {
		return nil
	}
	return mh.Author
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetContained() []fhir.Resource {
	if mh == nil {
		return nil
	}
	return mh.Contained
}

// GetDefinition returns the value of the field Definition.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetDefinition() *fhir.Canonical {
	if mh == nil {
		return nil
	}
	return mh.Definition
}

// GetDestination returns the value of the field Destination.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetDestination() []*MessageHeaderDestination {
	if mh == nil {
		return nil
	}
	return mh.Destination
}

// GetEnterer returns the value of the field Enterer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetEnterer() *fhir.Reference {
	if mh == nil {
		return nil
	}
	return mh.Enterer
}

// GetEvent returns the value of the field Event.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetEvent() fhir.Element {
	if mh == nil {
		return nil
	}
	return mh.Event
}

// GetEventCoding returns the value of the field Event.
func (mh *MessageHeader) GetEventCoding() *fhir.Coding {
	if mh == nil {
		return nil
	}
	val, ok := mh.Event.(*fhir.Coding)
	if !ok {
		return nil
	}
	return val
}

// GetEventURI returns the value of the field Event.
func (mh *MessageHeader) GetEventURI() *fhir.URI {
	if mh == nil {
		return nil
	}
	val, ok := mh.Event.(*fhir.URI)
	if !ok {
		return nil
	}
	return val
} // GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetExtension() []*fhir.Extension {
	if mh == nil {
		return nil
	}
	return mh.Extension
}

// GetFocus returns the value of the field Focus.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetFocus() []*fhir.Reference {
	if mh == nil {
		return nil
	}
	return mh.Focus
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetID() string {
	if mh == nil {
		return ""
	}
	return mh.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetImplicitRules() *fhir.URI {
	if mh == nil {
		return nil
	}
	return mh.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetLanguage() *fhir.Code {
	if mh == nil {
		return nil
	}
	return mh.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetMeta() *fhir.Meta {
	if mh == nil {
		return nil
	}
	return mh.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetModifierExtension() []*fhir.Extension {
	if mh == nil {
		return nil
	}
	return mh.ModifierExtension
}

// GetReason returns the value of the field Reason.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetReason() *fhir.CodeableConcept {
	if mh == nil {
		return nil
	}
	return mh.Reason
}

// GetResponse returns the value of the field Response.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetResponse() *MessageHeaderResponse {
	if mh == nil {
		return nil
	}
	return mh.Response
}

// GetResponsible returns the value of the field Responsible.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetResponsible() *fhir.Reference {
	if mh == nil {
		return nil
	}
	return mh.Responsible
}

// GetSender returns the value of the field Sender.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetSender() *fhir.Reference {
	if mh == nil {
		return nil
	}
	return mh.Sender
}

// GetSource returns the value of the field Source.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetSource() *MessageHeaderSource {
	if mh == nil {
		return nil
	}
	return mh.Source
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mh *MessageHeader) GetText() *fhir.Narrative {
	if mh == nil {
		return nil
	}
	return mh.Text
}

// Message destination application(s)// The destination application which the message is intended for.// There SHOULD be at least one destination, but in some circumstances, the
// source system is unaware of any particular destination system.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MessageHeader.json
type MessageHeaderDestination struct {

	// Indicates where the message should be routed to.
	Endpoint *fhir.URL `fhirpath:"endpoint"`

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

	// Human-readable name for the target system.
	Name *fhir.String `fhirpath:"name"`

	// Allows data conveyed by a message to be addressed to a particular person or
	// department when routing to a specific application isn't sufficient.
	Receiver *fhir.Reference `fhirpath:"receiver"`

	// Identifies the target end system in situations where the initial message
	// transmission is to an intermediary system.
	Target *fhir.Reference `fhirpath:"target"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetEndpoint returns the value of the field Endpoint.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhd *MessageHeaderDestination) GetEndpoint() *fhir.URL {
	if mhd == nil {
		return nil
	}
	return mhd.Endpoint
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhd *MessageHeaderDestination) GetExtension() []*fhir.Extension {
	if mhd == nil {
		return nil
	}
	return mhd.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhd *MessageHeaderDestination) GetID() string {
	if mhd == nil {
		return ""
	}
	return mhd.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhd *MessageHeaderDestination) GetModifierExtension() []*fhir.Extension {
	if mhd == nil {
		return nil
	}
	return mhd.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhd *MessageHeaderDestination) GetName() *fhir.String {
	if mhd == nil {
		return nil
	}
	return mhd.Name
}

// GetReceiver returns the value of the field Receiver.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhd *MessageHeaderDestination) GetReceiver() *fhir.Reference {
	if mhd == nil {
		return nil
	}
	return mhd.Receiver
}

// GetTarget returns the value of the field Target.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhd *MessageHeaderDestination) GetTarget() *fhir.Reference {
	if mhd == nil {
		return nil
	}
	return mhd.Target
}

// If this is a reply to prior message// Information about the message that this message is a response to. Only
// present if this message is a response.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MessageHeader.json
type MessageHeaderResponse struct {

	// Code that identifies the type of response to the message - whether it was
	// successful or not, and whether it should be resent or not.
	Code *fhir.Code `fhirpath:"code"`

	// Full details of any issues found in the message.
	Details *fhir.Reference `fhirpath:"details"`

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

	// The MessageHeader.id of the message to which this message is a response.
	Identifier *fhir.ID `fhirpath:"identifier"`

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
func (mhr *MessageHeaderResponse) GetCode() *fhir.Code {
	if mhr == nil {
		return nil
	}
	return mhr.Code
}

// GetDetails returns the value of the field Details.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhr *MessageHeaderResponse) GetDetails() *fhir.Reference {
	if mhr == nil {
		return nil
	}
	return mhr.Details
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhr *MessageHeaderResponse) GetExtension() []*fhir.Extension {
	if mhr == nil {
		return nil
	}
	return mhr.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhr *MessageHeaderResponse) GetID() string {
	if mhr == nil {
		return ""
	}
	return mhr.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhr *MessageHeaderResponse) GetIdentifier() *fhir.ID {
	if mhr == nil {
		return nil
	}
	return mhr.Identifier
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhr *MessageHeaderResponse) GetModifierExtension() []*fhir.Extension {
	if mhr == nil {
		return nil
	}
	return mhr.ModifierExtension
}

// Message source application// The source application from which this message originated.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MessageHeader.json
type MessageHeaderSource struct {

	// An e-mail, phone, website or other contact point to use to resolve issues
	// with message communications.
	Contact *fhir.ContactPoint `fhirpath:"contact"`

	// Identifies the routing target to send acknowledgements to.
	Endpoint *fhir.URL `fhirpath:"endpoint"`

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

	// Human-readable name for the source system.
	Name *fhir.String `fhirpath:"name"`

	// May include configuration or other information useful in debugging.
	Software *fhir.String `fhirpath:"software"`

	// Can convey versions of multiple systems in situations where a message passes
	// through multiple hands.
	Version *fhir.String `fhirpath:"version"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetContact returns the value of the field Contact.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhs *MessageHeaderSource) GetContact() *fhir.ContactPoint {
	if mhs == nil {
		return nil
	}
	return mhs.Contact
}

// GetEndpoint returns the value of the field Endpoint.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhs *MessageHeaderSource) GetEndpoint() *fhir.URL {
	if mhs == nil {
		return nil
	}
	return mhs.Endpoint
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhs *MessageHeaderSource) GetExtension() []*fhir.Extension {
	if mhs == nil {
		return nil
	}
	return mhs.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhs *MessageHeaderSource) GetID() string {
	if mhs == nil {
		return ""
	}
	return mhs.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhs *MessageHeaderSource) GetModifierExtension() []*fhir.Extension {
	if mhs == nil {
		return nil
	}
	return mhs.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhs *MessageHeaderSource) GetName() *fhir.String {
	if mhs == nil {
		return nil
	}
	return mhs.Name
}

// GetSoftware returns the value of the field Software.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhs *MessageHeaderSource) GetSoftware() *fhir.String {
	if mhs == nil {
		return nil
	}
	return mhs.Software
}

// GetVersion returns the value of the field Version.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mhs *MessageHeaderSource) GetVersion() *fhir.String {
	if mhs == nil {
		return nil
	}
	return mhs.Version
}

func (mh *MessageHeader) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (mh *MessageHeader) UnmarshalJSON(data []byte) error {
	var raw struct {
		Author      *fhir.Reference             `json:"author"`
		Contained   []fhir.Resource             `json:"contained"`
		Definition  *fhir.Canonical             `json:"definition"`
		Destination []*MessageHeaderDestination `json:"destination"`
		Enterer     *fhir.Reference             `json:"enterer"`
		EventCoding *fhir.Coding                `json:"eventCoding"`
		EventURI    *fhir.URI                   `json:"eventURI"`
		Extension   []*fhir.Extension           `json:"extension"`
		Focus       []*fhir.Reference           `json:"focus"`

		ID                string                 `json:"id"`
		ImplicitRules     *fhir.URI              `json:"implicitRules"`
		Language          *fhir.Code             `json:"language"`
		Meta              *fhir.Meta             `json:"meta"`
		ModifierExtension []*fhir.Extension      `json:"modifierExtension"`
		Reason            *fhir.CodeableConcept  `json:"reason"`
		Response          *MessageHeaderResponse `json:"response"`
		Responsible       *fhir.Reference        `json:"responsible"`
		Sender            *fhir.Reference        `json:"sender"`
		Source            *MessageHeaderSource   `json:"source"`
		Text              *fhir.Narrative        `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mh.Author = raw.Author
	mh.Contained = raw.Contained
	mh.Definition = raw.Definition
	mh.Destination = raw.Destination
	mh.Enterer = raw.Enterer
	mh.Event, err = validate.SelectOneOf[fhir.Element]("MessageHeader.event",
		raw.EventCoding,
		raw.EventURI)
	if err != nil {
		return err
	}
	mh.Extension = raw.Extension
	mh.Focus = raw.Focus
	mh.ID = raw.ID
	mh.ImplicitRules = raw.ImplicitRules
	mh.Language = raw.Language
	mh.Meta = raw.Meta
	mh.ModifierExtension = raw.ModifierExtension
	mh.Reason = raw.Reason
	mh.Response = raw.Response
	mh.Responsible = raw.Responsible
	mh.Sender = raw.Sender
	mh.Source = raw.Source
	mh.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*MessageHeader)(nil)
var _ json.Unmarshaler = (*MessageHeader)(nil)

func (mhd *MessageHeaderDestination) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (mhd *MessageHeaderDestination) UnmarshalJSON(data []byte) error {
	var raw struct {
		Endpoint  *fhir.URL         `json:"endpoint"`
		Extension []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		Name              *fhir.String      `json:"name"`
		Receiver          *fhir.Reference   `json:"receiver"`
		Target            *fhir.Reference   `json:"target"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mhd.Endpoint = raw.Endpoint
	mhd.Extension = raw.Extension
	mhd.ID = raw.ID
	mhd.ModifierExtension = raw.ModifierExtension
	mhd.Name = raw.Name
	mhd.Receiver = raw.Receiver
	mhd.Target = raw.Target
	return nil
}

var _ json.Marshaler = (*MessageHeaderDestination)(nil)
var _ json.Unmarshaler = (*MessageHeaderDestination)(nil)

func (mhr *MessageHeaderResponse) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (mhr *MessageHeaderResponse) UnmarshalJSON(data []byte) error {
	var raw struct {
		Code      *fhir.Code        `json:"code"`
		Details   *fhir.Reference   `json:"details"`
		Extension []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		Identifier        *fhir.ID          `json:"identifier"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mhr.Code = raw.Code
	mhr.Details = raw.Details
	mhr.Extension = raw.Extension
	mhr.ID = raw.ID
	mhr.Identifier = raw.Identifier
	mhr.ModifierExtension = raw.ModifierExtension
	return nil
}

var _ json.Marshaler = (*MessageHeaderResponse)(nil)
var _ json.Unmarshaler = (*MessageHeaderResponse)(nil)

func (mhs *MessageHeaderSource) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (mhs *MessageHeaderSource) UnmarshalJSON(data []byte) error {
	var raw struct {
		Contact   *fhir.ContactPoint `json:"contact"`
		Endpoint  *fhir.URL          `json:"endpoint"`
		Extension []*fhir.Extension  `json:"extension"`

		ID                string            `json:"id"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		Name              *fhir.String      `json:"name"`
		Software          *fhir.String      `json:"software"`
		Version           *fhir.String      `json:"version"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mhs.Contact = raw.Contact
	mhs.Endpoint = raw.Endpoint
	mhs.Extension = raw.Extension
	mhs.ID = raw.ID
	mhs.ModifierExtension = raw.ModifierExtension
	mhs.Name = raw.Name
	mhs.Software = raw.Software
	mhs.Version = raw.Version
	return nil
}

var _ json.Marshaler = (*MessageHeaderSource)(nil)
var _ json.Unmarshaler = (*MessageHeaderSource)(nil)
