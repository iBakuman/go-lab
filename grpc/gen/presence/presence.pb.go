// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: presence/presence.proto

package presence

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

type MessageA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *MessageA) Reset() {
	*x = MessageA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presence_presence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageA) ProtoMessage() {}

func (x *MessageA) ProtoReflect() protoreflect.Message {
	mi := &file_presence_presence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageA.ProtoReflect.Descriptor instead.
func (*MessageA) Descriptor() ([]byte, []int) {
	return file_presence_presence_proto_rawDescGZIP(), []int{0}
}

func (x *MessageA) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

type MessageB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32     `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B *MessageA `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C *int32    `protobuf:"varint,3,opt,name=c,proto3,oneof" json:"c,omitempty"`
}

func (x *MessageB) Reset() {
	*x = MessageB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presence_presence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageB) ProtoMessage() {}

func (x *MessageB) ProtoReflect() protoreflect.Message {
	mi := &file_presence_presence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageB.ProtoReflect.Descriptor instead.
func (*MessageB) Descriptor() ([]byte, []int) {
	return file_presence_presence_proto_rawDescGZIP(), []int{1}
}

func (x *MessageB) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *MessageB) GetB() *MessageA {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *MessageB) GetC() int32 {
	if x != nil && x.C != nil {
		return *x.C
	}
	return 0
}

type MessageC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Resp:
	//
	//	*MessageC_A
	//	*MessageC_B
	Resp isMessageC_Resp `protobuf_oneof:"resp"`
}

func (x *MessageC) Reset() {
	*x = MessageC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presence_presence_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageC) ProtoMessage() {}

func (x *MessageC) ProtoReflect() protoreflect.Message {
	mi := &file_presence_presence_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageC.ProtoReflect.Descriptor instead.
func (*MessageC) Descriptor() ([]byte, []int) {
	return file_presence_presence_proto_rawDescGZIP(), []int{2}
}

func (m *MessageC) GetResp() isMessageC_Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

func (x *MessageC) GetA() *MessageA {
	if x, ok := x.GetResp().(*MessageC_A); ok {
		return x.A
	}
	return nil
}

func (x *MessageC) GetB() *MessageB {
	if x, ok := x.GetResp().(*MessageC_B); ok {
		return x.B
	}
	return nil
}

type isMessageC_Resp interface {
	isMessageC_Resp()
}

type MessageC_A struct {
	A *MessageA `protobuf:"bytes,1,opt,name=a,proto3,oneof"`
}

type MessageC_B struct {
	B *MessageB `protobuf:"bytes,2,opt,name=b,proto3,oneof"`
}

func (*MessageC_A) isMessageC_Resp() {}

func (*MessageC_B) isMessageC_Resp() {}

var File_presence_presence_proto protoreflect.FileDescriptor

var file_presence_presence_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x22, 0x18, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x12,
	0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x22, 0x53, 0x0a,
	0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x20, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x52, 0x01, 0x62, 0x12, 0x11, 0x0a, 0x01, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x01, 0x63, 0x88, 0x01, 0x01, 0x42, 0x04, 0x0a, 0x02,
	0x5f, 0x63, 0x22, 0x5a, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x12, 0x22,
	0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x48, 0x00, 0x52,
	0x01, 0x61, 0x12, 0x22, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x48, 0x00, 0x52, 0x01, 0x62, 0x42, 0x06, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x42, 0x8b,
	0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42,
	0x0d, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x62, 0x61,
	0x6b, 0x75, 0x6d, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x61, 0x62, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0xa2, 0x02,
	0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0xca,
	0x02, 0x08, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0xe2, 0x02, 0x14, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x08, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_presence_presence_proto_rawDescOnce sync.Once
	file_presence_presence_proto_rawDescData = file_presence_presence_proto_rawDesc
)

func file_presence_presence_proto_rawDescGZIP() []byte {
	file_presence_presence_proto_rawDescOnce.Do(func() {
		file_presence_presence_proto_rawDescData = protoimpl.X.CompressGZIP(file_presence_presence_proto_rawDescData)
	})
	return file_presence_presence_proto_rawDescData
}

var file_presence_presence_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_presence_presence_proto_goTypes = []any{
	(*MessageA)(nil), // 0: presence.MessageA
	(*MessageB)(nil), // 1: presence.MessageB
	(*MessageC)(nil), // 2: presence.MessageC
}
var file_presence_presence_proto_depIdxs = []int32{
	0, // 0: presence.MessageB.b:type_name -> presence.MessageA
	0, // 1: presence.MessageC.a:type_name -> presence.MessageA
	1, // 2: presence.MessageC.b:type_name -> presence.MessageB
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_presence_presence_proto_init() }
func file_presence_presence_proto_init() {
	if File_presence_presence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_presence_presence_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MessageA); i {
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
		file_presence_presence_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MessageB); i {
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
		file_presence_presence_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MessageC); i {
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
	file_presence_presence_proto_msgTypes[1].OneofWrappers = []any{}
	file_presence_presence_proto_msgTypes[2].OneofWrappers = []any{
		(*MessageC_A)(nil),
		(*MessageC_B)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_presence_presence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_presence_presence_proto_goTypes,
		DependencyIndexes: file_presence_presence_proto_depIdxs,
		MessageInfos:      file_presence_presence_proto_msgTypes,
	}.Build()
	File_presence_presence_proto = out.File
	file_presence_presence_proto_rawDesc = nil
	file_presence_presence_proto_goTypes = nil
	file_presence_presence_proto_depIdxs = nil
}
