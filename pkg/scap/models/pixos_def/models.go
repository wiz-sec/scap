// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos
package pixos_def

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_def"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type LineTest struct {
	XMLName xml.Name `xml:line_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type LineObject struct {
	XMLName xml.Name `xml:line_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type LineState struct {
	XMLName xml.Name `xml:line_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	ShowSubcommand *oval_def.EntityStateStringType `xml:"show_subcommand"`

	ConfigLine *oval_def.EntityStateStringType `xml:"config_line"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VersionTest struct {
	XMLName xml.Name `xml:version_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VersionObject struct {
	XMLName xml.Name `xml:version_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VersionState struct {
	XMLName xml.Name `xml:version_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	PixRelease *oval_def.EntityStateStringType `xml:"pix_release"`

	PixMajorRelease *oval_def.EntityStateVersionType `xml:"pix_major_release"`

	PixMinorRelease *oval_def.EntityStateVersionType `xml:"pix_minor_release"`

	PixBuild *oval_def.EntityStateIntType `xml:"pix_build"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// XSD ComplexType declarations
