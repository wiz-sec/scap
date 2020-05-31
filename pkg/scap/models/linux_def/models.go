// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
package linux_def

import (
	"encoding/xml"
)

// Element
type ApparmorstatusTest struct {
	XMLName xml.Name `xml:apparmorstatus_test`
}

// Element
type ApparmorstatusObject struct {
	XMLName xml.Name `xml:apparmorstatus_object`
}

// Element
type ApparmorstatusState struct {
	XMLName xml.Name `xml:apparmorstatus_state`
}

// Element
type DpkginfoTest struct {
	XMLName xml.Name `xml:dpkginfo_test`
}

// Element
type DpkginfoObject struct {
	XMLName xml.Name `xml:dpkginfo_object`
}

// Element
type DpkginfoState struct {
	XMLName xml.Name `xml:dpkginfo_state`
}

// Element
type IflistenersTest struct {
	XMLName xml.Name `xml:iflisteners_test`
}

// Element
type IflistenersObject struct {
	XMLName xml.Name `xml:iflisteners_object`
}

// Element
type IflistenersState struct {
	XMLName xml.Name `xml:iflisteners_state`
}

// Element
type InetlisteningserversTest struct {
	XMLName xml.Name `xml:inetlisteningservers_test`
}

// Element
type InetlisteningserversObject struct {
	XMLName xml.Name `xml:inetlisteningservers_object`
}

// Element
type InetlisteningserversState struct {
	XMLName xml.Name `xml:inetlisteningservers_state`
}

// Element
type PartitionTest struct {
	XMLName xml.Name `xml:partition_test`
}

// Element
type PartitionObject struct {
	XMLName xml.Name `xml:partition_object`
}

// Element
type PartitionState struct {
	XMLName xml.Name `xml:partition_state`
}

// Element
type RpminfoTest struct {
	XMLName xml.Name `xml:rpminfo_test`
}

// Element
type RpminfoObject struct {
	XMLName xml.Name `xml:rpminfo_object`
}

// Element
type RpminfoState struct {
	XMLName xml.Name `xml:rpminfo_state`
}

// Element
type RpmverifyTest struct {
	XMLName xml.Name `xml:rpmverify_test`
}

// Element
type RpmverifyObject struct {
	XMLName xml.Name `xml:rpmverify_object`
}

// Element
type RpmverifyState struct {
	XMLName xml.Name `xml:rpmverify_state`
}

// Element
type RpmverifyfileTest struct {
	XMLName xml.Name `xml:rpmverifyfile_test`
}

// Element
type RpmverifyfileObject struct {
	XMLName xml.Name `xml:rpmverifyfile_object`
}

// Element
type RpmverifyfileState struct {
	XMLName xml.Name `xml:rpmverifyfile_state`
}

// Element
type RpmverifypackageTest struct {
	XMLName xml.Name `xml:rpmverifypackage_test`
}

// Element
type RpmverifypackageObject struct {
	XMLName xml.Name `xml:rpmverifypackage_object`
}

// Element
type RpmverifypackageState struct {
	XMLName xml.Name `xml:rpmverifypackage_state`
}

// Element
type SelinuxbooleanTest struct {
	XMLName xml.Name `xml:selinuxboolean_test`
}

// Element
type SelinuxbooleanObject struct {
	XMLName xml.Name `xml:selinuxboolean_object`
}

// Element
type SelinuxbooleanState struct {
	XMLName xml.Name `xml:selinuxboolean_state`
}

// Element
type SelinuxsecuritycontextTest struct {
	XMLName xml.Name `xml:selinuxsecuritycontext_test`
}

// Element
type SelinuxsecuritycontextObject struct {
	XMLName xml.Name `xml:selinuxsecuritycontext_object`
}

// Element
type SelinuxsecuritycontextState struct {
	XMLName xml.Name `xml:selinuxsecuritycontext_state`
}

// Element
type SlackwarepkginfoTest struct {
	XMLName xml.Name `xml:slackwarepkginfo_test`
}

// Element
type SlackwarepkginfoObject struct {
	XMLName xml.Name `xml:slackwarepkginfo_object`
}

// Element
type SlackwarepkginfoState struct {
	XMLName xml.Name `xml:slackwarepkginfo_state`
}

// Element
type SystemdunitdependencyTest struct {
	XMLName xml.Name `xml:systemdunitdependency_test`
}

// Element
type SystemdunitdependencyObject struct {
	XMLName xml.Name `xml:systemdunitdependency_object`
}

// Element
type SystemdunitdependencyState struct {
	XMLName xml.Name `xml:systemdunitdependency_state`
}

// Element
type SystemdunitpropertyTest struct {
	XMLName xml.Name `xml:systemdunitproperty_test`
}

// Element
type SystemdunitpropertyObject struct {
	XMLName xml.Name `xml:systemdunitproperty_object`
}

// Element
type SystemdunitpropertyState struct {
	XMLName xml.Name `xml:systemdunitproperty_state`
}

// XSD ComplexType declarations

type RpmInfoBehaviors struct {
	Filepaths string `xml:"filepaths,attr"`
}

type RpmVerifyBehaviors struct {
	Nodeps string `xml:"nodeps,attr"`

	Nodigest string `xml:"nodigest,attr"`

	Nofiles string `xml:"nofiles,attr"`

	Noscripts string `xml:"noscripts,attr"`

	Nosignature string `xml:"nosignature,attr"`

	Nolinkto string `xml:"nolinkto,attr"`

	Nomd5 string `xml:"nomd5,attr"`

	Nosize string `xml:"nosize,attr"`

	Nouser string `xml:"nouser,attr"`

	Nogroup string `xml:"nogroup,attr"`

	Nomtime string `xml:"nomtime,attr"`

	Nomode string `xml:"nomode,attr"`

	Nordev string `xml:"nordev,attr"`

	Noconfigfiles string `xml:"noconfigfiles,attr"`

	Noghostfiles string `xml:"noghostfiles,attr"`
}

type RpmVerifyFileBehaviors struct {
	Nolinkto string `xml:"nolinkto,attr"`

	Nomd5 string `xml:"nomd5,attr"`

	Nosize string `xml:"nosize,attr"`

	Nouser string `xml:"nouser,attr"`

	Nogroup string `xml:"nogroup,attr"`

	Nomtime string `xml:"nomtime,attr"`

	Nomode string `xml:"nomode,attr"`

	Nordev string `xml:"nordev,attr"`

	Noconfigfiles string `xml:"noconfigfiles,attr"`

	Noghostfiles string `xml:"noghostfiles,attr"`

	Nofiledigest string `xml:"nofiledigest,attr"`

	Nocaps string `xml:"nocaps,attr"`
}

type RpmVerifyPackageBehaviors struct {
	Nodeps string `xml:"nodeps,attr"`

	Nodigest string `xml:"nodigest,attr"`

	Noscripts string `xml:"noscripts,attr"`

	Nosignature string `xml:"nosignature,attr"`
}

type FileBehaviors struct {
	MaxDepth string `xml:"max_depth,attr"`

	Recurse string `xml:"recurse,attr"`

	RecurseDirection string `xml:"recurse_direction,attr"`

	RecurseFileSystem string `xml:"recurse_file_system,attr"`
}

type EntityStateRpmVerifyResultType struct {
}

type EntityStateProtocolType struct {
}
