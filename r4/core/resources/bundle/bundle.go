// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package bundle

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// A container for a collection of resources.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Bundle
//   - Source File: StructureDefinition-Bundle.json
type Bundle struct {

	// An entry in a bundle resource - will either contain a resource or
	// information about a resource (transactions and history only).
	Entry []*BundleEntry `fhirpath:"entry"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// A persistent identifier for the bundle that won't change as a bundle is
	// copied from server to server.
	Identifier *fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// A series of links that provide context to this bundle.
	Link []*BundleLink `fhirpath:"link"`

	// The metadata about the resource. This is content that is maintained by the
	// infrastructure. Changes to the content might not always be associated with
	// version changes to the resource.
	Meta *fhir.Meta `fhirpath:"meta"`

	// Digital Signature - base64 encoded. XML-DSig or a JWT.
	Signature *fhir.Signature `fhirpath:"signature"`

	// The date/time that the bundle was assembled - i.e. when the resources were
	// placed in the bundle.
	Timestamp *fhir.Instant `fhirpath:"timestamp"`

	// If a set of search matches, this is the total number of entries of type
	// 'match' across all pages in the search. It does not include search.mode =
	// 'include' or 'outcome' entries and it does not provide a count of the number
	// of entries in the Bundle.
	Total *fhir.UnsignedInt `fhirpath:"total"`

	// Indicates the purpose of this bundle - how it is intended to be used.
	Type *fhir.Code `fhirpath:"type"`

	profileimpl.BaseResource
}

// GetEntry returns the value of the field Entry.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b *Bundle) GetEntry() []*BundleEntry {
	if b == nil {
		return nil
	}
	return b.Entry
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b *Bundle) GetID() string {
	if b == nil {
		return ""
	}
	return b.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b *Bundle) GetIdentifier() *fhir.Identifier {
	if b == nil {
		return nil
	}
	return b.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b *Bundle) GetImplicitRules() *fhir.URI {
	if b == nil {
		return nil
	}
	return b.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b *Bundle) GetLanguage() *fhir.Code {
	if b == nil {
		return nil
	}
	return b.Language
}

// GetLink returns the value of the field Link.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b *Bundle) GetLink() []*BundleLink {
	if b == nil {
		return nil
	}
	return b.Link
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b *Bundle) GetMeta() *fhir.Meta {
	if b == nil {
		return nil
	}
	return b.Meta
}

// GetSignature returns the value of the field Signature.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b *Bundle) GetSignature() *fhir.Signature {
	if b == nil {
		return nil
	}
	return b.Signature
}

// GetTimestamp returns the value of the field Timestamp.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b *Bundle) GetTimestamp() *fhir.Instant {
	if b == nil {
		return nil
	}
	return b.Timestamp
}

// GetTotal returns the value of the field Total.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b *Bundle) GetTotal() *fhir.UnsignedInt {
	if b == nil {
		return nil
	}
	return b.Total
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (b *Bundle) GetType() *fhir.Code {
	if b == nil {
		return nil
	}
	return b.Type
}

// Entry in the bundle - will have a resource or information// An entry in a bundle resource - will either contain a resource or
// information about a resource (transactions and history only).
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Bundle.json
type BundleEntry struct {

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// The Absolute URL for the resource. The fullUrl SHALL NOT disagree with the
	// id in the resource - i.e. if the fullUrl is not a urn:uuid, the URL shall be
	// version-independent URL consistent with the Resource.id. The fullUrl is a
	// version independent reference to the resource. The fullUrl element SHALL
	// have a value except that: * fullUrl can be empty on a POST (although it does
	// not need to when specifying a temporary id for reference in the bundle) *
	// Results from operations might involve resources that are not identified.
	FullURL *fhir.URI `fhirpath:"fullUrl"`

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

	// Additional information about how this entry should be processed as part of a
	// transaction or batch. For history, it shows how the entry was processed to
	// create the version contained in the entry.
	Request *BundleEntryRequest `fhirpath:"request"`

	// The Resource for the entry. The purpose/meaning of the resource is
	// determined by the Bundle.type.
	Resource fhir.Resource `fhirpath:"resource"`

	// Indicates the results of processing the corresponding 'request' entry in the
	// batch or transaction being responded to or what the results of an operation
	// where when returning history.
	Response *BundleEntryResponse `fhirpath:"response"`

	// Information about the search process that lead to the creation of this
	// entry.
	Search *BundleEntrySearch `fhirpath:"search"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (be *BundleEntry) GetExtension() []*fhir.Extension {
	if be == nil {
		return nil
	}
	return be.Extension
}

// GetFullURL returns the value of the field FullURL.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (be *BundleEntry) GetFullURL() *fhir.URI {
	if be == nil {
		return nil
	}
	return be.FullURL
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (be *BundleEntry) GetID() string {
	if be == nil {
		return ""
	}
	return be.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (be *BundleEntry) GetModifierExtension() []*fhir.Extension {
	if be == nil {
		return nil
	}
	return be.ModifierExtension
}

// GetRequest returns the value of the field Request.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (be *BundleEntry) GetRequest() *BundleEntryRequest {
	if be == nil {
		return nil
	}
	return be.Request
}

// GetResource returns the value of the field Resource.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (be *BundleEntry) GetResource() fhir.Resource {
	if be == nil {
		return nil
	}
	return be.Resource
}

