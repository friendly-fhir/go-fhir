// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package diagnosticreport

import (
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// The findings and interpretation of diagnostic tests performed on patients,
// groups of patients, devices, and locations, and/or specimens derived from
// these. The report includes clinical context such as requesting and provider
// information, and some mix of atomic results, images, textual and coded
// interpretations, and formatted representation of diagnostic reports.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/DiagnosticReport
//   - Source File: StructureDefinition-DiagnosticReport.json
type DiagnosticReport struct {

	// Details concerning a service requested.
	BasedOn []*fhir.Reference `fhirpath:"basedOn"`

	// A code that classifies the clinical discipline, department or diagnostic
	// service that created the report (e.g. cardiology, biochemistry, hematology,
	// MRI). This is used for searching, sorting and display purposes.
	Category []*fhir.CodeableConcept `fhirpath:"category"`

	// A code or name that describes this diagnostic report.
	Code *fhir.CodeableConcept `fhirpath:"code"`

	// Concise and clinically contextualized summary conclusion
	// (interpretation/impression) of the diagnostic report.
	Conclusion *fhir.String `fhirpath:"conclusion"`

	// One or more codes that represent the summary conclusion
	// (interpretation/impression) of the diagnostic report.
	ConclusionCode []*fhir.CodeableConcept `fhirpath:"conclusionCode"`

	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, and nor can
	// they have their own independent transaction scope.
	Contained []fhir.Resource `fhirpath:"contained"`

	// The time or time-period the observed values are related to. When the subject
	// of the report is a patient, this is usually either the time of the procedure
	// or of specimen collection(s), but very often the source of the date/time is
	// not known, only the date/time itself.
	Effective fhir.Element `fhirpath:"effective"`

	// The healthcare event (e.g. a patient and healthcare provider interaction)
	// which this DiagnosticReport is about.
	Encounter *fhir.Reference `fhirpath:"encounter"`

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

	// Identifiers assigned to this report by the performer or other systems.
	Identifier []*fhir.Identifier `fhirpath:"identifier"`

	// One or more links to full details of any imaging performed during the
	// diagnostic investigation. Typically, this is imaging performed by DICOM
	// enabled modalities, but this is not required. A fully enabled PACS viewer
	// can use this information to provide views of the source images.
	ImagingStudy []*fhir.Reference `fhirpath:"imagingStudy"`

	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content.
	// Often, this is a reference to an implementation guide that defines the
	// special rules along with other profiles etc.
	ImplicitRules *fhir.URI `fhirpath:"implicitRules"`

	// The date and time that this version of the report was made available to
	// providers, typically after the report was reviewed and verified.
	Issued *fhir.Instant `fhirpath:"issued"`

	// The base language in which the resource is written.
	Language *fhir.Code `fhirpath:"language"`

	// A list of key images associated with this report. The images are generally
	// created during the diagnostic process, and may be directly of the patient,
	// or of treated specimens (i.e. slides of interest).
	Media []*DiagnosticReportMedia `fhirpath:"media"`

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

	// The diagnostic service that is responsible for issuing the report.
	Performer []*fhir.Reference `fhirpath:"performer"`

	// Rich text representation of the entire result as issued by the diagnostic
	// service. Multiple formats are allowed but they SHALL be semantically
	// equivalent.
	PresentedForm []*fhir.Attachment `fhirpath:"presentedForm"`

	// [Observations](observation.html) that are part of this diagnostic report.
	Result []*fhir.Reference `fhirpath:"result"`

	// The practitioner or organization that is responsible for the report's
	// conclusions and interpretations.
	ResultsInterpreter []*fhir.Reference `fhirpath:"resultsInterpreter"`

	// Details about the specimens on which this diagnostic report is based.
	Specimen []*fhir.Reference `fhirpath:"specimen"`

	// The status of the diagnostic report.
	Status *fhir.Code `fhirpath:"status"`

	// The subject of the report. Usually, but not always, this is a patient.
	// However, diagnostic services also perform analyses on specimens collected
	// from a variety of other sources.
	Subject *fhir.Reference `fhirpath:"subject"`

	// A human-readable narrative that contains a summary of the resource and can
	// be used to represent the content of the resource to a human. The narrative
	// need not encode all the structured data, but is required to contain
	// sufficient detail to make it "clinically safe" for a human to just read the
	// narrative. Resource definitions may define what content should be
	// represented in the narrative to ensure clinical safety.
	Text *fhir.Narrative `fhirpath:"text"`

	profileimpl.BaseDiagnosticReport
	profileimpl.BaseDomainResource
	profileimpl.BaseResource
}

// GetBasedOn returns the value of the field BasedOn.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetBasedOn() []*fhir.Reference {
	if dr == nil {
		return nil
	}
	return dr.BasedOn
}

// GetCategory returns the value of the field Category.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetCategory() []*fhir.CodeableConcept {
	if dr == nil {
		return nil
	}
	return dr.Category
}

// GetCode returns the value of the field Code.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetCode() *fhir.CodeableConcept {
	if dr == nil {
		return nil
	}
	return dr.Code
}

// GetConclusion returns the value of the field Conclusion.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetConclusion() *fhir.String {
	if dr == nil {
		return nil
	}
	return dr.Conclusion
}

