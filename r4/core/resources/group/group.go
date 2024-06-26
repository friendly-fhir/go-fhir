// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package group

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Represents a defined collection of entities that may be discussed or acted
// upon collectively but which are not expected to act collectively, and are
// not formally or legally recognized; i.e. a collection of entities that isn't
// an Organization.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Group
//   - Source File: StructureDefinition-Group.json
type Group struct {

	// Indicates whether the record for the group is available for use or is merely
	// being retained for historical purposes.
	Active *fhir.Boolean `fhirpath:"active"`

	// If true, indicates that the resource refers to a specific group of real
	// individuals. If false, the group defines a set of intended individuals.
	Actual *fhir.Boolean `fhirpath:"actual"`

	// Identifies traits whose presence r absence is shared by members of the
	// group.
	Characteristic []*GroupCharacteristic `fhirpath:"characteristic"`

	// Provides a specific type of resource the group includes; e.g. "cow",
	// "syringe", etc.
	Code *fhir.CodeableConcept `fhirpath:"code"`

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

	// A unique business identifier for this group.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// Entity responsible for defining and maintaining Group characteristics and/or
	// registered members.
	ManagingEntity *fhir.Reference `fhirpath:"managingEntity"`

	// Identifies the resource instances that are members of the group.
	Member []*GroupMember `fhirpath:"member"`

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

	// A label assigned to the group for human identification and communication.
	Name *fhir.String `fhirpath:"name"`

	// A count of the number of resource instances that are part of the group.
	Quantity *fhir.UnsignedInt `fhirpath:"quantity"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// Identifies the broad classification of the kind of resources the group
	// includes.
	Type *fhir.Code `fhirpath:"type"`

	profileimpl.BaseGroup
	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetActive returns the value of the field Active.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetActive() *fhir.Boolean {
	if g == nil {
		return nil
	}
	return g.Active
}

// GetActual returns the value of the field Actual.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetActual() *fhir.Boolean {
	if g == nil {
		return nil
	}
	return g.Actual
}

// GetCharacteristic returns the value of the field Characteristic.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetCharacteristic() []*GroupCharacteristic {
	if g == nil {
		return nil
	}
	return g.Characteristic
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetCode() *fhir.CodeableConcept {
	if g == nil {
		return nil
	}
	return g.Code
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetContained() []fhir.Resource {
	if g == nil {
		return nil
	}
	return g.Contained
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetExtension() []*fhir.Extension {
	if g == nil {
		return nil
	}
	return g.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetID() string {
	if g == nil {
		return ""
	}
	return g.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetIdentifier() []*fhir.Identifier {
	if g == nil {
		return nil
	}
	return g.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetImplicitRules() *fhir.URI {
	if g == nil {
		return nil
	}
	return g.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetLanguage() *fhir.Code {
	if g == nil {
		return nil
	}
	return g.Language
}

// GetManagingEntity returns the value of the field ManagingEntity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetManagingEntity() *fhir.Reference {
	if g == nil {
		return nil
	}
	return g.ManagingEntity
}

// GetMember returns the value of the field Member.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetMember() []*GroupMember {
	if g == nil {
		return nil
	}
	return g.Member
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetMeta() *fhir.Meta {
	if g == nil {
		return nil
	}
	return g.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetModifierExtension() []*fhir.Extension {
	if g == nil {
		return nil
	}
	return g.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetName() *fhir.String {
	if g == nil {
		return nil
	}
	return g.Name
}

// GetQuantity returns the value of the field Quantity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetQuantity() *fhir.UnsignedInt {
	if g == nil {
		return nil
	}
	return g.Quantity
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetText() *fhir.Narrative {
	if g == nil {
		return nil
	}
	return g.Text
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (g *Group) GetType() *fhir.Code {
	if g == nil {
		return nil
	}
	return g.Type
}

// Include / Exclude group members by Trait// Identifies traits whose presence r absence is shared by members of the
// group.// All the identified characteristics must be true for an entity to a member of
// the group.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Group.json
type GroupCharacteristic struct {

	// A code that identifies the kind of trait being asserted.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// If true, indicates the characteristic is one that is NOT held by members of
	// the group.
	Exclude *fhir.Boolean `fhirpath:"exclude"`

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

	// The period over which the characteristic is tested; e.g. the patient had an
	// operation during the month of June.
	Period *fhir.Period `fhirpath:"period"`

	// The value of the trait that holds (or does not hold - see 'exclude') for
	// members of the group.
	Value fhir.Element `fhirpath:"value"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gc *GroupCharacteristic) GetCode() *fhir.CodeableConcept {
	if gc == nil {
		return nil
	}
	return gc.Code
}

