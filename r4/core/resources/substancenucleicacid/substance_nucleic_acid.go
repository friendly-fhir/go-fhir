// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package substancenucleicacid

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// Nucleic acids are defined by three distinct elements: the base, sugar and
// linkage. Individual substance/moiety IDs will be created for each of these
// elements. The nucleotide sequence will be always entered in the 5’-3’
// direction.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/SubstanceNucleicAcid
//   - Source File: StructureDefinition-SubstanceNucleicAcid.json
type SubstanceNucleicAcid struct {

	// The area of hybridisation shall be described if applicable for double
	// stranded RNA or DNA. The number associated with the subunit followed by the
	// number associated to the residue shall be specified in increasing order. The
	// underscore “” shall be used as separator as follows: “Subunitnumber
	// Residue”.
	AreaOfHybridisation *fhir.String `fhirpath:"areaOfHybridisation"`

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

	// The number of linear sequences of nucleotides linked through phosphodiester
	// bonds shall be described. Subunits would be strands of nucleic acids that
	// are tightly associated typically through Watson-Crick base pairing. NOTE: If
	// not specified in the reference source, the assumption is that there is 1
	// subunit.
	NumberOfSubunits *fhir.Integer `fhirpath:"numberOfSubunits"`

	// (TBC).
	OligoNucleotideType *fhir.CodeableConcept `fhirpath:"oligoNucleotideType"`

	// The type of the sequence shall be specified based on a controlled
	// vocabulary.
	SequenceType *fhir.CodeableConcept `fhirpath:"sequenceType"`

	// Subunits are listed in order of decreasing length; sequences of the same
	// length will be ordered by molecular weight; subunits that have identical
	// sequences will be repeated multiple times.
	Subunit []*SubstanceNucleicAcidSubunit `fhirpath:"subunit"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetAreaOfHybridisation returns the value of the field AreaOfHybridisation.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetAreaOfHybridisation() *fhir.String {
	if sna == nil {
		return nil
	}
	return sna.AreaOfHybridisation
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetContained() []fhir.Resource {
	if sna == nil {
		return nil
	}
	return sna.Contained
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetExtension() []*fhir.Extension {
	if sna == nil {
		return nil
	}
	return sna.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetID() string {
	if sna == nil {
		return ""
	}
	return sna.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetImplicitRules() *fhir.URI {
	if sna == nil {
		return nil
	}
	return sna.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetLanguage() *fhir.Code {
	if sna == nil {
		return nil
	}
	return sna.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetMeta() *fhir.Meta {
	if sna == nil {
		return nil
	}
	return sna.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetModifierExtension() []*fhir.Extension {
	if sna == nil {
		return nil
	}
	return sna.ModifierExtension
}

// GetNumberOfSubunits returns the value of the field NumberOfSubunits.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetNumberOfSubunits() *fhir.Integer {
	if sna == nil {
		return nil
	}
	return sna.NumberOfSubunits
}

// GetOligoNucleotideType returns the value of the field OligoNucleotideType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetOligoNucleotideType() *fhir.CodeableConcept {
	if sna == nil {
		return nil
	}
	return sna.OligoNucleotideType
}

// GetSequenceType returns the value of the field SequenceType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetSequenceType() *fhir.CodeableConcept {
	if sna == nil {
		return nil
	}
	return sna.SequenceType
}

// GetSubunit returns the value of the field Subunit.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetSubunit() []*SubstanceNucleicAcidSubunit {
	if sna == nil {
		return nil
	}
	return sna.Subunit
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sna *SubstanceNucleicAcid) GetText() *fhir.Narrative {
	if sna == nil {
		return nil
	}
	return sna.Text
}

// Subunits are listed in order of decreasing length; sequences of the same
// length will be ordered by molecular weight; subunits that have identical
// sequences will be repeated multiple times// Subunits are listed in order of decreasing length; sequences of the same
// length will be ordered by molecular weight; subunits that have identical
// sequences will be repeated multiple times.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-SubstanceNucleicAcid.json
type SubstanceNucleicAcidSubunit struct {

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*fhir.Extension `fhirpath:"extension"`

	// The nucleotide present at the 5’ terminal shall be specified based on a
	// controlled vocabulary. Since the sequence is represented from the 5' to the
	// 3' end, the 5’ prime nucleotide is the letter at the first position in the
	// sequence. A separate representation would be redundant.
	FivePrime *fhir.CodeableConcept `fhirpath:"fivePrime"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	ID string `fhirpath:"id"`

	// The length of the sequence shall be captured.
	Length *fhir.Integer `fhirpath:"length"`

	// The linkages between sugar residues will also be captured.
	Linkage []*SubstanceNucleicAcidSubunitLinkage `fhirpath:"linkage"`

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

	// Actual nucleotide sequence notation from 5' to 3' end using standard single
	// letter codes. In addition to the base sequence, sugar and type of phosphate
	// or non-phosphate linkage should also be captured.
	Sequence *fhir.String `fhirpath:"sequence"`

	// (TBC).
	SequenceAttachment *fhir.Attachment `fhirpath:"sequenceAttachment"`

	// Index of linear sequences of nucleic acids in order of decreasing length.
	// Sequences of the same length will be ordered by molecular weight. Subunits
	// that have identical sequences will be repeated and have sequential
	// subscripts.
	Subunit *fhir.Integer `fhirpath:"subunit"`

	// 5.3.6.8.1 Sugar ID (Mandatory).
	Sugar []*SubstanceNucleicAcidSubunitSugar `fhirpath:"sugar"`

	// The nucleotide present at the 3’ terminal shall be specified based on a
	// controlled vocabulary. Since the sequence is represented from the 5' to the
	// 3' end, the 5’ prime nucleotide is the letter at the last position in the
	// sequence. A separate representation would be redundant.
	ThreePrime *fhir.CodeableConcept `fhirpath:"threePrime"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snas *SubstanceNucleicAcidSubunit) GetExtension() []*fhir.Extension {
	if snas == nil {
		return nil
	}
	return snas.Extension
}

// GetFivePrime returns the value of the field FivePrime.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snas *SubstanceNucleicAcidSubunit) GetFivePrime() *fhir.CodeableConcept {
	if snas == nil {
		return nil
	}
	return snas.FivePrime
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snas *SubstanceNucleicAcidSubunit) GetID() string {
	if snas == nil {
		return ""
	}
	return snas.ID
}

// GetLength returns the value of the field Length.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snas *SubstanceNucleicAcidSubunit) GetLength() *fhir.Integer {
	if snas == nil {
		return nil
	}
	return snas.Length
}

// GetLinkage returns the value of the field Linkage.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snas *SubstanceNucleicAcidSubunit) GetLinkage() []*SubstanceNucleicAcidSubunitLinkage {
	if snas == nil {
		return nil
	}
	return snas.Linkage
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snas *SubstanceNucleicAcidSubunit) GetModifierExtension() []*fhir.Extension {
	if snas == nil {
		return nil
	}
	return snas.ModifierExtension
}

// GetSequence returns the value of the field Sequence.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snas *SubstanceNucleicAcidSubunit) GetSequence() *fhir.String {
	if snas == nil {
		return nil
	}
	return snas.Sequence
}

// GetSequenceAttachment returns the value of the field SequenceAttachment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snas *SubstanceNucleicAcidSubunit) GetSequenceAttachment() *fhir.Attachment {
	if snas == nil {
		return nil
	}
	return snas.SequenceAttachment
}

// GetSubunit returns the value of the field Subunit.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snas *SubstanceNucleicAcidSubunit) GetSubunit() *fhir.Integer {
	if snas == nil {
		return nil
	}
	return snas.Subunit
}

// GetSugar returns the value of the field Sugar.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snas *SubstanceNucleicAcidSubunit) GetSugar() []*SubstanceNucleicAcidSubunitSugar {
	if snas == nil {
		return nil
	}
	return snas.Sugar
}

// GetThreePrime returns the value of the field ThreePrime.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snas *SubstanceNucleicAcidSubunit) GetThreePrime() *fhir.CodeableConcept {
	if snas == nil {
		return nil
	}
	return snas.ThreePrime
}

// The linkages between sugar residues will also be captured// The linkages between sugar residues will also be captured.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-SubstanceNucleicAcid.json
type SubstanceNucleicAcidSubunitLinkage struct {

	// The entity that links the sugar residues together should also be captured
	// for nearly all naturally occurring nucleic acid the linkage is a phosphate
	// group. For many synthetic oligonucleotides phosphorothioate linkages are
	// often seen. Linkage connectivity is assumed to be 3’-5’. If the linkage
	// is either 3’-3’ or 5’-5’ this should be specified.
	Connectivity *fhir.String `fhirpath:"connectivity"`

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

	// Each linkage will be registered as a fragment and have an ID.
	Identifier *fhir.Identifier `fhirpath:"identifier"`

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

	// Each linkage will be registered as a fragment and have at least one name. A
	// single name shall be assigned to each linkage.
	Name *fhir.String `fhirpath:"name"`

	// Residues shall be captured as described in 5.3.6.8.3.
	ResidueSite *fhir.String `fhirpath:"residueSite"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetConnectivity returns the value of the field Connectivity.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snasl *SubstanceNucleicAcidSubunitLinkage) GetConnectivity() *fhir.String {
	if snasl == nil {
		return nil
	}
	return snasl.Connectivity
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snasl *SubstanceNucleicAcidSubunitLinkage) GetExtension() []*fhir.Extension {
	if snasl == nil {
		return nil
	}
	return snasl.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snasl *SubstanceNucleicAcidSubunitLinkage) GetID() string {
	if snasl == nil {
		return ""
	}
	return snasl.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snasl *SubstanceNucleicAcidSubunitLinkage) GetIdentifier() *fhir.Identifier {
	if snasl == nil {
		return nil
	}
	return snasl.Identifier
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snasl *SubstanceNucleicAcidSubunitLinkage) GetModifierExtension() []*fhir.Extension {
	if snasl == nil {
		return nil
	}
	return snasl.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snasl *SubstanceNucleicAcidSubunitLinkage) GetName() *fhir.String {
	if snasl == nil {
		return nil
	}
	return snasl.Name
}