// GetConclusionCode returns the value of the field ConclusionCode.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetConclusionCode() []*fhir.CodeableConcept {
	if dr == nil {
		return nil
	}
	return dr.ConclusionCode
}

// GetContained returns the value of the field Contained.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetContained() []fhir.Resource {
	if dr == nil {
		return nil
	}
	return dr.Contained
}

// GetEffective returns the value of the field Effective.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetEffective() fhir.Element {
	if dr == nil {
		return nil
	}
	return dr.Effective
}

// GetEffectiveDateTime returns the value of the field Effective.
func (dr *DiagnosticReport) GetEffectiveDateTime() *fhir.DateTime {
	if dr == nil {
		return nil
	}
	val, ok := dr.Effective.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
}

// GetEffectivePeriod returns the value of the field Effective.
func (dr *DiagnosticReport) GetEffectivePeriod() *fhir.Period {
	if dr == nil {
		return nil
	}
	val, ok := dr.Effective.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
} // GetEncounter returns the value of the field Encounter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetEncounter() *fhir.Reference {
	if dr == nil {
		return nil
	}
	return dr.Encounter
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetExtension() []*fhir.Extension {
	if dr == nil {
		return nil
	}
	return dr.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetID() string {
	if dr == nil {
		return ""
	}
	return dr.ID
}

// GetIdentifier returns the value of the field Identifier.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetIdentifier() []*fhir.Identifier {
	if dr == nil {
		return nil
	}
	return dr.Identifier
}

// GetImagingStudy returns the value of the field ImagingStudy.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetImagingStudy() []*fhir.Reference {
	if dr == nil {
		return nil
	}
	return dr.ImagingStudy
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetImplicitRules() *fhir.URI {
	if dr == nil {
		return nil
	}
	return dr.ImplicitRules
}

// GetIssued returns the value of the field Issued.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetIssued() *fhir.Instant {
	if dr == nil {
		return nil
	}
	return dr.Issued
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetLanguage() *fhir.Code {
	if dr == nil {
		return nil
	}
	return dr.Language
}

// GetMedia returns the value of the field Media.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetMedia() []*DiagnosticReportMedia {
	if dr == nil {
		return nil
	}
	return dr.Media
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetMeta() *fhir.Meta {
	if dr == nil {
		return nil
	}
	return dr.Meta
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetModifierExtension() []*fhir.Extension {
	if dr == nil {
		return nil
	}
	return dr.ModifierExtension
}

// GetPerformer returns the value of the field Performer.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetPerformer() []*fhir.Reference {
	if dr == nil {
		return nil
	}
	return dr.Performer
}

// GetPresentedForm returns the value of the field PresentedForm.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetPresentedForm() []*fhir.Attachment {
	if dr == nil {
		return nil
	}
	return dr.PresentedForm
}

// GetResult returns the value of the field Result.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetResult() []*fhir.Reference {
	if dr == nil {
		return nil
	}
	return dr.Result
}

// GetResultsInterpreter returns the value of the field ResultsInterpreter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetResultsInterpreter() []*fhir.Reference {
	if dr == nil {
		return nil
	}
	return dr.ResultsInterpreter
}

// GetSpecimen returns the value of the field Specimen.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetSpecimen() []*fhir.Reference {
	if dr == nil {
		return nil
	}
	return dr.Specimen
}

// GetStatus returns the value of the field Status.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetStatus() *fhir.Code {
	if dr == nil {
		return nil
	}
	return dr.Status
}

// GetSubject returns the value of the field Subject.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetSubject() *fhir.Reference {
	if dr == nil {
		return nil
	}
	return dr.Subject
}

// GetText returns the value of the field Text.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (dr *DiagnosticReport) GetText() *fhir.Narrative {
	if dr == nil {
		return nil
	}
	return dr.Text
}

// Key images associated with this report// A list of key images associated with this report. The images are generally
// created during the diagnostic process, and may be directly of the patient,
// or of treated specimens (i.e. slides of interest).
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-DiagnosticReport.json
type DiagnosticReportMedia struct {

	// A comment about the image. Typically, this is used to provide an explanation
	// for why the image is included, or to draw the viewer's attention to
	// important features.
	Comment *fhir.String `fhirpath:"comment"`

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

	// Reference to the image source.
	Link *fhir.Reference `fhirpath:"link"`

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

// GetComment returns the value of the field Comment.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (drm *DiagnosticReportMedia) GetComment() *fhir.String {
	if drm == nil {
		return nil
	}
	return drm.Comment
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (drm *DiagnosticReportMedia) GetExtension() []*fhir.Extension {
	if drm == nil {
		return nil
	}
	return drm.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (drm *DiagnosticReportMedia) GetID() string {
	if drm == nil {
		return ""
	}
	return drm.ID
}

// GetLink returns the value of the field Link.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (drm *DiagnosticReportMedia) GetLink() *fhir.Reference {
	if drm == nil {
		return nil
	}
	return drm.Link
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (drm *DiagnosticReportMedia) GetModifierExtension() []*fhir.Extension {
	if drm == nil {
		return nil
	}
	return drm.ModifierExtension
}
