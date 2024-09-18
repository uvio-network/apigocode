// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pbf/vote/search.proto

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

// SearchI is the input for searching votes.
//
//	{
//	    "filter": {
//	        "paging": {
//	            "kind": "page",
//	            "start": "0",
//	            "stop": "49"
//	        }
//	    },
//	    "object": [
//	        {
//	            "intern": {
//	                "owner": "551265"
//	            }
//	        }
//	    ]
//	}
type SearchI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *SearchI_Filter   `protobuf:"bytes,100,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*SearchI_Object `protobuf:"bytes,200,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *SearchI) Reset() {
	*x = SearchI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI) ProtoMessage() {}

func (x *SearchI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI.ProtoReflect.Descriptor instead.
func (*SearchI) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{0}
}

func (x *SearchI) GetFilter() *SearchI_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *SearchI) GetObject() []*SearchI_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type SearchI_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paging *SearchI_Filter_Paging `protobuf:"bytes,100,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *SearchI_Filter) Reset() {
	*x = SearchI_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Filter) ProtoMessage() {}

func (x *SearchI_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Filter.ProtoReflect.Descriptor instead.
func (*SearchI_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{1}
}

func (x *SearchI_Filter) GetPaging() *SearchI_Filter_Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type SearchI_Filter_Paging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind  string `protobuf:"bytes,100,opt,name=kind,proto3" json:"kind,omitempty"`
	Start string `protobuf:"bytes,200,opt,name=start,proto3" json:"start,omitempty"`
	Stop  string `protobuf:"bytes,300,opt,name=stop,proto3" json:"stop,omitempty"`
}

func (x *SearchI_Filter_Paging) Reset() {
	*x = SearchI_Filter_Paging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Filter_Paging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Filter_Paging) ProtoMessage() {}

func (x *SearchI_Filter_Paging) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Filter_Paging.ProtoReflect.Descriptor instead.
func (*SearchI_Filter_Paging) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{2}
}

func (x *SearchI_Filter_Paging) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *SearchI_Filter_Paging) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *SearchI_Filter_Paging) GetStop() string {
	if x != nil {
		return x.Stop
	}
	return ""
}

type SearchI_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern *SearchI_Object_Intern `protobuf:"bytes,100,opt,name=intern,proto3" json:"intern,omitempty"`
	Public *SearchI_Object_Public `protobuf:"bytes,200,opt,name=public,proto3" json:"public,omitempty"`
	Symbol *SearchI_Object_Symbol `protobuf:"bytes,300,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *SearchI_Object) Reset() {
	*x = SearchI_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Object) ProtoMessage() {}

func (x *SearchI_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Object.ProtoReflect.Descriptor instead.
func (*SearchI_Object) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{3}
}

func (x *SearchI_Object) GetIntern() *SearchI_Object_Intern {
	if x != nil {
		return x.Intern
	}
	return nil
}

func (x *SearchI_Object) GetPublic() *SearchI_Object_Public {
	if x != nil {
		return x.Public
	}
	return nil
}

func (x *SearchI_Object) GetSymbol() *SearchI_Object_Symbol {
	if x != nil {
		return x.Symbol
	}
	return nil
}

type SearchI_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the ID of the vote object being searched. If searching for a
	// particular ID, the search query object must not contain any other fields.
	Id string `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
	// owner is the ID of the user having created the votes being searched. If
	// searching for votes created by a particular owner, the search query object
	// may also define public.claim in order to search for votes a user has cast
	// on a specific claim. Note that owner may be set to "self" in order to set
	// the internal user ID by reference of the caller's access token.
	Owner string `protobuf:"bytes,200,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *SearchI_Object_Intern) Reset() {
	*x = SearchI_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Object_Intern) ProtoMessage() {}

func (x *SearchI_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Object_Intern.ProtoReflect.Descriptor instead.
func (*SearchI_Object_Intern) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{4}
}

