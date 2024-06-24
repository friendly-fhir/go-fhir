// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package episodeofcare

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// An association between a patient and an organization / healthcare
// provider(s) during which time encounters may occur. The managing
// organization assumes a level of responsibility for the patient during this
// time.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/EpisodeOfCare
//   - Source File: StructureDefinition-EpisodeOfCare.json
type EpisodeOfCare struct {

	// The set of accounts that may be used for billing for this EpisodeOfCare.
	Account []*fhir.Reference `fhirpath:"account"`

	// The practitioner that is the care manager/care coordinator for this patient.
	CareManager *fhir.Reference `fhirpath:"careManager"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The list of diagnosis relevant to this episode of care.
	Diagnosis []*EpisodeOfCareDiagnosis `fhirpath:"diagnosis"`

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

	// The EpisodeOfCare may be known by different identifiers for different
	// contexts of use, such as when an external agency is tracking the Episode for
	// funding purposes.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// The organization that has assumed the specific responsibilities for the
	// specified duration.
	ManagingOrganization *fhir.Reference `fhirpath:"managingOrganization"`

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

	// The patient who is the focus of this episode of care.
	Patient *fhir.Reference `fhirpath:"patient"`

	// The interval during which the managing organization assumes the defined
	// responsibility.
	Period *fhir.Period `fhirpath:"period"`

	// Referral Request(s) that are fulfilled by this EpisodeOfCare, incoming
	// referrals.
	ReferralRequest []*fhir.Reference `fhirpath:"referralRequest"`

	// planned | waitlist | active | onhold | finished | cancelled.
	Status *fhir.Code `fhirpath:"status"`

	// The history of statuses that the EpisodeOfCare has been through (without
	// requiring processing the history of the resource).
	StatusHistory []*EpisodeOfCareStatusHistory `fhirpath:"statusHistory"`

	// The list of practitioners that may be facilitating this episode of care for
	// specific purposes.
	Team []*fhir.Reference `fhirpath:"team"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// A classification of the type of episode of care; e.g. specialist referral,
	// disease management, type of funded care.
	Type []*fhir.CodeableConcept `fhirpath:"type"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetAccount returns the value of the field Account.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetAccount() []*fhir.Reference {
	if eoc == nil {
		return nil
	}
	return eoc.Account
}

// GetCareManager returns the value of the field CareManager.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetCareManager() *fhir.Reference {
	if eoc == nil {
		return nil
	}
	return eoc.CareManager
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetContained() []fhir.Resource {
	if eoc == nil {
		return nil
	}
	return eoc.Contained
}

// GetDiagnosis returns the value of the field Diagnosis.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetDiagnosis() []*EpisodeOfCareDiagnosis {
	if eoc == nil {
		return nil
	}
	return eoc.Diagnosis
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetExtension() []*fhir.Extension {
	if eoc == nil {
		return nil
	}
	return eoc.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetID() string {
	if eoc == nil {
		return ""
	}
	return eoc.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetIdentifier() []*fhir.Identifier {
	if eoc == nil {
		return nil
	}
	return eoc.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetImplicitRules() *fhir.URI {
	if eoc == nil {
		return nil
	}
	return eoc.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetLanguage() *fhir.Code {
	if eoc == nil {
		return nil
	}
	return eoc.Language
}

// GetManagingOrganization returns the value of the field ManagingOrganization.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetManagingOrganization() *fhir.Reference {
	if eoc == nil {
		return nil
	}
	return eoc.ManagingOrganization
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetMeta() *fhir.Meta {
	if eoc == nil {
		return nil
	}
	return eoc.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetModifierExtension() []*fhir.Extension {
	if eoc == nil {
		return nil
	}
	return eoc.ModifierExtension
}

// GetPatient returns the value of the field Patient.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetPatient() *fhir.Reference {
	if eoc == nil {
		return nil
	}
	return eoc.Patient
}

// GetPeriod returns the value of the field Period.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetPeriod() *fhir.Period {
	if eoc == nil {
		return nil
	}
	return eoc.Period
}

// GetReferralRequest returns the value of the field ReferralRequest.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetReferralRequest() []*fhir.Reference {
	if eoc == nil {
		return nil
	}
	return eoc.ReferralRequest
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetStatus() *fhir.Code {
	if eoc == nil {
		return nil
	}
	return eoc.Status
}

// GetStatusHistory returns the value of the field StatusHistory.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetStatusHistory() []*EpisodeOfCareStatusHistory {
	if eoc == nil {
		return nil
	}
	return eoc.StatusHistory
}

// GetTeam returns the value of the field Team.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetTeam() []*fhir.Reference {
	if eoc == nil {
		return nil
	}
	return eoc.Team
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetText() *fhir.Narrative {
	if eoc == nil {
		return nil
	}
	return eoc.Text
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eoc *EpisodeOfCare) GetType() []*fhir.CodeableConcept {
	if eoc == nil {
		return nil
	}
	return eoc.Type
}

// The list of diagnosis relevant to this episode of care// The list of diagnosis relevant to this episode of care.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-EpisodeOfCare.json
type EpisodeOfCareDiagnosis struct {

	// A list of conditions/problems/diagnoses that this episode of care is
	// intended to be providing care for.
	Condition *fhir.Reference `fhirpath:"condition"`

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

	// Ranking of the diagnosis (for each role type).
	Rank *fhir.PositiveInt `fhirpath:"rank"`

	// Role that this diagnosis has within the episode of care (e.g. admission,
	// billing, discharge …).
	Role *fhir.CodeableConcept `fhirpath:"role"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCondition returns the value of the field Condition.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eocd *EpisodeOfCareDiagnosis) GetCondition() *fhir.Reference {
	if eocd == nil {
		return nil
	}
	return eocd.Condition
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eocd *EpisodeOfCareDiagnosis) GetExtension() []*fhir.Extension {
	if eocd == nil {
		return nil
	}
	return eocd.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eocd *EpisodeOfCareDiagnosis) GetID() string {
	if eocd == nil {
		return ""
	}
	return eocd.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eocd *EpisodeOfCareDiagnosis) GetModifierExtension() []*fhir.Extension {
	if eocd == nil {
		return nil
	}
	return eocd.ModifierExtension
}

