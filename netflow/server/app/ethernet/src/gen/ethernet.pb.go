// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: ethernet.proto

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

type Ethernet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AgentId  string   `protobuf:"bytes,2,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Name     string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DevName  string   `protobuf:"bytes,4,opt,name=dev_name,json=devName,proto3" json:"dev_name,omitempty"`
	Desc     string   `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	Ipv4     []string `protobuf:"bytes,6,rep,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6     []string `protobuf:"bytes,7,rep,name=ipv6,proto3" json:"ipv6,omitempty"`
	Hardware string   `protobuf:"bytes,8,opt,name=hardware,proto3" json:"hardware,omitempty"`
}

func (x *Ethernet) Reset() {
	*x = Ethernet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethernet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ethernet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ethernet) ProtoMessage() {}

func (x *Ethernet) ProtoReflect() protoreflect.Message {
	mi := &file_ethernet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ethernet.ProtoReflect.Descriptor instead.
func (*Ethernet) Descriptor() ([]byte, []int) {
	return file_ethernet_proto_rawDescGZIP(), []int{0}
}

func (x *Ethernet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ethernet) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *Ethernet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Ethernet) GetDevName() string {
	if x != nil {
		return x.DevName
	}
	return ""
}

func (x *Ethernet) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Ethernet) GetIpv4() []string {
	if x != nil {
		return x.Ipv4
	}
	return nil
}

func (x *Ethernet) GetIpv6() []string {
	if x != nil {
		return x.Ipv6
	}
	return nil
}

func (x *Ethernet) GetHardware() string {
	if x != nil {
		return x.Hardware
	}
	return ""
}

type GetEthernetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthernetId string `protobuf:"bytes,1,opt,name=Ethernet_id,json=EthernetId,proto3" json:"Ethernet_id,omitempty"`
}

func (x *GetEthernetRequest) Reset() {
	*x = GetEthernetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethernet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEthernetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEthernetRequest) ProtoMessage() {}

func (x *GetEthernetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ethernet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEthernetRequest.ProtoReflect.Descriptor instead.
func (*GetEthernetRequest) Descriptor() ([]byte, []int) {
	return file_ethernet_proto_rawDescGZIP(), []int{1}
}

func (x *GetEthernetRequest) GetEthernetId() string {
	if x != nil {
		return x.EthernetId
	}
	return ""
}

type GetEthernetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ethernet *Ethernet `protobuf:"bytes,1,opt,name=Ethernet,proto3" json:"Ethernet,omitempty"`
}

func (x *GetEthernetResponse) Reset() {
	*x = GetEthernetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethernet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEthernetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEthernetResponse) ProtoMessage() {}

func (x *GetEthernetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ethernet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEthernetResponse.ProtoReflect.Descriptor instead.
func (*GetEthernetResponse) Descriptor() ([]byte, []int) {
	return file_ethernet_proto_rawDescGZIP(), []int{2}
}

func (x *GetEthernetResponse) GetEthernet() *Ethernet {
	if x != nil {
		return x.Ethernet
	}
	return nil
}

type GetEthernetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEthernetsRequest) Reset() {
	*x = GetEthernetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethernet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEthernetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEthernetsRequest) ProtoMessage() {}

func (x *GetEthernetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ethernet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEthernetsRequest.ProtoReflect.Descriptor instead.
func (*GetEthernetsRequest) Descriptor() ([]byte, []int) {
	return file_ethernet_proto_rawDescGZIP(), []int{3}
}

type GetEthernetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ethernets []*Ethernet `protobuf:"bytes,1,rep,name=Ethernets,proto3" json:"Ethernets,omitempty"`
}

func (x *GetEthernetsResponse) Reset() {
	*x = GetEthernetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethernet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEthernetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEthernetsResponse) ProtoMessage() {}

func (x *GetEthernetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ethernet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEthernetsResponse.ProtoReflect.Descriptor instead.
func (*GetEthernetsResponse) Descriptor() ([]byte, []int) {
	return file_ethernet_proto_rawDescGZIP(), []int{4}
}

func (x *GetEthernetsResponse) GetEthernets() []*Ethernet {
	if x != nil {
		return x.Ethernets
	}
	return nil
}

type PostEthernetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ethernet *Ethernet `protobuf:"bytes,1,opt,name=Ethernet,proto3" json:"Ethernet,omitempty"`
}

func (x *PostEthernetRequest) Reset() {
	*x = PostEthernetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethernet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostEthernetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostEthernetRequest) ProtoMessage() {}

func (x *PostEthernetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ethernet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostEthernetRequest.ProtoReflect.Descriptor instead.
func (*PostEthernetRequest) Descriptor() ([]byte, []int) {
	return file_ethernet_proto_rawDescGZIP(), []int{5}
}

func (x *PostEthernetRequest) GetEthernet() *Ethernet {
	if x != nil {
		return x.Ethernet
	}
	return nil
}

type PostEthernetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ethernet *Ethernet `protobuf:"bytes,1,opt,name=Ethernet,proto3" json:"Ethernet,omitempty"`
}

func (x *PostEthernetResponse) Reset() {
	*x = PostEthernetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethernet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostEthernetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostEthernetResponse) ProtoMessage() {}

func (x *PostEthernetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ethernet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostEthernetResponse.ProtoReflect.Descriptor instead.
func (*PostEthernetResponse) Descriptor() ([]byte, []int) {
	return file_ethernet_proto_rawDescGZIP(), []int{6}
}

func (x *PostEthernetResponse) GetEthernet() *Ethernet {
	if x != nil {
		return x.Ethernet
	}
	return nil
}

type PutEthernetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthernetId string    `protobuf:"bytes,1,opt,name=Ethernet_id,json=EthernetId,proto3" json:"Ethernet_id,omitempty"`
	Ethernet   *Ethernet `protobuf:"bytes,2,opt,name=Ethernet,proto3" json:"Ethernet,omitempty"`
}

func (x *PutEthernetRequest) Reset() {
	*x = PutEthernetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethernet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutEthernetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutEthernetRequest) ProtoMessage() {}

func (x *PutEthernetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ethernet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutEthernetRequest.ProtoReflect.Descriptor instead.
func (*PutEthernetRequest) Descriptor() ([]byte, []int) {
	return file_ethernet_proto_rawDescGZIP(), []int{7}
}

func (x *PutEthernetRequest) GetEthernetId() string {
	if x != nil {
		return x.EthernetId
	}
	return ""
}

func (x *PutEthernetRequest) GetEthernet() *Ethernet {
	if x != nil {
		return x.Ethernet
	}
	return nil
}

type PutEthernetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ethernet *Ethernet `protobuf:"bytes,1,opt,name=Ethernet,proto3" json:"Ethernet,omitempty"`
}

func (x *PutEthernetResponse) Reset() {
	*x = PutEthernetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethernet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutEthernetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutEthernetResponse) ProtoMessage() {}

func (x *PutEthernetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ethernet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutEthernetResponse.ProtoReflect.Descriptor instead.
func (*PutEthernetResponse) Descriptor() ([]byte, []int) {
	return file_ethernet_proto_rawDescGZIP(), []int{8}
}

func (x *PutEthernetResponse) GetEthernet() *Ethernet {
	if x != nil {
		return x.Ethernet
	}
	return nil
}

type DeleteEthernetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthernetId string `protobuf:"bytes,1,opt,name=Ethernet_id,json=EthernetId,proto3" json:"Ethernet_id,omitempty"`
}

func (x *DeleteEthernetRequest) Reset() {
	*x = DeleteEthernetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethernet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEthernetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEthernetRequest) ProtoMessage() {}

func (x *DeleteEthernetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ethernet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEthernetRequest.ProtoReflect.Descriptor instead.
func (*DeleteEthernetRequest) Descriptor() ([]byte, []int) {
	return file_ethernet_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteEthernetRequest) GetEthernetId() string {
	if x != nil {
		return x.EthernetId
	}
	return ""
}

type DeleteEthernetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted uint32 `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteEthernetResponse) Reset() {
	*x = DeleteEthernetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethernet_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEthernetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEthernetResponse) ProtoMessage() {}

