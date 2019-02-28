// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/deprecation.proto

package grpc

import (
	proto "github.com/golang/protobuf/proto"
	protoapi "github.com/golang/protobuf/protoapi"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("grpc/deprecation.proto", xxx_File_grpc_deprecation_proto_rawdesc_gzipped)
}

var xxx_File_grpc_deprecation_proto_rawdesc = []byte{
	// 245 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x16, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x6c,
	0x0a, 0x11, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x1a, 0x03, 0x88, 0x02, 0x01, 0x42, 0x3f, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var xxx_File_grpc_deprecation_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_grpc_deprecation_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_grpc_deprecation_proto protoreflect.FileDescriptor

var xxx_File_grpc_deprecation_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: goproto.protoc.grpc.Request
	(*Response)(nil), // 1: goproto.protoc.grpc.Response
}
var xxx_File_grpc_deprecation_proto_depIdxs = []int32{
	0, // goproto.protoc.grpc.DeprecatedService.DeprecatedCall:input_type -> goproto.protoc.grpc.Request
	1, // goproto.protoc.grpc.DeprecatedService.DeprecatedCall:output_type -> goproto.protoc.grpc.Response
}

func init() {
	File_grpc_deprecation_proto = protoimpl.FileBuilder{
		RawDescriptor:     xxx_File_grpc_deprecation_proto_rawdesc,
		GoTypes:           xxx_File_grpc_deprecation_proto_goTypes,
		DependencyIndexes: xxx_File_grpc_deprecation_proto_depIdxs,
	}.Init()
	xxx_File_grpc_deprecation_proto_goTypes = nil
	xxx_File_grpc_deprecation_proto_depIdxs = nil
}
