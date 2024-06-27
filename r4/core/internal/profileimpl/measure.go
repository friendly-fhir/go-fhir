// GENERATED CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package profileimpl

// MeasureProfile is a base interface for the Measure
// profile resource, which is defined in the exported `profile` package.
//
// This interface is used so that it can be embedded in the exported profile
// package, and keep the interface a closed-type.
//
// This uses an unexported/empty function so that it may only be implemented
// by types in a sibling package by embedding [BaseMeasure].
type MeasureProfile interface {
	isMeasure()
}

// BaseMeasure is a base-embeddable struct that implements the
// unexported isMeasure() function, which enables the
// base profile interface to claim conformance.
type BaseMeasure struct{}

// isMeasure is an unexported function that is used to claim
// conformance to the MeasureProfile interface.
func (BaseMeasure) isMeasure() {}

var _ MeasureProfile = (*BaseMeasure)(nil)
