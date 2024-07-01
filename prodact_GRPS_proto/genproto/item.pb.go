// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: item.proto

package genproto

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

type I_Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *I_Void) Reset() {
	*x = I_Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *I_Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*I_Void) ProtoMessage() {}

func (x *I_Void) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use I_Void.ProtoReflect.Descriptor instead.
func (*I_Void) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{0}
}

type CreateItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasketID  string `protobuf:"bytes,1,opt,name=basketID,proto3" json:"basketID,omitempty"`
	ProductID string `protobuf:"bytes,2,opt,name=productID,proto3" json:"productID,omitempty"`
	Quantity  int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CreateItem) Reset() {
	*x = CreateItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItem) ProtoMessage() {}

func (x *CreateItem) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItem.ProtoReflect.Descriptor instead.
func (*CreateItem) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{1}
}

func (x *CreateItem) GetBasketID() string {
	if x != nil {
		return x.BasketID
	}
	return ""
}

func (x *CreateItem) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *CreateItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ByIdItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ByIdItem) Reset() {
	*x = ByIdItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByIdItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByIdItem) ProtoMessage() {}

func (x *ByIdItem) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByIdItem.ProtoReflect.Descriptor instead.
func (*ByIdItem) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{2}
}

func (x *ByIdItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FilterItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasketID string `protobuf:"bytes,1,opt,name=basketID,proto3" json:"basketID,omitempty"`
}

func (x *FilterItem) Reset() {
	*x = FilterItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterItem) ProtoMessage() {}

func (x *FilterItem) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterItem.ProtoReflect.Descriptor instead.
func (*FilterItem) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{3}
}

func (x *FilterItem) GetBasketID() string {
	if x != nil {
		return x.BasketID
	}
	return ""
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Basket   *Basket  `protobuf:"bytes,1,opt,name=basket,proto3" json:"basket,omitempty"`
	Product  *Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Quantity int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{4}
}

func (x *Item) GetBasket() *Basket {
	if x != nil {
		return x.Basket
	}
	return nil
}

func (x *Item) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type GetAllItemRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetAllItemRes) Reset() {
	*x = GetAllItemRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllItemRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllItemRes) ProtoMessage() {}

func (x *GetAllItemRes) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllItemRes.ProtoReflect.Descriptor instead.
func (*GetAllItemRes) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllItemRes) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpdateItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BasketID  string `protobuf:"bytes,2,opt,name=basketID,proto3" json:"basketID,omitempty"`
	ProductID string `protobuf:"bytes,3,opt,name=productID,proto3" json:"productID,omitempty"`
	Quantity  int32  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *UpdateItem) Reset() {
	*x = UpdateItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItem) ProtoMessage() {}

func (x *UpdateItem) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItem.ProtoReflect.Descriptor instead.
func (*UpdateItem) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateItem) GetBasketID() string {
	if x != nil {
		return x.BasketID
	}
	return ""
}

func (x *UpdateItem) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *UpdateItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_item_proto protoreflect.FileDescriptor

var file_item_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x0c, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x08, 0x0a, 0x06, 0x49, 0x5f, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x1a, 0x0a, 0x08, 0x42, 0x79, 0x49, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0a,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x22, 0x77, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x27,
	0x0a, 0x06, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52,
	0x06, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x34, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x72, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0xff, 0x01, 0x0a, 0x0b, 0x49, 0x74,
	0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x49, 0x5f, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x42,
	0x79, 0x49, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x35, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x2e, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x49, 0x5f, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2c, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x49, 0x5f, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x0b, 0x5a, 0x09, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_item_proto_rawDescOnce sync.Once
	file_item_proto_rawDescData = file_item_proto_rawDesc
)

func file_item_proto_rawDescGZIP() []byte {
	file_item_proto_rawDescOnce.Do(func() {
		file_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_item_proto_rawDescData)
	})
	return file_item_proto_rawDescData
}

var file_item_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_item_proto_goTypes = []interface{}{
	(*I_Void)(nil),        // 0: product.I_Void
	(*CreateItem)(nil),    // 1: product.CreateItem
	(*ByIdItem)(nil),      // 2: product.ByIdItem
	(*FilterItem)(nil),    // 3: product.FilterItem
	(*Item)(nil),          // 4: product.Item
	(*GetAllItemRes)(nil), // 5: product.GetAllItemRes
	(*UpdateItem)(nil),    // 6: product.UpdateItem
	(*Basket)(nil),        // 7: product.Basket
	(*Product)(nil),       // 8: product.Product
}
var file_item_proto_depIdxs = []int32{
	7, // 0: product.Item.basket:type_name -> product.Basket
	8, // 1: product.Item.product:type_name -> product.Product
	4, // 2: product.GetAllItemRes.items:type_name -> product.Item
	1, // 3: product.ItemService.Create:input_type -> product.CreateItem
	2, // 4: product.ItemService.GetById:input_type -> product.ByIdItem
	3, // 5: product.ItemService.GetAll:input_type -> product.FilterItem
	6, // 6: product.ItemService.Update:input_type -> product.UpdateItem
	2, // 7: product.ItemService.Delete:input_type -> product.ByIdItem
	0, // 8: product.ItemService.Create:output_type -> product.I_Void
	4, // 9: product.ItemService.GetById:output_type -> product.Item
	5, // 10: product.ItemService.GetAll:output_type -> product.GetAllItemRes
	0, // 11: product.ItemService.Update:output_type -> product.I_Void
	0, // 12: product.ItemService.Delete:output_type -> product.I_Void
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_item_proto_init() }
func file_item_proto_init() {
	if File_item_proto != nil {
		return
	}
	file_basket_proto_init()
	file_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*I_Void); i {
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
		file_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateItem); i {
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
		file_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByIdItem); i {
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
		file_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterItem); i {
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
		file_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllItemRes); i {
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
		file_item_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItem); i {
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
			RawDescriptor: file_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_item_proto_goTypes,
		DependencyIndexes: file_item_proto_depIdxs,
		MessageInfos:      file_item_proto_msgTypes,
	}.Build()
	File_item_proto = out.File
	file_item_proto_rawDesc = nil
	file_item_proto_goTypes = nil
	file_item_proto_depIdxs = nil
}
