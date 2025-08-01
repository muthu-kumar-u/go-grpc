// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/user/user.proto

package user

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	error1 "github.com/muthu-kumar-u/go-grpc/gen/go/proto/error"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ─── Requests ─────────────────────────────────────
type CreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Number        string                 `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_proto_user_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type ReadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	mi := &file_proto_user_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *ReadRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Number        string                 `protobuf:"bytes,4,opt,name=number,proto3" json:"number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	mi := &file_proto_user_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_proto_user_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// ─── Responses ─────────────────────────────────────
type CreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Error         *error1.Error          `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_proto_user_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *CreateResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateResponse) GetError() *error1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type ReadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Number        string                 `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	Error         *error1.Error          `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	mi := &file_proto_user_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *ReadResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReadResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ReadResponse) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *ReadResponse) GetError() *error1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Error         *error1.Error          `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	mi := &file_proto_user_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateResponse) GetError() *error1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Error         *error1.Error          `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_proto_user_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteResponse) GetError() *error1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_proto_user_user_proto protoreflect.FileDescriptor

const file_proto_user_user_proto_rawDesc = "" +
	"\n" +
	"\x15proto/user/user.proto\x12\x04user\x1a\x1cgoogle/api/annotations.proto\x1a\x17validate/validate.proto\x1a\x17proto/error/error.proto\"w\n" +
	"\rCreateRequest\x12\x1b\n" +
	"\x04name\x18\x01 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\x04name\x12\x1d\n" +
	"\x05email\x18\x02 \x01(\tB\a\xfaB\x04r\x02`\x01R\x05email\x12*\n" +
	"\x06number\x18\x03 \x01(\tB\x12\xfaB\x0fr\r2\b^[0-9]+$\x98\x01\n" +
	"R\x06number\"&\n" +
	"\vReadRequest\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x05B\a\xfaB\x04\x1a\x02 \x00R\x02id\"\x90\x01\n" +
	"\rUpdateRequest\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x05B\a\xfaB\x04\x1a\x02 \x00R\x02id\x12\x1b\n" +
	"\x04name\x18\x02 \x01(\tB\a\xfaB\x04r\x02\x10\x01R\x04name\x12\x1d\n" +
	"\x05email\x18\x03 \x01(\tB\a\xfaB\x04r\x02`\x01R\x05email\x12*\n" +
	"\x06number\x18\x04 \x01(\tB\x12\xfaB\x0fr\r2\b^[0-9]+$\x98\x01\n" +
	"R\x06number\"(\n" +
	"\rDeleteRequest\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x05B\a\xfaB\x04\x1a\x02 \x00R\x02id\"M\n" +
	"\x0eCreateResponse\x12\x17\n" +
	"\x02id\x18\x01 \x01(\x05B\a\xfaB\x04\x1a\x02 \x00R\x02id\x12\"\n" +
	"\x05error\x18\x02 \x01(\v2\f.error.ErrorR\x05error\"t\n" +
	"\fReadResponse\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x16\n" +
	"\x06number\x18\x03 \x01(\tR\x06number\x12\"\n" +
	"\x05error\x18\x04 \x01(\v2\f.error.ErrorR\x05error\"D\n" +
	"\x0eUpdateResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\"\n" +
	"\x05error\x18\x02 \x01(\v2\f.error.ErrorR\x05error\"N\n" +
	"\x0eDeleteResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12\"\n" +
	"\x05error\x18\x02 \x01(\v2\f.error.ErrorR\x05error2\xbc\x02\n" +
	"\vUserService\x12I\n" +
	"\x06Create\x12\x13.user.CreateRequest\x1a\x14.user.CreateResponse\"\x14\x82\xd3\xe4\x93\x02\x0e:\x01*\"\t/v1/users\x12E\n" +
	"\x04Read\x12\x11.user.ReadRequest\x1a\x12.user.ReadResponse\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/v1/users/{id}\x12N\n" +
	"\x06Update\x12\x13.user.UpdateRequest\x1a\x14.user.UpdateResponse\"\x19\x82\xd3\xe4\x93\x02\x13:\x01*\x1a\x0e/v1/users/{id}\x12K\n" +
	"\x06Delete\x12\x13.user.DeleteRequest\x1a\x14.user.DeleteResponse\"\x16\x82\xd3\xe4\x93\x02\x10*\x0e/v1/users/{id}B9Z7github.com/muthu-kumar-u/go-grpc/gen/go/proto/user;userb\x06proto3"

var (
	file_proto_user_user_proto_rawDescOnce sync.Once
	file_proto_user_user_proto_rawDescData []byte
)

func file_proto_user_user_proto_rawDescGZIP() []byte {
	file_proto_user_user_proto_rawDescOnce.Do(func() {
		file_proto_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_user_user_proto_rawDesc), len(file_proto_user_user_proto_rawDesc)))
	})
	return file_proto_user_user_proto_rawDescData
}

var file_proto_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_user_user_proto_goTypes = []any{
	(*CreateRequest)(nil),  // 0: user.CreateRequest
	(*ReadRequest)(nil),    // 1: user.ReadRequest
	(*UpdateRequest)(nil),  // 2: user.UpdateRequest
	(*DeleteRequest)(nil),  // 3: user.DeleteRequest
	(*CreateResponse)(nil), // 4: user.CreateResponse
	(*ReadResponse)(nil),   // 5: user.ReadResponse
	(*UpdateResponse)(nil), // 6: user.UpdateResponse
	(*DeleteResponse)(nil), // 7: user.DeleteResponse
	(*error1.Error)(nil),   // 8: error.Error
}
var file_proto_user_user_proto_depIdxs = []int32{
	8, // 0: user.CreateResponse.error:type_name -> error.Error
	8, // 1: user.ReadResponse.error:type_name -> error.Error
	8, // 2: user.UpdateResponse.error:type_name -> error.Error
	8, // 3: user.DeleteResponse.error:type_name -> error.Error
	0, // 4: user.UserService.Create:input_type -> user.CreateRequest
	1, // 5: user.UserService.Read:input_type -> user.ReadRequest
	2, // 6: user.UserService.Update:input_type -> user.UpdateRequest
	3, // 7: user.UserService.Delete:input_type -> user.DeleteRequest
	4, // 8: user.UserService.Create:output_type -> user.CreateResponse
	5, // 9: user.UserService.Read:output_type -> user.ReadResponse
	6, // 10: user.UserService.Update:output_type -> user.UpdateResponse
	7, // 11: user.UserService.Delete:output_type -> user.DeleteResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_user_user_proto_init() }
func file_proto_user_user_proto_init() {
	if File_proto_user_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_user_user_proto_rawDesc), len(file_proto_user_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_user_proto_goTypes,
		DependencyIndexes: file_proto_user_user_proto_depIdxs,
		MessageInfos:      file_proto_user_user_proto_msgTypes,
	}.Build()
	File_proto_user_user_proto = out.File
	file_proto_user_user_proto_goTypes = nil
	file_proto_user_user_proto_depIdxs = nil
}
