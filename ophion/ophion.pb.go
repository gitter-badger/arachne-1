// Code generated by protoc-gen-go.
// source: ophion.proto
// DO NOT EDIT!

/*
Package ophion is a generated protocol buffer package.

It is generated from these files:
	ophion.proto

It has these top-level messages:
	GraphQuery
	GraphStatement
	PropertyStatement
	HasStatement
	SelectStatement
	Vertex
	Edge
	QueryResult
*/
package ophion

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GraphQuery struct {
	Query []*GraphStatement `protobuf:"bytes,1,rep,name=query" json:"query,omitempty"`
}

func (m *GraphQuery) Reset()                    { *m = GraphQuery{} }
func (m *GraphQuery) String() string            { return proto.CompactTextString(m) }
func (*GraphQuery) ProtoMessage()               {}
func (*GraphQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GraphQuery) GetQuery() []*GraphStatement {
	if m != nil {
		return m.Query
	}
	return nil
}

type GraphStatement struct {
	// Types that are valid to be assigned to Statement:
	//	*GraphStatement_V
	//	*GraphStatement_E
	//	*GraphStatement_Label
	//	*GraphStatement_Has
	//	*GraphStatement_As
	//	*GraphStatement_In
	//	*GraphStatement_Out
	//	*GraphStatement_InEdge
	//	*GraphStatement_OutEdge
	//	*GraphStatement_InVertex
	//	*GraphStatement_OutVertex
	//	*GraphStatement_Select
	//	*GraphStatement_Limit
	//	*GraphStatement_Count
	//	*GraphStatement_AddV
	//	*GraphStatement_Property
	//	*GraphStatement_AddE
	//	*GraphStatement_To
	Statement isGraphStatement_Statement `protobuf_oneof:"statement"`
}

