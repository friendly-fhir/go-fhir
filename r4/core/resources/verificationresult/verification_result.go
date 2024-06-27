// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package verificationresult

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Describes validation requirements, source(s), status and dates for one or
// more elements.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/VerificationResult
//   - Source File: StructureDefinition-VerificationResult.json
type VerificationResult struct {

	// Information about the entity attesting to information.
	Attestation *VerificationResultAttestation `fhirpath:"attestation"`

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

	// The result if validation fails (fatal; warning; record only; none).
	FailureAction *fhir.CodeableConcept `fhirpath:"failureAction"`

	// Frequency of revalidation.
	Frequency *fhir.Timing `fhirpath:"frequency"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes.
	ID string `fhirpath:"id"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// The date/time validation was last completed (including failed validations).
	LastPerformed *fhir.DateTime `fhirpath:"lastPerformed"`

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

	// The frequency with which the target must be validated (none; initial;
	// periodic).
	Need *fhir.CodeableConcept `fhirpath:"need"`

	// The date when target is next validated, if appropriate.
	NextScheduled *fhir.Date `fhirpath:"nextScheduled"`

	// Information about the primary source(s) involved in validation.
	PrimarySource []*VerificationResultPrimarySource `fhirpath:"primarySource"`

	// The validation status of the target (attested; validated; in process;
	// requires revalidation; validation failed; revalidation failed).
	Status *fhir.Code `fhirpath:"status"`

	// When the validation status was updated.
	StatusDate *fhir.DateTime `fhirpath:"statusDate"`

	// A resource that was validated.
	Target []*fhir.Reference `fhirpath:"target"`

	// The fhirpath location(s) within the resource that was validated.
	TargetLocation []*fhir.String `fhirpath:"targetLocation"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// The primary process by which the target is validated (edit check; value set;
	// primary source; multiple sources; standalone; in context).
	ValidationProcess []*fhir.CodeableConcept `fhirpath:"validationProcess"`

	// What the target is validated against (nothing; primary source; multiple
	// sources).
	ValidationType *fhir.CodeableConcept `fhirpath:"validationType"`

	// Information about the entity validating information.
	Validator []*VerificationResultValidator `fhirpath:"validator"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetAttestation returns the value of the field Attestation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetAttestation() *VerificationResultAttestation {
	if vr == nil {
		return nil
	}
	return vr.Attestation
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetContained() []fhir.Resource {
	if vr == nil {
		return nil
	}
	return vr.Contained
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetExtension() []*fhir.Extension {
	if vr == nil {
		return nil
	}
	return vr.Extension
}

// GetFailureAction returns the value of the field FailureAction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetFailureAction() *fhir.CodeableConcept {
	if vr == nil {
		return nil
	}
	return vr.FailureAction
}

// GetFrequency returns the value of the field Frequency.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetFrequency() *fhir.Timing {
	if vr == nil {
		return nil
	}
	return vr.Frequency
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetID() string {
	if vr == nil {
		return ""
	}
	return vr.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetImplicitRules() *fhir.URI {
	if vr == nil {
		return nil
	}
	return vr.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetLanguage() *fhir.Code {
	if vr == nil {
		return nil
	}
	return vr.Language
}

// GetLastPerformed returns the value of the field LastPerformed.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetLastPerformed() *fhir.DateTime {
	if vr == nil {
		return nil
	}
	return vr.LastPerformed
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetMeta() *fhir.Meta {
	if vr == nil {
		return nil
	}
	return vr.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetModifierExtension() []*fhir.Extension {
	if vr == nil {
		return nil
	}
	return vr.ModifierExtension
}

// GetNeed returns the value of the field Need.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetNeed() *fhir.CodeableConcept {
	if vr == nil {
		return nil
	}
	return vr.Need
}

// GetNextScheduled returns the value of the field NextScheduled.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetNextScheduled() *fhir.Date {
	if vr == nil {
		return nil
	}
	return vr.NextScheduled
}

// GetPrimarySource returns the value of the field PrimarySource.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetPrimarySource() []*VerificationResultPrimarySource {
	if vr == nil {
		return nil
	}
	return vr.PrimarySource
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetStatus() *fhir.Code {
	if vr == nil {
		return nil
	}
	return vr.Status
}

// GetStatusDate returns the value of the field StatusDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetStatusDate() *fhir.DateTime {
	if vr == nil {
		return nil
	}
	return vr.StatusDate
}

// GetTarget returns the value of the field Target.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetTarget() []*fhir.Reference {
	if vr == nil {
		return nil
	}
	return vr.Target
}

// GetTargetLocation returns the value of the field TargetLocation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetTargetLocation() []*fhir.String {
	if vr == nil {
		return nil
	}
	return vr.TargetLocation
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetText() *fhir.Narrative {
	if vr == nil {
		return nil
	}
	return vr.Text
}

// GetValidationProcess returns the value of the field ValidationProcess.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetValidationProcess() []*fhir.CodeableConcept {
	if vr == nil {
		return nil
	}
	return vr.ValidationProcess
}

