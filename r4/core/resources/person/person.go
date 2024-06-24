// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package person

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Demographics and administrative information about a person independent of a
// specific health-related context.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Person
//   - Source File: StructureDefinition-Person.json
type Person struct {

	// Whether this person's record is in active use.
	Active *fhir.Boolean `fhirpath:"active"`

	// One or more addresses for the person.
	Address []*fhir.Address `fhirpath:"address"`

	// The birth date for the person.
	BirthDate *fhir.Date `fhirpath:"birthDate"`

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

	// Administrative Gender.
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

	// Link to a resource that concerns the same actual person.
	Link []*PersonLink `fhirpath:"link"`

	// The organization that is the custodian of the person record.
	ManagingOrganization *fhir.Reference `fhirpath:"managingOrganization"`

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

	// An image that can be displayed as a thumbnail of the person to enhance the
	// identification of the individual.
	Photo *fhir.Attachment `fhirpath:"photo"`

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
func (p *Person) GetActive() *fhir.Boolean {
	if p == nil {
		return nil
	}
	return p.Active
}

// GetAddress returns the value of the field Address.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetAddress() []*fhir.Address {
	if p == nil {
		return nil
	}
	return p.Address
}

// GetBirthDate returns the value of the field BirthDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetBirthDate() *fhir.Date {
	if p == nil {
		return nil
	}
	return p.BirthDate
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetContained() []fhir.Resource {
	if p == nil {
		return nil
	}
	return p.Contained
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetExtension() []*fhir.Extension {
	if p == nil {
		return nil
	}
	return p.Extension
}

// GetGender returns the value of the field Gender.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetGender() *fhir.Code {
	if p == nil {
		return nil
	}
	return p.Gender
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetIdentifier() []*fhir.Identifier {
	if p == nil {
		return nil
	}
	return p.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetImplicitRules() *fhir.URI {
	if p == nil {
		return nil
	}
	return p.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetLanguage() *fhir.Code {
	if p == nil {
		return nil
	}
	return p.Language
}

// GetLink returns the value of the field Link.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetLink() []*PersonLink {
	if p == nil {
		return nil
	}
	return p.Link
}

// GetManagingOrganization returns the value of the field ManagingOrganization.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetManagingOrganization() *fhir.Reference {
	if p == nil {
		return nil
	}
	return p.ManagingOrganization
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetMeta() *fhir.Meta {
	if p == nil {
		return nil
	}
	return p.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetModifierExtension() []*fhir.Extension {
	if p == nil {
		return nil
	}
	return p.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetName() []*fhir.HumanName {
	if p == nil {
		return nil
	}
	return p.Name
}

// GetPhoto returns the value of the field Photo.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetPhoto() *fhir.Attachment {
	if p == nil {
		return nil
	}
	return p.Photo
}

// GetTelecom returns the value of the field Telecom.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetTelecom() []*fhir.ContactPoint {
	if p == nil {
		return nil
	}
	return p.Telecom
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Person) GetText() *fhir.Narrative {
	if p == nil {
		return nil
	}
	return p.Text
}

// Link to a resource that concerns the same actual person// Link to a resource that concerns the same actual person.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Person.json
type PersonLink struct {

	// Level of assurance that this link is associated with the target resource.
	Assurance *fhir.Code `fhirpath:"assurance"`

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

	// The resource to which this actual person is associated.
	Target *fhir.Reference `fhirpath:"target"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAssurance returns the value of the field Assurance.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pl *PersonLink) GetAssurance() *fhir.Code {
	if pl == nil {
		return nil
	}
	return pl.Assurance
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pl *PersonLink) GetExtension() []*fhir.Extension {
	if pl == nil {
		return nil
	}
	return pl.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pl *PersonLink) GetID() string {
	if pl == nil {
		return ""
	}
	return pl.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pl *PersonLink) GetModifierExtension() []*fhir.Extension {
	if pl == nil {
		return nil
	}
	return pl.ModifierExtension
}

// GetTarget returns the value of the field Target.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pl *PersonLink) GetTarget() *fhir.Reference {
	if pl == nil {
		return nil
	}
	return pl.Target
}