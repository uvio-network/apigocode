// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pbf/post/update.proto

package post

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

// UpdateI is the input for updating posts.
//
//	{
//	    "object": [
//	        {
//	            "intern": {
//	                "id": "778237"
//	            },
//	            "public": {
//	                "hash": "0x1234"
//	            }
//	        }
//	    ]
//	}
type UpdateI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *UpdateI_Filter   `protobuf:"bytes,100,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*UpdateI_Object `protobuf:"bytes,200,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *UpdateI) Reset() {
	*x = UpdateI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_post_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateI) ProtoMessage() {}

func (x *UpdateI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_post_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateI.ProtoReflect.Descriptor instead.
func (*UpdateI) Descriptor() ([]byte, []int) {
	return file_pbf_post_update_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateI) GetFilter() *UpdateI_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *UpdateI) GetObject() []*UpdateI_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type UpdateI_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateI_Filter) Reset() {
	*x = UpdateI_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_post_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateI_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateI_Filter) ProtoMessage() {}

func (x *UpdateI_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_post_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateI_Filter.ProtoReflect.Descriptor instead.
func (*UpdateI_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_post_update_proto_rawDescGZIP(), []int{1}
}

type UpdateI_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern *UpdateI_Object_Intern `protobuf:"bytes,100,opt,name=intern,proto3" json:"intern,omitempty"`
	Public *UpdateI_Object_Public `protobuf:"bytes,200,opt,name=public,proto3" json:"public,omitempty"`
	Symbol *UpdateI_Object_Symbol `protobuf:"bytes,300,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *UpdateI_Object) Reset() {
	*x = UpdateI_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_post_update_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateI_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateI_Object) ProtoMessage() {}

func (x *UpdateI_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_post_update_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateI_Object.ProtoReflect.Descriptor instead.
func (*UpdateI_Object) Descriptor() ([]byte, []int) {
	return file_pbf_post_update_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateI_Object) GetIntern() *UpdateI_Object_Intern {
	if x != nil {
		return x.Intern
	}
	return nil
}

func (x *UpdateI_Object) GetPublic() *UpdateI_Object_Public {
	if x != nil {
		return x.Public
	}
	return nil
}

func (x *UpdateI_Object) GetSymbol() *UpdateI_Object_Symbol {
	if x != nil {
		return x.Symbol
	}
	return nil
}

type UpdateI_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the ID of the post object being updated.
	Id string `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateI_Object_Intern) Reset() {
	*x = UpdateI_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_post_update_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateI_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateI_Object_Intern) ProtoMessage() {}

func (x *UpdateI_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_post_update_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateI_Object_Intern.ProtoReflect.Descriptor instead.
func (*UpdateI_Object_Intern) Descriptor() ([]byte, []int) {
	return file_pbf_post_update_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateI_Object_Intern) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateI_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hash is a claim's confirmed onchain transaction hash. hash must not be
	// updated for comments. hash can only be updated once, if its prior value was
	// empty. Updating hash implies that the given claim got confirmed onchain,
	// and with it the claim's interim lifecycle phase will not be "pending"
	// anymore, but instead switch to the provided desired lifecycle phase.
	Hash string `protobuf:"bytes,100,opt,name=hash,proto3" json:"hash,omitempty"`
	// meta may contain onchain arbitrary meta data.
	Meta string `protobuf:"bytes,200,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *UpdateI_Object_Public) Reset() {
	*x = UpdateI_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_post_update_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateI_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateI_Object_Public) ProtoMessage() {}

func (x *UpdateI_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_post_update_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateI_Object_Public.ProtoReflect.Descriptor instead.
func (*UpdateI_Object_Public) Descriptor() ([]byte, []int) {
	return file_pbf_post_update_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateI_Object_Public) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *UpdateI_Object_Public) GetMeta() string {
	if x != nil {
		return x.Meta
	}
	return ""
}

type UpdateI_Object_Symbol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateI_Object_Symbol) Reset() {
	*x = UpdateI_Object_Symbol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_post_update_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateI_Object_Symbol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateI_Object_Symbol) ProtoMessage() {}

func (x *UpdateI_Object_Symbol) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_post_update_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateI_Object_Symbol.ProtoReflect.Descriptor instead.
func (*UpdateI_Object_Symbol) Descriptor() ([]byte, []int) {
	return file_pbf_post_update_proto_rawDescGZIP(), []int{5}
}

// UpdateO is the output for updating posts.
//
//	{
//	    "object": [
//	        {
//	            "intern": {
//	                "status": "updated"
//	            }
//	        }
//	    ]
//	}
type UpdateO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *UpdateO_Filter   `protobuf:"bytes,100,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*UpdateO_Object `protobuf:"bytes,200,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *UpdateO) Reset() {
	*x = UpdateO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_post_update_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateO) ProtoMessage() {}

func (x *UpdateO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_post_update_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateO.ProtoReflect.Descriptor instead.
func (*UpdateO) Descriptor() ([]byte, []int) {
	return file_pbf_post_update_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateO) GetFilter() *UpdateO_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *UpdateO) GetObject() []*UpdateO_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type UpdateO_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateO_Filter) Reset() {
	*x = UpdateO_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_post_update_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateO_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateO_Filter) ProtoMessage() {}

