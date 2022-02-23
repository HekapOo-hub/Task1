// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: human_service.proto

package humanpb

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

type Human struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required,max=40"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" validate:"required,max=40"`
	// @inject_tag: validate:"required"
	Age int32 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	Male bool `protobuf:"varint,3,opt,name=male,proto3" json:"male,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	ID string `protobuf:"bytes,4,opt,name=ID,proto3" json:"ID,omitempty" validate:"required"`
}

func (x *Human) Reset() {
	*x = Human{}
	if protoimpl.UnsafeEnabled {
		mi := &file_human_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Human) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Human) ProtoMessage() {}

func (x *Human) ProtoReflect() protoreflect.Message {
	mi := &file_human_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Human.ProtoReflect.Descriptor instead.
func (*Human) Descriptor() ([]byte, []int) {
	return file_human_service_proto_rawDescGZIP(), []int{0}
}

func (x *Human) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Human) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Human) GetMale() bool {
	if x != nil {
		return x.Male
	}
	return false
}

func (x *Human) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type UpdateHumanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required,max=40"
	OldName string `protobuf:"bytes,1,opt,name=oldName,proto3" json:"oldName,omitempty" validate:"required,max=40"`
	// @inject_tag: validate:"required"
	Human *Human `protobuf:"bytes,2,opt,name=human,proto3" json:"human,omitempty" validate:"required"`
}

func (x *UpdateHumanRequest) Reset() {
	*x = UpdateHumanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_human_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHumanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHumanRequest) ProtoMessage() {}

func (x *UpdateHumanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_human_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHumanRequest.ProtoReflect.Descriptor instead.
func (*UpdateHumanRequest) Descriptor() ([]byte, []int) {
	return file_human_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateHumanRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *UpdateHumanRequest) GetHuman() *Human {
	if x != nil {
		return x.Human
	}
	return nil
}

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required"
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty" validate:"required"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_human_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_human_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_human_service_proto_rawDescGZIP(), []int{2}
}

func (x *Name) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required"
	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validate:"required"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_human_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_human_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_human_service_proto_rawDescGZIP(), []int{3}
}

func (x *SignInRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Tokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Access  string `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Refresh string `protobuf:"bytes,2,opt,name=refresh,proto3" json:"refresh,omitempty"`
}

func (x *Tokens) Reset() {
	*x = Tokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_human_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tokens) ProtoMessage() {}

func (x *Tokens) ProtoReflect() protoreflect.Message {
	mi := &file_human_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tokens.ProtoReflect.Descriptor instead.
func (*Tokens) Descriptor() ([]byte, []int) {
	return file_human_service_proto_rawDescGZIP(), []int{4}
}

func (x *Tokens) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

func (x *Tokens) GetRefresh() string {
	if x != nil {
		return x.Refresh
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required"
	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	Role string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	ID string `protobuf:"bytes,4,opt,name=ID,proto3" json:"ID,omitempty" validate:"required"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_human_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_human_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_human_service_proto_rawDescGZIP(), []int{5}
}

func (x *User) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required"
	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validate:"required"`
	// @inject_tag: validate:"required,eqfield=Password"
	ConfirmPwd string `protobuf:"bytes,3,opt,name=confirmPwd,proto3" json:"confirmPwd,omitempty" validate:"required,eqfield=Password"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_human_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_human_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_human_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateUserRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetConfirmPwd() string {
	if x != nil {
		return x.ConfirmPwd
	}
	return ""
}

type Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required"
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty" validate:"required"`
}

func (x *Login) Reset() {
	*x = Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_human_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login) ProtoMessage() {}

func (x *Login) ProtoReflect() protoreflect.Message {
	mi := &file_human_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login.ProtoReflect.Descriptor instead.
func (*Login) Descriptor() ([]byte, []int) {
	return file_human_service_proto_rawDescGZIP(), []int{7}
}

func (x *Login) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required"
	OldLogin string `protobuf:"bytes,1,opt,name=oldLogin,proto3" json:"oldLogin,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	NewLogin string `protobuf:"bytes,2,opt,name=newLogin,proto3" json:"newLogin,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	NewPassword string `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty" validate:"required"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_human_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_human_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_human_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateUserRequest) GetOldLogin() string {
	if x != nil {
		return x.OldLogin
	}
	return ""
}

func (x *UpdateUserRequest) GetNewLogin() string {
	if x != nil {
		return x.NewLogin
	}
	return ""
}

