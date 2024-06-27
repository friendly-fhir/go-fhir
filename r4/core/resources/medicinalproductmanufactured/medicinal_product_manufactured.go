// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package medicinalproductmanufactured

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// The manufactured item as contained in the packaged medicinal product.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/MedicinalProductManufactured
//   - Source File: StructureDefinition-MedicinalProductManufactured.json
type MedicinalProductManufactured struct {

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

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// Ingredient.
	Ingredient []*fhir.Reference `fhirpath:"ingredient"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// Dose form as manufactured and before any transformation into the
	// pharmaceutical product.
	ManufacturedDoseForm *fhir.CodeableConcept `fhirpath:"manufacturedDoseForm"`

	// Manufacturer of the item (Note that this should be named "manufacturer" but
	// it currently causes technical issues).
	Manufacturer []*fhir.Reference `fhirpath:"manufacturer"`

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

	// Other codeable characteristics.
	OtherCharacteristics []*fhir.CodeableConcept `fhirpath:"otherCharacteristics"`

	// Dimensions, color etc.
	PhysicalCharacteristics *fhir.ProdCharacteristic `fhirpath:"physicalCharacteristics"`

	// The quantity or "count number" of the manufactured item.
	Quantity *fhir.Quantity `fhirpath:"quantity"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// The “real world” units in which the quantity of the manufactured item is
	// described.
	UnitOfPresentation *fhir.CodeableConcept `fhirpath:"unitOfPresentation"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetContained() []fhir.Resource {
	if mpm == nil {
		return nil
	}
	return mpm.Contained
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetExtension() []*fhir.Extension {
	if mpm == nil {
		return nil
	}
	return mpm.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetID() string {
	if mpm == nil {
		return ""
	}
	return mpm.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetImplicitRules() *fhir.URI {
	if mpm == nil {
		return nil
	}
	return mpm.ImplicitRules
}

// GetIngredient returns the value of the field Ingredient.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetIngredient() []*fhir.Reference {
	if mpm == nil {
		return nil
	}
	return mpm.Ingredient
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetLanguage() *fhir.Code {
	if mpm == nil {
		return nil
	}
	return mpm.Language
}

// GetManufacturedDoseForm returns the value of the field ManufacturedDoseForm.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetManufacturedDoseForm() *fhir.CodeableConcept {
	if mpm == nil {
		return nil
	}
	return mpm.ManufacturedDoseForm
}

// GetManufacturer returns the value of the field Manufacturer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetManufacturer() []*fhir.Reference {
	if mpm == nil {
		return nil
	}
	return mpm.Manufacturer
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetMeta() *fhir.Meta {
	if mpm == nil {
		return nil
	}
	return mpm.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetModifierExtension() []*fhir.Extension {
	if mpm == nil {
		return nil
	}
	return mpm.ModifierExtension
}

// GetOtherCharacteristics returns the value of the field OtherCharacteristics.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetOtherCharacteristics() []*fhir.CodeableConcept {
	if mpm == nil {
		return nil
	}
	return mpm.OtherCharacteristics
}

// GetPhysicalCharacteristics returns the value of the field PhysicalCharacteristics.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetPhysicalCharacteristics() *fhir.ProdCharacteristic {
	if mpm == nil {
		return nil
	}
	return mpm.PhysicalCharacteristics
}

// GetQuantity returns the value of the field Quantity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetQuantity() *fhir.Quantity {
	if mpm == nil {
		return nil
	}
	return mpm.Quantity
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetText() *fhir.Narrative {
	if mpm == nil {
		return nil
	}
	return mpm.Text
}

// GetUnitOfPresentation returns the value of the field UnitOfPresentation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpm *MedicinalProductManufactured) GetUnitOfPresentation() *fhir.CodeableConcept {
	if mpm == nil {
		return nil
	}
	return mpm.UnitOfPresentation
}

func (mpm *MedicinalProductManufactured) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (mpm *MedicinalProductManufactured) UnmarshalJSON(data []byte) error {
	var raw struct {
		Contained []fhir.Resource   `json:"contained"`
		Extension []*fhir.Extension `json:"extension"`

		ID                      string                   `json:"id"`
		ImplicitRules           *fhir.URI                `json:"implicitRules"`
		Ingredient              []*fhir.Reference        `json:"ingredient"`
		Language                *fhir.Code               `json:"language"`
		ManufacturedDoseForm    *fhir.CodeableConcept    `json:"manufacturedDoseForm"`
		Manufacturer            []*fhir.Reference        `json:"manufacturer"`
		Meta                    *fhir.Meta               `json:"meta"`
		ModifierExtension       []*fhir.Extension        `json:"modifierExtension"`
		OtherCharacteristics    []*fhir.CodeableConcept  `json:"otherCharacteristics"`
		PhysicalCharacteristics *fhir.ProdCharacteristic `json:"physicalCharacteristics"`
		Quantity                *fhir.Quantity           `json:"quantity"`
		Text                    *fhir.Narrative          `json:"text"`
		UnitOfPresentation      *fhir.CodeableConcept    `json:"unitOfPresentation"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mpm.Contained = raw.Contained
	mpm.Extension = raw.Extension
	mpm.ID = raw.ID
	mpm.ImplicitRules = raw.ImplicitRules
	mpm.Ingredient = raw.Ingredient
	mpm.Language = raw.Language
	mpm.ManufacturedDoseForm = raw.ManufacturedDoseForm
	mpm.Manufacturer = raw.Manufacturer
	mpm.Meta = raw.Meta
	mpm.ModifierExtension = raw.ModifierExtension
	mpm.OtherCharacteristics = raw.OtherCharacteristics
	mpm.PhysicalCharacteristics = raw.PhysicalCharacteristics
	mpm.Quantity = raw.Quantity
	mpm.Text = raw.Text
	mpm.UnitOfPresentation = raw.UnitOfPresentation
	return nil
}

var _ json.Marshaler = (*MedicinalProductManufactured)(nil)
var _ json.Unmarshaler = (*MedicinalProductManufactured)(nil)
