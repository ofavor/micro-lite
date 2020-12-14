// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: foo.proto

package foo

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  uint32 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_foo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_foo_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age   uint32 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Adult bool   `protobuf:"varint,3,opt,name=adult,proto3" json:"adult,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_foo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_foo_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Response) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Response) GetAdult() bool {
	if x != nil {
		return x.Adult
	}
	return false
}

var File_foo_proto protoreflect.FileDescriptor

var file_foo_proto_rawDesc = []byte{
	0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x66, 0x6f, 0x6f,
	0x22, 0x2f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x22, 0x46, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x75, 0x6c, 0x74, 0x32, 0x2b, 0x0a, 0x03, 0x46, 0x6f, 0x6f,
	0x12, 0x24, 0x0a, 0x03, 0x42, 0x61, 0x72, 0x12, 0x0c, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x66, 0x6f, 0x6f,
	0x3b, 0x66, 0x6f, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foo_proto_rawDescOnce sync.Once
	file_foo_proto_rawDescData = file_foo_proto_rawDesc
)

func file_foo_proto_rawDescGZIP() []byte {
	file_foo_proto_rawDescOnce.Do(func() {
		file_foo_proto_rawDescData = protoimpl.X.CompressGZIP(file_foo_proto_rawDescData)
	})
	return file_foo_proto_rawDescData
}

var file_foo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_foo_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: foo.Request
	(*Response)(nil), // 1: foo.Response
}
var file_foo_proto_depIdxs = []int32{
	0, // 0: foo.Foo.Bar:input_type -> foo.Request
	1, // 1: foo.Foo.Bar:output_type -> foo.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_foo_proto_init() }
func file_foo_proto_init() {
	if File_foo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_foo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_foo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_foo_proto_goTypes,
		DependencyIndexes: file_foo_proto_depIdxs,
		MessageInfos:      file_foo_proto_msgTypes,
	}.Build()
	File_foo_proto = out.File
	file_foo_proto_rawDesc = nil
	file_foo_proto_goTypes = nil
	file_foo_proto_depIdxs = nil
}