func (x *UpdateUserRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type FilePortion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Start bool   `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
}

func (x *FilePortion) Reset() {
	*x = FilePortion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_human_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilePortion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilePortion) ProtoMessage() {}

func (x *FilePortion) ProtoReflect() protoreflect.Message {
	mi := &file_human_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilePortion.ProtoReflect.Descriptor instead.
func (*FilePortion) Descriptor() ([]byte, []int) {
	return file_human_service_proto_rawDescGZIP(), []int{9}
}

func (x *FilePortion) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *FilePortion) GetStart() bool {
	if x != nil {
		return x.Start
	}
	return false
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_human_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_human_service_proto_msgTypes[10]
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
	return file_human_service_proto_rawDescGZIP(), []int{10}
}

var File_human_service_proto protoreflect.FileDescriptor

var file_human_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22,
	0x51, 0x0a, 0x05, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6d, 0x61,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x55, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x75, 0x6d, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x6c, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x48, 0x75, 0x6d,
	0x61, 0x6e, 0x52, 0x05, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x22, 0x1c, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3a, 0x0a, 0x06, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x22, 0x5c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x22, 0x65, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x77, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x77, 0x64, 0x22, 0x1d, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6d, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x65, 0x77, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x65, 0x77, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65,
	0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xcc, 0x05,
	0x0a, 0x0c, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x12, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x1a, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x75, 0x6d, 0x61, 0x6e,
	0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x2d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x12, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x12,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06,
	0x4c, 0x6f, 0x67, 0x4f, 0x75, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0c, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x2f, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_human_service_proto_rawDescOnce sync.Once
	file_human_service_proto_rawDescData = file_human_service_proto_rawDesc
)

func file_human_service_proto_rawDescGZIP() []byte {
	file_human_service_proto_rawDescOnce.Do(func() {
		file_human_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_human_service_proto_rawDescData)
	})
	return file_human_service_proto_rawDescData
}

var file_human_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_human_service_proto_goTypes = []interface{}{
	(*Human)(nil),              // 0: protobuf.Human
	(*UpdateHumanRequest)(nil), // 1: protobuf.UpdateHumanRequest
	(*Name)(nil),               // 2: protobuf.Name
	(*SignInRequest)(nil),      // 3: protobuf.SignInRequest
	(*Tokens)(nil),             // 4: protobuf.Tokens
	(*User)(nil),               // 5: protobuf.User
	(*CreateUserRequest)(nil),  // 6: protobuf.CreateUserRequest
	(*Login)(nil),              // 7: protobuf.Login
	(*UpdateUserRequest)(nil),  // 8: protobuf.UpdateUserRequest
	(*FilePortion)(nil),        // 9: protobuf.FilePortion
	(*Empty)(nil),              // 10: protobuf.Empty
}
var file_human_service_proto_depIdxs = []int32{
	0,  // 0: protobuf.UpdateHumanRequest.human:type_name -> protobuf.Human
	0,  // 1: protobuf.HumanService.CreateHuman:input_type -> protobuf.Human
	1,  // 2: protobuf.HumanService.UpdateHuman:input_type -> protobuf.UpdateHumanRequest
	2,  // 3: protobuf.HumanService.GetHuman:input_type -> protobuf.Name
	2,  // 4: protobuf.HumanService.DeleteHuman:input_type -> protobuf.Name
	3,  // 5: protobuf.HumanService.Authenticate:input_type -> protobuf.SignInRequest
	6,  // 6: protobuf.HumanService.CreateUser:input_type -> protobuf.CreateUserRequest
	7,  // 7: protobuf.HumanService.GetUser:input_type -> protobuf.Login
	8,  // 8: protobuf.HumanService.UpdateUser:input_type -> protobuf.UpdateUserRequest
	7,  // 9: protobuf.HumanService.DeleteUser:input_type -> protobuf.Login
	4,  // 10: protobuf.HumanService.Refresh:input_type -> protobuf.Tokens
	10, // 11: protobuf.HumanService.LogOut:input_type -> protobuf.Empty
	2,  // 12: protobuf.HumanService.DownloadFile:input_type -> protobuf.Name
	2,  // 13: protobuf.HumanService.UploadFile:input_type -> protobuf.Name
	10, // 14: protobuf.HumanService.CreateHuman:output_type -> protobuf.Empty
	10, // 15: protobuf.HumanService.UpdateHuman:output_type -> protobuf.Empty
	0,  // 16: protobuf.HumanService.GetHuman:output_type -> protobuf.Human
	10, // 17: protobuf.HumanService.DeleteHuman:output_type -> protobuf.Empty
	4,  // 18: protobuf.HumanService.Authenticate:output_type -> protobuf.Tokens
	10, // 19: protobuf.HumanService.CreateUser:output_type -> protobuf.Empty
	5,  // 20: protobuf.HumanService.GetUser:output_type -> protobuf.User
	10, // 21: protobuf.HumanService.UpdateUser:output_type -> protobuf.Empty
	10, // 22: protobuf.HumanService.DeleteUser:output_type -> protobuf.Empty
	4,  // 23: protobuf.HumanService.Refresh:output_type -> protobuf.Tokens
	10, // 24: protobuf.HumanService.LogOut:output_type -> protobuf.Empty
	9,  // 25: protobuf.HumanService.DownloadFile:output_type -> protobuf.FilePortion
	10, // 26: protobuf.HumanService.UploadFile:output_type -> protobuf.Empty
	14, // [14:27] is the sub-list for method output_type
	1,  // [1:14] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_human_service_proto_init() }
func file_human_service_proto_init() {
	if File_human_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_human_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Human); i {
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
		file_human_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHumanRequest); i {
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
		file_human_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_human_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
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
		file_human_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tokens); i {
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
		file_human_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_human_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_human_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Login); i {
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
		file_human_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_human_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilePortion); i {
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
		file_human_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_human_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_human_service_proto_goTypes,
		DependencyIndexes: file_human_service_proto_depIdxs,
		MessageInfos:      file_human_service_proto_msgTypes,
	}.Build()
	File_human_service_proto = out.File
	file_human_service_proto_rawDesc = nil
	file_human_service_proto_goTypes = nil
	file_human_service_proto_depIdxs = nil
}
