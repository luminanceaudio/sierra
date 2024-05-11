// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: apprequests.proto

package models

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

type GetSamplesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Samples []*Sample `protobuf:"bytes,1,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *GetSamplesResponse) Reset() {
	*x = GetSamplesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apprequests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSamplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSamplesResponse) ProtoMessage() {}

func (x *GetSamplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apprequests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSamplesResponse.ProtoReflect.Descriptor instead.
func (*GetSamplesResponse) Descriptor() ([]byte, []int) {
	return file_apprequests_proto_rawDescGZIP(), []int{0}
}

func (x *GetSamplesResponse) GetSamples() []*Sample {
	if x != nil {
		return x.Samples
	}
	return nil
}

var File_apprequests_proto protoreflect.FileDescriptor

var file_apprequests_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x70, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x69, 0x65, 0x72, 0x72, 0x61, 0x5f, 0x61, 0x70, 0x70, 0x1a,
	0x0d, 0x61, 0x70, 0x70, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x69, 0x65, 0x72, 0x72, 0x61, 0x5f, 0x61,
	0x70, 0x70, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apprequests_proto_rawDescOnce sync.Once
	file_apprequests_proto_rawDescData = file_apprequests_proto_rawDesc
)

func file_apprequests_proto_rawDescGZIP() []byte {
	file_apprequests_proto_rawDescOnce.Do(func() {
		file_apprequests_proto_rawDescData = protoimpl.X.CompressGZIP(file_apprequests_proto_rawDescData)
	})
	return file_apprequests_proto_rawDescData
}

var file_apprequests_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_apprequests_proto_goTypes = []interface{}{
	(*GetSamplesResponse)(nil), // 0: sierra_app.GetSamplesResponse
	(*Sample)(nil),             // 1: sierra_app.Sample
}
var file_apprequests_proto_depIdxs = []int32{
	1, // 0: sierra_app.GetSamplesResponse.samples:type_name -> sierra_app.Sample
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apprequests_proto_init() }
func file_apprequests_proto_init() {
	if File_apprequests_proto != nil {
		return
	}
	file_appbase_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apprequests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSamplesResponse); i {
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
			RawDescriptor: file_apprequests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apprequests_proto_goTypes,
		DependencyIndexes: file_apprequests_proto_depIdxs,
		MessageInfos:      file_apprequests_proto_msgTypes,
	}.Build()
	File_apprequests_proto = out.File
	file_apprequests_proto_rawDesc = nil
	file_apprequests_proto_goTypes = nil
	file_apprequests_proto_depIdxs = nil
}
