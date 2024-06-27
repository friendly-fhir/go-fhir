// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package substanceprotein

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// A SubstanceProtein is defined as a single unit of a linear amino acid
// sequence, or a combination of subunits that are either covalently linked or
// have a defined invariant stoichiometric relationship. This includes all
// synthetic, recombinant and purified SubstanceProteins of defined sequence,
// whether the use is therapeutic or prophylactic. This set of elements will be
// used to describe albumins, coagulation factors, cytokines, growth factors,
// peptide/SubstanceProtein hormones, enzymes, toxins, toxoids, recombinant
// vaccines, and immunomodulators.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/SubstanceProtein
//   - Source File: StructureDefinition-SubstanceProtein.json
type SubstanceProtein struct {

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The disulphide bond between two cysteine residues either on the same subunit
	// or on two different subunits shall be described. The position of the
	// disulfide bonds in the SubstanceProtein shall be listed in increasing order
	// of subunit number and position within subunit followed by the abbreviation
	// of the amino acids involved. The disulfide linkage positions shall actually
	// contain the amino acid Cysteine at the respective positions.
	DisulfideLinkage []*fhir.String `fhirpath:"disulfideLinkage"`

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

	// Number of linear sequences of amino acids linked through peptide bonds. The
	// number of subunits constituting the SubstanceProtein shall be described. It
	// is possible that the number of subunits can be variable.
	NumberOfSubunits *fhir.Integer `fhirpath:"numberOfSubunits"`

	// The SubstanceProtein descriptive elements will only be used when a complete
	// or partial amino acid sequence is available or derivable from a nucleic acid
	// sequence.
	SequenceType *fhir.CodeableConcept `fhirpath:"sequenceType"`

	// This subclause refers to the description of each subunit constituting the
	// SubstanceProtein. A subunit is a linear sequence of amino acids linked
	// through peptide bonds. The Subunit information shall be provided when the
	// finished SubstanceProtein is a complex of multiple sequences; subunits are
	// not used to delineate domains within a single sequence. Subunits are listed
	// in order of decreasing length; sequences of the same length will be ordered
	// by decreasing molecular weight; subunits that have identical sequences will
	// be repeated multiple times.
	Subunit []*SubstanceProteinSubunit `fhirpath:"subunit"`

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

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SubstanceProtein) GetContained() []fhir.Resource {
	if sp == nil {
		return nil
	}
	return sp.Contained
}

// GetDisulfideLinkage returns the value of the field DisulfideLinkage.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SubstanceProtein) GetDisulfideLinkage() []*fhir.String {
	if sp == nil {
		return nil
	}
	return sp.DisulfideLinkage
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SubstanceProtein) GetExtension() []*fhir.Extension {
	if sp == nil {
		return nil
	}
	return sp.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SubstanceProtein) GetID() string {
	if sp == nil {
		return ""
	}
	return sp.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SubstanceProtein) GetImplicitRules() *fhir.URI {
	if sp == nil {
		return nil
	}
	return sp.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SubstanceProtein) GetLanguage() *fhir.Code {
	if sp == nil {
		return nil
	}
	return sp.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SubstanceProtein) GetMeta() *fhir.Meta {
	if sp == nil {
		return nil
	}
	return sp.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SubstanceProtein) GetModifierExtension() []*fhir.Extension {
	if sp == nil {
		return nil
	}
	return sp.ModifierExtension
}

// GetNumberOfSubunits returns the value of the field NumberOfSubunits.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SubstanceProtein) GetNumberOfSubunits() *fhir.Integer {
	if sp == nil {
		return nil
	}
	return sp.NumberOfSubunits
}

// GetSequenceType returns the value of the field SequenceType.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SubstanceProtein) GetSequenceType() *fhir.CodeableConcept {
	if sp == nil {
		return nil
	}
	return sp.SequenceType
}

// GetSubunit returns the value of the field Subunit.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SubstanceProtein) GetSubunit() []*SubstanceProteinSubunit {
	if sp == nil {
		return nil
	}
	return sp.Subunit
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sp *SubstanceProtein) GetText() *fhir.Narrative {
	if sp == nil {
		return nil
	}
	return sp.Text
}

