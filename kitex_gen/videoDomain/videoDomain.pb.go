// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: videoDomain.proto

package videoDomain

import (
	context "context"
	entity "github.com/TremblingV5/DouTok/kitex_gen/entity"
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

type DoutokGetFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestTime int64 `protobuf:"varint,1,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	UserId     int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             // 请求feed流的用户id，若没有则置为0
}

func (x *DoutokGetFeedRequest) Reset() {
	*x = DoutokGetFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoDomain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoutokGetFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoutokGetFeedRequest) ProtoMessage() {}

func (x *DoutokGetFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_videoDomain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoutokGetFeedRequest.ProtoReflect.Descriptor instead.
func (*DoutokGetFeedRequest) Descriptor() ([]byte, []int) {
	return file_videoDomain_proto_rawDescGZIP(), []int{0}
}

func (x *DoutokGetFeedRequest) GetLatestTime() int64 {
	if x != nil {
		return x.LatestTime
	}
	return 0
}

func (x *DoutokGetFeedRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DoutokGetFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32           `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string          `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	VideoList  []*entity.Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
	NextTime   int64           `protobuf:"varint,4,opt,name=next_time,json=nextTime,proto3" json:"next_time,omitempty"`
}

func (x *DoutokGetFeedResponse) Reset() {
	*x = DoutokGetFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoDomain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoutokGetFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoutokGetFeedResponse) ProtoMessage() {}

func (x *DoutokGetFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoDomain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoutokGetFeedResponse.ProtoReflect.Descriptor instead.
func (*DoutokGetFeedResponse) Descriptor() ([]byte, []int) {
	return file_videoDomain_proto_rawDescGZIP(), []int{1}
}

func (x *DoutokGetFeedResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DoutokGetFeedResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *DoutokGetFeedResponse) GetVideoList() []*entity.Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *DoutokGetFeedResponse) GetNextTime() int64 {
	if x != nil {
		return x.NextTime
	}
	return 0
}

type DoutokAddPublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	UserId int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DoutokAddPublishRequest) Reset() {
	*x = DoutokAddPublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoDomain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoutokAddPublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoutokAddPublishRequest) ProtoMessage() {}

func (x *DoutokAddPublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_videoDomain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoutokAddPublishRequest.ProtoReflect.Descriptor instead.
func (*DoutokAddPublishRequest) Descriptor() ([]byte, []int) {
	return file_videoDomain_proto_rawDescGZIP(), []int{2}
}

func (x *DoutokAddPublishRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DoutokAddPublishRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DoutokAddPublishRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DoutokAddPublishRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DoutokAddPublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *DoutokAddPublishResponse) Reset() {
	*x = DoutokAddPublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoDomain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoutokAddPublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoutokAddPublishResponse) ProtoMessage() {}

func (x *DoutokAddPublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoDomain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoutokAddPublishResponse.ProtoReflect.Descriptor instead.
func (*DoutokAddPublishResponse) Descriptor() ([]byte, []int) {
	return file_videoDomain_proto_rawDescGZIP(), []int{3}
}

func (x *DoutokAddPublishResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DoutokAddPublishResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type DoutokListPublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DoutokListPublishRequest) Reset() {
	*x = DoutokListPublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoDomain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoutokListPublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoutokListPublishRequest) ProtoMessage() {}

func (x *DoutokListPublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_videoDomain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoutokListPublishRequest.ProtoReflect.Descriptor instead.
func (*DoutokListPublishRequest) Descriptor() ([]byte, []int) {
	return file_videoDomain_proto_rawDescGZIP(), []int{4}
}

func (x *DoutokListPublishRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DoutokListPublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32           `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string          `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	VideoList  []*entity.Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
}

func (x *DoutokListPublishResponse) Reset() {
	*x = DoutokListPublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoDomain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoutokListPublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoutokListPublishResponse) ProtoMessage() {}

func (x *DoutokListPublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoDomain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoutokListPublishResponse.ProtoReflect.Descriptor instead.
func (*DoutokListPublishResponse) Descriptor() ([]byte, []int) {
	return file_videoDomain_proto_rawDescGZIP(), []int{5}
}

func (x *DoutokListPublishResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DoutokListPublishResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *DoutokListPublishResponse) GetVideoList() []*entity.Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type DoutokCountPublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId []int64 `protobuf:"varint,1,rep,packed,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DoutokCountPublishRequest) Reset() {
	*x = DoutokCountPublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoDomain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoutokCountPublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoutokCountPublishRequest) ProtoMessage() {}

