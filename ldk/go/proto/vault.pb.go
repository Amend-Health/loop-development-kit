// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: vault.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VaultExistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Key     string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *VaultExistsRequest) Reset() {
	*x = VaultExistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultExistsRequest) ProtoMessage() {}

func (x *VaultExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultExistsRequest.ProtoReflect.Descriptor instead.
func (*VaultExistsRequest) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{0}
}

func (x *VaultExistsRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *VaultExistsRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type VaultExistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *VaultExistsResponse) Reset() {
	*x = VaultExistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultExistsResponse) ProtoMessage() {}

func (x *VaultExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultExistsResponse.ProtoReflect.Descriptor instead.
func (*VaultExistsResponse) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{1}
}

func (x *VaultExistsResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type VaultReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Key     string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *VaultReadRequest) Reset() {
	*x = VaultReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultReadRequest) ProtoMessage() {}

func (x *VaultReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultReadRequest.ProtoReflect.Descriptor instead.
func (*VaultReadRequest) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{2}
}

func (x *VaultReadRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *VaultReadRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type VaultReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *VaultReadResponse) Reset() {
	*x = VaultReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultReadResponse) ProtoMessage() {}

func (x *VaultReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultReadResponse.ProtoReflect.Descriptor instead.
func (*VaultReadResponse) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{3}
}

func (x *VaultReadResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type VaultDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Key     string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *VaultDeleteRequest) Reset() {
	*x = VaultDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultDeleteRequest) ProtoMessage() {}

func (x *VaultDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultDeleteRequest.ProtoReflect.Descriptor instead.
func (*VaultDeleteRequest) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{4}
}

func (x *VaultDeleteRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *VaultDeleteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type VaultWriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Key     string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value   string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *VaultWriteRequest) Reset() {
	*x = VaultWriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultWriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultWriteRequest) ProtoMessage() {}

func (x *VaultWriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultWriteRequest.ProtoReflect.Descriptor instead.
func (*VaultWriteRequest) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{5}
}

func (x *VaultWriteRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *VaultWriteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *VaultWriteRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_vault_proto protoreflect.FileDescriptor

var file_vault_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x50, 0x0a, 0x12, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x2d, 0x0a, 0x13, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x22, 0x4e, 0x0a, 0x10, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x29, 0x0a, 0x11, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x50, 0x0a, 0x12,
	0x56, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x65,
	0x0a, 0x11, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x8f, 0x02, 0x0a, 0x05, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x12,
	0x40, 0x0a, 0x0b, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x44, 0x0a, 0x0b, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x56, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x61, 0x64, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x56, 0x61, 0x75, 0x6c, 0x74,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61,
	0x75, 0x6c, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vault_proto_rawDescOnce sync.Once
	file_vault_proto_rawDescData = file_vault_proto_rawDesc
)

func file_vault_proto_rawDescGZIP() []byte {
	file_vault_proto_rawDescOnce.Do(func() {
		file_vault_proto_rawDescData = protoimpl.X.CompressGZIP(file_vault_proto_rawDescData)
	})
	return file_vault_proto_rawDescData
}