// GetResponse returns the value of the field Response.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (be *BundleEntry) GetResponse() *BundleEntryResponse {
	if be == nil {
		return nil
	}
	return be.Response
}

// GetSearch returns the value of the field Search.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (be *BundleEntry) GetSearch() *BundleEntrySearch {
	if be == nil {
		return nil
	}
	return be.Search
}

// Additional execution information (transaction/batch/history)// Additional information about how this entry should be processed as part of a
// transaction or batch. For history, it shows how the entry was processed to
// create the version contained in the entry.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Bundle.json
type BundleEntryRequest struct {

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

	// Only perform the operation if the Etag value matches. For more information,
	// see the API section ["Managing Resource Contention"](http.html#concurrency).
	IfMatch *fhir.String `fhirpath:"ifMatch"`

	// Only perform the operation if the last updated date matches. See the API
	// documentation for ["Conditional Read"](http.html#cread).
	IfModifiedSince *fhir.Instant `fhirpath:"ifModifiedSince"`

	// Instruct the server not to perform the create if a specified resource
	// already exists. For further information, see the API documentation for
	// ["Conditional Create"](http.html#ccreate). This is just the query portion of
	// the URL - what follows the "?" (not including the "?").
	IfNoneExist *fhir.String `fhirpath:"ifNoneExist"`

	// If the ETag values match, return a 304 Not Modified status. See the API
	// documentation for ["Conditional Read"](http.html#cread).
	IfNoneMatch *fhir.String `fhirpath:"ifNoneMatch"`

	// In a transaction or batch, this is the HTTP action to be executed for this
	// entry. In a history bundle, this indicates the HTTP action that occurred.
	Method *fhir.Code `fhirpath:"method"`

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

	// The URL for this entry, relative to the root (the address to which the
	// request is posted).
	URL *fhir.URI `fhirpath:"url"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryRequest) GetExtension() []*fhir.Extension {
	if ber == nil {
		return nil
	}
	return ber.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryRequest) GetID() string {
	if ber == nil {
		return ""
	}
	return ber.ID
}

// GetIfMatch returns the value of the field IfMatch.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryRequest) GetIfMatch() *fhir.String {
	if ber == nil {
		return nil
	}
	return ber.IfMatch
}

// GetIfModifiedSince returns the value of the field IfModifiedSince.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryRequest) GetIfModifiedSince() *fhir.Instant {
	if ber == nil {
		return nil
	}
	return ber.IfModifiedSince
}

// GetIfNoneExist returns the value of the field IfNoneExist.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryRequest) GetIfNoneExist() *fhir.String {
	if ber == nil {
		return nil
	}
	return ber.IfNoneExist
}

// GetIfNoneMatch returns the value of the field IfNoneMatch.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryRequest) GetIfNoneMatch() *fhir.String {
	if ber == nil {
		return nil
	}
	return ber.IfNoneMatch
}

// GetMethod returns the value of the field Method.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryRequest) GetMethod() *fhir.Code {
	if ber == nil {
		return nil
	}
	return ber.Method
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryRequest) GetModifierExtension() []*fhir.Extension {
	if ber == nil {
		return nil
	}
	return ber.ModifierExtension
}

// GetURL returns the value of the field URL.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryRequest) GetURL() *fhir.URI {
	if ber == nil {
		return nil
	}
	return ber.URL
}

// Results of execution (transaction/batch/history)// Indicates the results of processing the corresponding 'request' entry in the
// batch or transaction being responded to or what the results of an operation
// where when returning history.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Bundle.json
type BundleEntryResponse struct {

	// The Etag for the resource, if the operation for the entry produced a
	// versioned resource (see [Resource Metadata and
	// Versioning](http.html#versioning) and [Managing Resource
	// Contention](http.html#concurrency)).
	Etag *fhir.String `fhirpath:"etag"`

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

	// The date/time that the resource was modified on the server.
	LastModified *fhir.Instant `fhirpath:"lastModified"`

	// The location header created by processing this operation, populated if the
	// operation returns a location.
	Location *fhir.URI `fhirpath:"location"`

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

	// An OperationOutcome containing hints and warnings produced as part of
	// processing this entry in a batch or transaction.
	Outcome fhir.Resource `fhirpath:"outcome"`

	// The status code returned by processing this entry. The status SHALL start
	// with a 3 digit HTTP code (e.g. 404) and may contain the standard HTTP
	// description associated with the status code.
	Status *fhir.String `fhirpath:"status"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetEtag returns the value of the field Etag.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryResponse) GetEtag() *fhir.String {
	if ber == nil {
		return nil
	}
	return ber.Etag
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryResponse) GetExtension() []*fhir.Extension {
	if ber == nil {
		return nil
	}
	return ber.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryResponse) GetID() string {
	if ber == nil {
		return ""
	}
	return ber.ID
}

