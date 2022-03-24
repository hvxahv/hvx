// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: api/channel/v1alpha1/administrative.proto

package v1alpha1

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

type IsChannelAdministratorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *IsChannelAdministratorRequest) Reset() {
	*x = IsChannelAdministratorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsChannelAdministratorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsChannelAdministratorRequest) ProtoMessage() {}

func (x *IsChannelAdministratorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsChannelAdministratorRequest.ProtoReflect.Descriptor instead.
func (*IsChannelAdministratorRequest) Descriptor() ([]byte, []int) {
	return file_api_channel_v1alpha1_administrative_proto_rawDescGZIP(), []int{0}
}

func (x *IsChannelAdministratorRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *IsChannelAdministratorRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type IsChannelAdministratorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAdministrator bool `protobuf:"varint,1,opt,name=is_administrator,json=isAdministrator,proto3" json:"is_administrator,omitempty"`
}

func (x *IsChannelAdministratorResponse) Reset() {
	*x = IsChannelAdministratorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsChannelAdministratorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsChannelAdministratorResponse) ProtoMessage() {}

func (x *IsChannelAdministratorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsChannelAdministratorResponse.ProtoReflect.Descriptor instead.
func (*IsChannelAdministratorResponse) Descriptor() ([]byte, []int) {
	return file_api_channel_v1alpha1_administrative_proto_rawDescGZIP(), []int{1}
}

func (x *IsChannelAdministratorResponse) GetIsAdministrator() bool {
	if x != nil {
		return x.IsAdministrator
	}
	return false
}

type AddAdministratorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The channel id.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// ID of the channel owner.
	AdminAccountId string `protobuf:"bytes,2,opt,name=admin_account_id,json=adminAccountId,proto3" json:"admin_account_id,omitempty"`
	// Added administrator ID.
	AddAdminId string `protobuf:"bytes,3,opt,name=add_admin_id,json=addAdminId,proto3" json:"add_admin_id,omitempty"`
	IsOwner    bool   `protobuf:"varint,4,opt,name=is_owner,json=isOwner,proto3" json:"is_owner,omitempty"`
}

func (x *AddAdministratorRequest) Reset() {
	*x = AddAdministratorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAdministratorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAdministratorRequest) ProtoMessage() {}

func (x *AddAdministratorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAdministratorRequest.ProtoReflect.Descriptor instead.
func (*AddAdministratorRequest) Descriptor() ([]byte, []int) {
	return file_api_channel_v1alpha1_administrative_proto_rawDescGZIP(), []int{2}
}

func (x *AddAdministratorRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *AddAdministratorRequest) GetAdminAccountId() string {
	if x != nil {
		return x.AdminAccountId
	}
	return ""
}

func (x *AddAdministratorRequest) GetAddAdminId() string {
	if x != nil {
		return x.AddAdminId
	}
	return ""
}

func (x *AddAdministratorRequest) GetIsOwner() bool {
	if x != nil {
		return x.IsOwner
	}
	return false
}

type AddAdministratorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *AddAdministratorResponse) Reset() {
	*x = AddAdministratorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAdministratorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAdministratorResponse) ProtoMessage() {}

func (x *AddAdministratorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAdministratorResponse.ProtoReflect.Descriptor instead.
func (*AddAdministratorResponse) Descriptor() ([]byte, []int) {
	return file_api_channel_v1alpha1_administrative_proto_rawDescGZIP(), []int{3}
}

func (x *AddAdministratorResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AddAdministratorResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type RemoveAdministratorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId       string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	ChannelId     string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	RemoveAdminId string `protobuf:"bytes,3,opt,name=remove_admin_id,json=removeAdminId,proto3" json:"remove_admin_id,omitempty"`
}

func (x *RemoveAdministratorRequest) Reset() {
	*x = RemoveAdministratorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAdministratorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAdministratorRequest) ProtoMessage() {}

func (x *RemoveAdministratorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAdministratorRequest.ProtoReflect.Descriptor instead.
func (*RemoveAdministratorRequest) Descriptor() ([]byte, []int) {
	return file_api_channel_v1alpha1_administrative_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveAdministratorRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *RemoveAdministratorRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *RemoveAdministratorRequest) GetRemoveAdminId() string {
	if x != nil {
		return x.RemoveAdminId
	}
	return ""
}

type RemoveAdministratorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *RemoveAdministratorResponse) Reset() {
	*x = RemoveAdministratorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAdministratorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAdministratorResponse) ProtoMessage() {}

func (x *RemoveAdministratorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAdministratorResponse.ProtoReflect.Descriptor instead.
func (*RemoveAdministratorResponse) Descriptor() ([]byte, []int) {
	return file_api_channel_v1alpha1_administrative_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveAdministratorResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *RemoveAdministratorResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type GetAdministratorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *GetAdministratorsRequest) Reset() {
	*x = GetAdministratorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdministratorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdministratorsRequest) ProtoMessage() {}

func (x *GetAdministratorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdministratorsRequest.ProtoReflect.Descriptor instead.
func (*GetAdministratorsRequest) Descriptor() ([]byte, []int) {
	return file_api_channel_v1alpha1_administrative_proto_rawDescGZIP(), []int{6}
}

func (x *GetAdministratorsRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *GetAdministratorsRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type GetAdministratorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code           string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Administrators []string `protobuf:"bytes,2,rep,name=administrators,proto3" json:"administrators,omitempty"`
}

func (x *GetAdministratorsResponse) Reset() {
	*x = GetAdministratorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdministratorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdministratorsResponse) ProtoMessage() {}

func (x *GetAdministratorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_channel_v1alpha1_administrative_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdministratorsResponse.ProtoReflect.Descriptor instead.
func (*GetAdministratorsResponse) Descriptor() ([]byte, []int) {
	return file_api_channel_v1alpha1_administrative_proto_rawDescGZIP(), []int{7}
}

func (x *GetAdministratorsResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetAdministratorsResponse) GetAdministrators() []string {
	if x != nil {
		return x.Administrators
	}
	return nil
}

var File_api_channel_v1alpha1_administrative_proto protoreflect.FileDescriptor

var file_api_channel_v1alpha1_administrative_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x68, 0x76, 0x78,
	0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x1d, 0x49, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x4b, 0x0a, 0x1e, 0x49, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x9f,
	0x01, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x22, 0x44, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x7e, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x1b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x58, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x32, 0x92, 0x04, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x87, 0x01, 0x0a,
	0x16, 0x49, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x34, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e,
	0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x75, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2e, 0x2e, 0x68, 0x76, 0x78,
	0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x68, 0x76, 0x78,
	0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7e, 0x0a,
	0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x31, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x78, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x2f, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2f, 0x68, 0x76, 0x78,
	0x61, 0x68, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_channel_v1alpha1_administrative_proto_rawDescOnce sync.Once
	file_api_channel_v1alpha1_administrative_proto_rawDescData = file_api_channel_v1alpha1_administrative_proto_rawDesc
)

func file_api_channel_v1alpha1_administrative_proto_rawDescGZIP() []byte {
	file_api_channel_v1alpha1_administrative_proto_rawDescOnce.Do(func() {
		file_api_channel_v1alpha1_administrative_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_channel_v1alpha1_administrative_proto_rawDescData)
	})
	return file_api_channel_v1alpha1_administrative_proto_rawDescData
}

var file_api_channel_v1alpha1_administrative_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_channel_v1alpha1_administrative_proto_goTypes = []interface{}{
	(*IsChannelAdministratorRequest)(nil),  // 0: hvxahv.v1alpha1.proto.IsChannelAdministratorRequest
	(*IsChannelAdministratorResponse)(nil), // 1: hvxahv.v1alpha1.proto.IsChannelAdministratorResponse
	(*AddAdministratorRequest)(nil),        // 2: hvxahv.v1alpha1.proto.AddAdministratorRequest
	(*AddAdministratorResponse)(nil),       // 3: hvxahv.v1alpha1.proto.AddAdministratorResponse
	(*RemoveAdministratorRequest)(nil),     // 4: hvxahv.v1alpha1.proto.RemoveAdministratorRequest
	(*RemoveAdministratorResponse)(nil),    // 5: hvxahv.v1alpha1.proto.RemoveAdministratorResponse
	(*GetAdministratorsRequest)(nil),       // 6: hvxahv.v1alpha1.proto.GetAdministratorsRequest
	(*GetAdministratorsResponse)(nil),      // 7: hvxahv.v1alpha1.proto.GetAdministratorsResponse
}
var file_api_channel_v1alpha1_administrative_proto_depIdxs = []int32{
	0, // 0: hvxahv.v1alpha1.proto.AdministrativeService.IsChannelAdministrator:input_type -> hvxahv.v1alpha1.proto.IsChannelAdministratorRequest
	2, // 1: hvxahv.v1alpha1.proto.AdministrativeService.AddAdministrator:input_type -> hvxahv.v1alpha1.proto.AddAdministratorRequest
	4, // 2: hvxahv.v1alpha1.proto.AdministrativeService.RemoveAdministrator:input_type -> hvxahv.v1alpha1.proto.RemoveAdministratorRequest
	6, // 3: hvxahv.v1alpha1.proto.AdministrativeService.GetAdministrators:input_type -> hvxahv.v1alpha1.proto.GetAdministratorsRequest
	1, // 4: hvxahv.v1alpha1.proto.AdministrativeService.IsChannelAdministrator:output_type -> hvxahv.v1alpha1.proto.IsChannelAdministratorResponse
	3, // 5: hvxahv.v1alpha1.proto.AdministrativeService.AddAdministrator:output_type -> hvxahv.v1alpha1.proto.AddAdministratorResponse
	5, // 6: hvxahv.v1alpha1.proto.AdministrativeService.RemoveAdministrator:output_type -> hvxahv.v1alpha1.proto.RemoveAdministratorResponse
	7, // 7: hvxahv.v1alpha1.proto.AdministrativeService.GetAdministrators:output_type -> hvxahv.v1alpha1.proto.GetAdministratorsResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_channel_v1alpha1_administrative_proto_init() }
func file_api_channel_v1alpha1_administrative_proto_init() {
	if File_api_channel_v1alpha1_administrative_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_channel_v1alpha1_administrative_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsChannelAdministratorRequest); i {
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
		file_api_channel_v1alpha1_administrative_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsChannelAdministratorResponse); i {
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
		file_api_channel_v1alpha1_administrative_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAdministratorRequest); i {
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
		file_api_channel_v1alpha1_administrative_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAdministratorResponse); i {
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
		file_api_channel_v1alpha1_administrative_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAdministratorRequest); i {
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
		file_api_channel_v1alpha1_administrative_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAdministratorResponse); i {
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
		file_api_channel_v1alpha1_administrative_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdministratorsRequest); i {
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
		file_api_channel_v1alpha1_administrative_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdministratorsResponse); i {
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
			RawDescriptor: file_api_channel_v1alpha1_administrative_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_channel_v1alpha1_administrative_proto_goTypes,
		DependencyIndexes: file_api_channel_v1alpha1_administrative_proto_depIdxs,
		MessageInfos:      file_api_channel_v1alpha1_administrative_proto_msgTypes,
	}.Build()
	File_api_channel_v1alpha1_administrative_proto = out.File
	file_api_channel_v1alpha1_administrative_proto_rawDesc = nil
	file_api_channel_v1alpha1_administrative_proto_goTypes = nil
	file_api_channel_v1alpha1_administrative_proto_depIdxs = nil
}
