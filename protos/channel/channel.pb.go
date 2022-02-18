// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: channel.proto

package v1

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

type GetListChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetListChannelRequest) Reset() {
	*x = GetListChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListChannelRequest) ProtoMessage() {}

func (x *GetListChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListChannelRequest.ProtoReflect.Descriptor instead.
func (*GetListChannelRequest) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{0}
}

type ChannelModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ChannelModel) Reset() {
	*x = ChannelModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelModel) ProtoMessage() {}

func (x *ChannelModel) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelModel.ProtoReflect.Descriptor instead.
func (*ChannelModel) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChannelModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChannelModel) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetListChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelItems []*ChannelModel `protobuf:"bytes,1,rep,name=channel_items,json=channelItems,proto3" json:"channel_items,omitempty"`
}

func (x *GetListChannelResponse) Reset() {
	*x = GetListChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListChannelResponse) ProtoMessage() {}

func (x *GetListChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListChannelResponse.ProtoReflect.Descriptor instead.
func (*GetListChannelResponse) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{2}
}

func (x *GetListChannelResponse) GetChannelItems() []*ChannelModel {
	if x != nil {
		return x.ChannelItems
	}
	return nil
}

var File_channel_proto protoreflect.FileDescriptor

var file_channel_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x76, 0x31, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x17, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x57, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0c, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x6b, 0x0a, 0x0e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x21, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_channel_proto_rawDescOnce sync.Once
	file_channel_proto_rawDescData = file_channel_proto_rawDesc
)

func file_channel_proto_rawDescGZIP() []byte {
	file_channel_proto_rawDescOnce.Do(func() {
		file_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_channel_proto_rawDescData)
	})
	return file_channel_proto_rawDescData
}

var file_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_channel_proto_goTypes = []interface{}{
	(*GetListChannelRequest)(nil),  // 0: v1.channel.GetListChannelRequest
	(*ChannelModel)(nil),           // 1: v1.channel.ChannelModel
	(*GetListChannelResponse)(nil), // 2: v1.channel.GetListChannelResponse
}
var file_channel_proto_depIdxs = []int32{
	1, // 0: v1.channel.GetListChannelResponse.channel_items:type_name -> v1.channel.ChannelModel
	0, // 1: v1.channel.ChannelService.GetListChannel:input_type -> v1.channel.GetListChannelRequest
	2, // 2: v1.channel.ChannelService.GetListChannel:output_type -> v1.channel.GetListChannelResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_channel_proto_init() }
func file_channel_proto_init() {
	if File_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListChannelRequest); i {
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
		file_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelModel); i {
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
		file_channel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListChannelResponse); i {
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
			RawDescriptor: file_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_channel_proto_goTypes,
		DependencyIndexes: file_channel_proto_depIdxs,
		MessageInfos:      file_channel_proto_msgTypes,
	}.Build()
	File_channel_proto = out.File
	file_channel_proto_rawDesc = nil
	file_channel_proto_goTypes = nil
	file_channel_proto_depIdxs = nil
}
