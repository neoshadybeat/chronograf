// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: predicate.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Node_Type int32

const (
	NodeTypeLogicalExpression    Node_Type = 0
	NodeTypeComparisonExpression Node_Type = 1
	NodeTypeParenExpression      Node_Type = 2
	NodeTypeTagRef               Node_Type = 3
	NodeTypeLiteral              Node_Type = 4
	NodeTypeFieldRef             Node_Type = 5
)

var Node_Type_name = map[int32]string{
	0: "LOGICAL_EXPRESSION",
	1: "COMPARISON_EXPRESSION",
	2: "PAREN_EXPRESSION",
	3: "TAG_REF",
	4: "LITERAL",
	5: "FIELD_REF",
}
var Node_Type_value = map[string]int32{
	"LOGICAL_EXPRESSION":    0,
	"COMPARISON_EXPRESSION": 1,
	"PAREN_EXPRESSION":      2,
	"TAG_REF":               3,
	"LITERAL":               4,
	"FIELD_REF":             5,
}

func (x Node_Type) String() string {
	return proto.EnumName(Node_Type_name, int32(x))
}
func (Node_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorPredicate, []int{0, 0} }

type Node_Comparison int32

const (
	ComparisonEqual        Node_Comparison = 0
	ComparisonNotEqual     Node_Comparison = 1
	ComparisonStartsWith   Node_Comparison = 2
	ComparisonRegex        Node_Comparison = 3
	ComparisonNotRegex     Node_Comparison = 4
	ComparisonLess         Node_Comparison = 5
	ComparisonLessEqual    Node_Comparison = 6
	ComparisonGreater      Node_Comparison = 7
	ComparisonGreaterEqual Node_Comparison = 8
)

var Node_Comparison_name = map[int32]string{
	0: "EQUAL",
	1: "NOT_EQUAL",
	2: "STARTS_WITH",
	3: "REGEX",
	4: "NOT_REGEX",
	5: "LT",
	6: "LTE",
	7: "GT",
	8: "GTE",
}
var Node_Comparison_value = map[string]int32{
	"EQUAL":       0,
	"NOT_EQUAL":   1,
	"STARTS_WITH": 2,
	"REGEX":       3,
	"NOT_REGEX":   4,
	"LT":          5,
	"LTE":         6,
	"GT":          7,
	"GTE":         8,
}

func (x Node_Comparison) String() string {
	return proto.EnumName(Node_Comparison_name, int32(x))
}
func (Node_Comparison) EnumDescriptor() ([]byte, []int) { return fileDescriptorPredicate, []int{0, 1} }

// Logical operators apply to boolean values and combine to produce a single boolean result.
type Node_Logical int32

const (
	LogicalAnd Node_Logical = 0
	LogicalOr  Node_Logical = 1
)

var Node_Logical_name = map[int32]string{
	0: "AND",
	1: "OR",
}
var Node_Logical_value = map[string]int32{
	"AND": 0,
	"OR":  1,
}

func (x Node_Logical) String() string {
	return proto.EnumName(Node_Logical_name, int32(x))
}
func (Node_Logical) EnumDescriptor() ([]byte, []int) { return fileDescriptorPredicate, []int{0, 2} }

type Node struct {
	NodeType Node_Type `protobuf:"varint,1,opt,name=node_type,json=nodeType,proto3,enum=com.github.influxdata.influxdb.services.storage.Node_Type" json:"nodeType"`
	Children []*Node   `protobuf:"bytes,2,rep,name=children" json:"children,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Node_StringValue
	//	*Node_BooleanValue
	//	*Node_IntegerValue
	//	*Node_UnsignedValue
	//	*Node_FloatValue
	//	*Node_RegexValue
	//	*Node_TagRefValue
	//	*Node_FieldRefValue
	//	*Node_Logical_
	//	*Node_Comparison_
	Value isNode_Value `protobuf_oneof:"value"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptorPredicate, []int{0} }

type isNode_Value interface {
	isNode_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Node_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}
