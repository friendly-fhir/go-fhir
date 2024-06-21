// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.
package profileimpl

// QuestionnaireProfile is a base interface for the Questionnaire
// profile resource, which is defined in the exported `profile` package.
//
// This interface is used so that it can be embedded in the exported profile
// package, and keep the interface a closed-type.
//
// This uses an unexported/empty function so that it may only be implemented
// by types in a sibling package by embedding [BaseQuestionnaire].
type QuestionnaireProfile interface {
	isQuestionnaire()
}

// BaseQuestionnaire is a base-embeddable struct that implements the
// unexported isQuestionnaire() function, which enables the
// base profile interface to claim conformance.
type BaseQuestionnaire struct{}

// isQuestionnaire is an unexported function that is used to claim
// conformance to the QuestionnaireProfile interface.
func (BaseQuestionnaire) isQuestionnaire() {}

var _ QuestionnaireProfile = (*BaseQuestionnaire)(nil)
