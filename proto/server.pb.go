// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: proto/server.proto

package proto

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

type Direction int32

const (
	Direction_UP    Direction = 0
	Direction_DOWN  Direction = 1
	Direction_LEFT  Direction = 2
	Direction_RIGHT Direction = 3
	Direction_STOP  Direction = 4
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "UP",
		1: "DOWN",
		2: "LEFT",
		3: "RIGHT",
		4: "STOP",
	}
	Direction_value = map[string]int32{
		"UP":    0,
		"DOWN":  1,
		"LEFT":  2,
		"RIGHT": 3,
		"STOP":  4,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_server_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_proto_server_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{0}
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConnectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConnectRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string    `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Entities []*Entity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ConnectResponse) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

type Move struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direction Direction `protobuf:"varint,1,opt,name=direction,proto3,enum=proto.Direction" json:"direction,omitempty"`
}

func (x *Move) Reset() {
	*x = Move{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Move) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Move) ProtoMessage() {}

func (x *Move) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Move.ProtoReflect.Descriptor instead.
func (*Move) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{2}
}

func (x *Move) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_UP
}

type Coordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Coordinate) Reset() {
	*x = Coordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinate) ProtoMessage() {}

func (x *Coordinate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinate.ProtoReflect.Descriptor instead.
func (*Coordinate) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{3}
}

func (x *Coordinate) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coordinate) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Position *Coordinate `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Icon     string      `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{4}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetPosition() *Coordinate {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Player) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Entity:
	//
	//	*Entity_Player
	Entity isEntity_Entity `protobuf_oneof:"entity"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{5}
}

func (m *Entity) GetEntity() isEntity_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (x *Entity) GetPlayer() *Player {
	if x, ok := x.GetEntity().(*Entity_Player); ok {
		return x.Player
	}
	return nil
}

type isEntity_Entity interface {
	isEntity_Entity()
}

type Entity_Player struct {
	Player *Player `protobuf:"bytes,2,opt,name=player,proto3,oneof"`
}

func (*Entity_Player) isEntity_Entity() {}

type AddEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *AddEntity) Reset() {
	*x = AddEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEntity) ProtoMessage() {}

func (x *AddEntity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEntity.ProtoReflect.Descriptor instead.
func (*AddEntity) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{6}
}

func (x *AddEntity) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type UpdateEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *UpdateEntity) Reset() {
	*x = UpdateEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntity) ProtoMessage() {}

func (x *UpdateEntity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntity.ProtoReflect.Descriptor instead.
func (*UpdateEntity) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateEntity) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type RemoveEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveEntity) Reset() {
	*x = RemoveEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveEntity) ProtoMessage() {}

func (x *RemoveEntity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveEntity.ProtoReflect.Descriptor instead.
func (*RemoveEntity) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//
	//	*Request_Move
	Action isRequest_Action `protobuf_oneof:"action"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{9}
}

func (m *Request) GetAction() isRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *Request) GetMove() *Move {
	if x, ok := x.GetAction().(*Request_Move); ok {
		return x.Move
	}
	return nil
}

type isRequest_Action interface {
	isRequest_Action()
}

type Request_Move struct {
	Move *Move `protobuf:"bytes,1,opt,name=move,proto3,oneof"`
}

func (*Request_Move) isRequest_Action() {}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//
	//	*Response_AddEntity
	//	*Response_UpdateEntity
	//	*Response_RemoveEntity
	Action isResponse_Action `protobuf_oneof:"action"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{10}
}

func (m *Response) GetAction() isResponse_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *Response) GetAddEntity() *AddEntity {
	if x, ok := x.GetAction().(*Response_AddEntity); ok {
		return x.AddEntity
	}
	return nil
}

func (x *Response) GetUpdateEntity() *UpdateEntity {
	if x, ok := x.GetAction().(*Response_UpdateEntity); ok {
		return x.UpdateEntity
	}
	return nil
}

func (x *Response) GetRemoveEntity() *RemoveEntity {
	if x, ok := x.GetAction().(*Response_RemoveEntity); ok {
		return x.RemoveEntity
	}
	return nil
}

type isResponse_Action interface {
	isResponse_Action()
}

type Response_AddEntity struct {
	AddEntity *AddEntity `protobuf:"bytes,1,opt,name=addEntity,proto3,oneof"`
}

type Response_UpdateEntity struct {
	UpdateEntity *UpdateEntity `protobuf:"bytes,2,opt,name=updateEntity,proto3,oneof"`
}

