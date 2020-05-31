// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://cpe.mitre.org/dictionary/2.0
package cpe_dict

import (
	"encoding/xml"
)

// Element
type CpeList struct {
	XMLName xml.Name `xml:cpe-list`

	Generator *GeneratorType `xml:"generator"`

	CpeItem []CpeItem `xml:"cpe-item"`
}

// Element
type CpeItem struct {
	XMLName xml.Name `xml:cpe-item`

	Title []TextType `xml:"title"`

	Notes []NotesType `xml:"notes"`

	References *ReferencesType `xml:"references"`

	Check []CheckType `xml:"check"`
}

// Element
type Reference struct {
	XMLName xml.Name `xml:reference`

	Href string `xml:"href,attr"`
}

// XSD ComplexType declarations

type GeneratorType struct {
	ProductName string `xml:"product_name"`

	ProductVersion string `xml:"product_version"`

	SchemaVersion float64 `xml:"schema_version"`

	Timestamp string `xml:"timestamp"`
}

type ItemType struct {
	Name string `xml:"name,attr"`

	Deprecated string `xml:"deprecated,attr"`

	DeprecatedBy string `xml:"deprecated_by,attr"`

	DeprecationDate string `xml:"deprecation_date,attr"`

	Title []TextType `xml:"title"`

	Notes []NotesType `xml:"notes"`

	References *ReferencesType `xml:"references"`

	Check []CheckType `xml:"check"`
}

type ListType struct {
	Generator *GeneratorType `xml:"generator"`

	CpeItem []CpeItem `xml:"cpe-item"`
}

type TextType struct {
	XmlLang string `xml:",attr"`
}

type NotesType struct {
	XmlLang string `xml:",attr"`

	Note []string `xml:"note"`
}

type ReferencesType struct {
	Reference []Reference `xml:"reference"`
}

type CheckType struct {
	System string `xml:"system,attr"`

	Href string `xml:"href,attr"`
}