// This subclause refers to the description of each subunit constituting the
// SubstanceProtein. A subunit is a linear sequence of amino acids linked
// through peptide bonds. The Subunit information shall be provided when the
// finished SubstanceProtein is a complex of multiple sequences; subunits are
// not used to delineate domains within a single sequence. Subunits are listed
// in order of decreasing length; sequences of the same length will be ordered
// by decreasing molecular weight; subunits that have identical sequences will
// be repeated multiple times// This subclause refers to the description of each subunit constituting the
// SubstanceProtein. A subunit is a linear sequence of amino acids linked
// through peptide bonds. The Subunit information shall be provided when the
// finished SubstanceProtein is a complex of multiple sequences; subunits are
// not used to delineate domains within a single sequence. Subunits are listed
// in order of decreasing length; sequences of the same length will be ordered
// by decreasing molecular weight; subunits that have identical sequences will
// be repeated multiple times.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-SubstanceProtein.json
type SubstanceProteinSubunit struct {

	// The modification at the C-terminal shall be specified.
	CTerminalModification *fhir.String `fhirpath:"cTerminalModification"`

	// Unique identifier for molecular fragment modification based on the ISO 11238
	// Substance ID.
	CTerminalModificationID *fhir.Identifier `fhirpath:"cTerminalModificationId"`

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

	// Length of linear sequences of amino acids contained in the subunit.
	Length *fhir.Integer `fhirpath:"length"`

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

	// The name of the fragment modified at the N-terminal of the SubstanceProtein
	// shall be specified.
	NTerminalModification *fhir.String `fhirpath:"nTerminalModification"`

	// Unique identifier for molecular fragment modification based on the ISO 11238
	// Substance ID.
	NTerminalModificationID *fhir.Identifier `fhirpath:"nTerminalModificationId"`

	// The sequence information shall be provided enumerating the amino acids from
	// N- to C-terminal end using standard single-letter amino acid codes.
	// Uppercase shall be used for L-amino acids and lowercase for D-amino acids.
	// Transcribed SubstanceProteins will always be described using the translated
	// sequence; for synthetic peptide containing amino acids that are not
	// represented with a single letter code an X should be used within the
	// sequence. The modified amino acids will be distinguished by their position
	// in the sequence.
	Sequence *fhir.String `fhirpath:"sequence"`

	// The sequence information shall be provided enumerating the amino acids from
	// N- to C-terminal end using standard single-letter amino acid codes.
	// Uppercase shall be used for L-amino acids and lowercase for D-amino acids.
	// Transcribed SubstanceProteins will always be described using the translated
	// sequence; for synthetic peptide containing amino acids that are not
	// represented with a single letter code an X should be used within the
	// sequence. The modified amino acids will be distinguished by their position
	// in the sequence.
	SequenceAttachment *fhir.Attachment `fhirpath:"sequenceAttachment"`

	// Index of primary sequences of amino acids linked through peptide bonds in
	// order of decreasing length. Sequences of the same length will be ordered by
	// molecular weight. Subunits that have identical sequences will be repeated
	// and have sequential subscripts.
	Subunit *fhir.Integer `fhirpath:"subunit"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetCTerminalModification returns the value of the field CTerminalModification.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sps *SubstanceProteinSubunit) GetCTerminalModification() *fhir.String {
	if sps == nil {
		return nil
	}
	return sps.CTerminalModification
}

// GetCTerminalModificationID returns the value of the field CTerminalModificationID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sps *SubstanceProteinSubunit) GetCTerminalModificationID() *fhir.Identifier {
	if sps == nil {
		return nil
	}
	return sps.CTerminalModificationID
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sps *SubstanceProteinSubunit) GetExtension() []*fhir.Extension {
	if sps == nil {
		return nil
	}
	return sps.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sps *SubstanceProteinSubunit) GetID() string {
	if sps == nil {
		return ""
	}
	return sps.ID
}

// GetLength returns the value of the field Length.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sps *SubstanceProteinSubunit) GetLength() *fhir.Integer {
	if sps == nil {
		return nil
	}
	return sps.Length
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sps *SubstanceProteinSubunit) GetModifierExtension() []*fhir.Extension {
	if sps == nil {
		return nil
	}
	return sps.ModifierExtension
}

// GetNTerminalModification returns the value of the field NTerminalModification.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sps *SubstanceProteinSubunit) GetNTerminalModification() *fhir.String {
	if sps == nil {
		return nil
	}
	return sps.NTerminalModification
}

// GetNTerminalModificationID returns the value of the field NTerminalModificationID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sps *SubstanceProteinSubunit) GetNTerminalModificationID() *fhir.Identifier {
	if sps == nil {
		return nil
	}
	return sps.NTerminalModificationID
}

// GetSequence returns the value of the field Sequence.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sps *SubstanceProteinSubunit) GetSequence() *fhir.String {
	if sps == nil {
		return nil
	}
	return sps.Sequence
}

// GetSequenceAttachment returns the value of the field SequenceAttachment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sps *SubstanceProteinSubunit) GetSequenceAttachment() *fhir.Attachment {
	if sps == nil {
		return nil
	}
	return sps.SequenceAttachment
}

// GetSubunit returns the value of the field Subunit.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (sps *SubstanceProteinSubunit) GetSubunit() *fhir.Integer {
	if sps == nil {
		return nil
	}
	return sps.Subunit
}

func (sp *SubstanceProtein) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (sp *SubstanceProtein) UnmarshalJSON(data []byte) error {
	var raw struct {
		Contained        []fhir.Resource   `json:"contained"`
		DisulfideLinkage []*fhir.String    `json:"disulfideLinkage"`
		Extension        []*fhir.Extension `json:"extension"`

		ID                string                     `json:"id"`
		ImplicitRules     *fhir.URI                  `json:"implicitRules"`
		Language          *fhir.Code                 `json:"language"`
		Meta              *fhir.Meta                 `json:"meta"`
		ModifierExtension []*fhir.Extension          `json:"modifierExtension"`
		NumberOfSubunits  *fhir.Integer              `json:"numberOfSubunits"`
		SequenceType      *fhir.CodeableConcept      `json:"sequenceType"`
		Subunit           []*SubstanceProteinSubunit `json:"subunit"`
		Text              *fhir.Narrative            `json:"text"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	sp.Contained = raw.Contained
	sp.DisulfideLinkage = raw.DisulfideLinkage
	sp.Extension = raw.Extension
	sp.ID = raw.ID
	sp.ImplicitRules = raw.ImplicitRules
	sp.Language = raw.Language
	sp.Meta = raw.Meta
	sp.ModifierExtension = raw.ModifierExtension
	sp.NumberOfSubunits = raw.NumberOfSubunits
	sp.SequenceType = raw.SequenceType
	sp.Subunit = raw.Subunit
	sp.Text = raw.Text
	return nil
}