// GetLastModified returns the value of the field LastModified.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryResponse) GetLastModified() *fhir.Instant {
	if ber == nil {
		return nil
	}
	return ber.LastModified
}

// GetLocation returns the value of the field Location.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryResponse) GetLocation() *fhir.URI {
	if ber == nil {
		return nil
	}
	return ber.Location
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryResponse) GetModifierExtension() []*fhir.Extension {
	if ber == nil {
		return nil
	}
	return ber.ModifierExtension
}

// GetOutcome returns the value of the field Outcome.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryResponse) GetOutcome() fhir.Resource {
	if ber == nil {
		return nil
	}
	return ber.Outcome
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (ber *BundleEntryResponse) GetStatus() *fhir.String {
	if ber == nil {
		return nil
	}
	return ber.Status
}

// Search related information// Information about the search process that lead to the creation of this
// entry.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Bundle.json
type BundleEntrySearch struct {

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

	// Why this entry is in the result set - whether it's included as a match or
	// because of an _include requirement, or to convey information or warning
	// information about the search process.
	Mode *fhir.Code `fhirpath:"mode"`

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

	// When searching, the server's search ranking score for the entry.
	Score *fhir.Decimal `fhirpath:"score"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (bes *BundleEntrySearch) GetExtension() []*fhir.Extension {
	if bes == nil {
		return nil
	}
	return bes.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (bes *BundleEntrySearch) GetID() string {
	if bes == nil {
		return ""
	}
	return bes.ID
}

// GetMode returns the value of the field Mode.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (bes *BundleEntrySearch) GetMode() *fhir.Code {
	if bes == nil {
		return nil
	}
	return bes.Mode
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (bes *BundleEntrySearch) GetModifierExtension() []*fhir.Extension {
	if bes == nil {
		return nil
	}
	return bes.ModifierExtension
}

// GetScore returns the value of the field Score.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (bes *BundleEntrySearch) GetScore() *fhir.Decimal {
	if bes == nil {
		return nil
	}
	return bes.Score
}

// Links related to this Bundle// A series of links that provide context to this bundle.// Both Bundle.link and Bundle.entry.link are defined to support providing
// additional context when Bundles are used (e.g.
// [HATEOAS](http://en.wikipedia.org/wiki/HATEOAS)).
// Bundle.entry.link corresponds to links found in the HTTP header if the
// resource in the entry was [read](http.html#read) directly.
// This specification defines some specific uses of Bundle.link for
// [searching](search.html#conformance) and [paging](http.html#paging), but no
// specific uses for Bundle.entry.link, and no defined function in a
// transaction - the meaning is implementation specific.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Bundle.json
type BundleLink struct {

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

	// A name which details the functional use for this link - see
	// [http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1](http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1).
	Relation *fhir.String `fhirpath:"relation"`

	// The reference details for the link.
	URL *fhir.URI `fhirpath:"url"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (bl *BundleLink) GetExtension() []*fhir.Extension {
	if bl == nil {
		return nil
	}
	return bl.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (bl *BundleLink) GetID() string {
	if bl == nil {
		return ""
	}
	return bl.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (bl *BundleLink) GetModifierExtension() []*fhir.Extension {
	if bl == nil {
		return nil
	}
	return bl.ModifierExtension
}

// GetRelation returns the value of the field Relation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (bl *BundleLink) GetRelation() *fhir.String {
	if bl == nil {
		return nil
	}
	return bl.Relation
}

// GetURL returns the value of the field URL.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (bl *BundleLink) GetURL() *fhir.URI {
	if bl == nil {
		return nil
	}
	return bl.URL
}

func (b *Bundle) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (b *Bundle) UnmarshalJSON(data []byte) error {
	var raw struct {
		Entry []*BundleEntry `json:"entry"`

		ID            string            `json:"id"`
		Identifier    *fhir.Identifier  `json:"identifier"`
		ImplicitRules *fhir.URI         `json:"implicitRules"`
		Language      *fhir.Code        `json:"language"`
		Link          []*BundleLink     `json:"link"`
		Meta          *fhir.Meta        `json:"meta"`
		Signature     *fhir.Signature   `json:"signature"`
		Timestamp     *fhir.Instant     `json:"timestamp"`
		Total         *fhir.UnsignedInt `json:"total"`
		Type          *fhir.Code        `json:"type"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	b.Entry = raw.Entry
	b.ID = raw.ID
	b.Identifier = raw.Identifier
	b.ImplicitRules = raw.ImplicitRules
	b.Language = raw.Language
	b.Link = raw.Link
	b.Meta = raw.Meta
	b.Signature = raw.Signature
	b.Timestamp = raw.Timestamp
	b.Total = raw.Total
	b.Type = raw.Type
	return nil
}

var _ json.Marshaler = (*Bundle)(nil)
var _ json.Unmarshaler = (*Bundle)(nil)

func (be *BundleEntry) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (be *BundleEntry) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`
		FullURL   *fhir.URI         `json:"fullUrl"`

		ID                string               `json:"id"`
		ModifierExtension []*fhir.Extension    `json:"modifierExtension"`
		Request           *BundleEntryRequest  `json:"request"`
		Resource          fhir.Resource        `json:"resource"`
		Response          *BundleEntryResponse `json:"response"`
		Search            *BundleEntrySearch   `json:"search"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	be.Extension = raw.Extension
	be.FullURL = raw.FullURL
	be.ID = raw.ID
	be.ModifierExtension = raw.ModifierExtension
	be.Request = raw.Request
	be.Resource = raw.Resource
	be.Response = raw.Response
	be.Search = raw.Search
	return nil
}

var _ json.Marshaler = (*BundleEntry)(nil)
var _ json.Unmarshaler = (*BundleEntry)(nil)

func (ber *BundleEntryRequest) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (ber *BundleEntryRequest) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		IfMatch           *fhir.String      `json:"ifMatch"`
		IfModifiedSince   *fhir.Instant     `json:"ifModifiedSince"`
		IfNoneExist       *fhir.String      `json:"ifNoneExist"`
		IfNoneMatch       *fhir.String      `json:"ifNoneMatch"`
		Method            *fhir.Code        `json:"method"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		URL               *fhir.URI         `json:"url"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	ber.Extension = raw.Extension
	ber.ID = raw.ID
	ber.IfMatch = raw.IfMatch
	ber.IfModifiedSince = raw.IfModifiedSince
	ber.IfNoneExist = raw.IfNoneExist
	ber.IfNoneMatch = raw.IfNoneMatch
	ber.Method = raw.Method
	ber.ModifierExtension = raw.ModifierExtension
	ber.URL = raw.URL
	return nil
}

var _ json.Marshaler = (*BundleEntryRequest)(nil)
var _ json.Unmarshaler = (*BundleEntryRequest)(nil)

func (ber *BundleEntryResponse) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (ber *BundleEntryResponse) UnmarshalJSON(data []byte) error {
	var raw struct {
		Etag      *fhir.String      `json:"etag"`
		Extension []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		LastModified      *fhir.Instant     `json:"lastModified"`
		Location          *fhir.URI         `json:"location"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		Outcome           fhir.Resource     `json:"outcome"`
		Status            *fhir.String      `json:"status"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	ber.Etag = raw.Etag
	ber.Extension = raw.Extension
	ber.ID = raw.ID
	ber.LastModified = raw.LastModified
	ber.Location = raw.Location
	ber.ModifierExtension = raw.ModifierExtension
	ber.Outcome = raw.Outcome
	ber.Status = raw.Status
	return nil
}

var _ json.Marshaler = (*BundleEntryResponse)(nil)
var _ json.Unmarshaler = (*BundleEntryResponse)(nil)

func (bes *BundleEntrySearch) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (bes *BundleEntrySearch) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		Mode              *fhir.Code        `json:"mode"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		Score             *fhir.Decimal     `json:"score"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	bes.Extension = raw.Extension
	bes.ID = raw.ID
	bes.Mode = raw.Mode
	bes.ModifierExtension = raw.ModifierExtension
	bes.Score = raw.Score
	return nil
}

var _ json.Marshaler = (*BundleEntrySearch)(nil)
var _ json.Unmarshaler = (*BundleEntrySearch)(nil)

func (bl *BundleLink) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (bl *BundleLink) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		Relation          *fhir.String      `json:"relation"`
		URL               *fhir.URI         `json:"url"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	bl.Extension = raw.Extension
	bl.ID = raw.ID
	bl.ModifierExtension = raw.ModifierExtension
	bl.Relation = raw.Relation
	bl.URL = raw.URL
	return nil
}

var _ json.Marshaler = (*BundleLink)(nil)
var _ json.Unmarshaler = (*BundleLink)(nil)
