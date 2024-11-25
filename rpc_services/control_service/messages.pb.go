// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: control_service/messages.proto

package control_service

import (
	common_types "github.com/enbility/eebus-grpc/rpc_services/common_types"
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

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	mi := &file_control_service_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_service_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_control_service_messages_proto_rawDescGZIP(), []int{0}
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	mi := &file_control_service_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_control_service_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_control_service_messages_proto_rawDescGZIP(), []int{1}
}

type SetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port                    uint32                `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	VendorCode              string                `protobuf:"bytes,2,opt,name=vendor_code,json=vendorCode,proto3" json:"vendor_code,omitempty"`
	DeviceBrand             string                `protobuf:"bytes,3,opt,name=device_brand,json=deviceBrand,proto3" json:"device_brand,omitempty"`
	DeviceModel             string                `protobuf:"bytes,4,opt,name=device_model,json=deviceModel,proto3" json:"device_model,omitempty"`
	SerialNumber            string                `protobuf:"bytes,5,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	DeviceCategories        []DeviceCategory_Enum `protobuf:"varint,6,rep,packed,name=device_categories,json=deviceCategories,proto3,enum=control_service.DeviceCategory_Enum" json:"device_categories,omitempty"`
	DeviceType              DeviceType_Enum       `protobuf:"varint,7,opt,name=device_type,json=deviceType,proto3,enum=control_service.DeviceType_Enum" json:"device_type,omitempty"`
	EntityTypes             []EntityType_Enum     `protobuf:"varint,8,rep,packed,name=entity_types,json=entityTypes,proto3,enum=control_service.EntityType_Enum" json:"entity_types,omitempty"`
	HeartbeatTimeoutSeconds uint32                `protobuf:"varint,9,opt,name=heartbeat_timeout_seconds,json=heartbeatTimeoutSeconds,proto3" json:"heartbeat_timeout_seconds,omitempty"`
}

func (x *SetConfigRequest) Reset() {
	*x = SetConfigRequest{}
	mi := &file_control_service_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConfigRequest) ProtoMessage() {}

func (x *SetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_service_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetConfigRequest.ProtoReflect.Descriptor instead.
func (*SetConfigRequest) Descriptor() ([]byte, []int) {
	return file_control_service_messages_proto_rawDescGZIP(), []int{2}
}

func (x *SetConfigRequest) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *SetConfigRequest) GetVendorCode() string {
	if x != nil {
		return x.VendorCode
	}
	return ""
}

func (x *SetConfigRequest) GetDeviceBrand() string {
	if x != nil {
		return x.DeviceBrand
	}
	return ""
}

func (x *SetConfigRequest) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

func (x *SetConfigRequest) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *SetConfigRequest) GetDeviceCategories() []DeviceCategory_Enum {
	if x != nil {
		return x.DeviceCategories
	}
	return nil
}

func (x *SetConfigRequest) GetDeviceType() DeviceType_Enum {
	if x != nil {
		return x.DeviceType
	}
	return DeviceType_UNKNOWN
}

func (x *SetConfigRequest) GetEntityTypes() []EntityType_Enum {
	if x != nil {
		return x.EntityTypes
	}
	return nil
}

func (x *SetConfigRequest) GetHeartbeatTimeoutSeconds() uint32 {
	if x != nil {
		return x.HeartbeatTimeoutSeconds
	}
	return 0
}

type AddEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    EntityType_Enum             `protobuf:"varint,1,opt,name=type,proto3,enum=control_service.EntityType_Enum" json:"type,omitempty"`
	Address *common_types.EntityAddress `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AddEntityRequest) Reset() {
	*x = AddEntityRequest{}
	mi := &file_control_service_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEntityRequest) ProtoMessage() {}

func (x *AddEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_service_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEntityRequest.ProtoReflect.Descriptor instead.
func (*AddEntityRequest) Descriptor() ([]byte, []int) {
	return file_control_service_messages_proto_rawDescGZIP(), []int{3}
}

func (x *AddEntityRequest) GetType() EntityType_Enum {
	if x != nil {
		return x.Type
	}
	return EntityType_UNKNOWN
}

func (x *AddEntityRequest) GetAddress() *common_types.EntityAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

type RemoveEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *common_types.EntityAddress `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RemoveEntityRequest) Reset() {
	*x = RemoveEntityRequest{}
	mi := &file_control_service_messages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveEntityRequest) ProtoMessage() {}

func (x *RemoveEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_service_messages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveEntityRequest.ProtoReflect.Descriptor instead.
func (*RemoveEntityRequest) Descriptor() ([]byte, []int) {
	return file_control_service_messages_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveEntityRequest) GetAddress() *common_types.EntityAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

type AddUseCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityAddress *common_types.EntityAddress `protobuf:"bytes,1,opt,name=entity_address,json=entityAddress,proto3" json:"entity_address,omitempty"`
	UseCase       *UseCase                    `protobuf:"bytes,2,opt,name=use_case,json=useCase,proto3" json:"use_case,omitempty"`
}

