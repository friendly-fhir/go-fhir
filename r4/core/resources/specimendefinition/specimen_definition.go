// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package specimendefinition

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// A kind of specimen with associated set of requirements.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/SpecimenDefinition
//   - Source File: StructureDefinition-SpecimenDefinition.json
type SpecimenDefinition struct {

	// The action to be performed for collecting the specimen.
	Collection []*fhir.CodeableConcept `fhirpath:"collection"`

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

	// A business identifier associated with the kind of specimen.
	Identifier *fhir.Identifier `fhirpath:"identifier"`

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

	// Preparation of the patient for specimen collection.
	PatientPreparation []*fhir.CodeableConcept `fhirpath:"patientPreparation"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// Time aspect of specimen collection (duration or offset).
	TimeAspect *fhir.String `fhirpath:"timeAspect"`

	// The kind of material to be collected.
	TypeCollected *fhir.CodeableConcept `fhirpath:"typeCollected"`

	// Specimen conditioned in a container as expected by the testing laboratory.
	TypeTested []*SpecimenDefinitionTypeTested `fhirpath:"typeTested"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetCollection returns the value of the field Collection.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetCollection() []*fhir.CodeableConcept {
	if sd == nil {
		return nil
	}
	return sd.Collection
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetContained() []fhir.Resource {
	if sd == nil {
		return nil
	}
	return sd.Contained
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetExtension() []*fhir.Extension {
	if sd == nil {
		return nil
	}
	return sd.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetID() string {
	if sd == nil {
		return ""
	}
	return sd.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetIdentifier() *fhir.Identifier {
	if sd == nil {
		return nil
	}
	return sd.Identifier
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetImplicitRules() *fhir.URI {
	if sd == nil {
		return nil
	}
	return sd.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetLanguage() *fhir.Code {
	if sd == nil {
		return nil
	}
	return sd.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetMeta() *fhir.Meta {
	if sd == nil {
		return nil
	}
	return sd.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetModifierExtension() []*fhir.Extension {
	if sd == nil {
		return nil
	}
	return sd.ModifierExtension
}

// GetPatientPreparation returns the value of the field PatientPreparation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetPatientPreparation() []*fhir.CodeableConcept {
	if sd == nil {
		return nil
	}
	return sd.PatientPreparation
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetText() *fhir.Narrative {
	if sd == nil {
		return nil
	}
	return sd.Text
}

// GetTimeAspect returns the value of the field TimeAspect.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetTimeAspect() *fhir.String {
	if sd == nil {
		return nil
	}
	return sd.TimeAspect
}

// GetTypeCollected returns the value of the field TypeCollected.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetTypeCollected() *fhir.CodeableConcept {
	if sd == nil {
		return nil
	}
	return sd.TypeCollected
}

// GetTypeTested returns the value of the field TypeTested.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sd *SpecimenDefinition) GetTypeTested() []*SpecimenDefinitionTypeTested {
	if sd == nil {
		return nil
	}
	return sd.TypeTested
}

// Specimen in container intended for testing by lab// Specimen conditioned in a container as expected by the testing laboratory.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-SpecimenDefinition.json
type SpecimenDefinitionTypeTested struct {

	// The specimen's container.
	Container *SpecimenDefinitionTypeTestedContainer `fhirpath:"container"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// Set of instructions for preservation/transport of the specimen at a defined
	// temperature interval, prior the testing process.
	Handling []*SpecimenDefinitionTypeTestedHandling `fhirpath:"handling"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID string `fhirpath:"id"`

	// Primary of secondary specimen.
	IsDerived *fhir.Boolean `fhirpath:"isDerived"`

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

	// The preference for this type of conditioned specimen.
	Preference *fhir.Code `fhirpath:"preference"`

	// Criterion for rejection of the specimen in its container by the laboratory.
	RejectionCriterion []*fhir.CodeableConcept `fhirpath:"rejectionCriterion"`

	// Requirements for delivery and special handling of this kind of conditioned
	// specimen.
	Requirement *fhir.String `fhirpath:"requirement"`

	// The usual time that a specimen of this kind is retained after the ordered
	// tests are completed, for the purpose of additional testing.
	RetentionTime *fhir.Duration `fhirpath:"retentionTime"`

	// The kind of specimen conditioned for testing expected by lab.
	Type *fhir.CodeableConcept `fhirpath:"type"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetContainer returns the value of the field Container.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtt *SpecimenDefinitionTypeTested) GetContainer() *SpecimenDefinitionTypeTestedContainer {
	if sdtt == nil {
		return nil
	}
	return sdtt.Container
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtt *SpecimenDefinitionTypeTested) GetExtension() []*fhir.Extension {
	if sdtt == nil {
		return nil
	}
	return sdtt.Extension
}

// GetHandling returns the value of the field Handling.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtt *SpecimenDefinitionTypeTested) GetHandling() []*SpecimenDefinitionTypeTestedHandling {
	if sdtt == nil {
		return nil
	}
	return sdtt.Handling
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtt *SpecimenDefinitionTypeTested) GetID() string {
	if sdtt == nil {
		return ""
	}
	return sdtt.ID
}

// GetIsDerived returns the value of the field IsDerived.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtt *SpecimenDefinitionTypeTested) GetIsDerived() *fhir.Boolean {
	if sdtt == nil {
		return nil
	}
	return sdtt.IsDerived
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtt *SpecimenDefinitionTypeTested) GetModifierExtension() []*fhir.Extension {
	if sdtt == nil {
		return nil
	}
	return sdtt.ModifierExtension
}

// GetPreference returns the value of the field Preference.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtt *SpecimenDefinitionTypeTested) GetPreference() *fhir.Code {
	if sdtt == nil {
		return nil
	}
	return sdtt.Preference
}

// GetRejectionCriterion returns the value of the field RejectionCriterion.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtt *SpecimenDefinitionTypeTested) GetRejectionCriterion() []*fhir.CodeableConcept {
	if sdtt == nil {
		return nil
	}
	return sdtt.RejectionCriterion
}

// GetRequirement returns the value of the field Requirement.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtt *SpecimenDefinitionTypeTested) GetRequirement() *fhir.String {
	if sdtt == nil {
		return nil
	}
	return sdtt.Requirement
}

// GetRetentionTime returns the value of the field RetentionTime.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtt *SpecimenDefinitionTypeTested) GetRetentionTime() *fhir.Duration {
	if sdtt == nil {
		return nil
	}
	return sdtt.RetentionTime
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtt *SpecimenDefinitionTypeTested) GetType() *fhir.CodeableConcept {
	if sdtt == nil {
		return nil
	}
	return sdtt.Type
}

// The specimen's container// The specimen's container.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-SpecimenDefinition.json
type SpecimenDefinitionTypeTestedContainer struct {

	// Substance introduced in the kind of container to preserve, maintain or
	// enhance the specimen. Examples: Formalin, Citrate, EDTA.
	Additive []*SpecimenDefinitionTypeTestedContainerAdditive `fhirpath:"additive"`

	// Color of container cap.
	Cap *fhir.CodeableConcept `fhirpath:"cap"`

	// The capacity (volume or other measure) of this kind of container.
	Capacity *fhir.Quantity `fhirpath:"capacity"`

	// The textual description of the kind of container.
	Description *fhir.String `fhirpath:"description"`

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

	// The type of material of the container.
	Material *fhir.CodeableConcept `fhirpath:"material"`

	// The minimum volume to be conditioned in the container.
	MinimumVolume fhir.Element `fhirpath:"minimumVolume"`

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

	// Special processing that should be applied to the container for this kind of
	// specimen.
	Preparation *fhir.String `fhirpath:"preparation"`

	// The type of container used to contain this kind of specimen.
	Type *fhir.CodeableConcept `fhirpath:"type"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAdditive returns the value of the field Additive.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetAdditive() []*SpecimenDefinitionTypeTestedContainerAdditive {
	if sdttc == nil {
		return nil
	}
	return sdttc.Additive
}

// GetCap returns the value of the field Cap.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetCap() *fhir.CodeableConcept {
	if sdttc == nil {
		return nil
	}
	return sdttc.Cap
}

// GetCapacity returns the value of the field Capacity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetCapacity() *fhir.Quantity {
	if sdttc == nil {
		return nil
	}
	return sdttc.Capacity
}

// GetDescription returns the value of the field Description.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetDescription() *fhir.String {
	if sdttc == nil {
		return nil
	}
	return sdttc.Description
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetExtension() []*fhir.Extension {
	if sdttc == nil {
		return nil
	}
	return sdttc.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetID() string {
	if sdttc == nil {
		return ""
	}
	return sdttc.ID
}

// GetMaterial returns the value of the field Material.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetMaterial() *fhir.CodeableConcept {
	if sdttc == nil {
		return nil
	}
	return sdttc.Material
}

// GetMinimumVolume returns the value of the field MinimumVolume.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetMinimumVolume() fhir.Element {
	if sdttc == nil {
		return nil
	}
	return sdttc.MinimumVolume
}

// GetMinimumVolumeQuantity returns the value of the field MinimumVolume.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetMinimumVolumeQuantity() *fhir.Quantity {
	if sdttc == nil {
		return nil
	}
	val, ok := sdttc.MinimumVolume.(*fhir.Quantity)
	if !ok {
		return nil
	}
	return val
}

// GetMinimumVolumeString returns the value of the field MinimumVolume.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetMinimumVolumeString() *fhir.String {
	if sdttc == nil {
		return nil
	}
	val, ok := sdttc.MinimumVolume.(*fhir.String)
	if !ok {
		return nil
	}
	return val
} // GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetModifierExtension() []*fhir.Extension {
	if sdttc == nil {
		return nil
	}
	return sdttc.ModifierExtension
}

// GetPreparation returns the value of the field Preparation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetPreparation() *fhir.String {
	if sdttc == nil {
		return nil
	}
	return sdttc.Preparation
}

