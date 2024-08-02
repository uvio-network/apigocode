// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pbf/vote/delete.proto

package vote

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

// DeleteI is the input for deleting votes.
//
//	{
//	    "object": [
//	        {
//	            "intern": {
//	                "id": "778237"
//	            }
//	        }
//	    ]
//	}
type DeleteI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *DeleteI_Filter   `protobuf:"bytes,100,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*DeleteI_Object `protobuf:"bytes,200,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *DeleteI) Reset() {
	*x = DeleteI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_delete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI) ProtoMessage() {}

func (x *DeleteI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_delete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI.ProtoReflect.Descriptor instead.
func (*DeleteI) Descriptor() ([]byte, []int) {
	return file_pbf_vote_delete_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteI) GetFilter() *DeleteI_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *DeleteI) GetObject() []*DeleteI_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type DeleteI_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteI_Filter) Reset() {
	*x = DeleteI_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_delete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI_Filter) ProtoMessage() {}

func (x *DeleteI_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_delete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI_Filter.ProtoReflect.Descriptor instead.
func (*DeleteI_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_vote_delete_proto_rawDescGZIP(), []int{1}
}

type DeleteI_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern *DeleteI_Object_Intern `protobuf:"bytes,100,opt,name=intern,proto3" json:"intern,omitempty"`
	Public *DeleteI_Object_Public `protobuf:"bytes,200,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *DeleteI_Object) Reset() {
	*x = DeleteI_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_delete_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI_Object) ProtoMessage() {}

func (x *DeleteI_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_delete_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI_Object.ProtoReflect.Descriptor instead.
func (*DeleteI_Object) Descriptor() ([]byte, []int) {
	return file_pbf_vote_delete_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteI_Object) GetIntern() *DeleteI_Object_Intern {
	if x != nil {
		return x.Intern
	}
	return nil
}

func (x *DeleteI_Object) GetPublic() *DeleteI_Object_Public {
	if x != nil {
		return x.Public
	}
	return nil
}

type DeleteI_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the ID of the vote object being deleted.
	Id string `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteI_Object_Intern) Reset() {
	*x = DeleteI_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_delete_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI_Object_Intern) ProtoMessage() {}

func (x *DeleteI_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_delete_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI_Object_Intern.ProtoReflect.Descriptor instead.
func (*DeleteI_Object_Intern) Descriptor() ([]byte, []int) {
	return file_pbf_vote_delete_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteI_Object_Intern) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteI_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteI_Object_Public) Reset() {
	*x = DeleteI_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_delete_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI_Object_Public) ProtoMessage() {}

func (x *DeleteI_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_delete_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI_Object_Public.ProtoReflect.Descriptor instead.
func (*DeleteI_Object_Public) Descriptor() ([]byte, []int) {
	return file_pbf_vote_delete_proto_rawDescGZIP(), []int{4}
}

// DeleteO is the output for deleting votes.
//
//	{
//	    "object": [
//	        {
//	            "intern": {
//	                "status": "deleted"
//	            }
//	        }
//	    ]
//	}
type DeleteO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *DeleteO_Filter   `protobuf:"bytes,100,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*DeleteO_Object `protobuf:"bytes,200,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *DeleteO) Reset() {
	*x = DeleteO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_delete_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO) ProtoMessage() {}

func (x *DeleteO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_delete_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO.ProtoReflect.Descriptor instead.
func (*DeleteO) Descriptor() ([]byte, []int) {
	return file_pbf_vote_delete_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteO) GetFilter() *DeleteO_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *DeleteO) GetObject() []*DeleteO_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type DeleteO_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteO_Filter) Reset() {
	*x = DeleteO_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_delete_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO_Filter) ProtoMessage() {}

func (x *DeleteO_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_delete_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO_Filter.ProtoReflect.Descriptor instead.
func (*DeleteO_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_vote_delete_proto_rawDescGZIP(), []int{6}
}

type DeleteO_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern *DeleteO_Object_Intern `protobuf:"bytes,100,opt,name=intern,proto3" json:"intern,omitempty"`
	Public *DeleteO_Object_Public `protobuf:"bytes,200,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *DeleteO_Object) Reset() {
	*x = DeleteO_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_delete_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO_Object) ProtoMessage() {}

func (x *DeleteO_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_delete_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO_Object.ProtoReflect.Descriptor instead.
func (*DeleteO_Object) Descriptor() ([]byte, []int) {
	return file_pbf_vote_delete_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteO_Object) GetIntern() *DeleteO_Object_Intern {
	if x != nil {
		return x.Intern
	}
	return nil
}

func (x *DeleteO_Object) GetPublic() *DeleteO_Object_Public {
	if x != nil {
		return x.Public
	}
	return nil
}

type DeleteO_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// status is the resource status upon successful vote deletion.
	Status string `protobuf:"bytes,100,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteO_Object_Intern) Reset() {
	*x = DeleteO_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_delete_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO_Object_Intern) ProtoMessage() {}