func (x *SearchI_Object_Intern) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SearchI_Object_Intern) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type SearchI_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// claim is the ID of the claim that votes can be searched for. In other
	// words, providing a claim ID allows to search for all votes cast on the
	// specified claim. This search query can be further restricted by setting
	// intern.owner, which will then only return the votes that have been cast on
	// the given claim by that specific user.
	Claim string `protobuf:"bytes,100,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (x *SearchI_Object_Public) Reset() {
	*x = SearchI_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Object_Public) ProtoMessage() {}

func (x *SearchI_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Object_Public.ProtoReflect.Descriptor instead.
func (*SearchI_Object_Public) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{5}
}

func (x *SearchI_Object_Public) GetClaim() string {
	if x != nil {
		return x.Claim
	}
	return ""
}

type SearchI_Object_Symbol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchI_Object_Symbol) Reset() {
	*x = SearchI_Object_Symbol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Object_Symbol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Object_Symbol) ProtoMessage() {}

func (x *SearchI_Object_Symbol) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Object_Symbol.ProtoReflect.Descriptor instead.
func (*SearchI_Object_Symbol) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{6}
}

// SearchO is the output for searching votes.
//
//	{
//	    "object": [
//	        {
//	            "intern": {
//	                "created": "1689001255",
//	                "id": "778237",
//	                "owner": "551265"
//	            },
//	            "public": {
//	                "chain": "421614",
//	                "claim": "778237",
//	                "hash": "0x1234",
//	                "kind": "stake",
//	                "lifecycle": "onchain",
//	                "option": "true",
//	                "value": "1.5"
//	            }
//	        }
//	    ]
//	}
type SearchO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *SearchO_Filter   `protobuf:"bytes,100,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*SearchO_Object `protobuf:"bytes,200,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *SearchO) Reset() {
	*x = SearchO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO) ProtoMessage() {}

func (x *SearchO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO.ProtoReflect.Descriptor instead.
func (*SearchO) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{7}
}

func (x *SearchO) GetFilter() *SearchO_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *SearchO) GetObject() []*SearchO_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type SearchO_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchO_Filter) Reset() {
	*x = SearchO_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Filter) ProtoMessage() {}

func (x *SearchO_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Filter.ProtoReflect.Descriptor instead.
func (*SearchO_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{8}
}

type SearchO_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extern *SearchO_Object_Extern `protobuf:"bytes,100,opt,name=extern,proto3" json:"extern,omitempty"`
	Intern *SearchO_Object_Intern `protobuf:"bytes,200,opt,name=intern,proto3" json:"intern,omitempty"`
	Public *SearchO_Object_Public `protobuf:"bytes,300,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *SearchO_Object) Reset() {
	*x = SearchO_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Object) ProtoMessage() {}

func (x *SearchO_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Object.ProtoReflect.Descriptor instead.
func (*SearchO_Object) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{9}
}

func (x *SearchO_Object) GetExtern() *SearchO_Object_Extern {
	if x != nil {
		return x.Extern
	}
	return nil
}

func (x *SearchO_Object) GetIntern() *SearchO_Object_Intern {
	if x != nil {
		return x.Intern
	}
	return nil
}

func (x *SearchO_Object) GetPublic() *SearchO_Object_Public {
	if x != nil {
		return x.Public
	}
	return nil
}

type SearchO_Object_Extern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchO_Object_Extern) Reset() {
	*x = SearchO_Object_Extern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Object_Extern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Object_Extern) ProtoMessage() {}

func (x *SearchO_Object_Extern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Object_Extern.ProtoReflect.Descriptor instead.
func (*SearchO_Object_Extern) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{10}
}

type SearchO_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// created is the unix timestamp in seconds at which this vote got created.
	Created string `protobuf:"bytes,100,opt,name=created,proto3" json:"created,omitempty"`
	// id is the ID of the vote object being searched.
	Id string `protobuf:"bytes,200,opt,name=id,proto3" json:"id,omitempty"`
	// owner is the ID of the user who created this vote.
	Owner string `protobuf:"bytes,300,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *SearchO_Object_Intern) Reset() {
	*x = SearchO_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Object_Intern) ProtoMessage() {}

func (x *SearchO_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Object_Intern.ProtoReflect.Descriptor instead.
func (*SearchO_Object_Intern) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{11}
}

