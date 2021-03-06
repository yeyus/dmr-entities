// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entities.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	entities.proto

It has these top-level messages:
	SearchEntitiesRequest
	GetEntityRequest
	GetEntitiesBySystemIDRequest
	GetEntitiesByCallsignRequest
	EntitiesResponse
	Pagination
	Entity
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Namespace int32

const (
	Namespace_UNKNOWN_NAMESPACE Namespace = 0
	Namespace_RADIOID           Namespace = 1
	Namespace_DMRMARC           Namespace = 2
	Namespace_BRANDMEISTER      Namespace = 3
	Namespace_TGIF              Namespace = 4
)

var Namespace_name = map[int32]string{
	0: "UNKNOWN_NAMESPACE",
	1: "RADIOID",
	2: "DMRMARC",
	3: "BRANDMEISTER",
	4: "TGIF",
}
var Namespace_value = map[string]int32{
	"UNKNOWN_NAMESPACE": 0,
	"RADIOID":           1,
	"DMRMARC":           2,
	"BRANDMEISTER":      3,
	"TGIF":              4,
}

func (x Namespace) String() string {
	return proto.EnumName(Namespace_name, int32(x))
}
func (Namespace) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SystemType int32

const (
	SystemType_UNKNOWN_SYSTEM_TYPE SystemType = 0
	SystemType_DMR                 SystemType = 1
	SystemType_NXDN                SystemType = 2
)

var SystemType_name = map[int32]string{
	0: "UNKNOWN_SYSTEM_TYPE",
	1: "DMR",
	2: "NXDN",
}
var SystemType_value = map[string]int32{
	"UNKNOWN_SYSTEM_TYPE": 0,
	"DMR":  1,
	"NXDN": 2,
}

