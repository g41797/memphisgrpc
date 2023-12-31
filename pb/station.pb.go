// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: station.proto

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

type StorageOpt_StorageType int32

const (
	StorageOpt_Disk   StorageOpt_StorageType = 0
	StorageOpt_Memory StorageOpt_StorageType = 1
)

// Enum value maps for StorageOpt_StorageType.
var (
	StorageOpt_StorageType_name = map[int32]string{
		0: "Disk",
		1: "Memory",
	}
	StorageOpt_StorageType_value = map[string]int32{
		"Disk":   0,
		"Memory": 1,
	}
)

func (x StorageOpt_StorageType) Enum() *StorageOpt_StorageType {
	p := new(StorageOpt_StorageType)
	*p = x
	return p
}

func (x StorageOpt_StorageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageOpt_StorageType) Descriptor() protoreflect.EnumDescriptor {
	return file_station_proto_enumTypes[0].Descriptor()
}

func (StorageOpt_StorageType) Type() protoreflect.EnumType {
	return &file_station_proto_enumTypes[0]
}

func (x StorageOpt_StorageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageOpt_StorageType.Descriptor instead.
func (StorageOpt_StorageType) EnumDescriptor() ([]byte, []int) {
	return file_station_proto_rawDescGZIP(), []int{6, 0}
}

type Station struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Station) Reset() {
	*x = Station{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Station) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Station) ProtoMessage() {}

func (x *Station) ProtoReflect() protoreflect.Message {
	mi := &file_station_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Station.ProtoReflect.Descriptor instead.
func (*Station) Descriptor() ([]byte, []int) {
	return file_station_proto_rawDescGZIP(), []int{0}
}

func (x *Station) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MaxMessageAgeSecondsRet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds int32 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
}

func (x *MaxMessageAgeSecondsRet) Reset() {
	*x = MaxMessageAgeSecondsRet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxMessageAgeSecondsRet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxMessageAgeSecondsRet) ProtoMessage() {}

func (x *MaxMessageAgeSecondsRet) ProtoReflect() protoreflect.Message {
	mi := &file_station_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxMessageAgeSecondsRet.ProtoReflect.Descriptor instead.
func (*MaxMessageAgeSecondsRet) Descriptor() ([]byte, []int) {
	return file_station_proto_rawDescGZIP(), []int{1}
}

func (x *MaxMessageAgeSecondsRet) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

type MessagesRet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages int32 `protobuf:"varint,1,opt,name=messages,proto3" json:"messages,omitempty"`
}

func (x *MessagesRet) Reset() {
	*x = MessagesRet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagesRet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagesRet) ProtoMessage() {}