type Node_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof"`
}
type Node_IntegerValue struct {
	IntegerValue int64 `protobuf:"varint,5,opt,name=int_value,json=intValue,proto3,oneof"`
}
type Node_UnsignedValue struct {
	UnsignedValue uint64 `protobuf:"varint,6,opt,name=uint_value,json=uintValue,proto3,oneof"`
}
type Node_FloatValue struct {
	FloatValue float64 `protobuf:"fixed64,7,opt,name=float_value,json=floatValue,proto3,oneof"`
}
type Node_RegexValue struct {
	RegexValue string `protobuf:"bytes,8,opt,name=regex_value,json=regexValue,proto3,oneof"`
}
type Node_TagRefValue struct {
	TagRefValue string `protobuf:"bytes,9,opt,name=tag_ref_value,json=tagRefValue,proto3,oneof"`
}
type Node_FieldRefValue struct {
	FieldRefValue string `protobuf:"bytes,10,opt,name=field_ref_value,json=fieldRefValue,proto3,oneof"`
}
type Node_Logical_ struct {
	Logical Node_Logical `protobuf:"varint,11,opt,name=logical,proto3,enum=com.github.influxdata.influxdb.services.storage.Node_Logical,oneof"`
}
type Node_Comparison_ struct {
	Comparison Node_Comparison `protobuf:"varint,12,opt,name=comparison,proto3,enum=com.github.influxdata.influxdb.services.storage.Node_Comparison,oneof"`
}

func (*Node_StringValue) isNode_Value()   {}
func (*Node_BooleanValue) isNode_Value()  {}
func (*Node_IntegerValue) isNode_Value()  {}
func (*Node_UnsignedValue) isNode_Value() {}
func (*Node_FloatValue) isNode_Value()    {}
func (*Node_RegexValue) isNode_Value()    {}
func (*Node_TagRefValue) isNode_Value()   {}
func (*Node_FieldRefValue) isNode_Value() {}
func (*Node_Logical_) isNode_Value()      {}
func (*Node_Comparison_) isNode_Value()   {}

func (m *Node) GetValue() isNode_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Node) GetNodeType() Node_Type {
	if m != nil {
		return m.NodeType
	}
	return NodeTypeLogicalExpression
}

func (m *Node) GetChildren() []*Node {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Node) GetStringValue() string {
	if x, ok := m.GetValue().(*Node_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Node) GetBooleanValue() bool {
	if x, ok := m.GetValue().(*Node_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (m *Node) GetIntegerValue() int64 {
	if x, ok := m.GetValue().(*Node_IntegerValue); ok {
		return x.IntegerValue
	}
	return 0
}

func (m *Node) GetUnsignedValue() uint64 {
	if x, ok := m.GetValue().(*Node_UnsignedValue); ok {
		return x.UnsignedValue
	}
	return 0
}

func (m *Node) GetFloatValue() float64 {
	if x, ok := m.GetValue().(*Node_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *Node) GetRegexValue() string {
	if x, ok := m.GetValue().(*Node_RegexValue); ok {
		return x.RegexValue
	}
	return ""
}

func (m *Node) GetTagRefValue() string {
	if x, ok := m.GetValue().(*Node_TagRefValue); ok {
		return x.TagRefValue
	}
	return ""
}

func (m *Node) GetFieldRefValue() string {
	if x, ok := m.GetValue().(*Node_FieldRefValue); ok {
		return x.FieldRefValue
	}
	return ""
}

func (m *Node) GetLogical() Node_Logical {
	if x, ok := m.GetValue().(*Node_Logical_); ok {
		return x.Logical
	}
	return LogicalAnd
}

func (m *Node) GetComparison() Node_Comparison {
	if x, ok := m.GetValue().(*Node_Comparison_); ok {
		return x.Comparison
	}
	return ComparisonEqual
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Node) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Node_OneofMarshaler, _Node_OneofUnmarshaler, _Node_OneofSizer, []interface{}{
		(*Node_StringValue)(nil),
		(*Node_BooleanValue)(nil),
		(*Node_IntegerValue)(nil),
		(*Node_UnsignedValue)(nil),
		(*Node_FloatValue)(nil),
		(*Node_RegexValue)(nil),
		(*Node_TagRefValue)(nil),
		(*Node_FieldRefValue)(nil),
		(*Node_Logical_)(nil),
		(*Node_Comparison_)(nil),
	}
}

func _Node_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Node)
	// value
	switch x := m.Value.(type) {
	case *Node_StringValue:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.StringValue)
	case *Node_BooleanValue:
		t := uint64(0)
		if x.BooleanValue {
			t = 1
		}
		_ = b.EncodeVarint(4<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *Node_IntegerValue:
		_ = b.EncodeVarint(5<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.IntegerValue))
	case *Node_UnsignedValue:
		_ = b.EncodeVarint(6<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.UnsignedValue))
	case *Node_FloatValue:
		_ = b.EncodeVarint(7<<3 | proto.WireFixed64)
		_ = b.EncodeFixed64(math.Float64bits(x.FloatValue))
	case *Node_RegexValue:
		_ = b.EncodeVarint(8<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.RegexValue)
	case *Node_TagRefValue:
		_ = b.EncodeVarint(9<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.TagRefValue)
	case *Node_FieldRefValue:
		_ = b.EncodeVarint(10<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.FieldRefValue)
	case *Node_Logical_:
		_ = b.EncodeVarint(11<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Logical))
	case *Node_Comparison_:
		_ = b.EncodeVarint(12<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Comparison))
	case nil:
	default:
		return fmt.Errorf("Node.Value has unexpected type %T", x)
	}
	return nil
}

