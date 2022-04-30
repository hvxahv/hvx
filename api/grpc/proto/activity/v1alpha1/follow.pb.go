// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: proto/activity/v1alpha1/follow.proto

package v1alpha

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

type CreateFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateFollowRequest) Reset() {
	*x = CreateFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_activity_v1alpha1_follow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFollowRequest) ProtoMessage() {}

func (x *CreateFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_activity_v1alpha1_follow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFollowRequest.ProtoReflect.Descriptor instead.
func (*CreateFollowRequest) Descriptor() ([]byte, []int) {
	return file_proto_activity_v1alpha1_follow_proto_rawDescGZIP(), []int{0}
}

type CreateFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateFollowResponse) Reset() {
	*x = CreateFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_activity_v1alpha1_follow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFollowResponse) ProtoMessage() {}

func (x *CreateFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_activity_v1alpha1_follow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFollowResponse.ProtoReflect.Descriptor instead.
func (*CreateFollowResponse) Descriptor() ([]byte, []int) {
	return file_proto_activity_v1alpha1_follow_proto_rawDescGZIP(), []int{1}
}

type GetFollowersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFollowersRequest) Reset() {
	*x = GetFollowersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_activity_v1alpha1_follow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersRequest) ProtoMessage() {}

func (x *GetFollowersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_activity_v1alpha1_follow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersRequest.ProtoReflect.Descriptor instead.
func (*GetFollowersRequest) Descriptor() ([]byte, []int) {
	return file_proto_activity_v1alpha1_follow_proto_rawDescGZIP(), []int{2}
}

type GetFollowersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFollowersResponse) Reset() {
	*x = GetFollowersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_activity_v1alpha1_follow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersResponse) ProtoMessage() {}

func (x *GetFollowersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_activity_v1alpha1_follow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersResponse.ProtoReflect.Descriptor instead.
func (*GetFollowersResponse) Descriptor() ([]byte, []int) {
	return file_proto_activity_v1alpha1_follow_proto_rawDescGZIP(), []int{3}
}

type GetFollowingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFollowingRequest) Reset() {
	*x = GetFollowingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_activity_v1alpha1_follow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingRequest) ProtoMessage() {}

func (x *GetFollowingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_activity_v1alpha1_follow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingRequest.ProtoReflect.Descriptor instead.
func (*GetFollowingRequest) Descriptor() ([]byte, []int) {
	return file_proto_activity_v1alpha1_follow_proto_rawDescGZIP(), []int{4}
}

type GetFollowingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFollowingResponse) Reset() {
	*x = GetFollowingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_activity_v1alpha1_follow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingResponse) ProtoMessage() {}

func (x *GetFollowingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_activity_v1alpha1_follow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingResponse.ProtoReflect.Descriptor instead.
func (*GetFollowingResponse) Descriptor() ([]byte, []int) {
	return file_proto_activity_v1alpha1_follow_proto_rawDescGZIP(), []int{5}
}

var File_proto_activity_v1alpha1_follow_proto protoreflect.FileDescriptor

var file_proto_activity_v1alpha1_follow_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc9, 0x02, 0x0a, 0x06, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x69, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x2a, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x69, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73,
	0x12, 0x2a, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x68,
	0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x2e, 0x68, 0x76,
	0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_activity_v1alpha1_follow_proto_rawDescOnce sync.Once
	file_proto_activity_v1alpha1_follow_proto_rawDescData = file_proto_activity_v1alpha1_follow_proto_rawDesc
)

func file_proto_activity_v1alpha1_follow_proto_rawDescGZIP() []byte {
	file_proto_activity_v1alpha1_follow_proto_rawDescOnce.Do(func() {
		file_proto_activity_v1alpha1_follow_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_activity_v1alpha1_follow_proto_rawDescData)
	})
	return file_proto_activity_v1alpha1_follow_proto_rawDescData
}

var file_proto_activity_v1alpha1_follow_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_activity_v1alpha1_follow_proto_goTypes = []interface{}{
	(*CreateFollowRequest)(nil),  // 0: hvxahv.v1alpha1.proto.CreateFollowRequest
	(*CreateFollowResponse)(nil), // 1: hvxahv.v1alpha1.proto.CreateFollowResponse
	(*GetFollowersRequest)(nil),  // 2: hvxahv.v1alpha1.proto.GetFollowersRequest
	(*GetFollowersResponse)(nil), // 3: hvxahv.v1alpha1.proto.GetFollowersResponse
	(*GetFollowingRequest)(nil),  // 4: hvxahv.v1alpha1.proto.GetFollowingRequest
	(*GetFollowingResponse)(nil), // 5: hvxahv.v1alpha1.proto.GetFollowingResponse
}
var file_proto_activity_v1alpha1_follow_proto_depIdxs = []int32{
	0, // 0: hvxahv.v1alpha1.proto.Follow.CreateFollow:input_type -> hvxahv.v1alpha1.proto.CreateFollowRequest
	2, // 1: hvxahv.v1alpha1.proto.Follow.GetFollowers:input_type -> hvxahv.v1alpha1.proto.GetFollowersRequest
	4, // 2: hvxahv.v1alpha1.proto.Follow.GetFollowing:input_type -> hvxahv.v1alpha1.proto.GetFollowingRequest
	1, // 3: hvxahv.v1alpha1.proto.Follow.CreateFollow:output_type -> hvxahv.v1alpha1.proto.CreateFollowResponse
	3, // 4: hvxahv.v1alpha1.proto.Follow.GetFollowers:output_type -> hvxahv.v1alpha1.proto.GetFollowersResponse
	5, // 5: hvxahv.v1alpha1.proto.Follow.GetFollowing:output_type -> hvxahv.v1alpha1.proto.GetFollowingResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_activity_v1alpha1_follow_proto_init() }
func file_proto_activity_v1alpha1_follow_proto_init() {
	if File_proto_activity_v1alpha1_follow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_activity_v1alpha1_follow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFollowRequest); i {
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
		file_proto_activity_v1alpha1_follow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFollowResponse); i {
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
		file_proto_activity_v1alpha1_follow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersRequest); i {
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
		file_proto_activity_v1alpha1_follow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersResponse); i {
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
		file_proto_activity_v1alpha1_follow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowingRequest); i {
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
		file_proto_activity_v1alpha1_follow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowingResponse); i {
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
			RawDescriptor: file_proto_activity_v1alpha1_follow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_activity_v1alpha1_follow_proto_goTypes,
		DependencyIndexes: file_proto_activity_v1alpha1_follow_proto_depIdxs,
		MessageInfos:      file_proto_activity_v1alpha1_follow_proto_msgTypes,
	}.Build()
	File_proto_activity_v1alpha1_follow_proto = out.File
	file_proto_activity_v1alpha1_follow_proto_rawDesc = nil
	file_proto_activity_v1alpha1_follow_proto_goTypes = nil
	file_proto_activity_v1alpha1_follow_proto_depIdxs = nil
}
