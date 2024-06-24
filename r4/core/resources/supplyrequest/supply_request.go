// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package supplyrequest

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// A record of a request for a medication, substance or device used in the
// healthcare setting.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/SupplyRequest
//   - Source File: StructureDefinition-SupplyRequest.json
type SupplyRequest struct {

	// When the request was made.
	AuthoredOn *fhir.DateTime `fhirpath:"authoredOn"`

	// Category of supply, e.g. central, non-stock, etc. This is used to support
	// work flows associated with the supply process.
	Category *fhir.CodeableConcept `fhirpath:"category"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// Where the supply is expected to come from.
	DeliverFrom *fhir.Reference `fhirpath:"deliverFrom"`

	// Where the supply is destined to go.
	DeliverTo *fhir.Reference `fhirpath:"deliverTo"`

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

	// Business identifiers assigned to this SupplyRequest by the author and/or
	// other systems. These identifiers remain constant as the resource is updated
	// and propagates from server to server.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The item that is requested to be supplied. This is either a link to a
	// resource representing the details of the item or a code that identifies the
	// item from a known list.
	Item fhir.Element `fhirpath:"item"`

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

	// When the request should be fulfilled.
	Occurrence fhir.Element `fhirpath:"occurrence"`

	// Specific parameters for the ordered item. For example, the size of the
	// indicated item.
	Parameter []*SupplyRequestParameter `fhirpath:"parameter"`

	// Indicates how quickly this SupplyRequest should be addressed with respect to
	// other requests.
	Priority *fhir.Code `fhirpath:"priority"`

	// The amount that is being ordered of the indicated item.
	Quantity *fhir.Quantity `fhirpath:"quantity"`

	// The reason why the supply item was requested.
	ReasonCode []*fhir.CodeableConcept `fhirpath:"reasonCode"`

	// The reason why the supply item was requested.
	ReasonReference []*fhir.Reference `fhirpath:"reasonReference"`

	// The device, practitioner, etc. who initiated the request.
	Requester *fhir.Reference `fhirpath:"requester"`

	// Status of the supply request.
	Status *fhir.Code `fhirpath:"status"`

	// Who is intended to fulfill the request.
	Supplier []*fhir.Reference `fhirpath:"supplier"`

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

// GetAuthoredOn returns the value of the field AuthoredOn.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetAuthoredOn() *fhir.DateTime {
	if sr == nil {
		return nil
	}
	return sr.AuthoredOn
}

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetCategory() *fhir.CodeableConcept {
	if sr == nil {
		return nil
	}
	return sr.Category
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetContained() []fhir.Resource {
	if sr == nil {
		return nil
	}
	return sr.Contained
}

// GetDeliverFrom returns the value of the field DeliverFrom.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetDeliverFrom() *fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.DeliverFrom
}

// GetDeliverTo returns the value of the field DeliverTo.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetDeliverTo() *fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.DeliverTo
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetExtension() []*fhir.Extension {
	if sr == nil {
		return nil
	}
	return sr.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetID() string {
	if sr == nil {
		return ""
	}
	return sr.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetIdentifier() []*fhir.Identifier {
	if sr == nil {
		return nil
	}
	return sr.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetImplicitRules() *fhir.URI {
	if sr == nil {
		return nil
	}
	return sr.ImplicitRules
}

// GetItem returns the value of the field Item.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetItem() fhir.Element {
	if sr == nil {
		return nil
	}
	return sr.Item
}

// GetItemCodeableConcept returns the value of the field Item.
func (sr *SupplyRequest) GetItemCodeableConcept() *fhir.CodeableConcept {
	if sr == nil {
		return nil
	}
	val, ok := sr.Item.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetItemReference returns the value of the field Item.
func (sr *SupplyRequest) GetItemReference() *fhir.Reference {
	if sr == nil {
		return nil
	}
	val, ok := sr.Item.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
} // GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetLanguage() *fhir.Code {
	if sr == nil {
		return nil
	}
	return sr.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetMeta() *fhir.Meta {
	if sr == nil {
		return nil
	}
	return sr.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetModifierExtension() []*fhir.Extension {
	if sr == nil {
		return nil
	}
	return sr.ModifierExtension
}

// GetOccurrence returns the value of the field Occurrence.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetOccurrence() fhir.Element {
	if sr == nil {
		return nil
	}
	return sr.Occurrence
}

// GetOccurrenceDateTime returns the value of the field Occurrence.
func (sr *SupplyRequest) GetOccurrenceDateTime() *fhir.DateTime {
	if sr == nil {
		return nil
	}
	val, ok := sr.Occurrence.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
}

// GetOccurrencePeriod returns the value of the field Occurrence.
func (sr *SupplyRequest) GetOccurrencePeriod() *fhir.Period {
	if sr == nil {
		return nil
	}
	val, ok := sr.Occurrence.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
}

// GetOccurrenceTiming returns the value of the field Occurrence.
func (sr *SupplyRequest) GetOccurrenceTiming() *fhir.Timing {
	if sr == nil {
		return nil
	}
	val, ok := sr.Occurrence.(*fhir.Timing)
	if !ok {
		return nil
	}
	return val
} // GetParameter returns the value of the field Parameter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetParameter() []*SupplyRequestParameter {
	if sr == nil {
		return nil
	}
	return sr.Parameter
}

// GetPriority returns the value of the field Priority.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetPriority() *fhir.Code {
	if sr == nil {
		return nil
	}
	return sr.Priority
}

// GetQuantity returns the value of the field Quantity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetQuantity() *fhir.Quantity {
	if sr == nil {
		return nil
	}
	return sr.Quantity
}

// GetReasonCode returns the value of the field ReasonCode.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetReasonCode() []*fhir.CodeableConcept {
	if sr == nil {
		return nil
	}
	return sr.ReasonCode
}

// GetReasonReference returns the value of the field ReasonReference.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetReasonReference() []*fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.ReasonReference
}

// GetRequester returns the value of the field Requester.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetRequester() *fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.Requester
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetStatus() *fhir.Code {
	if sr == nil {
		return nil
	}
	return sr.Status
}

// GetSupplier returns the value of the field Supplier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetSupplier() []*fhir.Reference {
	if sr == nil {
		return nil
	}
	return sr.Supplier
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sr *SupplyRequest) GetText() *fhir.Narrative {
	if sr == nil {
		return nil
	}
	return sr.Text
}

// Ordered item details// Specific parameters for the ordered item. For example, the size of the
// indicated item.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-SupplyRequest.json
type SupplyRequestParameter struct {

	// A code or string that identifies the device detail being asserted.
	Code *fhir.CodeableConcept `fhirpath:"code"`

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

	// The value of the device detail.
	Value fhir.Element `fhirpath:"value"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (srp *SupplyRequestParameter) GetCode() *fhir.CodeableConcept {
	if srp == nil {
		return nil
	}
	return srp.Code
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (srp *SupplyRequestParameter) GetExtension() []*fhir.Extension {
	if srp == nil {
		return nil
	}
	return srp.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (srp *SupplyRequestParameter) GetID() string {
	if srp == nil {
		return ""
	}
	return srp.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (srp *SupplyRequestParameter) GetModifierExtension() []*fhir.Extension {
	if srp == nil {
		return nil
	}
	return srp.ModifierExtension
}

// GetValue returns the value of the field Value.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (srp *SupplyRequestParameter) GetValue() fhir.Element {
	if srp == nil {
		return nil
	}
	return srp.Value
}

// GetValueCodeableConcept returns the value of the field Value.
func (srp *SupplyRequestParameter) GetValueCodeableConcept() *fhir.CodeableConcept {
	if srp == nil {
		return nil
	}
	val, ok := srp.Value.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetValueQuantity returns the value of the field Value.
func (srp *SupplyRequestParameter) GetValueQuantity() *fhir.Quantity {
	if srp == nil {
		return nil
	}
	val, ok := srp.Value.(*fhir.Quantity)
	if !ok {
		return nil
	}
	return val
}

// GetValueRange returns the value of the field Value.
func (srp *SupplyRequestParameter) GetValueRange() *fhir.Range {
	if srp == nil {
		return nil
	}
	val, ok := srp.Value.(*fhir.Range)
	if !ok {
		return nil
	}
	return val
}

// GetValueBoolean returns the value of the field Value.
func (srp *SupplyRequestParameter) GetValueBoolean() *fhir.Boolean {
	if srp == nil {
		return nil
	}
	val, ok := srp.Value.(*fhir.Boolean)
	if !ok {
		return nil
	}
	return val
}
