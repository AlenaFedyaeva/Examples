// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: greet/greetpb/greet.proto

package greetpb

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

// Unary example
type Greeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *Greeting) Reset() {
	*x = Greeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Greeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Greeting) ProtoMessage() {}

func (x *Greeting) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Greeting.ProtoReflect.Descriptor instead.
func (*Greeting) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{0}
}

func (x *Greeting) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Greeting) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{1}
}

func (x *GreetRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{2}
}

func (x *GreetResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// Stream example
type GreetingManyTimes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *GreetingManyTimes) Reset() {
	*x = GreetingManyTimes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingManyTimes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingManyTimes) ProtoMessage() {}

func (x *GreetingManyTimes) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingManyTimes.ProtoReflect.Descriptor instead.
func (*GreetingManyTimes) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{3}
}

func (x *GreetingManyTimes) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GreetingManyTimes) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type GreetRequestManyTimes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GreetingMT *GreetingManyTimes `protobuf:"bytes,1,opt,name=greetingMT,proto3" json:"greetingMT,omitempty"`
}

func (x *GreetRequestManyTimes) Reset() {
	*x = GreetRequestManyTimes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequestManyTimes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequestManyTimes) ProtoMessage() {}

func (x *GreetRequestManyTimes) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequestManyTimes.ProtoReflect.Descriptor instead.
func (*GreetRequestManyTimes) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{4}
}

func (x *GreetRequestManyTimes) GetGreetingMT() *GreetingManyTimes {
	if x != nil {
		return x.GreetingMT
	}
	return nil
}

type GreetResponseManyTimes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetResponseManyTimes) Reset() {
	*x = GreetResponseManyTimes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponseManyTimes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponseManyTimes) ProtoMessage() {}

func (x *GreetResponseManyTimes) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponseManyTimes.ProtoReflect.Descriptor instead.
func (*GreetResponseManyTimes) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{5}
}

func (x *GreetResponseManyTimes) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

//Client streaming
type GreetingCliStreaming struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamName string `protobuf:"bytes,1,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
}

func (x *GreetingCliStreaming) Reset() {
	*x = GreetingCliStreaming{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingCliStreaming) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingCliStreaming) ProtoMessage() {}

func (x *GreetingCliStreaming) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingCliStreaming.ProtoReflect.Descriptor instead.
func (*GreetingCliStreaming) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{6}
}

func (x *GreetingCliStreaming) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

type GreetRequestCliStreaming struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *GreetingCliStreaming `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetRequestCliStreaming) Reset() {
	*x = GreetRequestCliStreaming{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequestCliStreaming) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequestCliStreaming) ProtoMessage() {}

func (x *GreetRequestCliStreaming) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequestCliStreaming.ProtoReflect.Descriptor instead.
func (*GreetRequestCliStreaming) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{7}
}

func (x *GreetRequestCliStreaming) GetGreeting() *GreetingCliStreaming {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type GreetResponseCliStriaming struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetResponseCliStriaming) Reset() {
	*x = GreetResponseCliStriaming{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponseCliStriaming) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponseCliStriaming) ProtoMessage() {}

func (x *GreetResponseCliStriaming) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponseCliStriaming.ProtoReflect.Descriptor instead.
func (*GreetResponseCliStriaming) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{8}
}

func (x *GreetResponseCliStriaming) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// Bi Directional Streaming
type GreetRequestBiDirect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GreetingBd *Greeting `protobuf:"bytes,1,opt,name=greeting_bd,json=greetingBd,proto3" json:"greeting_bd,omitempty"`
}

func (x *GreetRequestBiDirect) Reset() {
	*x = GreetRequestBiDirect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequestBiDirect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequestBiDirect) ProtoMessage() {}

func (x *GreetRequestBiDirect) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequestBiDirect.ProtoReflect.Descriptor instead.
func (*GreetRequestBiDirect) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{9}
}

func (x *GreetRequestBiDirect) GetGreetingBd() *Greeting {
	if x != nil {
		return x.GreetingBd
	}
	return nil
}

