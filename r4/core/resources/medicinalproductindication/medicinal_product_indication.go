// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package medicinalproductindication

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Indication for the Medicinal Product.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/MedicinalProductIndication
//   - Source File: StructureDefinition-MedicinalProductIndication.json
type MedicinalProductIndication struct {

	// Comorbidity (concurrent condition) or co-infection as part of the
	// indication.
	Comorbidity []*fhir.CodeableConcept `fhirpath:"comorbidity"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The status of the disease or symptom for which the indication applies.
	DiseaseStatus *fhir.CodeableConcept `fhirpath:"diseaseStatus"`

	// The disease, symptom or procedure that is the indication for treatment.
	DiseaseSymptomProcedure *fhir.CodeableConcept `fhirpath:"diseaseSymptomProcedure"`

	// Timing or duration information as part of the indication.
	Duration *fhir.Quantity `fhirpath:"duration"`

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

	// The intended effect, aim or strategy to be achieved by the indication.
	IntendedEffect *fhir.CodeableConcept `fhirpath:"intendedEffect"`

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

	// Information about the use of the medicinal product in relation to other
	// therapies described as part of the indication.
	OtherTherapy []*MedicinalProductIndicationOtherTherapy `fhirpath:"otherTherapy"`

	// The population group to which this applies.
	Population []*fhir.Population `fhirpath:"population"`

	// The medication for which this is an indication.
	Subject []*fhir.Reference `fhirpath:"subject"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	// Describe the undesirable effects of the medicinal product.
	UndesirableEffect []*fhir.Reference `fhirpath:"undesirableEffect"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetComorbidity returns the value of the field Comorbidity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetComorbidity() []*fhir.CodeableConcept {
	if mpi == nil {
		return nil
	}
	return mpi.Comorbidity
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetContained() []fhir.Resource {
	if mpi == nil {
		return nil
	}
	return mpi.Contained
}

// GetDiseaseStatus returns the value of the field DiseaseStatus.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetDiseaseStatus() *fhir.CodeableConcept {
	if mpi == nil {
		return nil
	}
	return mpi.DiseaseStatus
}

// GetDiseaseSymptomProcedure returns the value of the field DiseaseSymptomProcedure.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetDiseaseSymptomProcedure() *fhir.CodeableConcept {
	if mpi == nil {
		return nil
	}
	return mpi.DiseaseSymptomProcedure
}

// GetDuration returns the value of the field Duration.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetDuration() *fhir.Quantity {
	if mpi == nil {
		return nil
	}
	return mpi.Duration
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetExtension() []*fhir.Extension {
	if mpi == nil {
		return nil
	}
	return mpi.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetID() string {
	if mpi == nil {
		return ""
	}
	return mpi.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetImplicitRules() *fhir.URI {
	if mpi == nil {
		return nil
	}
	return mpi.ImplicitRules
}

// GetIntendedEffect returns the value of the field IntendedEffect.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetIntendedEffect() *fhir.CodeableConcept {
	if mpi == nil {
		return nil
	}
	return mpi.IntendedEffect
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetLanguage() *fhir.Code {
	if mpi == nil {
		return nil
	}
	return mpi.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetMeta() *fhir.Meta {
	if mpi == nil {
		return nil
	}
	return mpi.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetModifierExtension() []*fhir.Extension {
	if mpi == nil {
		return nil
	}
	return mpi.ModifierExtension
}

// GetOtherTherapy returns the value of the field OtherTherapy.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetOtherTherapy() []*MedicinalProductIndicationOtherTherapy {
	if mpi == nil {
		return nil
	}
	return mpi.OtherTherapy
}

// GetPopulation returns the value of the field Population.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetPopulation() []*fhir.Population {
	if mpi == nil {
		return nil
	}
	return mpi.Population
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetSubject() []*fhir.Reference {
	if mpi == nil {
		return nil
	}
	return mpi.Subject
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetText() *fhir.Narrative {
	if mpi == nil {
		return nil
	}
	return mpi.Text
}

// GetUndesirableEffect returns the value of the field UndesirableEffect.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpi *MedicinalProductIndication) GetUndesirableEffect() []*fhir.Reference {
	if mpi == nil {
		return nil
	}
	return mpi.UndesirableEffect
}

// Information about the use of the medicinal product in relation to other
// therapies described as part of the indication// Information about the use of the medicinal product in relation to other
// therapies described as part of the indication.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MedicinalProductIndication.json
type MedicinalProductIndicationOtherTherapy struct {

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

	// Reference to a specific medication (active substance, medicinal product or
	// class of products) as part of an indication or contraindication.
	Medication fhir.Element `fhirpath:"medication"`

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

	// The type of relationship between the medicinal product indication or
	// contraindication and another therapy.
	TherapyRelationshipType *fhir.CodeableConcept `fhirpath:"therapyRelationshipType"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpiot *MedicinalProductIndicationOtherTherapy) GetExtension() []*fhir.Extension {
	if mpiot == nil {
		return nil
	}
	return mpiot.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpiot *MedicinalProductIndicationOtherTherapy) GetID() string {
	if mpiot == nil {
		return ""
	}
	return mpiot.ID
}

// GetMedication returns the value of the field Medication.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpiot *MedicinalProductIndicationOtherTherapy) GetMedication() fhir.Element {
	if mpiot == nil {
		return nil
	}
	return mpiot.Medication
}

// GetMedicationCodeableConcept returns the value of the field Medication.
func (mpiot *MedicinalProductIndicationOtherTherapy) GetMedicationCodeableConcept() *fhir.CodeableConcept {
	if mpiot == nil {
		return nil
	}
	val, ok := mpiot.Medication.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetMedicationReference returns the value of the field Medication.
func (mpiot *MedicinalProductIndicationOtherTherapy) GetMedicationReference() *fhir.Reference {
	if mpiot == nil {
		return nil
	}
	val, ok := mpiot.Medication.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
} // GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpiot *MedicinalProductIndicationOtherTherapy) GetModifierExtension() []*fhir.Extension {
	if mpiot == nil {
		return nil
	}
	return mpiot.ModifierExtension
}

// GetTherapyRelationshipType returns the value of the field TherapyRelationshipType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpiot *MedicinalProductIndicationOtherTherapy) GetTherapyRelationshipType() *fhir.CodeableConcept {
	if mpiot == nil {
		return nil
	}
	return mpiot.TherapyRelationshipType
}
