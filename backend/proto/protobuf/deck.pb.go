// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: deck.proto

package protobuf

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

// Deck represents a flash card deck
type Deck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckId    string `protobuf:"bytes,1,opt,name=deckId,proto3" json:"deckId,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner     string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Deck) Reset() {
	*x = Deck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deck) ProtoMessage() {}

func (x *Deck) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deck.ProtoReflect.Descriptor instead.
func (*Deck) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{0}
}

func (x *Deck) GetDeckId() string {
	if x != nil {
		return x.DeckId
	}
	return ""
}

func (x *Deck) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Deck) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Deck) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// CreateDeckRequest represents a request to create a new deck
type CreateDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *CreateDeckRequest) Reset() {
	*x = CreateDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeckRequest) ProtoMessage() {}

func (x *CreateDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeckRequest.ProtoReflect.Descriptor instead.
func (*CreateDeckRequest) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDeckRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDeckRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

// CreateDeckResponse represents a response for deck creation
type CreateDeckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deck *Deck `protobuf:"bytes,1,opt,name=deck,proto3" json:"deck,omitempty"`
}

func (x *CreateDeckResponse) Reset() {
	*x = CreateDeckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeckResponse) ProtoMessage() {}

func (x *CreateDeckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeckResponse.ProtoReflect.Descriptor instead.
func (*CreateDeckResponse) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDeckResponse) GetDeck() *Deck {
	if x != nil {
		return x.Deck
	}
	return nil
}

// GetDeckRequest represents a request to get a deck
type GetDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckId string `protobuf:"bytes,1,opt,name=deckId,proto3" json:"deckId,omitempty"`
	Owner  string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *GetDeckRequest) Reset() {
	*x = GetDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeckRequest) ProtoMessage() {}

func (x *GetDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeckRequest.ProtoReflect.Descriptor instead.
func (*GetDeckRequest) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeckRequest) GetDeckId() string {
	if x != nil {
		return x.DeckId
	}
	return ""
}

func (x *GetDeckRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

// GetDeckResponse represents a response for getting a deck
type GetDeckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deck *Deck `protobuf:"bytes,1,opt,name=deck,proto3" json:"deck,omitempty"`
}

func (x *GetDeckResponse) Reset() {
	*x = GetDeckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeckResponse) ProtoMessage() {}

func (x *GetDeckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeckResponse.ProtoReflect.Descriptor instead.
func (*GetDeckResponse) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{4}
}

func (x *GetDeckResponse) GetDeck() *Deck {
	if x != nil {
		return x.Deck
	}
	return nil
}

// ListDecksRequest represents a request to list decks
type ListDecksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *ListDecksRequest) Reset() {
	*x = ListDecksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDecksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDecksRequest) ProtoMessage() {}

func (x *ListDecksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDecksRequest.ProtoReflect.Descriptor instead.
func (*ListDecksRequest) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{5}
}

func (x *ListDecksRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

// ListDecksResponse represents a response for listing decks
type ListDecksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decks []*Deck `protobuf:"bytes,1,rep,name=decks,proto3" json:"decks,omitempty"`
}

func (x *ListDecksResponse) Reset() {
	*x = ListDecksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDecksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDecksResponse) ProtoMessage() {}

func (x *ListDecksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDecksResponse.ProtoReflect.Descriptor instead.
func (*ListDecksResponse) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{6}
}

func (x *ListDecksResponse) GetDecks() []*Deck {
	if x != nil {
		return x.Decks
	}
	return nil
}

// UpdateDeckRequest represents a request to update a deck
type UpdateDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckId string `protobuf:"bytes,1,opt,name=deckId,proto3" json:"deckId,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner  string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *UpdateDeckRequest) Reset() {
	*x = UpdateDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeckRequest) ProtoMessage() {}

func (x *UpdateDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeckRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeckRequest) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateDeckRequest) GetDeckId() string {
	if x != nil {
		return x.DeckId
	}
	return ""
}

func (x *UpdateDeckRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDeckRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

// UpdateDeckResponse represents a response for updating a deck
type UpdateDeckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deck *Deck `protobuf:"bytes,1,opt,name=deck,proto3" json:"deck,omitempty"`
}

func (x *UpdateDeckResponse) Reset() {
	*x = UpdateDeckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeckResponse) ProtoMessage() {}

func (x *UpdateDeckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeckResponse.ProtoReflect.Descriptor instead.
func (*UpdateDeckResponse) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateDeckResponse) GetDeck() *Deck {
	if x != nil {
		return x.Deck
	}
	return nil
}

// DeleteDeckRequest represents a request to delete a deck
type DeleteDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckId string `protobuf:"bytes,1,opt,name=deckId,proto3" json:"deckId,omitempty"`
	Owner  string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *DeleteDeckRequest) Reset() {
	*x = DeleteDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeckRequest) ProtoMessage() {}

func (x *DeleteDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeckRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeckRequest) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteDeckRequest) GetDeckId() string {
	if x != nil {
		return x.DeckId
	}
	return ""
}

func (x *DeleteDeckRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

// DeleteDeckResponse represents a response for deleting a deck
type DeleteDeckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteDeckResponse) Reset() {
	*x = DeleteDeckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeckResponse) ProtoMessage() {}

