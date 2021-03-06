// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.9.0
// source: model.proto

//protoc --go_out=./ model.proto   执行
//指定所在包

package model

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

type Sex int32

const (
	Sex_UNKNOWN Sex = 0
	Sex_MAN     Sex = 1
	Sex_WOMAN   Sex = 2
)

// Enum value maps for Sex.
var (
	Sex_name = map[int32]string{
		0: "UNKNOWN",
		1: "MAN",
		2: "WOMAN",
	}
	Sex_value = map[string]int32{
		"UNKNOWN": 0,
		"MAN":     1,
		"WOMAN":   2,
	}
)

func (x Sex) Enum() *Sex {
	p := new(Sex)
	*p = x
	return p
}

func (x Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_model_proto_enumTypes[0].Descriptor()
}

func (Sex) Type() protoreflect.EnumType {
	return &file_model_proto_enumTypes[0]
}

func (x Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sex.Descriptor instead.
func (Sex) EnumDescriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Phone   string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Salt    string `protobuf:"bytes,3,opt,name=salt,proto3" json:"salt,omitempty"`
	CodePwd string `protobuf:"bytes,4,opt,name=codePwd,proto3" json:"codePwd,omitempty"`
	Name    string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Sex     Sex    `protobuf:"varint,6,opt,name=sex,proto3,enum=protobuf.Sex" json:"sex,omitempty"`
	Uptime  string `protobuf:"bytes,7,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Text    string `protobuf:"bytes,8,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Student) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Student) GetSalt() string {
	if x != nil {
		return x.Salt
	}
	return ""
}

func (x *Student) GetCodePwd() string {
	if x != nil {
		return x.CodePwd
	}
	return ""
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetSex() Sex {
	if x != nil {
		return x.Sex
	}
	return Sex_UNKNOWN
}

func (x *Student) GetUptime() string {
	if x != nil {
		return x.Uptime
	}
	return ""
}

func (x *Student) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Uptime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Uptime string `protobuf:"bytes,2,opt,name=uptime,proto3" json:"uptime,omitempty"`
}

func (x *Uptime) Reset() {
	*x = Uptime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uptime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uptime) ProtoMessage() {}

func (x *Uptime) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uptime.ProtoReflect.Descriptor instead.
func (*Uptime) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{1}
}

func (x *Uptime) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Uptime) GetUptime() string {
	if x != nil {
		return x.Uptime
	}
	return ""
}

var File_model_proto protoreflect.FileDescriptor

var file_model_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0xc2, 0x01, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x61, 0x6c,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x77, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x77, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x78, 0x52, 0x03, 0x73, 0x65, 0x78,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x34, 0x0a, 0x06,
	0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x2a, 0x26, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x4e, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x57, 0x4f, 0x4d, 0x41, 0x4e, 0x10, 0x02, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_model_proto_rawDescOnce sync.Once
	file_model_proto_rawDescData = file_model_proto_rawDesc
)

func file_model_proto_rawDescGZIP() []byte {
	file_model_proto_rawDescOnce.Do(func() {
		file_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_proto_rawDescData)
	})
	return file_model_proto_rawDescData
}

var file_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_proto_goTypes = []interface{}{
	(Sex)(0),        // 0: protobuf.Sex
	(*Student)(nil), // 1: protobuf.Student
	(*Uptime)(nil),  // 2: protobuf.uptime
}
var file_model_proto_depIdxs = []int32{
	0, // 0: protobuf.Student.sex:type_name -> protobuf.Sex
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_proto_init() }
func file_model_proto_init() {
	if File_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
		file_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uptime); i {
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
			RawDescriptor: file_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_proto_goTypes,
		DependencyIndexes: file_model_proto_depIdxs,
		EnumInfos:         file_model_proto_enumTypes,
		MessageInfos:      file_model_proto_msgTypes,
	}.Build()
	File_model_proto = out.File
	file_model_proto_rawDesc = nil
	file_model_proto_goTypes = nil
	file_model_proto_depIdxs = nil
}
