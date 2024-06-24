// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.

package profileimpl

// ValueSetProfile is a base interface for the ValueSet
// profile resource, which is defined in the exported `profile` package.
//
// This interface is used so that it can be embedded in the exported profile
// package, and keep the interface a closed-type.
//
// This uses an unexported/empty function so that it may only be implemented
// by types in a sibling package by embedding [BaseValueSet].
type ValueSetProfile interface {
	isValueSet()
}

// BaseValueSet is a base-embeddable struct that implements the
// unexported isValueSet() function, which enables the
// base profile interface to claim conformance.
type BaseValueSet struct{}

// isValueSet is an unexported function that is used to claim
// conformance to the ValueSetProfile interface.
func (BaseValueSet) isValueSet() {}

var _ ValueSetProfile = (*BaseValueSet)(nil)
