// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-common-5
package oval

import (
	"encoding/xml"
)

// Element
type DeprecatedInfo struct {
	XMLName xml.Name `xml:deprecated_info`

	Version string `xml:"version"`

	Reason string `xml:"reason"`

	Comment string `xml:"comment"`
}

// Element
type ElementMapping struct {
	XMLName xml.Name `xml:element_mapping`

	Test ElementMapItemType `xml:"test"`

	Object *ElementMapItemType `xml:"object"`

	State *ElementMapItemType `xml:"state"`

	Item *ElementMapItemType `xml:"item"`
}

// Element
type Notes struct {
	XMLName xml.Name `xml:notes`

	Note []string `xml:"note"`
}

// XSD ComplexType declarations

type ElementMapType struct {
	Test ElementMapItemType `xml:"test"`

	Object *ElementMapItemType `xml:"object"`

	State *ElementMapItemType `xml:"state"`

	Item *ElementMapItemType `xml:"item"`

	InnerXml string `xml:",innerxml"`
}

type ElementMapItemType struct {
	TargetNamespace *string `xml:"target_namespace,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type DeprecatedInfoType struct {
	Version string `xml:"version"`

	Reason string `xml:"reason"`

	Comment string `xml:"comment"`

	InnerXml string `xml:",innerxml"`
}

type GeneratorType struct {
	ProductName string `xml:"product_name"`

	ProductVersion string `xml:"product_version"`

	SchemaVersion []SchemaVersionType `xml:"schema_version"`

	Timestamp string `xml:"timestamp"`

	InnerXml string `xml:",innerxml"`
}

type SchemaVersionType struct {
	Platform *string `xml:"platform,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type MessageType struct {
	Level *string `xml:"level,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type NotesType struct {
	Note []string `xml:"note"`

	InnerXml string `xml:",innerxml"`
}
