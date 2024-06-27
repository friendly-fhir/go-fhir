// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package medicinalproductinteraction

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// The interactions of the medicinal product with other medicinal products, or
// other forms of interactions.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/MedicinalProductInteraction
//   - Source File: StructureDefinition-MedicinalProductInteraction.json
type MedicinalProductInteraction struct {

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The interaction described.
	Description *fhir.String `fhirpath:"description"`

	// The effect of the interaction, for example "reduced gastric absorption of
	// primary medication".
	Effect *fhir.CodeableConcept `fhirpath:"effect"`

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

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The incidence of the interaction, e.g. theoretical, observed.
	Incidence *fhir.CodeableConcept `fhirpath:"incidence"`

	// The specific medication, food or laboratory test that interacts.
	Interactant []*MedicinalProductInteractionInteractant `fhirpath:"interactant"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// Actions for managing the interaction.
	Management *fhir.CodeableConcept `fhirpath:"management"`

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

	// The medication for which this is a described interaction.
	Subject []*fhir.Reference `fhirpath:"subject"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// The type of the interaction e.g. drug-drug interaction, drug-food
	// interaction, drug-lab test interaction.
	Type *fhir.CodeableConcept `fhirpath:"type"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetContained() []fhir.Resource {
	if mpi == nil {
		return nil
	}
	return mpi.Contained
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetDescription() *fhir.String {
	if mpi == nil {
		return nil
	}
	return mpi.Description
}

// GetEffect returns the value of the field Effect.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetEffect() *fhir.CodeableConcept {
	if mpi == nil {
		return nil
	}
	return mpi.Effect
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetExtension() []*fhir.Extension {
	if mpi == nil {
		return nil
	}
	return mpi.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetID() string {
	if mpi == nil {
		return ""
	}
	return mpi.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetImplicitRules() *fhir.URI {
	if mpi == nil {
		return nil
	}
	return mpi.ImplicitRules
}

// GetIncidence returns the value of the field Incidence.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetIncidence() *fhir.CodeableConcept {
	if mpi == nil {
		return nil
	}
	return mpi.Incidence
}

// GetInteractant returns the value of the field Interactant.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetInteractant() []*MedicinalProductInteractionInteractant {
	if mpi == nil {
		return nil
	}
	return mpi.Interactant
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetLanguage() *fhir.Code {
	if mpi == nil {
		return nil
	}
	return mpi.Language
}

// GetManagement returns the value of the field Management.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetManagement() *fhir.CodeableConcept {
	if mpi == nil {
		return nil
	}
	return mpi.Management
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetMeta() *fhir.Meta {
	if mpi == nil {
		return nil
	}
	return mpi.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetModifierExtension() []*fhir.Extension {
	if mpi == nil {
		return nil
	}
	return mpi.ModifierExtension
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetSubject() []*fhir.Reference {
	if mpi == nil {
		return nil
	}
	return mpi.Subject
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetText() *fhir.Narrative {
	if mpi == nil {
		return nil
	}
	return mpi.Text
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductInteraction) GetType() *fhir.CodeableConcept {
	if mpi == nil {
		return nil
	}
	return mpi.Type
}

// The specific medication, food or laboratory test that interacts// The specific medication, food or laboratory test that interacts.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MedicinalProductInteraction.json
type MedicinalProductInteractionInteractant struct {

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

	// The specific medication, food or laboratory test that interacts.
	Item fhir.Element `fhirpath:"item"`

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

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpii *MedicinalProductInteractionInteractant) GetExtension() []*fhir.Extension {
	if mpii == nil {
		return nil
	}
	return mpii.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpii *MedicinalProductInteractionInteractant) GetID() string {
	if mpii == nil {
		return ""
	}
	return mpii.ID
}

// GetItem returns the value of the field Item.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpii *MedicinalProductInteractionInteractant) GetItem() fhir.Element {
	if mpii == nil {
		return nil
	}
	return mpii.Item
}

// GetItemReference returns the value of the field Item.
func (mpii *MedicinalProductInteractionInteractant) GetItemReference() *fhir.Reference {
	if mpii == nil {
		return nil
	}
	val, ok := mpii.Item.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
}

// GetItemCodeableConcept returns the value of the field Item.
func (mpii *MedicinalProductInteractionInteractant) GetItemCodeableConcept() *fhir.CodeableConcept {
	if mpii == nil {
		return nil
	}
	val, ok := mpii.Item.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
} // GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpii *MedicinalProductInteractionInteractant) GetModifierExtension() []*fhir.Extension {
	if mpii == nil {
		return nil
	}
	return mpii.ModifierExtension
}

func (mpi *MedicinalProductInteraction) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (mpi *MedicinalProductInteraction) UnmarshalJSON(data []byte) error {
	var raw struct {
		Contained   []fhir.Resource       `json:"contained"`
		Description *fhir.String          `json:"description"`
		Effect      *fhir.CodeableConcept `json:"effect"`
		Extension   []*fhir.Extension     `json:"extension"`

		ID                string                                    `json:"id"`
		ImplicitRules     *fhir.URI                                 `json:"implicitRules"`
		Incidence         *fhir.CodeableConcept                     `json:"incidence"`
		Interactant       []*MedicinalProductInteractionInteractant `json:"interactant"`
		Language          *fhir.Code                                `json:"language"`
		Management        *fhir.CodeableConcept                     `json:"management"`
		Meta              *fhir.Meta                                `json:"meta"`
		ModifierExtension []*fhir.Extension                         `json:"modifierExtension"`
		Subject           []*fhir.Reference                         `json:"subject"`
		Text              *fhir.Narrative                           `json:"text"`
		Type              *fhir.CodeableConcept                     `json:"type"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mpi.Contained = raw.Contained
	mpi.Description = raw.Description
	mpi.Effect = raw.Effect
	mpi.Extension = raw.Extension
	mpi.ID = raw.ID
	mpi.ImplicitRules = raw.ImplicitRules
	mpi.Incidence = raw.Incidence
	mpi.Interactant = raw.Interactant
	mpi.Language = raw.Language
	mpi.Management = raw.Management
	mpi.Meta = raw.Meta
	mpi.ModifierExtension = raw.ModifierExtension
	mpi.Subject = raw.Subject
	mpi.Text = raw.Text
	mpi.Type = raw.Type
	return nil
}

var _ json.Marshaler = (*MedicinalProductInteraction)(nil)
var _ json.Unmarshaler = (*MedicinalProductInteraction)(nil)

func (mpii *MedicinalProductInteractionInteractant) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (mpii *MedicinalProductInteractionInteractant) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                  string                `json:"id"`
		ItemReference       *fhir.Reference       `json:"itemReference"`
		ItemCodeableConcept *fhir.CodeableConcept `json:"itemCodeableConcept"`
		ModifierExtension   []*fhir.Extension     `json:"modifierExtension"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mpii.Extension = raw.Extension
	mpii.ID = raw.ID
	mpii.Item, err = validate.SelectOneOf[fhir.Element]("MedicinalProductInteraction.interactant.item",
		raw.ItemReference,
		raw.ItemCodeableConcept)
	if err != nil {
		return err
	}
	mpii.ModifierExtension = raw.ModifierExtension
	return nil
}

var _ json.Marshaler = (*MedicinalProductInteractionInteractant)(nil)
var _ json.Unmarshaler = (*MedicinalProductInteractionInteractant)(nil)
