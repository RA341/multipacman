// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: lobby/v1/lobby.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListLobbiesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListLobbiesRequest) Reset() {
	*x = ListLobbiesRequest{}
	mi := &file_lobby_v1_lobby_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLobbiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLobbiesRequest) ProtoMessage() {}

func (x *ListLobbiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_v1_lobby_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLobbiesRequest.ProtoReflect.Descriptor instead.
func (*ListLobbiesRequest) Descriptor() ([]byte, []int) {
	return file_lobby_v1_lobby_proto_rawDescGZIP(), []int{0}
}

type ListLobbiesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Lobbies       []*Lobby               `protobuf:"bytes,1,rep,name=lobbies,proto3" json:"lobbies,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListLobbiesResponse) Reset() {
	*x = ListLobbiesResponse{}
	mi := &file_lobby_v1_lobby_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLobbiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLobbiesResponse) ProtoMessage() {}

func (x *ListLobbiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_v1_lobby_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLobbiesResponse.ProtoReflect.Descriptor instead.
func (*ListLobbiesResponse) Descriptor() ([]byte, []int) {
	return file_lobby_v1_lobby_proto_rawDescGZIP(), []int{1}
}

func (x *ListLobbiesResponse) GetLobbies() []*Lobby {
	if x != nil {
		return x.Lobbies
	}
	return nil
}

type AddLobbiesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LobbyName     string                 `protobuf:"bytes,1,opt,name=lobby_name,json=lobbyName,proto3" json:"lobby_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddLobbiesRequest) Reset() {
	*x = AddLobbiesRequest{}
	mi := &file_lobby_v1_lobby_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddLobbiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLobbiesRequest) ProtoMessage() {}

func (x *AddLobbiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_v1_lobby_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLobbiesRequest.ProtoReflect.Descriptor instead.
func (*AddLobbiesRequest) Descriptor() ([]byte, []int) {
	return file_lobby_v1_lobby_proto_rawDescGZIP(), []int{2}
}

func (x *AddLobbiesRequest) GetLobbyName() string {
	if x != nil {
		return x.LobbyName
	}
	return ""
}

type AddLobbiesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddLobbiesResponse) Reset() {
	*x = AddLobbiesResponse{}
	mi := &file_lobby_v1_lobby_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddLobbiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLobbiesResponse) ProtoMessage() {}

func (x *AddLobbiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_v1_lobby_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLobbiesResponse.ProtoReflect.Descriptor instead.
func (*AddLobbiesResponse) Descriptor() ([]byte, []int) {
	return file_lobby_v1_lobby_proto_rawDescGZIP(), []int{3}
}

type DelLobbiesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Lobby         *Lobby                 `protobuf:"bytes,1,opt,name=lobby,proto3" json:"lobby,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelLobbiesRequest) Reset() {
	*x = DelLobbiesRequest{}
	mi := &file_lobby_v1_lobby_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelLobbiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelLobbiesRequest) ProtoMessage() {}

func (x *DelLobbiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_v1_lobby_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelLobbiesRequest.ProtoReflect.Descriptor instead.
func (*DelLobbiesRequest) Descriptor() ([]byte, []int) {
	return file_lobby_v1_lobby_proto_rawDescGZIP(), []int{4}
}

func (x *DelLobbiesRequest) GetLobby() *Lobby {
	if x != nil {
		return x.Lobby
	}
	return nil
}

type DelLobbiesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelLobbiesResponse) Reset() {
	*x = DelLobbiesResponse{}
	mi := &file_lobby_v1_lobby_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelLobbiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelLobbiesResponse) ProtoMessage() {}

func (x *DelLobbiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_v1_lobby_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelLobbiesResponse.ProtoReflect.Descriptor instead.
func (*DelLobbiesResponse) Descriptor() ([]byte, []int) {
	return file_lobby_v1_lobby_proto_rawDescGZIP(), []int{5}
}

type Lobby struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            uint64                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	LobbyName     string                 `protobuf:"bytes,2,opt,name=lobby_name,json=lobbyName,proto3" json:"lobby_name,omitempty"`
	OwnerName     string                 `protobuf:"bytes,4,opt,name=ownerName,proto3" json:"ownerName,omitempty"`
	OwnerId       uint64                 `protobuf:"varint,5,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Lobby) Reset() {
	*x = Lobby{}
	mi := &file_lobby_v1_lobby_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Lobby) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lobby) ProtoMessage() {}

func (x *Lobby) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_v1_lobby_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lobby.ProtoReflect.Descriptor instead.
func (*Lobby) Descriptor() ([]byte, []int) {
	return file_lobby_v1_lobby_proto_rawDescGZIP(), []int{6}
}

