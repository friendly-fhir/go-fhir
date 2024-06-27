// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package relatedperson

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Information about a person that is involved in the care for a patient, but
// who is not the target of healthcare, nor has a formal responsibility in the
// care process.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/RelatedPerson
//   - Source File: StructureDefinition-RelatedPerson.json
type RelatedPerson struct {

	// Whether this related person record is in active use.
	Active *fhir.Boolean `fhirpath:"active"`

	// Address where the related person can be contacted or visited.
	Address []*fhir.Address `fhirpath:"address"`

	// The date on which the related person was born.
	BirthDate *fhir.Date `fhirpath:"birthDate"`

	// A language which may be used to communicate with about the patient's health.
	Communication []*RelatedPersonCommunication `fhirpath:"communication"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// May be used to represent additional information that is not part of the
	// basic definition of the resource. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Administrative Gender - the gender that the person is considered to have for
	// administration and record keeping purposes.
	Gender *fhir.Code `fhirpath:"gender"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// Identifier for a person within a particular scope.
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

	// A name associated with the person.
	Name []*fhir.HumanName `fhirpath:"name"`

	// The patient this person is related to.
	Patient *fhir.Reference `fhirpath:"patient"`

	// The period of time during which this relationship is or was active. If there
	// are no dates defined, then the interval is unknown.
	Period *fhir.Period `fhirpath:"period"`

	// Image of the person.
	Photo []*fhir.Attachment `fhirpath:"photo"`

	// The nature of the relationship between a patient and the related person.
	Relationship []*fhir.CodeableConcept `fhirpath:"relationship"`

	// A contact detail for the person, e.g. a telephone number or an email
	// address.
	Telecom []*fhir.ContactPoint `fhirpath:"telecom"`

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

// GetActive returns the value of the field Active.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetActive() *fhir.Boolean {
	if rp == nil {
		return nil
	}
	return rp.Active
}

// GetAddress returns the value of the field Address.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetAddress() []*fhir.Address {
	if rp == nil {
		return nil
	}
	return rp.Address
}

// GetBirthDate returns the value of the field BirthDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetBirthDate() *fhir.Date {
	if rp == nil {
		return nil
	}
	return rp.BirthDate
}

