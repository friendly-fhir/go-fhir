// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package documentmanifest

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// A collection of documents compiled for a purpose together with metadata that
// applies to the collection.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/DocumentManifest
//   - Source File: StructureDefinition-DocumentManifest.json
type DocumentManifest struct {

	// Identifies who is the author of the manifest. Manifest author is not
	// necessarly the author of the references included.
	Author []*fhir.Reference `fhirpath:"author"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The list of Resources that consist of the parts of this manifest.
	Content []*fhir.Reference `fhirpath:"content"`

	// When the document manifest was created for submission to the server (not
	// necessarily the same thing as the actual resource last modified time, since
	// it may be modified, replicated, etc.).
	Created *fhir.DateTime `fhirpath:"created"`

	// Human-readable description of the source document. This is sometimes known
	// as the "title".
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

	// Other identifiers associated with the document manifest, including version
	// independent identifiers.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// A single identifier that uniquely identifies this manifest. Principally used
	// to refer to the manifest in non-FHIR contexts.
	MasterIdentifier *fhir.Identifier `fhirpath:"masterIdentifier"`

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

	// A patient, practitioner, or organization for which this set of documents is
	// intended.
	Recipient []*fhir.Reference `fhirpath:"recipient"`

	// Related identifiers or resources associated with the DocumentManifest.
	Related []*DocumentManifestRelated `fhirpath:"related"`

	// Identifies the source system, application, or software that produced the
	// document manifest.
	Source *fhir.URI `fhirpath:"source"`

	// The status of this document manifest.
	Status *fhir.Code `fhirpath:"status"`

	// Who or what the set of documents is about. The documents can be about a
	// person, (patient or healthcare practitioner), a device (i.e. machine) or
	// even a group of subjects (such as a document about a herd of farm animals,
	// or a set of patients that share a common exposure). If the documents cross
	// more than one subject, then more than one subject is allowed here (unusual
	// use case).
	Subject *fhir.Reference `fhirpath:"subject"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// The code specifying the type of clinical activity that resulted in placing
	// the associated content into the DocumentManifest.
	Type *fhir.CodeableConcept `fhirpath:"type"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetAuthor returns the value of the field Author.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetAuthor() []*fhir.Reference {
	if dm == nil {
		return nil
	}
	return dm.Author
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetContained() []fhir.Resource {
	if dm == nil {
		return nil
	}
	return dm.Contained
}

// GetContent returns the value of the field Content.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetContent() []*fhir.Reference {
	if dm == nil {
		return nil
	}
	return dm.Content
}

// GetCreated returns the value of the field Created.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetCreated() *fhir.DateTime {
	if dm == nil {
		return nil
	}
	return dm.Created
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetDescription() *fhir.String {
	if dm == nil {
		return nil
	}
	return dm.Description
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetExtension() []*fhir.Extension {
	if dm == nil {
		return nil
	}
	return dm.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetID() string {
	if dm == nil {
		return ""
	}
	return dm.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetIdentifier() []*fhir.Identifier {
	if dm == nil {
		return nil
	}
	return dm.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetImplicitRules() *fhir.URI {
	if dm == nil {
		return nil
	}
	return dm.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetLanguage() *fhir.Code {
	if dm == nil {
		return nil
	}
	return dm.Language
}

// GetMasterIdentifier returns the value of the field MasterIdentifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetMasterIdentifier() *fhir.Identifier {
	if dm == nil {
		return nil
	}
	return dm.MasterIdentifier
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetMeta() *fhir.Meta {
	if dm == nil {
		return nil
	}
	return dm.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetModifierExtension() []*fhir.Extension {
	if dm == nil {
		return nil
	}
	return dm.ModifierExtension
}

// GetRecipient returns the value of the field Recipient.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetRecipient() []*fhir.Reference {
	if dm == nil {
		return nil
	}
	return dm.Recipient
}

// GetRelated returns the value of the field Related.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetRelated() []*DocumentManifestRelated {
	if dm == nil {
		return nil
	}
	return dm.Related
}

// GetSource returns the value of the field Source.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetSource() *fhir.URI {
	if dm == nil {
		return nil
	}
	return dm.Source
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetStatus() *fhir.Code {
	if dm == nil {
		return nil
	}
	return dm.Status
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetSubject() *fhir.Reference {
	if dm == nil {
		return nil
	}
	return dm.Subject
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetText() *fhir.Narrative {
	if dm == nil {
		return nil
	}
	return dm.Text
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dm *DocumentManifest) GetType() *fhir.CodeableConcept {
	if dm == nil {
		return nil
	}
	return dm.Type
}

// Related things// Related identifiers or resources associated with the DocumentManifest.// May be identifiers or resources that caused the DocumentManifest to be
// created.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-DocumentManifest.json
type DocumentManifestRelated struct {

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

	// Related identifier to this DocumentManifest. For example, Order numbers,
	// accession numbers, XDW workflow numbers.
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

	// Related Resource to this DocumentManifest. For example, Order,
	// ServiceRequest, Procedure, EligibilityRequest, etc.
	Ref *fhir.Reference `fhirpath:"ref"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dmr *DocumentManifestRelated) GetExtension() []*fhir.Extension {
	if dmr == nil {
		return nil
	}
	return dmr.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dmr *DocumentManifestRelated) GetID() string {
	if dmr == nil {
		return ""
	}
	return dmr.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dmr *DocumentManifestRelated) GetIdentifier() *fhir.Identifier {
	if dmr == nil {
		return nil
	}
	return dmr.Identifier
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dmr *DocumentManifestRelated) GetModifierExtension() []*fhir.Extension {
	if dmr == nil {
		return nil
	}
	return dmr.ModifierExtension
}