// GetValidationType returns the value of the field ValidationType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetValidationType() *fhir.CodeableConcept {
	if vr == nil {
		return nil
	}
	return vr.ValidationType
}

// GetValidator returns the value of the field Validator.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vr *VerificationResult) GetValidator() []*VerificationResultValidator {
	if vr == nil {
		return nil
	}
	return vr.Validator
}

// Information about the entity attesting to information// Information about the entity attesting to information.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-VerificationResult.json
type VerificationResultAttestation struct {

	// The method by which attested information was submitted/retrieved (manual;
	// API; Push).
	CommunicationMethod *fhir.CodeableConcept `fhirpath:"communicationMethod"`

	// The date the information was attested to.
	Date *fhir.Date `fhirpath:"date"`

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

	// When the who is asserting on behalf of another (organization or individual).
	OnBehalfOf *fhir.Reference `fhirpath:"onBehalfOf"`

	// A digital identity certificate associated with the proxy entity submitting
	// attested information on behalf of the attestation source.
	ProxyIdentityCertificate *fhir.String `fhirpath:"proxyIdentityCertificate"`

	// Signed assertion by the proxy entity indicating that they have the right to
	// submit attested information on behalf of the attestation source.
	ProxySignature *fhir.Signature `fhirpath:"proxySignature"`

	// A digital identity certificate associated with the attestation source.
	SourceIdentityCertificate *fhir.String `fhirpath:"sourceIdentityCertificate"`

	// Signed assertion by the attestation source that they have attested to the
	// information.
	SourceSignature *fhir.Signature `fhirpath:"sourceSignature"`

	// The individual or organization attesting to information.
	Who *fhir.Reference `fhirpath:"who"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCommunicationMethod returns the value of the field CommunicationMethod.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vra *VerificationResultAttestation) GetCommunicationMethod() *fhir.CodeableConcept {
	if vra == nil {
		return nil
	}
	return vra.CommunicationMethod
}

// GetDate returns the value of the field Date.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vra *VerificationResultAttestation) GetDate() *fhir.Date {
	if vra == nil {
		return nil
	}
	return vra.Date
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vra *VerificationResultAttestation) GetExtension() []*fhir.Extension {
	if vra == nil {
		return nil
	}
	return vra.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vra *VerificationResultAttestation) GetID() string {
	if vra == nil {
		return ""
	}
	return vra.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vra *VerificationResultAttestation) GetModifierExtension() []*fhir.Extension {
	if vra == nil {
		return nil
	}
	return vra.ModifierExtension
}

// GetOnBehalfOf returns the value of the field OnBehalfOf.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vra *VerificationResultAttestation) GetOnBehalfOf() *fhir.Reference {
	if vra == nil {
		return nil
	}
	return vra.OnBehalfOf
}

// GetProxyIdentityCertificate returns the value of the field ProxyIdentityCertificate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vra *VerificationResultAttestation) GetProxyIdentityCertificate() *fhir.String {
	if vra == nil {
		return nil
	}
	return vra.ProxyIdentityCertificate
}

// GetProxySignature returns the value of the field ProxySignature.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vra *VerificationResultAttestation) GetProxySignature() *fhir.Signature {
	if vra == nil {
		return nil
	}
	return vra.ProxySignature
}

// GetSourceIdentityCertificate returns the value of the field SourceIdentityCertificate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vra *VerificationResultAttestation) GetSourceIdentityCertificate() *fhir.String {
	if vra == nil {
		return nil
	}
	return vra.SourceIdentityCertificate
}

// GetSourceSignature returns the value of the field SourceSignature.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vra *VerificationResultAttestation) GetSourceSignature() *fhir.Signature {
	if vra == nil {
		return nil
	}
	return vra.SourceSignature
}

// GetWho returns the value of the field Who.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vra *VerificationResultAttestation) GetWho() *fhir.Reference {
	if vra == nil {
		return nil
	}
	return vra.Who
}

// Information about the primary source(s) involved in validation// Information about the primary source(s) involved in validation.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-VerificationResult.json
type VerificationResultPrimarySource struct {

	// Ability of the primary source to push updates/alerts (yes; no;
	// undetermined).
	CanPushUpdates *fhir.CodeableConcept `fhirpath:"canPushUpdates"`

	// Method for communicating with the primary source (manual; API; Push).
	CommunicationMethod []*fhir.CodeableConcept `fhirpath:"communicationMethod"`

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

	// Type of alerts/updates the primary source can send (specific requested
	// changes; any changes; as defined by source).
	PushTypeAvailable []*fhir.CodeableConcept `fhirpath:"pushTypeAvailable"`

	// Type of primary source (License Board; Primary Education; Continuing
	// Education; Postal Service; Relationship owner; Registration Authority; legal
	// source; issuing source; authoritative source).
	Type []*fhir.CodeableConcept `fhirpath:"type"`

	// When the target was validated against the primary source.
	ValidationDate *fhir.DateTime `fhirpath:"validationDate"`

	// Status of the validation of the target against the primary source
	// (successful; failed; unknown).
	ValidationStatus *fhir.CodeableConcept `fhirpath:"validationStatus"`

	// Reference to the primary source.
	Who *fhir.Reference `fhirpath:"who"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCanPushUpdates returns the value of the field CanPushUpdates.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrps *VerificationResultPrimarySource) GetCanPushUpdates() *fhir.CodeableConcept {
	if vrps == nil {
		return nil
	}
	return vrps.CanPushUpdates
}

