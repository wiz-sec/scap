// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
package asa_def

import (
	"encoding/xml"
)

// Element
type AclTest struct {
	XMLName xml.Name `xml:acl_test`
}

// Element
type AclObject struct {
	XMLName xml.Name `xml:acl_object`
}

// Element
type AclState struct {
	XMLName xml.Name `xml:acl_state`
}

// Element
type ClassMapTest struct {
	XMLName xml.Name `xml:class_map_test`
}

// Element
type ClassMapObject struct {
	XMLName xml.Name `xml:class_map_object`
}

// Element
type ClassMapState struct {
	XMLName xml.Name `xml:class_map_state`
}

// Element
type InterfaceTest struct {
	XMLName xml.Name `xml:interface_test`
}

// Element
type InterfaceObject struct {
	XMLName xml.Name `xml:interface_object`
}

// Element
type InterfaceState struct {
	XMLName xml.Name `xml:interface_state`
}

// Element
type LineTest struct {
	XMLName xml.Name `xml:line_test`
}

// Element
type LineObject struct {
	XMLName xml.Name `xml:line_object`
}

// Element
type LineState struct {
	XMLName xml.Name `xml:line_state`
}

// Element
type PolicyMapTest struct {
	XMLName xml.Name `xml:policy_map_test`
}

// Element
type PolicyMapObject struct {
	XMLName xml.Name `xml:policy_map_object`
}

// Element
type PolicyMapState struct {
	XMLName xml.Name `xml:policy_map_state`
}

// Element
type ServicePolicyTest struct {
	XMLName xml.Name `xml:service_policy_test`
}

// Element
type ServicePolicyObject struct {
	XMLName xml.Name `xml:service_policy_object`
}

// Element
type ServicePolicyState struct {
	XMLName xml.Name `xml:service_policy_state`
}

// Element
type SnmpHostTest struct {
	XMLName xml.Name `xml:snmp_host_test`
}

// Element
type SnmpHostObject struct {
	XMLName xml.Name `xml:snmp_host_object`
}

// Element
type SnmpHostState struct {
	XMLName xml.Name `xml:snmp_host_state`
}

// Element
type SnmpUserTest struct {
	XMLName xml.Name `xml:snmp_user_test`
}

// Element
type SnmpUserObject struct {
	XMLName xml.Name `xml:snmp_user_object`
}

// Element
type SnmpUserState struct {
	XMLName xml.Name `xml:snmp_user_state`
}

// Element
type SnmpGroupTest struct {
	XMLName xml.Name `xml:snmp_group_test`
}

// Element
type SnmpGroupObject struct {
	XMLName xml.Name `xml:snmp_group_object`
}

// Element
type SnmpGroupState struct {
	XMLName xml.Name `xml:snmp_group_state`
}

// Element
type TcpMapTest struct {
	XMLName xml.Name `xml:tcp_map_test`
}

// Element
type TcpMapObject struct {
	XMLName xml.Name `xml:tcp_map_object`
}

// Element
type TcpMapState struct {
	XMLName xml.Name `xml:tcp_map_state`
}

// Element
type VersionTest struct {
	XMLName xml.Name `xml:version_test`
}

// Element
type VersionObject struct {
	XMLName xml.Name `xml:version_object`
}

// Element
type VersionState struct {
	XMLName xml.Name `xml:version_state`
}

// XSD ComplexType declarations

type EntityObjectAccessListIPVersionType string

type EntityStateAccessListIPVersionType string

type EntityStateAccessListUseType string

type EntityStateAccessListInterfaceDirectionType string

type EntityStateClassMapType string

type EntityStateInspectionType string

type EntityStateApplyServicePolicyType string

type EntityStateMatchType string

type EntityStateSNMPVersionStringType string

type EntityStateSNMPSecLevelStringType string

type EntityStateSNMPAuthStringType string

type EntityStateSNMPPrivStringType string