func _Node_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Node)
	switch tag {
	case 3: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Node_StringValue{x}
		return true, err
	case 4: // value.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Node_BooleanValue{x != 0}
		return true, err
	case 5: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Node_IntegerValue{int64(x)}
		return true, err
	case 6: // value.uint_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Node_UnsignedValue{x}
		return true, err
	case 7: // value.float_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &Node_FloatValue{math.Float64frombits(x)}
		return true, err
	case 8: // value.regex_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Node_RegexValue{x}
		return true, err
	case 9: // value.tag_ref_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Node_TagRefValue{x}
		return true, err
	case 10: // value.field_ref_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Node_FieldRefValue{x}
		return true, err
	case 11: // value.logical
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Node_Logical_{Node_Logical(x)}
		return true, err
	case 12: // value.comparison
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Node_Comparison_{Node_Comparison(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Node_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Node)
	// value
	switch x := m.Value.(type) {
	case *Node_StringValue:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Node_BooleanValue:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += 1
	case *Node_IntegerValue:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntegerValue))
	case *Node_UnsignedValue:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.UnsignedValue))
	case *Node_FloatValue:
		n += proto.SizeVarint(7<<3 | proto.WireFixed64)
		n += 8
	case *Node_RegexValue:
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RegexValue)))
		n += len(x.RegexValue)
	case *Node_TagRefValue:
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TagRefValue)))
		n += len(x.TagRefValue)
	case *Node_FieldRefValue:
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.FieldRefValue)))
		n += len(x.FieldRefValue)
	case *Node_Logical_:
		n += proto.SizeVarint(11<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Logical))
	case *Node_Comparison_:
		n += proto.SizeVarint(12<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Comparison))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Predicate struct {
	Root *Node `protobuf:"bytes,1,opt,name=root" json:"root,omitempty"`
}

func (m *Predicate) Reset()                    { *m = Predicate{} }
func (m *Predicate) String() string            { return proto.CompactTextString(m) }
func (*Predicate) ProtoMessage()               {}
func (*Predicate) Descriptor() ([]byte, []int) { return fileDescriptorPredicate, []int{1} }

