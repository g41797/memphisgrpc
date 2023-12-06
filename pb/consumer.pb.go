// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: consumer.proto

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

type Consumer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Consumer) Reset() {
	*x = Consumer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consumer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consumer) ProtoMessage() {}

func (x *Consumer) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consumer.ProtoReflect.Descriptor instead.
func (*Consumer) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{0}
}

func (x *Consumer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Station  *Station  `protobuf:"bytes,1,opt,name=station,proto3" json:"station,omitempty"`
	Consumer *Consumer `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
}

func (x *CreateConsumerRequest) Reset() {
	*x = CreateConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConsumerRequest) ProtoMessage() {}

func (x *CreateConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConsumerRequest.ProtoReflect.Descriptor instead.
func (*CreateConsumerRequest) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateConsumerRequest) GetStation() *Station {
	if x != nil {
		return x.Station
	}
	return nil
}

func (x *CreateConsumerRequest) GetConsumer() *Consumer {
	if x != nil {
		return x.Consumer
	}
	return nil
}

type ConsumeMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*ConsumeMessages_Start
	//	*ConsumeMessages_Stop
	Data isConsumeMessages_Data `protobuf_oneof:"data"`
}

func (x *ConsumeMessages) Reset() {
	*x = ConsumeMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeMessages) ProtoMessage() {}

func (x *ConsumeMessages) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeMessages.ProtoReflect.Descriptor instead.
func (*ConsumeMessages) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{2}
}

func (m *ConsumeMessages) GetData() isConsumeMessages_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ConsumeMessages) GetStart() *CreateConsumerRequest {
	if x, ok := x.GetData().(*ConsumeMessages_Start); ok {
		return x.Start
	}
	return nil
}

func (x *ConsumeMessages) GetStop() *Stop {
	if x, ok := x.GetData().(*ConsumeMessages_Stop); ok {
		return x.Stop
	}
	return nil
}

type isConsumeMessages_Data interface {
	isConsumeMessages_Data()
}

type ConsumeMessages_Start struct {
	Start *CreateConsumerRequest `protobuf:"bytes,1,opt,name=start,proto3,oneof"`
}

type ConsumeMessages_Stop struct {
	Stop *Stop `protobuf:"bytes,2,opt,name=stop,proto3,oneof"`
}

func (*ConsumeMessages_Start) isConsumeMessages_Data() {}

func (*ConsumeMessages_Stop) isConsumeMessages_Data() {}

type Wakeup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Wakeup) Reset() {
	*x = Wakeup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wakeup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wakeup) ProtoMessage() {}

func (x *Wakeup) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wakeup.ProtoReflect.Descriptor instead.
func (*Wakeup) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{3}
}

type ConsumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*ConsumeResponse_Msg
	//	*ConsumeResponse_Status
	//	*ConsumeResponse_Wakeup
	Data isConsumeResponse_Data `protobuf_oneof:"data"`
}

func (x *ConsumeResponse) Reset() {
	*x = ConsumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeResponse) ProtoMessage() {}

func (x *ConsumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeResponse.ProtoReflect.Descriptor instead.
func (*ConsumeResponse) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{4}
}

func (m *ConsumeResponse) GetData() isConsumeResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ConsumeResponse) GetMsg() *Msg {
	if x, ok := x.GetData().(*ConsumeResponse_Msg); ok {
		return x.Msg
	}
	return nil
}

func (x *ConsumeResponse) GetStatus() *Status {
	if x, ok := x.GetData().(*ConsumeResponse_Status); ok {
		return x.Status
	}
	return nil
}

func (x *ConsumeResponse) GetWakeup() *Wakeup {
	if x, ok := x.GetData().(*ConsumeResponse_Wakeup); ok {
		return x.Wakeup
	}
	return nil
}

type isConsumeResponse_Data interface {
	isConsumeResponse_Data()
}

type ConsumeResponse_Msg struct {
	Msg *Msg `protobuf:"bytes,1,opt,name=msg,proto3,oneof"`
}

type ConsumeResponse_Status struct {
	Status *Status `protobuf:"bytes,2,opt,name=status,proto3,oneof"`
}

type ConsumeResponse_Wakeup struct {
	Wakeup *Wakeup `protobuf:"bytes,3,opt,name=wakeup,proto3,oneof"`
}

func (*ConsumeResponse_Msg) isConsumeResponse_Data() {}

func (*ConsumeResponse_Status) isConsumeResponse_Data() {}

func (*ConsumeResponse_Wakeup) isConsumeResponse_Data() {}

var File_consumer_proto protoreflect.FileDescriptor

var file_consumer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x73, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x6d,
	0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x68, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x22, 0x6c, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x70,
	0x48, 0x00, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x08, 0x0a, 0x06, 0x57, 0x61, 0x6b, 0x65, 0x75, 0x70, 0x22, 0x82, 0x01, 0x0a, 0x0f, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x73, 0x67, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x24, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x24, 0x0a, 0x06, 0x77, 0x61, 0x6b, 0x65, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x61, 0x6b, 0x65, 0x75, 0x70, 0x48, 0x00, 0x52,
	0x06, 0x77, 0x61, 0x6b, 0x65, 0x75, 0x70, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x54, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x34,
	0x31, 0x37, 0x39, 0x37, 0x2f, 0x6d, 0x65, 0x6d, 0x70, 0x68, 0x69, 0x73, 0x67, 0x72, 0x70, 0x63,
	0x3b, 0x70, 0x62, 0xca, 0x02, 0x13, 0x67, 0x34, 0x31, 0x39, 0x37, 0x5c, 0x6d, 0x65, 0x6d, 0x70,
	0x68, 0x69, 0x73, 0x70, 0x68, 0x70, 0x5c, 0x70, 0x62, 0xe2, 0x02, 0x19, 0x67, 0x34, 0x31, 0x39,
	0x37, 0x5c, 0x6d, 0x65, 0x6d, 0x70, 0x68, 0x69, 0x73, 0x70, 0x68, 0x70, 0x5c, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consumer_proto_rawDescOnce sync.Once
	file_consumer_proto_rawDescData = file_consumer_proto_rawDesc
)

func file_consumer_proto_rawDescGZIP() []byte {
	file_consumer_proto_rawDescOnce.Do(func() {
		file_consumer_proto_rawDescData = protoimpl.X.CompressGZIP(file_consumer_proto_rawDescData)
	})
	return file_consumer_proto_rawDescData
}

var file_consumer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_consumer_proto_goTypes = []interface{}{
	(*Consumer)(nil),              // 0: pb.Consumer
	(*CreateConsumerRequest)(nil), // 1: pb.CreateConsumerRequest
	(*ConsumeMessages)(nil),       // 2: pb.ConsumeMessages
	(*Wakeup)(nil),                // 3: pb.Wakeup
	(*ConsumeResponse)(nil),       // 4: pb.ConsumeResponse
	(*Station)(nil),               // 5: pb.Station
	(*Stop)(nil),                  // 6: pb.Stop
	(*Msg)(nil),                   // 7: pb.Msg
	(*Status)(nil),                // 8: pb.Status
}
var file_consumer_proto_depIdxs = []int32{
	5, // 0: pb.CreateConsumerRequest.station:type_name -> pb.Station
	0, // 1: pb.CreateConsumerRequest.consumer:type_name -> pb.Consumer
	1, // 2: pb.ConsumeMessages.start:type_name -> pb.CreateConsumerRequest
	6, // 3: pb.ConsumeMessages.stop:type_name -> pb.Stop
	7, // 4: pb.ConsumeResponse.msg:type_name -> pb.Msg
	8, // 5: pb.ConsumeResponse.status:type_name -> pb.Status
	3, // 6: pb.ConsumeResponse.wakeup:type_name -> pb.Wakeup
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_consumer_proto_init() }
func file_consumer_proto_init() {
	if File_consumer_proto != nil {
		return
	}
	file_status_proto_init()
	file_station_proto_init()
	file_stop_proto_init()
	file_msg_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_consumer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consumer); i {
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
		file_consumer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConsumerRequest); i {
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
		file_consumer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeMessages); i {
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
		file_consumer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wakeup); i {
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
		file_consumer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeResponse); i {
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
	file_consumer_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ConsumeMessages_Start)(nil),
		(*ConsumeMessages_Stop)(nil),
	}
	file_consumer_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ConsumeResponse_Msg)(nil),
		(*ConsumeResponse_Status)(nil),
		(*ConsumeResponse_Wakeup)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_consumer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_consumer_proto_goTypes,
		DependencyIndexes: file_consumer_proto_depIdxs,
		MessageInfos:      file_consumer_proto_msgTypes,
	}.Build()
	File_consumer_proto = out.File
	file_consumer_proto_rawDesc = nil
	file_consumer_proto_goTypes = nil
	file_consumer_proto_depIdxs = nil
}