func (m *GraphStatement) Reset()                    { *m = GraphStatement{} }
func (m *GraphStatement) String() string            { return proto.CompactTextString(m) }
func (*GraphStatement) ProtoMessage()               {}
func (*GraphStatement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isGraphStatement_Statement interface {
	isGraphStatement_Statement()
}

type GraphStatement_V struct {
	V string `protobuf:"bytes,1,opt,name=V,oneof"`
}
type GraphStatement_E struct {
	E string `protobuf:"bytes,2,opt,name=E,oneof"`
}
type GraphStatement_Label struct {
	Label string `protobuf:"bytes,3,opt,name=label,oneof"`
}
type GraphStatement_Has struct {
	Has *HasStatement `protobuf:"bytes,4,opt,name=has,oneof"`
}
type GraphStatement_As struct {
	As string `protobuf:"bytes,5,opt,name=as,oneof"`
}
type GraphStatement_In struct {
	In string `protobuf:"bytes,6,opt,name=in,oneof"`
}
type GraphStatement_Out struct {
	Out string `protobuf:"bytes,7,opt,name=out,oneof"`
}
type GraphStatement_InEdge struct {
	InEdge string `protobuf:"bytes,8,opt,name=inEdge,oneof"`
}
type GraphStatement_OutEdge struct {
	OutEdge string `protobuf:"bytes,9,opt,name=outEdge,oneof"`
}
type GraphStatement_InVertex struct {
	InVertex string `protobuf:"bytes,10,opt,name=inVertex,oneof"`
}
type GraphStatement_OutVertex struct {
	OutVertex string `protobuf:"bytes,11,opt,name=outVertex,oneof"`
}
type GraphStatement_Select struct {
	Select *SelectStatement `protobuf:"bytes,12,opt,name=select,oneof"`
}
type GraphStatement_Limit struct {
	Limit int64 `protobuf:"varint,13,opt,name=limit,oneof"`
}
type GraphStatement_Count struct {
	Count string `protobuf:"bytes,14,opt,name=count,oneof"`
}
type GraphStatement_AddV struct {
	AddV string `protobuf:"bytes,15,opt,name=addV,oneof"`
}
type GraphStatement_Property struct {
	Property *PropertyStatement `protobuf:"bytes,16,opt,name=property,oneof"`
}
type GraphStatement_AddE struct {
	AddE string `protobuf:"bytes,17,opt,name=addE,oneof"`
}
type GraphStatement_To struct {
	To string `protobuf:"bytes,18,opt,name=to,oneof"`
}

func (*GraphStatement_V) isGraphStatement_Statement()         {}
func (*GraphStatement_E) isGraphStatement_Statement()         {}
func (*GraphStatement_Label) isGraphStatement_Statement()     {}
func (*GraphStatement_Has) isGraphStatement_Statement()       {}
func (*GraphStatement_As) isGraphStatement_Statement()        {}
func (*GraphStatement_In) isGraphStatement_Statement()        {}
func (*GraphStatement_Out) isGraphStatement_Statement()       {}
func (*GraphStatement_InEdge) isGraphStatement_Statement()    {}
func (*GraphStatement_OutEdge) isGraphStatement_Statement()   {}
func (*GraphStatement_InVertex) isGraphStatement_Statement()  {}
func (*GraphStatement_OutVertex) isGraphStatement_Statement() {}
func (*GraphStatement_Select) isGraphStatement_Statement()    {}
func (*GraphStatement_Limit) isGraphStatement_Statement()     {}
func (*GraphStatement_Count) isGraphStatement_Statement()     {}
func (*GraphStatement_AddV) isGraphStatement_Statement()      {}
func (*GraphStatement_Property) isGraphStatement_Statement()  {}
func (*GraphStatement_AddE) isGraphStatement_Statement()      {}
func (*GraphStatement_To) isGraphStatement_Statement()        {}

func (m *GraphStatement) GetStatement() isGraphStatement_Statement {
	if m != nil {
		return m.Statement
	}
	return nil
}

func (m *GraphStatement) GetV() string {
	if x, ok := m.GetStatement().(*GraphStatement_V); ok {
		return x.V
	}
	return ""
}

func (m *GraphStatement) GetE() string {
	if x, ok := m.GetStatement().(*GraphStatement_E); ok {
		return x.E
	}
	return ""
}

func (m *GraphStatement) GetLabel() string {
	if x, ok := m.GetStatement().(*GraphStatement_Label); ok {
		return x.Label
	}
	return ""
}

func (m *GraphStatement) GetHas() *HasStatement {
	if x, ok := m.GetStatement().(*GraphStatement_Has); ok {
		return x.Has
	}
	return nil
}

func (m *GraphStatement) GetAs() string {
	if x, ok := m.GetStatement().(*GraphStatement_As); ok {
		return x.As
	}
	return ""
}

func (m *GraphStatement) GetIn() string {
	if x, ok := m.GetStatement().(*GraphStatement_In); ok {
		return x.In
	}
	return ""
}

func (m *GraphStatement) GetOut() string {
	if x, ok := m.GetStatement().(*GraphStatement_Out); ok {
		return x.Out
	}
	return ""
}

func (m *GraphStatement) GetInEdge() string {
	if x, ok := m.GetStatement().(*GraphStatement_InEdge); ok {
		return x.InEdge
	}
	return ""
}

func (m *GraphStatement) GetOutEdge() string {
	if x, ok := m.GetStatement().(*GraphStatement_OutEdge); ok {
		return x.OutEdge
	}
	return ""
}

func (m *GraphStatement) GetInVertex() string {
	if x, ok := m.GetStatement().(*GraphStatement_InVertex); ok {
		return x.InVertex
	}
	return ""
}

func (m *GraphStatement) GetOutVertex() string {
	if x, ok := m.GetStatement().(*GraphStatement_OutVertex); ok {
		return x.OutVertex
	}
	return ""
}

func (m *GraphStatement) GetSelect() *SelectStatement {
	if x, ok := m.GetStatement().(*GraphStatement_Select); ok {
		return x.Select
	}
	return nil
}

func (m *GraphStatement) GetLimit() int64 {
	if x, ok := m.GetStatement().(*GraphStatement_Limit); ok {
		return x.Limit
	}
	return 0
}

func (m *GraphStatement) GetCount() string {
	if x, ok := m.GetStatement().(*GraphStatement_Count); ok {
		return x.Count
	}
	return ""
}

func (m *GraphStatement) GetAddV() string {
	if x, ok := m.GetStatement().(*GraphStatement_AddV); ok {
		return x.AddV
	}
	return ""
}

func (m *GraphStatement) GetProperty() *PropertyStatement {
	if x, ok := m.GetStatement().(*GraphStatement_Property); ok {
		return x.Property
	}
	return nil
}

func (m *GraphStatement) GetAddE() string {
	if x, ok := m.GetStatement().(*GraphStatement_AddE); ok {
		return x.AddE
	}
	return ""
}

func (m *GraphStatement) GetTo() string {
	if x, ok := m.GetStatement().(*GraphStatement_To); ok {
		return x.To
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GraphStatement) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GraphStatement_OneofMarshaler, _GraphStatement_OneofUnmarshaler, _GraphStatement_OneofSizer, []interface{}{
		(*GraphStatement_V)(nil),
		(*GraphStatement_E)(nil),
		(*GraphStatement_Label)(nil),
		(*GraphStatement_Has)(nil),
		(*GraphStatement_As)(nil),
		(*GraphStatement_In)(nil),
		(*GraphStatement_Out)(nil),
		(*GraphStatement_InEdge)(nil),
		(*GraphStatement_OutEdge)(nil),
		(*GraphStatement_InVertex)(nil),
		(*GraphStatement_OutVertex)(nil),
		(*GraphStatement_Select)(nil),
		(*GraphStatement_Limit)(nil),
		(*GraphStatement_Count)(nil),
		(*GraphStatement_AddV)(nil),
		(*GraphStatement_Property)(nil),
		(*GraphStatement_AddE)(nil),
		(*GraphStatement_To)(nil),
	}
}

func _GraphStatement_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GraphStatement)
	// statement
	switch x := m.Statement.(type) {
	case *GraphStatement_V:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.V)
	case *GraphStatement_E:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.E)
	case *GraphStatement_Label:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Label)
	case *GraphStatement_Has:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Has); err != nil {
			return err
		}
	case *GraphStatement_As:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.As)
	case *GraphStatement_In:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.In)
	case *GraphStatement_Out:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Out)
	case *GraphStatement_InEdge:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.InEdge)
	case *GraphStatement_OutEdge:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OutEdge)
	case *GraphStatement_InVertex:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.InVertex)
	case *GraphStatement_OutVertex:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OutVertex)
	case *GraphStatement_Select:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Select); err != nil {
			return err
		}
	case *GraphStatement_Limit:
		b.EncodeVarint(13<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Limit))
	case *GraphStatement_Count:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Count)
	case *GraphStatement_AddV:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AddV)
	case *GraphStatement_Property:
		b.EncodeVarint(16<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Property); err != nil {
			return err
		}
	case *GraphStatement_AddE:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AddE)
	case *GraphStatement_To:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.To)
	case nil:
	default:
		return fmt.Errorf("GraphStatement.Statement has unexpected type %T", x)
	}
	return nil
}

