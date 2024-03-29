// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: pb/example.proto

package pb

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

type ContactType int32

const (
	ContactType_UNKNOWN ContactType = 0
	ContactType_EMAIL   ContactType = 1
	ContactType_PHONE   ContactType = 2
)

// Enum value maps for ContactType.
var (
	ContactType_name = map[int32]string{
		0: "UNKNOWN",
		1: "EMAIL",
		2: "PHONE",
	}
	ContactType_value = map[string]int32{
		"UNKNOWN": 0,
		"EMAIL":   1,
		"PHONE":   2,
	}
)

func (x ContactType) Enum() *ContactType {
	p := new(ContactType)
	*p = x
	return p
}

func (x ContactType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContactType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_example_proto_enumTypes[0].Descriptor()
}

func (ContactType) Type() protoreflect.EnumType {
	return &file_pb_example_proto_enumTypes[0]
}

func (x ContactType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContactType.Descriptor instead.
func (ContactType) EnumDescriptor() ([]byte, []int) {
	return file_pb_example_proto_rawDescGZIP(), []int{0}
}

type MySampleKafkaMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string     `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string     `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Contacts    []*Contact `protobuf:"bytes,3,rep,name=contacts,proto3" json:"contacts,omitempty"`
	LuckyNumber uint32     `protobuf:"varint,4,opt,name=lucky_number,json=luckyNumber,proto3" json:"lucky_number,omitempty"`
}

func (x *MySampleKafkaMessage) Reset() {
	*x = MySampleKafkaMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MySampleKafkaMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MySampleKafkaMessage) ProtoMessage() {}

func (x *MySampleKafkaMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pb_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MySampleKafkaMessage.ProtoReflect.Descriptor instead.
func (*MySampleKafkaMessage) Descriptor() ([]byte, []int) {
	return file_pb_example_proto_rawDescGZIP(), []int{0}
}

func (x *MySampleKafkaMessage) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *MySampleKafkaMessage) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *MySampleKafkaMessage) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

func (x *MySampleKafkaMessage) GetLuckyNumber() uint32 {
	if x != nil {
		return x.LuckyNumber
	}
	return 0
}

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactType ContactType `protobuf:"varint,1,opt,name=contact_type,json=contactType,proto3,enum=my.package.ContactType" json:"contact_type,omitempty"`
	Value       string      `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_pb_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_pb_example_proto_rawDescGZIP(), []int{1}
}

func (x *Contact) GetContactType() ContactType {
	if x != nil {
		return x.ContactType
	}
	return ContactType_UNKNOWN
}

func (x *Contact) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_pb_example_proto protoreflect.FileDescriptor

var file_pb_example_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x79, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0xa6,
	0x01, 0x0a, 0x14, 0x4d, 0x79, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x79, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6c, 0x75, 0x63, 0x6b,
	0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x5b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x79, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x2a, 0x30, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x50,
	0x48, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x6e, 0x79, 0x61, 0x6f, 0x2d, 0x6c, 0x65, 0x65, 0x2f,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2d, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_example_proto_rawDescOnce sync.Once
	file_pb_example_proto_rawDescData = file_pb_example_proto_rawDesc
)

func file_pb_example_proto_rawDescGZIP() []byte {
	file_pb_example_proto_rawDescOnce.Do(func() {
		file_pb_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_example_proto_rawDescData)
	})
	return file_pb_example_proto_rawDescData
}

var file_pb_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_example_proto_goTypes = []interface{}{
	(ContactType)(0),             // 0: my.package.ContactType
	(*MySampleKafkaMessage)(nil), // 1: my.package.MySampleKafkaMessage
	(*Contact)(nil),              // 2: my.package.Contact
}
var file_pb_example_proto_depIdxs = []int32{
	2, // 0: my.package.MySampleKafkaMessage.contacts:type_name -> my.package.Contact
	0, // 1: my.package.Contact.contact_type:type_name -> my.package.ContactType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_example_proto_init() }
func file_pb_example_proto_init() {
	if File_pb_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MySampleKafkaMessage); i {
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
		file_pb_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
			RawDescriptor: file_pb_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_example_proto_goTypes,
		DependencyIndexes: file_pb_example_proto_depIdxs,
		EnumInfos:         file_pb_example_proto_enumTypes,
		MessageInfos:      file_pb_example_proto_msgTypes,
	}.Build()
	File_pb_example_proto = out.File
	file_pb_example_proto_rawDesc = nil
	file_pb_example_proto_goTypes = nil
	file_pb_example_proto_depIdxs = nil
}