// GetExclude returns the value of the field Exclude.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gc *GroupCharacteristic) GetExclude() *fhir.Boolean {
	if gc == nil {
		return nil
	}
	return gc.Exclude
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gc *GroupCharacteristic) GetExtension() []*fhir.Extension {
	if gc == nil {
		return nil
	}
	return gc.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gc *GroupCharacteristic) GetID() string {
	if gc == nil {
		return ""
	}
	return gc.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gc *GroupCharacteristic) GetModifierExtension() []*fhir.Extension {
	if gc == nil {
		return nil
	}
	return gc.ModifierExtension
}

// GetPeriod returns the value of the field Period.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gc *GroupCharacteristic) GetPeriod() *fhir.Period {
	if gc == nil {
		return nil
	}
	return gc.Period
}

// GetValue returns the value of the field Value.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gc *GroupCharacteristic) GetValue() fhir.Element {
	if gc == nil {
		return nil
	}
	return gc.Value
}

// GetValueCodeableConcept returns the value of the field Value.
func (gc *GroupCharacteristic) GetValueCodeableConcept() *fhir.CodeableConcept {
	if gc == nil {
		return nil
	}
	val, ok := gc.Value.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetValueBoolean returns the value of the field Value.
func (gc *GroupCharacteristic) GetValueBoolean() *fhir.Boolean {
	if gc == nil {
		return nil
	}
	val, ok := gc.Value.(*fhir.Boolean)
	if !ok {
		return nil
	}
	return val
}

// GetValueQuantity returns the value of the field Value.
func (gc *GroupCharacteristic) GetValueQuantity() *fhir.Quantity {
	if gc == nil {
		return nil
	}
	val, ok := gc.Value.(*fhir.Quantity)
	if !ok {
		return nil
	}
	return val
}

// GetValueRange returns the value of the field Value.
func (gc *GroupCharacteristic) GetValueRange() *fhir.Range {
	if gc == nil {
		return nil
	}
	val, ok := gc.Value.(*fhir.Range)
	if !ok {
		return nil
	}
	return val
}

// GetValueReference returns the value of the field Value.
func (gc *GroupCharacteristic) GetValueReference() *fhir.Reference {
	if gc == nil {
		return nil
	}
	val, ok := gc.Value.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
}

// Who or what is in group// Identifies the resource instances that are members of the group.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Group.json
type GroupMember struct {

	// A reference to the entity that is a member of the group. Must be consistent
	// with Group.type. If the entity is another group, then the type must be the
	// same.
	Entity *fhir.Reference `fhirpath:"entity"`

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

	// A flag to indicate that the member is no longer in the group, but previously
	// may have been a member.
	Inactive *fhir.Boolean `fhirpath:"inactive"`

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

	// The period that the member was in the group, if known.
	Period *fhir.Period `fhirpath:"period"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetEntity returns the value of the field Entity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gm *GroupMember) GetEntity() *fhir.Reference {
	if gm == nil {
		return nil
	}
	return gm.Entity
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gm *GroupMember) GetExtension() []*fhir.Extension {
	if gm == nil {
		return nil
	}
	return gm.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gm *GroupMember) GetID() string {
	if gm == nil {
		return ""
	}
	return gm.ID
}

// GetInactive returns the value of the field Inactive.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gm *GroupMember) GetInactive() *fhir.Boolean {
	if gm == nil {
		return nil
	}
	return gm.Inactive
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gm *GroupMember) GetModifierExtension() []*fhir.Extension {
	if gm == nil {
		return nil
	}
	return gm.ModifierExtension
}

// GetPeriod returns the value of the field Period.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (gm *GroupMember) GetPeriod() *fhir.Period {
	if gm == nil {
		return nil
	}
	return gm.Period
}

func (g *Group) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (g *Group) UnmarshalJSON(data []byte) error {
	var raw struct {
		Active         *fhir.Boolean          `json:"active"`
		Actual         *fhir.Boolean          `json:"actual"`
		Characteristic []*GroupCharacteristic `json:"characteristic"`
		Code           *fhir.CodeableConcept  `json:"code"`
		Contained      []fhir.Resource        `json:"contained"`
		Extension      []*fhir.Extension      `json:"extension"`

		ID                string             `json:"id"`
		Identifier        []*fhir.Identifier `json:"identifier"`
		ImplicitRules     *fhir.URI          `json:"implicitRules"`
		Language          *fhir.Code         `json:"language"`
		ManagingEntity    *fhir.Reference    `json:"managingEntity"`
		Member            []*GroupMember     `json:"member"`
		Meta              *fhir.Meta         `json:"meta"`
		ModifierExtension []*fhir.Extension  `json:"modifierExtension"`
		Name              *fhir.String       `json:"name"`
		Quantity          *fhir.UnsignedInt  `json:"quantity"`
		Text              *fhir.Narrative    `json:"text"`
		Type              *fhir.Code         `json:"type"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	g.Active = raw.Active
	g.Actual = raw.Actual
	g.Characteristic = raw.Characteristic
	g.Code = raw.Code
	g.Contained = raw.Contained
	g.Extension = raw.Extension
	g.ID = raw.ID
	g.Identifier = raw.Identifier
	g.ImplicitRules = raw.ImplicitRules
	g.Language = raw.Language
	g.ManagingEntity = raw.ManagingEntity
	g.Member = raw.Member
	g.Meta = raw.Meta
	g.ModifierExtension = raw.ModifierExtension
	g.Name = raw.Name
	g.Quantity = raw.Quantity
	g.Text = raw.Text
	g.Type = raw.Type
	return nil
}

var _ json.Marshaler = (*Group)(nil)
var _ json.Unmarshaler = (*Group)(nil)

func (gc *GroupCharacteristic) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (gc *GroupCharacteristic) UnmarshalJSON(data []byte) error {
	var raw struct {
		Code      *fhir.CodeableConcept `json:"code"`
		Exclude   *fhir.Boolean         `json:"exclude"`
		Extension []*fhir.Extension     `json:"extension"`

		ID                   string                `json:"id"`
		ModifierExtension    []*fhir.Extension     `json:"modifierExtension"`
		Period               *fhir.Period          `json:"period"`
		ValueCodeableConcept *fhir.CodeableConcept `json:"valueCodeableConcept"`
		ValueBoolean         *fhir.Boolean         `json:"valueBoolean"`
		ValueQuantity        *fhir.Quantity        `json:"valueQuantity"`
		ValueRange           *fhir.Range           `json:"valueRange"`
		ValueReference       *fhir.Reference       `json:"valueReference"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	gc.Code = raw.Code
	gc.Exclude = raw.Exclude
	gc.Extension = raw.Extension
	gc.ID = raw.ID
	gc.ModifierExtension = raw.ModifierExtension
	gc.Period = raw.Period
	gc.Value, err = validate.SelectOneOf[fhir.Element]("Group.characteristic.value",
		raw.ValueCodeableConcept,
		raw.ValueBoolean,
		raw.ValueQuantity,
		raw.ValueRange,
		raw.ValueReference)
	if err != nil {
		return err
	}
	return nil
}

var _ json.Marshaler = (*GroupCharacteristic)(nil)
var _ json.Unmarshaler = (*GroupCharacteristic)(nil)

func (gm *GroupMember) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (gm *GroupMember) UnmarshalJSON(data []byte) error {
	var raw struct {
		Entity    *fhir.Reference   `json:"entity"`
		Extension []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		Inactive          *fhir.Boolean     `json:"inactive"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		Period            *fhir.Period      `json:"period"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	gm.Entity = raw.Entity
	gm.Extension = raw.Extension
	gm.ID = raw.ID
	gm.Inactive = raw.Inactive
	gm.ModifierExtension = raw.ModifierExtension
	gm.Period = raw.Period
	return nil
}

var _ json.Marshaler = (*GroupMember)(nil)
var _ json.Unmarshaler = (*GroupMember)(nil)