// GetRank returns the value of the field Rank.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eocd *EpisodeOfCareDiagnosis) GetRank() *fhir.PositiveInt {
	if eocd == nil {
		return nil
	}
	return eocd.Rank
}

// GetRole returns the value of the field Role.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eocd *EpisodeOfCareDiagnosis) GetRole() *fhir.CodeableConcept {
	if eocd == nil {
		return nil
	}
	return eocd.Role
}

// Past list of status codes (the current status may be included to cover the
// start date of the status)// The history of statuses that the EpisodeOfCare has been through (without
// requiring processing the history of the resource).
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-EpisodeOfCare.json
type EpisodeOfCareStatusHistory struct {

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

	// The period during this EpisodeOfCare that the specific status applied.
	Period *fhir.Period `fhirpath:"period"`

	// planned | waitlist | active | onhold | finished | cancelled.
	Status *fhir.Code `fhirpath:"status"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eocsh *EpisodeOfCareStatusHistory) GetExtension() []*fhir.Extension {
	if eocsh == nil {
		return nil
	}
	return eocsh.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eocsh *EpisodeOfCareStatusHistory) GetID() string {
	if eocsh == nil {
		return ""
	}
	return eocsh.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eocsh *EpisodeOfCareStatusHistory) GetModifierExtension() []*fhir.Extension {
	if eocsh == nil {
		return nil
	}
	return eocsh.ModifierExtension
}

// GetPeriod returns the value of the field Period.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eocsh *EpisodeOfCareStatusHistory) GetPeriod() *fhir.Period {
	if eocsh == nil {
		return nil
	}
	return eocsh.Period
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (eocsh *EpisodeOfCareStatusHistory) GetStatus() *fhir.Code {
	if eocsh == nil {
		return nil
	}
	return eocsh.Status
}
