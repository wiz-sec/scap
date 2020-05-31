// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
package sol_def

import (
	"encoding/xml"

	"github.com/gocomply/scap/pkg/scap/models/oval_def"
)

// Element
type FacetTest struct {
	XMLName xml.Name `xml:facet_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type FacetObject struct {
	XMLName xml.Name `xml:facet_object`

	Set oval_def.Set `xml:"set"`
}

// Element
type FacetState struct {
	XMLName xml.Name `xml:facet_state`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Value *oval_def.EntityStateBoolType `xml:"value"`
}

// Element
type ImageTest struct {
	XMLName xml.Name `xml:image_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type ImageObject struct {
	XMLName xml.Name `xml:image_object`

	Set oval_def.Set `xml:"set"`
}

// Element
type ImageState struct {
	XMLName xml.Name `xml:image_state`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`
}

// Element
type IsainfoTest struct {
	XMLName xml.Name `xml:isainfo_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type IsainfoObject struct {
	XMLName xml.Name `xml:isainfo_object`
}

// Element
type IsainfoState struct {
	XMLName xml.Name `xml:isainfo_state`

	Bits *oval_def.EntityStateIntType `xml:"bits"`

	KernelIsa *oval_def.EntityStateStringType `xml:"kernel_isa"`

	ApplicationIsa *oval_def.EntityStateStringType `xml:"application_isa"`
}

// Element
type NddTest struct {
	XMLName xml.Name `xml:ndd_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type NddObject struct {
	XMLName xml.Name `xml:ndd_object`

	Set oval_def.Set `xml:"set"`
}

// Element
type NddState struct {
	XMLName xml.Name `xml:ndd_state`

	Device *oval_def.EntityStateStringType `xml:"device"`

	Instance *oval_def.EntityStateIntType `xml:"instance"`

	Parameter *oval_def.EntityStateStringType `xml:"parameter"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`
}

// Element
type PackageTest struct {
	XMLName xml.Name `xml:package_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type PackageObject struct {
	XMLName xml.Name `xml:package_object`

	Set oval_def.Set `xml:"set"`
}

// Element
type PackageState struct {
	XMLName xml.Name `xml:package_state`

	Pkginst *oval_def.EntityStateStringType `xml:"pkginst"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Category *oval_def.EntityStateStringType `xml:"category"`

	Version *oval_def.EntityStateStringType `xml:"version"`

	Vendor *oval_def.EntityStateStringType `xml:"vendor"`

	Description *oval_def.EntityStateStringType `xml:"description"`
}

// Element
type Package511Test struct {
	XMLName xml.Name `xml:package511_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type Package511Object struct {
	XMLName xml.Name `xml:package511_object`

	Set oval_def.Set `xml:"set"`
}

// Element
type Package511State struct {
	XMLName xml.Name `xml:package511_state`

	Publisher *oval_def.EntityStateStringType `xml:"publisher"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Version *oval_def.EntityStateVersionType `xml:"version"`

	Timestamp *oval_def.EntityStateStringType `xml:"timestamp"`

	Fmri *oval_def.EntityStateStringType `xml:"fmri"`

	Summary *oval_def.EntityStateStringType `xml:"summary"`

	Description *oval_def.EntityStateStringType `xml:"description"`

	Category *oval_def.EntityStateStringType `xml:"category"`

	UpdatesAvailable *oval_def.EntityStateBoolType `xml:"updates_available"`
}

// Element
type PackageavoidlistTest struct {
	XMLName xml.Name `xml:packageavoidlist_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type PackageavoidlistObject struct {
	XMLName xml.Name `xml:packageavoidlist_object`
}

// Element
type PackageavoidlistState struct {
	XMLName xml.Name `xml:packageavoidlist_state`

	Fmri *oval_def.EntityStateStringType `xml:"fmri"`
}

// Element
type PackagecheckTest struct {
	XMLName xml.Name `xml:packagecheck_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type PackagecheckObject struct {
	XMLName xml.Name `xml:packagecheck_object`

	Set oval_def.Set `xml:"set"`
}

// Element
type PackagecheckState struct {
	XMLName xml.Name `xml:packagecheck_state`

	Pkginst *oval_def.EntityStateStringType `xml:"pkginst"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	ChecksumDiffers *oval_def.EntityStateBoolType `xml:"checksum_differs"`

	SizeDiffers *oval_def.EntityStateBoolType `xml:"size_differs"`

	MtimeDiffers *oval_def.EntityStateBoolType `xml:"mtime_differs"`

	Uread *EntityStatePermissionCompareType `xml:"uread"`

	Uwrite *EntityStatePermissionCompareType `xml:"uwrite"`

	Uexec *EntityStatePermissionCompareType `xml:"uexec"`

	Gread *EntityStatePermissionCompareType `xml:"gread"`

	Gwrite *EntityStatePermissionCompareType `xml:"gwrite"`

	Gexec *EntityStatePermissionCompareType `xml:"gexec"`

	Oread *EntityStatePermissionCompareType `xml:"oread"`

	Owrite *EntityStatePermissionCompareType `xml:"owrite"`

	Oexec *EntityStatePermissionCompareType `xml:"oexec"`
}

// Element
type PackagefreezelistTest struct {
	XMLName xml.Name `xml:packagefreezelist_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type PackagefreezelistObject struct {
	XMLName xml.Name `xml:packagefreezelist_object`
}

// Element
type PackagefreezelistState struct {
	XMLName xml.Name `xml:packagefreezelist_state`

	Fmri *oval_def.EntityStateStringType `xml:"fmri"`
}

// Element
type PackagepublisherTest struct {
	XMLName xml.Name `xml:packagepublisher_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type PackagepublisherObject struct {
	XMLName xml.Name `xml:packagepublisher_object`

	Set oval_def.Set `xml:"set"`
}

// Element
type PackagepublisherState struct {
	XMLName xml.Name `xml:packagepublisher_state`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Type *EntityStatePublisherTypeType `xml:"type"`

	OriginUri *oval_def.EntityStateStringType `xml:"origin_uri"`

	Alias *oval_def.EntityStateStringType `xml:"alias"`

	SslKey *oval_def.EntityStateStringType `xml:"ssl_key"`

	SslCert *oval_def.EntityStateStringType `xml:"ssl_cert"`

	ClientUuid *EntityStateClientUUIDType `xml:"client_uuid"`

	CatalogUpdated *oval_def.EntityStateIntType `xml:"catalog_updated"`

	Enabled *oval_def.EntityStateBoolType `xml:"enabled"`

	Order *oval_def.EntityStateIntType `xml:"order"`

	Properties *oval_def.EntityStateRecordType `xml:"properties"`
}

// Element
type Patch54Test struct {
	XMLName xml.Name `xml:patch54_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type PatchTest struct {
	XMLName xml.Name `xml:patch_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type Patch54Object struct {
	XMLName xml.Name `xml:patch54_object`

	Set oval_def.Set `xml:"set"`
}

// Element
type PatchObject struct {
	XMLName xml.Name `xml:patch_object`

	Set oval_def.Set `xml:"set"`

	Base oval_def.EntityObjectIntType `xml:"base"`
}

// Element
type PatchState struct {
	XMLName xml.Name `xml:patch_state`

	Base *oval_def.EntityStateIntType `xml:"base"`

	Version *oval_def.EntityStateIntType `xml:"version"`
}

// Element
type SmfTest struct {
	XMLName xml.Name `xml:smf_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type SmfObject struct {
	XMLName xml.Name `xml:smf_object`

	Set oval_def.Set `xml:"set"`
}

// Element
type SmfState struct {
	XMLName xml.Name `xml:smf_state`

	Fmri *oval_def.EntityStateStringType `xml:"fmri"`

	ServiceName *oval_def.EntityStateStringType `xml:"service_name"`

	ServiceState *EntityStateSmfServiceStateType `xml:"service_state"`

	Protocol *oval_def.EntityStateStringType `xml:"protocol"`

	ServerExecutable *oval_def.EntityStateStringType `xml:"server_executable"`

	ServerArguements *oval_def.EntityStateStringType `xml:"server_arguements"`

	ExecAsUser *oval_def.EntityStateStringType `xml:"exec_as_user"`
}

// Element
type SmfpropertyTest struct {
	XMLName xml.Name `xml:smfproperty_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type SmfpropertyObject struct {
	XMLName xml.Name `xml:smfproperty_object`

	Set oval_def.Set `xml:"set"`
}

// Element
type SmfpropertyState struct {
	XMLName xml.Name `xml:smfproperty_state`

	Service *oval_def.EntityStateStringType `xml:"service"`

	Instance *oval_def.EntityStateStringType `xml:"instance"`

	Property *oval_def.EntityStateStringType `xml:"property"`

	Fmri *oval_def.EntityStateStringType `xml:"fmri"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`
}

// Element
type VariantTest struct {
	XMLName xml.Name `xml:variant_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type VariantObject struct {
	XMLName xml.Name `xml:variant_object`

	Set oval_def.Set `xml:"set"`
}

// Element
type VariantState struct {
	XMLName xml.Name `xml:variant_state`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`
}

// Element
type VirtualizationinfoTest struct {
	XMLName xml.Name `xml:virtualizationinfo_test`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`
}

// Element
type VirtualizationinfoObject struct {
	XMLName xml.Name `xml:virtualizationinfo_object`
}

// Element
type VirtualizationinfoState struct {
	XMLName xml.Name `xml:virtualizationinfo_state`

	Current *oval_def.EntityStateStringType `xml:"current"`

	Supported *EntityStateV12NEnvType `xml:"supported"`

	Parent *EntityStateV12NEnvType `xml:"parent"`

	LdomRole *EntityStateLDOMRoleType `xml:"ldom-role"`

	Properties *oval_def.EntityStateRecordType `xml:"properties"`
}

// XSD ComplexType declarations

type PackageCheckBehaviors struct {
	FileattributesOnly string `xml:"fileattributes_only,attr"`

	FilecontentsOnly string `xml:"filecontents_only,attr"`

	NoVolatileeditable string `xml:"no_volatileeditable,attr"`
}

type PatchBehaviors struct {
	Supersedence string `xml:"supersedence,attr"`
}

type EntityObjectPublisherTypeType struct {
}

type EntityStateClientUUIDType struct {
}

type EntityStatePermissionCompareType struct {
}

type EntityStatePublisherTypeType struct {
}

type EntityStateSmfServiceStateType struct {
}

type EntityStateV12NEnvType struct {
}

type EntityStateLDOMRoleType struct {
}