// GetRef returns the value of the field Ref.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dmr *DocumentManifestRelated) GetRef() *fhir.Reference {
	if dmr == nil {
		return nil
	}
	return dmr.Ref
}

func (dm *DocumentManifest) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (dm *DocumentManifest) UnmarshalJSON(data []byte) error {
	var raw struct {
		Author      []*fhir.Reference `json:"author"`
		Contained   []fhir.Resource   `json:"contained"`
		Content     []*fhir.Reference `json:"content"`
		Created     *fhir.DateTime    `json:"created"`
		Description *fhir.String      `json:"description"`
		Extension   []*fhir.Extension `json:"extension"`

		ID                string                     `json:"id"`
		Identifier        []*fhir.Identifier         `json:"identifier"`
		ImplicitRules     *fhir.URI                  `json:"implicitRules"`
		Language          *fhir.Code                 `json:"language"`
		MasterIdentifier  *fhir.Identifier           `json:"masterIdentifier"`
		Meta              *fhir.Meta                 `json:"meta"`
		ModifierExtension []*fhir.Extension          `json:"modifierExtension"`
		Recipient         []*fhir.Reference          `json:"recipient"`
		Related           []*DocumentManifestRelated `json:"related"`
		Source            *fhir.URI                  `json:"source"`
		Status            *fhir.Code                 `json:"status"`
		Subject           *fhir.Reference            `json:"subject"`
		Text              *fhir.Narrative            `json:"text"`
		Type              *fhir.CodeableConcept      `json:"type"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	dm.Author = raw.Author
	dm.Contained = raw.Contained
	dm.Content = raw.Content
	dm.Created = raw.Created
	dm.Description = raw.Description
	dm.Extension = raw.Extension
	dm.ID = raw.ID
	dm.Identifier = raw.Identifier
	dm.ImplicitRules = raw.ImplicitRules
	dm.Language = raw.Language
	dm.MasterIdentifier = raw.MasterIdentifier
	dm.Meta = raw.Meta
	dm.ModifierExtension = raw.ModifierExtension
	dm.Recipient = raw.Recipient
	dm.Related = raw.Related
	dm.Source = raw.Source
	dm.Status = raw.Status
	dm.Subject = raw.Subject
	dm.Text = raw.Text
	dm.Type = raw.Type
	return nil
}

var _ json.Marshaler = (*DocumentManifest)(nil)
var _ json.Unmarshaler = (*DocumentManifest)(nil)

func (dmr *DocumentManifestRelated) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (dmr *DocumentManifestRelated) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		Identifier        *fhir.Identifier  `json:"identifier"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		Ref               *fhir.Reference   `json:"ref"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	dmr.Extension = raw.Extension
	dmr.ID = raw.ID
	dmr.Identifier = raw.Identifier
	dmr.ModifierExtension = raw.ModifierExtension
	dmr.Ref = raw.Ref
	return nil
}

var _ json.Marshaler = (*DocumentManifestRelated)(nil)
var _ json.Unmarshaler = (*DocumentManifestRelated)(nil)