func (x *UpdateO_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_post_update_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateO_Filter.ProtoReflect.Descriptor instead.
func (*UpdateO_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_post_update_proto_rawDescGZIP(), []int{7}
}

type UpdateO_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern *UpdateO_Object_Intern `protobuf:"bytes,100,opt,name=intern,proto3" json:"intern,omitempty"`
	Public *UpdateO_Object_Public `protobuf:"bytes,200,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *UpdateO_Object) Reset() {
	*x = UpdateO_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_post_update_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateO_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateO_Object) ProtoMessage() {}

func (x *UpdateO_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_post_update_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateO_Object.ProtoReflect.Descriptor instead.
func (*UpdateO_Object) Descriptor() ([]byte, []int) {
	return file_pbf_post_update_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateO_Object) GetIntern() *UpdateO_Object_Intern {
	if x != nil {
		return x.Intern
	}
	return nil
}

func (x *UpdateO_Object) GetPublic() *UpdateO_Object_Public {
	if x != nil {
		return x.Public
	}
	return nil
}

type UpdateO_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// status is the resource status upon successful post modification.
	Status string `protobuf:"bytes,100,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateO_Object_Intern) Reset() {
	*x = UpdateO_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_post_update_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateO_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateO_Object_Intern) ProtoMessage() {}

func (x *UpdateO_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_post_update_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateO_Object_Intern.ProtoReflect.Descriptor instead.
func (*UpdateO_Object_Intern) Descriptor() ([]byte, []int) {
	return file_pbf_post_update_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateO_Object_Intern) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateO_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateO_Object_Public) Reset() {
	*x = UpdateO_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_post_update_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateO_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateO_Object_Public) ProtoMessage() {}

func (x *UpdateO_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_post_update_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateO_Object_Public.ProtoReflect.Descriptor instead.
func (*UpdateO_Object_Public) Descriptor() ([]byte, []int) {
	return file_pbf_post_update_proto_rawDescGZIP(), []int{10}
}

var File_pbf_post_update_proto protoreflect.FileDescriptor

var file_pbf_post_update_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x66, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x66, 0x0a,
	0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xb1, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12,
	0x34, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x5f,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0xac, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x53, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x27, 0x0a, 0x15, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x5f,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x12, 0x13, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22,
	0x66, 0x0a, 0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x7b, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x12, 0x34, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x2f, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pbf_post_update_proto_rawDescOnce sync.Once
	file_pbf_post_update_proto_rawDescData = file_pbf_post_update_proto_rawDesc
)

func file_pbf_post_update_proto_rawDescGZIP() []byte {
	file_pbf_post_update_proto_rawDescOnce.Do(func() {
		file_pbf_post_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_post_update_proto_rawDescData)
	})
	return file_pbf_post_update_proto_rawDescData
}

var file_pbf_post_update_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pbf_post_update_proto_goTypes = []interface{}{
	(*UpdateI)(nil),               // 0: post.UpdateI
	(*UpdateI_Filter)(nil),        // 1: post.UpdateI_Filter
	(*UpdateI_Object)(nil),        // 2: post.UpdateI_Object
	(*UpdateI_Object_Intern)(nil), // 3: post.UpdateI_Object_Intern
	(*UpdateI_Object_Public)(nil), // 4: post.UpdateI_Object_Public
	(*UpdateI_Object_Symbol)(nil), // 5: post.UpdateI_Object_Symbol
	(*UpdateO)(nil),               // 6: post.UpdateO
	(*UpdateO_Filter)(nil),        // 7: post.UpdateO_Filter
	(*UpdateO_Object)(nil),        // 8: post.UpdateO_Object
	(*UpdateO_Object_Intern)(nil), // 9: post.UpdateO_Object_Intern
	(*UpdateO_Object_Public)(nil), // 10: post.UpdateO_Object_Public
}
var file_pbf_post_update_proto_depIdxs = []int32{
	1,  // 0: post.UpdateI.filter:type_name -> post.UpdateI_Filter
	2,  // 1: post.UpdateI.object:type_name -> post.UpdateI_Object
	3,  // 2: post.UpdateI_Object.intern:type_name -> post.UpdateI_Object_Intern
	4,  // 3: post.UpdateI_Object.public:type_name -> post.UpdateI_Object_Public
	5,  // 4: post.UpdateI_Object.symbol:type_name -> post.UpdateI_Object_Symbol
	7,  // 5: post.UpdateO.filter:type_name -> post.UpdateO_Filter
	8,  // 6: post.UpdateO.object:type_name -> post.UpdateO_Object
	9,  // 7: post.UpdateO_Object.intern:type_name -> post.UpdateO_Object_Intern
	10, // 8: post.UpdateO_Object.public:type_name -> post.UpdateO_Object_Public
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_pbf_post_update_proto_init() }
func file_pbf_post_update_proto_init() {
	if File_pbf_post_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_post_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateI); i {
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
		file_pbf_post_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateI_Filter); i {
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
		file_pbf_post_update_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateI_Object); i {
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
		file_pbf_post_update_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateI_Object_Intern); i {
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
		file_pbf_post_update_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateI_Object_Public); i {
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
		file_pbf_post_update_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateI_Object_Symbol); i {
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
		file_pbf_post_update_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateO); i {
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
		file_pbf_post_update_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateO_Filter); i {
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
		file_pbf_post_update_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateO_Object); i {
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
		file_pbf_post_update_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateO_Object_Intern); i {
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
		file_pbf_post_update_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateO_Object_Public); i {
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
			RawDescriptor: file_pbf_post_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_post_update_proto_goTypes,
		DependencyIndexes: file_pbf_post_update_proto_depIdxs,
		MessageInfos:      file_pbf_post_update_proto_msgTypes,
	}.Build()
	File_pbf_post_update_proto = out.File
	file_pbf_post_update_proto_rawDesc = nil
	file_pbf_post_update_proto_goTypes = nil
	file_pbf_post_update_proto_depIdxs = nil
}