func (x *DeleteDeckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeckResponse.ProtoReflect.Descriptor instead.
func (*DeleteDeckResponse) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteDeckResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_deck_proto protoreflect.FileDescriptor

var file_deck_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x65,
	0x63, 0x6b, 0x22, 0x66, 0x0a, 0x04, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65,
	0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x63, 0x6b,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3d, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x22,
	0x3e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22,
	0x31, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x04, 0x64, 0x65,
	0x63, 0x6b, 0x22, 0x28, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x05, 0x64, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x05, 0x64, 0x65,
	0x63, 0x6b, 0x73, 0x22, 0x55, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x63, 0x6b,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x04, 0x64, 0x65, 0x63, 0x6b,
	0x22, 0x41, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0x9f, 0x02, 0x0a, 0x0b, 0x44, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63,
	0x6b, 0x12, 0x17, 0x2e, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x64, 0x65, 0x63,
	0x6b, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63,
	0x6b, 0x12, 0x14, 0x2e, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44,
	0x65, 0x63, 0x6b, 0x12, 0x3c, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x73,
	0x12, 0x16, 0x2e, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64, 0x65, 0x63, 0x6b, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x12,
	0x17, 0x2e, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x64, 0x65, 0x63, 0x6b, 0x2e,
	0x44, 0x65, 0x63, 0x6b, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65,
	0x63, 0x6b, 0x12, 0x17, 0x2e, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x65,
	0x63, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deck_proto_rawDescOnce sync.Once
	file_deck_proto_rawDescData = file_deck_proto_rawDesc
)

func file_deck_proto_rawDescGZIP() []byte {
	file_deck_proto_rawDescOnce.Do(func() {
		file_deck_proto_rawDescData = protoimpl.X.CompressGZIP(file_deck_proto_rawDescData)
	})
	return file_deck_proto_rawDescData
}

var file_deck_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_deck_proto_goTypes = []interface{}{
	(*Deck)(nil),               // 0: deck.Deck
	(*CreateDeckRequest)(nil),  // 1: deck.CreateDeckRequest
	(*CreateDeckResponse)(nil), // 2: deck.CreateDeckResponse
	(*GetDeckRequest)(nil),     // 3: deck.GetDeckRequest
	(*GetDeckResponse)(nil),    // 4: deck.GetDeckResponse
	(*ListDecksRequest)(nil),   // 5: deck.ListDecksRequest
	(*ListDecksResponse)(nil),  // 6: deck.ListDecksResponse
	(*UpdateDeckRequest)(nil),  // 7: deck.UpdateDeckRequest
	(*UpdateDeckResponse)(nil), // 8: deck.UpdateDeckResponse
	(*DeleteDeckRequest)(nil),  // 9: deck.DeleteDeckRequest
	(*DeleteDeckResponse)(nil), // 10: deck.DeleteDeckResponse
}
var file_deck_proto_depIdxs = []int32{
	0,  // 0: deck.CreateDeckResponse.deck:type_name -> deck.Deck
	0,  // 1: deck.GetDeckResponse.deck:type_name -> deck.Deck
	0,  // 2: deck.ListDecksResponse.decks:type_name -> deck.Deck
	0,  // 3: deck.UpdateDeckResponse.deck:type_name -> deck.Deck
	1,  // 4: deck.DeckService.CreateDeck:input_type -> deck.CreateDeckRequest
	3,  // 5: deck.DeckService.GetDeck:input_type -> deck.GetDeckRequest
	5,  // 6: deck.DeckService.ListDecks:input_type -> deck.ListDecksRequest
	7,  // 7: deck.DeckService.UpdateDeck:input_type -> deck.UpdateDeckRequest
	9,  // 8: deck.DeckService.DeleteDeck:input_type -> deck.DeleteDeckRequest
	0,  // 9: deck.DeckService.CreateDeck:output_type -> deck.Deck
	0,  // 10: deck.DeckService.GetDeck:output_type -> deck.Deck
	6,  // 11: deck.DeckService.ListDecks:output_type -> deck.ListDecksResponse
	0,  // 12: deck.DeckService.UpdateDeck:output_type -> deck.Deck
	10, // 13: deck.DeckService.DeleteDeck:output_type -> deck.DeleteDeckResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_deck_proto_init() }
func file_deck_proto_init() {
	if File_deck_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deck_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deck); i {
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
		file_deck_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeckRequest); i {
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
		file_deck_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeckResponse); i {
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
		file_deck_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeckRequest); i {
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
		file_deck_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeckResponse); i {
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
		file_deck_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDecksRequest); i {
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
		file_deck_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDecksResponse); i {
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
		file_deck_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeckRequest); i {
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
		file_deck_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeckResponse); i {
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
		file_deck_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeckRequest); i {
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
		file_deck_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeckResponse); i {
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
			RawDescriptor: file_deck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deck_proto_goTypes,
		DependencyIndexes: file_deck_proto_depIdxs,
		MessageInfos:      file_deck_proto_msgTypes,
	}.Build()
	File_deck_proto = out.File
	file_deck_proto_rawDesc = nil
	file_deck_proto_goTypes = nil
	file_deck_proto_depIdxs = nil
}
