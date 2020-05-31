// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#catos
package catos_sc

import (
	"encoding/xml"

	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

// Element
type LineItem struct {
	XMLName xml.Name `xml:line_item`

	ShowSubcommand *oval_sc.EntityItemStringType `xml:"show_subcommand"`

	ConfigLine *oval_sc.EntityItemStringType `xml:"config_line"`
}

// Element
type ModuleItem struct {
	XMLName xml.Name `xml:module_item`

	ModuleNumber *oval_sc.EntityItemIntType `xml:"module_number"`

	Type *oval_sc.EntityItemStringType `xml:"type"`

	Model *oval_sc.EntityItemStringType `xml:"model"`

	SoftwareMajorRelease *oval_sc.EntityItemVersionType `xml:"software_major_release"`

	SoftwareIndividualRelease *oval_sc.EntityItemIntType `xml:"software_individual_release"`

	SoftwareVersionId *oval_sc.EntityItemStringType `xml:"software_version_id"`

	HardwareMajorRelease *oval_sc.EntityItemVersionType `xml:"hardware_major_release"`

	HardwareIndividualRelease *oval_sc.EntityItemIntType `xml:"hardware_individual_release"`

	FirmwareMajorRelease *oval_sc.EntityItemVersionType `xml:"firmware_major_release"`

	FirmwareIndividualRelease *oval_sc.EntityItemIntType `xml:"firmware_individual_release"`
}

// Element
type VersionItem struct {
	XMLName xml.Name `xml:version_item`

	SwitchSeries *oval_sc.EntityItemStringType `xml:"switch_series"`

	ImageName *oval_sc.EntityItemStringType `xml:"image_name"`

	CatosRelease *oval_sc.EntityItemVersionType `xml:"catos_release"`

	CatosMajorRelease *oval_sc.EntityItemVersionType `xml:"catos_major_release"`

	CatosIndividualRelease *oval_sc.EntityItemIntType `xml:"catos_individual_release"`

	CatosVersionId *oval_sc.EntityItemStringType `xml:"catos_version_id"`
}

// XSD ComplexType declarations
