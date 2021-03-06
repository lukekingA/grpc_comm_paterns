// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: talkingservice.proto

package talkingservice

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talkingservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_talkingservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_talkingservice_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	UserSaid string `protobuf:"bytes,2,opt,name=userSaid,proto3" json:"userSaid,omitempty"`
	ChatId   string `protobuf:"bytes,3,opt,name=chatId,proto3" json:"chatId,omitempty"`
}

func (x *ChatRequest) Reset() {
	*x = ChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talkingservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequest) ProtoMessage() {}

func (x *ChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talkingservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRequest.ProtoReflect.Descriptor instead.
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return file_talkingservice_proto_rawDescGZIP(), []int{1}
}

func (x *ChatRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ChatRequest) GetUserSaid() string {
	if x != nil {
		return x.UserSaid
	}
	return ""
}

func (x *ChatRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type ChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User          *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	ServerReplied string `protobuf:"bytes,2,opt,name=serverReplied,proto3" json:"serverReplied,omitempty"`
	ChatId        string `protobuf:"bytes,3,opt,name=chatId,proto3" json:"chatId,omitempty"`
	ResponseId    string `protobuf:"bytes,4,opt,name=ResponseId,proto3" json:"ResponseId,omitempty"`
}

func (x *ChatResponse) Reset() {
	*x = ChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talkingservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse) ProtoMessage() {}

func (x *ChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talkingservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse.ProtoReflect.Descriptor instead.
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return file_talkingservice_proto_rawDescGZIP(), []int{2}
}

func (x *ChatResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ChatResponse) GetServerReplied() string {
	if x != nil {
		return x.ServerReplied
	}
	return ""
}

func (x *ChatResponse) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *ChatResponse) GetResponseId() string {
	if x != nil {
		return x.ResponseId
	}
	return ""
}

var File_talkingservice_proto protoreflect.FileDescriptor

var file_talkingservice_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x6b, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x53, 0x61, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x53, 0x61, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22,
	0x96, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x32, 0x59, 0x0a, 0x0e, 0x54, 0x61, 0x6c, 0x6b,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x04, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x7b, 0x5a, 0x79, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x67, 0x73,
	0x2e, 0x6b, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x4c, 0x75, 0x6b, 0x65, 0x2f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x2f, 0x32, 0x30, 0x32, 0x31, 0x2f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x61, 0x6e, 0x64, 0x64,
	0x65, 0x76, 0x6f, 0x70, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x2f, 0x64, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3b, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_talkingservice_proto_rawDescOnce sync.Once
	file_talkingservice_proto_rawDescData = file_talkingservice_proto_rawDesc
)

func file_talkingservice_proto_rawDescGZIP() []byte {
	file_talkingservice_proto_rawDescOnce.Do(func() {
		file_talkingservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_talkingservice_proto_rawDescData)
	})
	return file_talkingservice_proto_rawDescData
}

var file_talkingservice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_talkingservice_proto_goTypes = []interface{}{
	(*User)(nil),         // 0: talkingservice.User
	(*ChatRequest)(nil),  // 1: talkingservice.ChatRequest
	(*ChatResponse)(nil), // 2: talkingservice.ChatResponse
}
var file_talkingservice_proto_depIdxs = []int32{
	0, // 0: talkingservice.ChatRequest.user:type_name -> talkingservice.User
	0, // 1: talkingservice.ChatResponse.user:type_name -> talkingservice.User
	1, // 2: talkingservice.TalkingService.Chat:input_type -> talkingservice.ChatRequest
	2, // 3: talkingservice.TalkingService.Chat:output_type -> talkingservice.ChatResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_talkingservice_proto_init() }
func file_talkingservice_proto_init() {
	if File_talkingservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_talkingservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_talkingservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRequest); i {
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
		file_talkingservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResponse); i {
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
			RawDescriptor: file_talkingservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_talkingservice_proto_goTypes,
		DependencyIndexes: file_talkingservice_proto_depIdxs,
		MessageInfos:      file_talkingservice_proto_msgTypes,
	}.Build()
	File_talkingservice_proto = out.File
	file_talkingservice_proto_rawDesc = nil
	file_talkingservice_proto_goTypes = nil
	file_talkingservice_proto_depIdxs = nil
}