type GreetResponseBiDirecrt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetResponseBiDirecrt) Reset() {
	*x = GreetResponseBiDirecrt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponseBiDirecrt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponseBiDirecrt) ProtoMessage() {}

func (x *GreetResponseBiDirecrt) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponseBiDirecrt.ProtoReflect.Descriptor instead.
func (*GreetResponseBiDirecrt) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{10}
}

func (x *GreetResponseBiDirecrt) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// Deadline
type GreetRequestDeadline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GreetingDeadline *Greeting `protobuf:"bytes,1,opt,name=greeting_deadline,json=greetingDeadline,proto3" json:"greeting_deadline,omitempty"`
}

func (x *GreetRequestDeadline) Reset() {
	*x = GreetRequestDeadline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequestDeadline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequestDeadline) ProtoMessage() {}

func (x *GreetRequestDeadline) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequestDeadline.ProtoReflect.Descriptor instead.
func (*GreetRequestDeadline) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{11}
}

func (x *GreetRequestDeadline) GetGreetingDeadline() *Greeting {
	if x != nil {
		return x.GreetingDeadline
	}
	return nil
}

type GreetResponseDeadline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetResponseDeadline) Reset() {
	*x = GreetResponseDeadline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_greet_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponseDeadline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponseDeadline) ProtoMessage() {}

func (x *GreetResponseDeadline) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_greet_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponseDeadline.ProtoReflect.Descriptor instead.
func (*GreetResponseDeadline) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_greet_proto_rawDescGZIP(), []int{12}
}

func (x *GreetResponseDeadline) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_greet_greetpb_greet_proto protoreflect.FileDescriptor

var file_greet_greetpb_greet_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2f,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x22, 0x46, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x0c, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x27, 0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x4f, 0x0a, 0x11, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x6e, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x51, 0x0a, 0x15, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4d,
	0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x0a, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x4d, 0x54, 0x22, 0x30, 0x0a, 0x16, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x37, 0x0a, 0x14, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x43, 0x6c, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x53, 0x0a, 0x18, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x6c, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x37, 0x0a, 0x08, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6c,
	0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x22, 0x33, 0x0a, 0x19, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6c, 0x69, 0x53, 0x74, 0x72, 0x69, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x48, 0x0a, 0x14, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x12, 0x30, 0x0a, 0x0b, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x42, 0x64, 0x22, 0x30, 0x0a, 0x16, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63, 0x72, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x54, 0x0a, 0x14, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3c, 0x0a,
	0x11, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x2f, 0x0a, 0x15, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xc9, 0x03, 0x0a,
	0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a,
	0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x52, 0x0a, 0x09, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4c,
	0x6f, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x1a, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6c, 0x69, 0x53, 0x74, 0x72,
	0x69, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x28, 0x01, 0x12, 0x51, 0x0a, 0x0d, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x69,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x72, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x37, 0x0a,
	0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x45, 0x72, 0x72, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x11, 0x47, 0x72, 0x65, 0x65, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greet_greetpb_greet_proto_rawDescOnce sync.Once
	file_greet_greetpb_greet_proto_rawDescData = file_greet_greetpb_greet_proto_rawDesc
)

func file_greet_greetpb_greet_proto_rawDescGZIP() []byte {
	file_greet_greetpb_greet_proto_rawDescOnce.Do(func() {
		file_greet_greetpb_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_greet_greetpb_greet_proto_rawDescData)
	})
	return file_greet_greetpb_greet_proto_rawDescData
}