func (x *DoutokCountPublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_videoDomain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoutokCountPublishRequest.ProtoReflect.Descriptor instead.
func (*DoutokCountPublishRequest) Descriptor() ([]byte, []int) {
	return file_videoDomain_proto_rawDescGZIP(), []int{6}
}

func (x *DoutokCountPublishRequest) GetUserId() []int64 {
	if x != nil {
		return x.UserId
	}
	return nil
}

type DoutokCountPublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32           `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string          `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Result     map[int64]int64 `protobuf:"bytes,3,rep,name=result,proto3" json:"result,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *DoutokCountPublishResponse) Reset() {
	*x = DoutokCountPublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoDomain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoutokCountPublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoutokCountPublishResponse) ProtoMessage() {}

func (x *DoutokCountPublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoDomain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoutokCountPublishResponse.ProtoReflect.Descriptor instead.
func (*DoutokCountPublishResponse) Descriptor() ([]byte, []int) {
	return file_videoDomain_proto_rawDescGZIP(), []int{7}
}

func (x *DoutokCountPublishResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DoutokCountPublishResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *DoutokCountPublishResponse) GetResult() map[int64]int64 {
	if x != nil {
		return x.Result
	}
	return nil
}

type DoutokGetVideoInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *DoutokGetVideoInfoRequest) Reset() {
	*x = DoutokGetVideoInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoDomain_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoutokGetVideoInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoutokGetVideoInfoRequest) ProtoMessage() {}

func (x *DoutokGetVideoInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_videoDomain_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoutokGetVideoInfoRequest.ProtoReflect.Descriptor instead.
func (*DoutokGetVideoInfoRequest) Descriptor() ([]byte, []int) {
	return file_videoDomain_proto_rawDescGZIP(), []int{8}
}

func (x *DoutokGetVideoInfoRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

var File_videoDomain_proto protoreflect.FileDescriptor

var file_videoDomain_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x1a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53,
	0x0a, 0x17, 0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x65, 0x65,
	0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x18, 0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x67,
	0x65, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67,
	0x12, 0x2c, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x73, 0x0a, 0x1a, 0x64,
	0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x5d, 0x0a, 0x1b, 0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22,
	0x36, 0x0a, 0x1b, 0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x1c, 0x64, 0x6f, 0x75, 0x74,
	0x6f, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x1c, 0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xea, 0x01, 0x0a, 0x1d, 0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73,
	0x67, 0x12, 0x4e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3a, 0x0a, 0x1d,
	0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x32, 0xe3, 0x03, 0x0a, 0x12, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x56, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x12, 0x24, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f,
	0x67, 0x65, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x64,
	0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x27, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x64, 0x6f, 0x75,
	0x74, 0x6f, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x28, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x0c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x29, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x64, 0x6f, 0x75, 0x74, 0x6f,
	0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x64, 0x6f, 0x75, 0x74, 0x6f, 0x6b, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x72, 0x65,
	0x6d, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x56, 0x35, 0x2f, 0x44, 0x6f, 0x75, 0x54, 0x6f, 0x6b, 0x2f,
	0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_videoDomain_proto_rawDescOnce sync.Once
	file_videoDomain_proto_rawDescData = file_videoDomain_proto_rawDesc
)

func file_videoDomain_proto_rawDescGZIP() []byte {
	file_videoDomain_proto_rawDescOnce.Do(func() {
		file_videoDomain_proto_rawDescData = protoimpl.X.CompressGZIP(file_videoDomain_proto_rawDescData)
	})
	return file_videoDomain_proto_rawDescData
}

var file_videoDomain_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_videoDomain_proto_goTypes = []interface{}{
	(*DoutokGetFeedRequest)(nil),       // 0: videoDomain.doutok_get_feed_request
	(*DoutokGetFeedResponse)(nil),      // 1: videoDomain.doutok_get_feed_response
	(*DoutokAddPublishRequest)(nil),    // 2: videoDomain.doutok_add_publish_request
	(*DoutokAddPublishResponse)(nil),   // 3: videoDomain.doutok_add_publish_response
	(*DoutokListPublishRequest)(nil),   // 4: videoDomain.doutok_list_publish_request
	(*DoutokListPublishResponse)(nil),  // 5: videoDomain.doutok_list_publish_response
	(*DoutokCountPublishRequest)(nil),  // 6: videoDomain.doutok_count_publish_request
	(*DoutokCountPublishResponse)(nil), // 7: videoDomain.doutok_count_publish_response
	(*DoutokGetVideoInfoRequest)(nil),  // 8: videoDomain.doutok_get_video_info_request
	nil,                                // 9: videoDomain.doutok_count_publish_response.ResultEntry
	(*entity.Video)(nil),               // 10: entity.Video
}
var file_videoDomain_proto_depIdxs = []int32{
	10, // 0: videoDomain.doutok_get_feed_response.video_list:type_name -> entity.Video
	10, // 1: videoDomain.doutok_list_publish_response.video_list:type_name -> entity.Video
	9,  // 2: videoDomain.doutok_count_publish_response.result:type_name -> videoDomain.doutok_count_publish_response.ResultEntry
	0,  // 3: videoDomain.VideoDomainService.GetFeed:input_type -> videoDomain.doutok_get_feed_request
	2,  // 4: videoDomain.VideoDomainService.AddPublish:input_type -> videoDomain.doutok_add_publish_request
	4,  // 5: videoDomain.VideoDomainService.ListPublish:input_type -> videoDomain.doutok_list_publish_request
	6,  // 6: videoDomain.VideoDomainService.CountPublish:input_type -> videoDomain.doutok_count_publish_request
	8,  // 7: videoDomain.VideoDomainService.GetVideoInfo:input_type -> videoDomain.doutok_get_video_info_request
	1,  // 8: videoDomain.VideoDomainService.GetFeed:output_type -> videoDomain.doutok_get_feed_response
	3,  // 9: videoDomain.VideoDomainService.AddPublish:output_type -> videoDomain.doutok_add_publish_response
	5,  // 10: videoDomain.VideoDomainService.ListPublish:output_type -> videoDomain.doutok_list_publish_response
	7,  // 11: videoDomain.VideoDomainService.CountPublish:output_type -> videoDomain.doutok_count_publish_response
	10, // 12: videoDomain.VideoDomainService.GetVideoInfo:output_type -> entity.Video
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_videoDomain_proto_init() }
func file_videoDomain_proto_init() {
	if File_videoDomain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_videoDomain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoutokGetFeedRequest); i {
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
		file_videoDomain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoutokGetFeedResponse); i {
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
		file_videoDomain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoutokAddPublishRequest); i {
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
		file_videoDomain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoutokAddPublishResponse); i {
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
		file_videoDomain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoutokListPublishRequest); i {
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
		file_videoDomain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoutokListPublishResponse); i {
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
		file_videoDomain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoutokCountPublishRequest); i {
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
		file_videoDomain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoutokCountPublishResponse); i {
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
		file_videoDomain_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoutokGetVideoInfoRequest); i {
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
			RawDescriptor: file_videoDomain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_videoDomain_proto_goTypes,
		DependencyIndexes: file_videoDomain_proto_depIdxs,
		MessageInfos:      file_videoDomain_proto_msgTypes,
	}.Build()
	File_videoDomain_proto = out.File
	file_videoDomain_proto_rawDesc = nil
	file_videoDomain_proto_goTypes = nil
	file_videoDomain_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.4.4. DO NOT EDIT.

type VideoDomainService interface {
	GetFeed(ctx context.Context, req *DoutokGetFeedRequest) (res *DoutokGetFeedResponse, err error)
	AddPublish(ctx context.Context, req *DoutokAddPublishRequest) (res *DoutokAddPublishResponse, err error)
	ListPublish(ctx context.Context, req *DoutokListPublishRequest) (res *DoutokListPublishResponse, err error)
	CountPublish(ctx context.Context, req *DoutokCountPublishRequest) (res *DoutokCountPublishResponse, err error)
	GetVideoInfo(ctx context.Context, req *DoutokGetVideoInfoRequest) (res *entity.Video, err error)
}