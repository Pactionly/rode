// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: proto/v1alpha1/rode-attest.proto

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type AttestPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy      string `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	ResourceURI string `protobuf:"bytes,2,opt,name=resourceURI,proto3" json:"resourceURI,omitempty"`
}

func (x *AttestPolicyRequest) Reset() {
	*x = AttestPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_rode_attest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestPolicyRequest) ProtoMessage() {}

func (x *AttestPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_rode_attest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestPolicyRequest.ProtoReflect.Descriptor instead.
func (*AttestPolicyRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_rode_attest_proto_rawDescGZIP(), []int{0}
}

func (x *AttestPolicyRequest) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *AttestPolicyRequest) GetResourceURI() string {
	if x != nil {
		return x.ResourceURI
	}
	return ""
}

type AttestPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allow        bool                       `protobuf:"varint,1,opt,name=allow,proto3" json:"allow,omitempty"`
	Changed      bool                       `protobuf:"varint,2,opt,name=changed,proto3" json:"changed,omitempty"`
	Attestations []*AttestPolicyAttestation `protobuf:"bytes,3,rep,name=attestations,proto3" json:"attestations,omitempty"`
}

func (x *AttestPolicyResponse) Reset() {
	*x = AttestPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_rode_attest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestPolicyResponse) ProtoMessage() {}

func (x *AttestPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_rode_attest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestPolicyResponse.ProtoReflect.Descriptor instead.
func (*AttestPolicyResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_rode_attest_proto_rawDescGZIP(), []int{1}
}

func (x *AttestPolicyResponse) GetAllow() bool {
	if x != nil {
		return x.Allow
	}
	return false
}

func (x *AttestPolicyResponse) GetChanged() bool {
	if x != nil {
		return x.Changed
	}
	return false
}

func (x *AttestPolicyResponse) GetAttestations() []*AttestPolicyAttestation {
	if x != nil {
		return x.Attestations
	}
	return nil
}

type AttestPolicyAttestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allow      bool                     `protobuf:"varint,1,opt,name=allow,proto3" json:"allow,omitempty"`
	Created    *timestamp.Timestamp     `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Violations []*AttestPolicyViolation `protobuf:"bytes,3,rep,name=violations,proto3" json:"violations,omitempty"`
}

func (x *AttestPolicyAttestation) Reset() {
	*x = AttestPolicyAttestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_rode_attest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestPolicyAttestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestPolicyAttestation) ProtoMessage() {}

func (x *AttestPolicyAttestation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_rode_attest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestPolicyAttestation.ProtoReflect.Descriptor instead.
func (*AttestPolicyAttestation) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_rode_attest_proto_rawDescGZIP(), []int{2}
}

func (x *AttestPolicyAttestation) GetAllow() bool {
	if x != nil {
		return x.Allow
	}
	return false
}

func (x *AttestPolicyAttestation) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *AttestPolicyAttestation) GetViolations() []*AttestPolicyViolation {
	if x != nil {
		return x.Violations
	}
	return nil
}

type AttestPolicyViolation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Link        string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *AttestPolicyViolation) Reset() {
	*x = AttestPolicyViolation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_rode_attest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestPolicyViolation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestPolicyViolation) ProtoMessage() {}

func (x *AttestPolicyViolation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_rode_attest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestPolicyViolation.ProtoReflect.Descriptor instead.
func (*AttestPolicyViolation) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_rode_attest_proto_rawDescGZIP(), []int{3}
}

func (x *AttestPolicyViolation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AttestPolicyViolation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttestPolicyViolation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AttestPolicyViolation) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_proto_v1alpha1_rode_attest_proto protoreflect.FileDescriptor

var file_proto_v1alpha1_rode_attest_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x72, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x13, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x52, 0x49,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x55, 0x52, 0x49, 0x22, 0x92, 0x01, 0x0a, 0x14, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x0c,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x17, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x44, 0x0a, 0x0a, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x76, 0x69, 0x6f, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x71, 0x0a, 0x15, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x6f, 0x64,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1alpha1_rode_attest_proto_rawDescOnce sync.Once
	file_proto_v1alpha1_rode_attest_proto_rawDescData = file_proto_v1alpha1_rode_attest_proto_rawDesc
)

func file_proto_v1alpha1_rode_attest_proto_rawDescGZIP() []byte {
	file_proto_v1alpha1_rode_attest_proto_rawDescOnce.Do(func() {
		file_proto_v1alpha1_rode_attest_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1alpha1_rode_attest_proto_rawDescData)
	})
	return file_proto_v1alpha1_rode_attest_proto_rawDescData
}

var file_proto_v1alpha1_rode_attest_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_v1alpha1_rode_attest_proto_goTypes = []interface{}{
	(*AttestPolicyRequest)(nil),     // 0: rode.v1alpha1.AttestPolicyRequest
	(*AttestPolicyResponse)(nil),    // 1: rode.v1alpha1.AttestPolicyResponse
	(*AttestPolicyAttestation)(nil), // 2: rode.v1alpha1.AttestPolicyAttestation
	(*AttestPolicyViolation)(nil),   // 3: rode.v1alpha1.AttestPolicyViolation
	(*timestamp.Timestamp)(nil),     // 4: google.protobuf.Timestamp
}
var file_proto_v1alpha1_rode_attest_proto_depIdxs = []int32{
	2, // 0: rode.v1alpha1.AttestPolicyResponse.attestations:type_name -> rode.v1alpha1.AttestPolicyAttestation
	4, // 1: rode.v1alpha1.AttestPolicyAttestation.created:type_name -> google.protobuf.Timestamp
	3, // 2: rode.v1alpha1.AttestPolicyAttestation.violations:type_name -> rode.v1alpha1.AttestPolicyViolation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_v1alpha1_rode_attest_proto_init() }
func file_proto_v1alpha1_rode_attest_proto_init() {
	if File_proto_v1alpha1_rode_attest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1alpha1_rode_attest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestPolicyRequest); i {
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
		file_proto_v1alpha1_rode_attest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestPolicyResponse); i {
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
		file_proto_v1alpha1_rode_attest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestPolicyAttestation); i {
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
		file_proto_v1alpha1_rode_attest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestPolicyViolation); i {
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
			RawDescriptor: file_proto_v1alpha1_rode_attest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_v1alpha1_rode_attest_proto_goTypes,
		DependencyIndexes: file_proto_v1alpha1_rode_attest_proto_depIdxs,
		MessageInfos:      file_proto_v1alpha1_rode_attest_proto_msgTypes,
	}.Build()
	File_proto_v1alpha1_rode_attest_proto = out.File
	file_proto_v1alpha1_rode_attest_proto_rawDesc = nil
	file_proto_v1alpha1_rode_attest_proto_goTypes = nil
	file_proto_v1alpha1_rode_attest_proto_depIdxs = nil
}