func (x *AddUseCaseRequest) Reset() {
	*x = AddUseCaseRequest{}
	mi := &file_control_service_messages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddUseCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUseCaseRequest) ProtoMessage() {}

func (x *AddUseCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_service_messages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUseCaseRequest.ProtoReflect.Descriptor instead.
func (*AddUseCaseRequest) Descriptor() ([]byte, []int) {
	return file_control_service_messages_proto_rawDescGZIP(), []int{5}
}

func (x *AddUseCaseRequest) GetEntityAddress() *common_types.EntityAddress {
	if x != nil {
		return x.EntityAddress
	}
	return nil
}

func (x *AddUseCaseRequest) GetUseCase() *UseCase {
	if x != nil {
		return x.UseCase
	}
	return nil
}

type AddUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *AddUseCaseResponse) Reset() {
	*x = AddUseCaseResponse{}
	mi := &file_control_service_messages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUseCaseResponse) ProtoMessage() {}

func (x *AddUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_control_service_messages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUseCaseResponse.ProtoReflect.Descriptor instead.
func (*AddUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_control_service_messages_proto_rawDescGZIP(), []int{6}
}

func (x *AddUseCaseResponse) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type RegisterRemoteSkiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteSki string `protobuf:"bytes,1,opt,name=remote_ski,json=remoteSki,proto3" json:"remote_ski,omitempty"`
}

func (x *RegisterRemoteSkiRequest) Reset() {
	*x = RegisterRemoteSkiRequest{}
	mi := &file_control_service_messages_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRemoteSkiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRemoteSkiRequest) ProtoMessage() {}

func (x *RegisterRemoteSkiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_service_messages_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRemoteSkiRequest.ProtoReflect.Descriptor instead.
func (*RegisterRemoteSkiRequest) Descriptor() ([]byte, []int) {
	return file_control_service_messages_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterRemoteSkiRequest) GetRemoteSki() string {
	if x != nil {
		return x.RemoteSki
	}
	return ""
}

type SubscribeUseCaseEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityAddress *common_types.EntityAddress `protobuf:"bytes,1,opt,name=entity_address,json=entityAddress,proto3" json:"entity_address,omitempty"`
	UseCase       *UseCase                    `protobuf:"bytes,2,opt,name=use_case,json=useCase,proto3" json:"use_case,omitempty"`
}

func (x *SubscribeUseCaseEventsRequest) Reset() {
	*x = SubscribeUseCaseEventsRequest{}
	mi := &file_control_service_messages_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeUseCaseEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeUseCaseEventsRequest) ProtoMessage() {}

func (x *SubscribeUseCaseEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_service_messages_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeUseCaseEventsRequest.ProtoReflect.Descriptor instead.
func (*SubscribeUseCaseEventsRequest) Descriptor() ([]byte, []int) {
	return file_control_service_messages_proto_rawDescGZIP(), []int{8}
}

func (x *SubscribeUseCaseEventsRequest) GetEntityAddress() *common_types.EntityAddress {
	if x != nil {
		return x.EntityAddress
	}
	return nil
}

func (x *SubscribeUseCaseEventsRequest) GetUseCase() *UseCase {
	if x != nil {
		return x.UseCase
	}
	return nil
}

type SubscribeUseCaseEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteSki           string                      `protobuf:"bytes,1,opt,name=remote_ski,json=remoteSki,proto3" json:"remote_ski,omitempty"`
	RemoteEntityAddress *common_types.EntityAddress `protobuf:"bytes,2,opt,name=remote_entity_address,json=remoteEntityAddress,proto3" json:"remote_entity_address,omitempty"`
	UseCaseEvent        *UseCaseEvent               `protobuf:"bytes,3,opt,name=use_case_event,json=useCaseEvent,proto3" json:"use_case_event,omitempty"`
}

func (x *SubscribeUseCaseEventsResponse) Reset() {
	*x = SubscribeUseCaseEventsResponse{}
	mi := &file_control_service_messages_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeUseCaseEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeUseCaseEventsResponse) ProtoMessage() {}