func _GraphStatement_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GraphStatement)
	switch tag {
	case 1: // statement.V
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_V{x}
		return true, err
	case 2: // statement.E
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_E{x}
		return true, err
	case 3: // statement.label
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_Label{x}
		return true, err
	case 4: // statement.has
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HasStatement)
		err := b.DecodeMessage(msg)
		m.Statement = &GraphStatement_Has{msg}
		return true, err
	case 5: // statement.as
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_As{x}
		return true, err
	case 6: // statement.in
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_In{x}
		return true, err
	case 7: // statement.out
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_Out{x}
		return true, err
	case 8: // statement.inEdge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_InEdge{x}
		return true, err
	case 9: // statement.outEdge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_OutEdge{x}
		return true, err
	case 10: // statement.inVertex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_InVertex{x}
		return true, err
	case 11: // statement.outVertex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_OutVertex{x}
		return true, err
	case 12: // statement.select
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SelectStatement)
		err := b.DecodeMessage(msg)
		m.Statement = &GraphStatement_Select{msg}
		return true, err
	case 13: // statement.limit
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Statement = &GraphStatement_Limit{int64(x)}
		return true, err
	case 14: // statement.count
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_Count{x}
		return true, err
	case 15: // statement.addV
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_AddV{x}
		return true, err
	case 16: // statement.property
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PropertyStatement)
		err := b.DecodeMessage(msg)
		m.Statement = &GraphStatement_Property{msg}
		return true, err
	case 17: // statement.addE
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_AddE{x}
		return true, err
	case 18: // statement.to
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_To{x}
		return true, err
	default:
		return false, nil
	}
}