// GetCommunicationMethod returns the value of the field CommunicationMethod.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrps *VerificationResultPrimarySource) GetCommunicationMethod() []*fhir.CodeableConcept {
	if vrps == nil {
		return nil
	}
	return vrps.CommunicationMethod
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrps *VerificationResultPrimarySource) GetExtension() []*fhir.Extension {
	if vrps == nil {
		return nil
	}
	return vrps.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrps *VerificationResultPrimarySource) GetID() string {
	if vrps == nil {
		return ""
	}
	return vrps.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrps *VerificationResultPrimarySource) GetModifierExtension() []*fhir.Extension {
	if vrps == nil {
		return nil
	}
	return vrps.ModifierExtension
}

// GetPushTypeAvailable returns the value of the field PushTypeAvailable.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrps *VerificationResultPrimarySource) GetPushTypeAvailable() []*fhir.CodeableConcept {
	if vrps == nil {
		return nil
	}
	return vrps.PushTypeAvailable
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrps *VerificationResultPrimarySource) GetType() []*fhir.CodeableConcept {
	if vrps == nil {
		return nil
	}
	return vrps.Type
}

// GetValidationDate returns the value of the field ValidationDate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrps *VerificationResultPrimarySource) GetValidationDate() *fhir.DateTime {
	if vrps == nil {
		return nil
	}
	return vrps.ValidationDate
}

// GetValidationStatus returns the value of the field ValidationStatus.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrps *VerificationResultPrimarySource) GetValidationStatus() *fhir.CodeableConcept {
	if vrps == nil {
		return nil
	}
	return vrps.ValidationStatus
}

// GetWho returns the value of the field Who.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrps *VerificationResultPrimarySource) GetWho() *fhir.Reference {
	if vrps == nil {
		return nil
	}
	return vrps.Who
}

