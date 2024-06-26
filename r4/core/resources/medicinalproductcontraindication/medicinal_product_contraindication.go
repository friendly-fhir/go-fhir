// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package medicinalproductcontraindication

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// The clinical particulars - indications, contraindications etc. of a
// medicinal product, including for regulatory purposes.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/MedicinalProductContraindication
//   - Source File: StructureDefinition-MedicinalProductContraindication.json
type MedicinalProductContraindication struct {

	// A comorbidity (concurrent condition) or coinfection.
	Comorbidity []*fhir.CodeableConcept `fhirpath:"comorbidity"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The disease, symptom or procedure for the contraindication.
	Disease *fhir.CodeableConcept `fhirpath:"disease"`

	// The status of the disease or symptom for the contraindication.
	DiseaseStatus *fhir.CodeableConcept `fhirpath:"diseaseStatus"`

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
	OtherTherapy []*MedicinalProductContraindicationOtherTherapy `fhirpath:"otherTherapy"`

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

	// Information about the use of the medicinal product in relation to other
	// therapies as part of the indication.
	TherapeuticIndication []*fhir.Reference `fhirpath:"therapeuticIndication"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetComorbidity returns the value of the field Comorbidity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetComorbidity() []*fhir.CodeableConcept {
	if mpc == nil {
		return nil
	}
	return mpc.Comorbidity
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetContained() []fhir.Resource {
	if mpc == nil {
		return nil
	}
	return mpc.Contained
}

// GetDisease returns the value of the field Disease.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetDisease() *fhir.CodeableConcept {
	if mpc == nil {
		return nil
	}
	return mpc.Disease
}

// GetDiseaseStatus returns the value of the field DiseaseStatus.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetDiseaseStatus() *fhir.CodeableConcept {
	if mpc == nil {
		return nil
	}
	return mpc.DiseaseStatus
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetExtension() []*fhir.Extension {
	if mpc == nil {
		return nil
	}
	return mpc.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetID() string {
	if mpc == nil {
		return ""
	}
	return mpc.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetImplicitRules() *fhir.URI {
	if mpc == nil {
		return nil
	}
	return mpc.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetLanguage() *fhir.Code {
	if mpc == nil {
		return nil
	}
	return mpc.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetMeta() *fhir.Meta {
	if mpc == nil {
		return nil
	}
	return mpc.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetModifierExtension() []*fhir.Extension {
	if mpc == nil {
		return nil
	}
	return mpc.ModifierExtension
}

// GetOtherTherapy returns the value of the field OtherTherapy.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetOtherTherapy() []*MedicinalProductContraindicationOtherTherapy {
	if mpc == nil {
		return nil
	}
	return mpc.OtherTherapy
}

// GetPopulation returns the value of the field Population.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetPopulation() []*fhir.Population {
	if mpc == nil {
		return nil
	}
	return mpc.Population
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetSubject() []*fhir.Reference {
	if mpc == nil {
		return nil
	}
	return mpc.Subject
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetText() *fhir.Narrative {
	if mpc == nil {
		return nil
	}
	return mpc.Text
}

// GetTherapeuticIndication returns the value of the field TherapeuticIndication.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpc *MedicinalProductContraindication) GetTherapeuticIndication() []*fhir.Reference {
	if mpc == nil {
		return nil
	}
	return mpc.TherapeuticIndication
}

// Information about the use of the medicinal product in relation to other
// therapies described as part of the indication// Information about the use of the medicinal product in relation to other
// therapies described as part of the indication.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-MedicinalProductContraindication.json
type MedicinalProductContraindicationOtherTherapy struct {

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
func (mpcot *MedicinalProductContraindicationOtherTherapy) GetExtension() []*fhir.Extension {
	if mpcot == nil {
		return nil
	}
	return mpcot.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpcot *MedicinalProductContraindicationOtherTherapy) GetID() string {
	if mpcot == nil {
		return ""
	}
	return mpcot.ID
}

// GetMedication returns the value of the field Medication.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpcot *MedicinalProductContraindicationOtherTherapy) GetMedication() fhir.Element {
	if mpcot == nil {
		return nil
	}
	return mpcot.Medication
}

// GetMedicationCodeableConcept returns the value of the field Medication.
func (mpcot *MedicinalProductContraindicationOtherTherapy) GetMedicationCodeableConcept() *fhir.CodeableConcept {
	if mpcot == nil {
		return nil
	}
	val, ok := mpcot.Medication.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetMedicationReference returns the value of the field Medication.
func (mpcot *MedicinalProductContraindicationOtherTherapy) GetMedicationReference() *fhir.Reference {
	if mpcot == nil {
		return nil
	}
	val, ok := mpcot.Medication.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
} // GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpcot *MedicinalProductContraindicationOtherTherapy) GetModifierExtension() []*fhir.Extension {
	if mpcot == nil {
		return nil
	}
	return mpcot.ModifierExtension
}

// GetTherapyRelationshipType returns the value of the field TherapyRelationshipType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (mpcot *MedicinalProductContraindicationOtherTherapy) GetTherapyRelationshipType() *fhir.CodeableConcept {
	if mpcot == nil {
		return nil
	}
	return mpcot.TherapyRelationshipType
}

func (mpc *MedicinalProductContraindication) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (mpc *MedicinalProductContraindication) UnmarshalJSON(data []byte) error {
	var raw struct {
		Comorbidity   []*fhir.CodeableConcept `json:"comorbidity"`
		Contained     []fhir.Resource         `json:"contained"`
		Disease       *fhir.CodeableConcept   `json:"disease"`
		DiseaseStatus *fhir.CodeableConcept   `json:"diseaseStatus"`
		Extension     []*fhir.Extension       `json:"extension"`

		ID                    string                                          `json:"id"`
		ImplicitRules         *fhir.URI                                       `json:"implicitRules"`
		Language              *fhir.Code                                      `json:"language"`
		Meta                  *fhir.Meta                                      `json:"meta"`
		ModifierExtension     []*fhir.Extension                               `json:"modifierExtension"`
		OtherTherapy          []*MedicinalProductContraindicationOtherTherapy `json:"otherTherapy"`
		Population            []*fhir.Population                              `json:"population"`
		Subject               []*fhir.Reference                               `json:"subject"`
		Text                  *fhir.Narrative                                 `json:"text"`
		TherapeuticIndication []*fhir.Reference                               `json:"therapeuticIndication"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mpc.Comorbidity = raw.Comorbidity
	mpc.Contained = raw.Contained
	mpc.Disease = raw.Disease
	mpc.DiseaseStatus = raw.DiseaseStatus
	mpc.Extension = raw.Extension
	mpc.ID = raw.ID
	mpc.ImplicitRules = raw.ImplicitRules
	mpc.Language = raw.Language
	mpc.Meta = raw.Meta
	mpc.ModifierExtension = raw.ModifierExtension
	mpc.OtherTherapy = raw.OtherTherapy
	mpc.Population = raw.Population
	mpc.Subject = raw.Subject
	mpc.Text = raw.Text
	mpc.TherapeuticIndication = raw.TherapeuticIndication
	return nil
}

var _ json.Marshaler = (*MedicinalProductContraindication)(nil)
var _ json.Unmarshaler = (*MedicinalProductContraindication)(nil)

func (mpcot *MedicinalProductContraindicationOtherTherapy) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (mpcot *MedicinalProductContraindicationOtherTherapy) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                        string                `json:"id"`
		MedicationCodeableConcept *fhir.CodeableConcept `json:"medicationCodeableConcept"`
		MedicationReference       *fhir.Reference       `json:"medicationReference"`
		ModifierExtension         []*fhir.Extension     `json:"modifierExtension"`
		TherapyRelationshipType   *fhir.CodeableConcept `json:"therapyRelationshipType"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	mpcot.Extension = raw.Extension
	mpcot.ID = raw.ID
	mpcot.Medication, err = validate.SelectOneOf[fhir.Element]("MedicinalProductContraindication.otherTherapy.medication",
		raw.MedicationCodeableConcept,
		raw.MedicationReference)
	if err != nil {
		return err
	}
	mpcot.ModifierExtension = raw.ModifierExtension
	mpcot.TherapyRelationshipType = raw.TherapyRelationshipType
	return nil
}

var _ json.Marshaler = (*MedicinalProductContraindicationOtherTherapy)(nil)
var _ json.Unmarshaler = (*MedicinalProductContraindicationOtherTherapy)(nil)
