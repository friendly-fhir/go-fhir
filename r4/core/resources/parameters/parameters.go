// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package parameters

import (
	"github.com/friendly-fhir/go-fhir/internal/validate"
	"github.com/friendly-fhir/go-fhir/r4/core"
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"

	"encoding/json"
)

// This resource is a non-persisted resource used to pass information into and
// back from an [operation](operations.html). It has no other use, and there is
// no RESTful endpoint associated with it.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Parameters
//   - Source File: StructureDefinition-Parameters.json
type Parameters struct {

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

	// A parameter passed to or received from the operation.
	Parameter []*ParametersParameter `fhirpath:"parameter"`

	profileimpl.BaseResource
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Parameters) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}

// GetImplicitRules returns the value of the field ImplicitRules.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Parameters) GetImplicitRules() *fhir.URI {
	if p == nil {
		return nil
	}
	return p.ImplicitRules
}

// GetLanguage returns the value of the field Language.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Parameters) GetLanguage() *fhir.Code {
	if p == nil {
		return nil
	}
	return p.Language
}

// GetMeta returns the value of the field Meta.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Parameters) GetMeta() *fhir.Meta {
	if p == nil {
		return nil
	}
	return p.Meta
}

// GetParameter returns the value of the field Parameter.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (p *Parameters) GetParameter() []*ParametersParameter {
	if p == nil {
		return nil
	}
	return p.Parameter
}

