// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: api/accounts/v1alpha1/accounts.proto

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

type DeleteData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	ActorID  uint64 `protobuf:"varint,2,opt,name=actorID,proto3" json:"actorID,omitempty"`
}

func (x *DeleteData) Reset() {
	*x = DeleteData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteData) ProtoMessage() {}

func (x *DeleteData) ProtoReflect() protoreflect.Message {
	mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteData.ProtoReflect.Descriptor instead.
func (*DeleteData) Descriptor() ([]byte, []int) {
	return file_api_accounts_v1alpha1_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DeleteData) GetActorID() uint64 {
	if x != nil {
		return x.ActorID
	}
	return 0
}

type UpdatePasswordData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UpdatePasswordData) Reset() {
	*x = UpdatePasswordData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasswordData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswordData) ProtoMessage() {}

func (x *UpdatePasswordData) ProtoReflect() protoreflect.Message {
	mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswordData.ProtoReflect.Descriptor instead.
func (*UpdatePasswordData) Descriptor() ([]byte, []int) {
	return file_api_accounts_v1alpha1_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePasswordData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdatePasswordData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateUsernameData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username       string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	TargetUsername string `protobuf:"bytes,2,opt,name=targetUsername,proto3" json:"targetUsername,omitempty"`
	ActorID        uint64 `protobuf:"varint,3,opt,name=actorID,proto3" json:"actorID,omitempty"`
}

func (x *UpdateUsernameData) Reset() {
	*x = UpdateUsernameData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUsernameData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUsernameData) ProtoMessage() {}

func (x *UpdateUsernameData) ProtoReflect() protoreflect.Message {
	mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUsernameData.ProtoReflect.Descriptor instead.
func (*UpdateUsernameData) Descriptor() ([]byte, []int) {
	return file_api_accounts_v1alpha1_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUsernameData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateUsernameData) GetTargetUsername() string {
	if x != nil {
		return x.TargetUsername
	}
	return ""
}

func (x *UpdateUsernameData) GetActorID() uint64 {
	if x != nil {
		return x.ActorID
	}
	return 0
}

type ActorID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorID uint64 `protobuf:"varint,1,opt,name=actorID,proto3" json:"actorID,omitempty"`
}

func (x *ActorID) Reset() {
	*x = ActorID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorID) ProtoMessage() {}

func (x *ActorID) ProtoReflect() protoreflect.Message {
	mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorID.ProtoReflect.Descriptor instead.
func (*ActorID) Descriptor() ([]byte, []int) {
	return file_api_accounts_v1alpha1_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *ActorID) GetActorID() uint64 {
	if x != nil {
		return x.ActorID
	}
	return 0
}

type ActorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PreferredUsername string `protobuf:"bytes,2,opt,name=preferredUsername,proto3" json:"preferredUsername,omitempty"`
	Domain            string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Avatar            string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Name              string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Summary           string `protobuf:"bytes,6,opt,name=summary,proto3" json:"summary,omitempty"`
	Inbox             string `protobuf:"bytes,7,opt,name=inbox,proto3" json:"inbox,omitempty"`
	PublicKey         string `protobuf:"bytes,8,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	MatrixId          string `protobuf:"bytes,9,opt,name=matrixId,proto3" json:"matrixId,omitempty"`
	MatrixToken       string `protobuf:"bytes,10,opt,name=matrixToken,proto3" json:"matrixToken,omitempty"`
	ActorType         string `protobuf:"bytes,11,opt,name=actorType,proto3" json:"actorType,omitempty"`
}

func (x *ActorData) Reset() {
	*x = ActorData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorData) ProtoMessage() {}

func (x *ActorData) ProtoReflect() protoreflect.Message {
	mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorData.ProtoReflect.Descriptor instead.
func (*ActorData) Descriptor() ([]byte, []int) {
	return file_api_accounts_v1alpha1_accounts_proto_rawDescGZIP(), []int{4}
}

func (x *ActorData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ActorData) GetPreferredUsername() string {
	if x != nil {
		return x.PreferredUsername
	}
	return ""
}

func (x *ActorData) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ActorData) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *ActorData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ActorData) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *ActorData) GetInbox() string {
	if x != nil {
		return x.Inbox
	}
	return ""
}

func (x *ActorData) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *ActorData) GetMatrixId() string {
	if x != nil {
		return x.MatrixId
	}
	return ""
}

func (x *ActorData) GetMatrixToken() string {
	if x != nil {
		return x.MatrixToken
	}
	return ""
}

func (x *ActorData) GetActorType() string {
	if x != nil {
		return x.ActorType
	}
	return ""
}

// NewAccountData Data needed to create a new account.
type CreateAccountData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Mail     string `protobuf:"bytes,3,opt,name=mail,proto3" json:"mail,omitempty"`
}

func (x *CreateAccountData) Reset() {
	*x = CreateAccountData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountData) ProtoMessage() {}

func (x *CreateAccountData) ProtoReflect() protoreflect.Message {
	mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountData.ProtoReflect.Descriptor instead.
func (*CreateAccountData) Descriptor() ([]byte, []int) {
	return file_api_accounts_v1alpha1_accounts_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAccountData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateAccountData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateAccountData) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

// AccountsReply The return value, code and message of the Accounts operation.
type AccountsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AccountsReply) Reset() {
	*x = AccountsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsReply) ProtoMessage() {}

func (x *AccountsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountsReply.ProtoReflect.Descriptor instead.
func (*AccountsReply) Descriptor() ([]byte, []int) {
	return file_api_accounts_v1alpha1_accounts_proto_rawDescGZIP(), []int{6}
}

func (x *AccountsReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AccountsReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// AccountUsername Functions operated by username.
type AccountUsername struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AccountUsername) Reset() {
	*x = AccountUsername{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountUsername) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountUsername) ProtoMessage() {}

func (x *AccountUsername) ProtoReflect() protoreflect.Message {
	mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountUsername.ProtoReflect.Descriptor instead.
func (*AccountUsername) Descriptor() ([]byte, []int) {
	return file_api_accounts_v1alpha1_accounts_proto_rawDescGZIP(), []int{7}
}

func (x *AccountUsername) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// AccountsData List of all data of the account service
type AccountData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Mail      string `protobuf:"bytes,3,opt,name=mail,proto3" json:"mail,omitempty"`
	ActorId   uint64 `protobuf:"varint,5,opt,name=actorId,proto3" json:"actorId,omitempty"`
	IsPrivate bool   `protobuf:"varint,6,opt,name=isPrivate,proto3" json:"isPrivate,omitempty"`
}

func (x *AccountData) Reset() {
	*x = AccountData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountData) ProtoMessage() {}

func (x *AccountData) ProtoReflect() protoreflect.Message {
	mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountData.ProtoReflect.Descriptor instead.
func (*AccountData) Descriptor() ([]byte, []int) {
	return file_api_accounts_v1alpha1_accounts_proto_rawDescGZIP(), []int{8}
}

func (x *AccountData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AccountData) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *AccountData) GetActorId() uint64 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *AccountData) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

// AuthReply Account login, use email and password.
type AuthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mail     string `protobuf:"bytes,1,opt,name=mail,proto3" json:"mail,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthData) Reset() {
	*x = AuthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthData) ProtoMessage() {}

func (x *AuthData) ProtoReflect() protoreflect.Message {
	mi := &file_api_accounts_v1alpha1_accounts_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthData.ProtoReflect.Descriptor instead.
func (*AuthData) Descriptor() ([]byte, []int) {
	return file_api_accounts_v1alpha1_accounts_proto_rawDescGZIP(), []int{9}
}

func (x *AuthData) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *AuthData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_api_accounts_v1alpha1_accounts_proto protoreflect.FileDescriptor

var file_api_accounts_v1alpha1_accounts_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49,
	0x44, 0x22, 0x4c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x72, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x44, 0x22, 0x23, 0x0a, 0x07, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x22, 0xb7, 0x02, 0x0a, 0x09, 0x41, 0x63, 0x74,
	0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x5f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x3d, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x2d, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x85, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x0a, 0x08, 0x41, 0x75, 0x74,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0xd5, 0x06, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x5a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x68,
	0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x24, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x54,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68,
	0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x24, 0x2e, 0x68,
	0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x24, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x2e, 0x68, 0x76,
	0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x24, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x53,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68,
	0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x24, 0x2e, 0x68, 0x76,
	0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x2e,
	0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x22, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x1b, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x2e, 0x68, 0x76, 0x78,
	0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x1a, 0x20, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x63,
	0x74, 0x6f, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x1a, 0x20, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x24, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x73, 0x69,
	0x73, 0x6d, 0x2f, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_accounts_v1alpha1_accounts_proto_rawDescOnce sync.Once
	file_api_accounts_v1alpha1_accounts_proto_rawDescData = file_api_accounts_v1alpha1_accounts_proto_rawDesc
)

func file_api_accounts_v1alpha1_accounts_proto_rawDescGZIP() []byte {
	file_api_accounts_v1alpha1_accounts_proto_rawDescOnce.Do(func() {
		file_api_accounts_v1alpha1_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_accounts_v1alpha1_accounts_proto_rawDescData)
	})
	return file_api_accounts_v1alpha1_accounts_proto_rawDescData
}

