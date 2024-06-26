// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Base StructureDefinition for SubstanceAmount Type: Chemical substances are a
// single substance type whose primary defining element is the molecular
// structure. Chemical substances shall be defined on the basis of their
// complete covalent molecular structure; the presence of a salt (counter-ion)
// and/or solvates (water, alcohols) is also captured. Purity, grade, physical
// form or particle size are not taken into account in the definition of a
// chemical substance or in the assignment of a Substance ID.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/SubstanceAmount
//   - Source File: StructureDefinition-SubstanceAmount.json
type SubstanceAmount struct {

	// A textual comment on a numeric value.
	AmountText *String `fhirpath:"amountText"`

	// Most elements that require a quantitative value will also have a field
	// called amount type. Amount type should always be specified because the
	// actual value of the amount is often dependent on it. EXAMPLE: In capturing
	// the actual relative amounts of substances or molecular fragments it is
	// essential to indicate whether the amount refers to a mole ratio or weight
	// ratio. For any given element an effort should be made to use same the amount
	// type for all related definitional elements.
	AmountType *CodeableConcept `fhirpath:"amountType"`

	// Used to capture quantitative values for a variety of elements. If only
	// limits are given, the arithmetic mean would be the average. If only a single
	// definite value for a given element is given, it would be captured in this
	// field.
	Amount Element `fhirpath:"amount"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*Extension `fhirpath:"extension"`

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
	ModifierExtension []*Extension `fhirpath:"modifierExtension"`

	// Reference range of possible or expected values.
	ReferenceRange *SubstanceAmountReferenceRange `fhirpath:"referenceRange"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAmountText returns the value of the field AmountText.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sa *SubstanceAmount) GetAmountText() *String {
	if sa == nil {
		return nil
	}
	return sa.AmountText
}

// GetAmountType returns the value of the field AmountType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sa *SubstanceAmount) GetAmountType() *CodeableConcept {
	if sa == nil {
		return nil
	}
	return sa.AmountType
}

// GetAmount returns the value of the field Amount.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sa *SubstanceAmount) GetAmount() Element {
	if sa == nil {
		return nil
	}
	return sa.Amount
}

// GetAmountQuantity returns the value of the field Amount.
func (sa *SubstanceAmount) GetAmountQuantity() *Quantity {
	if sa == nil {
		return nil
	}
	val, ok := sa.Amount.(*Quantity)
	if !ok {
		return nil
	}
	return val
}

// GetAmountRange returns the value of the field Amount.
func (sa *SubstanceAmount) GetAmountRange() *Range {
	if sa == nil {
		return nil
	}
	val, ok := sa.Amount.(*Range)
	if !ok {
		return nil
	}
	return val
}

// GetAmountString returns the value of the field Amount.
func (sa *SubstanceAmount) GetAmountString() *String {
	if sa == nil {
		return nil
	}
	val, ok := sa.Amount.(*String)
	if !ok {
		return nil
	}
	return val
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sa *SubstanceAmount) GetExtension() []*Extension {
	if sa == nil {
		return nil
	}
	return sa.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sa *SubstanceAmount) GetID() string {
	if sa == nil {
		return ""
	}
	return sa.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sa *SubstanceAmount) GetModifierExtension() []*Extension {
	if sa == nil {
		return nil
	}
	return sa.ModifierExtension
}

// GetReferenceRange returns the value of the field ReferenceRange.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sa *SubstanceAmount) GetReferenceRange() *SubstanceAmountReferenceRange {
	if sa == nil {
		return nil
	}
	return sa.ReferenceRange
}

// Reference range of possible or expected values// Reference range of possible or expected values.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-SubstanceAmount.json
type SubstanceAmountReferenceRange struct {

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*Extension `fhirpath:"extension"`

	// Upper limit possible or expected.
	HighLimit *Quantity `fhirpath:"highLimit"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID string `fhirpath:"id"`

	// Lower limit possible or expected.
	LowLimit *Quantity `fhirpath:"lowLimit"`

	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sarr *SubstanceAmountReferenceRange) GetExtension() []*Extension {
	if sarr == nil {
		return nil
	}
	return sarr.Extension
}

// GetHighLimit returns the value of the field HighLimit.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sarr *SubstanceAmountReferenceRange) GetHighLimit() *Quantity {
	if sarr == nil {
		return nil
	}
	return sarr.HighLimit
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sarr *SubstanceAmountReferenceRange) GetID() string {
	if sarr == nil {
		return ""
	}
	return sarr.ID
}

// GetLowLimit returns the value of the field LowLimit.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sarr *SubstanceAmountReferenceRange) GetLowLimit() *Quantity {
	if sarr == nil {
		return nil
	}
	return sarr.LowLimit
}

func (sa *SubstanceAmount) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (sa *SubstanceAmount) UnmarshalJSON(data []byte) error {
	var raw struct {
		AmountText     *String          `json:"amountText"`
		AmountType     *CodeableConcept `json:"amountType"`
		AmountQuantity *Quantity        `json:"amountQuantity"`
		AmountRange    *Range           `json:"amountRange"`
		AmountString   *String          `json:"amountString"`
		Extension      []*Extension     `json:"extension"`

		ID                string                         `json:"id"`
		ModifierExtension []*Extension                   `json:"modifierExtension"`
		ReferenceRange    *SubstanceAmountReferenceRange `json:"referenceRange"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	sa.AmountText = raw.AmountText
	sa.AmountType = raw.AmountType
	sa.Amount, err = validate.SelectOneOf[Element]("SubstanceAmount.amount",
		raw.AmountQuantity,
		raw.AmountRange,
		raw.AmountString)
	if err != nil {
		return err
	}
	sa.Extension = raw.Extension
	sa.ID = raw.ID
	sa.ModifierExtension = raw.ModifierExtension
	sa.ReferenceRange = raw.ReferenceRange
	return nil
}

var _ json.Marshaler = (*SubstanceAmount)(nil)
var _ json.Unmarshaler = (*SubstanceAmount)(nil)

func (sarr *SubstanceAmountReferenceRange) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (sarr *SubstanceAmountReferenceRange) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*Extension `json:"extension"`
		HighLimit *Quantity    `json:"highLimit"`

		ID       string    `json:"id"`
		LowLimit *Quantity `json:"lowLimit"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	sarr.Extension = raw.Extension
	sarr.HighLimit = raw.HighLimit
	sarr.ID = raw.ID
	sarr.LowLimit = raw.LowLimit
	return nil
}

var _ json.Marshaler = (*SubstanceAmountReferenceRange)(nil)
var _ json.Unmarshaler = (*SubstanceAmountReferenceRange)(nil)