// GetType returns the value of the field Type.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttc *SpecimenDefinitionTypeTestedContainer) GetType() *fhir.CodeableConcept {
	if sdttc == nil {
		return nil
	}
	return sdttc.Type
}

// Additive associated with container// Substance introduced in the kind of container to preserve, maintain or
// enhance the specimen. Examples: Formalin, Citrate, EDTA.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-SpecimenDefinition.json
type SpecimenDefinitionTypeTestedContainerAdditive struct {

	// Substance introduced in the kind of container to preserve, maintain or
	// enhance the specimen. Examples: Formalin, Citrate, EDTA.
	Additive fhir.Element `fhirpath:"additive"`

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

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetAdditive returns the value of the field Additive.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttca *SpecimenDefinitionTypeTestedContainerAdditive) GetAdditive() fhir.Element {
	if sdttca == nil {
		return nil
	}
	return sdttca.Additive
}

// GetAdditiveCodeableConcept returns the value of the field Additive.
func (sdttca *SpecimenDefinitionTypeTestedContainerAdditive) GetAdditiveCodeableConcept() *fhir.CodeableConcept {
	if sdttca == nil {
		return nil
	}
	val, ok := sdttca.Additive.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetAdditiveReference returns the value of the field Additive.
func (sdttca *SpecimenDefinitionTypeTestedContainerAdditive) GetAdditiveReference() *fhir.Reference {
	if sdttca == nil {
		return nil
	}
	val, ok := sdttca.Additive.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
} // GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttca *SpecimenDefinitionTypeTestedContainerAdditive) GetExtension() []*fhir.Extension {
	if sdttca == nil {
		return nil
	}
	return sdttca.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttca *SpecimenDefinitionTypeTestedContainerAdditive) GetID() string {
	if sdttca == nil {
		return ""
	}
	return sdttca.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdttca *SpecimenDefinitionTypeTestedContainerAdditive) GetModifierExtension() []*fhir.Extension {
	if sdttca == nil {
		return nil
	}
	return sdttca.ModifierExtension
}