// GetResidueSite returns the value of the field ResidueSite.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snasl *SubstanceNucleicAcidSubunitLinkage) GetResidueSite() *fhir.String {
	if snasl == nil {
		return nil
	}
	return snasl.ResidueSite
}

// 5.3.6.8.1 Sugar ID (Mandatory)// 5.3.6.8.1 Sugar ID (Mandatory).
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-SubstanceNucleicAcid.json
type SubstanceNucleicAcidSubunitSugar struct {

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

	// The Substance ID of the sugar or sugar-like component that make up the
	// nucleotide.
	Identifier *fhir.Identifier `fhirpath:"identifier"`

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

	// The name of the sugar or sugar-like component that make up the nucleotide.
	Name *fhir.String `fhirpath:"name"`

	// The residues that contain a given sugar will be captured. The order of given
	// residues will be captured in the 5‘-3‘direction consistent with the base
	// sequences listed above.
	ResidueSite *fhir.String `fhirpath:"residueSite"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snass *SubstanceNucleicAcidSubunitSugar) GetExtension() []*fhir.Extension {
	if snass == nil {
		return nil
	}
	return snass.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snass *SubstanceNucleicAcidSubunitSugar) GetID() string {
	if snass == nil {
		return ""
	}
	return snass.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snass *SubstanceNucleicAcidSubunitSugar) GetIdentifier() *fhir.Identifier {
	if snass == nil {
		return nil
	}
	return snass.Identifier
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snass *SubstanceNucleicAcidSubunitSugar) GetModifierExtension() []*fhir.Extension {
	if snass == nil {
		return nil
	}
	return snass.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snass *SubstanceNucleicAcidSubunitSugar) GetName() *fhir.String {
	if snass == nil {
		return nil
	}
	return snass.Name
}

// GetResidueSite returns the value of the field ResidueSite.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (snass *SubstanceNucleicAcidSubunitSugar) GetResidueSite() *fhir.String {
	if snass == nil {
		return nil
	}
	return snass.ResidueSite
}

func (sna *SubstanceNucleicAcid) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (sna *SubstanceNucleicAcid) UnmarshalJSON(data []byte) error {
	var raw struct {
		AreaOfHybridisation *fhir.String      `json:"areaOfHybridisation"`
		Contained           []fhir.Resource   `json:"contained"`
		Extension           []*fhir.Extension `json:"extension"`

		ID                  string                         `json:"id"`
		ImplicitRules       *fhir.URI                      `json:"implicitRules"`
		Language            *fhir.Code                     `json:"language"`
		Meta                *fhir.Meta                     `json:"meta"`
		ModifierExtension   []*fhir.Extension              `json:"modifierExtension"`
		NumberOfSubunits    *fhir.Integer                  `json:"numberOfSubunits"`
		OligoNucleotideType *fhir.CodeableConcept          `json:"oligoNucleotideType"`
		SequenceType        *fhir.CodeableConcept          `json:"sequenceType"`
		Subunit             []*SubstanceNucleicAcidSubunit `json:"subunit"`
		Text                *fhir.Narrative                `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	sna.AreaOfHybridisation = raw.AreaOfHybridisation
	sna.Contained = raw.Contained
	sna.Extension = raw.Extension
	sna.ID = raw.ID
	sna.ImplicitRules = raw.ImplicitRules
	sna.Language = raw.Language
	sna.Meta = raw.Meta
	sna.ModifierExtension = raw.ModifierExtension
	sna.NumberOfSubunits = raw.NumberOfSubunits
	sna.OligoNucleotideType = raw.OligoNucleotideType
	sna.SequenceType = raw.SequenceType
	sna.Subunit = raw.Subunit
	sna.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*SubstanceNucleicAcid)(nil)
