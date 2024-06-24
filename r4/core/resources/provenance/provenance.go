// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package provenance

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Provenance of a resource is a record that describes entities and processes
// involved in producing and delivering or otherwise influencing that resource.
// Provenance provides a critical foundation for assessing authenticity,
// enabling trust, and allowing reproducibility. Provenance assertions are a
// form of contextual metadata and can themselves become important records with
// their own provenance. Provenance statement indicates clinical significance
// in terms of confidence in authenticity, reliability, and trustworthiness,
// integrity, and stage in lifecycle (e.g. Document Completion - has the
// artifact been legally authenticated), all of which may impact security,
// privacy, and trust policies.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Provenance
//   - Source File: StructureDefinition-Provenance.json
type Provenance struct {

	// An activity is something that occurs over a period of time and acts upon or
	// with entities; it may include consuming, processing, transforming,
	// modifying, relocating, using, or generating entities.
	Activity *fhir.CodeableConcept `fhirpath:"activity"`

	// An actor taking a role in an activity for which it can be assigned some
	// degree of responsibility for the activity taking place.
	Agent []*ProvenanceAgent `fhirpath:"agent"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// An entity used in this activity.
	Entity []*ProvenanceEntity `fhirpath:"entity"`

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

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// Where the activity occurred, if relevant.
	Location *fhir.Reference `fhirpath:"location"`

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

	// The period during which the activity occurred.
	Occurred fhir.Element `fhirpath:"occurred"`

	// Policy or plan the activity was defined by. Typically, a single activity may
	// have multiple applicable policy documents, such as patient consent,
	// guarantor funding, etc.
	Policy []*fhir.URI `fhirpath:"policy"`

	// The reason that the activity was taking place.
	Reason []*fhir.CodeableConcept `fhirpath:"reason"`

	// The instant of time at which the activity was recorded.
	Recorded *fhir.Instant `fhirpath:"recorded"`

	// A digital signature on the target Reference(s). The signer should match a
	// Provenance.agent. The purpose of the signature is indicated.
	Signature []*fhir.Signature `fhirpath:"signature"`

	// The Reference(s) that were generated or updated by the activity described in
	// this resource. A provenance can point to more than one target if multiple
	// resources were created/updated by the same activity.
	Target []*fhir.Reference `fhirpath:"target"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	profileimpl.BaseProvenance
	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetActivity returns the value of the field Activity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetActivity() *fhir.CodeableConcept {
	if p == nil {
		return nil
	}
	return p.Activity
}

// GetAgent returns the value of the field Agent.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetAgent() []*ProvenanceAgent {
	if p == nil {
		return nil
	}
	return p.Agent
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetContained() []fhir.Resource {
	if p == nil {
		return nil
	}
	return p.Contained
}

// GetEntity returns the value of the field Entity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetEntity() []*ProvenanceEntity {
	if p == nil {
		return nil
	}
	return p.Entity
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetExtension() []*fhir.Extension {
	if p == nil {
		return nil
	}
	return p.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetImplicitRules() *fhir.URI {
	if p == nil {
		return nil
	}
	return p.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetLanguage() *fhir.Code {
	if p == nil {
		return nil
	}
	return p.Language
}

// GetLocation returns the value of the field Location.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetLocation() *fhir.Reference {
	if p == nil {
		return nil
	}
	return p.Location
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetMeta() *fhir.Meta {
	if p == nil {
		return nil
	}
	return p.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetModifierExtension() []*fhir.Extension {
	if p == nil {
		return nil
	}
	return p.ModifierExtension
}

// GetOccurred returns the value of the field Occurred.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetOccurred() fhir.Element {
	if p == nil {
		return nil
	}
	return p.Occurred
}

// GetOccurredPeriod returns the value of the field Occurred.
func (p *Provenance) GetOccurredPeriod() *fhir.Period {
	if p == nil {
		return nil
	}
	val, ok := p.Occurred.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
}

// GetOccurredDateTime returns the value of the field Occurred.
func (p *Provenance) GetOccurredDateTime() *fhir.DateTime {
	if p == nil {
		return nil
	}
	val, ok := p.Occurred.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
} // GetPolicy returns the value of the field Policy.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetPolicy() []*fhir.URI {
	if p == nil {
		return nil
	}
	return p.Policy
}

// GetReason returns the value of the field Reason.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetReason() []*fhir.CodeableConcept {
	if p == nil {
		return nil
	}
	return p.Reason
}

// GetRecorded returns the value of the field Recorded.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetRecorded() *fhir.Instant {
	if p == nil {
		return nil
	}
	return p.Recorded
}

