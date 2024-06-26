// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Base StructureDefinition for Attachment Type: For referring to data content
// defined in other formats.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Attachment
//   - Source File: StructureDefinition-Attachment.json
type Attachment struct {

	// Identifies the type of the data in the attachment and allows a method to be
	// chosen to interpret or render the data. Includes mime type parameters such
	// as charset where appropriate.
	ContentType *Code `fhirpath:"contentType"`

	// The date that the attachment was first created.
	Creation *DateTime `fhirpath:"creation"`

	// The actual data of the attachment - a sequence of bytes, base64 encoded.
	Data *Base64Binary `fhirpath:"data"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*Extension `fhirpath:"extension"`

	// The calculated hash of the data using SHA-1. Represented using base64.
	Hash *Base64Binary `fhirpath:"hash"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID string `fhirpath:"id"`

	// The human language of the content. The value can be any valid value
	// according to BCP 47.
	Language *Code `fhirpath:"language"`

	// The number of bytes of data that make up this attachment (before base64
	// encoding, if that is done).
	Size *UnsignedInt `fhirpath:"size"`

	// A label or set of text to display in place of the data.
	Title *String `fhirpath:"title"`

	// A location where the data can be accessed.
	URL *URL `fhirpath:"url"`

	profileimpl.BaseElement
}

// GetContentType returns the value of the field ContentType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Attachment) GetContentType() *Code {
	if a == nil {
		return nil
	}
	return a.ContentType
}

// GetCreation returns the value of the field Creation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Attachment) GetCreation() *DateTime {
	if a == nil {
		return nil
	}
	return a.Creation
}

// GetData returns the value of the field Data.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Attachment) GetData() *Base64Binary {
	if a == nil {
		return nil
	}
	return a.Data
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Attachment) GetExtension() []*Extension {
	if a == nil {
		return nil
	}
	return a.Extension
}

// GetHash returns the value of the field Hash.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Attachment) GetHash() *Base64Binary {
	if a == nil {
		return nil
	}
	return a.Hash
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Attachment) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Attachment) GetLanguage() *Code {
	if a == nil {
		return nil
	}
	return a.Language
}

// GetSize returns the value of the field Size.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Attachment) GetSize() *UnsignedInt {
	if a == nil {
		return nil
	}
	return a.Size
}

// GetTitle returns the value of the field Title.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Attachment) GetTitle() *String {
	if a == nil {
		return nil
	}
	return a.Title
}

// GetURL returns the value of the field URL.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (a *Attachment) GetURL() *URL {
	if a == nil {
		return nil
	}
	return a.URL
}

func (a *Attachment) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (a *Attachment) UnmarshalJSON(data []byte) error {
	var raw struct {
		ContentType *Code         `json:"contentType"`
		Creation    *DateTime     `json:"creation"`
		Data        *Base64Binary `json:"data"`
		Extension   []*Extension  `json:"extension"`
		Hash        *Base64Binary `json:"hash"`

		ID       string       `json:"id"`
		Language *Code        `json:"language"`
		Size     *UnsignedInt `json:"size"`
		Title    *String      `json:"title"`
		URL      *URL         `json:"url"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	a.ContentType = raw.ContentType
	a.Creation = raw.Creation
	a.Data = raw.Data
	a.Extension = raw.Extension
	a.Hash = raw.Hash
	a.ID = raw.ID
	a.Language = raw.Language
	a.Size = raw.Size
	a.Title = raw.Title
	a.URL = raw.URL
	return nil
}

var _ json.Marshaler = (*Attachment)(nil)
var _ json.Unmarshaler = (*Attachment)(nil)