var file_api_accounts_v1alpha1_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_accounts_v1alpha1_accounts_proto_goTypes = []interface{}{
	(*DeleteData)(nil),         // 0: hvxahv.v1alpha1.proto.DeleteData
	(*UpdatePasswordData)(nil), // 1: hvxahv.v1alpha1.proto.UpdatePasswordData
	(*UpdateUsernameData)(nil), // 2: hvxahv.v1alpha1.proto.UpdateUsernameData
	(*ActorID)(nil),            // 3: hvxahv.v1alpha1.proto.ActorID
	(*ActorData)(nil),          // 4: hvxahv.v1alpha1.proto.ActorData
	(*CreateAccountData)(nil),  // 5: hvxahv.v1alpha1.proto.CreateAccountData
	(*AccountsReply)(nil),      // 6: hvxahv.v1alpha1.proto.AccountsReply
	(*AccountUsername)(nil),    // 7: hvxahv.v1alpha1.proto.AccountUsername
	(*AccountData)(nil),        // 8: hvxahv.v1alpha1.proto.AccountData
	(*AuthData)(nil),           // 9: hvxahv.v1alpha1.proto.AuthData
}
var file_api_accounts_v1alpha1_accounts_proto_depIdxs = []int32{
	5, // 0: hvxahv.v1alpha1.proto.Accounts.Create:input_type -> hvxahv.v1alpha1.proto.CreateAccountData
	8, // 1: hvxahv.v1alpha1.proto.Accounts.Update:input_type -> hvxahv.v1alpha1.proto.AccountData
	1, // 2: hvxahv.v1alpha1.proto.Accounts.UpdatePassword:input_type -> hvxahv.v1alpha1.proto.UpdatePasswordData
	2, // 3: hvxahv.v1alpha1.proto.Accounts.UpdateUsername:input_type -> hvxahv.v1alpha1.proto.UpdateUsernameData
	0, // 4: hvxahv.v1alpha1.proto.Accounts.Delete:input_type -> hvxahv.v1alpha1.proto.DeleteData
	7, // 5: hvxahv.v1alpha1.proto.Accounts.FindAccountsByUsername:input_type -> hvxahv.v1alpha1.proto.AccountUsername
	7, // 6: hvxahv.v1alpha1.proto.Accounts.FindActorByAccountsUsername:input_type -> hvxahv.v1alpha1.proto.AccountUsername
	3, // 7: hvxahv.v1alpha1.proto.Accounts.FindActorByID:input_type -> hvxahv.v1alpha1.proto.ActorID
	9, // 8: hvxahv.v1alpha1.proto.Accounts.Login:input_type -> hvxahv.v1alpha1.proto.AuthData
	6, // 9: hvxahv.v1alpha1.proto.Accounts.Create:output_type -> hvxahv.v1alpha1.proto.AccountsReply
	6, // 10: hvxahv.v1alpha1.proto.Accounts.Update:output_type -> hvxahv.v1alpha1.proto.AccountsReply
	6, // 11: hvxahv.v1alpha1.proto.Accounts.UpdatePassword:output_type -> hvxahv.v1alpha1.proto.AccountsReply
	6, // 12: hvxahv.v1alpha1.proto.Accounts.UpdateUsername:output_type -> hvxahv.v1alpha1.proto.AccountsReply
	6, // 13: hvxahv.v1alpha1.proto.Accounts.Delete:output_type -> hvxahv.v1alpha1.proto.AccountsReply
	8, // 14: hvxahv.v1alpha1.proto.Accounts.FindAccountsByUsername:output_type -> hvxahv.v1alpha1.proto.AccountData
	4, // 15: hvxahv.v1alpha1.proto.Accounts.FindActorByAccountsUsername:output_type -> hvxahv.v1alpha1.proto.ActorData
	4, // 16: hvxahv.v1alpha1.proto.Accounts.FindActorByID:output_type -> hvxahv.v1alpha1.proto.ActorData
	6, // 17: hvxahv.v1alpha1.proto.Accounts.Login:output_type -> hvxahv.v1alpha1.proto.AccountsReply
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_accounts_v1alpha1_accounts_proto_init() }
func file_api_accounts_v1alpha1_accounts_proto_init() {
	if File_api_accounts_v1alpha1_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_accounts_v1alpha1_accounts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteData); i {
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
		file_api_accounts_v1alpha1_accounts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePasswordData); i {
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
		file_api_accounts_v1alpha1_accounts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUsernameData); i {
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
		file_api_accounts_v1alpha1_accounts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorID); i {
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
		file_api_accounts_v1alpha1_accounts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorData); i {
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
		file_api_accounts_v1alpha1_accounts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountData); i {
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
		file_api_accounts_v1alpha1_accounts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountsReply); i {
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
		file_api_accounts_v1alpha1_accounts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountUsername); i {
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
		file_api_accounts_v1alpha1_accounts_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountData); i {
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
		file_api_accounts_v1alpha1_accounts_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthData); i {
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
			RawDescriptor: file_api_accounts_v1alpha1_accounts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_accounts_v1alpha1_accounts_proto_goTypes,
		DependencyIndexes: file_api_accounts_v1alpha1_accounts_proto_depIdxs,
		MessageInfos:      file_api_accounts_v1alpha1_accounts_proto_msgTypes,
	}.Build()
	File_api_accounts_v1alpha1_accounts_proto = out.File
	file_api_accounts_v1alpha1_accounts_proto_rawDesc = nil
	file_api_accounts_v1alpha1_accounts_proto_goTypes = nil
	file_api_accounts_v1alpha1_accounts_proto_depIdxs = nil
}