func (m *Predicate) GetRoot() *Node {
	if m != nil {
		return m.Root
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "com.github.influxdata.influxdb.services.storage.Node")
	proto.RegisterType((*Predicate)(nil), "com.github.influxdata.influxdb.services.storage.Predicate")
	proto.RegisterEnum("com.github.influxdata.influxdb.services.storage.Node_Type", Node_Type_name, Node_Type_value)
	proto.RegisterEnum("com.github.influxdata.influxdb.services.storage.Node_Comparison", Node_Comparison_name, Node_Comparison_value)
	proto.RegisterEnum("com.github.influxdata.influxdb.services.storage.Node_Logical", Node_Logical_name, Node_Logical_value)
}
func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NodeType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPredicate(dAtA, i, uint64(m.NodeType))
	}
	if len(m.Children) > 0 {
		for _, msg := range m.Children {
			dAtA[i] = 0x12
			i++
			i = encodeVarintPredicate(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Value != nil {
		nn1, err := m.Value.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *Node_StringValue) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x1a
	i++
	i = encodeVarintPredicate(dAtA, i, uint64(len(m.StringValue)))
	i += copy(dAtA[i:], m.StringValue)
	return i, nil
}
func (m *Node_BooleanValue) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x20
	i++
	if m.BooleanValue {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func (m *Node_IntegerValue) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x28
	i++
	i = encodeVarintPredicate(dAtA, i, uint64(m.IntegerValue))
	return i, nil
}
func (m *Node_UnsignedValue) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x30
	i++
	i = encodeVarintPredicate(dAtA, i, uint64(m.UnsignedValue))
	return i, nil
}
func (m *Node_FloatValue) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x39
	i++
	i = encodeFixed64Predicate(dAtA, i, uint64(math.Float64bits(float64(m.FloatValue))))
	return i, nil
}
func (m *Node_RegexValue) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x42
	i++
	i = encodeVarintPredicate(dAtA, i, uint64(len(m.RegexValue)))
	i += copy(dAtA[i:], m.RegexValue)
	return i, nil
}
func (m *Node_TagRefValue) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x4a
	i++
	i = encodeVarintPredicate(dAtA, i, uint64(len(m.TagRefValue)))
	i += copy(dAtA[i:], m.TagRefValue)
	return i, nil
}
func (m *Node_FieldRefValue) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x52
	i++
	i = encodeVarintPredicate(dAtA, i, uint64(len(m.FieldRefValue)))
	i += copy(dAtA[i:], m.FieldRefValue)
	return i, nil
}
func (m *Node_Logical_) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x58
	i++
	i = encodeVarintPredicate(dAtA, i, uint64(m.Logical))
	return i, nil
}
func (m *Node_Comparison_) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x60
	i++
	i = encodeVarintPredicate(dAtA, i, uint64(m.Comparison))
	return i, nil
}
func (m *Predicate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Predicate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Root != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPredicate(dAtA, i, uint64(m.Root.Size()))
		n2, err := m.Root.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64Predicate(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Predicate(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPredicate(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Node) Size() (n int) {
	var l int
	_ = l
	if m.NodeType != 0 {
		n += 1 + sovPredicate(uint64(m.NodeType))
	}
	if len(m.Children) > 0 {
		for _, e := range m.Children {
			l = e.Size()
			n += 1 + l + sovPredicate(uint64(l))
		}
	}
	if m.Value != nil {
		n += m.Value.Size()
	}
	return n
}

func (m *Node_StringValue) Size() (n int) {
	var l int
	_ = l
	l = len(m.StringValue)
	n += 1 + l + sovPredicate(uint64(l))
	return n
}
func (m *Node_BooleanValue) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}
func (m *Node_IntegerValue) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovPredicate(uint64(m.IntegerValue))
	return n
}
func (m *Node_UnsignedValue) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovPredicate(uint64(m.UnsignedValue))
	return n
}
func (m *Node_FloatValue) Size() (n int) {
	var l int
	_ = l
	n += 9
	return n
}
func (m *Node_RegexValue) Size() (n int) {
	var l int
	_ = l
	l = len(m.RegexValue)
	n += 1 + l + sovPredicate(uint64(l))
	return n
}
func (m *Node_TagRefValue) Size() (n int) {
	var l int
	_ = l
	l = len(m.TagRefValue)
	n += 1 + l + sovPredicate(uint64(l))
	return n
}
func (m *Node_FieldRefValue) Size() (n int) {
	var l int
	_ = l
	l = len(m.FieldRefValue)
	n += 1 + l + sovPredicate(uint64(l))
	return n
}
func (m *Node_Logical_) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovPredicate(uint64(m.Logical))
	return n
}
func (m *Node_Comparison_) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovPredicate(uint64(m.Comparison))
	return n
}
func (m *Predicate) Size() (n int) {
	var l int
	_ = l
	if m.Root != nil {
		l = m.Root.Size()
		n += 1 + l + sovPredicate(uint64(l))
	}
	return n
}

