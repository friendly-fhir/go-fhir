// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package list

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// A list is a curated collection of resources.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/List
//   - Source File: StructureDefinition-List.json
type List struct {

	// This code defines the purpose of the list - why it was created.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The date that the list was prepared.
	Date *fhir.DateTime `fhirpath:"date"`

	// If the list is empty, why the list is empty.
	EmptyReason *fhir.CodeableConcept `fhirpath:"emptyReason"`

	// The encounter that is the context in which this list was created.
	Encounter *fhir.Reference `fhirpath:"encounter"`

	// Entries in this list.
	Entry []*ListEntry `fhirpath:"entry"`

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

	// Identifier for the List assigned for business purposes outside the context
	// of FHIR.
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

	// How this list was prepared - whether it is a working list that is suitable
	// for being maintained on an ongoing basis, or if it represents a snapshot of
	// a list of items from another source, or whether it is a prepared list where
	// items may be marked as added, modified or deleted.
	Mode *fhir.Code `fhirpath:"mode"`

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

	// Comments that apply to the overall list.
	Note []*fhir.Annotation `fhirpath:"note"`

	// What order applies to the items in the list.
	OrderedBy *fhir.CodeableConcept `fhirpath:"orderedBy"`

	// The entity responsible for deciding what the contents of the list were.
	// Where the list was created by a human, this is the same as the author of the
	// list.
	Source *fhir.Reference `fhirpath:"source"`

	// Indicates the current state of this list.
	Status *fhir.Code `fhirpath:"status"`

	// The common subject (or patient) of the resources that are in the list if
	// there is one.
	Subject *fhir.Reference `fhirpath:"subject"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// A label for the list assigned by the author.
	Title *fhir.String `fhirpath:"title"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetCode() *fhir.CodeableConcept {
	if l == nil {
		return nil
	}
	return l.Code
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetContained() []fhir.Resource {
	if l == nil {
		return nil
	}
	return l.Contained
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetDate() *fhir.DateTime {
	if l == nil {
		return nil
	}
	return l.Date
}

// GetEmptyReason returns the value of the field EmptyReason.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetEmptyReason() *fhir.CodeableConcept {
	if l == nil {
		return nil
	}
	return l.EmptyReason
}

// GetEncounter returns the value of the field Encounter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetEncounter() *fhir.Reference {
	if l == nil {
		return nil
	}
	return l.Encounter
}

// GetEntry returns the value of the field Entry.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetEntry() []*ListEntry {
	if l == nil {
		return nil
	}
	return l.Entry
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetExtension() []*fhir.Extension {
	if l == nil {
		return nil
	}
	return l.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetID() string {
	if l == nil {
		return ""
	}
	return l.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetIdentifier() []*fhir.Identifier {
	if l == nil {
		return nil
	}
	return l.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetImplicitRules() *fhir.URI {
	if l == nil {
		return nil
	}
	return l.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetLanguage() *fhir.Code {
	if l == nil {
		return nil
	}
	return l.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetMeta() *fhir.Meta {
	if l == nil {
		return nil
	}
	return l.Meta
}

// GetMode returns the value of the field Mode.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetMode() *fhir.Code {
	if l == nil {
		return nil
	}
	return l.Mode
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetModifierExtension() []*fhir.Extension {
	if l == nil {
		return nil
	}
	return l.ModifierExtension
}

// GetNote returns the value of the field Note.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetNote() []*fhir.Annotation {
	if l == nil {
		return nil
	}
	return l.Note
}

// GetOrderedBy returns the value of the field OrderedBy.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetOrderedBy() *fhir.CodeableConcept {
	if l == nil {
		return nil
	}
	return l.OrderedBy
}

// GetSource returns the value of the field Source.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetSource() *fhir.Reference {
	if l == nil {
		return nil
	}
	return l.Source
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetStatus() *fhir.Code {
	if l == nil {
		return nil
	}
	return l.Status
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetSubject() *fhir.Reference {
	if l == nil {
		return nil
	}
	return l.Subject
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetText() *fhir.Narrative {
	if l == nil {
		return nil
	}
	return l.Text
}

// GetTitle returns the value of the field Title.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (l *List) GetTitle() *fhir.String {
	if l == nil {
		return nil
	}
	return l.Title
}

// Entries in the list// Entries in this list.// If there are no entries in the list, an emptyReason SHOULD be provided.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-List.json
type ListEntry struct {

	// When this item was added to the list.
	Date *fhir.DateTime `fhirpath:"date"`

	// True if this item is marked as deleted in the list.
	Deleted *fhir.Boolean `fhirpath:"deleted"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// The flag allows the system constructing the list to indicate the role and
	// significance of the item in the list.
	Flag *fhir.CodeableConcept `fhirpath:"flag"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID string `fhirpath:"id"`

	// A reference to the actual resource from which data was derived.
	Item *fhir.Reference `fhirpath:"item"`

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

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (le *ListEntry) GetDate() *fhir.DateTime {
	if le == nil {
		return nil
	}
	return le.Date
}

// GetDeleted returns the value of the field Deleted.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (le *ListEntry) GetDeleted() *fhir.Boolean {
	if le == nil {
		return nil
	}
	return le.Deleted
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (le *ListEntry) GetExtension() []*fhir.Extension {
	if le == nil {
		return nil
	}
	return le.Extension
}

// GetFlag returns the value of the field Flag.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (le *ListEntry) GetFlag() *fhir.CodeableConcept {
	if le == nil {
		return nil
	}
	return le.Flag
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (le *ListEntry) GetID() string {
	if le == nil {
		return ""
	}
	return le.ID
}

// GetItem returns the value of the field Item.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (le *ListEntry) GetItem() *fhir.Reference {
	if le == nil {
		return nil
	}
	return le.Item
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (le *ListEntry) GetModifierExtension() []*fhir.Extension {
	if le == nil {
		return nil
	}
	return le.ModifierExtension
}

func (l *List) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (l *List) UnmarshalJSON(data []byte) error {
	var raw struct {
		Code        *fhir.CodeableConcept `json:"code"`
		Contained   []fhir.Resource       `json:"contained"`
		Date        *fhir.DateTime        `json:"date"`
		EmptyReason *fhir.CodeableConcept `json:"emptyReason"`
		Encounter   *fhir.Reference       `json:"encounter"`
		Entry       []*ListEntry          `json:"entry"`
		Extension   []*fhir.Extension     `json:"extension"`

		ID                string                `json:"id"`
		Identifier        []*fhir.Identifier    `json:"identifier"`
		ImplicitRules     *fhir.URI             `json:"implicitRules"`
		Language          *fhir.Code            `json:"language"`
		Meta              *fhir.Meta            `json:"meta"`
		Mode              *fhir.Code            `json:"mode"`
		ModifierExtension []*fhir.Extension     `json:"modifierExtension"`
		Note              []*fhir.Annotation    `json:"note"`
		OrderedBy         *fhir.CodeableConcept `json:"orderedBy"`
		Source            *fhir.Reference       `json:"source"`
		Status            *fhir.Code            `json:"status"`
		Subject           *fhir.Reference       `json:"subject"`
		Text              *fhir.Narrative       `json:"text"`
		Title             *fhir.String          `json:"title"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	l.Code = raw.Code
	l.Contained = raw.Contained
	l.Date = raw.Date
	l.EmptyReason = raw.EmptyReason
	l.Encounter = raw.Encounter
	l.Entry = raw.Entry
	l.Extension = raw.Extension
	l.ID = raw.ID
	l.Identifier = raw.Identifier
	l.ImplicitRules = raw.ImplicitRules
	l.Language = raw.Language
	l.Meta = raw.Meta
	l.Mode = raw.Mode
	l.ModifierExtension = raw.ModifierExtension
	l.Note = raw.Note
	l.OrderedBy = raw.OrderedBy
	l.Source = raw.Source
	l.Status = raw.Status
	l.Subject = raw.Subject
	l.Text = raw.Text
	l.Title = raw.Title
	return nil
}

var _ json.Marshaler = (*List)(nil)
var _ json.Unmarshaler = (*List)(nil)

func (le *ListEntry) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (le *ListEntry) UnmarshalJSON(data []byte) error {
	var raw struct {
		Date      *fhir.DateTime        `json:"date"`
		Deleted   *fhir.Boolean         `json:"deleted"`
		Extension []*fhir.Extension     `json:"extension"`
		Flag      *fhir.CodeableConcept `json:"flag"`

		ID                string            `json:"id"`
		Item              *fhir.Reference   `json:"item"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	le.Date = raw.Date
	le.Deleted = raw.Deleted
	le.Extension = raw.Extension
	le.Flag = raw.Flag
	le.ID = raw.ID
	le.Item = raw.Item
	le.ModifierExtension = raw.ModifierExtension
	return nil
}

var _ json.Marshaler = (*ListEntry)(nil)
var _ json.Unmarshaler = (*ListEntry)(nil)
