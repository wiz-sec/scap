// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
package aix_def

import (
	"encoding/xml"
)

// Element
type InterimFixTest struct {
	XMLName xml.Name `xml:interim_fix_test`
}

// Element
type InterimFixObject struct {
	XMLName xml.Name `xml:interim_fix_object`
}

// Element
type InterimFixState struct {
	XMLName xml.Name `xml:interim_fix_state`
}

// Element
type FilesetTest struct {
	XMLName xml.Name `xml:fileset_test`
}

// Element
type FilesetObject struct {
	XMLName xml.Name `xml:fileset_object`
}

// Element
type FilesetState struct {
	XMLName xml.Name `xml:fileset_state`
}

// Element
type FixTest struct {
	XMLName xml.Name `xml:fix_test`
}

// Element
type FixObject struct {
	XMLName xml.Name `xml:fix_object`
}

// Element
type FixState struct {
	XMLName xml.Name `xml:fix_state`
}

// Element
type NoTest struct {
	XMLName xml.Name `xml:no_test`
}

// Element
type NoObject struct {
	XMLName xml.Name `xml:no_object`
}

// Element
type NoState struct {
	XMLName xml.Name `xml:no_state`
}

// Element
type OslevelTest struct {
	XMLName xml.Name `xml:oslevel_test`
}

// Element
type OslevelObject struct {
	XMLName xml.Name `xml:oslevel_object`
}

// Element
type OslevelState struct {
	XMLName xml.Name `xml:oslevel_state`
}

// XSD ComplexType declarations

type EntityStateFilesetStateType struct {
}

type EntityStateFixInstallationStatusType struct {
}

type EntityStateInterimFixStateType struct {
}