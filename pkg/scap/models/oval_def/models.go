// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5
package oval_def

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type OvalDefinitions struct {
	XMLName xml.Name `xml:oval_definitions`

	Generator oval.GeneratorType `xml:"generator"`

	Definitions *DefinitionsType `xml:"definitions"`

	Tests *TestsType `xml:"tests"`

	Objects *ObjectsType `xml:"objects"`

	States *StatesType `xml:"states"`

	Variables *VariablesType `xml:"variables"`

	Signature *xml_dsig.Signature `xml:"Signature"`
}

// Element
type Notes struct {
	XMLName xml.Name `xml:notes`

	Note []string `xml:"note"`
}

// Element
type Definition struct {
	XMLName xml.Name `xml:definition`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Class string `xml:"class,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Metadata MetadataType `xml:"metadata"`

	Notes *oval.Notes `xml:"notes"`

	Criteria *CriteriaType `xml:"criteria"`
}

// Element
type Test struct {
	XMLName xml.Name `xml:test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Object struct {
	XMLName xml.Name `xml:object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Set struct {
	XMLName xml.Name `xml:set`

	SetOperator string `xml:"set_operator,attr,omitempty"`

	Set []Set `xml:"set"`

	ObjectReference string `xml:"object_reference"`

	Filter []Filter `xml:"filter"`
}

// Element
type Filter struct {
	XMLName xml.Name `xml:filter`

	Action string `xml:"action,attr,omitempty"`

	Text string `xml:",chardata"`
}

// Element
type State struct {
	XMLName xml.Name `xml:state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Variable struct {
	XMLName xml.Name `xml:variable`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Datatype string `xml:"datatype,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type ExternalVariable struct {
	XMLName xml.Name `xml:external_variable`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Datatype string `xml:"datatype,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type ConstantVariable struct {
	XMLName xml.Name `xml:constant_variable`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Datatype string `xml:"datatype,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Value []ValueType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type LocalVariable struct {
	XMLName xml.Name `xml:local_variable`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Datatype string `xml:"datatype,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// XSD ComplexType declarations

type DefinitionsType struct {
	Definition []Definition `xml:"definition"`

	InnerXml string `xml:",innerxml"`
}

type DefinitionType struct {
	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Class string `xml:"class,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Metadata MetadataType `xml:"metadata"`

	Notes *oval.Notes `xml:"notes"`

	Criteria *CriteriaType `xml:"criteria"`

	InnerXml string `xml:",innerxml"`
}

type MetadataType struct {
	Title string `xml:"title"`

	Affected []AffectedType `xml:"affected"`

	Reference []ReferenceType `xml:"reference"`

	Description string `xml:"description"`

	InnerXml string `xml:",innerxml"`
}

type AffectedType struct {
	Family string `xml:"family,attr"`

	Platform []string `xml:"platform"`

	Product []string `xml:"product"`

	InnerXml string `xml:",innerxml"`
}

type ReferenceType struct {
	Source string `xml:"source,attr"`

	RefId string `xml:"ref_id,attr"`

	RefUrl string `xml:"ref_url,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type CriteriaType struct {
	ApplicabilityCheck string `xml:"applicability_check,attr,omitempty"`

	Operator string `xml:"operator,attr,omitempty"`

	Negate string `xml:"negate,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Criteria []CriteriaType `xml:"criteria"`

	Criterion []CriterionType `xml:"criterion"`

	ExtendDefinition []ExtendDefinitionType `xml:"extend_definition"`

	InnerXml string `xml:",innerxml"`
}

type CriterionType struct {
	ApplicabilityCheck string `xml:"applicability_check,attr,omitempty"`

	TestRef string `xml:"test_ref,attr"`

	Negate string `xml:"negate,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type ExtendDefinitionType struct {
	ApplicabilityCheck string `xml:"applicability_check,attr,omitempty"`

	DefinitionRef string `xml:"definition_ref,attr"`

	Negate string `xml:"negate,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type TestsType struct {
	Test []Test `xml:"test"`

	InnerXml string `xml:",innerxml"`
}

type TestType struct {
	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type ObjectRefType struct {
	ObjectRef string `xml:"object_ref,attr"`

	InnerXml string `xml:",innerxml"`
}

type StateRefType struct {
	StateRef string `xml:"state_ref,attr"`

	InnerXml string `xml:",innerxml"`
}

type ObjectsType struct {
	Object []Object `xml:"object"`

	InnerXml string `xml:",innerxml"`
}

type ObjectType struct {
	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type StatesType struct {
	State []State `xml:"state"`

	InnerXml string `xml:",innerxml"`
}

type StateType struct {
	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type VariablesType struct {
	Variable []Variable `xml:"variable"`

	InnerXml string `xml:",innerxml"`
}

type VariableType struct {
	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Datatype string `xml:"datatype,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type PossibleValueType struct {
	Hint string `xml:"hint,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type PossibleRestrictionType struct {
	Operator string `xml:"operator,attr,omitempty"`

	Hint string `xml:"hint,attr"`

	Restriction []RestrictionType `xml:"restriction"`

	InnerXml string `xml:",innerxml"`
}

type RestrictionType struct {
	Operation string `xml:"operation,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type ValueType struct {
	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type LiteralComponentType struct {
	Datatype string `xml:"datatype,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type ObjectComponentType struct {
	ObjectRef string `xml:"object_ref,attr"`

	ItemField string `xml:"item_field,attr"`

	RecordField string `xml:"record_field,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type VariableComponentType struct {
	VarRef string `xml:"var_ref,attr"`

	InnerXml string `xml:",innerxml"`
}

type ArithmeticFunctionType struct {
	ArithmeticOperation string `xml:"arithmetic_operation,attr"`

	InnerXml string `xml:",innerxml"`
}

type BeginFunctionType struct {
	Character string `xml:"character,attr"`

	InnerXml string `xml:",innerxml"`
}

type ConcatFunctionType struct {
	InnerXml string `xml:",innerxml"`
}

type EndFunctionType struct {
	Character string `xml:"character,attr"`

	InnerXml string `xml:",innerxml"`
}

type EscapeRegexFunctionType struct {
	InnerXml string `xml:",innerxml"`
}

type SplitFunctionType struct {
	Delimiter string `xml:"delimiter,attr"`

	InnerXml string `xml:",innerxml"`
}

type SubstringFunctionType struct {
	SubstringStart string `xml:"substring_start,attr"`

	SubstringLength string `xml:"substring_length,attr"`

	InnerXml string `xml:",innerxml"`
}

type TimeDifferenceFunctionType struct {
	Format1 string `xml:"format_1,attr,omitempty"`

	Format2 string `xml:"format_2,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type RegexCaptureFunctionType struct {
	Pattern string `xml:"pattern,attr"`

	InnerXml string `xml:",innerxml"`
}

type UniqueFunctionType struct {
	InnerXml string `xml:",innerxml"`
}

type CountFunctionType struct {
	InnerXml string `xml:",innerxml"`
}

type GlobToRegexFunctionType struct {
	GlobNoescape string `xml:"glob_noescape,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntitySimpleBaseType struct {
	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type EntityComplexBaseType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityObjectIPAddressType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityObjectIPAddressStringType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityObjectAnySimpleType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityObjectBinaryType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityObjectBoolType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityObjectFloatType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityObjectIntType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityObjectStringType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityObjectVersionType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityObjectRecordType struct {
	Field []EntityObjectFieldType `xml:"field"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectFieldType struct {
	Name string `xml:"name,attr"`

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type EntityStateSimpleBaseType struct {
	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type EntityStateComplexBaseType struct {
	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateIPAddressType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateIPAddressStringType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateAnySimpleType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateBinaryType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateBoolType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateFloatType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateIntType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateEVRStringType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateDebianEVRStringType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateVersionType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateFileSetRevisionType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateIOSVersionType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateStringType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateRecordType struct {
	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Field []EntityStateFieldType `xml:"field"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateFieldType struct {
	Name string `xml:"name,attr"`

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}
