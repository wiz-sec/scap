// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
package iosxe_def

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_def"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type GlobalTest struct {
	XMLName xml.Name `xml:global_test`

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
type GlobalObject struct {
	XMLName xml.Name `xml:global_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type GlobalState struct {
	XMLName xml.Name `xml:global_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	GlobalCommand *oval_def.EntityStateStringType `xml:"global_command"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

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

	Platform *oval_def.EntityStateStringType `xml:"platform"`

	Rp *oval_def.EntityStateIntType `xml:"rp"`

	Pkg *oval_def.EntityStateStringType `xml:"pkg"`

	VersionString *oval_def.EntityStateStringType `xml:"version_string"`

	MajorRelease *oval_def.EntityStateIntType `xml:"major_release"`

	Release *oval_def.EntityStateIntType `xml:"release"`

	Rebuild *oval_def.EntityStateIntType `xml:"rebuild"`

	Train *oval_def.EntityStateStringType `xml:"train"`

	IosRelease *oval_def.EntityStateStringType `xml:"ios_release"`

	IosTrain *oval_def.EntityStateStringType `xml:"ios_train"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InterfaceTest struct {
	XMLName xml.Name `xml:interface_test`

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
type InterfaceObject struct {
	XMLName xml.Name `xml:interface_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InterfaceState struct {
	XMLName xml.Name `xml:interface_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	IpDirectedBroadcast *oval_def.EntityStateBoolType `xml:"ip_directed_broadcast"`

	ProxyArp *oval_def.EntityStateBoolType `xml:"proxy_arp"`

	Shutdown *oval_def.EntityStateBoolType `xml:"shutdown"`

	HardwareAddr *oval_def.EntityStateStringType `xml:"hardware_addr"`

	Ipv4Address *oval_def.EntityStateIPAddressStringType `xml:"ipv4_address"`

	Ipv6Address *oval_def.EntityStateIPAddressStringType `xml:"ipv6_address"`

	Ipv4AccessList *oval_def.EntityStateStringType `xml:"ipv4_access_list"`

	Ipv6AccessList *oval_def.EntityStateStringType `xml:"ipv6_access_list"`

	CryptoMap *oval_def.EntityStateStringType `xml:"crypto_map"`

	Ipv4UrpfCommand *oval_def.EntityStateStringType `xml:"ipv4_urpf_command"`

	Ipv6UrpfCommand *oval_def.EntityStateStringType `xml:"ipv6_urpf_command"`

	UrpfCommand *oval_def.EntityStateStringType `xml:"urpf_command"`

	SwitchportTrunkEncapsulation *EntityStateTrunkEncapType `xml:"switchport_trunk_encapsulation"`

	SwitchportMode *EntityStateSwitchportModeType `xml:"switchport_mode"`

	SwitchportNativeVlan *InterfaceStateSwitchportNativeVlan `xml:"switchport_native_vlan"`

	SwitchportAccessVlan *InterfaceStateSwitchportAccessVlan `xml:"switchport_access_vlan"`

	SwitchportTrunkedVlans *oval_def.EntityStateStringType `xml:"switchport_trunked_vlans"`

	SwitchportPrunedVlans *oval_def.EntityStateStringType `xml:"switchport_pruned_vlans"`

	SwitchportPortSecurity *oval_def.EntityStateStringType `xml:"switchport_port_security"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SectionTest struct {
	XMLName xml.Name `xml:section_test`

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
type SectionObject struct {
	XMLName xml.Name `xml:section_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SectionState struct {
	XMLName xml.Name `xml:section_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	SectionCommand *oval_def.EntityStateStringType `xml:"section_command"`

	SectionConfigLines *oval_def.EntityStateStringType `xml:"section_config_lines"`

	ConfigLine *oval_def.EntityStateStringType `xml:"config_line"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RouterTest struct {
	XMLName xml.Name `xml:router_test`

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
type RouterObject struct {
	XMLName xml.Name `xml:router_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RouterState struct {
	XMLName xml.Name `xml:router_state`

	Id2 string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Protocol EntityStateRoutingProtocolType `xml:"protocol"`

	Id *oval_def.EntityStateIntType `xml:"id"`

	Network *oval_def.EntityStateStringType `xml:"network"`

	BgpNeighbor *oval_def.EntityStateStringType `xml:"bgp_neighbor"`

	OspfAuthenticationArea *RouterStateOspfAuthenticationArea `xml:"ospf_authentication_area"`

	RouterConfigLines *oval_def.EntityStateStringType `xml:"router_config_lines"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type BgpneighborTest struct {
	XMLName xml.Name `xml:bgpneighbor_test`

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
type BgpneighborObject struct {
	XMLName xml.Name `xml:bgpneighbor_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type BgpneighborState struct {
	XMLName xml.Name `xml:bgpneighbor_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Neighbor *oval_def.EntityStateStringType `xml:"neighbor"`

	Password *oval_def.EntityStateStringType `xml:"password"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RoutingprotocolauthintfTest struct {
	XMLName xml.Name `xml:routingprotocolauthintf_test`

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
type RoutingprotocolauthintfObject struct {
	XMLName xml.Name `xml:routingprotocolauthintf_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RoutingprotocolauthintfState struct {
	XMLName xml.Name `xml:routingprotocolauthintf_state`

	Id2 string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Interface *oval_def.EntityStateStringType `xml:"interface"`

	Protocol *EntityStateRoutingProtocolType `xml:"protocol"`

	Id *oval_def.EntityStateIntType `xml:"id"`

	AuthType *EntityStateRoutingAuthTypeStringType `xml:"auth_type"`

	OspfArea *RoutingprotocolauthintfStateOspfArea `xml:"ospf_area"`

	KeyChain *oval_def.EntityStateStringType `xml:"key_chain"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type AclTest struct {
	XMLName xml.Name `xml:acl_test`

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
type AclObject struct {
	XMLName xml.Name `xml:acl_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type AclState struct {
	XMLName xml.Name `xml:acl_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	IpVersion *EntityStateAccessListIPVersionType `xml:"ip_version"`

	Use *EntityStateAccessListUseType `xml:"use"`

	UsedIn *oval_def.EntityStateStringType `xml:"used_in"`

	InterfaceDirection *EntityStateAccessListInterfaceDirectionType `xml:"interface_direction"`

	AclConfigLines *oval_def.EntityStateStringType `xml:"acl_config_lines"`

	ConfigLine *oval_def.EntityStateStringType `xml:"config_line"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SnmphostTest struct {
	XMLName xml.Name `xml:snmphost_test`

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
type SnmphostObject struct {
	XMLName xml.Name `xml:snmphost_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SnmphostState struct {
	XMLName xml.Name `xml:snmphost_state`

	Id string `xml:"id,attr"`

	Version2 string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Host *oval_def.EntityStateStringType `xml:"host"`

	CommunityOrUser *oval_def.EntityStateStringType `xml:"community_or_user"`

	Version *EntityStateSNMPVersionStringType `xml:"version"`

	Snmpv3SecLevel *EntityStateSNMPSecLevelStringType `xml:"snmpv3_sec_level"`

	Traps *oval_def.EntityStateStringType `xml:"traps"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SnmpcommunityTest struct {
	XMLName xml.Name `xml:snmpcommunity_test`

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
type SnmpcommunityObject struct {
	XMLName xml.Name `xml:snmpcommunity_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SnmpcommunityState struct {
	XMLName xml.Name `xml:snmpcommunity_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	View *oval_def.EntityStateStringType `xml:"view"`

	Mode *EntityStateSNMPModeStringType `xml:"mode"`

	Ipv4Acl *oval_def.EntityStateStringType `xml:"ipv4_acl"`

	Ipv6Acl *oval_def.EntityStateStringType `xml:"ipv6_acl"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SnmpuserTest struct {
	XMLName xml.Name `xml:snmpuser_test`

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
type SnmpuserObject struct {
	XMLName xml.Name `xml:snmpuser_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SnmpuserState struct {
	XMLName xml.Name `xml:snmpuser_state`

	Id string `xml:"id,attr"`

	Version2 string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Group *oval_def.EntityStateStringType `xml:"group"`

	Version *EntityStateSNMPVersionStringType `xml:"version"`

	Ipv4Acl *oval_def.EntityStateStringType `xml:"ipv4_acl"`

	Ipv6Acl *oval_def.EntityStateStringType `xml:"ipv6_acl"`

	Priv *EntityStateSNMPPrivStringType `xml:"priv"`

	Auth *EntityStateSNMPAuthStringType `xml:"auth"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SnmpgroupTest struct {
	XMLName xml.Name `xml:snmpgroup_test`

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
type SnmpgroupObject struct {
	XMLName xml.Name `xml:snmpgroup_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SnmpgroupState struct {
	XMLName xml.Name `xml:snmpgroup_state`

	Id string `xml:"id,attr"`

	Version2 string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Version *EntityStateSNMPVersionStringType `xml:"version"`

	Snmpv3SecLevel *EntityStateSNMPSecLevelStringType `xml:"snmpv3_sec_level"`

	Ipv4Acl *oval_def.EntityStateStringType `xml:"ipv4_acl"`

	Ipv6Acl *oval_def.EntityStateStringType `xml:"ipv6_acl"`

	ReadView *oval_def.EntityStateStringType `xml:"read_view"`

	WriteView *oval_def.EntityStateStringType `xml:"write_view"`

	NotifyView *oval_def.EntityStateStringType `xml:"notify_view"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SnmpviewTest struct {
	XMLName xml.Name `xml:snmpview_test`

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
type SnmpviewObject struct {
	XMLName xml.Name `xml:snmpview_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SnmpviewState struct {
	XMLName xml.Name `xml:snmpview_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	MibFamily *oval_def.EntityStateStringType `xml:"mib_family"`

	Include *oval_def.EntityStateBoolType `xml:"include"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InterfaceStateSwitchportNativeVlan struct {
	XMLName xml.Name `xml:switchport_native_vlan`
}

// Element
type InterfaceStateSwitchportAccessVlan struct {
	XMLName xml.Name `xml:switchport_access_vlan`
}

// Element
type RouterStateOspfAuthenticationArea struct {
	XMLName xml.Name `xml:ospf_authentication_area`
}

// Element
type RoutingprotocolauthintfStateOspfArea struct {
	XMLName xml.Name `xml:ospf_area`
}

// XSD ComplexType declarations

type EntityObjectAccessListIPVersionType struct {
}

type EntityObjectRoutingProtocolType struct {
}

type EntityStateTrunkEncapType struct {
}

type EntityStateSwitchportModeType struct {
}

type EntityStateRoutingProtocolType struct {
}

type EntityStateRoutingAuthTypeStringType struct {
}

type EntityStateSNMPVersionStringType struct {
}

type EntityStateSNMPSecLevelStringType struct {
}

type EntityStateSNMPModeStringType struct {
}

type EntityStateSNMPAuthStringType struct {
}

type EntityStateSNMPPrivStringType struct {
}

type EntityStateAccessListIPVersionType struct {
}

type EntityStateAccessListUseType struct {
}

type EntityStateAccessListInterfaceDirectionType struct {
}
