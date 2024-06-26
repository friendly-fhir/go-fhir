// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// This is the base resource type for everything.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Resource
//   - Source File: StructureDefinition-Resource.json
type Resource interface {
	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	GetID() string

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	GetImplicitRules() *URI

	// The base language in which the resource is written.
	GetLanguage() *Code

	// The metadata about the resource. This is content that is maintained by the
	// infrastructure. Changes to the content might not always be associated with
	// version changes to the resource.
	GetMeta() *Meta

	// Implement base ResourceProfile
	profileimpl.ResourceProfile
}
