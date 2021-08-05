// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: testsvc/test/v1/enums.proto

package testsvc_test_v1

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

type PersonType int32

const (
	PersonType_PERSON_TYPE_INVALID PersonType = 0
	PersonType_PERSON_TYPE_ONE     PersonType = 1
	PersonType_PERSON_TYPE_TWO     PersonType = 2
)

// Enum value maps for PersonType.
var (
	PersonType_name = map[int32]string{
		0: "PERSON_TYPE_INVALID",
		1: "PERSON_TYPE_ONE",
		2: "PERSON_TYPE_TWO",
	}
	PersonType_value = map[string]int32{
		"PERSON_TYPE_INVALID": 0,
		"PERSON_TYPE_ONE":     1,
		"PERSON_TYPE_TWO":     2,
	}
)

func (x PersonType) Enum() *PersonType {
	p := new(PersonType)
	*p = x
	return p
}

func (x PersonType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PersonType) Descriptor() protoreflect.EnumDescriptor {
	return file_testsvc_test_v1_enums_proto_enumTypes[0].Descriptor()
}

func (PersonType) Type() protoreflect.EnumType {
	return &file_testsvc_test_v1_enums_proto_enumTypes[0]
}

func (x PersonType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PersonType.Descriptor instead.
func (PersonType) EnumDescriptor() ([]byte, []int) {
	return file_testsvc_test_v1_enums_proto_rawDescGZIP(), []int{0}
}

type NestedEnumTestMessage_PersonType int32

const (
	NestedEnumTestMessage_PERSON_TYPE_INVALID NestedEnumTestMessage_PersonType = 0
	NestedEnumTestMessage_PERSON_TYPE_ONE     NestedEnumTestMessage_PersonType = 1
	NestedEnumTestMessage_PERSON_TYPE_TWO     NestedEnumTestMessage_PersonType = 2
)

// Enum value maps for NestedEnumTestMessage_PersonType.
var (
	NestedEnumTestMessage_PersonType_name = map[int32]string{
		0: "PERSON_TYPE_INVALID",
		1: "PERSON_TYPE_ONE",
		2: "PERSON_TYPE_TWO",
	}
	NestedEnumTestMessage_PersonType_value = map[string]int32{
		"PERSON_TYPE_INVALID": 0,
		"PERSON_TYPE_ONE":     1,
		"PERSON_TYPE_TWO":     2,
	}
)

func (x NestedEnumTestMessage_PersonType) Enum() *NestedEnumTestMessage_PersonType {
	p := new(NestedEnumTestMessage_PersonType)
	*p = x
	return p
}

func (x NestedEnumTestMessage_PersonType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NestedEnumTestMessage_PersonType) Descriptor() protoreflect.EnumDescriptor {
	return file_testsvc_test_v1_enums_proto_enumTypes[1].Descriptor()
}

func (NestedEnumTestMessage_PersonType) Type() protoreflect.EnumType {
	return &file_testsvc_test_v1_enums_proto_enumTypes[1]
}

func (x NestedEnumTestMessage_PersonType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NestedEnumTestMessage_PersonType.Descriptor instead.
func (NestedEnumTestMessage_PersonType) EnumDescriptor() ([]byte, []int) {
	return file_testsvc_test_v1_enums_proto_rawDescGZIP(), []int{0, 0}
}

type NestedEnumTestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type NestedEnumTestMessage_PersonType `protobuf:"varint,1,opt,name=type,proto3,enum=testsvc.person.v1.NestedEnumTestMessage_PersonType" json:"type,omitempty"`
}

func (x *NestedEnumTestMessage) Reset() {
	*x = NestedEnumTestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testsvc_test_v1_enums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedEnumTestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedEnumTestMessage) ProtoMessage() {}

func (x *NestedEnumTestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_testsvc_test_v1_enums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedEnumTestMessage.ProtoReflect.Descriptor instead.
func (*NestedEnumTestMessage) Descriptor() ([]byte, []int) {
	return file_testsvc_test_v1_enums_proto_rawDescGZIP(), []int{0}
}

func (x *NestedEnumTestMessage) GetType() NestedEnumTestMessage_PersonType {
	if x != nil {
		return x.Type
	}
	return NestedEnumTestMessage_PERSON_TYPE_INVALID
}

type GetPersonTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId int32 `protobuf:"varint,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
}

func (x *GetPersonTypeRequest) Reset() {
	*x = GetPersonTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testsvc_test_v1_enums_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonTypeRequest) ProtoMessage() {}

func (x *GetPersonTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testsvc_test_v1_enums_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonTypeRequest.ProtoReflect.Descriptor instead.
func (*GetPersonTypeRequest) Descriptor() ([]byte, []int) {
	return file_testsvc_test_v1_enums_proto_rawDescGZIP(), []int{1}
}

func (x *GetPersonTypeRequest) GetPersonId() int32 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

type GetPersonTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type PersonType `protobuf:"varint,1,opt,name=type,proto3,enum=testsvc.person.v1.PersonType" json:"type,omitempty"`
}

func (x *GetPersonTypeResponse) Reset() {
	*x = GetPersonTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testsvc_test_v1_enums_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonTypeResponse) ProtoMessage() {}

func (x *GetPersonTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testsvc_test_v1_enums_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonTypeResponse.ProtoReflect.Descriptor instead.
func (*GetPersonTypeResponse) Descriptor() ([]byte, []int) {
	return file_testsvc_test_v1_enums_proto_rawDescGZIP(), []int{2}
}

func (x *GetPersonTypeResponse) GetType() PersonType {
	if x != nil {
		return x.Type
	}
	return PersonType_PERSON_TYPE_INVALID
}

var File_testsvc_test_v1_enums_proto protoreflect.FileDescriptor

var file_testsvc_test_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x73, 0x76, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x22, 0xb1, 0x01, 0x0a, 0x15, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x54,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x76, 0x63, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x4f, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x45,
	0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x57, 0x4f, 0x10, 0x02, 0x22, 0x33, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x4f, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f,
	0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x4e, 0x45, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x32, 0x75, 0x0a, 0x11, 0x54, 0x65, 0x73, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x76, 0x63, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testsvc_test_v1_enums_proto_rawDescOnce sync.Once
	file_testsvc_test_v1_enums_proto_rawDescData = file_testsvc_test_v1_enums_proto_rawDesc
)

func file_testsvc_test_v1_enums_proto_rawDescGZIP() []byte {
	file_testsvc_test_v1_enums_proto_rawDescOnce.Do(func() {
		file_testsvc_test_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_testsvc_test_v1_enums_proto_rawDescData)
	})
	return file_testsvc_test_v1_enums_proto_rawDescData
}

var file_testsvc_test_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_testsvc_test_v1_enums_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_testsvc_test_v1_enums_proto_goTypes = []interface{}{
	(PersonType)(0),                       // 0: testsvc.person.v1.PersonType
	(NestedEnumTestMessage_PersonType)(0), // 1: testsvc.person.v1.NestedEnumTestMessage.PersonType
	(*NestedEnumTestMessage)(nil),         // 2: testsvc.person.v1.NestedEnumTestMessage
	(*GetPersonTypeRequest)(nil),          // 3: testsvc.person.v1.GetPersonTypeRequest
	(*GetPersonTypeResponse)(nil),         // 4: testsvc.person.v1.GetPersonTypeResponse
}
var file_testsvc_test_v1_enums_proto_depIdxs = []int32{
	1, // 0: testsvc.person.v1.NestedEnumTestMessage.type:type_name -> testsvc.person.v1.NestedEnumTestMessage.PersonType
	0, // 1: testsvc.person.v1.GetPersonTypeResponse.type:type_name -> testsvc.person.v1.PersonType
	3, // 2: testsvc.person.v1.TestPersonService.GetPerson:input_type -> testsvc.person.v1.GetPersonTypeRequest
	4, // 3: testsvc.person.v1.TestPersonService.GetPerson:output_type -> testsvc.person.v1.GetPersonTypeResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_testsvc_test_v1_enums_proto_init() }
func file_testsvc_test_v1_enums_proto_init() {
	if File_testsvc_test_v1_enums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testsvc_test_v1_enums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedEnumTestMessage); i {
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
		file_testsvc_test_v1_enums_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonTypeRequest); i {
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
		file_testsvc_test_v1_enums_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonTypeResponse); i {
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
			RawDescriptor: file_testsvc_test_v1_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_testsvc_test_v1_enums_proto_goTypes,
		DependencyIndexes: file_testsvc_test_v1_enums_proto_depIdxs,
		EnumInfos:         file_testsvc_test_v1_enums_proto_enumTypes,
		MessageInfos:      file_testsvc_test_v1_enums_proto_msgTypes,
	}.Build()
	File_testsvc_test_v1_enums_proto = out.File
	file_testsvc_test_v1_enums_proto_rawDesc = nil
	file_testsvc_test_v1_enums_proto_goTypes = nil
	file_testsvc_test_v1_enums_proto_depIdxs = nil
}