func _GraphStatement_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GraphStatement)
	// statement
	switch x := m.Statement.(type) {
	case *GraphStatement_V:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.V)))
		n += len(x.V)
	case *GraphStatement_E:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.E)))
		n += len(x.E)
	case *GraphStatement_Label:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Label)))
		n += len(x.Label)
	case *GraphStatement_Has:
		s := proto.Size(x.Has)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GraphStatement_As:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.As)))
		n += len(x.As)
	case *GraphStatement_In:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.In)))
		n += len(x.In)
	case *GraphStatement_Out:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Out)))
		n += len(x.Out)
	case *GraphStatement_InEdge:
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.InEdge)))
		n += len(x.InEdge)
	case *GraphStatement_OutEdge:
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.OutEdge)))
		n += len(x.OutEdge)
	case *GraphStatement_InVertex:
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.InVertex)))
		n += len(x.InVertex)
	case *GraphStatement_OutVertex:
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.OutVertex)))
		n += len(x.OutVertex)
	case *GraphStatement_Select:
		s := proto.Size(x.Select)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GraphStatement_Limit:
		n += proto.SizeVarint(13<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Limit))
	case *GraphStatement_Count:
		n += proto.SizeVarint(14<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Count)))
		n += len(x.Count)
	case *GraphStatement_AddV:
		n += proto.SizeVarint(15<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AddV)))
		n += len(x.AddV)
	case *GraphStatement_Property:
		s := proto.Size(x.Property)
		n += proto.SizeVarint(16<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GraphStatement_AddE:
		n += proto.SizeVarint(17<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AddE)))
		n += len(x.AddE)
	case *GraphStatement_To:
		n += proto.SizeVarint(18<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.To)))
		n += len(x.To)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PropertyStatement struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// TODO: change from string values to 'any' data type
	// google.protobuf.Any value = 2;
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *PropertyStatement) Reset()                    { *m = PropertyStatement{} }
func (m *PropertyStatement) String() string            { return proto.CompactTextString(m) }
func (*PropertyStatement) ProtoMessage()               {}
func (*PropertyStatement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PropertyStatement) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PropertyStatement) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type HasStatement struct {
	Key    string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Within []string `protobuf:"bytes,2,rep,name=within" json:"within,omitempty"`
}

func (m *HasStatement) Reset()                    { *m = HasStatement{} }
func (m *HasStatement) String() string            { return proto.CompactTextString(m) }
func (*HasStatement) ProtoMessage()               {}
func (*HasStatement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HasStatement) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HasStatement) GetWithin() []string {
	if m != nil {
		return m.Within
	}
	return nil
}

type SelectStatement struct {
	Steps []string `protobuf:"bytes,1,rep,name=steps" json:"steps,omitempty"`
}

func (m *SelectStatement) Reset()                    { *m = SelectStatement{} }
func (m *SelectStatement) String() string            { return proto.CompactTextString(m) }
func (*SelectStatement) ProtoMessage()               {}
func (*SelectStatement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SelectStatement) GetSteps() []string {
	if m != nil {
		return m.Steps
	}
	return nil
}

type Vertex struct {
	Gid   string `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	// map<string, google.protobuf.Any> properties = 3;
	Properties map[string]string `protobuf:"bytes,3,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Vertex) Reset()                    { *m = Vertex{} }
func (m *Vertex) String() string            { return proto.CompactTextString(m) }
func (*Vertex) ProtoMessage()               {}
func (*Vertex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Vertex) GetGid() string {
	if m != nil {
		return m.Gid
	}
	return ""
}

func (m *Vertex) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Vertex) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

type Edge struct {
	Gid   string `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	In    string `protobuf:"bytes,3,opt,name=in" json:"in,omitempty"`
	Out   string `protobuf:"bytes,4,opt,name=out" json:"out,omitempty"`
	// map<string, google.protobuf.Any> properties = 5;
	Properties map[string]string `protobuf:"bytes,5,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Edge) Reset()                    { *m = Edge{} }
func (m *Edge) String() string            { return proto.CompactTextString(m) }
func (*Edge) ProtoMessage()               {}
func (*Edge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Edge) GetGid() string {
	if m != nil {
		return m.Gid
	}
	return ""
}

func (m *Edge) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Edge) GetIn() string {
	if m != nil {
		return m.In
	}
	return ""
}

func (m *Edge) GetOut() string {
	if m != nil {
		return m.Out
	}
	return ""
}

