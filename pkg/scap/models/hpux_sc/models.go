// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#hpux
package hpux_sc

import (
	"encoding/xml"
)

// Element
type GetconfItem struct {
	XMLName xml.Name `xml:getconf_item`
}

// Element
type NddItem struct {
	XMLName xml.Name `xml:ndd_item`
}

// Element
type PatchItem struct {
	XMLName xml.Name `xml:patch_item`
}

// Element
type SwlistItem struct {
	XMLName xml.Name `xml:swlist_item`
}

// Element
type TrustedItem struct {
	XMLName xml.Name `xml:trusted_item`
}

// XSD ComplexType declarations