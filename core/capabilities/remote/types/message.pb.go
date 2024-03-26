// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: core/capabilities/remote/types/message.proto

package types

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

type Error int32

const (
	Error_OK                   Error = 0
	Error_VALIDATION_FAILED    Error = 1
	Error_CAPABILITY_NOT_FOUND Error = 2
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0: "OK",
		1: "VALIDATION_FAILED",
		2: "CAPABILITY_NOT_FOUND",
	}
	Error_value = map[string]int32{
		"OK":                   0,
		"VALIDATION_FAILED":    1,
		"CAPABILITY_NOT_FOUND": 2,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_core_capabilities_remote_types_message_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_core_capabilities_remote_types_message_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_core_capabilities_remote_types_message_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Body      []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"` // proto-encoded MessageBody to sign
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_capabilities_remote_types_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_core_capabilities_remote_types_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_core_capabilities_remote_types_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Message) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type MessageBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version      uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Sender       []byte `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver     []byte `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	MessageId    []byte `protobuf:"bytes,4,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"` // scoped to (don_id, capability_id)
	CapabilityId string `protobuf:"bytes,5,opt,name=capability_id,json=capabilityId,proto3" json:"capability_id,omitempty"`
	DonId        string `protobuf:"bytes,6,opt,name=don_id,json=donId,proto3" json:"don_id,omitempty"` // where the capability actually lives
	Method       string `protobuf:"bytes,7,opt,name=method,proto3" json:"method,omitempty"`
	Timestamp    int64  `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Payload      []byte `protobuf:"bytes,9,opt,name=payload,proto3" json:"payload,omitempty"`
	Error        Error  `protobuf:"varint,10,opt,name=error,proto3,enum=remote.Error" json:"error,omitempty"`
}

func (x *MessageBody) Reset() {
	*x = MessageBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_capabilities_remote_types_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageBody) ProtoMessage() {}

func (x *MessageBody) ProtoReflect() protoreflect.Message {
	mi := &file_core_capabilities_remote_types_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageBody.ProtoReflect.Descriptor instead.
func (*MessageBody) Descriptor() ([]byte, []int) {
	return file_core_capabilities_remote_types_message_proto_rawDescGZIP(), []int{1}
}

func (x *MessageBody) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *MessageBody) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *MessageBody) GetReceiver() []byte {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *MessageBody) GetMessageId() []byte {
	if x != nil {
		return x.MessageId
	}
	return nil
}

func (x *MessageBody) GetCapabilityId() string {
	if x != nil {
		return x.CapabilityId
	}
	return ""
}

func (x *MessageBody) GetDonId() string {
	if x != nil {
		return x.DonId
	}
	return ""
}

func (x *MessageBody) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *MessageBody) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MessageBody) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *MessageBody) GetError() Error {
	if x != nil {
		return x.Error
	}
	return Error_OK
}

var File_core_capabilities_remote_types_message_proto protoreflect.FileDescriptor

var file_core_capabilities_remote_types_message_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x22, 0x3b, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x22, 0xab, 0x02, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2a, 0x40, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x41, 0x50,
	0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x02, 0x42, 0x20, 0x5a, 0x1e, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_capabilities_remote_types_message_proto_rawDescOnce sync.Once
	file_core_capabilities_remote_types_message_proto_rawDescData = file_core_capabilities_remote_types_message_proto_rawDesc
)

func file_core_capabilities_remote_types_message_proto_rawDescGZIP() []byte {
	file_core_capabilities_remote_types_message_proto_rawDescOnce.Do(func() {
		file_core_capabilities_remote_types_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_capabilities_remote_types_message_proto_rawDescData)
	})
	return file_core_capabilities_remote_types_message_proto_rawDescData
}

var file_core_capabilities_remote_types_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_capabilities_remote_types_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_core_capabilities_remote_types_message_proto_goTypes = []interface{}{
	(Error)(0),          // 0: remote.Error
	(*Message)(nil),     // 1: remote.Message
	(*MessageBody)(nil), // 2: remote.MessageBody
}
var file_core_capabilities_remote_types_message_proto_depIdxs = []int32{
	0, // 0: remote.MessageBody.error:type_name -> remote.Error
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_capabilities_remote_types_message_proto_init() }
func file_core_capabilities_remote_types_message_proto_init() {
	if File_core_capabilities_remote_types_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_capabilities_remote_types_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_core_capabilities_remote_types_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageBody); i {
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
			RawDescriptor: file_core_capabilities_remote_types_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_capabilities_remote_types_message_proto_goTypes,
		DependencyIndexes: file_core_capabilities_remote_types_message_proto_depIdxs,
		EnumInfos:         file_core_capabilities_remote_types_message_proto_enumTypes,
		MessageInfos:      file_core_capabilities_remote_types_message_proto_msgTypes,
	}.Build()
	File_core_capabilities_remote_types_message_proto = out.File
	file_core_capabilities_remote_types_message_proto_rawDesc = nil
	file_core_capabilities_remote_types_message_proto_goTypes = nil
	file_core_capabilities_remote_types_message_proto_depIdxs = nil
}