func (x *DeleteEthernetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ethernet_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEthernetResponse.ProtoReflect.Descriptor instead.
func (*DeleteEthernetResponse) Descriptor() ([]byte, []int) {
	return file_ethernet_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteEthernetResponse) GetDeleted() uint32 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

var File_ethernet_proto protoreflect.FileDescriptor

var file_ethernet_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbc, 0x01, 0x0a, 0x08, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x64, 0x65, 0x76, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x70, 0x76, 0x34, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x36, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x70, 0x76, 0x36, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x22,
	0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x74, 0x68,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x08, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x08, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x52, 0x09, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x73, 0x22, 0x3c, 0x0a, 0x13,
	0x50, 0x6f, 0x73, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x52, 0x08, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x22, 0x3d, 0x0a, 0x14, 0x50, 0x6f,
	0x73, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52,
	0x08, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x22, 0x5c, 0x0a, 0x12, 0x50, 0x75, 0x74,
	0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x08, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x08, 0x45,
	0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x22, 0x3c, 0x0a, 0x13, 0x50, 0x75, 0x74, 0x45, 0x74,
	0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x08, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x08, 0x45, 0x74, 0x68,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x22, 0x38, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x22,
	0x32, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x32, 0xc2, 0x02, 0x0a, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45, 0x74,
	0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x73, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x74, 0x68,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x12, 0x14,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x50,
	0x75, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x50, 0x75, 0x74,
	0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x50, 0x75, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ethernet_proto_rawDescOnce sync.Once
	file_ethernet_proto_rawDescData = file_ethernet_proto_rawDesc
)

