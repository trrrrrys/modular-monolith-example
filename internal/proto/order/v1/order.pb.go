// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: order/v1/order.proto

package order

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

type OrderItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *OrderItemRequest) Reset() {
	*x = OrderItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemRequest) ProtoMessage() {}

func (x *OrderItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemRequest.ProtoReflect.Descriptor instead.
func (*OrderItemRequest) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItemRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *OrderItemRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type OrderItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ItemId  string `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	UserId  string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *OrderItemResponse) Reset() {
	*x = OrderItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemResponse) ProtoMessage() {}

func (x *OrderItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemResponse.ProtoReflect.Descriptor instead.
func (*OrderItemResponse) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderItemResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderItemResponse) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *OrderItemResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_order_v1_order_proto protoreflect.FileDescriptor

var file_order_v1_order_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x22, 0x44, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x54, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4c,
	0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x72,
	0x72, 0x72, 0x72, 0x79, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x61, 0x72, 0x2d, 0x6d, 0x6f,
	0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_v1_order_proto_rawDescOnce sync.Once
	file_order_v1_order_proto_rawDescData = file_order_v1_order_proto_rawDesc
)

func file_order_v1_order_proto_rawDescGZIP() []byte {
	file_order_v1_order_proto_rawDescOnce.Do(func() {
		file_order_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_v1_order_proto_rawDescData)
	})
	return file_order_v1_order_proto_rawDescData
}

var file_order_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_order_v1_order_proto_goTypes = []interface{}{
	(*OrderItemRequest)(nil),  // 0: order.v1.OrderItemRequest
	(*OrderItemResponse)(nil), // 1: order.v1.OrderItemResponse
}
var file_order_v1_order_proto_depIdxs = []int32{
	0, // 0: order.v1.OrderService.OrderItem:input_type -> order.v1.OrderItemRequest
	1, // 1: order.v1.OrderService.OrderItem:output_type -> order.v1.OrderItemResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_v1_order_proto_init() }
func file_order_v1_order_proto_init() {
	if File_order_v1_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_v1_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemRequest); i {
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
		file_order_v1_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemResponse); i {
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
			RawDescriptor: file_order_v1_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_v1_order_proto_goTypes,
		DependencyIndexes: file_order_v1_order_proto_depIdxs,
		MessageInfos:      file_order_v1_order_proto_msgTypes,
	}.Build()
	File_order_v1_order_proto = out.File
	file_order_v1_order_proto_rawDesc = nil
	file_order_v1_order_proto_goTypes = nil
	file_order_v1_order_proto_depIdxs = nil
}