func (x *Lobby) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Lobby) GetLobbyName() string {
	if x != nil {
		return x.LobbyName
	}
	return ""
}

func (x *Lobby) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

func (x *Lobby) GetOwnerId() uint64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *Lobby) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_lobby_v1_lobby_proto protoreflect.FileDescriptor

var file_lobby_v1_lobby_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x62, 0x62, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x76, 0x31,
	0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f,
	0x62, 0x62, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x07, 0x6c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52,
	0x07, 0x6c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4c,
	0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12,
	0x41, 0x64, 0x64, 0x4c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3a, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x4c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x6c, 0x6f, 0x62, 0x62, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x05, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x22, 0x14,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x4c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x05, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1d,
	0x0a, 0x0a, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x32, 0xf3, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x62,
	0x62, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x47, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x62, 0x62,
	0x79, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x4c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x62,
	0x62, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x12, 0x1b, 0x2e,
	0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x4c, 0x6f, 0x62, 0x62,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x6f, 0x62,
	0x62, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x4c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x8c, 0x01, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x41, 0x33, 0x34, 0x31, 0x2f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x61, 0x63, 0x6d, 0x61, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4c, 0x58, 0x58,
	0xaa, 0x02, 0x08, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x4c, 0x6f,
	0x62, 0x62, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09,
	0x4c, 0x6f, 0x62, 0x62, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_lobby_v1_lobby_proto_rawDescOnce sync.Once
	file_lobby_v1_lobby_proto_rawDescData = file_lobby_v1_lobby_proto_rawDesc
)

func file_lobby_v1_lobby_proto_rawDescGZIP() []byte {
	file_lobby_v1_lobby_proto_rawDescOnce.Do(func() {
		file_lobby_v1_lobby_proto_rawDescData = protoimpl.X.CompressGZIP(file_lobby_v1_lobby_proto_rawDescData)
	})
	return file_lobby_v1_lobby_proto_rawDescData
}

var file_lobby_v1_lobby_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_lobby_v1_lobby_proto_goTypes = []any{
	(*ListLobbiesRequest)(nil),  // 0: lobby.v1.ListLobbiesRequest
	(*ListLobbiesResponse)(nil), // 1: lobby.v1.ListLobbiesResponse
	(*AddLobbiesRequest)(nil),   // 2: lobby.v1.AddLobbiesRequest
	(*AddLobbiesResponse)(nil),  // 3: lobby.v1.AddLobbiesResponse
	(*DelLobbiesRequest)(nil),   // 4: lobby.v1.DelLobbiesRequest
	(*DelLobbiesResponse)(nil),  // 5: lobby.v1.DelLobbiesResponse
	(*Lobby)(nil),               // 6: lobby.v1.Lobby
}
var file_lobby_v1_lobby_proto_depIdxs = []int32{
	6, // 0: lobby.v1.ListLobbiesResponse.lobbies:type_name -> lobby.v1.Lobby
	6, // 1: lobby.v1.DelLobbiesRequest.lobby:type_name -> lobby.v1.Lobby
	0, // 2: lobby.v1.LobbyService.ListLobbies:input_type -> lobby.v1.ListLobbiesRequest
	2, // 3: lobby.v1.LobbyService.AddLobby:input_type -> lobby.v1.AddLobbiesRequest
	4, // 4: lobby.v1.LobbyService.DeleteLobby:input_type -> lobby.v1.DelLobbiesRequest
	1, // 5: lobby.v1.LobbyService.ListLobbies:output_type -> lobby.v1.ListLobbiesResponse
	3, // 6: lobby.v1.LobbyService.AddLobby:output_type -> lobby.v1.AddLobbiesResponse
	5, // 7: lobby.v1.LobbyService.DeleteLobby:output_type -> lobby.v1.DelLobbiesResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_lobby_v1_lobby_proto_init() }
func file_lobby_v1_lobby_proto_init() {
	if File_lobby_v1_lobby_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lobby_v1_lobby_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lobby_v1_lobby_proto_goTypes,
		DependencyIndexes: file_lobby_v1_lobby_proto_depIdxs,
		MessageInfos:      file_lobby_v1_lobby_proto_msgTypes,
	}.Build()
	File_lobby_v1_lobby_proto = out.File
	file_lobby_v1_lobby_proto_rawDesc = nil
	file_lobby_v1_lobby_proto_goTypes = nil
	file_lobby_v1_lobby_proto_depIdxs = nil
}