var _ json.Marshaler = (*SubstanceProtein)(nil)
var _ json.Unmarshaler = (*SubstanceProtein)(nil)

func (sps *SubstanceProteinSubunit) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (sps *SubstanceProteinSubunit) UnmarshalJSON(data []byte) error {
	var raw struct {
		CTerminalModification   *fhir.String      `json:"cTerminalModification"`
		CTerminalModificationID *fhir.Identifier  `json:"cTerminalModificationId"`
		Extension               []*fhir.Extension `json:"extension"`

		ID                      string            `json:"id"`
		Length                  *fhir.Integer     `json:"length"`
		ModifierExtension       []*fhir.Extension `json:"modifierExtension"`
		NTerminalModification   *fhir.String      `json:"nTerminalModification"`
		NTerminalModificationID *fhir.Identifier  `json:"nTerminalModificationId"`
		Sequence                *fhir.String      `json:"sequence"`
		SequenceAttachment      *fhir.Attachment  `json:"sequenceAttachment"`
		Subunit                 *fhir.Integer     `json:"subunit"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	sps.CTerminalModification = raw.CTerminalModification
	sps.CTerminalModificationID = raw.CTerminalModificationID
	sps.Extension = raw.Extension
	sps.ID = raw.ID
	sps.Length = raw.Length
	sps.ModifierExtension = raw.ModifierExtension
	sps.NTerminalModification = raw.NTerminalModification
	sps.NTerminalModificationID = raw.NTerminalModificationID
	sps.Sequence = raw.Sequence
	sps.SequenceAttachment = raw.SequenceAttachment
	sps.Subunit = raw.Subunit
	return nil
}

var _ json.Marshaler = (*SubstanceProteinSubunit)(nil)
var _ json.Unmarshaler = (*SubstanceProteinSubunit)(nil)
