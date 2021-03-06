// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.1
// source: trace-sampling-manager.proto

package trace_sampling_manager

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Port     int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trace_sampling_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_trace_sampling_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_trace_sampling_manager_proto_rawDescGZIP(), []int{0}
}

func (x *Host) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Host) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type HostsToTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts       []*Host `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	ComponentID string  `protobuf:"bytes,2,opt,name=componentID,proto3" json:"componentID,omitempty"`
}

func (x *HostsToTraceRequest) Reset() {
	*x = HostsToTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trace_sampling_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostsToTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostsToTraceRequest) ProtoMessage() {}

func (x *HostsToTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trace_sampling_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostsToTraceRequest.ProtoReflect.Descriptor instead.
func (*HostsToTraceRequest) Descriptor() ([]byte, []int) {
	return file_trace_sampling_manager_proto_rawDescGZIP(), []int{1}
}

func (x *HostsToTraceRequest) GetHosts() []*Host {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *HostsToTraceRequest) GetComponentID() string {
	if x != nil {
		return x.ComponentID
	}
	return ""
}

type RemoveHostsToTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts       []*Host `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	ComponentID string  `protobuf:"bytes,2,opt,name=componentID,proto3" json:"componentID,omitempty"`
}

func (x *RemoveHostsToTraceRequest) Reset() {
	*x = RemoveHostsToTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trace_sampling_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveHostsToTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveHostsToTraceRequest) ProtoMessage() {}

func (x *RemoveHostsToTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trace_sampling_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveHostsToTraceRequest.ProtoReflect.Descriptor instead.
func (*RemoveHostsToTraceRequest) Descriptor() ([]byte, []int) {
	return file_trace_sampling_manager_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveHostsToTraceRequest) GetHosts() []*Host {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *RemoveHostsToTraceRequest) GetComponentID() string {
	if x != nil {
		return x.ComponentID
	}
	return ""
}

type GetHostsToTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComponentID string `protobuf:"bytes,1,opt,name=componentID,proto3" json:"componentID,omitempty"`
}

func (x *GetHostsToTraceRequest) Reset() {
	*x = GetHostsToTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trace_sampling_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostsToTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostsToTraceRequest) ProtoMessage() {}

