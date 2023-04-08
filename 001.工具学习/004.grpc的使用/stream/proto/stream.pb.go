// proto 文件的解析语法, 不指定默认为 proto2, proto2 的协议比 proto3 的协议更复杂

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: grpc/stream/proto/stream.proto

// grpc 代码中 servername 的一部分(前缀)， 相当于命名空间

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Gender int32

const (
	// 男
	Gender_MALE Gender = 0 // 在 proto3 语法里， 第一个值必须时 0
	// 女
	Gender_FEMALE Gender = 1
	// 未知
	Gender_UNKNOWN Gender = 3
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
		3: "UNKNOWN",
	}
	Gender_value = map[string]int32{
		"MALE":    0,
		"FEMALE":  1,
		"UNKNOWN": 3,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_stream_proto_stream_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_grpc_stream_proto_stream_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_grpc_stream_proto_stream_proto_rawDescGZIP(), []int{0}
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provice string `protobuf:"bytes,1,opt,name=provice,proto3" json:"provice,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_stream_proto_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_stream_proto_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_grpc_stream_proto_stream_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetProvice() string {
	if x != nil {
		return x.Provice
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [1-15] 占用一个字节， 数字越大，占用越多， 在使用时尽量不超过15个字段
	// 如果大于15个字段，尽量将使用频繁的字段写在前面， 也有几率构建小的占用
	Name     string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 字段的标志
	Gender   Gender                 `protobuf:"varint,2,opt,name=gender,proto3,enum=stream.Gender" json:"gender,omitempty"`
	Age      uint32                 `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Birthday *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"` // 时间
	Addr     *Address               `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr,omitempty"`
	Hoppy    []string               `protobuf:"bytes,7,rep,name=hoppy,proto3" json:"hoppy,omitempty"` // repeated 会被转化为数组
	Data     map[string]*anypb.Any  `protobuf:"bytes,8,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_stream_proto_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_stream_proto_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_grpc_stream_proto_stream_proto_rawDescGZIP(), []int{1}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_MALE
}

func (x *HelloRequest) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *HelloRequest) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *HelloRequest) GetAddr() *Address {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *HelloRequest) GetHoppy() []string {
	if x != nil {
		return x.Hoppy
	}
	return nil
}

func (x *HelloRequest) GetData() map[string]*anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type HelloReplay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReplay) Reset() {
	*x = HelloReplay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_stream_proto_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReplay) ProtoMessage() {}

func (x *HelloReplay) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_stream_proto_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReplay.ProtoReflect.Descriptor instead.
func (*HelloReplay) Descriptor() ([]byte, []int) {
	return file_grpc_stream_proto_stream_proto_rawDescGZIP(), []int{2}
}

func (x *HelloReplay) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_stream_proto_stream_proto protoreflect.FileDescriptor

var file_grpc_stream_proto_stream_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0xf4, 0x02,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x12, 0x23, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x70,
	0x70, 0x79, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x70, 0x70, 0x79, 0x12,
	0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x4d, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x4a, 0x04, 0x08, 0x64, 0x10, 0x65, 0x4a, 0x04, 0x08, 0x67, 0x10, 0x68, 0x4a, 0x06,
	0x08, 0x96, 0x01, 0x10, 0xc9, 0x01, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x27, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x2b, 0x0a,
	0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x32, 0x99, 0x02, 0x0a, 0x07, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x14, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x14, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x61, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x45, 0x0a, 0x14, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x14,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x47, 0x0a,
	0x14, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x54, 0x77, 0x6f, 0x57, 0x61, 0x79, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x18, 0x5a, 0x16, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_stream_proto_stream_proto_rawDescOnce sync.Once
	file_grpc_stream_proto_stream_proto_rawDescData = file_grpc_stream_proto_stream_proto_rawDesc
)

func file_grpc_stream_proto_stream_proto_rawDescGZIP() []byte {
	file_grpc_stream_proto_stream_proto_rawDescOnce.Do(func() {
		file_grpc_stream_proto_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_stream_proto_stream_proto_rawDescData)
	})
	return file_grpc_stream_proto_stream_proto_rawDescData
}

var file_grpc_stream_proto_stream_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_stream_proto_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_stream_proto_stream_proto_goTypes = []interface{}{
	(Gender)(0),                   // 0: stream.Gender
	(*Address)(nil),               // 1: stream.Address
	(*HelloRequest)(nil),          // 2: stream.HelloRequest
	(*HelloReplay)(nil),           // 3: stream.HelloReplay
	nil,                           // 4: stream.HelloRequest.DataEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 6: google.protobuf.Any
}
var file_grpc_stream_proto_stream_proto_depIdxs = []int32{
	0, // 0: stream.HelloRequest.gender:type_name -> stream.Gender
	5, // 1: stream.HelloRequest.birthday:type_name -> google.protobuf.Timestamp
	1, // 2: stream.HelloRequest.addr:type_name -> stream.Address
	4, // 3: stream.HelloRequest.data:type_name -> stream.HelloRequest.DataEntry
	6, // 4: stream.HelloRequest.DataEntry.value:type_name -> google.protobuf.Any
	2, // 5: stream.Greeter.SayHello:input_type -> stream.HelloRequest
	2, // 6: stream.Greeter.SayHelloClientStream:input_type -> stream.HelloRequest
	2, // 7: stream.Greeter.SayHelloServerStream:input_type -> stream.HelloRequest
	2, // 8: stream.Greeter.SayHelloTwoWayStream:input_type -> stream.HelloRequest
	3, // 9: stream.Greeter.SayHello:output_type -> stream.HelloReplay
	3, // 10: stream.Greeter.SayHelloClientStream:output_type -> stream.HelloReplay
	3, // 11: stream.Greeter.SayHelloServerStream:output_type -> stream.HelloReplay
	3, // 12: stream.Greeter.SayHelloTwoWayStream:output_type -> stream.HelloReplay
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grpc_stream_proto_stream_proto_init() }
func file_grpc_stream_proto_stream_proto_init() {
	if File_grpc_stream_proto_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_stream_proto_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_grpc_stream_proto_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_grpc_stream_proto_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReplay); i {
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
			RawDescriptor: file_grpc_stream_proto_stream_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_stream_proto_stream_proto_goTypes,
		DependencyIndexes: file_grpc_stream_proto_stream_proto_depIdxs,
		EnumInfos:         file_grpc_stream_proto_stream_proto_enumTypes,
		MessageInfos:      file_grpc_stream_proto_stream_proto_msgTypes,
	}.Build()
	File_grpc_stream_proto_stream_proto = out.File
	file_grpc_stream_proto_stream_proto_rawDesc = nil
	file_grpc_stream_proto_stream_proto_goTypes = nil
	file_grpc_stream_proto_stream_proto_depIdxs = nil
}
