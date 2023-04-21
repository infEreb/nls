// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: agent.proto

package gen

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

type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Status   string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{0}
}

func (x *Agent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Agent) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Agent) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId string `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
}

func (x *GetAgentRequest) Reset() {
	*x = GetAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentRequest) ProtoMessage() {}

func (x *GetAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentRequest.ProtoReflect.Descriptor instead.
func (*GetAgentRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{1}
}

func (x *GetAgentRequest) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

type GetAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agent *Agent `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
}

func (x *GetAgentResponse) Reset() {
	*x = GetAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentResponse) ProtoMessage() {}

func (x *GetAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentResponse.ProtoReflect.Descriptor instead.
func (*GetAgentResponse) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{2}
}

func (x *GetAgentResponse) GetAgent() *Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

type GetAgentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAgentsRequest) Reset() {
	*x = GetAgentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentsRequest) ProtoMessage() {}

func (x *GetAgentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentsRequest.ProtoReflect.Descriptor instead.
func (*GetAgentsRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{3}
}

type GetAgentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agent []*Agent `protobuf:"bytes,1,rep,name=agent,proto3" json:"agent,omitempty"`
}

func (x *GetAgentsResponse) Reset() {
	*x = GetAgentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentsResponse) ProtoMessage() {}

func (x *GetAgentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentsResponse.ProtoReflect.Descriptor instead.
func (*GetAgentsResponse) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{4}
}

func (x *GetAgentsResponse) GetAgent() []*Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

type PostAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agent *Agent `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
}

func (x *PostAgentRequest) Reset() {
	*x = PostAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostAgentRequest) ProtoMessage() {}

func (x *PostAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostAgentRequest.ProtoReflect.Descriptor instead.
func (*PostAgentRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{5}
}

func (x *PostAgentRequest) GetAgent() *Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

type PostAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agent *Agent `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
}

func (x *PostAgentResponse) Reset() {
	*x = PostAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostAgentResponse) ProtoMessage() {}

func (x *PostAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostAgentResponse.ProtoReflect.Descriptor instead.
func (*PostAgentResponse) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{6}
}

func (x *PostAgentResponse) GetAgent() *Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

type PutAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agent *Agent `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
}

func (x *PutAgentRequest) Reset() {
	*x = PutAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutAgentRequest) ProtoMessage() {}

func (x *PutAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutAgentRequest.ProtoReflect.Descriptor instead.
func (*PutAgentRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{7}
}

func (x *PutAgentRequest) GetAgent() *Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

type PutAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agent *Agent `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
}

func (x *PutAgentResponse) Reset() {
	*x = PutAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutAgentResponse) ProtoMessage() {}

func (x *PutAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutAgentResponse.ProtoReflect.Descriptor instead.
func (*PutAgentResponse) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{8}
}

func (x *PutAgentResponse) GetAgent() *Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

type DeleteAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId string `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
}

func (x *DeleteAgentRequest) Reset() {
	*x = DeleteAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAgentRequest) ProtoMessage() {}

func (x *DeleteAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAgentRequest.ProtoReflect.Descriptor instead.
func (*DeleteAgentRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteAgentRequest) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

type DeleteAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted int32 `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteAgentResponse) Reset() {
	*x = DeleteAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAgentResponse) ProtoMessage() {}

func (x *DeleteAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAgentResponse.ProtoReflect.Descriptor instead.
func (*DeleteAgentResponse) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteAgentResponse) GetDeleted() int32 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a,
	0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x22, 0x30, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x22, 0x31, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x2f, 0x0a, 0x0f, 0x50, 0x75, 0x74, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x10, 0x50, 0x75, 0x74, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0x92, 0x02, 0x0a, 0x0c,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x11,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x50, 0x75, 0x74, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x10, 0x2e, 0x50, 0x75, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x50, 0x75, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_proto_rawDescOnce sync.Once
	file_agent_proto_rawDescData = file_agent_proto_rawDesc
)

func file_agent_proto_rawDescGZIP() []byte {
	file_agent_proto_rawDescOnce.Do(func() {
		file_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_proto_rawDescData)
	})
	return file_agent_proto_rawDescData
}

var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_agent_proto_goTypes = []interface{}{
	(*Agent)(nil),               // 0: Agent
	(*GetAgentRequest)(nil),     // 1: GetAgentRequest
	(*GetAgentResponse)(nil),    // 2: GetAgentResponse
	(*GetAgentsRequest)(nil),    // 3: GetAgentsRequest
	(*GetAgentsResponse)(nil),   // 4: GetAgentsResponse
	(*PostAgentRequest)(nil),    // 5: PostAgentRequest
	(*PostAgentResponse)(nil),   // 6: PostAgentResponse
	(*PutAgentRequest)(nil),     // 7: PutAgentRequest
	(*PutAgentResponse)(nil),    // 8: PutAgentResponse
	(*DeleteAgentRequest)(nil),  // 9: DeleteAgentRequest
	(*DeleteAgentResponse)(nil), // 10: DeleteAgentResponse
}
var file_agent_proto_depIdxs = []int32{
	0,  // 0: GetAgentResponse.agent:type_name -> Agent
	0,  // 1: GetAgentsResponse.agent:type_name -> Agent
	0,  // 2: PostAgentRequest.agent:type_name -> Agent
	0,  // 3: PostAgentResponse.agent:type_name -> Agent
	0,  // 4: PutAgentRequest.agent:type_name -> Agent
	0,  // 5: PutAgentResponse.agent:type_name -> Agent
	1,  // 6: AgentService.GetAgent:input_type -> GetAgentRequest
	3,  // 7: AgentService.GetAgents:input_type -> GetAgentsRequest
	5,  // 8: AgentService.PostAgent:input_type -> PostAgentRequest
	7,  // 9: AgentService.PutAgent:input_type -> PutAgentRequest
	9,  // 10: AgentService.DeleteAgent:input_type -> DeleteAgentRequest
	2,  // 11: AgentService.GetAgent:output_type -> GetAgentResponse
	4,  // 12: AgentService.GetAgents:output_type -> GetAgentsResponse
	6,  // 13: AgentService.PostAgent:output_type -> PostAgentResponse
	8,  // 14: AgentService.PutAgent:output_type -> PutAgentResponse
	10, // 15: AgentService.DeleteAgent:output_type -> DeleteAgentResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_agent_proto_init() }
func file_agent_proto_init() {
	if File_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agent); i {
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
		file_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgentRequest); i {
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
		file_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgentResponse); i {
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
		file_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgentsRequest); i {
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
		file_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgentsResponse); i {
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
		file_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostAgentRequest); i {
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
		file_agent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostAgentResponse); i {
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
		file_agent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutAgentRequest); i {
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
		file_agent_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutAgentResponse); i {
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
		file_agent_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAgentRequest); i {
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
		file_agent_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAgentResponse); i {
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
			RawDescriptor: file_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_proto_goTypes,
		DependencyIndexes: file_agent_proto_depIdxs,
		MessageInfos:      file_agent_proto_msgTypes,
	}.Build()
	File_agent_proto = out.File
	file_agent_proto_rawDesc = nil
	file_agent_proto_goTypes = nil
	file_agent_proto_depIdxs = nil
}