func (x *SearchO_Object_Intern) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *SearchO_Object_Intern) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SearchO_Object_Intern) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type SearchO_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// chain is the chain ID on which a vote got expressed.
	Chain string `protobuf:"bytes,100,opt,name=chain,proto3" json:"chain,omitempty"`
	// claim is the ID of the referenced claim being voted on. Note that the
	// lifecycle of the referenced claim must be compliant with the vote kind
	// provided.
	Claim string `protobuf:"bytes,200,opt,name=claim,proto3" json:"claim,omitempty"`
	// hash is the onchain transaction hash for this vote. Setting hash implies
	// that the vote got confirmed onchain, and with it the lifecycle phase
	// "onchain" will be inferred automatically.
	Hash string `protobuf:"bytes,300,opt,name=hash,proto3" json:"hash,omitempty"`
	// kind is the type of vote, e.g. "stake" or "truth" on which a vote is cast.
	// Note that kind must be compliant with the lifecycle of the referenced
	// claim.
	//
	//	"stake" defines votes for users expressing their own opinions. A vote
	//	may have set kind "stake" to express opinions on claims of either kind
	//	"adjourn", "dispute", "nullify" or "propose".
	//
	//	"truth" defines votes for users verifying real world events. A vote may
	//	have set kind "stake" to verify real events on claims of kind
	//	"resolve".
	Kind string `protobuf:"bytes,400,opt,name=kind,proto3" json:"kind,omitempty"`
	// lifecycle describes the vote's status within the system. This field may
	// look like one of the examples below.
	//
	//	pending
	//	onchain
	//
	// All votes start with the interim system status "pending". Those pending
	// votes were posted offchain, but have not yet been confirmed onchain.
	//
	//	"pending" describes votes that have only been posted offchain
	//
	//	"onchain" describes votes that have been confirmed onchain
	Lifecycle string `protobuf:"bytes,500,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// meta may contain onchain arbitrary meta data.
	Meta string `protobuf:"bytes,600,opt,name=meta,proto3" json:"meta,omitempty"`
	// option is the side of the vote being cast, e.g. "true" or "false". If
	// option is "true", then the vote is in agreement. If option is "false", then
	// the vote is in disagreement.
	Option string `protobuf:"bytes,700,opt,name=option,proto3" json:"option,omitempty"`
	// value is the weight of the vote being cast. If kind is "stake", then value
	// might be any positive number. If kind is "truth", then value must be 1.
	Value string `protobuf:"bytes,800,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SearchO_Object_Public) Reset() {
	*x = SearchO_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_vote_search_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Object_Public) ProtoMessage() {}

func (x *SearchO_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_vote_search_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Object_Public.ProtoReflect.Descriptor instead.
func (*SearchO_Object_Public) Descriptor() ([]byte, []int) {
	return file_pbf_vote_search_proto_rawDescGZIP(), []int{12}
}

func (x *SearchO_Object_Public) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *SearchO_Object_Public) GetClaim() string {
	if x != nil {
		return x.Claim
	}
	return ""
}

func (x *SearchO_Object_Public) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *SearchO_Object_Public) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *SearchO_Object_Public) GetLifecycle() string {
	if x != nil {
		return x.Lifecycle
	}
	return ""
}

func (x *SearchO_Object_Public) GetMeta() string {
	if x != nil {
		return x.Meta
	}
	return ""
}

func (x *SearchO_Object_Public) GetOption() string {
	if x != nil {
		return x.Option
	}
	return ""
}

func (x *SearchO_Object_Public) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_pbf_vote_search_proto protoreflect.FileDescriptor

var file_pbf_vote_search_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x66, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x22, 0x66, 0x0a,
	0x07, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x45, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49,
	0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x57, 0x0a, 0x15,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x15, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x13, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x74, 0x6f, 0x70, 0x22, 0xb1, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x34, 0x0a,
	0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0xac, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x3e, 0x0a, 0x15, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x15, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0xc8, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x15, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x22, 0x66, 0x0a, 0x07, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x12, 0x2c, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76,
	0x6f, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f,
	0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xb1, 0x01, 0x0a, 0x0e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x33,
	0x0a, 0x06, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x06, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x18, 0xc8, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x6f, 0x74,
	0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22,
	0x17, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x22, 0x59, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0f, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x22, 0xd2, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x12, 0x15, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0xc8, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x13, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12,
	0x13, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x90, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x18, 0xf4, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x12, 0x13, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0xd8, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0xbc, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x15, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0xa0, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x76,
	0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbf_vote_search_proto_rawDescOnce sync.Once
	file_pbf_vote_search_proto_rawDescData = file_pbf_vote_search_proto_rawDesc
)

