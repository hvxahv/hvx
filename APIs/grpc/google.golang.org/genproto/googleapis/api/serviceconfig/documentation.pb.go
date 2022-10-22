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
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: google/api/documentation.proto

package serviceconfig

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

// `Documentation` provides the information for describing a service.
//
// Example:
// <pre><code>documentation:
//
//	summary: >
//	  The Google Calendar API gives access
//	  to most calendar features.
//	pages:
//	- name: Overview
//	  content: &#40;== include google/foo/overview.md ==&#41;
//	- name: Tutorial
//	  content: &#40;== include google/foo/tutorial.md ==&#41;
//	  subpages;
//	  - name: Java
//	    content: &#40;== include google/foo/tutorial_java.md ==&#41;
//	rules:
//	- selector: google.calendar.Calendar.Get
//	  description: >
//	    ...
//	- selector: google.calendar.Calendar.Put
//	  description: >
//	    ...
//
// </code></pre>
// Documentation is provided in markdown syntax. In addition to
// standard markdown features, definition lists, tables and fenced
// code blocks are supported. Section headers can be provided and are
// interpreted relative to the section nesting of the context where
// a documentation fragment is embedded.
//
// Documentation from the IDL is merged with documentation defined
// via the config at normalization time, where documentation provided
// by config rules overrides IDL provided.
//
// A number of constructs specific to the API platform are supported
// in documentation text.
//
// In order to reference a proto element, the following
// notation can be used:
// <pre><code>&#91;fully.qualified.proto.name]&#91;]</code></pre>
// To override the display text used for the link, this can be used:
// <pre><code>&#91;display text]&#91;fully.qualified.proto.name]</code></pre>
// Text can be excluded from doc using the following notation:
// <pre><code>&#40;-- internal comment --&#41;</code></pre>
//
// A few directives are available in documentation. Note that
// directives must appear on a single line to be properly
// identified. The `include` directive includes a markdown file from
// an external source:
// <pre><code>&#40;== include path/to/file ==&#41;</code></pre>
// The `resource_for` directive marks a message to be the resource of
// a collection in REST view. If it is not specified, tools attempt
// to infer the resource from the operations in a collection:
// <pre><code>&#40;== resource_for v1.shelves.books ==&#41;</code></pre>
// The directive `suppress_warning` does not directly affect documentation
// and is documented together with service config validation.
type Documentation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A short summary of what the service does. Can only be provided by
	// plain text.
	Summary string `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	// The top level pages for the documentation set.
	Pages []*Page `protobuf:"bytes,5,rep,name=pages,proto3" json:"pages,omitempty"`
	// A list of documentation rules that apply to individual API elements.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*DocumentationRule `protobuf:"bytes,3,rep,name=rules,proto3" json:"rules,omitempty"`
	// The URL to the root of documentation.
	DocumentationRootUrl string `protobuf:"bytes,4,opt,name=documentation_root_url,json=documentationRootUrl,proto3" json:"documentation_root_url,omitempty"`
	// Specifies the service root url if the default one (the service name
	// from the yaml file) is not suitable. This can be seen in any fully
	// specified service urls as well as sections that show a base that other
	// urls are relative to.
	ServiceRootUrl string `protobuf:"bytes,6,opt,name=service_root_url,json=serviceRootUrl,proto3" json:"service_root_url,omitempty"`
	// Declares a single overview page. For example:
	// <pre><code>documentation:
	//
	//	summary: ...
	//	overview: &#40;== include overview.md ==&#41;
	//
	// </code></pre>
	// This is a shortcut for the following declaration (using pages style):
	// <pre><code>documentation:
	//
	//	summary: ...
	//	pages:
	//	- name: Overview
	//	  content: &#40;== include overview.md ==&#41;
	//
	// </code></pre>
	// Note: you cannot specify both `overview` field and `pages` field.
	Overview string `protobuf:"bytes,2,opt,name=overview,proto3" json:"overview,omitempty"`
}

func (x *Documentation) Reset() {
	*x = Documentation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_documentation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Documentation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Documentation) ProtoMessage() {}

func (x *Documentation) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_documentation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Documentation.ProtoReflect.Descriptor instead.
func (*Documentation) Descriptor() ([]byte, []int) {
	return file_google_api_documentation_proto_rawDescGZIP(), []int{0}
}

func (x *Documentation) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Documentation) GetPages() []*Page {
	if x != nil {
		return x.Pages
	}
	return nil
}

func (x *Documentation) GetRules() []*DocumentationRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *Documentation) GetDocumentationRootUrl() string {
	if x != nil {
		return x.DocumentationRootUrl
	}
	return ""
}

func (x *Documentation) GetServiceRootUrl() string {
	if x != nil {
		return x.ServiceRootUrl
	}
	return ""
}

func (x *Documentation) GetOverview() string {
	if x != nil {
		return x.Overview
	}
	return ""
}

// A documentation rule provides information about individual API elements.
type DocumentationRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The selector is a comma-separated list of patterns. Each pattern is a
	// qualified name of the element which may end in "*", indicating a wildcard.
	// Wildcards are only allowed at the end and for a whole component of the
	// qualified name, i.e. "foo.*" is ok, but not "foo.b*" or "foo.*.bar". A
	// wildcard will match one or more components. To specify a default for all
	// applicable elements, the whole pattern "*" is used.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Description of the selected API(s).
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Deprecation description of the selected element(s). It can be provided if
	// an element is marked as `deprecated`.
	DeprecationDescription string `protobuf:"bytes,3,opt,name=deprecation_description,json=deprecationDescription,proto3" json:"deprecation_description,omitempty"`
}

func (x *DocumentationRule) Reset() {
	*x = DocumentationRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_documentation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentationRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentationRule) ProtoMessage() {}

func (x *DocumentationRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_documentation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentationRule.ProtoReflect.Descriptor instead.
func (*DocumentationRule) Descriptor() ([]byte, []int) {
	return file_google_api_documentation_proto_rawDescGZIP(), []int{1}
}

func (x *DocumentationRule) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *DocumentationRule) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DocumentationRule) GetDeprecationDescription() string {
	if x != nil {
		return x.DeprecationDescription
	}
	return ""
}

// Represents a documentation page. A page can contain subpages to represent
// nested documentation set structure.
type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the page. It will be used as an identity of the page to
	// generate URI of the page, text of the link to this page in navigation,
	// etc. The full page name (start from the root page name to this page
	// concatenated with `.`) can be used as reference to the page in your
	// documentation. For example:
	// <pre><code>pages:
	//   - name: Tutorial
	//     content: &#40;== include tutorial.md ==&#41;
	//     subpages:
	//   - name: Java
	//     content: &#40;== include tutorial_java.md ==&#41;
	//
	// </code></pre>
	// You can reference `Java` page using Markdown reference link syntax:
	// `[Java][Tutorial.Java]`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The Markdown content of the page. You can use <code>&#40;== include {path}
	// ==&#41;</code> to include content from a Markdown file.
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	// Subpages of this page. The order of subpages specified here will be
	// honored in the generated docset.
	Subpages []*Page `protobuf:"bytes,3,rep,name=subpages,proto3" json:"subpages,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_documentation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_documentation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_google_api_documentation_proto_rawDescGZIP(), []int{2}
}