// Operation Parameter// A parameter passed to or received from the operation.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Parameters.json
type ParametersParameter struct {

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

	// The name of the parameter (reference to the operation definition).
	Name *fhir.String `fhirpath:"name"`

	// If the parameter is a whole resource.
	Resource fhir.Resource `fhirpath:"resource"`

	// If the parameter is a data type.
	Value fhir.Element `fhirpath:"value"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

// GetExtension returns the value of the field Extension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pp *ParametersParameter) GetExtension() []*fhir.Extension {
	if pp == nil {
		return nil
	}
	return pp.Extension
}

// GetID returns the value of the field ID.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pp *ParametersParameter) GetID() string {
	if pp == nil {
		return ""
	}
	return pp.ID
}

// GetModifierExtension returns the value of the field ModifierExtension.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pp *ParametersParameter) GetModifierExtension() []*fhir.Extension {
	if pp == nil {
		return nil
	}
	return pp.ModifierExtension
}

// GetName returns the value of the field Name.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pp *ParametersParameter) GetName() *fhir.String {
	if pp == nil {
		return nil
	}
	return pp.Name
}

// GetResource returns the value of the field Resource.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pp *ParametersParameter) GetResource() fhir.Resource {
	if pp == nil {
		return nil
	}
	return pp.Resource
}

// GetValue returns the value of the field Value.
// This function is safe to call on nil pointers, and will return the zero value
// instead.
func (pp *ParametersParameter) GetValue() fhir.Element {
	if pp == nil {
		return nil
	}
	return pp.Value
}

// GetValueBase64Binary returns the value of the field Value.
func (pp *ParametersParameter) GetValueBase64Binary() *fhir.Base64Binary {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Base64Binary)
	if !ok {
		return nil
	}
	return val
}

// GetValueBoolean returns the value of the field Value.
func (pp *ParametersParameter) GetValueBoolean() *fhir.Boolean {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Boolean)
	if !ok {
		return nil
	}
	return val
}

// GetValueCanonical returns the value of the field Value.
func (pp *ParametersParameter) GetValueCanonical() *fhir.Canonical {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Canonical)
	if !ok {
		return nil
	}
	return val
}

// GetValueCode returns the value of the field Value.
func (pp *ParametersParameter) GetValueCode() *fhir.Code {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Code)
	if !ok {
		return nil
	}
	return val
}

// GetValueDate returns the value of the field Value.
func (pp *ParametersParameter) GetValueDate() *fhir.Date {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Date)
	if !ok {
		return nil
	}
	return val
}

// GetValueDateTime returns the value of the field Value.
func (pp *ParametersParameter) GetValueDateTime() *fhir.DateTime {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.DateTime)
	if !ok {
		return nil
	}
	return val
}

// GetValueDecimal returns the value of the field Value.
func (pp *ParametersParameter) GetValueDecimal() *fhir.Decimal {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Decimal)
	if !ok {
		return nil
	}
	return val
}

// GetValueID returns the value of the field Value.
func (pp *ParametersParameter) GetValueID() *fhir.ID {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.ID)
	if !ok {
		return nil
	}
	return val
}

// GetValueInstant returns the value of the field Value.
func (pp *ParametersParameter) GetValueInstant() *fhir.Instant {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Instant)
	if !ok {
		return nil
	}
	return val
}

// GetValueInteger returns the value of the field Value.
func (pp *ParametersParameter) GetValueInteger() *fhir.Integer {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Integer)
	if !ok {
		return nil
	}
	return val
}

// GetValueMarkdown returns the value of the field Value.
func (pp *ParametersParameter) GetValueMarkdown() *fhir.Markdown {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Markdown)
	if !ok {
		return nil
	}
	return val
}

// GetValueOID returns the value of the field Value.
func (pp *ParametersParameter) GetValueOID() *fhir.OID {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.OID)
	if !ok {
		return nil
	}
	return val
}

// GetValuePositiveInt returns the value of the field Value.
func (pp *ParametersParameter) GetValuePositiveInt() *fhir.PositiveInt {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.PositiveInt)
	if !ok {
		return nil
	}
	return val
}

// GetValueString returns the value of the field Value.
func (pp *ParametersParameter) GetValueString() *fhir.String {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.String)
	if !ok {
		return nil
	}
	return val
}

// GetValueTime returns the value of the field Value.
func (pp *ParametersParameter) GetValueTime() *fhir.Time {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Time)
	if !ok {
		return nil
	}
	return val
}

// GetValueUnsignedInt returns the value of the field Value.
func (pp *ParametersParameter) GetValueUnsignedInt() *fhir.UnsignedInt {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.UnsignedInt)
	if !ok {
		return nil
	}
	return val
}

// GetValueURI returns the value of the field Value.
func (pp *ParametersParameter) GetValueURI() *fhir.URI {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.URI)
	if !ok {
		return nil
	}
	return val
}

// GetValueURL returns the value of the field Value.
func (pp *ParametersParameter) GetValueURL() *fhir.URL {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.URL)
	if !ok {
		return nil
	}
	return val
}

// GetValueUUID returns the value of the field Value.
func (pp *ParametersParameter) GetValueUUID() *fhir.UUID {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.UUID)
	if !ok {
		return nil
	}
	return val
}

// GetValueAddress returns the value of the field Value.
func (pp *ParametersParameter) GetValueAddress() *fhir.Address {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Address)
	if !ok {
		return nil
	}
	return val
}

// GetValueAge returns the value of the field Value.
func (pp *ParametersParameter) GetValueAge() *fhir.Age {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Age)
	if !ok {
		return nil
	}
	return val
}

// GetValueAnnotation returns the value of the field Value.
func (pp *ParametersParameter) GetValueAnnotation() *fhir.Annotation {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Annotation)
	if !ok {
		return nil
	}
	return val
}

// GetValueAttachment returns the value of the field Value.
func (pp *ParametersParameter) GetValueAttachment() *fhir.Attachment {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Attachment)
	if !ok {
		return nil
	}
	return val
}

// GetValueCodeableConcept returns the value of the field Value.
func (pp *ParametersParameter) GetValueCodeableConcept() *fhir.CodeableConcept {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.CodeableConcept)
	if !ok {
		return nil
	}
	return val
}

// GetValueCoding returns the value of the field Value.
func (pp *ParametersParameter) GetValueCoding() *fhir.Coding {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Coding)
	if !ok {
		return nil
	}
	return val
}

// GetValueContactPoint returns the value of the field Value.
func (pp *ParametersParameter) GetValueContactPoint() *fhir.ContactPoint {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.ContactPoint)
	if !ok {
		return nil
	}
	return val
}

// GetValueCount returns the value of the field Value.
func (pp *ParametersParameter) GetValueCount() *fhir.Count {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Count)
	if !ok {
		return nil
	}
	return val
}

// GetValueDistance returns the value of the field Value.
func (pp *ParametersParameter) GetValueDistance() *fhir.Distance {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Distance)
	if !ok {
		return nil
	}
	return val
}

// GetValueDuration returns the value of the field Value.
func (pp *ParametersParameter) GetValueDuration() *fhir.Duration {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Duration)
	if !ok {
		return nil
	}
	return val
}

// GetValueHumanName returns the value of the field Value.
func (pp *ParametersParameter) GetValueHumanName() *fhir.HumanName {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.HumanName)
	if !ok {
		return nil
	}
	return val
}

// GetValueIdentifier returns the value of the field Value.
func (pp *ParametersParameter) GetValueIdentifier() *fhir.Identifier {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Identifier)
	if !ok {
		return nil
	}
	return val
}

// GetValueMoney returns the value of the field Value.
func (pp *ParametersParameter) GetValueMoney() *fhir.Money {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Money)
	if !ok {
		return nil
	}
	return val
}

// GetValuePeriod returns the value of the field Value.
func (pp *ParametersParameter) GetValuePeriod() *fhir.Period {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Period)
	if !ok {
		return nil
	}
	return val
}

// GetValueQuantity returns the value of the field Value.
func (pp *ParametersParameter) GetValueQuantity() *fhir.Quantity {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Quantity)
	if !ok {
		return nil
	}
	return val
}

// GetValueRange returns the value of the field Value.
func (pp *ParametersParameter) GetValueRange() *fhir.Range {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Range)
	if !ok {
		return nil
	}
	return val
}

// GetValueRatio returns the value of the field Value.
func (pp *ParametersParameter) GetValueRatio() *fhir.Ratio {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Ratio)
	if !ok {
		return nil
	}
	return val
}

// GetValueReference returns the value of the field Value.
func (pp *ParametersParameter) GetValueReference() *fhir.Reference {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Reference)
	if !ok {
		return nil
	}
	return val
}

// GetValueSampledData returns the value of the field Value.
func (pp *ParametersParameter) GetValueSampledData() *fhir.SampledData {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.SampledData)
	if !ok {
		return nil
	}
	return val
}

// GetValueSignature returns the value of the field Value.
func (pp *ParametersParameter) GetValueSignature() *fhir.Signature {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Signature)
	if !ok {
		return nil
	}
	return val
}

// GetValueTiming returns the value of the field Value.
func (pp *ParametersParameter) GetValueTiming() *fhir.Timing {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Timing)
	if !ok {
		return nil
	}
	return val
}

// GetValueContactDetail returns the value of the field Value.
func (pp *ParametersParameter) GetValueContactDetail() *fhir.ContactDetail {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.ContactDetail)
	if !ok {
		return nil
	}
	return val
}

// GetValueContributor returns the value of the field Value.
func (pp *ParametersParameter) GetValueContributor() *fhir.Contributor {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Contributor)
	if !ok {
		return nil
	}
	return val
}

// GetValueDataRequirement returns the value of the field Value.
func (pp *ParametersParameter) GetValueDataRequirement() *fhir.DataRequirement {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.DataRequirement)
	if !ok {
		return nil
	}
	return val
}

// GetValueExpression returns the value of the field Value.
func (pp *ParametersParameter) GetValueExpression() *fhir.Expression {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Expression)
	if !ok {
		return nil
	}
	return val
}

// GetValueParameterDefinition returns the value of the field Value.
func (pp *ParametersParameter) GetValueParameterDefinition() *fhir.ParameterDefinition {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.ParameterDefinition)
	if !ok {
		return nil
	}
	return val
}

// GetValueRelatedArtifact returns the value of the field Value.
func (pp *ParametersParameter) GetValueRelatedArtifact() *fhir.RelatedArtifact {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.RelatedArtifact)
	if !ok {
		return nil
	}
	return val
}

// GetValueTriggerDefinition returns the value of the field Value.
func (pp *ParametersParameter) GetValueTriggerDefinition() *fhir.TriggerDefinition {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.TriggerDefinition)
	if !ok {
		return nil
	}
	return val
}

// GetValueUsageContext returns the value of the field Value.
func (pp *ParametersParameter) GetValueUsageContext() *fhir.UsageContext {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.UsageContext)
	if !ok {
		return nil
	}
	return val
}

// GetValueDosage returns the value of the field Value.
func (pp *ParametersParameter) GetValueDosage() *fhir.Dosage {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Dosage)
	if !ok {
		return nil
	}
	return val
}

// GetValueMeta returns the value of the field Value.
func (pp *ParametersParameter) GetValueMeta() *fhir.Meta {
	if pp == nil {
		return nil
	}
	val, ok := pp.Value.(*fhir.Meta)
	if !ok {
		return nil
	}
	return val
}

func (p *Parameters) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (p *Parameters) UnmarshalJSON(data []byte) error {
	var raw struct {
		ID            string                 `json:"id"`
		ImplicitRules *fhir.URI              `json:"implicitRules"`
		Language      *fhir.Code             `json:"language"`
		Meta          *fhir.Meta             `json:"meta"`
		Parameter     []*ParametersParameter `json:"parameter"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	p.ID = raw.ID
	p.ImplicitRules = raw.ImplicitRules
	p.Language = raw.Language
	p.Meta = raw.Meta
	p.Parameter = raw.Parameter
	return nil
}