// GetCommunication returns the value of the field Communication.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetCommunication() []*RelatedPersonCommunication {
	if rp == nil {
		return nil
	}
	return rp.Communication
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetContained() []fhir.Resource {
	if rp == nil {
		return nil
	}
	return rp.Contained
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetExtension() []*fhir.Extension {
	if rp == nil {
		return nil
	}
	return rp.Extension
}

// GetGender returns the value of the field Gender.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetGender() *fhir.Code {
	if rp == nil {
		return nil
	}
	return rp.Gender
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetID() string {
	if rp == nil {
		return ""
	}
	return rp.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetIdentifier() []*fhir.Identifier {
	if rp == nil {
		return nil
	}
	return rp.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetImplicitRules() *fhir.URI {
	if rp == nil {
		return nil
	}
	return rp.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetLanguage() *fhir.Code {
	if rp == nil {
		return nil
	}
	return rp.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetMeta() *fhir.Meta {
	if rp == nil {
		return nil
	}
	return rp.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetModifierExtension() []*fhir.Extension {
	if rp == nil {
		return nil
	}
	return rp.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetName() []*fhir.HumanName {
	if rp == nil {
		return nil
	}
	return rp.Name
}

// GetPatient returns the value of the field Patient.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetPatient() *fhir.Reference {
	if rp == nil {
		return nil
	}
	return rp.Patient
}

// GetPeriod returns the value of the field Period.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetPeriod() *fhir.Period {
	if rp == nil {
		return nil
	}
	return rp.Period
}

// GetPhoto returns the value of the field Photo.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetPhoto() []*fhir.Attachment {
	if rp == nil {
		return nil
	}
	return rp.Photo
}

// GetRelationship returns the value of the field Relationship.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetRelationship() []*fhir.CodeableConcept {
	if rp == nil {
		return nil
	}
	return rp.Relationship
}

// GetTelecom returns the value of the field Telecom.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetTelecom() []*fhir.ContactPoint {
	if rp == nil {
		return nil
	}
	return rp.Telecom
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rp *RelatedPerson) GetText() *fhir.Narrative {
	if rp == nil {
		return nil
	}
	return rp.Text
}

// A language which may be used to communicate with about the patient's health// A language which may be used to communicate with about the patient's health.// If no language is specified, this *implies* that the default local language
// is spoken. If you need to convey proficiency for multiple modes, then you
// need multiple RelatedPerson.Communication associations. If the RelatedPerson
// does not speak the default local language, then the Interpreter Required
// Standard can be used to explicitly declare that an interpreter is required.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-RelatedPerson.json
type RelatedPersonCommunication struct {

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

	// The ISO-639-1 alpha 2 code in lower case for the language, optionally
	// followed by a hyphen and the ISO-3166-1 alpha 2 code for the region in upper
	// case; e.g. "en" for English, or "en-US" for American English versus "en-EN"
	// for England English.
	Language *fhir.CodeableConcept `fhirpath:"language"`

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

	// Indicates whether or not the patient prefers this language (over other
	// languages he masters up a certain level).
	Preferred *fhir.Boolean `fhirpath:"preferred"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rpc *RelatedPersonCommunication) GetExtension() []*fhir.Extension {
	if rpc == nil {
		return nil
	}
	return rpc.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rpc *RelatedPersonCommunication) GetID() string {
	if rpc == nil {
		return ""
	}
	return rpc.ID
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rpc *RelatedPersonCommunication) GetLanguage() *fhir.CodeableConcept {
	if rpc == nil {
		return nil
	}
	return rpc.Language
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rpc *RelatedPersonCommunication) GetModifierExtension() []*fhir.Extension {
	if rpc == nil {
		return nil
	}
	return rpc.ModifierExtension
}

// GetPreferred returns the value of the field Preferred.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (rpc *RelatedPersonCommunication) GetPreferred() *fhir.Boolean {
	if rpc == nil {
		return nil
	}
	return rpc.Preferred
}

func (rp *RelatedPerson) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (rp *RelatedPerson) UnmarshalJSON(data []byte) error {
	var raw struct {
		Active        *fhir.Boolean                 `json:"active"`
		Address       []*fhir.Address               `json:"address"`
		BirthDate     *fhir.Date                    `json:"birthDate"`
		Communication []*RelatedPersonCommunication `json:"communication"`
		Contained     []fhir.Resource               `json:"contained"`
		Extension     []*fhir.Extension             `json:"extension"`
		Gender        *fhir.Code                    `json:"gender"`

		ID                string                  `json:"id"`
		Identifier        []*fhir.Identifier      `json:"identifier"`
		ImplicitRules     *fhir.URI               `json:"implicitRules"`
		Language          *fhir.Code              `json:"language"`
		Meta              *fhir.Meta              `json:"meta"`
		ModifierExtension []*fhir.Extension       `json:"modifierExtension"`
		Name              []*fhir.HumanName       `json:"name"`
		Patient           *fhir.Reference         `json:"patient"`
		Period            *fhir.Period            `json:"period"`
		Photo             []*fhir.Attachment      `json:"photo"`
		Relationship      []*fhir.CodeableConcept `json:"relationship"`
		Telecom           []*fhir.ContactPoint    `json:"telecom"`
		Text              *fhir.Narrative         `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	rp.Active = raw.Active
	rp.Address = raw.Address
	rp.BirthDate = raw.BirthDate
	rp.Communication = raw.Communication
	rp.Contained = raw.Contained
	rp.Extension = raw.Extension
	rp.Gender = raw.Gender
	rp.ID = raw.ID
	rp.Identifier = raw.Identifier
	rp.ImplicitRules = raw.ImplicitRules
	rp.Language = raw.Language
	rp.Meta = raw.Meta
	rp.ModifierExtension = raw.ModifierExtension
	rp.Name = raw.Name
	rp.Patient = raw.Patient
	rp.Period = raw.Period
	rp.Photo = raw.Photo
	rp.Relationship = raw.Relationship
	rp.Telecom = raw.Telecom
	rp.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*RelatedPerson)(nil)
var _ json.Unmarshaler = (*RelatedPerson)(nil)

func (rpc *RelatedPersonCommunication) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (rpc *RelatedPersonCommunication) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                string                `json:"id"`
		Language          *fhir.CodeableConcept `json:"language"`
		ModifierExtension []*fhir.Extension     `json:"modifierExtension"`
		Preferred         *fhir.Boolean         `json:"preferred"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	rpc.Extension = raw.Extension
	rpc.ID = raw.ID
	rpc.Language = raw.Language
	rpc.ModifierExtension = raw.ModifierExtension
	rpc.Preferred = raw.Preferred
	return nil
}

var _ json.Marshaler = (*RelatedPersonCommunication)(nil)
var _ json.Unmarshaler = (*RelatedPersonCommunication)(nil)
