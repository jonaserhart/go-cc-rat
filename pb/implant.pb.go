// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: implant.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_implant_proto protoreflect.FileDescriptor

var file_implant_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x69, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x69, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x67, 0x0a, 0x07, 0x49, 0x6d, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x12, 0x2e, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x2c, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x08, 0x5a, 0x06, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var file_implant_proto_goTypes = []any{
	(*Empty)(nil),   // 0: common.Empty
	(*Command)(nil), // 1: common.Command
}
var file_implant_proto_depIdxs = []int32{
	0, // 0: implant.Implant.FetchCommand:input_type -> common.Empty
	1, // 1: implant.Implant.SendOutput:input_type -> common.Command
	1, // 2: implant.Implant.FetchCommand:output_type -> common.Command
	0, // 3: implant.Implant.SendOutput:output_type -> common.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_implant_proto_init() }
func file_implant_proto_init() {
	if File_implant_proto != nil {
		return
	}
	file_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_implant_proto_rawDesc), len(file_implant_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_implant_proto_goTypes,
		DependencyIndexes: file_implant_proto_depIdxs,
	}.Build()
	File_implant_proto = out.File
	file_implant_proto_goTypes = nil
	file_implant_proto_depIdxs = nil
}