func (m *Edge) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

type QueryResult struct {
	// Types that are valid to be assigned to Result:
	//	*QueryResult_Struct
	//	*QueryResult_Vertex
	//	*QueryResult_Edge
	//	*QueryResult_IntValue
	//	*QueryResult_FloatValue
	//	*QueryResult_StrValue
	Result isQueryResult_Result `protobuf_oneof:"result"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isQueryResult_Result interface {
	isQueryResult_Result()
}

type QueryResult_Struct struct {
	Struct *google_protobuf1.Struct `protobuf:"bytes,1,opt,name=struct,oneof"`
}
type QueryResult_Vertex struct {
	Vertex *Vertex `protobuf:"bytes,2,opt,name=vertex,oneof"`
}
type QueryResult_Edge struct {
	Edge *Edge `protobuf:"bytes,3,opt,name=edge,oneof"`
}
type QueryResult_IntValue struct {
	IntValue int64 `protobuf:"varint,4,opt,name=int_value,json=intValue,oneof"`
}
type QueryResult_FloatValue struct {
	FloatValue float64 `protobuf:"fixed64,5,opt,name=float_value,json=floatValue,oneof"`
}
type QueryResult_StrValue struct {
	StrValue string `protobuf:"bytes,6,opt,name=str_value,json=strValue,oneof"`
}

func (*QueryResult_Struct) isQueryResult_Result()     {}
func (*QueryResult_Vertex) isQueryResult_Result()     {}
func (*QueryResult_Edge) isQueryResult_Result()       {}
func (*QueryResult_IntValue) isQueryResult_Result()   {}
func (*QueryResult_FloatValue) isQueryResult_Result() {}
func (*QueryResult_StrValue) isQueryResult_Result()   {}

func (m *QueryResult) GetResult() isQueryResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *QueryResult) GetStruct() *google_protobuf1.Struct {
	if x, ok := m.GetResult().(*QueryResult_Struct); ok {
		return x.Struct
	}
	return nil
}

func (m *QueryResult) GetVertex() *Vertex {
	if x, ok := m.GetResult().(*QueryResult_Vertex); ok {
		return x.Vertex
	}
	return nil
}

func (m *QueryResult) GetEdge() *Edge {
	if x, ok := m.GetResult().(*QueryResult_Edge); ok {
		return x.Edge
	}
	return nil
}

func (m *QueryResult) GetIntValue() int64 {
	if x, ok := m.GetResult().(*QueryResult_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *QueryResult) GetFloatValue() float64 {
	if x, ok := m.GetResult().(*QueryResult_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *QueryResult) GetStrValue() string {
	if x, ok := m.GetResult().(*QueryResult_StrValue); ok {
		return x.StrValue
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*QueryResult) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _QueryResult_OneofMarshaler, _QueryResult_OneofUnmarshaler, _QueryResult_OneofSizer, []interface{}{
		(*QueryResult_Struct)(nil),
		(*QueryResult_Vertex)(nil),
		(*QueryResult_Edge)(nil),
		(*QueryResult_IntValue)(nil),
		(*QueryResult_FloatValue)(nil),
		(*QueryResult_StrValue)(nil),
	}
}

func _QueryResult_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*QueryResult)
	// result
	switch x := m.Result.(type) {
	case *QueryResult_Struct:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Struct); err != nil {
			return err
		}
	case *QueryResult_Vertex:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Vertex); err != nil {
			return err
		}
	case *QueryResult_Edge:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Edge); err != nil {
			return err
		}
	case *QueryResult_IntValue:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *QueryResult_FloatValue:
		b.EncodeVarint(5<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.FloatValue))
	case *QueryResult_StrValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StrValue)
	case nil:
	default:
		return fmt.Errorf("QueryResult.Result has unexpected type %T", x)
	}
	return nil
}

func _QueryResult_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*QueryResult)
	switch tag {
	case 1: // result.struct
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Struct)
		err := b.DecodeMessage(msg)
		m.Result = &QueryResult_Struct{msg}
		return true, err
	case 2: // result.vertex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Vertex)
		err := b.DecodeMessage(msg)
		m.Result = &QueryResult_Vertex{msg}
		return true, err
	case 3: // result.edge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Edge)
		err := b.DecodeMessage(msg)
		m.Result = &QueryResult_Edge{msg}
		return true, err
	case 4: // result.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Result = &QueryResult_IntValue{int64(x)}
		return true, err
	case 5: // result.float_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Result = &QueryResult_FloatValue{math.Float64frombits(x)}
		return true, err
	case 6: // result.str_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Result = &QueryResult_StrValue{x}
		return true, err
	default:
		return false, nil
	}
}

func _QueryResult_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*QueryResult)
	// result
	switch x := m.Result.(type) {
	case *QueryResult_Struct:
		s := proto.Size(x.Struct)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *QueryResult_Vertex:
		s := proto.Size(x.Vertex)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *QueryResult_Edge:
		s := proto.Size(x.Edge)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *QueryResult_IntValue:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntValue))
	case *QueryResult_FloatValue:
		n += proto.SizeVarint(5<<3 | proto.WireFixed64)
		n += 8
	case *QueryResult_StrValue:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StrValue)))
		n += len(x.StrValue)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*GraphQuery)(nil), "ophion.GraphQuery")
	proto.RegisterType((*GraphStatement)(nil), "ophion.GraphStatement")
	proto.RegisterType((*PropertyStatement)(nil), "ophion.PropertyStatement")
	proto.RegisterType((*HasStatement)(nil), "ophion.HasStatement")
	proto.RegisterType((*SelectStatement)(nil), "ophion.SelectStatement")
	proto.RegisterType((*Vertex)(nil), "ophion.Vertex")
	proto.RegisterType((*Edge)(nil), "ophion.Edge")
	proto.RegisterType((*QueryResult)(nil), "ophion.QueryResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Query service

type QueryClient interface {
	Traversal(ctx context.Context, in *GraphQuery, opts ...grpc.CallOption) (Query_TraversalClient, error)
}

type queryClient struct {
	cc *grpc.ClientConn
}

func NewQueryClient(cc *grpc.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Traversal(ctx context.Context, in *GraphQuery, opts ...grpc.CallOption) (Query_TraversalClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Query_serviceDesc.Streams[0], c.cc, "/ophion.Query/Traversal", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryTraversalClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_TraversalClient interface {
	Recv() (*QueryResult, error)
	grpc.ClientStream
}

type queryTraversalClient struct {
	grpc.ClientStream
}

func (x *queryTraversalClient) Recv() (*QueryResult, error) {
	m := new(QueryResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Query service

type QueryServer interface {
	Traversal(*GraphQuery, Query_TraversalServer) error
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Traversal_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GraphQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).Traversal(m, &queryTraversalServer{stream})
}

type Query_TraversalServer interface {
	Send(*QueryResult) error
	grpc.ServerStream
}

type queryTraversalServer struct {
	grpc.ServerStream
}

func (x *queryTraversalServer) Send(m *QueryResult) error {
	return x.ServerStream.SendMsg(m)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ophion.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Traversal",
			Handler:       _Query_Traversal_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ophion.proto",
}

func init() { proto.RegisterFile("ophion.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 726 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xae, 0xe3, 0xc4, 0x8d, 0x27, 0xf9, 0x25, 0xe9, 0xfe, 0xa2, 0xb2, 0x44, 0x6d, 0x15, 0x7c,
	0x21, 0x42, 0x90, 0xd0, 0x72, 0xa0, 0x2a, 0x7f, 0x0e, 0x95, 0x22, 0x72, 0x04, 0x17, 0xe5, 0xc2,
	0x01, 0xb9, 0xcd, 0x36, 0x59, 0xe1, 0x7a, 0x83, 0x77, 0x1d, 0xc8, 0x15, 0x89, 0x27, 0xe0, 0x41,
	0x78, 0x12, 0x4e, 0xbc, 0x02, 0x2f, 0xc1, 0x0d, 0xed, 0xec, 0x6e, 0x48, 0x0a, 0x87, 0x5e, 0xb8,
	0x79, 0xbe, 0xef, 0x9b, 0x19, 0xcf, 0xec, 0xcc, 0x40, 0x5d, 0xcc, 0x67, 0x5c, 0x64, 0xfd, 0x79,
	0x2e, 0x94, 0x20, 0x81, 0xb1, 0x3a, 0x7b, 0x53, 0x21, 0xa6, 0x29, 0x1b, 0x24, 0x73, 0x3e, 0x48,
	0xb2, 0x4c, 0xa8, 0x44, 0x71, 0x91, 0x49, 0xa3, 0x5a, 0xb1, 0x68, 0x9d, 0x17, 0x97, 0x03, 0xa9,
	0xf2, 0xe2, 0x42, 0x19, 0x36, 0x3a, 0x01, 0x78, 0x91, 0x27, 0xf3, 0xd9, 0xab, 0x82, 0xe5, 0x4b,
	0x72, 0x1f, 0x2a, 0xef, 0xf5, 0x07, 0xf5, 0xba, 0x7e, 0xaf, 0x76, 0xb4, 0xdb, 0xb7, 0xf9, 0x50,
	0x72, 0xa6, 0x12, 0xc5, 0xae, 0x58, 0xa6, 0x62, 0x23, 0x8a, 0x3e, 0x97, 0xa1, 0xb1, 0xc9, 0x90,
	0x06, 0x78, 0x63, 0xea, 0x75, 0xbd, 0x5e, 0x38, 0xda, 0x8a, 0xbd, 0xb1, 0xb6, 0x87, 0xb4, 0xe4,
	0xec, 0x21, 0xd9, 0x85, 0x4a, 0x9a, 0x9c, 0xb3, 0x94, 0xfa, 0x16, 0x33, 0x26, 0xe9, 0x81, 0x3f,
	0x4b, 0x24, 0x2d, 0x77, 0xbd, 0x5e, 0xed, 0xa8, 0xed, 0xd2, 0x8e, 0x12, 0xb9, 0x0a, 0x3d, 0xda,
	0x8a, 0xb5, 0x84, 0xb4, 0xa0, 0x94, 0x48, 0x5a, 0xb1, 0xee, 0x25, 0x83, 0xf0, 0x8c, 0x06, 0x0e,
	0xe1, 0x19, 0x21, 0xe0, 0x8b, 0x42, 0xd1, 0x6d, 0x0b, 0x69, 0x83, 0x50, 0x08, 0x78, 0x36, 0x9c,
	0x4c, 0x19, 0xad, 0x5a, 0xd8, 0xda, 0xa4, 0x03, 0xdb, 0xa2, 0x50, 0x48, 0x85, 0x96, 0x72, 0x00,
	0xd9, 0x83, 0x2a, 0xcf, 0xc6, 0x2c, 0x57, 0xec, 0x23, 0x05, 0x4b, 0xae, 0x10, 0x72, 0x00, 0xa1,
	0x28, 0x94, 0xa5, 0x6b, 0x96, 0xfe, 0x0d, 0x91, 0x43, 0x08, 0x24, 0x4b, 0xd9, 0x85, 0xa2, 0x75,
	0x2c, 0xec, 0x96, 0x2b, 0xec, 0x0c, 0xd1, 0xf5, 0xda, 0xac, 0x10, 0x1b, 0xc4, 0xaf, 0xb8, 0xa2,
	0xff, 0x75, 0xbd, 0x9e, 0x8f, 0x0d, 0xd2, 0xa6, 0xc6, 0x2f, 0x44, 0x91, 0x29, 0xda, 0x70, 0x8d,
	0x43, 0x93, 0xb4, 0xa1, 0x9c, 0x4c, 0x26, 0x63, 0xda, 0xb4, 0x30, 0x5a, 0xe4, 0x31, 0x54, 0xe7,
	0xb9, 0x98, 0xb3, 0x5c, 0x2d, 0x69, 0x0b, 0x53, 0xdf, 0x76, 0xa9, 0x5f, 0x5a, 0x7c, 0x3d, 0xf9,
	0x4a, 0x6c, 0xc3, 0x0d, 0xe9, 0xce, 0x5a, 0xb8, 0xa1, 0xee, 0xb0, 0x12, 0x94, 0xb8, 0x0e, 0x2b,
	0x71, 0x5a, 0x83, 0x50, 0xba, 0x00, 0xd1, 0x13, 0xd8, 0xf9, 0x23, 0x2a, 0x69, 0x81, 0xff, 0x8e,
	0x2d, 0xcd, 0x2c, 0xc4, 0xfa, 0x93, 0xb4, 0xa1, 0xb2, 0x48, 0xd2, 0x82, 0x99, 0x79, 0x88, 0x8d,
	0x11, 0x1d, 0x43, 0x7d, 0xfd, 0x99, 0xff, 0xe2, 0xb7, 0x0b, 0xc1, 0x07, 0xae, 0x66, 0x3c, 0xa3,
	0xa5, 0xae, 0xdf, 0x0b, 0x63, 0x6b, 0x45, 0x77, 0xa1, 0x79, 0xad, 0x8f, 0x3a, 0x85, 0x54, 0x6c,
	0x2e, 0x71, 0x7e, 0xc3, 0xd8, 0x18, 0xd1, 0x57, 0x0f, 0x02, 0xfb, 0x22, 0x2d, 0xf0, 0xa7, 0x7c,
	0xe2, 0xa2, 0x4f, 0xf9, 0x44, 0xbb, 0x98, 0x89, 0xb4, 0x7f, 0x65, 0xe6, 0xf1, 0x39, 0x80, 0xed,
	0x09, 0x67, 0x92, 0xfa, 0xb8, 0x0d, 0x07, 0xae, 0x85, 0x26, 0x96, 0xeb, 0x24, 0x67, 0x72, 0x98,
	0xa9, 0x7c, 0x19, 0xaf, 0x79, 0x74, 0x9e, 0x41, 0xf3, 0x1a, 0x7d, 0xd3, 0x86, 0x9c, 0x94, 0x8e,
	0xbd, 0xe8, 0x9b, 0x07, 0x65, 0x9c, 0xbf, 0x9b, 0xfe, 0x6f, 0x03, 0x77, 0x00, 0x97, 0x0a, 0x37,
	0xa0, 0x65, 0x36, 0xa0, 0x6c, 0xfc, 0xf4, 0xfc, 0x3f, 0xdd, 0xa8, 0xa8, 0x82, 0x15, 0xed, 0xb9,
	0x8a, 0x74, 0xae, 0x7f, 0x59, 0xcf, 0x4f, 0x0f, 0x6a, 0x78, 0x61, 0x62, 0x26, 0x8b, 0x54, 0xe1,
	0x62, 0xe0, 0x15, 0x42, 0x77, 0xbd, 0x18, 0xe6, 0x48, 0xf5, 0xdd, 0x91, 0xea, 0x9f, 0x21, 0x8d,
	0x8b, 0x81, 0x5f, 0xa4, 0x07, 0xc1, 0xc2, 0x2c, 0x5a, 0x09, 0x5d, 0x1a, 0x9b, 0xaf, 0xa1, 0x95,
	0x86, 0x27, 0x11, 0x94, 0x99, 0x5e, 0x66, 0x1f, 0x75, 0xf5, 0xf5, 0x1a, 0xf5, 0x44, 0x6b, 0x8e,
	0xec, 0x43, 0xc8, 0x33, 0xf5, 0xd6, 0xfc, 0x6e, 0xd9, 0xae, 0x5a, 0x95, 0x67, 0x6a, 0xac, 0x11,
	0x72, 0x07, 0x6a, 0x97, 0xa9, 0x48, 0x9c, 0x40, 0x5f, 0x1b, 0x6f, 0xb4, 0x15, 0x03, 0x82, 0x46,
	0xb2, 0xaf, 0x37, 0x20, 0xb7, 0x02, 0x77, 0x7c, 0xaa, 0x52, 0xe5, 0x48, 0x9f, 0x56, 0x21, 0xc8,
	0xb1, 0xd6, 0xa3, 0x37, 0x50, 0x31, 0xc7, 0x35, 0x86, 0xf0, 0x75, 0x9e, 0x2c, 0x58, 0x2e, 0x93,
	0x94, 0x90, 0x8d, 0xd3, 0x8a, 0x82, 0xce, 0xff, 0x0e, 0x5b, 0x6b, 0x55, 0xd4, 0xf9, 0xf4, 0xfd,
	0xc7, 0x97, 0x52, 0x3b, 0x6a, 0x0e, 0x16, 0x87, 0x83, 0xa9, 0x16, 0x3f, 0xc0, 0xeb, 0x7b, 0xe2,
	0xdd, 0x7b, 0xe8, 0x9d, 0x07, 0xd8, 0xb0, 0x47, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x63,
	0x29, 0x95, 0x19, 0x06, 0x00, 0x00,
}
