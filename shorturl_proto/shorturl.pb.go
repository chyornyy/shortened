// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: shorturl_proto/shorturl.proto

package shorturl

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

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shorturl    string `protobuf:"bytes,1,opt,name=shorturl,proto3" json:"shorturl,omitempty"`
	Originalurl string `protobuf:"bytes,2,opt,name=originalurl,proto3" json:"originalurl,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorturl_proto_shorturl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_proto_shorturl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_shorturl_proto_shorturl_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetShorturl() string {
	if x != nil {
		return x.Shorturl
	}
	return ""
}

func (x *Link) GetOriginalurl() string {
	if x != nil {
		return x.Originalurl
	}
	return ""
}

type CreateLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link *Link `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *CreateLinkRequest) Reset() {
	*x = CreateLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorturl_proto_shorturl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkRequest) ProtoMessage() {}

func (x *CreateLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_proto_shorturl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateLinkRequest) Descriptor() ([]byte, []int) {
	return file_shorturl_proto_shorturl_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLinkRequest) GetLink() *Link {
	if x != nil {
		return x.Link
	}
	return nil
}

type CreateLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link *Link `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *CreateLinkResponse) Reset() {
	*x = CreateLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorturl_proto_shorturl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkResponse) ProtoMessage() {}

func (x *CreateLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_proto_shorturl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkResponse.ProtoReflect.Descriptor instead.
func (*CreateLinkResponse) Descriptor() ([]byte, []int) {
	return file_shorturl_proto_shorturl_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLinkResponse) GetLink() *Link {
	if x != nil {
		return x.Link
	}
	return nil
}

type ReadLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shorturl string `protobuf:"bytes,1,opt,name=shorturl,proto3" json:"shorturl,omitempty"`
}

func (x *ReadLinkRequest) Reset() {
	*x = ReadLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorturl_proto_shorturl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadLinkRequest) ProtoMessage() {}

func (x *ReadLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_proto_shorturl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadLinkRequest.ProtoReflect.Descriptor instead.
func (*ReadLinkRequest) Descriptor() ([]byte, []int) {
	return file_shorturl_proto_shorturl_proto_rawDescGZIP(), []int{3}
}

func (x *ReadLinkRequest) GetShorturl() string {
	if x != nil {
		return x.Shorturl
	}
	return ""
}

type ReadLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link *Link `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *ReadLinkResponse) Reset() {
	*x = ReadLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorturl_proto_shorturl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadLinkResponse) ProtoMessage() {}

func (x *ReadLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_proto_shorturl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadLinkResponse.ProtoReflect.Descriptor instead.
func (*ReadLinkResponse) Descriptor() ([]byte, []int) {
	return file_shorturl_proto_shorturl_proto_rawDescGZIP(), []int{4}
}

func (x *ReadLinkResponse) GetLink() *Link {
	if x != nil {
		return x.Link
	}
	return nil
}

type DeleteLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shorturl string `protobuf:"bytes,1,opt,name=shorturl,proto3" json:"shorturl,omitempty"`
}

func (x *DeleteLinkRequest) Reset() {
	*x = DeleteLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorturl_proto_shorturl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLinkRequest) ProtoMessage() {}

func (x *DeleteLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_proto_shorturl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLinkRequest.ProtoReflect.Descriptor instead.
func (*DeleteLinkRequest) Descriptor() ([]byte, []int) {
	return file_shorturl_proto_shorturl_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteLinkRequest) GetShorturl() string {
	if x != nil {
		return x.Shorturl
	}
	return ""
}

type DeleteLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteLinkResponse) Reset() {
	*x = DeleteLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorturl_proto_shorturl_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLinkResponse) ProtoMessage() {}

