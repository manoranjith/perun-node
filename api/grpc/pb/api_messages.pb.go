// Copyright (c) 2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/hyperledger-labs/perun-node
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
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: api_messages.proto

// Package pb contains proto3 definitions for user API and the corresponding
// generated code for grpc server and client.

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

// This type is defined as the enumeration of all messages in funding and
// watching service, in order to be able to parse the messages in api/tcp
// package.
type API_Messages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*API_Messages_FundingReq
	//	*API_Messages_FundingResp
	Msg isAPI_Messages_Msg `protobuf_oneof:"msg"`
}

func (x *API_Messages) Reset() {
	*x = API_Messages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *API_Messages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*API_Messages) ProtoMessage() {}

func (x *API_Messages) ProtoReflect() protoreflect.Message {
	mi := &file_api_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use API_Messages.ProtoReflect.Descriptor instead.
func (*API_Messages) Descriptor() ([]byte, []int) {
	return file_api_messages_proto_rawDescGZIP(), []int{0}
}

func (m *API_Messages) GetMsg() isAPI_Messages_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *API_Messages) GetFundingReq() *FundReq {
	if x, ok := x.GetMsg().(*API_Messages_FundingReq); ok {
		return x.FundingReq
	}
	return nil
}

func (x *API_Messages) GetFundingResp() *FundResp {
	if x, ok := x.GetMsg().(*API_Messages_FundingResp); ok {
		return x.FundingResp
	}
	return nil
}

type isAPI_Messages_Msg interface {
	isAPI_Messages_Msg()
}

type API_Messages_FundingReq struct {
	FundingReq *FundReq `protobuf:"bytes,1,opt,name=funding_req,json=fundingReq,proto3,oneof"`
}

type API_Messages_FundingResp struct {
	FundingResp *FundResp `protobuf:"bytes,2,opt,name=funding_resp,json=fundingResp,proto3,oneof"`
}

func (*API_Messages_FundingReq) isAPI_Messages_Msg() {}

func (*API_Messages_FundingResp) isAPI_Messages_Msg() {}

var File_api_messages_proto protoreflect.FileDescriptor

var file_api_messages_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x15, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x78, 0x0a, 0x0c, 0x41, 0x50, 0x49, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x2e, 0x0a, 0x0b, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x31, 0x0a, 0x0c, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x48, 0x00, 0x52, 0x0b, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_messages_proto_rawDescOnce sync.Once
	file_api_messages_proto_rawDescData = file_api_messages_proto_rawDesc
)

func file_api_messages_proto_rawDescGZIP() []byte {
	file_api_messages_proto_rawDescOnce.Do(func() {
		file_api_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_messages_proto_rawDescData)
	})
	return file_api_messages_proto_rawDescData
}

var file_api_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_messages_proto_goTypes = []interface{}{
	(*API_Messages)(nil), // 0: pb.API_Messages
	(*FundReq)(nil),      // 1: pb.FundReq
	(*FundResp)(nil),     // 2: pb.FundResp
}
var file_api_messages_proto_depIdxs = []int32{
	1, // 0: pb.API_Messages.funding_req:type_name -> pb.FundReq
	2, // 1: pb.API_Messages.funding_resp:type_name -> pb.FundResp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_messages_proto_init() }
func file_api_messages_proto_init() {
	if File_api_messages_proto != nil {
		return
	}
	file_funding_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*API_Messages); i {
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
	file_api_messages_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*API_Messages_FundingReq)(nil),
		(*API_Messages_FundingResp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_messages_proto_goTypes,
		DependencyIndexes: file_api_messages_proto_depIdxs,
		MessageInfos:      file_api_messages_proto_msgTypes,
	}.Build()
	File_api_messages_proto = out.File
	file_api_messages_proto_rawDesc = nil
	file_api_messages_proto_goTypes = nil
	file_api_messages_proto_depIdxs = nil
}
