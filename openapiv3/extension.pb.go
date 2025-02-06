// Copyright 2022 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: openapiv3/extension.proto

package openapi_v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Service struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Label         string                 `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Service) Reset() {
	*x = Service{}
	mi := &file_openapiv3_extension_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_openapiv3_extension_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_openapiv3_extension_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var file_openapiv3_extension_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*Service)(nil),
		Field:         1143,
		Name:          "openapi.v3.service",
		Tag:           "bytes,1143,opt,name=service",
		Filename:      "openapiv3/extension.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional openapi.v3.Service service = 1143;
	E_Service = &file_openapiv3_extension_proto_extTypes[0]
)

var File_openapiv3_extension_proto protoreflect.FileDescriptor

var file_openapiv3_extension_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x33, 0x1a, 0x19, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x33, 0x2f, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x76, 0x33, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3a, 0x4f, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xf7, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x9f, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x72, 0x6f, 0x65, 0x6e, 0x72, 0x69,
	0x6e, 0x7a, 0x65, 0x6d, 0x61, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x3b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x33, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5c, 0x56,
	0x33, 0xe2, 0x02, 0x16, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5c, 0x56, 0x33, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x4f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_openapiv3_extension_proto_rawDescOnce sync.Once
	file_openapiv3_extension_proto_rawDescData []byte
)

func file_openapiv3_extension_proto_rawDescGZIP() []byte {
	file_openapiv3_extension_proto_rawDescOnce.Do(func() {
		file_openapiv3_extension_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_openapiv3_extension_proto_rawDesc), len(file_openapiv3_extension_proto_rawDesc)))
	})
	return file_openapiv3_extension_proto_rawDescData
}

var file_openapiv3_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_openapiv3_extension_proto_goTypes = []any{
	(*Service)(nil),                     // 0: openapi.v3.Service
	(*descriptorpb.ServiceOptions)(nil), // 1: google.protobuf.ServiceOptions
}
var file_openapiv3_extension_proto_depIdxs = []int32{
	1, // 0: openapi.v3.service:extendee -> google.protobuf.ServiceOptions
	0, // 1: openapi.v3.service:type_name -> openapi.v3.Service
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_openapiv3_extension_proto_init() }
func file_openapiv3_extension_proto_init() {
	if File_openapiv3_extension_proto != nil {
		return
	}
	file_openapiv3_OpenAPIv3_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_openapiv3_extension_proto_rawDesc), len(file_openapiv3_extension_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_openapiv3_extension_proto_goTypes,
		DependencyIndexes: file_openapiv3_extension_proto_depIdxs,
		MessageInfos:      file_openapiv3_extension_proto_msgTypes,
		ExtensionInfos:    file_openapiv3_extension_proto_extTypes,
	}.Build()
	File_openapiv3_extension_proto = out.File
	file_openapiv3_extension_proto_goTypes = nil
	file_openapiv3_extension_proto_depIdxs = nil
}
