// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Base StructureDefinition for RelatedArtifact Type: Related artifacts such as
// additional documentation, justification, or bibliographic references.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/RelatedArtifact
//   - Source File: StructureDefinition-RelatedArtifact.json
type RelatedArtifact struct {

	// A bibliographic citation for the related artifact. This text SHOULD be
	// formatted according to an accepted citation format.
	Citation *Markdown `fhirpath:"citation"`

	// A brief description of the document or knowledge resource being referenced,
	// suitable for display to a consumer.
	Display *String `fhirpath:"display"`

	// The document being referenced, represented as an attachment. This is
	// exclusive with the resource element.
	Document *Attachment `fhirpath:"document"`

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

	// A short label that can be used to reference the citation from elsewhere in
	// the containing artifact, such as a footnote index.
	Label *String `fhirpath:"label"`

	// The related resource, such as a library, value set, profile, or other
	// knowledge resource.
	Resource *Canonical `fhirpath:"resource"`

	// The type of relationship to the related artifact.
	Type *Code `fhirpath:"type"`

	// A url for the artifact that can be followed to access the actual content.
	URL *URL `fhirpath:"url"`

	profileimpl.BaseElement
}

// GetCitation returns the value of the field Citation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ra *RelatedArtifact) GetCitation() *Markdown {
	if ra == nil {
		return nil
	}
	return ra.Citation
}

// GetDisplay returns the value of the field Display.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ra *RelatedArtifact) GetDisplay() *String {
	if ra == nil {
		return nil
	}
	return ra.Display
}

// GetDocument returns the value of the field Document.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ra *RelatedArtifact) GetDocument() *Attachment {
	if ra == nil {
		return nil
	}
	return ra.Document
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ra *RelatedArtifact) GetExtension() []*Extension {
	if ra == nil {
		return nil
	}
	return ra.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ra *RelatedArtifact) GetID() string {
	if ra == nil {
		return ""
	}
	return ra.ID
}

// GetLabel returns the value of the field Label.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ra *RelatedArtifact) GetLabel() *String {
	if ra == nil {
		return nil
	}
	return ra.Label
}

// GetResource returns the value of the field Resource.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ra *RelatedArtifact) GetResource() *Canonical {
	if ra == nil {
		return nil
	}
	return ra.Resource
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ra *RelatedArtifact) GetType() *Code {
	if ra == nil {
		return nil
	}
	return ra.Type
}

// GetURL returns the value of the field URL.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ra *RelatedArtifact) GetURL() *URL {
	if ra == nil {
		return nil
	}
	return ra.URL
}

func (ra *RelatedArtifact) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (ra *RelatedArtifact) UnmarshalJSON(data []byte) error {
	var raw struct {
		Citation  *Markdown    `json:"citation"`
		Display   *String      `json:"display"`
		Document  *Attachment  `json:"document"`
		Extension []*Extension `json:"extension"`

		ID       string     `json:"id"`
		Label    *String    `json:"label"`
		Resource *Canonical `json:"resource"`
		Type     *Code      `json:"type"`
		URL      *URL       `json:"url"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	ra.Citation = raw.Citation
	ra.Display = raw.Display
	ra.Document = raw.Document
	ra.Extension = raw.Extension
	ra.ID = raw.ID
	ra.Label = raw.Label
	ra.Resource = raw.Resource
	ra.Type = raw.Type
	ra.URL = raw.URL
	return nil
}

var _ json.Marshaler = (*RelatedArtifact)(nil)
var _ json.Unmarshaler = (*RelatedArtifact)(nil)
