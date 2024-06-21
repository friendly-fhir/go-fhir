// GENERATE CODE - DO NOT EDIT
//
// This file is generated by the FHIR Fhenix code generator tool, which can be
// located at https://github.com/friendly-fhir/fhenix.
package fhir

import (
	"github.com/friendly-fhir/go-fhir/r4/core/internal/profileimpl"
)

// Base StructureDefinition for Address Type: An address expressed using postal
// conventions (as opposed to GPS or other location definition formats). This
// data type may be used to convey addresses for use in delivering mail as well
// as for visiting locations which might not be valid for mail delivery. There
// are a variety of postal address formats defined around the world.
//
// Meta Details:
//
//   - Package IG: hl7.fhir.r4.core@4.0.1
//   - StructureDefinition: http://hl7.org/fhir/StructureDefinition/Address
//   - Source File: StructureDefinition-Address.json
type Address struct {

	// The name of the city, town, suburb, village or other community or delivery
	// center.
	City *String `json:"city"`

	// Country - a nation as commonly understood or generally accepted.
	Country *String `json:"country"`

	// The name of the administrative area (county).
	District *String `json:"district"`

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

	// This component contains the house number, apartment number, street name,
	// street direction, P.O. Box number, delivery hints, and similar address
	// information.
	Line []*String `json:"line"`

	// Time period when address was/is in use.
	Period *Period `json:"period"`

	// A postal code designating a region defined by the postal service.
	PostalCode *String `json:"postalCode"`

	// Sub-unit of a country with limited sovereignty in a federally organized
	// country. A code may be used if codes are in common use (e.g. US 2 letter
	// state codes).
	State *String `json:"state"`

	// Specifies the entire address as it should be displayed e.g. on a postal
	// label. This may be provided instead of or as well as the specific parts.
	Text *String `json:"text"`

	// Distinguishes between physical addresses (those you can visit) and mailing
	// addresses (e.g. PO Boxes and care-of addresses). Most addresses are both.
	Type *Code `json:"type"`

	// The purpose of this address.
	Use *Code `json:"use"`

	profileimpl.BaseElement
}

func (a *Address) GetCity() *String {
	if a == nil {
		return nil
	}
	return a.City
}

func (a *Address) GetCountry() *String {
	if a == nil {
		return nil
	}
	return a.Country
}

func (a *Address) GetDistrict() *String {
	if a == nil {
		return nil
	}
	return a.District
}

func (a *Address) GetExtension() []*Extension {
	if a == nil {
		return nil
	}
	return a.Extension
}

func (a *Address) GetId() string {
	if a == nil {
		return ""
	}
	return a.Id
}

func (a *Address) GetLine() []*String {
	if a == nil {
		return nil
	}
	return a.Line
}

func (a *Address) GetPeriod() *Period {
	if a == nil {
		return nil
	}
	return a.Period
}

func (a *Address) GetPostalCode() *String {
	if a == nil {
		return nil
	}
	return a.PostalCode
}

func (a *Address) GetState() *String {
	if a == nil {
		return nil
	}
	return a.State
}

func (a *Address) GetText() *String {
	if a == nil {
		return nil
	}
	return a.Text
}

func (a *Address) GetType() *Code {
	if a == nil {
		return nil
	}
	return a.Type
}

func (a *Address) GetUse() *Code {
	if a == nil {
		return nil
	}
	return a.Use
}