var file_vault_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vault_proto_goTypes = []interface{}{
	(*VaultExistsRequest)(nil),  // 0: proto.VaultExistsRequest
	(*VaultExistsResponse)(nil), // 1: proto.VaultExistsResponse
	(*VaultReadRequest)(nil),    // 2: proto.VaultReadRequest
	(*VaultReadResponse)(nil),   // 3: proto.VaultReadResponse
	(*VaultDeleteRequest)(nil),  // 4: proto.VaultDeleteRequest
	(*VaultWriteRequest)(nil),   // 5: proto.VaultWriteRequest
	(*Session)(nil),             // 6: proto.Session
	(*emptypb.Empty)(nil),       // 7: google.protobuf.Empty
}
var file_vault_proto_depIdxs = []int32{
	6, // 0: proto.VaultExistsRequest.session:type_name -> proto.Session
	6, // 1: proto.VaultReadRequest.session:type_name -> proto.Session
	6, // 2: proto.VaultDeleteRequest.session:type_name -> proto.Session
	6, // 3: proto.VaultWriteRequest.session:type_name -> proto.Session
	4, // 4: proto.Vault.VaultDelete:input_type -> proto.VaultDeleteRequest
	0, // 5: proto.Vault.VaultExists:input_type -> proto.VaultExistsRequest
	2, // 6: proto.Vault.VaultRead:input_type -> proto.VaultReadRequest
	5, // 7: proto.Vault.VaultWrite:input_type -> proto.VaultWriteRequest
	7, // 8: proto.Vault.VaultDelete:output_type -> google.protobuf.Empty
	1, // 9: proto.Vault.VaultExists:output_type -> proto.VaultExistsResponse
	3, // 10: proto.Vault.VaultRead:output_type -> proto.VaultReadResponse
	7, // 11: proto.Vault.VaultWrite:output_type -> google.protobuf.Empty
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_vault_proto_init() }
func file_vault_proto_init() {
	if File_vault_proto != nil {
		return
	}
	file_session_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vault_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultExistsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vault_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultExistsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vault_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultReadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vault_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultReadResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vault_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vault_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultWriteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vault_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vault_proto_goTypes,
		DependencyIndexes: file_vault_proto_depIdxs,
		MessageInfos:      file_vault_proto_msgTypes,
	}.Build()
	File_vault_proto = out.File
	file_vault_proto_rawDesc = nil
	file_vault_proto_goTypes = nil
	file_vault_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VaultClient is the client API for Vault service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VaultClient interface {
	// Delete the value of a specific key in the vault
	VaultDelete(ctx context.Context, in *VaultDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Check in the vault to determine if a key has a value
	VaultExists(ctx context.Context, in *VaultExistsRequest, opts ...grpc.CallOption) (*VaultExistsResponse, error)
	// Read the value of a specific vault key
	VaultRead(ctx context.Context, in *VaultReadRequest, opts ...grpc.CallOption) (*VaultReadResponse, error)
	// Write the value of a key to the vault
	VaultWrite(ctx context.Context, in *VaultWriteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type vaultClient struct {
	cc grpc.ClientConnInterface
}

func NewVaultClient(cc grpc.ClientConnInterface) VaultClient {
	return &vaultClient{cc}
}

func (c *vaultClient) VaultDelete(ctx context.Context, in *VaultDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.Vault/VaultDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vaultClient) VaultExists(ctx context.Context, in *VaultExistsRequest, opts ...grpc.CallOption) (*VaultExistsResponse, error) {
	out := new(VaultExistsResponse)
	err := c.cc.Invoke(ctx, "/proto.Vault/VaultExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vaultClient) VaultRead(ctx context.Context, in *VaultReadRequest, opts ...grpc.CallOption) (*VaultReadResponse, error) {
	out := new(VaultReadResponse)
	err := c.cc.Invoke(ctx, "/proto.Vault/VaultRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vaultClient) VaultWrite(ctx context.Context, in *VaultWriteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.Vault/VaultWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VaultServer is the server API for Vault service.
type VaultServer interface {
	// Delete the value of a specific key in the vault
	VaultDelete(context.Context, *VaultDeleteRequest) (*emptypb.Empty, error)
	// Check in the vault to determine if a key has a value
	VaultExists(context.Context, *VaultExistsRequest) (*VaultExistsResponse, error)
	// Read the value of a specific vault key
	VaultRead(context.Context, *VaultReadRequest) (*VaultReadResponse, error)
	// Write the value of a key to the vault
	VaultWrite(context.Context, *VaultWriteRequest) (*emptypb.Empty, error)
}

// UnimplementedVaultServer can be embedded to have forward compatible implementations.
type UnimplementedVaultServer struct {
}

func (*UnimplementedVaultServer) VaultDelete(context.Context, *VaultDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VaultDelete not implemented")
}
func (*UnimplementedVaultServer) VaultExists(context.Context, *VaultExistsRequest) (*VaultExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VaultExists not implemented")
}
func (*UnimplementedVaultServer) VaultRead(context.Context, *VaultReadRequest) (*VaultReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VaultRead not implemented")
}
func (*UnimplementedVaultServer) VaultWrite(context.Context, *VaultWriteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VaultWrite not implemented")
}

func RegisterVaultServer(s *grpc.Server, srv VaultServer) {
	s.RegisterService(&_Vault_serviceDesc, srv)
}

func _Vault_VaultDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VaultDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultServer).VaultDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Vault/VaultDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultServer).VaultDelete(ctx, req.(*VaultDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vault_VaultExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VaultExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultServer).VaultExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Vault/VaultExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultServer).VaultExists(ctx, req.(*VaultExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vault_VaultRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VaultReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultServer).VaultRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Vault/VaultRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultServer).VaultRead(ctx, req.(*VaultReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vault_VaultWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VaultWriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultServer).VaultWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Vault/VaultWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultServer).VaultWrite(ctx, req.(*VaultWriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Vault_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Vault",
	HandlerType: (*VaultServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VaultDelete",
			Handler:    _Vault_VaultDelete_Handler,
		},
		{
			MethodName: "VaultExists",
			Handler:    _Vault_VaultExists_Handler,
		},
		{
			MethodName: "VaultRead",
			Handler:    _Vault_VaultRead_Handler,
		},
		{
			MethodName: "VaultWrite",
			Handler:    _Vault_VaultWrite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vault.proto",
}
