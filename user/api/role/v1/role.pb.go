// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: user/api/role/v1/role.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RoleFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdIn   []string `protobuf:"bytes,1,rep,name=id_in,json=idIn,proto3" json:"id_in,omitempty"`
	NameIn []string `protobuf:"bytes,2,rep,name=name_in,json=nameIn,proto3" json:"name_in,omitempty"`
}

func (x *RoleFilter) Reset() {
	*x = RoleFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_role_v1_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleFilter) ProtoMessage() {}

func (x *RoleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_role_v1_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleFilter.ProtoReflect.Descriptor instead.
func (*RoleFilter) Descriptor() ([]byte, []int) {
	return file_user_api_role_v1_role_proto_rawDescGZIP(), []int{0}
}

func (x *RoleFilter) GetIdIn() []string {
	if x != nil {
		return x.IdIn
	}
	return nil
}

func (x *RoleFilter) GetNameIn() []string {
	if x != nil {
		return x.NameIn
	}
	return nil
}

type ListRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageOffset int32                  `protobuf:"varint,1,opt,name=page_offset,json=pageOffset,proto3" json:"page_offset,omitempty"`
	PageSize   int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Search     string                 `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Sort       []string               `protobuf:"bytes,4,rep,name=sort,proto3" json:"sort,omitempty"`
	Fields     *fieldmaskpb.FieldMask `protobuf:"bytes,5,opt,name=fields,proto3" json:"fields,omitempty"`
	Filter     *RoleFilter            `protobuf:"bytes,6,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListRolesRequest) Reset() {
	*x = ListRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_role_v1_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRolesRequest) ProtoMessage() {}

func (x *ListRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_role_v1_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRolesRequest.ProtoReflect.Descriptor instead.
func (*ListRolesRequest) Descriptor() ([]byte, []int) {
	return file_user_api_role_v1_role_proto_rawDescGZIP(), []int{1}
}

func (x *ListRolesRequest) GetPageOffset() int32 {
	if x != nil {
		return x.PageOffset
	}
	return 0
}

func (x *ListRolesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRolesRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListRolesRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *ListRolesRequest) GetFields() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ListRolesRequest) GetFilter() *RoleFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalSize  int32   `protobuf:"varint,1,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	FilterSize int32   `protobuf:"varint,2,opt,name=filter_size,json=filterSize,proto3" json:"filter_size,omitempty"`
	Items      []*Role `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListRolesResponse) Reset() {
	*x = ListRolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_role_v1_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRolesResponse) ProtoMessage() {}