var _ json.Marshaler = (*Parameters)(nil)
var _ json.Unmarshaler = (*Parameters)(nil)

func (pp *ParametersParameter) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (pp *ParametersParameter) UnmarshalJSON(data []byte) error {
	var raw struct {
		Extension []*fhir.Extension `json:"extension"`

		ID                       string                    `json:"id"`
		ModifierExtension        []*fhir.Extension         `json:"modifierExtension"`
		Name                     *fhir.String              `json:"name"`
		Resource                 fhir.Resource             `json:"resource"`
		ValueBase64Binary        *fhir.Base64Binary        `json:"valueBase64Binary"`
		ValueBoolean             *fhir.Boolean             `json:"valueBoolean"`
		ValueCanonical           *fhir.Canonical           `json:"valueCanonical"`
		ValueCode                *fhir.Code                `json:"valueCode"`
		ValueDate                *fhir.Date                `json:"valueDate"`
		ValueDateTime            *fhir.DateTime            `json:"valueDateTime"`
		ValueDecimal             *fhir.Decimal             `json:"valueDecimal"`
		ValueID                  *fhir.ID                  `json:"valueID"`
		ValueInstant             *fhir.Instant             `json:"valueInstant"`
		ValueInteger             *fhir.Integer             `json:"valueInteger"`
		ValueMarkdown            *fhir.Markdown            `json:"valueMarkdown"`
		ValueOID                 *fhir.OID                 `json:"valueOID"`
		ValuePositiveInt         *fhir.PositiveInt         `json:"valuePositiveInt"`
		ValueString              *fhir.String              `json:"valueString"`
		ValueTime                *fhir.Time                `json:"valueTime"`
		ValueUnsignedInt         *fhir.UnsignedInt         `json:"valueUnsignedInt"`
		ValueURI                 *fhir.URI                 `json:"valueURI"`
		ValueURL                 *fhir.URL                 `json:"valueURL"`
		ValueUUID                *fhir.UUID                `json:"valueUUID"`
		ValueAddress             *fhir.Address             `json:"valueAddress"`
		ValueAge                 *fhir.Age                 `json:"valueAge"`
		ValueAnnotation          *fhir.Annotation          `json:"valueAnnotation"`
		ValueAttachment          *fhir.Attachment          `json:"valueAttachment"`
		ValueCodeableConcept     *fhir.CodeableConcept     `json:"valueCodeableConcept"`
		ValueCoding              *fhir.Coding              `json:"valueCoding"`
		ValueContactPoint        *fhir.ContactPoint        `json:"valueContactPoint"`
		ValueCount               *fhir.Count               `json:"valueCount"`
		ValueDistance            *fhir.Distance            `json:"valueDistance"`
		ValueDuration            *fhir.Duration            `json:"valueDuration"`
		ValueHumanName           *fhir.HumanName           `json:"valueHumanName"`
		ValueIdentifier          *fhir.Identifier          `json:"valueIdentifier"`
		ValueMoney               *fhir.Money               `json:"valueMoney"`
		ValuePeriod              *fhir.Period              `json:"valuePeriod"`
		ValueQuantity            *fhir.Quantity            `json:"valueQuantity"`
		ValueRange               *fhir.Range               `json:"valueRange"`
		ValueRatio               *fhir.Ratio               `json:"valueRatio"`
		ValueReference           *fhir.Reference           `json:"valueReference"`
		ValueSampledData         *fhir.SampledData         `json:"valueSampledData"`
		ValueSignature           *fhir.Signature           `json:"valueSignature"`
		ValueTiming              *fhir.Timing              `json:"valueTiming"`
		ValueContactDetail       *fhir.ContactDetail       `json:"valueContactDetail"`
		ValueContributor         *fhir.Contributor         `json:"valueContributor"`
		ValueDataRequirement     *fhir.DataRequirement     `json:"valueDataRequirement"`
		ValueExpression          *fhir.Expression          `json:"valueExpression"`
		ValueParameterDefinition *fhir.ParameterDefinition `json:"valueParameterDefinition"`
		ValueRelatedArtifact     *fhir.RelatedArtifact     `json:"valueRelatedArtifact"`
		ValueTriggerDefinition   *fhir.TriggerDefinition   `json:"valueTriggerDefinition"`
		ValueUsageContext        *fhir.UsageContext        `json:"valueUsageContext"`
		ValueDosage              *fhir.Dosage              `json:"valueDosage"`
		ValueMeta                *fhir.Meta                `json:"valueMeta"`
	}

	var err error
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	pp.Extension = raw.Extension
	pp.ID = raw.ID
	pp.ModifierExtension = raw.ModifierExtension
	pp.Name = raw.Name
	pp.Resource = raw.Resource
	pp.Value, err = validate.SelectOneOf[fhir.Element]("Parameters.parameter.value",
		raw.ValueBase64Binary,
		raw.ValueBoolean,
		raw.ValueCanonical,
		raw.ValueCode,
		raw.ValueDate,
		raw.ValueDateTime,
		raw.ValueDecimal,
		raw.ValueID,
		raw.ValueInstant,
		raw.ValueInteger,
		raw.ValueMarkdown,
		raw.ValueOID,
		raw.ValuePositiveInt,
		raw.ValueString,
		raw.ValueTime,
		raw.ValueUnsignedInt,
		raw.ValueURI,
		raw.ValueURL,
		raw.ValueUUID,
		raw.ValueAddress,
		raw.ValueAge,
		raw.ValueAnnotation,
		raw.ValueAttachment,
		raw.ValueCodeableConcept,
		raw.ValueCoding,
		raw.ValueContactPoint,
		raw.ValueCount,
		raw.ValueDistance,
		raw.ValueDuration,
		raw.ValueHumanName,
		raw.ValueIdentifier,
		raw.ValueMoney,
		raw.ValuePeriod,
		raw.ValueQuantity,
		raw.ValueRange,
		raw.ValueRatio,
		raw.ValueReference,
		raw.ValueSampledData,
		raw.ValueSignature,
		raw.ValueTiming,
		raw.ValueContactDetail,
		raw.ValueContributor,
		raw.ValueDataRequirement,
		raw.ValueExpression,
		raw.ValueParameterDefinition,
		raw.ValueRelatedArtifact,
		raw.ValueTriggerDefinition,
		raw.ValueUsageContext,
		raw.ValueDosage,
		raw.ValueMeta)
	if err != nil {
		return err
	}
	return nil
}

var _ json.Marshaler = (*ParametersParameter)(nil)
var _ json.Unmarshaler = (*ParametersParameter)(nil)