func (x *DeleteO_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_delete_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO_Object_Intern.ProtoReflect.Descriptor instead.
func (*DeleteO_Object_Intern) Descriptor() ([]byte, []int) {
	return file_pbf_vote_delete_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteO_Object_Intern) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DeleteO_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteO_Object_Public) Reset() {
	*x = DeleteO_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_delete_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO_Object_Public) ProtoMessage() {}

func (x *DeleteO_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_delete_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO_Object_Public.ProtoReflect.Descriptor instead.
func (*DeleteO_Object_Public) Descriptor() ([]byte, []int) {
	return file_pbf_vote_delete_proto_rawDescGZIP(), []int{9}
}

var File_pbf_vote_delete_proto protoreflect.FileDescriptor

var file_pbf_vote_delete_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x66, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x22, 0x66, 0x0a,
	0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x7b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x34,
	0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x66, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x2d, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x10,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0x7b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52,
	0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x2f, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x17,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x76, 0x6f,
	0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbf_vote_delete_proto_rawDescOnce sync.Once
	file_pbf_vote_delete_proto_rawDescData = file_pbf_vote_delete_proto_rawDesc
)

func file_pbf_vote_delete_proto_rawDescGZIP() []byte {
	file_pbf_vote_delete_proto_rawDescOnce.Do(func() {
		file_pbf_vote_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_vote_delete_proto_rawDescData)
	})
	return file_pbf_vote_delete_proto_rawDescData
}

var file_pbf_vote_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pbf_vote_delete_proto_goTypes = []interface{}{
	(*DeleteI)(nil),               // 0: vote.DeleteI
	(*DeleteI_Filter)(nil),        // 1: vote.DeleteI_Filter
	(*DeleteI_Object)(nil),        // 2: vote.DeleteI_Object
	(*DeleteI_Object_Intern)(nil), // 3: vote.DeleteI_Object_Intern
	(*DeleteI_Object_Public)(nil), // 4: vote.DeleteI_Object_Public
	(*DeleteO)(nil),               // 5: vote.DeleteO
	(*DeleteO_Filter)(nil),        // 6: vote.DeleteO_Filter
	(*DeleteO_Object)(nil),        // 7: vote.DeleteO_Object
	(*DeleteO_Object_Intern)(nil), // 8: vote.DeleteO_Object_Intern
	(*DeleteO_Object_Public)(nil), // 9: vote.DeleteO_Object_Public
}
var file_pbf_vote_delete_proto_depIdxs = []int32{
	1, // 0: vote.DeleteI.filter:type_name -> vote.DeleteI_Filter
	2, // 1: vote.DeleteI.object:type_name -> vote.DeleteI_Object
	3, // 2: vote.DeleteI_Object.intern:type_name -> vote.DeleteI_Object_Intern
	4, // 3: vote.DeleteI_Object.public:type_name -> vote.DeleteI_Object_Public
	6, // 4: vote.DeleteO.filter:type_name -> vote.DeleteO_Filter
	7, // 5: vote.DeleteO.object:type_name -> vote.DeleteO_Object
	8, // 6: vote.DeleteO_Object.intern:type_name -> vote.DeleteO_Object_Intern
	9, // 7: vote.DeleteO_Object.public:type_name -> vote.DeleteO_Object_Public
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pbf_vote_delete_proto_init() }
func file_pbf_vote_delete_proto_init() {
	if File_pbf_vote_delete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_vote_delete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI); i {
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
		file_pbf_vote_delete_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI_Filter); i {
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
		file_pbf_vote_delete_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI_Object); i {
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
		file_pbf_vote_delete_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI_Object_Intern); i {
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
		file_pbf_vote_delete_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI_Object_Public); i {
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
		file_pbf_vote_delete_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO); i {
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
		file_pbf_vote_delete_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO_Filter); i {
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
		file_pbf_vote_delete_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO_Object); i {
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
		file_pbf_vote_delete_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO_Object_Intern); i {
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
		file_pbf_vote_delete_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO_Object_Public); i {
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
			RawDescriptor: file_pbf_vote_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_vote_delete_proto_goTypes,
		DependencyIndexes: file_pbf_vote_delete_proto_depIdxs,
		MessageInfos:      file_pbf_vote_delete_proto_msgTypes,
	}.Build()
	File_pbf_vote_delete_proto = out.File
	file_pbf_vote_delete_proto_rawDesc = nil
	file_pbf_vote_delete_proto_goTypes = nil
	file_pbf_vote_delete_proto_depIdxs = nil
}