func (x *DeleteLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_proto_shorturl_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLinkResponse.ProtoReflect.Descriptor instead.
func (*DeleteLinkResponse) Descriptor() ([]byte, []int) {
	return file_shorturl_proto_shorturl_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteLinkResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_shorturl_proto_shorturl_proto protoreflect.FileDescriptor

var file_shorturl_proto_shorturl_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x75, 0x72, 0x6c, 0x22, 0x34, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x22, 0x35, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x2d, 0x0a, 0x0f, 0x52, 0x65, 0x61,
	0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x22, 0x33, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x2f, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x22, 0x2e,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xd5,
	0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x64, 0x2e, 0x72, 0x75, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shorturl_proto_shorturl_proto_rawDescOnce sync.Once
	file_shorturl_proto_shorturl_proto_rawDescData = file_shorturl_proto_shorturl_proto_rawDesc
)

func file_shorturl_proto_shorturl_proto_rawDescGZIP() []byte {
	file_shorturl_proto_shorturl_proto_rawDescOnce.Do(func() {
		file_shorturl_proto_shorturl_proto_rawDescData = protoimpl.X.CompressGZIP(file_shorturl_proto_shorturl_proto_rawDescData)
	})
	return file_shorturl_proto_shorturl_proto_rawDescData
}

var file_shorturl_proto_shorturl_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_shorturl_proto_shorturl_proto_goTypes = []interface{}{
	(*Link)(nil),               // 0: proto.Link
	(*CreateLinkRequest)(nil),  // 1: proto.CreateLinkRequest
	(*CreateLinkResponse)(nil), // 2: proto.CreateLinkResponse
	(*ReadLinkRequest)(nil),    // 3: proto.ReadLinkRequest
	(*ReadLinkResponse)(nil),   // 4: proto.ReadLinkResponse
	(*DeleteLinkRequest)(nil),  // 5: proto.DeleteLinkRequest
	(*DeleteLinkResponse)(nil), // 6: proto.DeleteLinkResponse
}
var file_shorturl_proto_shorturl_proto_depIdxs = []int32{
	0, // 0: proto.CreateLinkRequest.link:type_name -> proto.Link
	0, // 1: proto.CreateLinkResponse.link:type_name -> proto.Link
	0, // 2: proto.ReadLinkResponse.link:type_name -> proto.Link
	1, // 3: proto.LinkService.CreateLink:input_type -> proto.CreateLinkRequest
	3, // 4: proto.LinkService.GetLink:input_type -> proto.ReadLinkRequest
	5, // 5: proto.LinkService.DeleteLink:input_type -> proto.DeleteLinkRequest
	2, // 6: proto.LinkService.CreateLink:output_type -> proto.CreateLinkResponse
	4, // 7: proto.LinkService.GetLink:output_type -> proto.ReadLinkResponse
	6, // 8: proto.LinkService.DeleteLink:output_type -> proto.DeleteLinkResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_shorturl_proto_shorturl_proto_init() }
func file_shorturl_proto_shorturl_proto_init() {
	if File_shorturl_proto_shorturl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shorturl_proto_shorturl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
		file_shorturl_proto_shorturl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinkRequest); i {
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
		file_shorturl_proto_shorturl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinkResponse); i {
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
		file_shorturl_proto_shorturl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadLinkRequest); i {
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
		file_shorturl_proto_shorturl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadLinkResponse); i {
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
		file_shorturl_proto_shorturl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLinkRequest); i {
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
		file_shorturl_proto_shorturl_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLinkResponse); i {
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
			RawDescriptor: file_shorturl_proto_shorturl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shorturl_proto_shorturl_proto_goTypes,
		DependencyIndexes: file_shorturl_proto_shorturl_proto_depIdxs,
		MessageInfos:      file_shorturl_proto_shorturl_proto_msgTypes,
	}.Build()
	File_shorturl_proto_shorturl_proto = out.File
	file_shorturl_proto_shorturl_proto_rawDesc = nil
	file_shorturl_proto_shorturl_proto_goTypes = nil
	file_shorturl_proto_shorturl_proto_depIdxs = nil
}