// Information about the entity validating information// Information about the entity validating information.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-VerificationResult.json
type VerificationResultValidator struct {

	// Signed assertion by the validator that they have validated the information.
	AttestationSignature *fhir.Signature `fhirpath:"attestationSignature"`

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

	// A digital identity certificate associated with the validator.
	IdentityCertificate *fhir.String `fhirpath:"identityCertificate"`

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

	// Reference to the organization validating information.
	Organization *fhir.Reference `fhirpath:"organization"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAttestationSignature returns the value of the field AttestationSignature.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrv *VerificationResultValidator) GetAttestationSignature() *fhir.Signature {
	if vrv == nil {
		return nil
	}
	return vrv.AttestationSignature
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrv *VerificationResultValidator) GetExtension() []*fhir.Extension {
	if vrv == nil {
		return nil
	}
	return vrv.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrv *VerificationResultValidator) GetID() string {
	if vrv == nil {
		return ""
	}
	return vrv.ID
}

// GetIdentityCertificate returns the value of the field IdentityCertificate.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrv *VerificationResultValidator) GetIdentityCertificate() *fhir.String {
	if vrv == nil {
		return nil
	}
	return vrv.IdentityCertificate
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrv *VerificationResultValidator) GetModifierExtension() []*fhir.Extension {
	if vrv == nil {
		return nil
	}
	return vrv.ModifierExtension
}

// GetOrganization returns the value of the field Organization.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (vrv *VerificationResultValidator) GetOrganization() *fhir.Reference {
	if vrv == nil {
		return nil
	}
	return vrv.Organization
}

func (vr *VerificationResult) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (vr *VerificationResult) UnmarshalJSON(data []byte) error {
	var raw struct {
		Attestation   *VerificationResultAttestation `json:"attestation"`
		Contained     []fhir.Resource                `json:"contained"`
		Extension     []*fhir.Extension              `json:"extension"`
		FailureAction *fhir.CodeableConcept          `json:"failureAction"`
		Frequency     *fhir.Timing                   `json:"frequency"`

		ID                string                             `json:"id"`
		ImplicitRules     *fhir.URI                          `json:"implicitRules"`
		Language          *fhir.Code                         `json:"language"`
		LastPerformed     *fhir.DateTime                     `json:"lastPerformed"`
		Meta              *fhir.Meta                         `json:"meta"`
		ModifierExtension []*fhir.Extension                  `json:"modifierExtension"`
		Need              *fhir.CodeableConcept              `json:"need"`
		NextScheduled     *fhir.Date                         `json:"nextScheduled"`
		PrimarySource     []*VerificationResultPrimarySource `json:"primarySource"`
		Status            *fhir.Code                         `json:"status"`
		StatusDate        *fhir.DateTime                     `json:"statusDate"`
		Target            []*fhir.Reference                  `json:"target"`
		TargetLocation    []*fhir.String                     `json:"targetLocation"`
		Text              *fhir.Narrative                    `json:"text"`
		ValidationProcess []*fhir.CodeableConcept            `json:"validationProcess"`
		ValidationType    *fhir.CodeableConcept              `json:"validationType"`
		Validator         []*VerificationResultValidator     `json:"validator"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	vr.Attestation = raw.Attestation
	vr.Contained = raw.Contained
	vr.Extension = raw.Extension
	vr.FailureAction = raw.FailureAction
	vr.Frequency = raw.Frequency
	vr.ID = raw.ID
	vr.ImplicitRules = raw.ImplicitRules
	vr.Language = raw.Language
	vr.LastPerformed = raw.LastPerformed
	vr.Meta = raw.Meta
	vr.ModifierExtension = raw.ModifierExtension
	vr.Need = raw.Need
	vr.NextScheduled = raw.NextScheduled
	vr.PrimarySource = raw.PrimarySource
	vr.Status = raw.Status
	vr.StatusDate = raw.StatusDate
	vr.Target = raw.Target
	vr.TargetLocation = raw.TargetLocation
	vr.Text = raw.Text
	vr.ValidationProcess = raw.ValidationProcess
	vr.ValidationType = raw.ValidationType
	vr.Validator = raw.Validator
	return nil
}

var _ json.Marshaler = (*VerificationResult)(nil)
var _ json.Unmarshaler = (*VerificationResult)(nil)