func (x *MessagesRet) ProtoReflect() protoreflect.Message {
	mi := &file_station_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagesRet.ProtoReflect.Descriptor instead.
func (*MessagesRet) Descriptor() ([]byte, []int) {
	return file_station_proto_rawDescGZIP(), []int{2}
}

func (x *MessagesRet) GetMessages() int32 {
	if x != nil {
		return x.Messages
	}
	return 0
}

type BytesRet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes int32 `protobuf:"varint,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *BytesRet) Reset() {
	*x = BytesRet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesRet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesRet) ProtoMessage() {}

func (x *BytesRet) ProtoReflect() protoreflect.Message {
	mi := &file_station_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesRet.ProtoReflect.Descriptor instead.
func (*BytesRet) Descriptor() ([]byte, []int) {
	return file_station_proto_rawDescGZIP(), []int{3}
}

func (x *BytesRet) GetBytes() int32 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

type AckBasedRet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AckBasedRet) Reset() {
	*x = AckBasedRet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckBasedRet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckBasedRet) ProtoMessage() {}

func (x *AckBasedRet) ProtoReflect() protoreflect.Message {
	mi := &file_station_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckBasedRet.ProtoReflect.Descriptor instead.
func (*AckBasedRet) Descriptor() ([]byte, []int) {
	return file_station_proto_rawDescGZIP(), []int{4}
}

type RetentionOpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Retentions:
	//
	//	*RetentionOpt_Mmasret
	//	*RetentionOpt_Msgret
	//	*RetentionOpt_Bret
	//	*RetentionOpt_Abret
	Retentions isRetentionOpt_Retentions `protobuf_oneof:"retentions"`
}

func (x *RetentionOpt) Reset() {
	*x = RetentionOpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetentionOpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetentionOpt) ProtoMessage() {}

func (x *RetentionOpt) ProtoReflect() protoreflect.Message {
	mi := &file_station_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetentionOpt.ProtoReflect.Descriptor instead.
func (*RetentionOpt) Descriptor() ([]byte, []int) {
	return file_station_proto_rawDescGZIP(), []int{5}
}

func (m *RetentionOpt) GetRetentions() isRetentionOpt_Retentions {
	if m != nil {
		return m.Retentions
	}
	return nil
}

func (x *RetentionOpt) GetMmasret() *MaxMessageAgeSecondsRet {
	if x, ok := x.GetRetentions().(*RetentionOpt_Mmasret); ok {
		return x.Mmasret
	}
	return nil
}

func (x *RetentionOpt) GetMsgret() *MessagesRet {
	if x, ok := x.GetRetentions().(*RetentionOpt_Msgret); ok {
		return x.Msgret
	}
	return nil
}

func (x *RetentionOpt) GetBret() *BytesRet {
	if x, ok := x.GetRetentions().(*RetentionOpt_Bret); ok {
		return x.Bret
	}
	return nil
}

func (x *RetentionOpt) GetAbret() *AckBasedRet {
	if x, ok := x.GetRetentions().(*RetentionOpt_Abret); ok {
		return x.Abret
	}
	return nil
}

type isRetentionOpt_Retentions interface {
	isRetentionOpt_Retentions()
}

type RetentionOpt_Mmasret struct {
	Mmasret *MaxMessageAgeSecondsRet `protobuf:"bytes,1,opt,name=mmasret,proto3,oneof"`
}

type RetentionOpt_Msgret struct {
	Msgret *MessagesRet `protobuf:"bytes,2,opt,name=msgret,proto3,oneof"`
}

type RetentionOpt_Bret struct {
	Bret *BytesRet `protobuf:"bytes,3,opt,name=bret,proto3,oneof"`
}

type RetentionOpt_Abret struct {
	Abret *AckBasedRet `protobuf:"bytes,4,opt,name=abret,proto3,oneof"`
}

func (*RetentionOpt_Mmasret) isRetentionOpt_Retentions() {}

func (*RetentionOpt_Msgret) isRetentionOpt_Retentions() {}

func (*RetentionOpt_Bret) isRetentionOpt_Retentions() {}

func (*RetentionOpt_Abret) isRetentionOpt_Retentions() {}

type StorageOpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageType StorageOpt_StorageType `protobuf:"varint,1,opt,name=storageType,proto3,enum=pb.StorageOpt_StorageType" json:"storageType,omitempty"`
}

func (x *StorageOpt) Reset() {
	*x = StorageOpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageOpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageOpt) ProtoMessage() {}

func (x *StorageOpt) ProtoReflect() protoreflect.Message {
	mi := &file_station_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageOpt.ProtoReflect.Descriptor instead.
func (*StorageOpt) Descriptor() ([]byte, []int) {
	return file_station_proto_rawDescGZIP(), []int{6}
}

func (x *StorageOpt) GetStorageType() StorageOpt_StorageType {
	if x != nil {
		return x.StorageType
	}
	return StorageOpt_Disk
}

type PartitionOpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PartitionOpt) Reset() {
	*x = PartitionOpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartitionOpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionOpt) ProtoMessage() {}

func (x *PartitionOpt) ProtoReflect() protoreflect.Message {
	mi := &file_station_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionOpt.ProtoReflect.Descriptor instead.
func (*PartitionOpt) Descriptor() ([]byte, []int) {
	return file_station_proto_rawDescGZIP(), []int{7}
}

func (x *PartitionOpt) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type StationOpions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storage   *StorageOpt   `protobuf:"bytes,1,opt,name=storage,proto3,oneof" json:"storage,omitempty"`
	Retention *RetentionOpt `protobuf:"bytes,2,opt,name=retention,proto3,oneof" json:"retention,omitempty"`
}

func (x *StationOpions) Reset() {
	*x = StationOpions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationOpions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationOpions) ProtoMessage() {}

func (x *StationOpions) ProtoReflect() protoreflect.Message {
	mi := &file_station_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationOpions.ProtoReflect.Descriptor instead.
func (*StationOpions) Descriptor() ([]byte, []int) {
	return file_station_proto_rawDescGZIP(), []int{8}
}

func (x *StationOpions) GetStorage() *StorageOpt {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *StationOpions) GetRetention() *RetentionOpt {
	if x != nil {
		return x.Retention
	}
	return nil
}

type CreateStationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Station *Station       `protobuf:"bytes,1,opt,name=station,proto3" json:"station,omitempty"`
	Options *StationOpions `protobuf:"bytes,2,opt,name=options,proto3,oneof" json:"options,omitempty"`
}

func (x *CreateStationRequest) Reset() {
	*x = CreateStationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStationRequest) ProtoMessage() {}

func (x *CreateStationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_station_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStationRequest.ProtoReflect.Descriptor instead.
func (*CreateStationRequest) Descriptor() ([]byte, []int) {
	return file_station_proto_rawDescGZIP(), []int{9}
}

func (x *CreateStationRequest) GetStation() *Station {
	if x != nil {
		return x.Station
	}
	return nil
}

func (x *CreateStationRequest) GetOptions() *StationOpions {
	if x != nil {
		return x.Options
	}
	return nil
}

type DestroyStationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Station *Station `protobuf:"bytes,1,opt,name=station,proto3" json:"station,omitempty"`
}

func (x *DestroyStationRequest) Reset() {
	*x = DestroyStationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyStationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyStationRequest) ProtoMessage() {}

func (x *DestroyStationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_station_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyStationRequest.ProtoReflect.Descriptor instead.
func (*DestroyStationRequest) Descriptor() ([]byte, []int) {
	return file_station_proto_rawDescGZIP(), []int{10}
}

func (x *DestroyStationRequest) GetStation() *Station {
	if x != nil {
		return x.Station
	}
	return nil
}

var File_station_proto protoreflect.FileDescriptor

var file_station_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x1d, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x33, 0x0a, 0x17, 0x4d, 0x61, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x41, 0x67, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x29, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x20, 0x0a, 0x08, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x63, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x64,
	0x52, 0x65, 0x74, 0x22, 0xcd, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x74, 0x12, 0x37, 0x0a, 0x07, 0x6d, 0x6d, 0x61, 0x73, 0x72, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x78, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x67, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x52,
	0x65, 0x74, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x6d, 0x61, 0x73, 0x72, 0x65, 0x74, 0x12, 0x29, 0x0a,
	0x06, 0x6d, 0x73, 0x67, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x6d, 0x73, 0x67, 0x72, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x62, 0x72, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x74, 0x48, 0x00, 0x52, 0x04, 0x62, 0x72, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x05,
	0x61, 0x62, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x63, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x64, 0x52, 0x65, 0x74, 0x48, 0x00, 0x52, 0x05,
	0x61, 0x62, 0x72, 0x65, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x6f, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x23, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x10, 0x01, 0x22, 0x26, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x8d, 0x01, 0x0a,
	0x0d, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d,
	0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x48,
	0x00, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a,
	0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x48, 0x01, 0x52, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7b, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x69, 0x6f, 0x6e, 0x73, 0x48,
	0x00, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3e, 0x0a, 0x15, 0x44, 0x65, 0x73,
	0x74, 0x72, 0x6f, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x48, 0x5a, 0x20, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x34, 0x31, 0x37, 0x39, 0x37, 0x2f, 0x6d,
	0x65, 0x6d, 0x70, 0x68, 0x69, 0x73, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x70, 0x62, 0xca, 0x02, 0x0d,
	0x4d, 0x65, 0x6d, 0x70, 0x68, 0x69, 0x73, 0x70, 0x68, 0x70, 0x5c, 0x50, 0x62, 0xe2, 0x02, 0x13,
	0x4d, 0x65, 0x6d, 0x70, 0x68, 0x69, 0x73, 0x70, 0x68, 0x70, 0x5c, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_station_proto_rawDescOnce sync.Once
	file_station_proto_rawDescData = file_station_proto_rawDesc
)

func file_station_proto_rawDescGZIP() []byte {
	file_station_proto_rawDescOnce.Do(func() {
		file_station_proto_rawDescData = protoimpl.X.CompressGZIP(file_station_proto_rawDescData)
	})
	return file_station_proto_rawDescData
}

var file_station_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_station_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_station_proto_goTypes = []interface{}{
	(StorageOpt_StorageType)(0),     // 0: pb.StorageOpt.StorageType
	(*Station)(nil),                 // 1: pb.Station
	(*MaxMessageAgeSecondsRet)(nil), // 2: pb.MaxMessageAgeSecondsRet
	(*MessagesRet)(nil),             // 3: pb.MessagesRet
	(*BytesRet)(nil),                // 4: pb.BytesRet
	(*AckBasedRet)(nil),             // 5: pb.AckBasedRet
	(*RetentionOpt)(nil),            // 6: pb.RetentionOpt
	(*StorageOpt)(nil),              // 7: pb.StorageOpt
	(*PartitionOpt)(nil),            // 8: pb.PartitionOpt
	(*StationOpions)(nil),           // 9: pb.StationOpions
	(*CreateStationRequest)(nil),    // 10: pb.CreateStationRequest
	(*DestroyStationRequest)(nil),   // 11: pb.DestroyStationRequest
}
var file_station_proto_depIdxs = []int32{
	2,  // 0: pb.RetentionOpt.mmasret:type_name -> pb.MaxMessageAgeSecondsRet
	3,  // 1: pb.RetentionOpt.msgret:type_name -> pb.MessagesRet
	4,  // 2: pb.RetentionOpt.bret:type_name -> pb.BytesRet
	5,  // 3: pb.RetentionOpt.abret:type_name -> pb.AckBasedRet
	0,  // 4: pb.StorageOpt.storageType:type_name -> pb.StorageOpt.StorageType
	7,  // 5: pb.StationOpions.storage:type_name -> pb.StorageOpt
	6,  // 6: pb.StationOpions.retention:type_name -> pb.RetentionOpt
	1,  // 7: pb.CreateStationRequest.station:type_name -> pb.Station
	9,  // 8: pb.CreateStationRequest.options:type_name -> pb.StationOpions
	1,  // 9: pb.DestroyStationRequest.station:type_name -> pb.Station
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_station_proto_init() }
func file_station_proto_init() {
	if File_station_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_station_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Station); i {
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
		file_station_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxMessageAgeSecondsRet); i {
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
		file_station_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagesRet); i {
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
		file_station_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesRet); i {
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
		file_station_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckBasedRet); i {
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
		file_station_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetentionOpt); i {
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
		file_station_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageOpt); i {
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
		file_station_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartitionOpt); i {
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
		file_station_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationOpions); i {
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
		file_station_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStationRequest); i {
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
		file_station_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyStationRequest); i {
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
	file_station_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*RetentionOpt_Mmasret)(nil),
		(*RetentionOpt_Msgret)(nil),
		(*RetentionOpt_Bret)(nil),
		(*RetentionOpt_Abret)(nil),
	}
	file_station_proto_msgTypes[8].OneofWrappers = []interface{}{}
	file_station_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_station_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_station_proto_goTypes,
		DependencyIndexes: file_station_proto_depIdxs,
		EnumInfos:         file_station_proto_enumTypes,
		MessageInfos:      file_station_proto_msgTypes,
	}.Build()
	File_station_proto = out.File
	file_station_proto_rawDesc = nil
	file_station_proto_goTypes = nil
	file_station_proto_depIdxs = nil
}
