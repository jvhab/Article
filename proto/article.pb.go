// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/article.proto

package proto

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

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Body        string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Counts      int32    `protobuf:"varint,4,opt,name=counts,proto3" json:"counts,omitempty"`
	TagList     []string `protobuf:"bytes,5,rep,name=tagList,proto3" json:"tagList,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_proto_article_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Article) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Article) GetCounts() int32 {
	if x != nil {
		return x.Counts
	}
	return 0
}

func (x *Article) GetTagList() []string {
	if x != nil {
		return x.TagList
	}
	return nil
}

type CreateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *CreateArticleRequest) Reset() {
	*x = CreateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleRequest) ProtoMessage() {}

func (x *CreateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleRequest.ProtoReflect.Descriptor instead.
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return file_proto_article_proto_rawDescGZIP(), []int{1}
}

func (x *CreateArticleRequest) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type CreateArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateArticleResponse) Reset() {
	*x = CreateArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleResponse) ProtoMessage() {}

func (x *CreateArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleResponse.ProtoReflect.Descriptor instead.
func (*CreateArticleResponse) Descriptor() ([]byte, []int) {
	return file_proto_article_proto_rawDescGZIP(), []int{2}
}

func (x *CreateArticleResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GetArticleRequest) Reset() {
	*x = GetArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleRequest) ProtoMessage() {}

func (x *GetArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleRequest.ProtoReflect.Descriptor instead.
func (*GetArticleRequest) Descriptor() ([]byte, []int) {
	return file_proto_article_proto_rawDescGZIP(), []int{3}
}

func (x *GetArticleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *GetArticleResponse) Reset() {
	*x = GetArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleResponse) ProtoMessage() {}

func (x *GetArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleResponse.ProtoReflect.Descriptor instead.
func (*GetArticleResponse) Descriptor() ([]byte, []int) {
	return file_proto_article_proto_rawDescGZIP(), []int{4}
}

func (x *GetArticleResponse) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type UpdateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title *Article `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UpdateArticleRequest) Reset() {
	*x = UpdateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleRequest) ProtoMessage() {}

func (x *UpdateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleRequest.ProtoReflect.Descriptor instead.
func (*UpdateArticleRequest) Descriptor() ([]byte, []int) {
	return file_proto_article_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateArticleRequest) GetTitle() *Article {
	if x != nil {
		return x.Title
	}
	return nil
}

type UpdateArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *UpdateArticleResponse) Reset() {
	*x = UpdateArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleResponse) ProtoMessage() {}

func (x *UpdateArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleResponse.ProtoReflect.Descriptor instead.
func (*UpdateArticleResponse) Descriptor() ([]byte, []int) {
	return file_proto_article_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateArticleResponse) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type DeleteArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *DeleteArticleRequest) Reset() {
	*x = DeleteArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleRequest) ProtoMessage() {}

func (x *DeleteArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleRequest.ProtoReflect.Descriptor instead.
func (*DeleteArticleRequest) Descriptor() ([]byte, []int) {
	return file_proto_article_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteArticleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type DeleteArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *DeleteArticleResponse) Reset() {
	*x = DeleteArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_article_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleResponse) ProtoMessage() {}

func (x *DeleteArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_article_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleResponse.ProtoReflect.Descriptor instead.
func (*DeleteArticleResponse) Descriptor() ([]byte, []int) {
	return file_proto_article_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteArticleResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_proto_article_proto protoreflect.FileDescriptor

var file_proto_article_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x87,
	0x01, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x2d, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x3e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x43, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x2c, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x2d, 0x0a, 0x15, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x32, 0xc7, 0x02, 0x0a, 0x0e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1d,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_article_proto_rawDescOnce sync.Once
	file_proto_article_proto_rawDescData = file_proto_article_proto_rawDesc
)

func file_proto_article_proto_rawDescGZIP() []byte {
	file_proto_article_proto_rawDescOnce.Do(func() {
		file_proto_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_article_proto_rawDescData)
	})
	return file_proto_article_proto_rawDescData
}

var file_proto_article_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_article_proto_goTypes = []interface{}{
	(*Article)(nil),               // 0: article.Article
	(*CreateArticleRequest)(nil),  // 1: article.CreateArticleRequest
	(*CreateArticleResponse)(nil), // 2: article.CreateArticleResponse
	(*GetArticleRequest)(nil),     // 3: article.GetArticleRequest
	(*GetArticleResponse)(nil),    // 4: article.GetArticleResponse
	(*UpdateArticleRequest)(nil),  // 5: article.UpdateArticleRequest
	(*UpdateArticleResponse)(nil), // 6: article.UpdateArticleResponse
	(*DeleteArticleRequest)(nil),  // 7: article.DeleteArticleRequest
	(*DeleteArticleResponse)(nil), // 8: article.DeleteArticleResponse
}
var file_proto_article_proto_depIdxs = []int32{
	0, // 0: article.CreateArticleRequest.article:type_name -> article.Article
	0, // 1: article.GetArticleResponse.article:type_name -> article.Article
	0, // 2: article.UpdateArticleRequest.title:type_name -> article.Article
	0, // 3: article.UpdateArticleResponse.article:type_name -> article.Article
	1, // 4: article.ArticleService.CreateArticle:input_type -> article.CreateArticleRequest
	3, // 5: article.ArticleService.GetArticle:input_type -> article.GetArticleRequest
	5, // 6: article.ArticleService.UpdateArticle:input_type -> article.UpdateArticleRequest
	7, // 7: article.ArticleService.DeleteArticle:input_type -> article.DeleteArticleRequest
	2, // 8: article.ArticleService.CreateArticle:output_type -> article.CreateArticleResponse
	4, // 9: article.ArticleService.GetArticle:output_type -> article.GetArticleResponse
	6, // 10: article.ArticleService.UpdateArticle:output_type -> article.UpdateArticleResponse
	8, // 11: article.ArticleService.DeleteArticle:output_type -> article.DeleteArticleResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_article_proto_init() }
func file_proto_article_proto_init() {
	if File_proto_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_proto_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleRequest); i {
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
		file_proto_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleResponse); i {
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
		file_proto_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleRequest); i {
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
		file_proto_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleResponse); i {
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
		file_proto_article_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleRequest); i {
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
		file_proto_article_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleResponse); i {
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
		file_proto_article_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleRequest); i {
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
		file_proto_article_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleResponse); i {
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
			RawDescriptor: file_proto_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_article_proto_goTypes,
		DependencyIndexes: file_proto_article_proto_depIdxs,
		MessageInfos:      file_proto_article_proto_msgTypes,
	}.Build()
	File_proto_article_proto = out.File
	file_proto_article_proto_rawDesc = nil
	file_proto_article_proto_goTypes = nil
	file_proto_article_proto_depIdxs = nil
}