// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api_delegate.proto

package api

import (
	pagination "github.com/iotexproject/iotex-analyser-api/api/pagination"
	_ "github.com/ysugimoto/grpc-graphql-gateway/graphql"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type DelegateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartEpoch             uint64                 `protobuf:"varint,1,opt,name=startEpoch,proto3" json:"startEpoch,omitempty"`
	EpochCount             uint64                 `protobuf:"varint,2,opt,name=epochCount,proto3" json:"epochCount,omitempty"`
	DelegateName           string                 `protobuf:"bytes,3,opt,name=delegateName,proto3" json:"delegateName,omitempty"`
	Pagination             *pagination.Pagination `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Percentage             uint64                 `protobuf:"varint,5,opt,name=percentage,proto3" json:"percentage,omitempty"`
	IncludeBlockReward     bool                   `protobuf:"varint,6,opt,name=includeBlockReward,proto3" json:"includeBlockReward,omitempty"`
	IncludeFoundationBonus bool                   `protobuf:"varint,7,opt,name=includeFoundationBonus,proto3" json:"includeFoundationBonus,omitempty"`
}

func (x *DelegateRequest) Reset() {
	*x = DelegateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_delegate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegateRequest) ProtoMessage() {}

func (x *DelegateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_delegate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelegateRequest.ProtoReflect.Descriptor instead.
func (*DelegateRequest) Descriptor() ([]byte, []int) {
	return file_api_delegate_proto_rawDescGZIP(), []int{0}
}

func (x *DelegateRequest) GetStartEpoch() uint64 {
	if x != nil {
		return x.StartEpoch
	}
	return 0
}

func (x *DelegateRequest) GetEpochCount() uint64 {
	if x != nil {
		return x.EpochCount
	}
	return 0
}

func (x *DelegateRequest) GetDelegateName() string {
	if x != nil {
		return x.DelegateName
	}
	return ""
}

func (x *DelegateRequest) GetPagination() *pagination.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *DelegateRequest) GetPercentage() uint64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *DelegateRequest) GetIncludeBlockReward() bool {
	if x != nil {
		return x.IncludeBlockReward
	}
	return false
}

func (x *DelegateRequest) GetIncludeFoundationBonus() bool {
	if x != nil {
		return x.IncludeFoundationBonus
	}
	return false
}

type BucketInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoterEthAddress   string `protobuf:"bytes,1,opt,name=voterEthAddress,proto3" json:"voterEthAddress,omitempty"`
	VoterIotexAddress string `protobuf:"bytes,2,opt,name=voterIotexAddress,proto3" json:"voterIotexAddress,omitempty"`
	IsNative          bool   `protobuf:"varint,3,opt,name=isNative,proto3" json:"isNative,omitempty"`
	Votes             string `protobuf:"bytes,4,opt,name=votes,proto3" json:"votes,omitempty"`
	WeightedVotes     string `protobuf:"bytes,5,opt,name=weightedVotes,proto3" json:"weightedVotes,omitempty"`
	RemainingDuration string `protobuf:"bytes,6,opt,name=remainingDuration,proto3" json:"remainingDuration,omitempty"`
	StartTime         string `protobuf:"bytes,7,opt,name=startTime,proto3" json:"startTime,omitempty"`
	Decay             bool   `protobuf:"varint,8,opt,name=decay,proto3" json:"decay,omitempty"`
	BucketID          uint64 `protobuf:"varint,9,opt,name=bucketID,proto3" json:"bucketID,omitempty"`
}

func (x *BucketInfo) Reset() {
	*x = BucketInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_delegate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketInfo) ProtoMessage() {}

func (x *BucketInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_delegate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketInfo.ProtoReflect.Descriptor instead.
func (*BucketInfo) Descriptor() ([]byte, []int) {
	return file_api_delegate_proto_rawDescGZIP(), []int{1}
}

func (x *BucketInfo) GetVoterEthAddress() string {
	if x != nil {
		return x.VoterEthAddress
	}
	return ""
}

func (x *BucketInfo) GetVoterIotexAddress() string {
	if x != nil {
		return x.VoterIotexAddress
	}
	return ""
}

func (x *BucketInfo) GetIsNative() bool {
	if x != nil {
		return x.IsNative
	}
	return false
}

func (x *BucketInfo) GetVotes() string {
	if x != nil {
		return x.Votes
	}
	return ""
}

func (x *BucketInfo) GetWeightedVotes() string {
	if x != nil {
		return x.WeightedVotes
	}
	return ""
}

func (x *BucketInfo) GetRemainingDuration() string {
	if x != nil {
		return x.RemainingDuration
	}
	return ""
}

func (x *BucketInfo) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *BucketInfo) GetDecay() bool {
	if x != nil {
		return x.Decay
	}
	return false
}

func (x *BucketInfo) GetBucketID() uint64 {
	if x != nil {
		return x.BucketID
	}
	return 0
}

type BucketInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpochNumber uint64        `protobuf:"varint,1,opt,name=epochNumber,proto3" json:"epochNumber,omitempty"`
	Count       uint64        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	BucketInfo  []*BucketInfo `protobuf:"bytes,3,rep,name=bucketInfo,proto3" json:"bucketInfo,omitempty"`
}

func (x *BucketInfoList) Reset() {
	*x = BucketInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_delegate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketInfoList) ProtoMessage() {}

func (x *BucketInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_api_delegate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketInfoList.ProtoReflect.Descriptor instead.
func (*BucketInfoList) Descriptor() ([]byte, []int) {
	return file_api_delegate_proto_rawDescGZIP(), []int{2}
}

func (x *BucketInfoList) GetEpochNumber() uint64 {
	if x != nil {
		return x.EpochNumber
	}
	return 0
}

func (x *BucketInfoList) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *BucketInfoList) GetBucketInfo() []*BucketInfo {
	if x != nil {
		return x.BucketInfo
	}
	return nil
}

type DelegateRewardDistribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoterEthAddress   string `protobuf:"bytes,1,opt,name=voterEthAddress,proto3" json:"voterEthAddress,omitempty"`
	VoterIotexAddress string `protobuf:"bytes,2,opt,name=voterIotexAddress,proto3" json:"voterIotexAddress,omitempty"`
	Amount            string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DelegateRewardDistribution) Reset() {
	*x = DelegateRewardDistribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_delegate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegateRewardDistribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegateRewardDistribution) ProtoMessage() {}

func (x *DelegateRewardDistribution) ProtoReflect() protoreflect.Message {
	mi := &file_api_delegate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelegateRewardDistribution.ProtoReflect.Descriptor instead.
func (*DelegateRewardDistribution) Descriptor() ([]byte, []int) {
	return file_api_delegate_proto_rawDescGZIP(), []int{3}
}

func (x *DelegateRewardDistribution) GetVoterEthAddress() string {
	if x != nil {
		return x.VoterEthAddress
	}
	return ""
}

func (x *DelegateRewardDistribution) GetVoterIotexAddress() string {
	if x != nil {
		return x.VoterIotexAddress
	}
	return ""
}

func (x *DelegateRewardDistribution) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type DelegateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist              bool                          `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
	Count              uint64                        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	BucketInfoList     []*BucketInfoList             `protobuf:"bytes,3,rep,name=bucketInfoList,proto3" json:"bucketInfoList,omitempty"`
	RewardDistribution []*DelegateRewardDistribution `protobuf:"bytes,4,rep,name=rewardDistribution,proto3" json:"rewardDistribution,omitempty"`
}