func (x SystemType) String() string {
	return proto.EnumName(SystemType_name, int32(x))
}
func (SystemType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type EntityType int32

const (
	EntityType_UNKNOWN_ENTITY_TYPE EntityType = 0
	EntityType_HAM                 EntityType = 1
	EntityType_TALKGROUP           EntityType = 2
)

var EntityType_name = map[int32]string{
	0: "UNKNOWN_ENTITY_TYPE",
	1: "HAM",
	2: "TALKGROUP",
}
var EntityType_value = map[string]int32{
	"UNKNOWN_ENTITY_TYPE": 0,
	"HAM":       1,
	"TALKGROUP": 2,
}

func (x EntityType) String() string {
	return proto.EnumName(EntityType_name, int32(x))
}
func (EntityType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SearchEntitiesRequest struct {
	Query      string     `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	CountryIso string     `protobuf:"bytes,2,opt,name=country_iso,json=countryIso" json:"country_iso,omitempty"`
	Type       EntityType `protobuf:"varint,3,opt,name=type,enum=api.EntityType" json:"type,omitempty"`
	System     SystemType `protobuf:"varint,4,opt,name=system,enum=api.SystemType" json:"system,omitempty"`
	Namespace  Namespace  `protobuf:"varint,7,opt,name=namespace,enum=api.Namespace" json:"namespace,omitempty"`
	PageSize   uint32     `protobuf:"varint,5,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	PageStart  uint32     `protobuf:"varint,6,opt,name=page_start,json=pageStart" json:"page_start,omitempty"`
}

func (m *SearchEntitiesRequest) Reset()                    { *m = SearchEntitiesRequest{} }
func (m *SearchEntitiesRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchEntitiesRequest) ProtoMessage()               {}
func (*SearchEntitiesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SearchEntitiesRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchEntitiesRequest) GetCountryIso() string {
	if m != nil {
		return m.CountryIso
	}
	return ""
}

func (m *SearchEntitiesRequest) GetType() EntityType {
	if m != nil {
		return m.Type
	}
	return EntityType_UNKNOWN_ENTITY_TYPE
}

func (m *SearchEntitiesRequest) GetSystem() SystemType {
	if m != nil {
		return m.System
	}
	return SystemType_UNKNOWN_SYSTEM_TYPE
}

func (m *SearchEntitiesRequest) GetNamespace() Namespace {
	if m != nil {
		return m.Namespace
	}
	return Namespace_UNKNOWN_NAMESPACE
}

func (m *SearchEntitiesRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchEntitiesRequest) GetPageStart() uint32 {
	if m != nil {
		return m.PageStart
	}
	return 0
}

type GetEntityRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetEntityRequest) Reset()                    { *m = GetEntityRequest{} }
func (m *GetEntityRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntityRequest) ProtoMessage()               {}
func (*GetEntityRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetEntityRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetEntitiesBySystemIDRequest struct {
	SystemId []uint64 `protobuf:"varint,1,rep,packed,name=system_id,json=systemId" json:"system_id,omitempty"`
}

func (m *GetEntitiesBySystemIDRequest) Reset()                    { *m = GetEntitiesBySystemIDRequest{} }
func (m *GetEntitiesBySystemIDRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntitiesBySystemIDRequest) ProtoMessage()               {}
func (*GetEntitiesBySystemIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetEntitiesBySystemIDRequest) GetSystemId() []uint64 {
	if m != nil {
		return m.SystemId
	}
	return nil
}

type GetEntitiesByCallsignRequest struct {
	Callsign []string `protobuf:"bytes,1,rep,name=callsign" json:"callsign,omitempty"`
}

func (m *GetEntitiesByCallsignRequest) Reset()                    { *m = GetEntitiesByCallsignRequest{} }
func (m *GetEntitiesByCallsignRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntitiesByCallsignRequest) ProtoMessage()               {}
func (*GetEntitiesByCallsignRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetEntitiesByCallsignRequest) GetCallsign() []string {
	if m != nil {
		return m.Callsign
	}
	return nil
}

type EntitiesResponse struct {
	Entities   []*Entity   `protobuf:"bytes,1,rep,name=entities" json:"entities,omitempty"`
	Pagination *Pagination `protobuf:"bytes,2,opt,name=pagination" json:"pagination,omitempty"`
}

func (m *EntitiesResponse) Reset()                    { *m = EntitiesResponse{} }
func (m *EntitiesResponse) String() string            { return proto.CompactTextString(m) }
func (*EntitiesResponse) ProtoMessage()               {}
func (*EntitiesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EntitiesResponse) GetEntities() []*Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *EntitiesResponse) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type Pagination struct {
	Total     uint32 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
	PageStart uint32 `protobuf:"varint,2,opt,name=page_start,json=pageStart" json:"page_start,omitempty"`
	PageSize  uint32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *Pagination) Reset()                    { *m = Pagination{} }
func (m *Pagination) String() string            { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()               {}
func (*Pagination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Pagination) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Pagination) GetPageStart() uint32 {
	if m != nil {
		return m.PageStart
	}
	return 0
}

func (m *Pagination) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type Entity struct {
	Id         uint64     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SystemId   uint64     `protobuf:"varint,2,opt,name=system_id,json=systemId" json:"system_id,omitempty"`
	System     SystemType `protobuf:"varint,3,opt,name=system,enum=api.SystemType" json:"system,omitempty"`
	Type       EntityType `protobuf:"varint,4,opt,name=type,enum=api.EntityType" json:"type,omitempty"`
	Namespace  Namespace  `protobuf:"varint,5,opt,name=namespace,enum=api.Namespace" json:"namespace,omitempty"`
	Callsign   string     `protobuf:"bytes,6,opt,name=callsign" json:"callsign,omitempty"`
	Name       string     `protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
	Surname    string     `protobuf:"bytes,8,opt,name=surname" json:"surname,omitempty"`
	Country    string     `protobuf:"bytes,9,opt,name=country" json:"country,omitempty"`
	CountryIso string     `protobuf:"bytes,10,opt,name=country_iso,json=countryIso" json:"country_iso,omitempty"`
	State      string     `protobuf:"bytes,11,opt,name=state" json:"state,omitempty"`
	City       string     `protobuf:"bytes,12,opt,name=city" json:"city,omitempty"`
}

func (m *Entity) Reset()                    { *m = Entity{} }
func (m *Entity) String() string            { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()               {}
func (*Entity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Entity) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Entity) GetSystemId() uint64 {
	if m != nil {
		return m.SystemId
	}
	return 0
}

func (m *Entity) GetSystem() SystemType {
	if m != nil {
		return m.System
	}
	return SystemType_UNKNOWN_SYSTEM_TYPE
}

func (m *Entity) GetType() EntityType {
	if m != nil {
		return m.Type
	}
	return EntityType_UNKNOWN_ENTITY_TYPE
}

func (m *Entity) GetNamespace() Namespace {
	if m != nil {
		return m.Namespace
	}
	return Namespace_UNKNOWN_NAMESPACE
}

func (m *Entity) GetCallsign() string {
	if m != nil {
		return m.Callsign
	}
	return ""
}

func (m *Entity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entity) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *Entity) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Entity) GetCountryIso() string {
	if m != nil {
		return m.CountryIso
	}
	return ""
}

func (m *Entity) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Entity) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchEntitiesRequest)(nil), "api.SearchEntitiesRequest")
	proto.RegisterType((*GetEntityRequest)(nil), "api.GetEntityRequest")
	proto.RegisterType((*GetEntitiesBySystemIDRequest)(nil), "api.GetEntitiesBySystemIDRequest")
	proto.RegisterType((*GetEntitiesByCallsignRequest)(nil), "api.GetEntitiesByCallsignRequest")
	proto.RegisterType((*EntitiesResponse)(nil), "api.EntitiesResponse")
	proto.RegisterType((*Pagination)(nil), "api.Pagination")
	proto.RegisterType((*Entity)(nil), "api.Entity")
	proto.RegisterEnum("api.Namespace", Namespace_name, Namespace_value)
	proto.RegisterEnum("api.SystemType", SystemType_name, SystemType_value)
	proto.RegisterEnum("api.EntityType", EntityType_name, EntityType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Entities service

type EntitiesClient interface {
	SearchEntity(ctx context.Context, in *SearchEntitiesRequest, opts ...grpc.CallOption) (*EntitiesResponse, error)
	GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*Entity, error)
	GetEntitiesBySystemID(ctx context.Context, in *GetEntitiesBySystemIDRequest, opts ...grpc.CallOption) (*EntitiesResponse, error)
	GetEntitiesByCallsign(ctx context.Context, in *GetEntitiesByCallsignRequest, opts ...grpc.CallOption) (*EntitiesResponse, error)
}

type entitiesClient struct {
	cc *grpc.ClientConn
}

func NewEntitiesClient(cc *grpc.ClientConn) EntitiesClient {
	return &entitiesClient{cc}
}

func (c *entitiesClient) SearchEntity(ctx context.Context, in *SearchEntitiesRequest, opts ...grpc.CallOption) (*EntitiesResponse, error) {
	out := new(EntitiesResponse)
	err := grpc.Invoke(ctx, "/api.Entities/SearchEntity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entitiesClient) GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := grpc.Invoke(ctx, "/api.Entities/GetEntity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entitiesClient) GetEntitiesBySystemID(ctx context.Context, in *GetEntitiesBySystemIDRequest, opts ...grpc.CallOption) (*EntitiesResponse, error) {
	out := new(EntitiesResponse)
	err := grpc.Invoke(ctx, "/api.Entities/GetEntitiesBySystemID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entitiesClient) GetEntitiesByCallsign(ctx context.Context, in *GetEntitiesByCallsignRequest, opts ...grpc.CallOption) (*EntitiesResponse, error) {
	out := new(EntitiesResponse)
	err := grpc.Invoke(ctx, "/api.Entities/GetEntitiesByCallsign", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Entities service

type EntitiesServer interface {
	SearchEntity(context.Context, *SearchEntitiesRequest) (*EntitiesResponse, error)
	GetEntity(context.Context, *GetEntityRequest) (*Entity, error)
	GetEntitiesBySystemID(context.Context, *GetEntitiesBySystemIDRequest) (*EntitiesResponse, error)
	GetEntitiesByCallsign(context.Context, *GetEntitiesByCallsignRequest) (*EntitiesResponse, error)
}

func RegisterEntitiesServer(s *grpc.Server, srv EntitiesServer) {
	s.RegisterService(&_Entities_serviceDesc, srv)
}

func _Entities_SearchEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntitiesServer).SearchEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Entities/SearchEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntitiesServer).SearchEntity(ctx, req.(*SearchEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Entities_GetEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntitiesServer).GetEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Entities/GetEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntitiesServer).GetEntity(ctx, req.(*GetEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Entities_GetEntitiesBySystemID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntitiesBySystemIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntitiesServer).GetEntitiesBySystemID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Entities/GetEntitiesBySystemID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntitiesServer).GetEntitiesBySystemID(ctx, req.(*GetEntitiesBySystemIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Entities_GetEntitiesByCallsign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntitiesByCallsignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntitiesServer).GetEntitiesByCallsign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Entities/GetEntitiesByCallsign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntitiesServer).GetEntitiesByCallsign(ctx, req.(*GetEntitiesByCallsignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Entities_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Entities",
	HandlerType: (*EntitiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchEntity",
			Handler:    _Entities_SearchEntity_Handler,
		},
		{
			MethodName: "GetEntity",
			Handler:    _Entities_GetEntity_Handler,
		},
		{
			MethodName: "GetEntitiesBySystemID",
			Handler:    _Entities_GetEntitiesBySystemID_Handler,
		},
		{
			MethodName: "GetEntitiesByCallsign",
			Handler:    _Entities_GetEntitiesByCallsign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entities.proto",
}

func init() { proto.RegisterFile("entities.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x6e, 0x1c, 0xe7, 0xe0, 0x49, 0x9b, 0xdf, 0xff, 0xd2, 0x08, 0x2b, 0x05, 0x51, 0xcc, 0x05,
	0x51, 0x85, 0x8a, 0x28, 0x37, 0x08, 0xc4, 0x85, 0x1b, 0x9b, 0x62, 0x15, 0xbb, 0xd1, 0xda, 0x15,
	0xed, 0x0d, 0x91, 0x49, 0x57, 0x65, 0xa5, 0xd4, 0x76, 0xb3, 0x9b, 0x0b, 0xf7, 0x01, 0x78, 0x1a,
	0x1e, 0x8a, 0x47, 0x41, 0x5e, 0x1f, 0xe2, 0x98, 0xb4, 0xdc, 0x79, 0xbe, 0xf9, 0x66, 0x3c, 0xfb,
	0xcd, 0x01, 0xfa, 0x24, 0xe4, 0x94, 0x53, 0xc2, 0x0e, 0xe3, 0x45, 0xc4, 0x23, 0xd4, 0x0c, 0x62,
	0xaa, 0xff, 0x94, 0x60, 0xe0, 0x91, 0x60, 0x31, 0xfb, 0x61, 0xe5, 0x5e, 0x4c, 0x6e, 0x97, 0x84,
	0x71, 0xb4, 0x0b, 0xad, 0xdb, 0x25, 0x59, 0x24, 0x5a, 0x63, 0xbf, 0x31, 0x52, 0x70, 0x66, 0xa0,
	0x67, 0xd0, 0x9b, 0x45, 0xcb, 0x90, 0x2f, 0x92, 0x29, 0x65, 0x91, 0x26, 0x09, 0x1f, 0xe4, 0x90,
	0xcd, 0x22, 0xf4, 0x02, 0x64, 0x9e, 0xc4, 0x44, 0x6b, 0xee, 0x37, 0x46, 0xfd, 0xa3, 0xff, 0x0e,
	0x83, 0x98, 0x1e, 0x8a, 0xd4, 0x89, 0x9f, 0xc4, 0x04, 0x0b, 0x27, 0x7a, 0x09, 0x6d, 0x96, 0x30,
	0x4e, 0x6e, 0x34, 0xb9, 0x42, 0xf3, 0x04, 0x24, 0x68, 0xb9, 0x1b, 0xbd, 0x02, 0x25, 0x0c, 0x6e,
	0x08, 0x8b, 0x83, 0x19, 0xd1, 0x3a, 0x82, 0xdb, 0x17, 0x5c, 0xb7, 0x40, 0xf1, 0x8a, 0x80, 0xf6,
	0x40, 0x89, 0x83, 0x6b, 0x32, 0x65, 0xf4, 0x8e, 0x68, 0xad, 0xfd, 0xc6, 0x68, 0x07, 0x77, 0x53,
	0xc0, 0xa3, 0x77, 0x04, 0x3d, 0x05, 0xc8, 0x9c, 0x3c, 0x58, 0x70, 0xad, 0x2d, 0xbc, 0x82, 0xee,
	0xa5, 0x80, 0xae, 0x83, 0x7a, 0x42, 0x78, 0x56, 0x69, 0x21, 0x41, 0x1f, 0x24, 0x7a, 0x25, 0xde,
	0x2f, 0x63, 0x89, 0x5e, 0xe9, 0x1f, 0xe0, 0x49, 0xc1, 0xa1, 0x84, 0x1d, 0x27, 0x59, 0xc1, 0xb6,
	0x59, 0xf0, 0xf7, 0x40, 0xc9, 0xea, 0x9e, 0x8a, 0xb0, 0xe6, 0x48, 0xc6, 0xdd, 0x0c, 0xb0, 0xaf,
	0xf4, 0xf7, 0xb5, 0xe0, 0x71, 0x30, 0x9f, 0x33, 0x7a, 0x1d, 0x16, 0xc1, 0x43, 0xe8, 0xce, 0x72,
	0x48, 0xc4, 0x2a, 0xb8, 0xb4, 0xf5, 0x39, 0xa8, 0xab, 0xf6, 0xb0, 0x38, 0x0a, 0x59, 0xaa, 0x61,
	0xb7, 0x68, 0xa8, 0xe0, 0xf7, 0x8e, 0x7a, 0x15, 0xb1, 0x71, 0xe9, 0x44, 0xaf, 0xc5, 0xc3, 0x69,
	0x18, 0x70, 0x1a, 0x85, 0xa2, 0x63, 0xbd, 0x5c, 0xf0, 0x49, 0x09, 0xe3, 0x0a, 0x45, 0xff, 0x06,
	0xb0, 0xf2, 0xa4, 0x73, 0xc0, 0x23, 0x1e, 0xcc, 0x85, 0x0e, 0x3b, 0x38, 0x33, 0x6a, 0x6a, 0x4a,
	0x35, 0x35, 0xd7, 0x3b, 0xd1, 0x5c, 0xef, 0x84, 0xfe, 0x5b, 0x82, 0x76, 0x56, 0x65, 0x5d, 0xe1,
	0x75, 0x05, 0x25, 0x01, 0x97, 0x0a, 0x56, 0xa6, 0xa6, 0xf9, 0xf0, 0xd4, 0x14, 0x33, 0x28, 0x3f,
	0x34, 0x83, 0x6b, 0xa3, 0xd5, 0xfa, 0xd7, 0x68, 0x55, 0xbb, 0xd3, 0x16, 0x43, 0x5f, 0xda, 0x08,
	0x81, 0x9c, 0x12, 0xc5, 0x7c, 0x2a, 0x58, 0x7c, 0x23, 0x0d, 0x3a, 0x6c, 0xb9, 0x10, 0x70, 0x57,
	0xc0, 0x85, 0x99, 0x7a, 0xf2, 0x75, 0xd1, 0x94, 0xcc, 0x93, 0x9b, 0xf5, 0xdd, 0x82, 0xbf, 0x76,
	0x6b, 0x17, 0x5a, 0x8c, 0x07, 0x9c, 0x68, 0xbd, 0x6c, 0x25, 0x85, 0x91, 0xfe, 0x7e, 0x46, 0x79,
	0xa2, 0x6d, 0x67, 0xbf, 0x4f, 0xbf, 0x0f, 0x2e, 0x40, 0x29, 0x9f, 0x81, 0x06, 0xf0, 0xff, 0xb9,
	0x7b, 0xea, 0x9e, 0x7d, 0x75, 0xa7, 0xae, 0xe1, 0x58, 0xde, 0xc4, 0x18, 0x5b, 0xea, 0x16, 0xea,
	0x41, 0x07, 0x1b, 0xa6, 0x7d, 0x66, 0x9b, 0x6a, 0x23, 0x35, 0x4c, 0x07, 0x3b, 0x06, 0x1e, 0xab,
	0x12, 0x52, 0x61, 0xfb, 0x18, 0x1b, 0xae, 0xe9, 0x58, 0xb6, 0xe7, 0x5b, 0x58, 0x6d, 0xa2, 0x2e,
	0xc8, 0xfe, 0x89, 0xfd, 0x49, 0x95, 0x0f, 0xde, 0x01, 0xac, 0x14, 0x47, 0x8f, 0xe1, 0x51, 0x91,
	0xda, 0xbb, 0xf4, 0x7c, 0xcb, 0x99, 0xfa, 0x97, 0x93, 0x34, 0x79, 0x07, 0x9a, 0xa6, 0x83, 0xd5,
	0x46, 0x1a, 0xe9, 0x5e, 0x98, 0xae, 0x2a, 0x1d, 0x7c, 0x04, 0x58, 0x35, 0xa1, 0x1a, 0x69, 0xb9,
	0xbe, 0xed, 0x5f, 0x56, 0x22, 0x3f, 0x1b, 0x8e, 0xda, 0x40, 0x3b, 0xa0, 0xf8, 0xc6, 0x97, 0xd3,
	0x13, 0x7c, 0x76, 0x3e, 0x51, 0xa5, 0xa3, 0x5f, 0x12, 0x74, 0x8b, 0x25, 0x40, 0x63, 0xd8, 0xae,
	0x5c, 0xad, 0x04, 0x0d, 0xb3, 0x51, 0xd8, 0x74, 0xc8, 0x86, 0x83, 0x55, 0xff, 0x2b, 0xfb, 0xa3,
	0x6f, 0xa1, 0x37, 0xa0, 0x94, 0x2b, 0x8f, 0x32, 0x56, 0xfd, 0x04, 0x0c, 0xab, 0x3b, 0xa5, 0x6f,
	0x21, 0x0f, 0x06, 0x1b, 0x2f, 0x00, 0x7a, 0xbe, 0x16, 0xbe, 0xe9, 0x3a, 0xdc, 0x5f, 0x47, 0x3d,
	0x69, 0x71, 0x19, 0x36, 0x25, 0xad, 0x5d, 0x8d, 0x7b, 0x93, 0x7e, 0x6f, 0x8b, 0x23, 0xff, 0xf6,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x84, 0xdd, 0x0a, 0xf6, 0x05, 0x00, 0x00,
}