func (x *Page) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Page) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Page) GetSubpages() []*Page {
	if x != nil {
		return x.Subpages
	}
	return nil
}

var File_google_api_documentation_proto protoreflect.FileDescriptor

var file_google_api_documentation_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x82, 0x02, 0x0a,
	0x0d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x33, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6f, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x22, 0x8a, 0x01, 0x0a, 0x11, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x17, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x62,
	0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x08, 0x73, 0x75, 0x62, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x42, 0x74, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x42, 0x12, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0xa2, 0x02, 0x04, 0x47, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_documentation_proto_rawDescOnce sync.Once
	file_google_api_documentation_proto_rawDescData = file_google_api_documentation_proto_rawDesc
)

func file_google_api_documentation_proto_rawDescGZIP() []byte {
	file_google_api_documentation_proto_rawDescOnce.Do(func() {
		file_google_api_documentation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_documentation_proto_rawDescData)
	})
	return file_google_api_documentation_proto_rawDescData
}

var file_google_api_documentation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_api_documentation_proto_goTypes = []interface{}{
	(*Documentation)(nil),     // 0: google.api.Documentation
	(*DocumentationRule)(nil), // 1: google.api.DocumentationRule
	(*Page)(nil),              // 2: google.api.Page
}
var file_google_api_documentation_proto_depIdxs = []int32{
	2, // 0: google.api.Documentation.pages:type_name -> google.api.Page
	1, // 1: google.api.Documentation.rules:type_name -> google.api.DocumentationRule
	2, // 2: google.api.Page.subpages:type_name -> google.api.Page
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_api_documentation_proto_init() }
func file_google_api_documentation_proto_init() {
	if File_google_api_documentation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_documentation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Documentation); i {
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
		file_google_api_documentation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentationRule); i {
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
		file_google_api_documentation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
			RawDescriptor: file_google_api_documentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_documentation_proto_goTypes,
		DependencyIndexes: file_google_api_documentation_proto_depIdxs,
		MessageInfos:      file_google_api_documentation_proto_msgTypes,
	}.Build()
	File_google_api_documentation_proto = out.File
	file_google_api_documentation_proto_rawDesc = nil
	file_google_api_documentation_proto_goTypes = nil
	file_google_api_documentation_proto_depIdxs = nil
}