func (x *DelegateResponse) Reset() {
	*x = DelegateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_delegate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegateResponse) ProtoMessage() {}

func (x *DelegateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_delegate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelegateResponse.ProtoReflect.Descriptor instead.
func (*DelegateResponse) Descriptor() ([]byte, []int) {
	return file_api_delegate_proto_rawDescGZIP(), []int{4}
}

func (x *DelegateResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *DelegateResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *DelegateResponse) GetBucketInfoList() []*BucketInfoList {
	if x != nil {
		return x.BucketInfoList
	}
	return nil
}

func (x *DelegateResponse) GetRewardDistribution() []*DelegateRewardDistribution {
	if x != nil {
		return x.RewardDistribution
	}
	return nil
}

var File_api_delegate_proto protoreflect.FileDescriptor

var file_api_delegate_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x02, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x70, 0x6f, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x36, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x6e,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73,
	0x22, 0xba, 0x02, 0x0a, 0x0a, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x28, 0x0a, 0x0f, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x45, 0x74, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x45,
	0x74, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x76, 0x6f, 0x74,
	0x65, 0x72, 0x49, 0x6f, 0x74, 0x65, 0x78, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x49, 0x6f, 0x74, 0x65, 0x78,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4e, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4e, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x65, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x2c, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x65, 0x63, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x63, 0x61,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x22, 0x79, 0x0a,
	0x0e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x1a, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x76, 0x6f, 0x74, 0x65, 0x72,
	0x45, 0x74, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x45, 0x74, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x2c, 0x0a, 0x11, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x49, 0x6f, 0x74, 0x65, 0x78, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x6f,
	0x74, 0x65, 0x72, 0x49, 0x6f, 0x74, 0x65, 0x78, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xcc, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0e, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x12, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x12, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x80, 0x02, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x0a, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0xba, 0x43, 0x0c, 0x12, 0x0a, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x3a, 0x01, 0x2a,
	0x12, 0x77, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x4b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x12,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0xba, 0x43,
	0x0d, 0x12, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x4b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x25, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4b,
	0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_delegate_proto_rawDescOnce sync.Once
	file_api_delegate_proto_rawDescData = file_api_delegate_proto_rawDesc
)

