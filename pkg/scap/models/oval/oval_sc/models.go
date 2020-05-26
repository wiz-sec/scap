// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5
package oval_sc

import (
	"encoding/xml"

	"github.com/gocomply/scap/pkg/scap/models/oval/ds"

	"github.com/gocomply/scap/pkg/scap/models/oval/oval"
)

// Element
type OvalSystemCharacteristics struct {
	XMLName xml.Name `xml:oval_system_characteristics`

	Generator oval.GeneratorType `xml:"generator"`

	SystemInfo SystemInfoType `xml:"system_info"`

	CollectedObjects CollectedObjectsType `xml:"collected_objects"`

	SystemData SystemDataType `xml:"system_data"`

	Signature ds.Signature `xml:"Signature"`
}

// Element
type Item struct {
	XMLName xml.Name `xml:item`
}

// XSD ComplexType declarations

type SystemInfoType string

type InterfacesType string

type InterfaceType string

type CollectedObjectsType string

type ObjectType string

type VariableValueType string

type ReferenceType string

type SystemDataType string

type ItemType string

type EntityItemSimpleBaseType string

type EntityItemComplexBaseType string

type EntityItemIPAddressType string

type EntityItemIPAddressStringType string

type EntityItemAnySimpleType string

type EntityItemBinaryType string

type EntityItemBoolType string

type EntityItemFloatType string

type EntityItemIntType string

type EntityItemStringType string

type EntityItemRecordType string

type EntityItemFieldType string

type EntityItemVersionType string

type EntityItemFilesetRevisionType string

type EntityItemIOSVersionType string

type EntityItemEVRStringType string

type EntityItemDebianEVRStringType string