// Specimen handling before testing// Set of instructions for preservation/transport of the specimen at a defined
// temperature interval, prior the testing process.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-SpecimenDefinition.json
type SpecimenDefinitionTypeTestedHandling struct {

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

	// Additional textual instructions for the preservation or transport of the
	// specimen. For instance, 'Protect from light exposure'.
	Instruction *fhir.String `fhirpath:"instruction"`

	// The maximum time interval of preservation of the specimen with these
	// conditions.
	MaxDuration *fhir.Duration `fhirpath:"maxDuration"`

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

	// It qualifies the interval of temperature, which characterizes an occurrence
	// of handling. Conditions that are not related to temperature may be handled
	// in the instruction element.
	TemperatureQualifier *fhir.CodeableConcept `fhirpath:"temperatureQualifier"`

	// The temperature interval for this set of handling instructions.
	TemperatureRange *fhir.Range `fhirpath:"temperatureRange"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtth *SpecimenDefinitionTypeTestedHandling) GetExtension() []*fhir.Extension {
	if sdtth == nil {
		return nil
	}
	return sdtth.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtth *SpecimenDefinitionTypeTestedHandling) GetID() string {
	if sdtth == nil {
		return ""
	}
	return sdtth.ID
}

// GetInstruction returns the value of the field Instruction.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtth *SpecimenDefinitionTypeTestedHandling) GetInstruction() *fhir.String {
	if sdtth == nil {
		return nil
	}
	return sdtth.Instruction
}

// GetMaxDuration returns the value of the field MaxDuration.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtth *SpecimenDefinitionTypeTestedHandling) GetMaxDuration() *fhir.Duration {
	if sdtth == nil {
		return nil
	}
	return sdtth.MaxDuration
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtth *SpecimenDefinitionTypeTestedHandling) GetModifierExtension() []*fhir.Extension {
	if sdtth == nil {
		return nil
	}
	return sdtth.ModifierExtension
}

// GetTemperatureQualifier returns the value of the field TemperatureQualifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtth *SpecimenDefinitionTypeTestedHandling) GetTemperatureQualifier() *fhir.CodeableConcept {
	if sdtth == nil {
		return nil
	}
	return sdtth.TemperatureQualifier
}

// GetTemperatureRange returns the value of the field TemperatureRange.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sdtth *SpecimenDefinitionTypeTestedHandling) GetTemperatureRange() *fhir.Range {
	if sdtth == nil {
		return nil
	}
	return sdtth.TemperatureRange
}