func file_api_delegate_proto_rawDescGZIP() []byte {
	file_api_delegate_proto_rawDescOnce.Do(func() {
		file_api_delegate_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_delegate_proto_rawDescData)
	})
	return file_api_delegate_proto_rawDescData
}

var file_api_delegate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_delegate_proto_goTypes = []interface{}{
	(*DelegateRequest)(nil),            // 0: api.DelegateRequest
	(*BucketInfo)(nil),                 // 1: api.BucketInfo
	(*BucketInfoList)(nil),             // 2: api.BucketInfoList
	(*DelegateRewardDistribution)(nil), // 3: api.DelegateRewardDistribution
	(*DelegateResponse)(nil),           // 4: api.DelegateResponse
	(*pagination.Pagination)(nil),      // 5: pagination.Pagination
}
var file_api_delegate_proto_depIdxs = []int32{
	5, // 0: api.DelegateRequest.pagination:type_name -> pagination.Pagination
	1, // 1: api.BucketInfoList.bucketInfo:type_name -> api.BucketInfo
	2, // 2: api.DelegateResponse.bucketInfoList:type_name -> api.BucketInfoList
	3, // 3: api.DelegateResponse.rewardDistribution:type_name -> api.DelegateRewardDistribution
	0, // 4: api.DelegateService.BucketInfo:input_type -> api.DelegateRequest
	0, // 5: api.DelegateService.BookKeeping:input_type -> api.DelegateRequest
	4, // 6: api.DelegateService.BucketInfo:output_type -> api.DelegateResponse
	4, // 7: api.DelegateService.BookKeeping:output_type -> api.DelegateResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_delegate_proto_init() }
func file_api_delegate_proto_init() {
	if File_api_delegate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_delegate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegateRequest); i {
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
		file_api_delegate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketInfo); i {
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
		file_api_delegate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketInfoList); i {
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
		file_api_delegate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegateRewardDistribution); i {
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
		file_api_delegate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegateResponse); i {
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
			RawDescriptor: file_api_delegate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_delegate_proto_goTypes,
		DependencyIndexes: file_api_delegate_proto_depIdxs,
		MessageInfos:      file_api_delegate_proto_msgTypes,
	}.Build()
	File_api_delegate_proto = out.File
	file_api_delegate_proto_rawDesc = nil
	file_api_delegate_proto_goTypes = nil
	file_api_delegate_proto_depIdxs = nil
}