// GetSignature returns the value of the field Signature.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetSignature() []*fhir.Signature {
	if p == nil {
		return nil
	}
	return p.Signature
}

// GetTarget returns the value of the field Target.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetTarget() []*fhir.Reference {
	if p == nil {
		return nil
	}
	return p.Target
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Provenance) GetText() *fhir.Narrative {
	if p == nil {
		return nil
	}
	return p.Text
}

// Actor involved// An actor taking a role in an activity for which it can be assigned some
// degree of responsibility for the activity taking place.// Several agents may be associated (i.e. has some responsibility for an
// activity) with an activity and vice-versa.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Provenance.json
type ProvenanceAgent struct {

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

	// The individual, device, or organization for whom the change was made.
	OnBehalfOf *fhir.Reference `fhirpath:"onBehalfOf"`

	// The function of the agent with respect to the activity. The security role
	// enabling the agent with respect to the activity.
	Role []*fhir.CodeableConcept `fhirpath:"role"`

	// The participation the agent had with respect to the activity.
	Type *fhir.CodeableConcept `fhirpath:"type"`

	// The individual, device or organization that participated in the event.
	Who *fhir.Reference `fhirpath:"who"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pa *ProvenanceAgent) GetExtension() []*fhir.Extension {
	if pa == nil {
		return nil
	}
	return pa.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pa *ProvenanceAgent) GetID() string {
	if pa == nil {
		return ""
	}
	return pa.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pa *ProvenanceAgent) GetModifierExtension() []*fhir.Extension {
	if pa == nil {
		return nil
	}
	return pa.ModifierExtension
}

// GetOnBehalfOf returns the value of the field OnBehalfOf.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pa *ProvenanceAgent) GetOnBehalfOf() *fhir.Reference {
	if pa == nil {
		return nil
	}
	return pa.OnBehalfOf
}

// GetRole returns the value of the field Role.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pa *ProvenanceAgent) GetRole() []*fhir.CodeableConcept {
	if pa == nil {
		return nil
	}
	return pa.Role
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pa *ProvenanceAgent) GetType() *fhir.CodeableConcept {
	if pa == nil {
		return nil
	}
	return pa.Type
}

// GetWho returns the value of the field Who.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pa *ProvenanceAgent) GetWho() *fhir.Reference {
	if pa == nil {
		return nil
	}
	return pa.Who
}

// An entity used in this activity// An entity used in this activity.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Provenance.json
type ProvenanceEntity struct {

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

	// How the entity was used during the activity.
	Role *fhir.Code `fhirpath:"role"`

	// Identity of the Entity used. May be a logical or physical uri and maybe
	// absolute or relative.
	What *fhir.Reference `fhirpath:"what"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pe *ProvenanceEntity) GetExtension() []*fhir.Extension {
	if pe == nil {
		return nil
	}
	return pe.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pe *ProvenanceEntity) GetID() string {
	if pe == nil {
		return ""
	}
	return pe.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pe *ProvenanceEntity) GetModifierExtension() []*fhir.Extension {
	if pe == nil {
		return nil
	}
	return pe.ModifierExtension
}

// GetRole returns the value of the field Role.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pe *ProvenanceEntity) GetRole() *fhir.Code {
	if pe == nil {
		return nil
	}
	return pe.Role
}

// GetWhat returns the value of the field What.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pe *ProvenanceEntity) GetWhat() *fhir.Reference {
	if pe == nil {
		return nil
	}
	return pe.What
}