func (x *ListRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_role_v1_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRolesResponse.ProtoReflect.Descriptor instead.
func (*ListRolesResponse) Descriptor() ([]byte, []int) {
	return file_user_api_role_v1_role_proto_rawDescGZIP(), []int{2}
}

func (x *ListRolesResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *ListRolesResponse) GetFilterSize() int32 {
	if x != nil {
		return x.FilterSize
	}
	return 0
}

func (x *ListRolesResponse) GetItems() []*Role {
	if x != nil {
		return x.Items
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_role_v1_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_role_v1_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_user_api_role_v1_role_proto_rawDescGZIP(), []int{3}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetRoleRequest) Reset() {
	*x = GetRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_role_v1_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleRequest) ProtoMessage() {}

func (x *GetRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_role_v1_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleRequest.ProtoReflect.Descriptor instead.
func (*GetRoleRequest) Descriptor() ([]byte, []int) {
	return file_user_api_role_v1_role_proto_rawDescGZIP(), []int{4}
}

func (x *GetRoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_role_v1_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_role_v1_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_user_api_role_v1_role_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role       *UpdateRole            `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateRoleRequest) Reset() {
	*x = UpdateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_role_v1_role_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleRequest) ProtoMessage() {}

func (x *UpdateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_role_v1_role_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoleRequest) Descriptor() ([]byte, []int) {
	return file_user_api_role_v1_role_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRoleRequest) GetRole() *UpdateRole {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *UpdateRoleRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdateRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateRole) Reset() {
	*x = UpdateRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_role_v1_role_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRole) ProtoMessage() {}

func (x *UpdateRole) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_role_v1_role_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRole.ProtoReflect.Descriptor instead.
func (*UpdateRole) Descriptor() ([]byte, []int) {
	return file_user_api_role_v1_role_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRole) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRole) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRoleRequest) Reset() {
	*x = DeleteRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_role_v1_role_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRequest) ProtoMessage() {}

func (x *DeleteRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_role_v1_role_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return file_user_api_role_v1_role_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteRoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRoleResponse) Reset() {
	*x = DeleteRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_role_v1_role_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleResponse) ProtoMessage() {}

func (x *DeleteRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_role_v1_role_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleResponse.ProtoReflect.Descriptor instead.
func (*DeleteRoleResponse) Descriptor() ([]byte, []int) {
	return file_user_api_role_v1_role_proto_rawDescGZIP(), []int{9}
}

var File_user_api_role_v1_role_proto protoreflect.FileDescriptor

var file_user_api_role_v1_role_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0a, 0x52, 0x6f, 0x6c,
	0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x64, 0x5f, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x69, 0x64, 0x49, 0x6e, 0x12, 0x17, 0x0a, 0x07,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x61, 0x6d, 0x65, 0x49, 0x6e, 0x22, 0xe6, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x81,
	0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x2a, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8c, 0x01,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x30, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xba, 0x04, 0x0a, 0x0b, 0x52, 0x6f,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x5a, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x5e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x81, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x36,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x1a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x2f, 0x7b, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x5a, 0x17, 0x32,
	0x12, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x7b, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x6e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x78, 0x69, 0x61, 0x6f, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x61, 0x61, 0x73, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_api_role_v1_role_proto_rawDescOnce sync.Once
	file_user_api_role_v1_role_proto_rawDescData = file_user_api_role_v1_role_proto_rawDesc
)

func file_user_api_role_v1_role_proto_rawDescGZIP() []byte {
	file_user_api_role_v1_role_proto_rawDescOnce.Do(func() {
		file_user_api_role_v1_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_api_role_v1_role_proto_rawDescData)
	})
	return file_user_api_role_v1_role_proto_rawDescData
}

var file_user_api_role_v1_role_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_api_role_v1_role_proto_goTypes = []interface{}{
	(*RoleFilter)(nil),            // 0: user.api.role.v1.RoleFilter
	(*ListRolesRequest)(nil),      // 1: user.api.role.v1.ListRolesRequest
	(*ListRolesResponse)(nil),     // 2: user.api.role.v1.ListRolesResponse
	(*Role)(nil),                  // 3: user.api.role.v1.Role
	(*GetRoleRequest)(nil),        // 4: user.api.role.v1.GetRoleRequest
	(*CreateRoleRequest)(nil),     // 5: user.api.role.v1.CreateRoleRequest
	(*UpdateRoleRequest)(nil),     // 6: user.api.role.v1.UpdateRoleRequest
	(*UpdateRole)(nil),            // 7: user.api.role.v1.UpdateRole
	(*DeleteRoleRequest)(nil),     // 8: user.api.role.v1.DeleteRoleRequest
	(*DeleteRoleResponse)(nil),    // 9: user.api.role.v1.DeleteRoleResponse
	(*fieldmaskpb.FieldMask)(nil), // 10: google.protobuf.FieldMask
}
var file_user_api_role_v1_role_proto_depIdxs = []int32{
	10, // 0: user.api.role.v1.ListRolesRequest.fields:type_name -> google.protobuf.FieldMask
	0,  // 1: user.api.role.v1.ListRolesRequest.filter:type_name -> user.api.role.v1.RoleFilter
	3,  // 2: user.api.role.v1.ListRolesResponse.items:type_name -> user.api.role.v1.Role
	7,  // 3: user.api.role.v1.UpdateRoleRequest.role:type_name -> user.api.role.v1.UpdateRole
	10, // 4: user.api.role.v1.UpdateRoleRequest.update_mask:type_name -> google.protobuf.FieldMask
	1,  // 5: user.api.role.v1.RoleService.ListRoles:input_type -> user.api.role.v1.ListRolesRequest
	4,  // 6: user.api.role.v1.RoleService.GetRole:input_type -> user.api.role.v1.GetRoleRequest
	5,  // 7: user.api.role.v1.RoleService.CreateRole:input_type -> user.api.role.v1.CreateRoleRequest
	6,  // 8: user.api.role.v1.RoleService.UpdateRole:input_type -> user.api.role.v1.UpdateRoleRequest
	8,  // 9: user.api.role.v1.RoleService.DeleteRole:input_type -> user.api.role.v1.DeleteRoleRequest
	2,  // 10: user.api.role.v1.RoleService.ListRoles:output_type -> user.api.role.v1.ListRolesResponse
	3,  // 11: user.api.role.v1.RoleService.GetRole:output_type -> user.api.role.v1.Role
	3,  // 12: user.api.role.v1.RoleService.CreateRole:output_type -> user.api.role.v1.Role
	3,  // 13: user.api.role.v1.RoleService.UpdateRole:output_type -> user.api.role.v1.Role
	9,  // 14: user.api.role.v1.RoleService.DeleteRole:output_type -> user.api.role.v1.DeleteRoleResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_user_api_role_v1_role_proto_init() }
func file_user_api_role_v1_role_proto_init() {
	if File_user_api_role_v1_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_api_role_v1_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleFilter); i {
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
		file_user_api_role_v1_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRolesRequest); i {
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
		file_user_api_role_v1_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRolesResponse); i {
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
		file_user_api_role_v1_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_user_api_role_v1_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleRequest); i {
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
		file_user_api_role_v1_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
		file_user_api_role_v1_role_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleRequest); i {
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
		file_user_api_role_v1_role_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRole); i {
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
		file_user_api_role_v1_role_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleRequest); i {
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
		file_user_api_role_v1_role_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleResponse); i {
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
			RawDescriptor: file_user_api_role_v1_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_api_role_v1_role_proto_goTypes,
		DependencyIndexes: file_user_api_role_v1_role_proto_depIdxs,
		MessageInfos:      file_user_api_role_v1_role_proto_msgTypes,
	}.Build()
	File_user_api_role_v1_role_proto = out.File
	file_user_api_role_v1_role_proto_rawDesc = nil
	file_user_api_role_v1_role_proto_goTypes = nil
	file_user_api_role_v1_role_proto_depIdxs = nil
}
