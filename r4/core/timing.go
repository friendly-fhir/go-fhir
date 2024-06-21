// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.
package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Base StructureDefinition for Timing Type: Specifies an event that may occur
// multiple times. Timing schedules are used to record when things are planned,
// expected or requested to occur. The most common usage is in dosage
// instructions for medications. They are also used when planning care of
// various kinds, and may be used for reporting the schedule to which past
// regular activities were carried out.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Timing
//   - Source File: StructureDefinition-Timing.json
type Timing struct {

	// A code for the timing schedule (or just text in code.text). Some codes such
	// as BID are ubiquitous, but many institutions define their own additional
	// codes. If a code is provided, the code is understood to be a complete
	// statement of whatever is specified in the structured timing data, and either
	// the code or the data may be used to interpret the Timing, with the exception
	// that .repeat.bounds still applies over the code (and is not contained in the
	// code).
	Code *CodeableConcept `json:"code"`

	// Identifies specific times when the event occurs.
	Event []*DateTime `json:"event"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*Extension `json:"extension"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	Id string `json:"id"`

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
	ModifierExtension []*Extension `json:"modifierExtension"`

	// A set of rules that describe when the event is scheduled.
	Repeat Element `json:"repeat"`

	profileimpl.BaseBackboneElement
	profileimpl.BaseElement
}

func (t *Timing) GetCode() *CodeableConcept {
	if t == nil {
		return nil
	}
	return t.Code
}

func (t *Timing) GetEvent() []*DateTime {
	if t == nil {
		return nil
	}
	return t.Event
}

func (t *Timing) GetExtension() []*Extension {
	if t == nil {
		return nil
	}
	return t.Extension
}

func (t *Timing) GetId() string {
	if t == nil {
		return ""
	}
	return t.Id
}

func (t *Timing) GetModifierExtension() []*Extension {
	if t == nil {
		return nil
	}
	return t.ModifierExtension
}

func (t *Timing) GetRepeat() Element {
	if t == nil {
		return nil
	}
	return t.Repeat
}

// When the event is to occur// A set of rules that describe when the event is scheduled.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition:
//   - Source File: StructureDefinition-Timing.json
type TimingRepeat struct {

	// Either a duration for the length of the timing schedule, a range of possible
	// length, or outer bounds for start and/or end limits of the timing schedule.
	Bounds Element `json:"bounds"`

	// A total count of the desired number of repetitions across the duration of
	// the entire timing specification. If countMax is present, this element
	// indicates the lower bound of the allowed range of count values.
	Count *PositiveInt `json:"count"`

	// If present, indicates that the count is a range - so to perform the action
	// between [count] and [countMax] times.
	CountMax *PositiveInt `json:"countMax"`

	// If one or more days of week is provided, then the action happens only on the
	// specified day(s).
	DayOfWeek []*Code `json:"dayOfWeek"`

	// How long this thing happens for when it happens. If durationMax is present,
	// this element indicates the lower bound of the allowed range of the duration.
	Duration *Decimal `json:"duration"`

	// If present, indicates that the duration is a range - so to perform the
	// action between [duration] and [durationMax] time length.
	DurationMax *Decimal `json:"durationMax"`

	// The units of time for the duration, in UCUM units.
	DurationUnit *Code `json:"durationUnit"`

	// May be used to represent additional information that is not part of the
	// basic definition of the element. To make the use of extensions safe and
	// manageable, there is a strict set of governance applied to the definition
	// and use of extensions. Though any implementer can define an extension, there
	// is a set of requirements that SHALL be met as part of the definition of the
	// extension.
	Extension []*Extension `json:"extension"`

	// The number of times to repeat the action within the specified period. If
	// frequencyMax is present, this element indicates the lower bound of the
	// allowed range of the frequency.
	Frequency *PositiveInt `json:"frequency"`

	// If present, indicates that the frequency is a range - so to repeat between
	// [frequency] and [frequencyMax] times within the period or period range.
	FrequencyMax *PositiveInt `json:"frequencyMax"`

	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	Id string `json:"id"`

	// The number of minutes from the event. If the event code does not indicate
	// whether the minutes is before or after the event, then the offset is assumed
	// to be after the event.
	Offset *UnsignedInt `json:"offset"`

	// Indicates the duration of time over which repetitions are to occur; e.g. to
	// express "3 times per day", 3 would be the frequency and "1 day" would be the
	// period. If periodMax is present, this element indicates the lower bound of
	// the allowed range of the period length.
	Period *Decimal `json:"period"`

	// If present, indicates that the period is a range from [period] to
	// [periodMax], allowing expressing concepts such as "do this once every 3-5
	// days.
	PeriodMax *Decimal `json:"periodMax"`

	// The units of time for the period in UCUM units.
	PeriodUnit *Code `json:"periodUnit"`

	// Specified time of day for action to take place.
	TimeOfDay []*Time `json:"timeOfDay"`

	// An approximate time period during the day, potentially linked to an event of
	// daily living that indicates when the action should occur.
	When []*Code `json:"when"`

	profileimpl.BaseElement
}

