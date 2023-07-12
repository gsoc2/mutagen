// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: synchronization/core/permissions_mode.proto

package core

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

// PermissionsMode specifies the mode for handling permission propagation.
type PermissionsMode int32

const (
	// PermissionsMode_PermissionsModeDefault represents an unspecified
	// permissions mode. It is not valid for use with Scan. It should be
	// converted to one of the following values based on the desired default
	// behavior.
	PermissionsMode_PermissionsModeDefault PermissionsMode = 0
	// PermissionsMode_PermissionsModePortable specifies that permissions should
	// be propagated in a portable fashion. This means that only executability
	// bits are managed by Mutagen and that manual specifications for ownership
	// and base file permissions are used.
	PermissionsMode_PermissionsModePortable PermissionsMode = 1
	// PermissionsMode_PermissionsModeManual specifies that only manual
	// permission specifications should be used. In this case, Mutagen does not
	// perform any propagation of permissions.
	PermissionsMode_PermissionsModeManual PermissionsMode = 2
)

// Enum value maps for PermissionsMode.
var (
	PermissionsMode_name = map[int32]string{
		0: "PermissionsModeDefault",
		1: "PermissionsModePortable",
		2: "PermissionsModeManual",
	}
	PermissionsMode_value = map[string]int32{
		"PermissionsModeDefault":  0,
		"PermissionsModePortable": 1,
		"PermissionsModeManual":   2,
	}
)

func (x PermissionsMode) Enum() *PermissionsMode {
	p := new(PermissionsMode)
	*p = x
	return p
}

func (x PermissionsMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionsMode) Descriptor() protoreflect.EnumDescriptor {
	return file_synchronization_core_permissions_mode_proto_enumTypes[0].Descriptor()
}

func (PermissionsMode) Type() protoreflect.EnumType {
	return &file_synchronization_core_permissions_mode_proto_enumTypes[0]
}

func (x PermissionsMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionsMode.Descriptor instead.
func (PermissionsMode) EnumDescriptor() ([]byte, []int) {
	return file_synchronization_core_permissions_mode_proto_rawDescGZIP(), []int{0}
}

var File_synchronization_core_permissions_mode_proto protoreflect.FileDescriptor

var file_synchronization_core_permissions_mode_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63,
	0x6f, 0x72, 0x65, 0x2a, 0x65, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x4d, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x6f,
	0x64, 0x65, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x10, 0x02, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e,
	0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_core_permissions_mode_proto_rawDescOnce sync.Once
	file_synchronization_core_permissions_mode_proto_rawDescData = file_synchronization_core_permissions_mode_proto_rawDesc
)

func file_synchronization_core_permissions_mode_proto_rawDescGZIP() []byte {
	file_synchronization_core_permissions_mode_proto_rawDescOnce.Do(func() {
		file_synchronization_core_permissions_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_core_permissions_mode_proto_rawDescData)
	})
	return file_synchronization_core_permissions_mode_proto_rawDescData
}

var file_synchronization_core_permissions_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_synchronization_core_permissions_mode_proto_goTypes = []interface{}{
	(PermissionsMode)(0), // 0: core.PermissionsMode
}
var file_synchronization_core_permissions_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_synchronization_core_permissions_mode_proto_init() }
func file_synchronization_core_permissions_mode_proto_init() {
	if File_synchronization_core_permissions_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_synchronization_core_permissions_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_core_permissions_mode_proto_goTypes,
		DependencyIndexes: file_synchronization_core_permissions_mode_proto_depIdxs,
		EnumInfos:         file_synchronization_core_permissions_mode_proto_enumTypes,
	}.Build()
	File_synchronization_core_permissions_mode_proto = out.File
	file_synchronization_core_permissions_mode_proto_rawDesc = nil
	file_synchronization_core_permissions_mode_proto_goTypes = nil
	file_synchronization_core_permissions_mode_proto_depIdxs = nil
}
