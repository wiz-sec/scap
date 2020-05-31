// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#freebsd
package freebsd_def

import (
	"encoding/xml"

	"github.com/gocomply/scap/pkg/scap/models/oval_def"
)

// Element
type PortinfoTest struct {
	XMLName xml.Name `xml:portinfo_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type PortinfoObject struct {
	XMLName xml.Name `xml:portinfo_object`

	Set oval_def.Set `xml:"set"`
}

// Element
type PortinfoState struct {
	XMLName xml.Name `xml:portinfo_state`

	Pkginst *oval_def.EntityStateStringType `xml:"pkginst"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Category *oval_def.EntityStateStringType `xml:"category"`

	Version *PortinfoStateVersion `xml:"version"`

	Vendor *oval_def.EntityStateStringType `xml:"vendor"`

	Description *oval_def.EntityStateStringType `xml:"description"`
}

// Element
type PortinfoStateVersion struct {
	XMLName xml.Name `xml:version`
}

// XSD ComplexType declarations
