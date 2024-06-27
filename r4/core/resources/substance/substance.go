// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package substance

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// A homogeneous material with a definite composition.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Substance
//   - Source File: StructureDefinition-Substance.json
type Substance struct {

	// A code that classifies the general type of substance. This is used for
	// searching, sorting and display purposes.
	Category []*fhir.CodeableConcept `fhirpath:"category"`

	// A code (or set of codes) that identify this substance.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// A description of the substance - its appearance, handling requirements, and
	// other usage notes.
	Description *fhir.String `fhirpath:"description"`

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

	// Unique identifier for the substance.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// A substance can be composed of other substances.
	Ingredient []*SubstanceIngredient `fhirpath:"ingredient"`

	// Substance may be used to describe a kind of substance, or a specific
	// package/container of the substance: an instance.
	Instance []*SubstanceInstance `fhirpath:"instance"`

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

	// A code to indicate if the substance is actively used.
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

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetCategory() []*fhir.CodeableConcept {
	if s == nil {
		return nil
	}
	return s.Category
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetCode() *fhir.CodeableConcept {
	if s == nil {
		return nil
	}
	return s.Code
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetContained() []fhir.Resource {
	if s == nil {
		return nil
	}
	return s.Contained
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetDescription() *fhir.String {
	if s == nil {
		return nil
	}
	return s.Description
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetExtension() []*fhir.Extension {
	if s == nil {
		return nil
	}
	return s.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetID() string {
	if s == nil {
		return ""
	}
	return s.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetIdentifier() []*fhir.Identifier {
	if s == nil {
		return nil
	}
	return s.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetImplicitRules() *fhir.URI {
	if s == nil {
		return nil
	}
	return s.ImplicitRules
}

// GetIngredient returns the value of the field Ingredient.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetIngredient() []*SubstanceIngredient {
	if s == nil {
		return nil
	}
	return s.Ingredient
}

// GetInstance returns the value of the field Instance.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetInstance() []*SubstanceInstance {
	if s == nil {
		return nil
	}
	return s.Instance
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetLanguage() *fhir.Code {
	if s == nil {
		return nil
	}
	return s.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetMeta() *fhir.Meta {
	if s == nil {
		return nil
	}
	return s.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetModifierExtension() []*fhir.Extension {
	if s == nil {
		return nil
	}
	return s.ModifierExtension
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetStatus() *fhir.Code {
	if s == nil {
		return nil
	}
	return s.Status
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (s *Substance) GetText() *fhir.Narrative {
	if s == nil {
		return nil
	}
	return s.Text
}

// Composition information about the substance// A substance can be composed of other substances.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Substance.json
type SubstanceIngredient struct {

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

	// The amount of the ingredient in the substance - a concentration ratio.
	Quantity *fhir.Ratio `fhirpath:"quantity"`

	// Another substance that is a component of this substance.
	Substance fhir.Element `fhirpath:"substance"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (si *SubstanceIngredient) GetExtension() []*fhir.Extension {
	if si == nil {
		return nil
	}
	return si.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (si *SubstanceIngredient) GetID() string {
	if si == nil {
		return ""
	}
	return si.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (si *SubstanceIngredient) GetModifierExtension() []*fhir.Extension {
	if si == nil {
		return nil
	}
	return si.ModifierExtension
}

// GetQuantity returns the value of the field Quantity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (si *SubstanceIngredient) GetQuantity() *fhir.Ratio {
	if si == nil {
		return nil
	}
	return si.Quantity
}

// GetSubstance returns the value of the field Substance.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (si *SubstanceIngredient) GetSubstance() fhir.Element {
	if si == nil {
		return nil
	}
	return si.Substance
}

// GetSubstanceCodeableConcept returns the value of the field Substance.
func (si *SubstanceIngredient) GetSubstanceCodeableConcept() *fhir.CodeableConcept {
	if si == nil {
		return nil
	}
	val, ok := si.Substance.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetSubstanceReference returns the value of the field Substance.
func (si *SubstanceIngredient) GetSubstanceReference() *fhir.Reference {
	if si == nil {
		return nil
	}
	val, ok := si.Substance.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
}

// If this describes a specific package/container of the substance// Substance may be used to describe a kind of substance, or a specific
// package/container of the substance: an instance.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Substance.json
type SubstanceInstance struct {

	// When the substance is no longer valid to use. For some substances, a single
	// arbitrary date is used for expiry.
	Expiry *fhir.DateTime `fhirpath:"expiry"`

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

	// Identifier associated with the package/container (usually a label affixed
	// directly).
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

	// The amount of the substance.
	Quantity *fhir.Quantity `fhirpath:"quantity"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExpiry returns the value of the field Expiry.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (si *SubstanceInstance) GetExpiry() *fhir.DateTime {
	if si == nil {
		return nil
	}
	return si.Expiry
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (si *SubstanceInstance) GetExtension() []*fhir.Extension {
	if si == nil {
		return nil
	}
	return si.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (si *SubstanceInstance) GetID() string {
	if si == nil {
		return ""
	}
	return si.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (si *SubstanceInstance) GetIdentifier() *fhir.Identifier {
	if si == nil {
		return nil
	}
	return si.Identifier
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (si *SubstanceInstance) GetModifierExtension() []*fhir.Extension {
	if si == nil {
		return nil
	}
	return si.ModifierExtension
}

// GetQuantity returns the value of the field Quantity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (si *SubstanceInstance) GetQuantity() *fhir.Quantity {
	if si == nil {
		return nil
	}
	return si.Quantity
}

func (s *Substance) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (s *Substance) UnmarshalJSON(data []byte) error {
	var raw struct {
		Category    []*fhir.CodeableConcept `json:"category"`
		Code        *fhir.CodeableConcept   `json:"code"`
		Contained   []fhir.Resource         `json:"contained"`
		Description *fhir.String            `json:"description"`
		Extension   []*fhir.Extension       `json:"extension"`

		ID                string                 `json:"id"`
		Identifier        []*fhir.Identifier     `json:"identifier"`
		ImplicitRules     *fhir.URI              `json:"implicitRules"`
		Ingredient        []*SubstanceIngredient `json:"ingredient"`
		Instance          []*SubstanceInstance   `json:"instance"`
		Language          *fhir.Code             `json:"language"`
		Meta              *fhir.Meta             `json:"meta"`
		ModifierExtension []*fhir.Extension      `json:"modifierExtension"`
		Status            *fhir.Code             `json:"status"`
		Text              *fhir.Narrative        `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	s.Category = raw.Category
	s.Code = raw.Code
	s.Contained = raw.Contained
	s.Description = raw.Description
	s.Extension = raw.Extension
	s.ID = raw.ID
	s.Identifier = raw.Identifier
	s.ImplicitRules = raw.ImplicitRules
	s.Ingredient = raw.Ingredient
	s.Instance = raw.Instance
	s.Language = raw.Language
	s.Meta = raw.Meta
	s.ModifierExtension = raw.ModifierExtension
	s.Status = raw.Status
	s.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*Substance)(nil)
var _ json.Unmarshaler = (*Substance)(nil)

func (si *SubstanceIngredient) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (si *SubstanceIngredient) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                       string                `json:"id"`
		ModifierExtension        []*fhir.Extension     `json:"modifierExtension"`
		Quantity                 *fhir.Ratio           `json:"quantity"`
		SubstanceCodeableConcept *fhir.CodeableConcept `json:"substanceCodeableConcept"`
		SubstanceReference       *fhir.Reference       `json:"substanceReference"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	si.Extension = raw.Extension
	si.ID = raw.ID
	si.ModifierExtension = raw.ModifierExtension
	si.Quantity = raw.Quantity
	si.Substance, err = validate.SelectOneOf[fhir.Element]("Substance.ingredient.substance",
		raw.SubstanceCodeableConcept,
		raw.SubstanceReference)
	if err != nil {
		return err
	}
	return nil
}

var _ json.Marshaler = (*SubstanceIngredient)(nil)
var _ json.Unmarshaler = (*SubstanceIngredient)(nil)

func (si *SubstanceInstance) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (si *SubstanceInstance) UnmarshalJSON(data []byte) error {
	var raw struct {
		Expiry    *fhir.DateTime    `json:"expiry"`
		Extension []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		Identifier        *fhir.Identifier  `json:"identifier"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		Quantity          *fhir.Quantity    `json:"quantity"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	si.Expiry = raw.Expiry
	si.Extension = raw.Extension
	si.ID = raw.ID
	si.Identifier = raw.Identifier
	si.ModifierExtension = raw.ModifierExtension
	si.Quantity = raw.Quantity
	return nil
}

var _ json.Marshaler = (*SubstanceInstance)(nil)
var _ json.Unmarshaler = (*SubstanceInstance)(nil)