func (v *TimingRepeat) GetBoundsDuration() *Duration {
	if v == nil {
		return nil
	}
	result, ok := v.Bounds.(*Duration)
	if ok {
		return result
	}
	return nil
}
func (v *TimingRepeat) GetBoundsRange() *Range {
	if v == nil {
		return nil
	}
	result, ok := v.Bounds.(*Range)
	if ok {
		return result
	}
	return nil
}
func (v *TimingRepeat) GetBoundsPeriod() *Period {
	if v == nil {
		return nil
	}
	result, ok := v.Bounds.(*Period)
	if ok {
		return result
	}
	return nil
}
func (t *TimingRepeat) GetBounds() Element {
	if t == nil {
		return nil
	}
	return t.Bounds
}

func (t *TimingRepeat) GetCount() *PositiveInt {
	if t == nil {
		return nil
	}
	return t.Count
}

func (t *TimingRepeat) GetCountMax() *PositiveInt {
	if t == nil {
		return nil
	}
	return t.CountMax
}

func (t *TimingRepeat) GetDayOfWeek() []*Code {
	if t == nil {
		return nil
	}
	return t.DayOfWeek
}

func (t *TimingRepeat) GetDuration() *Decimal {
	if t == nil {
		return nil
	}
	return t.Duration
}

func (t *TimingRepeat) GetDurationMax() *Decimal {
	if t == nil {
		return nil
	}
	return t.DurationMax
}

func (t *TimingRepeat) GetDurationUnit() *Code {
	if t == nil {
		return nil
	}
	return t.DurationUnit
}

func (t *TimingRepeat) GetExtension() []*Extension {
	if t == nil {
		return nil
	}
	return t.Extension
}

func (t *TimingRepeat) GetFrequency() *PositiveInt {
	if t == nil {
		return nil
	}
	return t.Frequency
}

func (t *TimingRepeat) GetFrequencyMax() *PositiveInt {
	if t == nil {
		return nil
	}
	return t.FrequencyMax
}

func (t *TimingRepeat) GetId() string {
	if t == nil {
		return ""
	}
	return t.Id
}

func (t *TimingRepeat) GetOffset() *UnsignedInt {
	if t == nil {
		return nil
	}
	return t.Offset
}

func (t *TimingRepeat) GetPeriod() *Decimal {
	if t == nil {
		return nil
	}
	return t.Period
}

func (t *TimingRepeat) GetPeriodMax() *Decimal {
	if t == nil {
		return nil
	}
	return t.PeriodMax
}

func (t *TimingRepeat) GetPeriodUnit() *Code {
	if t == nil {
		return nil
	}
	return t.PeriodUnit
}

func (t *TimingRepeat) GetTimeOfDay() []*Time {
	if t == nil {
		return nil
	}
	return t.TimeOfDay
}

func (t *TimingRepeat) GetWhen() []*Code {
	if t == nil {
		return nil
	}
	return t.When
}
