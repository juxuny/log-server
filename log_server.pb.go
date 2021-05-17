// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: log_server.proto

package log_server

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

type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App  string `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_log_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_log_server_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *AddReq) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type AddResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AddResp) Reset() {
	*x = AddResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResp) ProtoMessage() {}

func (x *AddResp) ProtoReflect() protoreflect.Message {
	mi := &file_log_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResp.ProtoReflect.Descriptor instead.
func (*AddResp) Descriptor() ([]byte, []int) {
	return file_log_server_proto_rawDescGZIP(), []int{1}
}

func (x *AddResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_log_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_log_server_proto_rawDescGZIP(), []int{2}
}

type PingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingResp) Reset() {
	*x = PingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResp) ProtoMessage() {}

func (x *PingResp) ProtoReflect() protoreflect.Message {
	mi := &file_log_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResp.ProtoReflect.Descriptor instead.
func (*PingResp) Descriptor() ([]byte, []int) {
	return file_log_server_proto_rawDescGZIP(), []int{3}
}

type InfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoReq) Reset() {
	*x = InfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReq) ProtoMessage() {}

func (x *InfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_log_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoReq.ProtoReflect.Descriptor instead.
func (*InfoReq) Descriptor() ([]byte, []int) {
	return file_log_server_proto_rawDescGZIP(), []int{4}
}

type InfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoResp) Reset() {
	*x = InfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResp) ProtoMessage() {}

func (x *InfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_log_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResp.ProtoReflect.Descriptor instead.
func (*InfoResp) Descriptor() ([]byte, []int) {
	return file_log_server_proto_rawDescGZIP(), []int{5}
}

var File_log_server_proto protoreflect.FileDescriptor

var file_log_server_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x2e,
	0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1b,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x09, 0x0a, 0x07, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x22, 0x0a, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x09, 0x0a, 0x07, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x22, 0x0a, 0x0a,
	0x08, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x32, 0xa7, 0x01, 0x0a, 0x09, 0x4c, 0x6f,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x12,
	0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6c, 0x6f,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6a, 0x75, 0x78, 0x75, 0x6e, 0x79, 0x2f, 0x6c, 0x6f, 0x67, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_log_server_proto_rawDescOnce sync.Once
	file_log_server_proto_rawDescData = file_log_server_proto_rawDesc
)

func file_log_server_proto_rawDescGZIP() []byte {
	file_log_server_proto_rawDescOnce.Do(func() {
		file_log_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_server_proto_rawDescData)
	})
	return file_log_server_proto_rawDescData
}

var file_log_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_log_server_proto_goTypes = []interface{}{
	(*AddReq)(nil),   // 0: log_server.AddReq
	(*AddResp)(nil),  // 1: log_server.AddResp
	(*PingReq)(nil),  // 2: log_server.PingReq
	(*PingResp)(nil), // 3: log_server.PingResp
	(*InfoReq)(nil),  // 4: log_server.InfoReq
	(*InfoResp)(nil), // 5: log_server.InfoResp
}
var file_log_server_proto_depIdxs = []int32{
	0, // 0: log_server.LogServer.Add:input_type -> log_server.AddReq
	2, // 1: log_server.LogServer.Ping:input_type -> log_server.PingReq
	4, // 2: log_server.LogServer.Info:input_type -> log_server.InfoReq
	1, // 3: log_server.LogServer.Add:output_type -> log_server.AddResp
	3, // 4: log_server.LogServer.Ping:output_type -> log_server.PingResp
	5, // 5: log_server.LogServer.Info:output_type -> log_server.InfoResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_log_server_proto_init() }
func file_log_server_proto_init() {
	if File_log_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_log_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResp); i {
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
		file_log_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
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
		file_log_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResp); i {
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
		file_log_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoReq); i {
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
		file_log_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResp); i {
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
			RawDescriptor: file_log_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_log_server_proto_goTypes,
		DependencyIndexes: file_log_server_proto_depIdxs,
		MessageInfos:      file_log_server_proto_msgTypes,
	}.Build()
	File_log_server_proto = out.File
	file_log_server_proto_rawDesc = nil
	file_log_server_proto_goTypes = nil
	file_log_server_proto_depIdxs = nil
}
