// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: internal/adapters/grpc/order/protofiles/order.proto

package orderpb

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

type Blank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Blank) Reset() {
	*x = Blank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blank) ProtoMessage() {}

func (x *Blank) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blank.ProtoReflect.Descriptor instead.
func (*Blank) Descriptor() ([]byte, []int) {
	return file_internal_adapters_grpc_order_protofiles_order_proto_rawDescGZIP(), []int{0}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string  `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Price      float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax        float32 `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice float32 `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
	CreatedAt  string  `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_internal_adapters_grpc_order_protofiles_order_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetTax() float32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *Order) GetFinalPrice() float32 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

func (x *Order) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax   float32 `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_internal_adapters_grpc_order_protofiles_order_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrderRequest) GetTax() float32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

type OrderList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *OrderList) Reset() {
	*x = OrderList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderList) ProtoMessage() {}

func (x *OrderList) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderList.ProtoReflect.Descriptor instead.
func (*OrderList) Descriptor() ([]byte, []int) {
	return file_internal_adapters_grpc_order_protofiles_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderList) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_internal_adapters_grpc_order_protofiles_order_proto protoreflect.FileDescriptor

var file_internal_adapters_grpc_order_protofiles_order_proto_rawDesc = []byte{
	0x0a, 0x33, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x22, 0x07,
	0x0a, 0x05, 0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x22, 0x8a, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x74, 0x61, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x3c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x74,
	0x61, 0x78, 0x22, 0x33, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x32, 0x80, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x62,
	0x6c, 0x61, 0x6e, 0x6b, 0x1a, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_adapters_grpc_order_protofiles_order_proto_rawDescOnce sync.Once
	file_internal_adapters_grpc_order_protofiles_order_proto_rawDescData = file_internal_adapters_grpc_order_protofiles_order_proto_rawDesc
)

func file_internal_adapters_grpc_order_protofiles_order_proto_rawDescGZIP() []byte {
	file_internal_adapters_grpc_order_protofiles_order_proto_rawDescOnce.Do(func() {
		file_internal_adapters_grpc_order_protofiles_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_adapters_grpc_order_protofiles_order_proto_rawDescData)
	})
	return file_internal_adapters_grpc_order_protofiles_order_proto_rawDescData
}

var file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_adapters_grpc_order_protofiles_order_proto_goTypes = []interface{}{
	(*Blank)(nil),              // 0: orderpb.blank
	(*Order)(nil),              // 1: orderpb.Order
	(*CreateOrderRequest)(nil), // 2: orderpb.CreateOrderRequest
	(*OrderList)(nil),          // 3: orderpb.OrderList
}
var file_internal_adapters_grpc_order_protofiles_order_proto_depIdxs = []int32{
	1, // 0: orderpb.OrderList.orders:type_name -> orderpb.Order
	2, // 1: orderpb.OrderService.CreateOrder:input_type -> orderpb.CreateOrderRequest
	0, // 2: orderpb.OrderService.ListOrders:input_type -> orderpb.blank
	1, // 3: orderpb.OrderService.CreateOrder:output_type -> orderpb.Order
	3, // 4: orderpb.OrderService.ListOrders:output_type -> orderpb.OrderList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_adapters_grpc_order_protofiles_order_proto_init() }
func file_internal_adapters_grpc_order_protofiles_order_proto_init() {
	if File_internal_adapters_grpc_order_protofiles_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blank); i {
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
		file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderList); i {
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
			RawDescriptor: file_internal_adapters_grpc_order_protofiles_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_adapters_grpc_order_protofiles_order_proto_goTypes,
		DependencyIndexes: file_internal_adapters_grpc_order_protofiles_order_proto_depIdxs,
		MessageInfos:      file_internal_adapters_grpc_order_protofiles_order_proto_msgTypes,
	}.Build()
	File_internal_adapters_grpc_order_protofiles_order_proto = out.File
	file_internal_adapters_grpc_order_protofiles_order_proto_rawDesc = nil
	file_internal_adapters_grpc_order_protofiles_order_proto_goTypes = nil
	file_internal_adapters_grpc_order_protofiles_order_proto_depIdxs = nil
}