func (vra *VerificationResultAttestation) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (vra *VerificationResultAttestation) UnmarshalJSON(data []byte) error {
	var raw struct {
		CommunicationMethod *fhir.CodeableConcept `json:"communicationMethod"`
		Date                *fhir.Date            `json:"date"`
		Extension           []*fhir.Extension     `json:"extension"`

		ID                        string            `json:"id"`
		ModifierExtension         []*fhir.Extension `json:"modifierExtension"`
		OnBehalfOf                *fhir.Reference   `json:"onBehalfOf"`
		ProxyIdentityCertificate  *fhir.String      `json:"proxyIdentityCertificate"`
		ProxySignature            *fhir.Signature   `json:"proxySignature"`
		SourceIdentityCertificate *fhir.String      `json:"sourceIdentityCertificate"`
		SourceSignature           *fhir.Signature   `json:"sourceSignature"`
		Who                       *fhir.Reference   `json:"who"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	vra.CommunicationMethod = raw.CommunicationMethod
	vra.Date = raw.Date
	vra.Extension = raw.Extension
	vra.ID = raw.ID
	vra.ModifierExtension = raw.ModifierExtension
	vra.OnBehalfOf = raw.OnBehalfOf
	vra.ProxyIdentityCertificate = raw.ProxyIdentityCertificate
	vra.ProxySignature = raw.ProxySignature
	vra.SourceIdentityCertificate = raw.SourceIdentityCertificate
	vra.SourceSignature = raw.SourceSignature
	vra.Who = raw.Who
	return nil
}

var _ json.Marshaler = (*VerificationResultAttestation)(nil)
var _ json.Unmarshaler = (*VerificationResultAttestation)(nil)

func (vrps *VerificationResultPrimarySource) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (vrps *VerificationResultPrimarySource) UnmarshalJSON(data []byte) error {
	var raw struct {
		CanPushUpdates      *fhir.CodeableConcept   `json:"canPushUpdates"`
		CommunicationMethod []*fhir.CodeableConcept `json:"communicationMethod"`
		Extension           []*fhir.Extension       `json:"extension"`

		ID                string                  `json:"id"`
		ModifierExtension []*fhir.Extension       `json:"modifierExtension"`
		PushTypeAvailable []*fhir.CodeableConcept `json:"pushTypeAvailable"`
		Type              []*fhir.CodeableConcept `json:"type"`
		ValidationDate    *fhir.DateTime          `json:"validationDate"`
		ValidationStatus  *fhir.CodeableConcept   `json:"validationStatus"`
		Who               *fhir.Reference         `json:"who"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	vrps.CanPushUpdates = raw.CanPushUpdates
	vrps.CommunicationMethod = raw.CommunicationMethod
	vrps.Extension = raw.Extension
	vrps.ID = raw.ID
	vrps.ModifierExtension = raw.ModifierExtension
	vrps.PushTypeAvailable = raw.PushTypeAvailable
	vrps.Type = raw.Type
	vrps.ValidationDate = raw.ValidationDate
	vrps.ValidationStatus = raw.ValidationStatus
	vrps.Who = raw.Who
	return nil
}

var _ json.Marshaler = (*VerificationResultPrimarySource)(nil)
var _ json.Unmarshaler = (*VerificationResultPrimarySource)(nil)

func (vrv *VerificationResultValidator) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (vrv *VerificationResultValidator) UnmarshalJSON(data []byte) error {
	var raw struct {
		AttestationSignature *fhir.Signature   `json:"attestationSignature"`
		Extension            []*fhir.Extension `json:"extension"`

		ID                  string            `json:"id"`
		IdentityCertificate *fhir.String      `json:"identityCertificate"`
		ModifierExtension   []*fhir.Extension `json:"modifierExtension"`
		Organization        *fhir.Reference   `json:"organization"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	vrv.AttestationSignature = raw.AttestationSignature
	vrv.Extension = raw.Extension
	vrv.ID = raw.ID
	vrv.IdentityCertificate = raw.IdentityCertificate
	vrv.ModifierExtension = raw.ModifierExtension
	vrv.Organization = raw.Organization
	return nil
}

var _ json.Marshaler = (*VerificationResultValidator)(nil)
var _ json.Unmarshaler = (*VerificationResultValidator)(nil)