func (x *GetHostsToTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trace_sampling_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostsToTraceRequest.ProtoReflect.Descriptor instead.
func (*GetHostsToTraceRequest) Descriptor() ([]byte, []int) {
	return file_trace_sampling_manager_proto_rawDescGZIP(), []int{3}
}

func (x *GetHostsToTraceRequest) GetComponentID() string {
	if x != nil {
		return x.ComponentID
	}
	return ""
}

type GetHostsToTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts []*Host `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
}

func (x *GetHostsToTraceResponse) Reset() {
	*x = GetHostsToTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trace_sampling_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostsToTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostsToTraceResponse) ProtoMessage() {}

func (x *GetHostsToTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trace_sampling_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostsToTraceResponse.ProtoReflect.Descriptor instead.
func (*GetHostsToTraceResponse) Descriptor() ([]byte, []int) {
	return file_trace_sampling_manager_proto_rawDescGZIP(), []int{4}
}

func (x *GetHostsToTraceResponse) GetHosts() []*Host {
	if x != nil {
		return x.Hosts
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trace_sampling_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_trace_sampling_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_trace_sampling_manager_proto_rawDescGZIP(), []int{5}
}

var File_trace_sampling_manager_proto protoreflect.FileDescriptor

var file_trace_sampling_manager_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67,
	0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x6b,
	0x0a, 0x13, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x48, 0x6f,
	0x73, 0x74, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x71, 0x0a, 0x19, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x3a,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x4d, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x48, 0x6f,
	0x73, 0x74, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0xd7, 0x02, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x5f, 0x0a, 0x0f, 0x53,
	0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x2b,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x54, 0x6f, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x12,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x54, 0x6f, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x12, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x74, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73,
	0x74, 0x73, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x2e, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x54, 0x6f, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x54, 0x6f, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trace_sampling_manager_proto_rawDescOnce sync.Once
	file_trace_sampling_manager_proto_rawDescData = file_trace_sampling_manager_proto_rawDesc
)

func file_trace_sampling_manager_proto_rawDescGZIP() []byte {
	file_trace_sampling_manager_proto_rawDescOnce.Do(func() {
		file_trace_sampling_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_trace_sampling_manager_proto_rawDescData)
	})
	return file_trace_sampling_manager_proto_rawDescData
}

var file_trace_sampling_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_trace_sampling_manager_proto_goTypes = []interface{}{
	(*Host)(nil),                      // 0: trace_sampling_manager.Host
	(*HostsToTraceRequest)(nil),       // 1: trace_sampling_manager.HostsToTraceRequest
	(*RemoveHostsToTraceRequest)(nil), // 2: trace_sampling_manager.RemoveHostsToTraceRequest
	(*GetHostsToTraceRequest)(nil),    // 3: trace_sampling_manager.GetHostsToTraceRequest
	(*GetHostsToTraceResponse)(nil),   // 4: trace_sampling_manager.GetHostsToTraceResponse
	(*Empty)(nil),                     // 5: trace_sampling_manager.Empty
}
var file_trace_sampling_manager_proto_depIdxs = []int32{
	0, // 0: trace_sampling_manager.HostsToTraceRequest.hosts:type_name -> trace_sampling_manager.Host
	0, // 1: trace_sampling_manager.RemoveHostsToTraceRequest.hosts:type_name -> trace_sampling_manager.Host
	0, // 2: trace_sampling_manager.GetHostsToTraceResponse.hosts:type_name -> trace_sampling_manager.Host
	1, // 3: trace_sampling_manager.TraceSamplingManager.SetHostsToTrace:input_type -> trace_sampling_manager.HostsToTraceRequest
	2, // 4: trace_sampling_manager.TraceSamplingManager.RemoveHostsToTrace:input_type -> trace_sampling_manager.RemoveHostsToTraceRequest
	3, // 5: trace_sampling_manager.TraceSamplingManager.GetHostsToTrace:input_type -> trace_sampling_manager.GetHostsToTraceRequest
	5, // 6: trace_sampling_manager.TraceSamplingManager.SetHostsToTrace:output_type -> trace_sampling_manager.Empty
	5, // 7: trace_sampling_manager.TraceSamplingManager.RemoveHostsToTrace:output_type -> trace_sampling_manager.Empty
	4, // 8: trace_sampling_manager.TraceSamplingManager.GetHostsToTrace:output_type -> trace_sampling_manager.GetHostsToTraceResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_trace_sampling_manager_proto_init() }
func file_trace_sampling_manager_proto_init() {
	if File_trace_sampling_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trace_sampling_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
		file_trace_sampling_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostsToTraceRequest); i {
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
		file_trace_sampling_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveHostsToTraceRequest); i {
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
		file_trace_sampling_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostsToTraceRequest); i {
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
		file_trace_sampling_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostsToTraceResponse); i {
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
		file_trace_sampling_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_trace_sampling_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trace_sampling_manager_proto_goTypes,
		DependencyIndexes: file_trace_sampling_manager_proto_depIdxs,
		MessageInfos:      file_trace_sampling_manager_proto_msgTypes,
	}.Build()
	File_trace_sampling_manager_proto = out.File
	file_trace_sampling_manager_proto_rawDesc = nil
	file_trace_sampling_manager_proto_goTypes = nil
	file_trace_sampling_manager_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TraceSamplingManagerClient is the client API for TraceSamplingManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceSamplingManagerClient interface {
	SetHostsToTrace(ctx context.Context, in *HostsToTraceRequest, opts ...grpc.CallOption) (*Empty, error)
	RemoveHostsToTrace(ctx context.Context, in *RemoveHostsToTraceRequest, opts ...grpc.CallOption) (*Empty, error)
	GetHostsToTrace(ctx context.Context, in *GetHostsToTraceRequest, opts ...grpc.CallOption) (*GetHostsToTraceResponse, error)
}

type traceSamplingManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTraceSamplingManagerClient(cc grpc.ClientConnInterface) TraceSamplingManagerClient {
	return &traceSamplingManagerClient{cc}
}

func (c *traceSamplingManagerClient) SetHostsToTrace(ctx context.Context, in *HostsToTraceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/trace_sampling_manager.TraceSamplingManager/SetHostsToTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceSamplingManagerClient) RemoveHostsToTrace(ctx context.Context, in *RemoveHostsToTraceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/trace_sampling_manager.TraceSamplingManager/RemoveHostsToTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceSamplingManagerClient) GetHostsToTrace(ctx context.Context, in *GetHostsToTraceRequest, opts ...grpc.CallOption) (*GetHostsToTraceResponse, error) {
	out := new(GetHostsToTraceResponse)
	err := c.cc.Invoke(ctx, "/trace_sampling_manager.TraceSamplingManager/GetHostsToTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceSamplingManagerServer is the server API for TraceSamplingManager service.
type TraceSamplingManagerServer interface {
	SetHostsToTrace(context.Context, *HostsToTraceRequest) (*Empty, error)
	RemoveHostsToTrace(context.Context, *RemoveHostsToTraceRequest) (*Empty, error)
	GetHostsToTrace(context.Context, *GetHostsToTraceRequest) (*GetHostsToTraceResponse, error)
}

// UnimplementedTraceSamplingManagerServer can be embedded to have forward compatible implementations.
type UnimplementedTraceSamplingManagerServer struct {
}

func (*UnimplementedTraceSamplingManagerServer) SetHostsToTrace(context.Context, *HostsToTraceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHostsToTrace not implemented")
}
func (*UnimplementedTraceSamplingManagerServer) RemoveHostsToTrace(context.Context, *RemoveHostsToTraceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveHostsToTrace not implemented")
}
func (*UnimplementedTraceSamplingManagerServer) GetHostsToTrace(context.Context, *GetHostsToTraceRequest) (*GetHostsToTraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostsToTrace not implemented")
}

func RegisterTraceSamplingManagerServer(s *grpc.Server, srv TraceSamplingManagerServer) {
	s.RegisterService(&_TraceSamplingManager_serviceDesc, srv)
}

func _TraceSamplingManager_SetHostsToTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostsToTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceSamplingManagerServer).SetHostsToTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trace_sampling_manager.TraceSamplingManager/SetHostsToTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceSamplingManagerServer).SetHostsToTrace(ctx, req.(*HostsToTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceSamplingManager_RemoveHostsToTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveHostsToTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceSamplingManagerServer).RemoveHostsToTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trace_sampling_manager.TraceSamplingManager/RemoveHostsToTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceSamplingManagerServer).RemoveHostsToTrace(ctx, req.(*RemoveHostsToTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceSamplingManager_GetHostsToTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHostsToTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceSamplingManagerServer).GetHostsToTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trace_sampling_manager.TraceSamplingManager/GetHostsToTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceSamplingManagerServer).GetHostsToTrace(ctx, req.(*GetHostsToTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TraceSamplingManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trace_sampling_manager.TraceSamplingManager",
	HandlerType: (*TraceSamplingManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetHostsToTrace",
			Handler:    _TraceSamplingManager_SetHostsToTrace_Handler,
		},
		{
			MethodName: "RemoveHostsToTrace",
			Handler:    _TraceSamplingManager_RemoveHostsToTrace_Handler,
		},
		{
			MethodName: "GetHostsToTrace",
			Handler:    _TraceSamplingManager_GetHostsToTrace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trace-sampling-manager.proto",
}