func file_pbf_vote_search_proto_rawDescGZIP() []byte {
	file_pbf_vote_search_proto_rawDescOnce.Do(func() {
		file_pbf_vote_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_vote_search_proto_rawDescData)
	})
	return file_pbf_vote_search_proto_rawDescData
}

var file_pbf_vote_search_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pbf_vote_search_proto_goTypes = []interface{}{
	(*SearchI)(nil),               // 0: vote.SearchI
	(*SearchI_Filter)(nil),        // 1: vote.SearchI_Filter
	(*SearchI_Filter_Paging)(nil), // 2: vote.SearchI_Filter_Paging
	(*SearchI_Object)(nil),        // 3: vote.SearchI_Object
	(*SearchI_Object_Intern)(nil), // 4: vote.SearchI_Object_Intern
	(*SearchI_Object_Public)(nil), // 5: vote.SearchI_Object_Public
	(*SearchI_Object_Symbol)(nil), // 6: vote.SearchI_Object_Symbol
	(*SearchO)(nil),               // 7: vote.SearchO
	(*SearchO_Filter)(nil),        // 8: vote.SearchO_Filter
	(*SearchO_Object)(nil),        // 9: vote.SearchO_Object
	(*SearchO_Object_Extern)(nil), // 10: vote.SearchO_Object_Extern
	(*SearchO_Object_Intern)(nil), // 11: vote.SearchO_Object_Intern
	(*SearchO_Object_Public)(nil), // 12: vote.SearchO_Object_Public
}
var file_pbf_vote_search_proto_depIdxs = []int32{
	1,  // 0: vote.SearchI.filter:type_name -> vote.SearchI_Filter
	3,  // 1: vote.SearchI.object:type_name -> vote.SearchI_Object
	2,  // 2: vote.SearchI_Filter.paging:type_name -> vote.SearchI_Filter_Paging
	4,  // 3: vote.SearchI_Object.intern:type_name -> vote.SearchI_Object_Intern
	5,  // 4: vote.SearchI_Object.public:type_name -> vote.SearchI_Object_Public
	6,  // 5: vote.SearchI_Object.symbol:type_name -> vote.SearchI_Object_Symbol
	8,  // 6: vote.SearchO.filter:type_name -> vote.SearchO_Filter
	9,  // 7: vote.SearchO.object:type_name -> vote.SearchO_Object
	10, // 8: vote.SearchO_Object.extern:type_name -> vote.SearchO_Object_Extern
	11, // 9: vote.SearchO_Object.intern:type_name -> vote.SearchO_Object_Intern
	12, // 10: vote.SearchO_Object.public:type_name -> vote.SearchO_Object_Public
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_pbf_vote_search_proto_init() }
func file_pbf_vote_search_proto_init() {
	if File_pbf_vote_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_vote_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI); i {
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
		file_pbf_vote_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Filter); i {
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
		file_pbf_vote_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Filter_Paging); i {
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
		file_pbf_vote_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Object); i {
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
		file_pbf_vote_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Object_Intern); i {
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
		file_pbf_vote_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Object_Public); i {
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
		file_pbf_vote_search_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Object_Symbol); i {
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
		file_pbf_vote_search_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO); i {
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
		file_pbf_vote_search_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Filter); i {
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
		file_pbf_vote_search_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Object); i {
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
		file_pbf_vote_search_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Object_Extern); i {
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
		file_pbf_vote_search_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Object_Intern); i {
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
		file_pbf_vote_search_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Object_Public); i {
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
			RawDescriptor: file_pbf_vote_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_vote_search_proto_goTypes,
		DependencyIndexes: file_pbf_vote_search_proto_depIdxs,
		MessageInfos:      file_pbf_vote_search_proto_msgTypes,
	}.Build()
	File_pbf_vote_search_proto = out.File
	file_pbf_vote_search_proto_rawDesc = nil
	file_pbf_vote_search_proto_goTypes = nil
	file_pbf_vote_search_proto_depIdxs = nil
}