func sovPredicate(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPredicate(x uint64) (n int) {
	return sovPredicate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPredicate
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeType", wireType)
			}
			m.NodeType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeType |= (Node_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPredicate
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Children = append(m.Children, &Node{})
			if err := m.Children[len(m.Children)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPredicate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &Node_StringValue{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BooleanValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Value = &Node_BooleanValue{b}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntegerValue", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &Node_IntegerValue{v}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnsignedValue", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &Node_UnsignedValue{v}
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatValue", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Value = &Node_FloatValue{float64(math.Float64frombits(v))}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegexValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPredicate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &Node_RegexValue{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagRefValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPredicate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &Node_TagRefValue{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldRefValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPredicate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &Node_FieldRefValue{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logical", wireType)
			}
			var v Node_Logical
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (Node_Logical(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &Node_Logical_{v}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comparison", wireType)
			}
			var v Node_Comparison
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (Node_Comparison(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &Node_Comparison_{v}
		default:
			iNdEx = preIndex
			skippy, err := skipPredicate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPredicate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Predicate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPredicate
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Predicate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Predicate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPredicate
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Root == nil {
				m.Root = &Node{}
			}
			if err := m.Root.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPredicate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPredicate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPredicate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPredicate
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPredicate
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPredicate
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPredicate(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPredicate = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPredicate   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("predicate.proto", fileDescriptorPredicate) }

var fileDescriptorPredicate = []byte{
	// 883 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x86, 0x45, 0x5d, 0x2c, 0xe9, 0xc8, 0x17, 0x66, 0x12, 0xc7, 0x0a, 0xdb, 0x48, 0x03, 0x07,
	0x05, 0xd4, 0x45, 0x65, 0xd8, 0xad, 0x37, 0x0d, 0x8a, 0x56, 0x72, 0x68, 0x59, 0x00, 0x2b, 0x29,
	0x14, 0x73, 0x69, 0x37, 0x02, 0x25, 0x8d, 0x68, 0x02, 0x34, 0x47, 0x19, 0x8e, 0x02, 0xe7, 0x0d,
	0x0a, 0xae, 0xba, 0x2f, 0xb8, 0xea, 0xcb, 0x74, 0x53, 0xa0, 0x4f, 0x20, 0x14, 0xea, 0xae, 0x7d,
	0x89, 0x82, 0xc3, 0x9b, 0xdc, 0x76, 0x53, 0xef, 0xe6, 0x9c, 0xf9, 0xbf, 0xff, 0xcc, 0x1c, 0x1e,
	0x92, 0x70, 0xb0, 0x64, 0x64, 0x6e, 0xcf, 0x4c, 0x4e, 0xda, 0x4b, 0x46, 0x39, 0x45, 0x27, 0x33,
	0x7a, 0xd3, 0xb6, 0x6c, 0x7e, 0xbd, 0x9a, 0xb6, 0x6d, 0x77, 0xe1, 0xac, 0x6e, 0xe7, 0x26, 0x37,
	0x93, 0xe5, 0xb4, 0xed, 0x11, 0xf6, 0xde, 0x9e, 0x11, 0xaf, 0xed, 0x71, 0xca, 0x4c, 0x8b, 0x28,
	0x9f, 0xc5, 0xe2, 0x19, 0xbd, 0x39, 0xb1, 0xa8, 0x45, 0x4f, 0x84, 0xcf, 0x74, 0xb5, 0x10, 0x91,
	0x08, 0xc4, 0x2a, 0xf2, 0x3f, 0xfe, 0xab, 0x06, 0xc5, 0x01, 0x9d, 0x13, 0xf4, 0x0e, 0xaa, 0x2e,
	0x9d, 0x93, 0x09, 0xff, 0xb0, 0x24, 0x75, 0x09, 0x4b, 0xad, 0xfd, 0xb3, 0x2f, 0xdb, 0xff, 0xb3,
	0x78, 0x3b, 0x74, 0x6a, 0x1b, 0x1f, 0x96, 0xa4, 0x5b, 0xdf, 0xac, 0x9b, 0x95, 0x30, 0x0c, 0xa3,
	0x3f, 0xd7, 0xcd, 0x8a, 0x1b, 0xaf, 0xf5, 0x74, 0x85, 0x5e, 0x42, 0x65, 0x76, 0x6d, 0x3b, 0x73,
	0x46, 0xdc, 0x7a, 0x1e, 0x17, 0x5a, 0xb5, 0xb3, 0xf3, 0x7b, 0x55, 0xd4, 0x53, 0x1b, 0xf4, 0x05,
	0xec, 0x7a, 0x9c, 0xd9, 0xae, 0x35, 0x79, 0x6f, 0x3a, 0x2b, 0x52, 0x2f, 0x60, 0xa9, 0x55, 0xed,
	0x1e, 0x6c, 0xd6, 0xcd, 0xda, 0x58, 0xe4, 0x5f, 0x87, 0xe9, 0xab, 0x9c, 0x5e, 0xf3, 0xb2, 0x10,
	0x9d, 0x02, 0x4c, 0x29, 0x75, 0x62, 0xa6, 0x88, 0xa5, 0x56, 0xa5, 0x2b, 0x6f, 0xd6, 0xcd, 0xdd,
	0x2e, 0xa5, 0x0e, 0x31, 0xdd, 0x04, 0xaa, 0x86, 0xaa, 0x08, 0x39, 0x81, 0xaa, 0xed, 0xf2, 0x98,
	0x28, 0x61, 0xa9, 0x55, 0x88, 0x88, 0xbe, 0xcb, 0x89, 0x45, 0x58, 0x42, 0x54, 0x6c, 0x97, 0x47,
	0xc0, 0x19, 0xc0, 0x2a, 0x23, 0x76, 0xb0, 0xd4, 0x2a, 0x76, 0x1f, 0x6c, 0xd6, 0xcd, 0xbd, 0x57,
	0xae, 0x67, 0x5b, 0x2e, 0x99, 0xa7, 0x45, 0x56, 0x29, 0x73, 0x0a, 0xb5, 0x85, 0x43, 0xcd, 0x04,
	0x2a, 0x63, 0xa9, 0x25, 0x75, 0xf7, 0x37, 0xeb, 0x26, 0x5c, 0x86, 0xe9, 0x84, 0x80, 0x45, 0x1a,
	0x85, 0x08, 0x23, 0x16, 0xb9, 0x8d, 0x91, 0x8a, 0xb8, 0xbf, 0x40, 0xf4, 0x30, 0x9d, 0x22, 0x2c,
	0x8d, 0xd0, 0x39, 0xec, 0x71, 0xd3, 0x9a, 0x30, 0xb2, 0x88, 0xa1, 0x6a, 0xd6, 0x34, 0xc3, 0xb4,
	0x74, 0xb2, 0x48, 0x9b, 0xc6, 0xb3, 0x10, 0x3d, 0x87, 0x83, 0x85, 0x4d, 0x9c, 0xf9, 0x16, 0x08,
	0x02, 0x14, 0xb7, 0xba, 0x0c, 0xb7, 0xb6, 0xd0, 0xbd, 0xc5, 0x76, 0x02, 0x7d, 0x07, 0x65, 0x87,
	0x5a, 0xf6, 0xcc, 0x74, 0xea, 0x35, 0x31, 0x6b, 0x5f, 0xdd, 0x6f, 0xd6, 0xb4, 0xc8, 0xe4, 0x2a,
	0xa7, 0x27, 0x7e, 0x68, 0x0a, 0x30, 0xa3, 0x37, 0x4b, 0x93, 0xd9, 0x1e, 0x75, 0xeb, 0xbb, 0xc2,
	0xfd, 0x9b, 0xfb, 0xb9, 0x5f, 0xa4, 0x3e, 0x61, 0xcb, 0x32, 0xd7, 0xe3, 0x9f, 0xf2, 0x50, 0x14,
	0x23, 0x7c, 0x0e, 0x48, 0x1b, 0xf6, 0xfa, 0x17, 0x1d, 0x6d, 0xa2, 0xbe, 0x1d, 0xe9, 0xea, 0x78,
	0xdc, 0x1f, 0x0e, 0xe4, 0x9c, 0xf2, 0xd4, 0x0f, 0xf0, 0x93, 0x64, 0xfc, 0xe3, 0x43, 0xaa, 0xb7,
	0x4b, 0x46, 0x3c, 0xcf, 0xa6, 0x2e, 0x7a, 0x0e, 0x87, 0x17, 0xc3, 0x6f, 0x47, 0x1d, 0xbd, 0x3f,
	0x1e, 0x0e, 0xb6, 0x49, 0x49, 0xc1, 0x7e, 0x80, 0x3f, 0x4e, 0xc8, 0xec, 0x00, 0x5b, 0xf0, 0x29,
	0xc8, 0xa3, 0x8e, 0xae, 0xde, 0xe1, 0xf2, 0xca, 0x47, 0x7e, 0x80, 0x8f, 0x12, 0x6e, 0x64, 0x32,
	0xb2, 0x8d, 0x34, 0xa1, 0x6c, 0x74, 0x7a, 0x13, 0x5d, 0xbd, 0x94, 0x0b, 0x0a, 0xf2, 0x03, 0xbc,
	0x9f, 0x28, 0xa3, 0x07, 0x8c, 0x30, 0x94, 0xb5, 0xbe, 0xa1, 0xea, 0x1d, 0x4d, 0x2e, 0x2a, 0x0f,
	0xfd, 0x00, 0x1f, 0xa4, 0x87, 0xb7, 0x39, 0x61, 0xa6, 0x83, 0x9e, 0x41, 0xf5, 0xb2, 0xaf, 0x6a,
	0x2f, 0x84, 0x49, 0x49, 0x79, 0xe4, 0x07, 0x58, 0x4e, 0x34, 0xc9, 0xc3, 0x56, 0x8a, 0x3f, 0xfc,
	0xdc, 0xc8, 0x1d, 0xff, 0x9a, 0x07, 0xc8, 0x4e, 0x8e, 0x1a, 0x50, 0x52, 0x5f, 0xbe, 0xea, 0x68,
	0x72, 0x2e, 0x72, 0xde, 0xba, 0xd4, 0xbb, 0x95, 0xe9, 0xa0, 0x4f, 0xa0, 0x3a, 0x18, 0x1a, 0x93,
	0x48, 0x23, 0x29, 0x8f, 0xfd, 0x00, 0xa3, 0x4c, 0x33, 0xa0, 0x3c, 0x92, 0x7d, 0x0a, 0xb5, 0xb1,
	0xd1, 0xd1, 0x8d, 0xf1, 0xe4, 0x4d, 0xdf, 0xb8, 0x92, 0xf3, 0x4a, 0xdd, 0x0f, 0xf0, 0xa3, 0x4c,
	0x38, 0xe6, 0x26, 0xe3, 0xde, 0x1b, 0x9b, 0x5f, 0x87, 0x15, 0x75, 0xb5, 0xa7, 0xbe, 0x95, 0x0b,
	0xff, 0xac, 0x28, 0x5e, 0x82, 0xa4, 0x62, 0xa4, 0x29, 0xfe, 0x47, 0xc5, 0x48, 0xa6, 0x40, 0x5e,
	0x33, 0xe4, 0x52, 0xd4, 0xb0, 0x6c, 0x5f, 0x23, 0x9e, 0x87, 0x30, 0x14, 0x34, 0x43, 0x95, 0x77,
	0x94, 0x23, 0x3f, 0xc0, 0x0f, 0xef, 0x6e, 0x46, 0xe7, 0x7d, 0x0a, 0xf9, 0x9e, 0x21, 0x97, 0x95,
	0x43, 0x3f, 0xc0, 0x0f, 0x32, 0x41, 0x8f, 0x11, 0x93, 0x13, 0x86, 0x9e, 0x41, 0xa1, 0x67, 0xa8,
	0x72, 0x45, 0x51, 0xfc, 0x00, 0x3f, 0xfe, 0xd7, 0xbe, 0xf0, 0x88, 0xfb, 0xf9, 0x35, 0x94, 0xe3,
	0x11, 0x42, 0x47, 0x50, 0xe8, 0x0c, 0x5e, 0xc8, 0x39, 0x65, 0xdf, 0x0f, 0x30, 0xc4, 0xd9, 0x8e,
	0x3b, 0x47, 0x87, 0x90, 0x1f, 0xea, 0xb2, 0xa4, 0xec, 0xf9, 0x01, 0xae, 0xc6, 0xf9, 0x21, 0x8b,
	0x0c, 0xba, 0x65, 0x28, 0x89, 0x17, 0xf4, 0xf8, 0x35, 0x54, 0x47, 0xc9, 0x0f, 0x06, 0xf5, 0xa1,
	0xc8, 0x28, 0xe5, 0xe2, 0x63, 0x7f, 0xef, 0x4f, 0xaf, 0xb0, 0xe8, 0x3e, 0xf9, 0x65, 0xd3, 0x90,
	0x7e, 0xdb, 0x34, 0xa4, 0xdf, 0x37, 0x0d, 0xe9, 0xc7, 0x3f, 0x1a, 0xb9, 0xef, 0xcb, 0xb1, 0x6a,
	0xba, 0x23, 0xfe, 0x33, 0x9f, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x37, 0x3b, 0xbc, 0x43, 0xda,
	0x06, 0x00, 0x00,
}
