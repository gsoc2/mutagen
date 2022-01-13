// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: url/url.proto

package url

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

// Kind indicates the kind of a URL.
type Kind int32

const (
	// Synchronization indicates a synchronization URL.
	Kind_Synchronization Kind = 0
	// Forwarding indicates a forwarding URL.
	Kind_Forwarding Kind = 1
)

// Enum value maps for Kind.
var (
	Kind_name = map[int32]string{
		0: "Synchronization",
		1: "Forwarding",
	}
	Kind_value = map[string]int32{
		"Synchronization": 0,
		"Forwarding":      1,
	}
)

func (x Kind) Enum() *Kind {
	p := new(Kind)
	*p = x
	return p
}

func (x Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_url_url_proto_enumTypes[0].Descriptor()
}

func (Kind) Type() protoreflect.EnumType {
	return &file_url_url_proto_enumTypes[0]
}

func (x Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Kind.Descriptor instead.
func (Kind) EnumDescriptor() ([]byte, []int) {
	return file_url_url_proto_rawDescGZIP(), []int{0}
}

// Protocol indicates a location type.
type Protocol int32

const (
	// Local indicates that the resource is on the local system.
	Protocol_Local Protocol = 0
	// SSH indicates that the resource is accessible via SSH.
	Protocol_SSH Protocol = 1
	// Docker indicates that the resource is inside a Docker container.
	Protocol_Docker Protocol = 11
)

// Enum value maps for Protocol.
var (
	Protocol_name = map[int32]string{
		0:  "Local",
		1:  "SSH",
		11: "Docker",
	}
	Protocol_value = map[string]int32{
		"Local":  0,
		"SSH":    1,
		"Docker": 11,
	}
)

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_url_url_proto_enumTypes[1].Descriptor()
}

func (Protocol) Type() protoreflect.EnumType {
	return &file_url_url_proto_enumTypes[1]
}

func (x Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Protocol.Descriptor instead.
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return file_url_url_proto_rawDescGZIP(), []int{1}
}

// URL represents a pointer to a resource. It should be considered immutable.
type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kind indicates the URL kind.
	// NOTE: This field number is out of order for historical reasons.
	Kind Kind `protobuf:"varint,7,opt,name=kind,proto3,enum=url.Kind" json:"kind,omitempty"`
	// Protocol indicates a location type.
	Protocol Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=url.Protocol" json:"protocol,omitempty"`
	// User is the user under which a resource should be accessed.
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// Host is protocol-specific, but generally indicates the location of the
	// remote.
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	// Port indicates a TCP port via which to access the remote location, if
	// applicable.
	Port uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	// Path indicates the path of a resource.
	Path string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	// Environment contains captured environment variable information. It is not
	// a required component and its contents and their behavior depend on the
	// transport implementation.
	Environment map[string]string `protobuf:"bytes,6,rep,name=environment,proto3" json:"environment,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Parameters are internal transport parameters. These are set for URLs
	// generated internally that require additional metadata. Parameters are not
	// required and their behavior is dependent on the transport implementation.
	Parameters map[string]string `protobuf:"bytes,8,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *URL) Reset() {
	*x = URL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_url_url_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URL.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_url_url_proto_rawDescGZIP(), []int{0}
}

func (x *URL) GetKind() Kind {
	if x != nil {
		return x.Kind
	}
	return Kind_Synchronization
}

func (x *URL) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_Local
}

func (x *URL) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *URL) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *URL) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *URL) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *URL) GetEnvironment() map[string]string {
	if x != nil {
		return x.Environment
	}
	return nil
}

func (x *URL) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

var File_url_url_proto protoreflect.FileDescriptor

var file_url_url_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x75, 0x72, 0x6c, 0x2f, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x95, 0x03, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x1d, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x75, 0x72, 0x6c,
	0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x75, 0x72, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x3b, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x72,
	0x6c, 0x2e, 0x55, 0x52, 0x4c, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x55, 0x52,
	0x4c, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3e, 0x0a,
	0x10, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a,
	0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x2b, 0x0a, 0x04,
	0x4b, 0x69, 0x6e, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x2a, 0x2a, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x53, 0x53, 0x48, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x10, 0x0b, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d,
	0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x72, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_url_url_proto_rawDescOnce sync.Once
	file_url_url_proto_rawDescData = file_url_url_proto_rawDesc
)

func file_url_url_proto_rawDescGZIP() []byte {
	file_url_url_proto_rawDescOnce.Do(func() {
		file_url_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_url_url_proto_rawDescData)
	})
	return file_url_url_proto_rawDescData
}

var file_url_url_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_url_url_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_url_url_proto_goTypes = []interface{}{
	(Kind)(0),     // 0: url.Kind
	(Protocol)(0), // 1: url.Protocol
	(*URL)(nil),   // 2: url.URL
	nil,           // 3: url.URL.EnvironmentEntry
	nil,           // 4: url.URL.ParametersEntry
}
var file_url_url_proto_depIdxs = []int32{
	0, // 0: url.URL.kind:type_name -> url.Kind
	1, // 1: url.URL.protocol:type_name -> url.Protocol
	3, // 2: url.URL.environment:type_name -> url.URL.EnvironmentEntry
	4, // 3: url.URL.parameters:type_name -> url.URL.ParametersEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_url_url_proto_init() }
func file_url_url_proto_init() {
	if File_url_url_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_url_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URL); i {
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
			RawDescriptor: file_url_url_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_url_url_proto_goTypes,
		DependencyIndexes: file_url_url_proto_depIdxs,
		EnumInfos:         file_url_url_proto_enumTypes,
		MessageInfos:      file_url_url_proto_msgTypes,
	}.Build()
	File_url_url_proto = out.File
	file_url_url_proto_rawDesc = nil
	file_url_url_proto_goTypes = nil
	file_url_url_proto_depIdxs = nil
}