func (x *SubscribeUseCaseEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_control_service_messages_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeUseCaseEventsResponse.ProtoReflect.Descriptor instead.
func (*SubscribeUseCaseEventsResponse) Descriptor() ([]byte, []int) {
	return file_control_service_messages_proto_rawDescGZIP(), []int{9}
}

func (x *SubscribeUseCaseEventsResponse) GetRemoteSki() string {
	if x != nil {
		return x.RemoteSki
	}
	return ""
}

func (x *SubscribeUseCaseEventsResponse) GetRemoteEntityAddress() *common_types.EntityAddress {
	if x != nil {
		return x.RemoteEntityAddress
	}
	return nil
}

func (x *SubscribeUseCaseEventsResponse) GetUseCaseEvent() *UseCaseEvent {
	if x != nil {
		return x.UseCaseEvent
	}
	return nil
}

var File_control_service_messages_proto protoreflect.FileDescriptor

var file_control_service_messages_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x1b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc9, 0x03, 0x0a, 0x10, 0x53, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x51,
	0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52,
	0x10, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0b, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x19, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x68, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x7f, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x35, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4c, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x43,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x0d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x07, 0x75, 0x73, 0x65, 0x43,
	0x61, 0x73, 0x65, 0x22, 0x30, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x39, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x6b, 0x69,
	0x22, 0x98, 0x01, 0x0a, 0x1d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73,
	0x65, 0x43, 0x61, 0x73, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x43, 0x61,
	0x73, 0x65, 0x52, 0x07, 0x75, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x22, 0xd5, 0x01, 0x0a, 0x1e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x12, 0x4f, 0x0a,
	0x15, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x43,
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x45, 0x56, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_control_service_messages_proto_rawDescOnce sync.Once
	file_control_service_messages_proto_rawDescData = file_control_service_messages_proto_rawDesc
)

func file_control_service_messages_proto_rawDescGZIP() []byte {
	file_control_service_messages_proto_rawDescOnce.Do(func() {
		file_control_service_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_control_service_messages_proto_rawDescData)
	})
	return file_control_service_messages_proto_rawDescData
}

var file_control_service_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_control_service_messages_proto_goTypes = []any{
	(*EmptyRequest)(nil),                   // 0: control_service.EmptyRequest
	(*EmptyResponse)(nil),                  // 1: control_service.EmptyResponse
	(*SetConfigRequest)(nil),               // 2: control_service.SetConfigRequest
	(*AddEntityRequest)(nil),               // 3: control_service.AddEntityRequest
	(*RemoveEntityRequest)(nil),            // 4: control_service.RemoveEntityRequest
	(*AddUseCaseRequest)(nil),              // 5: control_service.AddUseCaseRequest
	(*AddUseCaseResponse)(nil),             // 6: control_service.AddUseCaseResponse
	(*RegisterRemoteSkiRequest)(nil),       // 7: control_service.RegisterRemoteSkiRequest
	(*SubscribeUseCaseEventsRequest)(nil),  // 8: control_service.SubscribeUseCaseEventsRequest
	(*SubscribeUseCaseEventsResponse)(nil), // 9: control_service.SubscribeUseCaseEventsResponse
	(DeviceCategory_Enum)(0),               // 10: control_service.DeviceCategory.Enum
	(DeviceType_Enum)(0),                   // 11: control_service.DeviceType.Enum
	(EntityType_Enum)(0),                   // 12: control_service.EntityType.Enum
	(*common_types.EntityAddress)(nil),     // 13: common_types.EntityAddress
	(*UseCase)(nil),                        // 14: control_service.UseCase
	(*UseCaseEvent)(nil),                   // 15: control_service.UseCaseEvent
}
var file_control_service_messages_proto_depIdxs = []int32{
	10, // 0: control_service.SetConfigRequest.device_categories:type_name -> control_service.DeviceCategory.Enum
	11, // 1: control_service.SetConfigRequest.device_type:type_name -> control_service.DeviceType.Enum
	12, // 2: control_service.SetConfigRequest.entity_types:type_name -> control_service.EntityType.Enum
	12, // 3: control_service.AddEntityRequest.type:type_name -> control_service.EntityType.Enum
	13, // 4: control_service.AddEntityRequest.address:type_name -> common_types.EntityAddress
	13, // 5: control_service.RemoveEntityRequest.address:type_name -> common_types.EntityAddress
	13, // 6: control_service.AddUseCaseRequest.entity_address:type_name -> common_types.EntityAddress
	14, // 7: control_service.AddUseCaseRequest.use_case:type_name -> control_service.UseCase
	13, // 8: control_service.SubscribeUseCaseEventsRequest.entity_address:type_name -> common_types.EntityAddress
	14, // 9: control_service.SubscribeUseCaseEventsRequest.use_case:type_name -> control_service.UseCase
	13, // 10: control_service.SubscribeUseCaseEventsResponse.remote_entity_address:type_name -> common_types.EntityAddress
	15, // 11: control_service.SubscribeUseCaseEventsResponse.use_case_event:type_name -> control_service.UseCaseEvent
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_control_service_messages_proto_init() }
func file_control_service_messages_proto_init() {
	if File_control_service_messages_proto != nil {
		return
	}
	file_control_service_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_control_service_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_control_service_messages_proto_goTypes,
		DependencyIndexes: file_control_service_messages_proto_depIdxs,
		MessageInfos:      file_control_service_messages_proto_msgTypes,
	}.Build()
	File_control_service_messages_proto = out.File
	file_control_service_messages_proto_rawDesc = nil
	file_control_service_messages_proto_goTypes = nil
	file_control_service_messages_proto_depIdxs = nil
}
