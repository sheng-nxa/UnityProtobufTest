// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: testMsg.proto

package wsserver

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

// Command id
type Command int32

const (
	Command_Login Command = 0 // Login
	Command_Ping  Command = 1 // ping
	Command_Pong  Command = 2 // pong
)

// Enum value maps for Command.
var (
	Command_name = map[int32]string{
		0: "Login",
		1: "Ping",
		2: "Pong",
	}
	Command_value = map[string]int32{
		"Login": 0,
		"Ping":  1,
		"Pong":  2,
	}
)

func (x Command) Enum() *Command {
	p := new(Command)
	*p = x
	return p
}

func (x Command) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command) Descriptor() protoreflect.EnumDescriptor {
	return file_testMsg_proto_enumTypes[0].Descriptor()
}

func (Command) Type() protoreflect.EnumType {
	return &file_testMsg_proto_enumTypes[0]
}

func (x Command) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command.Descriptor instead.
func (Command) EnumDescriptor() ([]byte, []int) {
	return file_testMsg_proto_rawDescGZIP(), []int{0}
}

// ping
type CommandPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *CommandPing) Reset() {
	*x = CommandPing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testMsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandPing) ProtoMessage() {}

func (x *CommandPing) ProtoReflect() protoreflect.Message {
	mi := &file_testMsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandPing.ProtoReflect.Descriptor instead.
func (*CommandPing) Descriptor() ([]byte, []int) {
	return file_testMsg_proto_rawDescGZIP(), []int{0}
}

func (x *CommandPing) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

// pong
type CommandPong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *CommandPong) Reset() {
	*x = CommandPong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testMsg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandPong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandPong) ProtoMessage() {}

func (x *CommandPong) ProtoReflect() protoreflect.Message {
	mi := &file_testMsg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandPong.ProtoReflect.Descriptor instead.
func (*CommandPong) Descriptor() ([]byte, []int) {
	return file_testMsg_proto_rawDescGZIP(), []int{1}
}

func (x *CommandPong) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

// login
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testMsg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testMsg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_testMsg_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

// user
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string   `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Orders   []*Order `protobuf:"bytes,2,rep,name=Orders,proto3" json:"Orders,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testMsg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_testMsg_proto_msgTypes[3]
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
	return file_testMsg_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *User) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   int64  `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	OrderName string `protobuf:"bytes,2,opt,name=OrderName,proto3" json:"OrderName,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testMsg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_testMsg_proto_msgTypes[4]
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
	return file_testMsg_proto_rawDescGZIP(), []int{4}
}

func (x *Order) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Order) GetOrderName() string {
	if x != nil {
		return x.OrderName
	}
	return ""
}

var File_testMsg_proto protoreflect.FileDescriptor

var file_testMsg_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2a, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x22, 0x3f, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x2a, 0x28, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x09, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x10, 0x02, 0x42, 0x1b, 0x5a, 0x19, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3b,
	0x77, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testMsg_proto_rawDescOnce sync.Once
	file_testMsg_proto_rawDescData = file_testMsg_proto_rawDesc
)

func file_testMsg_proto_rawDescGZIP() []byte {
	file_testMsg_proto_rawDescOnce.Do(func() {
		file_testMsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_testMsg_proto_rawDescData)
	})
	return file_testMsg_proto_rawDescData
}

var file_testMsg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_testMsg_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_testMsg_proto_goTypes = []interface{}{
	(Command)(0),         // 0: main.Command
	(*CommandPing)(nil),  // 1: main.CommandPing
	(*CommandPong)(nil),  // 2: main.CommandPong
	(*LoginRequest)(nil), // 3: main.LoginRequest
	(*User)(nil),         // 4: main.User
	(*Order)(nil),        // 5: main.Order
}
var file_testMsg_proto_depIdxs = []int32{
	5, // 0: main.User.Orders:type_name -> main.Order
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_testMsg_proto_init() }
func file_testMsg_proto_init() {
	if File_testMsg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testMsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandPing); i {
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
		file_testMsg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandPong); i {
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
		file_testMsg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_testMsg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_testMsg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testMsg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testMsg_proto_goTypes,
		DependencyIndexes: file_testMsg_proto_depIdxs,
		EnumInfos:         file_testMsg_proto_enumTypes,
		MessageInfos:      file_testMsg_proto_msgTypes,
	}.Build()
	File_testMsg_proto = out.File
	file_testMsg_proto_rawDesc = nil
	file_testMsg_proto_goTypes = nil
	file_testMsg_proto_depIdxs = nil
}