type Response_RemoveEntity struct {
	RemoveEntity *RemoveEntity `protobuf:"bytes,3,opt,name=removeEntity,proto3,oneof"`
}

func (*Response_AddEntity) isResponse_Action() {}

func (*Response_UpdateEntity) isResponse_Action() {}

func (*Response_RemoveEntity) isResponse_Action() {}

var File_proto_server_proto protoreflect.FileDescriptor

var file_proto_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x52, 0x0a,
	0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x29, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x22, 0x36, 0x0a, 0x04, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x28, 0x0a, 0x0a, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x79, 0x22, 0x6f, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x27,
	0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0x32, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x35, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x1e, 0x0a, 0x0c,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f,
	0x76, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbc, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52, 0x09, 0x61, 0x64, 0x64, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48, 0x00,
	0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x39,
	0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2a, 0x3c, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x54, 0x4f, 0x50, 0x10,
	0x04, 0x32, 0x7a, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x12, 0x3a, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x19, 0x5a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_server_proto_rawDescOnce sync.Once
	file_proto_server_proto_rawDescData = file_proto_server_proto_rawDesc
)

func file_proto_server_proto_rawDescGZIP() []byte {
	file_proto_server_proto_rawDescOnce.Do(func() {
		file_proto_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_server_proto_rawDescData)
	})
	return file_proto_server_proto_rawDescData
}

var file_proto_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_server_proto_goTypes = []interface{}{
	(Direction)(0),          // 0: proto.Direction
	(*ConnectRequest)(nil),  // 1: proto.ConnectRequest
	(*ConnectResponse)(nil), // 2: proto.ConnectResponse
	(*Move)(nil),            // 3: proto.Move
	(*Coordinate)(nil),      // 4: proto.Coordinate
	(*Player)(nil),          // 5: proto.Player
	(*Entity)(nil),          // 6: proto.Entity
	(*AddEntity)(nil),       // 7: proto.AddEntity
	(*UpdateEntity)(nil),    // 8: proto.UpdateEntity
	(*RemoveEntity)(nil),    // 9: proto.RemoveEntity
	(*Request)(nil),         // 10: proto.Request
	(*Response)(nil),        // 11: proto.Response
}
var file_proto_server_proto_depIdxs = []int32{
	6,  // 0: proto.ConnectResponse.entities:type_name -> proto.Entity
	0,  // 1: proto.Move.direction:type_name -> proto.Direction
	4,  // 2: proto.Player.position:type_name -> proto.Coordinate
	5,  // 3: proto.Entity.player:type_name -> proto.Player
	6,  // 4: proto.AddEntity.entity:type_name -> proto.Entity
	6,  // 5: proto.UpdateEntity.entity:type_name -> proto.Entity
	3,  // 6: proto.Request.move:type_name -> proto.Move
	7,  // 7: proto.Response.addEntity:type_name -> proto.AddEntity
	8,  // 8: proto.Response.updateEntity:type_name -> proto.UpdateEntity
	9,  // 9: proto.Response.removeEntity:type_name -> proto.RemoveEntity
	1,  // 10: proto.GameBackend.Connect:input_type -> proto.ConnectRequest
	10, // 11: proto.GameBackend.Stream:input_type -> proto.Request
	2,  // 12: proto.GameBackend.Connect:output_type -> proto.ConnectResponse
	11, // 13: proto.GameBackend.Stream:output_type -> proto.Response
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_server_proto_init() }
func file_proto_server_proto_init() {
	if File_proto_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_proto_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
		file_proto_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Move); i {
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
		file_proto_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinate); i {
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
		file_proto_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_proto_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_proto_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEntity); i {
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
		file_proto_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntity); i {
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
		file_proto_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveEntity); i {
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
		file_proto_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_server_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
	file_proto_server_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Entity_Player)(nil),
	}
	file_proto_server_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*Request_Move)(nil),
	}
	file_proto_server_proto_msgTypes[10].OneofWrappers = []interface{}{
		(*Response_AddEntity)(nil),
		(*Response_UpdateEntity)(nil),
		(*Response_RemoveEntity)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_server_proto_goTypes,
		DependencyIndexes: file_proto_server_proto_depIdxs,
		EnumInfos:         file_proto_server_proto_enumTypes,
		MessageInfos:      file_proto_server_proto_msgTypes,
	}.Build()
	File_proto_server_proto = out.File
	file_proto_server_proto_rawDesc = nil
	file_proto_server_proto_goTypes = nil
	file_proto_server_proto_depIdxs = nil
}