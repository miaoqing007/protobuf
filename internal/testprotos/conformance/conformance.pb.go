// Code generated by protoc-gen-go. DO NOT EDIT.
// source: conformance/conformance.proto

package conformance_proto

import (
	proto "github.com/golang/protobuf/proto"
	protoapi "github.com/golang/protobuf/protoapi"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type WireFormat int32

const (
	WireFormat_UNSPECIFIED WireFormat = 0
	WireFormat_PROTOBUF    WireFormat = 1
	WireFormat_JSON        WireFormat = 2
)

func (e WireFormat) Type() protoreflect.EnumType {
	return xxx_File_conformance_conformance_proto_enumTypes[0]
}
func (e WireFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var WireFormat_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "PROTOBUF",
	2: "JSON",
}

var WireFormat_value = map[string]int32{
	"UNSPECIFIED": 0,
	"PROTOBUF":    1,
	"JSON":        2,
}

func (x WireFormat) String() string {
	return proto.EnumName(WireFormat_name, int32(x))
}

func (WireFormat) EnumDescriptor() ([]byte, []int) {
	return xxx_File_conformance_conformance_proto_rawdesc_gzipped, []int{0}
}

// Represents a single test case's input.  The testee should:
//
//   1. parse this proto (which should always succeed)
//   2. parse the protobuf or JSON payload in "payload" (which may fail)
//   3. if the parse succeeded, serialize the message in the requested format.
type ConformanceRequest struct {
	// The payload (whether protobuf of JSON) is always for a
	// protobuf_test_messages.proto3.TestAllTypes proto (as defined in
	// src/google/protobuf/proto3_test_messages.proto).
	//
	// TODO(haberman): if/when we expand the conformance tests to support proto2,
	// we will want to include a field that lets the payload/response be a
	// protobuf_test_messages.proto2.TestAllTypes message instead.
	//
	// Types that are valid to be assigned to Payload:
	//	*ConformanceRequest_ProtobufPayload
	//	*ConformanceRequest_JsonPayload
	Payload isConformanceRequest_Payload `protobuf_oneof:"payload"`
	// Which format should the testee serialize its message to?
	RequestedOutputFormat WireFormat `protobuf:"varint,3,opt,name=requested_output_format,json=requestedOutputFormat,proto3,enum=conformance.WireFormat" json:"requested_output_format,omitempty"`
	// The full name for the test message to use; for the moment, either:
	// protobuf_test_messages.proto3.TestAllTypesProto3 or
	// protobuf_test_messages.proto2.TestAllTypesProto2.
	MessageType          string   `protobuf:"bytes,4,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConformanceRequest) ProtoReflect() protoreflect.Message {
	return xxx_File_conformance_conformance_proto_messageTypes[0].MessageOf(m)
}
func (m *ConformanceRequest) Reset()         { *m = ConformanceRequest{} }
func (m *ConformanceRequest) String() string { return proto.CompactTextString(m) }
func (*ConformanceRequest) ProtoMessage()    {}
func (*ConformanceRequest) Descriptor() ([]byte, []int) {
	return xxx_File_conformance_conformance_proto_rawdesc_gzipped, []int{0}
}

func (m *ConformanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConformanceRequest.Unmarshal(m, b)
}
func (m *ConformanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConformanceRequest.Marshal(b, m, deterministic)
}
func (m *ConformanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConformanceRequest.Merge(m, src)
}
func (m *ConformanceRequest) XXX_Size() int {
	return xxx_messageInfo_ConformanceRequest.Size(m)
}
func (m *ConformanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConformanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConformanceRequest proto.InternalMessageInfo

type isConformanceRequest_Payload interface {
	isConformanceRequest_Payload()
}

type ConformanceRequest_ProtobufPayload struct {
	ProtobufPayload []byte `protobuf:"bytes,1,opt,name=protobuf_payload,json=protobufPayload,proto3,oneof"`
}

type ConformanceRequest_JsonPayload struct {
	JsonPayload string `protobuf:"bytes,2,opt,name=json_payload,json=jsonPayload,proto3,oneof"`
}

func (*ConformanceRequest_ProtobufPayload) isConformanceRequest_Payload() {}

func (*ConformanceRequest_JsonPayload) isConformanceRequest_Payload() {}

func (m *ConformanceRequest) GetPayload() isConformanceRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ConformanceRequest) GetProtobufPayload() []byte {
	if x, ok := m.GetPayload().(*ConformanceRequest_ProtobufPayload); ok {
		return x.ProtobufPayload
	}
	return nil
}

func (m *ConformanceRequest) GetJsonPayload() string {
	if x, ok := m.GetPayload().(*ConformanceRequest_JsonPayload); ok {
		return x.JsonPayload
	}
	return ""
}

func (m *ConformanceRequest) GetRequestedOutputFormat() WireFormat {
	if m != nil {
		return m.RequestedOutputFormat
	}
	return WireFormat_UNSPECIFIED
}

func (m *ConformanceRequest) GetMessageType() string {
	if m != nil {
		return m.MessageType
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConformanceRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConformanceRequest_ProtobufPayload)(nil),
		(*ConformanceRequest_JsonPayload)(nil),
	}
}

// Represents a single test case's output.
type ConformanceResponse struct {
	// Types that are valid to be assigned to Result:
	// This string should be set to indicate parsing failed.  The string can
	// provide more information about the parse error if it is available.
	//
	// Setting this string does not necessarily mean the testee failed the
	// test.  Some of the test cases are intentionally invalid input.
	//	*ConformanceResponse_ParseError
	// If the input was successfully parsed but errors occurred when
	// serializing it to the requested output format, set the error message in
	// this field.
	//	*ConformanceResponse_SerializeError
	// This should be set if some other error occurred.  This will always
	// indicate that the test failed.  The string can provide more information
	// about the failure.
	//	*ConformanceResponse_RuntimeError
	// If the input was successfully parsed and the requested output was
	// protobuf, serialize it to protobuf and set it in this field.
	//	*ConformanceResponse_ProtobufPayload
	// If the input was successfully parsed and the requested output was JSON,
	// serialize to JSON and set it in this field.
	//	*ConformanceResponse_JsonPayload
	// For when the testee skipped the test, likely because a certain feature
	// wasn't supported, like JSON input/output.
	//	*ConformanceResponse_Skipped
	Result               isConformanceResponse_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ConformanceResponse) ProtoReflect() protoreflect.Message {
	return xxx_File_conformance_conformance_proto_messageTypes[1].MessageOf(m)
}
func (m *ConformanceResponse) Reset()         { *m = ConformanceResponse{} }
func (m *ConformanceResponse) String() string { return proto.CompactTextString(m) }
func (*ConformanceResponse) ProtoMessage()    {}
func (*ConformanceResponse) Descriptor() ([]byte, []int) {
	return xxx_File_conformance_conformance_proto_rawdesc_gzipped, []int{1}
}

func (m *ConformanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConformanceResponse.Unmarshal(m, b)
}
func (m *ConformanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConformanceResponse.Marshal(b, m, deterministic)
}
func (m *ConformanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConformanceResponse.Merge(m, src)
}
func (m *ConformanceResponse) XXX_Size() int {
	return xxx_messageInfo_ConformanceResponse.Size(m)
}
func (m *ConformanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConformanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConformanceResponse proto.InternalMessageInfo

type isConformanceResponse_Result interface {
	isConformanceResponse_Result()
}

type ConformanceResponse_ParseError struct {
	ParseError string `protobuf:"bytes,1,opt,name=parse_error,json=parseError,proto3,oneof"`
}

type ConformanceResponse_SerializeError struct {
	SerializeError string `protobuf:"bytes,6,opt,name=serialize_error,json=serializeError,proto3,oneof"`
}

type ConformanceResponse_RuntimeError struct {
	RuntimeError string `protobuf:"bytes,2,opt,name=runtime_error,json=runtimeError,proto3,oneof"`
}

type ConformanceResponse_ProtobufPayload struct {
	ProtobufPayload []byte `protobuf:"bytes,3,opt,name=protobuf_payload,json=protobufPayload,proto3,oneof"`
}

type ConformanceResponse_JsonPayload struct {
	JsonPayload string `protobuf:"bytes,4,opt,name=json_payload,json=jsonPayload,proto3,oneof"`
}

type ConformanceResponse_Skipped struct {
	Skipped string `protobuf:"bytes,5,opt,name=skipped,proto3,oneof"`
}

func (*ConformanceResponse_ParseError) isConformanceResponse_Result() {}

func (*ConformanceResponse_SerializeError) isConformanceResponse_Result() {}

func (*ConformanceResponse_RuntimeError) isConformanceResponse_Result() {}

func (*ConformanceResponse_ProtobufPayload) isConformanceResponse_Result() {}

func (*ConformanceResponse_JsonPayload) isConformanceResponse_Result() {}

func (*ConformanceResponse_Skipped) isConformanceResponse_Result() {}

func (m *ConformanceResponse) GetResult() isConformanceResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ConformanceResponse) GetParseError() string {
	if x, ok := m.GetResult().(*ConformanceResponse_ParseError); ok {
		return x.ParseError
	}
	return ""
}

func (m *ConformanceResponse) GetSerializeError() string {
	if x, ok := m.GetResult().(*ConformanceResponse_SerializeError); ok {
		return x.SerializeError
	}
	return ""
}

func (m *ConformanceResponse) GetRuntimeError() string {
	if x, ok := m.GetResult().(*ConformanceResponse_RuntimeError); ok {
		return x.RuntimeError
	}
	return ""
}

func (m *ConformanceResponse) GetProtobufPayload() []byte {
	if x, ok := m.GetResult().(*ConformanceResponse_ProtobufPayload); ok {
		return x.ProtobufPayload
	}
	return nil
}

func (m *ConformanceResponse) GetJsonPayload() string {
	if x, ok := m.GetResult().(*ConformanceResponse_JsonPayload); ok {
		return x.JsonPayload
	}
	return ""
}

func (m *ConformanceResponse) GetSkipped() string {
	if x, ok := m.GetResult().(*ConformanceResponse_Skipped); ok {
		return x.Skipped
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConformanceResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConformanceResponse_ParseError)(nil),
		(*ConformanceResponse_SerializeError)(nil),
		(*ConformanceResponse_RuntimeError)(nil),
		(*ConformanceResponse_ProtobufPayload)(nil),
		(*ConformanceResponse_JsonPayload)(nil),
		(*ConformanceResponse_Skipped)(nil),
	}
}

func init() {
	proto.RegisterFile("conformance/conformance.proto", xxx_File_conformance_conformance_proto_rawdesc_gzipped)
	proto.RegisterEnum("conformance.WireFormat", WireFormat_name, WireFormat_value)
	proto.RegisterType((*ConformanceRequest)(nil), "conformance.ConformanceRequest")
	proto.RegisterType((*ConformanceResponse)(nil), "conformance.ConformanceResponse")
}

var xxx_File_conformance_conformance_proto_rawdesc = []byte{
	// 716 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x1d, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xe5, 0x01, 0x0a,
	0x12, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x23, 0x0a, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x4f, 0x0a, 0x17, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x57, 0x69, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52,
	0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x82, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0b,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x29, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0d, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x2b, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23,
	0x0a, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x42,
	0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x35, 0x0a, 0x0a, 0x57, 0x69, 0x72,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x54,
	0x4f, 0x42, 0x55, 0x46, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x02,
	0x42, 0x72, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var xxx_File_conformance_conformance_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_conformance_conformance_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_conformance_conformance_proto protoreflect.FileDescriptor

var xxx_File_conformance_conformance_proto_enumTypes [1]protoreflect.EnumType
var xxx_File_conformance_conformance_proto_messageTypes [2]protoimpl.MessageType
var xxx_File_conformance_conformance_proto_goTypes = []interface{}{
	(WireFormat)(0),             // 0: conformance.WireFormat
	(*ConformanceRequest)(nil),  // 1: conformance.ConformanceRequest
	(*ConformanceResponse)(nil), // 2: conformance.ConformanceResponse
}
var xxx_File_conformance_conformance_proto_depIdxs = []int32{
	0, // conformance.ConformanceRequest.requested_output_format:type_name -> conformance.WireFormat
}

func init() {
	var messageTypes [2]protoreflect.MessageType
	File_conformance_conformance_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_conformance_conformance_proto_rawdesc,
		GoTypes:            xxx_File_conformance_conformance_proto_goTypes,
		DependencyIndexes:  xxx_File_conformance_conformance_proto_depIdxs,
		EnumOutputTypes:    xxx_File_conformance_conformance_proto_enumTypes[:],
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_File_conformance_conformance_proto_goTypes[1:][:2]
	for i, mt := range messageTypes[:] {
		xxx_File_conformance_conformance_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_conformance_conformance_proto_messageTypes[i].PBType = mt
	}
	xxx_File_conformance_conformance_proto_goTypes = nil
	xxx_File_conformance_conformance_proto_depIdxs = nil
}
