// Copyright 2015 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: google/api/log.proto

package serviceconfig

import (
	label "google.golang.org/genproto/googleapis/api/label"
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

// A description of a log type. Example in YAML format:
//
//     - name: library.googleapis.com/activity_history
//       description: The history of borrowing and returning library items.
//       display_name: Activity
//       labels:
//       - key: /customer_id
//         description: Identifier of a library customer
type LogDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the log. It must be less than 512 characters long and can
	// include the following characters: upper- and lower-case alphanumeric
	// characters [A-Za-z0-9], and punctuation characters including
	// slash, underscore, hyphen, period [/_-.].
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The set of labels that are available to describe a specific log entry.
	// Runtime requests that contain labels not specified here are
	// considered invalid.
	Labels []*label.LabelDescriptor `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	// A human-readable description of this log. This information appears in
	// the documentation and can contain details.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The human-readable name for this log. This information appears on
	// the user interface and should be concise.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *LogDescriptor) Reset() {
	*x = LogDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogDescriptor) ProtoMessage() {}

func (x *LogDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogDescriptor.ProtoReflect.Descriptor instead.
func (*LogDescriptor) Descriptor() ([]byte, []int) {
	return file_google_api_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogDescriptor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LogDescriptor) GetLabels() []*label.LabelDescriptor {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *LogDescriptor) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LogDescriptor) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_google_api_log_proto protoreflect.FileDescriptor

var file_google_api_log_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x0d, 0x4c,
	0x6f, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x33, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x6a, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x08, 0x4c, 0x6f,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xa2,
	0x02, 0x04, 0x47, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_log_proto_rawDescOnce sync.Once
	file_google_api_log_proto_rawDescData = file_google_api_log_proto_rawDesc
)

func file_google_api_log_proto_rawDescGZIP() []byte {
	file_google_api_log_proto_rawDescOnce.Do(func() {
		file_google_api_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_log_proto_rawDescData)
	})
	return file_google_api_log_proto_rawDescData
}

var file_google_api_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_api_log_proto_goTypes = []interface{}{
	(*LogDescriptor)(nil),         // 0: google.api.LogDescriptor
	(*label.LabelDescriptor)(nil), // 1: google.api.LabelDescriptor
}
var file_google_api_log_proto_depIdxs = []int32{
	1, // 0: google.api.LogDescriptor.labels:type_name -> google.api.LabelDescriptor
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_api_log_proto_init() }
func file_google_api_log_proto_init() {
	if File_google_api_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogDescriptor); i {
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
			RawDescriptor: file_google_api_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_log_proto_goTypes,
		DependencyIndexes: file_google_api_log_proto_depIdxs,
		MessageInfos:      file_google_api_log_proto_msgTypes,
	}.Build()
	File_google_api_log_proto = out.File
	file_google_api_log_proto_rawDesc = nil
	file_google_api_log_proto_goTypes = nil
	file_google_api_log_proto_depIdxs = nil
}
