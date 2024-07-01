// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: basket.proto

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

type B_Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *B_Void) Reset() {
	*x = B_Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *B_Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*B_Void) ProtoMessage() {}

func (x *B_Void) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use B_Void.ProtoReflect.Descriptor instead.
func (*B_Void) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{0}
}

type CreatedBasket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreatedBasket) Reset() {
	*x = CreatedBasket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedBasket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedBasket) ProtoMessage() {}

func (x *CreatedBasket) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedBasket.ProtoReflect.Descriptor instead.
func (*CreatedBasket) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{1}
}

func (x *CreatedBasket) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ByIdBasket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ByIdBasket) Reset() {
	*x = ByIdBasket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByIdBasket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByIdBasket) ProtoMessage() {}

func (x *ByIdBasket) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByIdBasket.ProtoReflect.Descriptor instead.
func (*ByIdBasket) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{2}
}

func (x *ByIdBasket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Basket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID *User  `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *Basket) Reset() {
	*x = Basket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Basket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Basket) ProtoMessage() {}

func (x *Basket) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Basket.ProtoReflect.Descriptor instead.
func (*Basket) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{3}
}

func (x *Basket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Basket) GetUserID() *User {
	if x != nil {
		return x.UserID
	}
	return nil
}

type FilterBasket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FilterBasket) Reset() {
	*x = FilterBasket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterBasket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterBasket) ProtoMessage() {}

func (x *FilterBasket) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterBasket.ProtoReflect.Descriptor instead.
func (*FilterBasket) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{4}
}

func (x *FilterBasket) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetAllBasketRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Baskets []*Basket `protobuf:"bytes,1,rep,name=baskets,proto3" json:"baskets,omitempty"`
}

func (x *GetAllBasketRes) Reset() {
	*x = GetAllBasketRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBasketRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBasketRes) ProtoMessage() {}

func (x *GetAllBasketRes) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBasketRes.ProtoReflect.Descriptor instead.
func (*GetAllBasketRes) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllBasketRes) GetBaskets() []*Basket {
	if x != nil {
		return x.Baskets
	}
	return nil
}

var File_basket_proto protoreflect.FileDescriptor

var file_basket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x08, 0x0a, 0x06, 0x42, 0x5f, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x28, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x42, 0x79, 0x49, 0x64, 0x42,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x06, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x25, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x27, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x3c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x42, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x52, 0x07, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x32, 0xaa, 0x02,
	0x0a, 0x0d, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x42, 0x5f, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x42, 0x79, 0x49, 0x64, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x1a, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x1a,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x42, 0x5f, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x42, 0x79, 0x49,
	0x64, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x42, 0x5f, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basket_proto_rawDescOnce sync.Once
	file_basket_proto_rawDescData = file_basket_proto_rawDesc
)

func file_basket_proto_rawDescGZIP() []byte {
	file_basket_proto_rawDescOnce.Do(func() {
		file_basket_proto_rawDescData = protoimpl.X.CompressGZIP(file_basket_proto_rawDescData)
	})
	return file_basket_proto_rawDescData
}

var file_basket_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_basket_proto_goTypes = []interface{}{
	(*B_Void)(nil),          // 0: product.B_Void
	(*CreatedBasket)(nil),   // 1: product.CreatedBasket
	(*ByIdBasket)(nil),      // 2: product.ByIdBasket
	(*Basket)(nil),          // 3: product.Basket
	(*FilterBasket)(nil),    // 4: product.FilterBasket
	(*GetAllBasketRes)(nil), // 5: product.GetAllBasketRes
	(*User)(nil),            // 6: product.User
}
var file_basket_proto_depIdxs = []int32{
	6, // 0: product.Basket.userID:type_name -> product.User
	3, // 1: product.GetAllBasketRes.baskets:type_name -> product.Basket
	1, // 2: product.BasketService.CreateBasket:input_type -> product.CreatedBasket
	2, // 3: product.BasketService.GetByID:input_type -> product.ByIdBasket
	4, // 4: product.BasketService.GetAll:input_type -> product.FilterBasket
	2, // 5: product.BasketService.UpdateBasket:input_type -> product.ByIdBasket
	2, // 6: product.BasketService.DeleteBasket:input_type -> product.ByIdBasket
	0, // 7: product.BasketService.CreateBasket:output_type -> product.B_Void
	3, // 8: product.BasketService.GetByID:output_type -> product.Basket
	5, // 9: product.BasketService.GetAll:output_type -> product.GetAllBasketRes
	0, // 10: product.BasketService.UpdateBasket:output_type -> product.B_Void
	0, // 11: product.BasketService.DeleteBasket:output_type -> product.B_Void
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_basket_proto_init() }
func file_basket_proto_init() {
	if File_basket_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_basket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*B_Void); i {
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
		file_basket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedBasket); i {
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
		file_basket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByIdBasket); i {
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
		file_basket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Basket); i {
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
		file_basket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterBasket); i {
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
		file_basket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBasketRes); i {
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
			RawDescriptor: file_basket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basket_proto_goTypes,
		DependencyIndexes: file_basket_proto_depIdxs,
		MessageInfos:      file_basket_proto_msgTypes,
	}.Build()
	File_basket_proto = out.File
	file_basket_proto_rawDesc = nil
	file_basket_proto_goTypes = nil
	file_basket_proto_depIdxs = nil
}