var file_greet_greetpb_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_greet_greetpb_greet_proto_goTypes = []interface{}{
	(*Greeting)(nil),                  // 0: greet.Greeting
	(*GreetRequest)(nil),              // 1: greet.GreetRequest
	(*GreetResponse)(nil),             // 2: greet.GreetResponse
	(*GreetingManyTimes)(nil),         // 3: greet.GreetingManyTimes
	(*GreetRequestManyTimes)(nil),     // 4: greet.GreetRequestManyTimes
	(*GreetResponseManyTimes)(nil),    // 5: greet.GreetResponseManyTimes
	(*GreetingCliStreaming)(nil),      // 6: greet.GreetingCliStreaming
	(*GreetRequestCliStreaming)(nil),  // 7: greet.GreetRequestCliStreaming
	(*GreetResponseCliStriaming)(nil), // 8: greet.GreetResponseCliStriaming
	(*GreetRequestBiDirect)(nil),      // 9: greet.GreetRequestBiDirect
	(*GreetResponseBiDirecrt)(nil),    // 10: greet.GreetResponseBiDirecrt
	(*GreetRequestDeadline)(nil),      // 11: greet.GreetRequestDeadline
	(*GreetResponseDeadline)(nil),     // 12: greet.GreetResponseDeadline
}
var file_greet_greetpb_greet_proto_depIdxs = []int32{
	0,  // 0: greet.GreetRequest.greeting:type_name -> greet.Greeting
	3,  // 1: greet.GreetRequestManyTimes.greetingMT:type_name -> greet.GreetingManyTimes
	6,  // 2: greet.GreetRequestCliStreaming.greeting:type_name -> greet.GreetingCliStreaming
	0,  // 3: greet.GreetRequestBiDirect.greeting_bd:type_name -> greet.Greeting
	0,  // 4: greet.GreetRequestDeadline.greeting_deadline:type_name -> greet.Greeting
	1,  // 5: greet.GreetService.Greet:input_type -> greet.GreetRequest
	4,  // 6: greet.GreetService.GreetManyTimes:input_type -> greet.GreetRequestManyTimes
	7,  // 7: greet.GreetService.GreetLong:input_type -> greet.GreetRequestCliStreaming
	9,  // 8: greet.GreetService.GreetBiDirect:input_type -> greet.GreetRequestBiDirect
	1,  // 9: greet.GreetService.GreetErr:input_type -> greet.GreetRequest
	11, // 10: greet.GreetService.GreetWithDeadline:input_type -> greet.GreetRequestDeadline
	2,  // 11: greet.GreetService.Greet:output_type -> greet.GreetResponse
	5,  // 12: greet.GreetService.GreetManyTimes:output_type -> greet.GreetResponseManyTimes
	8,  // 13: greet.GreetService.GreetLong:output_type -> greet.GreetResponseCliStriaming
	10, // 14: greet.GreetService.GreetBiDirect:output_type -> greet.GreetResponseBiDirecrt
	2,  // 15: greet.GreetService.GreetErr:output_type -> greet.GreetResponse
	12, // 16: greet.GreetService.GreetWithDeadline:output_type -> greet.GreetResponseDeadline
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_greet_greetpb_greet_proto_init() }
func file_greet_greetpb_greet_proto_init() {
	if File_greet_greetpb_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greet_greetpb_greet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Greeting); i {
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
		file_greet_greetpb_greet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequest); i {
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
		file_greet_greetpb_greet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponse); i {
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
		file_greet_greetpb_greet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingManyTimes); i {
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
		file_greet_greetpb_greet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequestManyTimes); i {
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
		file_greet_greetpb_greet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponseManyTimes); i {
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
		file_greet_greetpb_greet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingCliStreaming); i {
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
		file_greet_greetpb_greet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequestCliStreaming); i {
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
		file_greet_greetpb_greet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponseCliStriaming); i {
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
		file_greet_greetpb_greet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequestBiDirect); i {
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
		file_greet_greetpb_greet_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponseBiDirecrt); i {
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
		file_greet_greetpb_greet_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequestDeadline); i {
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
		file_greet_greetpb_greet_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponseDeadline); i {
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
			RawDescriptor: file_greet_greetpb_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greet_greetpb_greet_proto_goTypes,
		DependencyIndexes: file_greet_greetpb_greet_proto_depIdxs,
		MessageInfos:      file_greet_greetpb_greet_proto_msgTypes,
	}.Build()
	File_greet_greetpb_greet_proto = out.File
	file_greet_greetpb_greet_proto_rawDesc = nil
	file_greet_greetpb_greet_proto_goTypes = nil
	file_greet_greetpb_greet_proto_depIdxs = nil
}