var _ json.Unmarshaler = (*SubstanceNucleicAcid)(nil)

func (snas *SubstanceNucleicAcidSubunit) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (snas *SubstanceNucleicAcidSubunit) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension     `json:"extension"`
		FivePrime *fhir.CodeableConcept `json:"fivePrime"`

		ID                 string                                `json:"id"`
		Length             *fhir.Integer                         `json:"length"`
		Linkage            []*SubstanceNucleicAcidSubunitLinkage `json:"linkage"`
		ModifierExtension  []*fhir.Extension                     `json:"modifierExtension"`
		Sequence           *fhir.String                          `json:"sequence"`
		SequenceAttachment *fhir.Attachment                      `json:"sequenceAttachment"`
		Subunit            *fhir.Integer                         `json:"subunit"`
		Sugar              []*SubstanceNucleicAcidSubunitSugar   `json:"sugar"`
		ThreePrime         *fhir.CodeableConcept                 `json:"threePrime"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	snas.Extension = raw.Extension
	snas.FivePrime = raw.FivePrime
	snas.ID = raw.ID
	snas.Length = raw.Length
	snas.Linkage = raw.Linkage
	snas.ModifierExtension = raw.ModifierExtension
	snas.Sequence = raw.Sequence
	snas.SequenceAttachment = raw.SequenceAttachment
	snas.Subunit = raw.Subunit
	snas.Sugar = raw.Sugar
	snas.ThreePrime = raw.ThreePrime
	return nil
}

var _ json.Marshaler = (*SubstanceNucleicAcidSubunit)(nil)
var _ json.Unmarshaler = (*SubstanceNucleicAcidSubunit)(nil)

func (snasl *SubstanceNucleicAcidSubunitLinkage) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (snasl *SubstanceNucleicAcidSubunitLinkage) UnmarshalJSON(data []byte) error {
	var raw struct {
		Connectivity *fhir.String      `json:"connectivity"`
		Extension    []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		Identifier        *fhir.Identifier  `json:"identifier"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		Name              *fhir.String      `json:"name"`
		ResidueSite       *fhir.String      `json:"residueSite"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	snasl.Connectivity = raw.Connectivity
	snasl.Extension = raw.Extension
	snasl.ID = raw.ID
	snasl.Identifier = raw.Identifier
	snasl.ModifierExtension = raw.ModifierExtension
	snasl.Name = raw.Name
	snasl.ResidueSite = raw.ResidueSite
	return nil
}

var _ json.Marshaler = (*SubstanceNucleicAcidSubunitLinkage)(nil)
var _ json.Unmarshaler = (*SubstanceNucleicAcidSubunitLinkage)(nil)

func (snass *SubstanceNucleicAcidSubunitSugar) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (snass *SubstanceNucleicAcidSubunitSugar) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                string            `json:"id"`
		Identifier        *fhir.Identifier  `json:"identifier"`
		ModifierExtension []*fhir.Extension `json:"modifierExtension"`
		Name              *fhir.String      `json:"name"`
		ResidueSite       *fhir.String      `json:"residueSite"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	snass.Extension = raw.Extension
	snass.ID = raw.ID
	snass.Identifier = raw.Identifier
	snass.ModifierExtension = raw.ModifierExtension
	snass.Name = raw.Name
	snass.ResidueSite = raw.ResidueSite
	return nil
}

var _ json.Marshaler = (*SubstanceNucleicAcidSubunitSugar)(nil)
var _ json.Unmarshaler = (*SubstanceNucleicAcidSubunitSugar)(nil)
