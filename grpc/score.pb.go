// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: score.proto

package main

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

type ListMatchesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *ListMatchesRequest) Reset() {
	*x = ListMatchesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMatchesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMatchesRequest) ProtoMessage() {}

func (x *ListMatchesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_score_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMatchesRequest.ProtoReflect.Descriptor instead.
func (*ListMatchesRequest) Descriptor() ([]byte, []int) {
	return file_score_proto_rawDescGZIP(), []int{0}
}

func (x *ListMatchesRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type MatchScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score string `protobuf:"bytes,1,opt,name=score,proto3" json:"score,omitempty"`
	Live  bool   `protobuf:"varint,2,opt,name=live,proto3" json:"live,omitempty"`
}

func (x *MatchScoreResponse) Reset() {
	*x = MatchScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchScoreResponse) ProtoMessage() {}

func (x *MatchScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_score_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchScoreResponse.ProtoReflect.Descriptor instead.
func (*MatchScoreResponse) Descriptor() ([]byte, []int) {
	return file_score_proto_rawDescGZIP(), []int{1}
}

func (x *MatchScoreResponse) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *MatchScoreResponse) GetLive() bool {
	if x != nil {
		return x.Live
	}
	return false
}

type ListMatchesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scores []*MatchScoreResponse `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
}

func (x *ListMatchesResponse) Reset() {
	*x = ListMatchesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMatchesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMatchesResponse) ProtoMessage() {}

func (x *ListMatchesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_score_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMatchesResponse.ProtoReflect.Descriptor instead.
func (*ListMatchesResponse) Descriptor() ([]byte, []int) {
	return file_score_proto_rawDescGZIP(), []int{2}
}

func (x *ListMatchesResponse) GetScores() []*MatchScoreResponse {
	if x != nil {
		return x.Scores
	}
	return nil
}

var File_score_proto protoreflect.FileDescriptor

var file_score_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x3e, 0x0a,
	0x12, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x76,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x69, 0x76, 0x65, 0x22, 0x42, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x32, 0x48, 0x0a, 0x0c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x38, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x12, 0x13, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_score_proto_rawDescOnce sync.Once
	file_score_proto_rawDescData = file_score_proto_rawDesc
)

func file_score_proto_rawDescGZIP() []byte {
	file_score_proto_rawDescOnce.Do(func() {
		file_score_proto_rawDescData = protoimpl.X.CompressGZIP(file_score_proto_rawDescData)
	})
	return file_score_proto_rawDescData
}

var file_score_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_score_proto_goTypes = []interface{}{
	(*ListMatchesRequest)(nil),  // 0: ListMatchesRequest
	(*MatchScoreResponse)(nil),  // 1: MatchScoreResponse
	(*ListMatchesResponse)(nil), // 2: ListMatchesResponse
}
var file_score_proto_depIdxs = []int32{
	1, // 0: ListMatchesResponse.scores:type_name -> MatchScoreResponse
	0, // 1: ScoreService.ListMatches:input_type -> ListMatchesRequest
	2, // 2: ScoreService.ListMatches:output_type -> ListMatchesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_score_proto_init() }
func file_score_proto_init() {
	if File_score_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_score_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMatchesRequest); i {
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
		file_score_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchScoreResponse); i {
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
		file_score_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMatchesResponse); i {
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
			RawDescriptor: file_score_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_score_proto_goTypes,
		DependencyIndexes: file_score_proto_depIdxs,
		MessageInfos:      file_score_proto_msgTypes,
	}.Build()
	File_score_proto = out.File
	file_score_proto_rawDesc = nil
	file_score_proto_goTypes = nil
	file_score_proto_depIdxs = nil
}
