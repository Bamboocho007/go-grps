// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: weather/v1/weather.proto

package weatherpb

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

type GetWeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Degrees int32  `protobuf:"varint,4,opt,name=degrees,proto3" json:"degrees,omitempty"`
}

func (x *GetWeatherRequest) Reset() {
	*x = GetWeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_v1_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherRequest) ProtoMessage() {}

func (x *GetWeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_v1_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherRequest.ProtoReflect.Descriptor instead.
func (*GetWeatherRequest) Descriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{0}
}

func (x *GetWeatherRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetWeatherRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *GetWeatherRequest) GetDegrees() int32 {
	if x != nil {
		return x.Degrees
	}
	return 0
}

type GetWeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Degrees int32  `protobuf:"varint,4,opt,name=degrees,proto3" json:"degrees,omitempty"`
}

func (x *GetWeatherResponse) Reset() {
	*x = GetWeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_v1_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherResponse) ProtoMessage() {}

func (x *GetWeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_v1_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherResponse.ProtoReflect.Descriptor instead.
func (*GetWeatherResponse) Descriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{1}
}

func (x *GetWeatherResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetWeatherResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetWeatherResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *GetWeatherResponse) GetDegrees() int32 {
	if x != nil {
		return x.Degrees
	}
	return 0
}

type PutWeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Degrees int32  `protobuf:"varint,4,opt,name=degrees,proto3" json:"degrees,omitempty"`
}

func (x *PutWeatherRequest) Reset() {
	*x = PutWeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_v1_weather_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutWeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutWeatherRequest) ProtoMessage() {}

func (x *PutWeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_v1_weather_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutWeatherRequest.ProtoReflect.Descriptor instead.
func (*PutWeatherRequest) Descriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{2}
}

func (x *PutWeatherRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *PutWeatherRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *PutWeatherRequest) GetDegrees() int32 {
	if x != nil {
		return x.Degrees
	}
	return 0
}

type PutWeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Degrees int32  `protobuf:"varint,4,opt,name=degrees,proto3" json:"degrees,omitempty"`
}

func (x *PutWeatherResponse) Reset() {
	*x = PutWeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_v1_weather_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutWeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutWeatherResponse) ProtoMessage() {}

func (x *PutWeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_v1_weather_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutWeatherResponse.ProtoReflect.Descriptor instead.
func (*PutWeatherResponse) Descriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{3}
}

func (x *PutWeatherResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PutWeatherResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *PutWeatherResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *PutWeatherResponse) GetDegrees() int32 {
	if x != nil {
		return x.Degrees
	}
	return 0
}

type DeleteWeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteWeatherRequest) Reset() {
	*x = DeleteWeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_v1_weather_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWeatherRequest) ProtoMessage() {}

func (x *DeleteWeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_v1_weather_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWeatherRequest.ProtoReflect.Descriptor instead.
func (*DeleteWeatherRequest) Descriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteWeatherRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteWeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteWeatherResponse) Reset() {
	*x = DeleteWeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_v1_weather_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWeatherResponse) ProtoMessage() {}

func (x *DeleteWeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_v1_weather_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWeatherResponse.ProtoReflect.Descriptor instead.
func (*DeleteWeatherResponse) Descriptor() ([]byte, []int) {
	return file_weather_v1_weather_proto_rawDescGZIP(), []int{5}
}

var File_weather_v1_weather_proto protoreflect.FileDescriptor

var file_weather_v1_weather_proto_rawDesc = []byte{
	0x0a, 0x18, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x5b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x67,
	0x72, 0x65, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x67, 0x72,
	0x65, 0x65, 0x73, 0x22, 0x6c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x67, 0x72, 0x65,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65,
	0x73, 0x22, 0x5b, 0x0a, 0x11, 0x50, 0x75, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x73, 0x22, 0x6c,
	0x0a, 0x12, 0x50, 0x75, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8b, 0x02,
	0x0a, 0x13, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x12, 0x1d, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x75, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x75, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x6d, 0x62, 0x6f, 0x6f,
	0x63, 0x68, 0x6f, 0x30, 0x30, 0x37, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_weather_v1_weather_proto_rawDescOnce sync.Once
	file_weather_v1_weather_proto_rawDescData = file_weather_v1_weather_proto_rawDesc
)

func file_weather_v1_weather_proto_rawDescGZIP() []byte {
	file_weather_v1_weather_proto_rawDescOnce.Do(func() {
		file_weather_v1_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_v1_weather_proto_rawDescData)
	})
	return file_weather_v1_weather_proto_rawDescData
}

var file_weather_v1_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_weather_v1_weather_proto_goTypes = []interface{}{
	(*GetWeatherRequest)(nil),     // 0: weather.v1.GetWeatherRequest
	(*GetWeatherResponse)(nil),    // 1: weather.v1.GetWeatherResponse
	(*PutWeatherRequest)(nil),     // 2: weather.v1.PutWeatherRequest
	(*PutWeatherResponse)(nil),    // 3: weather.v1.PutWeatherResponse
	(*DeleteWeatherRequest)(nil),  // 4: weather.v1.DeleteWeatherRequest
	(*DeleteWeatherResponse)(nil), // 5: weather.v1.DeleteWeatherResponse
}
var file_weather_v1_weather_proto_depIdxs = []int32{
	0, // 0: weather.v1.WeatherStoreService.GetWeather:input_type -> weather.v1.GetWeatherRequest
	2, // 1: weather.v1.WeatherStoreService.PutWeather:input_type -> weather.v1.PutWeatherRequest
	4, // 2: weather.v1.WeatherStoreService.DeleteWeather:input_type -> weather.v1.DeleteWeatherRequest
	1, // 3: weather.v1.WeatherStoreService.GetWeather:output_type -> weather.v1.GetWeatherResponse
	3, // 4: weather.v1.WeatherStoreService.PutWeather:output_type -> weather.v1.PutWeatherResponse
	5, // 5: weather.v1.WeatherStoreService.DeleteWeather:output_type -> weather.v1.DeleteWeatherResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_weather_v1_weather_proto_init() }
func file_weather_v1_weather_proto_init() {
	if File_weather_v1_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_v1_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeatherRequest); i {
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
		file_weather_v1_weather_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeatherResponse); i {
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
		file_weather_v1_weather_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutWeatherRequest); i {
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
		file_weather_v1_weather_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutWeatherResponse); i {
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
		file_weather_v1_weather_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWeatherRequest); i {
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
		file_weather_v1_weather_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWeatherResponse); i {
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
			RawDescriptor: file_weather_v1_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_v1_weather_proto_goTypes,
		DependencyIndexes: file_weather_v1_weather_proto_depIdxs,
		MessageInfos:      file_weather_v1_weather_proto_msgTypes,
	}.Build()
	File_weather_v1_weather_proto = out.File
	file_weather_v1_weather_proto_rawDesc = nil
	file_weather_v1_weather_proto_goTypes = nil
	file_weather_v1_weather_proto_depIdxs = nil
}
