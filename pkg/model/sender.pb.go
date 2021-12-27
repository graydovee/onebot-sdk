// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: sender.proto

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

type SenderGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	发送者 QQ 号
	UserID int64 `protobuf:"varint,1,opt,name=UserID,json=user_id,proto3" json:"UserID,omitempty"`
	//昵称
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	//群名片／备注
	Card string `protobuf:"bytes,3,opt,name=card,proto3" json:"card,omitempty"`
	//性别，male 或 female 或 unknown
	Sex string `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty"`
	//	年龄
	Age int32 `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	//群名片／备注
	Area string `protobuf:"bytes,6,opt,name=area,proto3" json:"area,omitempty"`
	//成员等级
	Level string `protobuf:"bytes,7,opt,name=level,proto3" json:"level,omitempty"`
	//角色 owner/admin/member
	Role string `protobuf:"bytes,8,opt,name=role,proto3" json:"role,omitempty"`
	//头衔
	Title string `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *SenderGRPC) Reset() {
	*x = SenderGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sender_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SenderGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SenderGRPC) ProtoMessage() {}

func (x *SenderGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_sender_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SenderGRPC.ProtoReflect.Descriptor instead.
func (*SenderGRPC) Descriptor() ([]byte, []int) {
	return file_sender_proto_rawDescGZIP(), []int{0}
}

func (x *SenderGRPC) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *SenderGRPC) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SenderGRPC) GetCard() string {
	if x != nil {
		return x.Card
	}
	return ""
}

func (x *SenderGRPC) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *SenderGRPC) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *SenderGRPC) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *SenderGRPC) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *SenderGRPC) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *SenderGRPC) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_sender_proto protoreflect.FileDescriptor

var file_sender_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xcd, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x47, 0x52, 0x50, 0x43, 0x12, 0x17, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x61, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sender_proto_rawDescOnce sync.Once
	file_sender_proto_rawDescData = file_sender_proto_rawDesc
)

func file_sender_proto_rawDescGZIP() []byte {
	file_sender_proto_rawDescOnce.Do(func() {
		file_sender_proto_rawDescData = protoimpl.X.CompressGZIP(file_sender_proto_rawDescData)
	})
	return file_sender_proto_rawDescData
}

var file_sender_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sender_proto_goTypes = []interface{}{
	(*SenderGRPC)(nil), // 0: model.SenderGRPC
}
var file_sender_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sender_proto_init() }
func file_sender_proto_init() {
	if File_sender_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sender_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SenderGRPC); i {
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
			RawDescriptor: file_sender_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sender_proto_goTypes,
		DependencyIndexes: file_sender_proto_depIdxs,
		MessageInfos:      file_sender_proto_msgTypes,
	}.Build()
	File_sender_proto = out.File
	file_sender_proto_rawDesc = nil
	file_sender_proto_goTypes = nil
	file_sender_proto_depIdxs = nil
}