func file_ethernet_proto_rawDescGZIP() []byte {
	file_ethernet_proto_rawDescOnce.Do(func() {
		file_ethernet_proto_rawDescData = protoimpl.X.CompressGZIP(file_ethernet_proto_rawDescData)
	})
	return file_ethernet_proto_rawDescData
}

var file_ethernet_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_ethernet_proto_goTypes = []interface{}{
	(*Ethernet)(nil),               // 0: Ethernet
	(*GetEthernetRequest)(nil),     // 1: GetEthernetRequest
	(*GetEthernetResponse)(nil),    // 2: GetEthernetResponse
	(*GetEthernetsRequest)(nil),    // 3: GetEthernetsRequest
	(*GetEthernetsResponse)(nil),   // 4: GetEthernetsResponse
	(*PostEthernetRequest)(nil),    // 5: PostEthernetRequest
	(*PostEthernetResponse)(nil),   // 6: PostEthernetResponse
	(*PutEthernetRequest)(nil),     // 7: PutEthernetRequest
	(*PutEthernetResponse)(nil),    // 8: PutEthernetResponse
	(*DeleteEthernetRequest)(nil),  // 9: DeleteEthernetRequest
	(*DeleteEthernetResponse)(nil), // 10: DeleteEthernetResponse
}
var file_ethernet_proto_depIdxs = []int32{
	0,  // 0: GetEthernetResponse.Ethernet:type_name -> Ethernet
	0,  // 1: GetEthernetsResponse.Ethernets:type_name -> Ethernet
	0,  // 2: PostEthernetRequest.Ethernet:type_name -> Ethernet
	0,  // 3: PostEthernetResponse.Ethernet:type_name -> Ethernet
	0,  // 4: PutEthernetRequest.Ethernet:type_name -> Ethernet
	0,  // 5: PutEthernetResponse.Ethernet:type_name -> Ethernet
	1,  // 6: EthernetService.GetEthernet:input_type -> GetEthernetRequest
	3,  // 7: EthernetService.GetEthernets:input_type -> GetEthernetsRequest
	5,  // 8: EthernetService.PostEthernet:input_type -> PostEthernetRequest
	7,  // 9: EthernetService.PutEthernet:input_type -> PutEthernetRequest
	9,  // 10: EthernetService.DeleteEthernet:input_type -> DeleteEthernetRequest
	2,  // 11: EthernetService.GetEthernet:output_type -> GetEthernetResponse
	4,  // 12: EthernetService.GetEthernets:output_type -> GetEthernetsResponse
	6,  // 13: EthernetService.PostEthernet:output_type -> PostEthernetResponse
	8,  // 14: EthernetService.PutEthernet:output_type -> PutEthernetResponse
	10, // 15: EthernetService.DeleteEthernet:output_type -> DeleteEthernetResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_ethernet_proto_init() }
func file_ethernet_proto_init() {
	if File_ethernet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ethernet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ethernet); i {
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
		file_ethernet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEthernetRequest); i {
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
		file_ethernet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEthernetResponse); i {
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
		file_ethernet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEthernetsRequest); i {
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
		file_ethernet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEthernetsResponse); i {
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
		file_ethernet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostEthernetRequest); i {
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
		file_ethernet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostEthernetResponse); i {
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
		file_ethernet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutEthernetRequest); i {
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
		file_ethernet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutEthernetResponse); i {
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
		file_ethernet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEthernetRequest); i {
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
		file_ethernet_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEthernetResponse); i {
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
			RawDescriptor: file_ethernet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ethernet_proto_goTypes,
		DependencyIndexes: file_ethernet_proto_depIdxs,
		MessageInfos:      file_ethernet_proto_msgTypes,
	}.Build()
	File_ethernet_proto = out.File
	file_ethernet_proto_rawDesc = nil
	file_ethernet_proto_goTypes = nil
	file_ethernet_proto_depIdxs = nil
}