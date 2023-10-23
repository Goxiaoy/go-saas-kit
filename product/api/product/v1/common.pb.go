// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: product/api/product/v1/common.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_product_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_product_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_product_api_product_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Media) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Media) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Media) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Badge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code  string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Label string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Badge) Reset() {
	*x = Badge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_product_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Badge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Badge) ProtoMessage() {}

func (x *Badge) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_product_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Badge.ProtoReflect.Descriptor instead.
func (*Badge) Descriptor() ([]byte, []int) {
	return file_product_api_product_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *Badge) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Badge) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Badge) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type Keyword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text  string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Refer string `protobuf:"bytes,3,opt,name=refer,proto3" json:"refer,omitempty"`
}

func (x *Keyword) Reset() {
	*x = Keyword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_product_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keyword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keyword) ProtoMessage() {}

func (x *Keyword) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_product_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keyword.ProtoReflect.Descriptor instead.
func (*Keyword) Descriptor() ([]byte, []int) {
	return file_product_api_product_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *Keyword) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Keyword) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Keyword) GetRefer() string {
	if x != nil {
		return x.Refer
	}
	return ""
}

type CampaignRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Rule  string           `protobuf:"bytes,2,opt,name=rule,proto3" json:"rule,omitempty"`
	Extra *structpb.Struct `protobuf:"bytes,3,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *CampaignRule) Reset() {
	*x = CampaignRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_product_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignRule) ProtoMessage() {}

func (x *CampaignRule) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_product_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignRule.ProtoReflect.Descriptor instead.
func (*CampaignRule) Descriptor() ([]byte, []int) {
	return file_product_api_product_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *CampaignRule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CampaignRule) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *CampaignRule) GetExtra() *structpb.Struct {
	if x != nil {
		return x.Extra
	}
	return nil
}

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InStock      bool   `protobuf:"varint,2,opt,name=in_stock,json=inStock,proto3" json:"in_stock,omitempty"`
	Level        string `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	Amount       int32  `protobuf:"varint,10,opt,name=amount,proto3" json:"amount,omitempty"`
	DeliveryCode string `protobuf:"bytes,11,opt,name=delivery_code,json=deliveryCode,proto3" json:"delivery_code,omitempty"`
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_product_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_product_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_product_api_product_v1_common_proto_rawDescGZIP(), []int{4}
}

func (x *Stock) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Stock) GetInStock() bool {
	if x != nil {
		return x.InStock
	}
	return false
}

func (x *Stock) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Stock) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Stock) GetDeliveryCode() string {
	if x != nil {
		return x.DeliveryCode
	}
	return ""
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyCode      string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	DefaultAmount     int32  `protobuf:"varint,2,opt,name=default_amount,json=defaultAmount,proto3" json:"default_amount,omitempty"`
	DiscountedAmount  int32  `protobuf:"varint,3,opt,name=discounted_amount,json=discountedAmount,proto3" json:"discounted_amount,omitempty"`
	DiscountText      string `protobuf:"bytes,5,opt,name=discount_text,json=discountText,proto3" json:"discount_text,omitempty"`
	DenyMoreDiscounts bool   `protobuf:"varint,6,opt,name=deny_more_discounts,json=denyMoreDiscounts,proto3" json:"deny_more_discounts,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_product_v1_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_product_v1_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_product_api_product_v1_common_proto_rawDescGZIP(), []int{5}
}

func (x *Price) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *Price) GetDefaultAmount() int32 {
	if x != nil {
		return x.DefaultAmount
	}
	return 0
}

func (x *Price) GetDiscountedAmount() int32 {
	if x != nil {
		return x.DiscountedAmount
	}
	return 0
}

func (x *Price) GetDiscountText() string {
	if x != nil {
		return x.DiscountText
	}
	return ""
}

func (x *Price) GetDenyMoreDiscounts() bool {
	if x != nil {
		return x.DenyMoreDiscounts
	}
	return false
}

var File_product_api_product_v1_common_proto protoreflect.FileDescriptor

var file_product_api_product_v1_common_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x41, 0x0a, 0x05, 0x42, 0x61,
	0x64, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x43, 0x0a,
	0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x22, 0x61, 0x0a, 0x0c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x9a, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x29, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xfa, 0x42, 0x10, 0x72, 0x0e,
	0x52, 0x03, 0x6f, 0x75, 0x74, 0x52, 0x02, 0x69, 0x6e, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0xd5, 0x01, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65,
	0x6e, 0x79, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x65, 0x6e, 0x79, 0x4d, 0x6f, 0x72,
	0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_api_product_v1_common_proto_rawDescOnce sync.Once
	file_product_api_product_v1_common_proto_rawDescData = file_product_api_product_v1_common_proto_rawDesc
)

func file_product_api_product_v1_common_proto_rawDescGZIP() []byte {
	file_product_api_product_v1_common_proto_rawDescOnce.Do(func() {
		file_product_api_product_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_api_product_v1_common_proto_rawDescData)
	})
	return file_product_api_product_v1_common_proto_rawDescData
}

var file_product_api_product_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_api_product_v1_common_proto_goTypes = []interface{}{
	(*Media)(nil),           // 0: product.api.product.v1.Media
	(*Badge)(nil),           // 1: product.api.product.v1.Badge
	(*Keyword)(nil),         // 2: product.api.product.v1.Keyword
	(*CampaignRule)(nil),    // 3: product.api.product.v1.CampaignRule
	(*Stock)(nil),           // 4: product.api.product.v1.Stock
	(*Price)(nil),           // 5: product.api.product.v1.Price
	(*structpb.Struct)(nil), // 6: google.protobuf.Struct
}
var file_product_api_product_v1_common_proto_depIdxs = []int32{
	6, // 0: product.api.product.v1.CampaignRule.extra:type_name -> google.protobuf.Struct
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_api_product_v1_common_proto_init() }
func file_product_api_product_v1_common_proto_init() {
	if File_product_api_product_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_api_product_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Media); i {
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
		file_product_api_product_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Badge); i {
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
		file_product_api_product_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keyword); i {
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
		file_product_api_product_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignRule); i {
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
		file_product_api_product_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stock); i {
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
		file_product_api_product_v1_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
			RawDescriptor: file_product_api_product_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_api_product_v1_common_proto_goTypes,
		DependencyIndexes: file_product_api_product_v1_common_proto_depIdxs,
		MessageInfos:      file_product_api_product_v1_common_proto_msgTypes,
	}.Build()
	File_product_api_product_v1_common_proto = out.File
	file_product_api_product_v1_common_proto_rawDesc = nil
	file_product_api_product_v1_common_proto_goTypes = nil
	file_product_api_product_v1_common_proto_depIdxs = nil
}
