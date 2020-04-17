// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/main.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Direction int32

const (
	Direction_UP    Direction = 0
	Direction_DOWN  Direction = 1
	Direction_LEFT  Direction = 2
	Direction_RIGHT Direction = 3
	Direction_STOP  Direction = 4
)

var Direction_name = map[int32]string{
	0: "UP",
	1: "DOWN",
	2: "LEFT",
	3: "RIGHT",
	4: "STOP",
}

var Direction_value = map[string]int32{
	"UP":    0,
	"DOWN":  1,
	"LEFT":  2,
	"RIGHT": 3,
	"STOP":  4,
}

func (x Direction) String() string {
	return proto.EnumName(Direction_name, int32(x))
}

func (Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{0}
}

type Coordinate struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinate) Reset()         { *m = Coordinate{} }
func (m *Coordinate) String() string { return proto.CompactTextString(m) }
func (*Coordinate) ProtoMessage()    {}
func (*Coordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{0}
}

func (m *Coordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinate.Unmarshal(m, b)
}
func (m *Coordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinate.Marshal(b, m, deterministic)
}
func (m *Coordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinate.Merge(m, src)
}
func (m *Coordinate) XXX_Size() int {
	return xxx_messageInfo_Coordinate.Size(m)
}
func (m *Coordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinate.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinate proto.InternalMessageInfo

func (m *Coordinate) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Coordinate) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Player struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Position             *Coordinate `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Icon                 string      `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{1}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player) GetPosition() *Coordinate {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Player) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

type Laser struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Direction            Direction            `protobuf:"varint,2,opt,name=direction,proto3,enum=proto.Direction" json:"direction,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	InitialPosition      *Coordinate          `protobuf:"bytes,4,opt,name=initialPosition,proto3" json:"initialPosition,omitempty"`
	OwnerId              string               `protobuf:"bytes,5,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Laser) Reset()         { *m = Laser{} }
func (m *Laser) String() string { return proto.CompactTextString(m) }
func (*Laser) ProtoMessage()    {}
func (*Laser) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{2}
}

func (m *Laser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Laser.Unmarshal(m, b)
}
func (m *Laser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Laser.Marshal(b, m, deterministic)
}
func (m *Laser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Laser.Merge(m, src)
}
func (m *Laser) XXX_Size() int {
	return xxx_messageInfo_Laser.Size(m)
}
func (m *Laser) XXX_DiscardUnknown() {
	xxx_messageInfo_Laser.DiscardUnknown(m)
}

var xxx_messageInfo_Laser proto.InternalMessageInfo

func (m *Laser) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Laser) GetDirection() Direction {
	if m != nil {
		return m.Direction
	}
	return Direction_UP
}

func (m *Laser) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Laser) GetInitialPosition() *Coordinate {
	if m != nil {
		return m.InitialPosition
	}
	return nil
}

func (m *Laser) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Entity struct {
	// Types that are valid to be assigned to Entity:
	//	*Entity_Player
	//	*Entity_Laser
	Entity               isEntity_Entity `protobuf_oneof:"entity"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{3}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

type isEntity_Entity interface {
	isEntity_Entity()
}

type Entity_Player struct {
	Player *Player `protobuf:"bytes,2,opt,name=player,proto3,oneof"`
}

type Entity_Laser struct {
	Laser *Laser `protobuf:"bytes,3,opt,name=laser,proto3,oneof"`
}

func (*Entity_Player) isEntity_Entity() {}

func (*Entity_Laser) isEntity_Entity() {}

func (m *Entity) GetEntity() isEntity_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Entity) GetPlayer() *Player {
	if x, ok := m.GetEntity().(*Entity_Player); ok {
		return x.Player
	}
	return nil
}

func (m *Entity) GetLaser() *Laser {
	if x, ok := m.GetEntity().(*Entity_Laser); ok {
		return x.Laser
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Entity) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Entity_Player)(nil),
		(*Entity_Laser)(nil),
	}
}

type ConnectRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{4}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ConnectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConnectRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ConnectResponse struct {
	Token                string    `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Entities             []*Entity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ConnectResponse) Reset()         { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()    {}
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{5}
}

func (m *ConnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectResponse.Unmarshal(m, b)
}
func (m *ConnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectResponse.Marshal(b, m, deterministic)
}
func (m *ConnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectResponse.Merge(m, src)
}
func (m *ConnectResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectResponse.Size(m)
}
func (m *ConnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectResponse proto.InternalMessageInfo

func (m *ConnectResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ConnectResponse) GetEntities() []*Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

type Move struct {
	Direction            Direction `protobuf:"varint,1,opt,name=direction,proto3,enum=proto.Direction" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Move) Reset()         { *m = Move{} }
func (m *Move) String() string { return proto.CompactTextString(m) }
func (*Move) ProtoMessage()    {}
func (*Move) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{6}
}

func (m *Move) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Move.Unmarshal(m, b)
}
func (m *Move) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Move.Marshal(b, m, deterministic)
}
func (m *Move) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Move.Merge(m, src)
}
func (m *Move) XXX_Size() int {
	return xxx_messageInfo_Move.Size(m)
}
func (m *Move) XXX_DiscardUnknown() {
	xxx_messageInfo_Move.DiscardUnknown(m)
}

var xxx_messageInfo_Move proto.InternalMessageInfo

func (m *Move) GetDirection() Direction {
	if m != nil {
		return m.Direction
	}
	return Direction_UP
}

type AddEntity struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddEntity) Reset()         { *m = AddEntity{} }
func (m *AddEntity) String() string { return proto.CompactTextString(m) }
func (*AddEntity) ProtoMessage()    {}
func (*AddEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{7}
}

func (m *AddEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEntity.Unmarshal(m, b)
}
func (m *AddEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEntity.Marshal(b, m, deterministic)
}
func (m *AddEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEntity.Merge(m, src)
}
func (m *AddEntity) XXX_Size() int {
	return xxx_messageInfo_AddEntity.Size(m)
}
func (m *AddEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEntity.DiscardUnknown(m)
}

var xxx_messageInfo_AddEntity proto.InternalMessageInfo

func (m *AddEntity) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type UpdateEntity struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEntity) Reset()         { *m = UpdateEntity{} }
func (m *UpdateEntity) String() string { return proto.CompactTextString(m) }
func (*UpdateEntity) ProtoMessage()    {}
func (*UpdateEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{8}
}

func (m *UpdateEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEntity.Unmarshal(m, b)
}
func (m *UpdateEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEntity.Marshal(b, m, deterministic)
}
func (m *UpdateEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEntity.Merge(m, src)
}
func (m *UpdateEntity) XXX_Size() int {
	return xxx_messageInfo_UpdateEntity.Size(m)
}
func (m *UpdateEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEntity.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEntity proto.InternalMessageInfo

func (m *UpdateEntity) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type RemoveEntity struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveEntity) Reset()         { *m = RemoveEntity{} }
func (m *RemoveEntity) String() string { return proto.CompactTextString(m) }
func (*RemoveEntity) ProtoMessage()    {}
func (*RemoveEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{9}
}

func (m *RemoveEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveEntity.Unmarshal(m, b)
}
func (m *RemoveEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveEntity.Marshal(b, m, deterministic)
}
func (m *RemoveEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveEntity.Merge(m, src)
}
func (m *RemoveEntity) XXX_Size() int {
	return xxx_messageInfo_RemoveEntity.Size(m)
}
func (m *RemoveEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveEntity.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveEntity proto.InternalMessageInfo

func (m *RemoveEntity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PlayerRespawn struct {
	Player               *Player  `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	KilledById           string   `protobuf:"bytes,2,opt,name=killedById,proto3" json:"killedById,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerRespawn) Reset()         { *m = PlayerRespawn{} }
func (m *PlayerRespawn) String() string { return proto.CompactTextString(m) }
func (*PlayerRespawn) ProtoMessage()    {}
func (*PlayerRespawn) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{10}
}

func (m *PlayerRespawn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerRespawn.Unmarshal(m, b)
}
func (m *PlayerRespawn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerRespawn.Marshal(b, m, deterministic)
}
func (m *PlayerRespawn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerRespawn.Merge(m, src)
}
func (m *PlayerRespawn) XXX_Size() int {
	return xxx_messageInfo_PlayerRespawn.Size(m)
}
func (m *PlayerRespawn) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerRespawn.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerRespawn proto.InternalMessageInfo

func (m *PlayerRespawn) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *PlayerRespawn) GetKilledById() string {
	if m != nil {
		return m.KilledById
	}
	return ""
}

type RoundOver struct {
	RoundWinnerId        string               `protobuf:"bytes,1,opt,name=roundWinnerId,proto3" json:"roundWinnerId,omitempty"`
	NewRoundAt           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=newRoundAt,proto3" json:"newRoundAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RoundOver) Reset()         { *m = RoundOver{} }
func (m *RoundOver) String() string { return proto.CompactTextString(m) }
func (*RoundOver) ProtoMessage()    {}
func (*RoundOver) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{11}
}

func (m *RoundOver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundOver.Unmarshal(m, b)
}
func (m *RoundOver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundOver.Marshal(b, m, deterministic)
}
func (m *RoundOver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundOver.Merge(m, src)
}
func (m *RoundOver) XXX_Size() int {
	return xxx_messageInfo_RoundOver.Size(m)
}
func (m *RoundOver) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundOver.DiscardUnknown(m)
}

var xxx_messageInfo_RoundOver proto.InternalMessageInfo

func (m *RoundOver) GetRoundWinnerId() string {
	if m != nil {
		return m.RoundWinnerId
	}
	return ""
}

func (m *RoundOver) GetNewRoundAt() *timestamp.Timestamp {
	if m != nil {
		return m.NewRoundAt
	}
	return nil
}

type RoundStart struct {
	Players              []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RoundStart) Reset()         { *m = RoundStart{} }
func (m *RoundStart) String() string { return proto.CompactTextString(m) }
func (*RoundStart) ProtoMessage()    {}
func (*RoundStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{12}
}

func (m *RoundStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundStart.Unmarshal(m, b)
}
func (m *RoundStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundStart.Marshal(b, m, deterministic)
}
func (m *RoundStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundStart.Merge(m, src)
}
func (m *RoundStart) XXX_Size() int {
	return xxx_messageInfo_RoundStart.Size(m)
}
func (m *RoundStart) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundStart.DiscardUnknown(m)
}

var xxx_messageInfo_RoundStart proto.InternalMessageInfo

func (m *RoundStart) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

type Request struct {
	// Types that are valid to be assigned to Action:
	//	*Request_Move
	//	*Request_Laser
	Action               isRequest_Action `protobuf_oneof:"action"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{13}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type isRequest_Action interface {
	isRequest_Action()
}

type Request_Move struct {
	Move *Move `protobuf:"bytes,1,opt,name=move,proto3,oneof"`
}

type Request_Laser struct {
	Laser *Laser `protobuf:"bytes,2,opt,name=laser,proto3,oneof"`
}

func (*Request_Move) isRequest_Action() {}

func (*Request_Laser) isRequest_Action() {}

func (m *Request) GetAction() isRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *Request) GetMove() *Move {
	if x, ok := m.GetAction().(*Request_Move); ok {
		return x.Move
	}
	return nil
}

func (m *Request) GetLaser() *Laser {
	if x, ok := m.GetAction().(*Request_Laser); ok {
		return x.Laser
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Request) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Request_Move)(nil),
		(*Request_Laser)(nil),
	}
}

type Response struct {
	// Types that are valid to be assigned to Action:
	//	*Response_AddEntity
	//	*Response_UpdateEntity
	//	*Response_RemoveEntity
	//	*Response_PlayerRespawn
	//	*Response_RoundOver
	//	*Response_RoundStart
	Action               isResponse_Action `protobuf_oneof:"action"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{14}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

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

type Response_PlayerRespawn struct {
	PlayerRespawn *PlayerRespawn `protobuf:"bytes,4,opt,name=playerRespawn,proto3,oneof"`
}

type Response_RoundOver struct {
	RoundOver *RoundOver `protobuf:"bytes,5,opt,name=roundOver,proto3,oneof"`
}

type Response_RoundStart struct {
	RoundStart *RoundStart `protobuf:"bytes,6,opt,name=roundStart,proto3,oneof"`
}

func (*Response_AddEntity) isResponse_Action() {}

func (*Response_UpdateEntity) isResponse_Action() {}

func (*Response_RemoveEntity) isResponse_Action() {}

func (*Response_PlayerRespawn) isResponse_Action() {}

func (*Response_RoundOver) isResponse_Action() {}

func (*Response_RoundStart) isResponse_Action() {}

func (m *Response) GetAction() isResponse_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *Response) GetAddEntity() *AddEntity {
	if x, ok := m.GetAction().(*Response_AddEntity); ok {
		return x.AddEntity
	}
	return nil
}

func (m *Response) GetUpdateEntity() *UpdateEntity {
	if x, ok := m.GetAction().(*Response_UpdateEntity); ok {
		return x.UpdateEntity
	}
	return nil
}

func (m *Response) GetRemoveEntity() *RemoveEntity {
	if x, ok := m.GetAction().(*Response_RemoveEntity); ok {
		return x.RemoveEntity
	}
	return nil
}

func (m *Response) GetPlayerRespawn() *PlayerRespawn {
	if x, ok := m.GetAction().(*Response_PlayerRespawn); ok {
		return x.PlayerRespawn
	}
	return nil
}

func (m *Response) GetRoundOver() *RoundOver {
	if x, ok := m.GetAction().(*Response_RoundOver); ok {
		return x.RoundOver
	}
	return nil
}

func (m *Response) GetRoundStart() *RoundStart {
	if x, ok := m.GetAction().(*Response_RoundStart); ok {
		return x.RoundStart
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Response) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Response_AddEntity)(nil),
		(*Response_UpdateEntity)(nil),
		(*Response_RemoveEntity)(nil),
		(*Response_PlayerRespawn)(nil),
		(*Response_RoundOver)(nil),
		(*Response_RoundStart)(nil),
	}
}

func init() {
	proto.RegisterEnum("proto.Direction", Direction_name, Direction_value)
	proto.RegisterType((*Coordinate)(nil), "proto.Coordinate")
	proto.RegisterType((*Player)(nil), "proto.Player")
	proto.RegisterType((*Laser)(nil), "proto.Laser")
	proto.RegisterType((*Entity)(nil), "proto.Entity")
	proto.RegisterType((*ConnectRequest)(nil), "proto.ConnectRequest")
	proto.RegisterType((*ConnectResponse)(nil), "proto.ConnectResponse")
	proto.RegisterType((*Move)(nil), "proto.Move")
	proto.RegisterType((*AddEntity)(nil), "proto.AddEntity")
	proto.RegisterType((*UpdateEntity)(nil), "proto.UpdateEntity")
	proto.RegisterType((*RemoveEntity)(nil), "proto.RemoveEntity")
	proto.RegisterType((*PlayerRespawn)(nil), "proto.PlayerRespawn")
	proto.RegisterType((*RoundOver)(nil), "proto.RoundOver")
	proto.RegisterType((*RoundStart)(nil), "proto.RoundStart")
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() {
	proto.RegisterFile("proto/main.proto", fileDescriptor_098391ad7281b52b)
}

var fileDescriptor_098391ad7281b52b = []byte{
	// 768 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xeb, 0x44,
	0x14, 0xb5, 0x5d, 0xc7, 0x89, 0x6f, 0x93, 0x34, 0x0c, 0x0f, 0x64, 0x65, 0xf1, 0x28, 0x16, 0xe8,
	0x05, 0x24, 0x92, 0xa7, 0x3c, 0x15, 0x41, 0xe9, 0xa6, 0x5f, 0xd4, 0x95, 0x0a, 0x8d, 0xa6, 0x29,
	0xdd, 0xb0, 0x71, 0xeb, 0xa1, 0x1a, 0x35, 0x9e, 0x31, 0xf6, 0xa4, 0x69, 0xfe, 0x29, 0x5b, 0xfe,
	0x09, 0x9a, 0x0f, 0x7f, 0xa4, 0x14, 0x0a, 0x2b, 0xcf, 0x9d, 0x39, 0xf7, 0xde, 0xb9, 0xe7, 0x1c,
	0x0f, 0x0c, 0xb2, 0x9c, 0x0b, 0x3e, 0x49, 0x63, 0xca, 0xc6, 0x6a, 0x89, 0x5a, 0xea, 0x33, 0xfc,
	0xec, 0x9e, 0xf3, 0xfb, 0x05, 0x99, 0xa8, 0xe8, 0x76, 0xf9, 0xdb, 0x44, 0xd0, 0x94, 0x14, 0x22,
	0x4e, 0x33, 0x8d, 0x0b, 0x47, 0x00, 0xc7, 0x9c, 0xe7, 0x09, 0x65, 0xb1, 0x20, 0xa8, 0x0b, 0xf6,
	0x53, 0x60, 0xef, 0xda, 0xa3, 0x16, 0xb6, 0x9f, 0x64, 0xb4, 0x0e, 0x1c, 0x1d, 0xad, 0x43, 0x0e,
	0xde, 0x6c, 0x11, 0xaf, 0x49, 0x8e, 0xfa, 0xe0, 0xd0, 0x44, 0xc1, 0x7c, 0xec, 0xd0, 0x04, 0x21,
	0x70, 0x59, 0x9c, 0x12, 0x05, 0xf5, 0xb1, 0x5a, 0xa3, 0x6f, 0xa0, 0x93, 0xf1, 0x82, 0x0a, 0xca,
	0x59, 0xb0, 0xb5, 0x6b, 0x8f, 0xb6, 0xa7, 0x1f, 0xe9, 0x8e, 0xe3, 0xba, 0x1d, 0xae, 0x20, 0xb2,
	0x04, 0xbd, 0xe3, 0x2c, 0x70, 0x75, 0x09, 0xb9, 0x0e, 0xff, 0xb0, 0xa1, 0x75, 0x11, 0x17, 0x2f,
	0x34, 0x1c, 0x83, 0x9f, 0xd0, 0x9c, 0xdc, 0xa9, 0xea, 0xb2, 0x6b, 0x7f, 0x3a, 0x30, 0xd5, 0x4f,
	0xca, 0x7d, 0x5c, 0x43, 0xd0, 0x77, 0xe0, 0x17, 0x22, 0xce, 0xc5, 0x9c, 0xa6, 0xc4, 0xdc, 0x66,
	0x38, 0xd6, 0xcc, 0x8c, 0x4b, 0x66, 0xc6, 0xf3, 0x92, 0x19, 0x5c, 0x83, 0xd1, 0x0f, 0xb0, 0x43,
	0x19, 0x15, 0x34, 0x5e, 0xcc, 0xca, 0x69, 0xdc, 0x7f, 0x9a, 0xe6, 0x39, 0x12, 0x05, 0xd0, 0xe6,
	0x2b, 0x46, 0xf2, 0xf3, 0x24, 0x68, 0xa9, 0xbb, 0x97, 0x61, 0x18, 0x83, 0x77, 0xca, 0x04, 0x15,
	0x6b, 0xf4, 0x0e, 0xbc, 0x4c, 0xb1, 0xaa, 0xe6, 0xd8, 0x9e, 0xf6, 0x4c, 0x5d, 0x4d, 0x75, 0x64,
	0x61, 0x73, 0x8c, 0xbe, 0x80, 0xd6, 0x42, 0x92, 0x61, 0xee, 0xdf, 0x35, 0x38, 0x45, 0x50, 0x64,
	0x61, 0x7d, 0x78, 0xd4, 0x01, 0x8f, 0xa8, 0xc2, 0xe1, 0x0c, 0xfa, 0xc7, 0x9c, 0x31, 0x72, 0x27,
	0x30, 0xf9, 0x7d, 0x49, 0x0a, 0xf1, 0x9f, 0x64, 0x1b, 0x42, 0x27, 0x8b, 0x8b, 0x62, 0xc5, 0xf3,
	0x44, 0x35, 0xf2, 0x71, 0x15, 0x87, 0x18, 0x76, 0xaa, 0x8a, 0x45, 0xc6, 0x59, 0x41, 0xd0, 0x1b,
	0x68, 0x09, 0xfe, 0x40, 0x98, 0xa9, 0xaa, 0x03, 0xf4, 0x15, 0x74, 0xd4, 0x25, 0x28, 0x29, 0x02,
	0x67, 0x77, 0xab, 0x31, 0x95, 0x1e, 0x1a, 0x57, 0xc7, 0xe1, 0xb7, 0xe0, 0xfe, 0xc4, 0x1f, 0xc9,
	0xa6, 0xa2, 0xf6, 0xab, 0x8a, 0x86, 0x53, 0xf0, 0x0f, 0x93, 0xc4, 0x70, 0xf8, 0x65, 0x39, 0xb4,
	0xca, 0xfc, 0x5b, 0xb7, 0x92, 0x91, 0x3d, 0xe8, 0x5e, 0x67, 0x49, 0x2c, 0xc8, 0xff, 0x4b, 0x7b,
	0x0b, 0x5d, 0x4c, 0x52, 0xfe, 0x58, 0xa6, 0x3d, 0xa3, 0x31, 0xfc, 0x05, 0x7a, 0x5a, 0x2c, 0xc9,
	0x4a, 0xbc, 0x62, 0xb2, 0xae, 0x91, 0xd4, 0x7e, 0x41, 0xd2, 0x4a, 0xd0, 0xb7, 0x00, 0x0f, 0x74,
	0xb1, 0x20, 0xc9, 0xd1, 0xfa, 0x3c, 0x31, 0x22, 0x34, 0x76, 0xc2, 0x14, 0x7c, 0xcc, 0x97, 0x2c,
	0xb9, 0x7c, 0x54, 0xea, 0xf7, 0x72, 0x19, 0xdc, 0x50, 0xa6, 0x0d, 0xa5, 0xfb, 0x6f, 0x6e, 0xa2,
	0x7d, 0x00, 0x46, 0x56, 0x2a, 0xeb, 0x50, 0x18, 0x43, 0xfd, 0x9b, 0xd1, 0x1b, 0xe8, 0x70, 0x0f,
	0x40, 0x2d, 0xaf, 0xa4, 0xf7, 0xd1, 0x3b, 0x68, 0xeb, 0x6b, 0x16, 0x81, 0xbd, 0xa1, 0xa0, 0x19,
	0xa2, 0x3c, 0x0d, 0x7f, 0x85, 0x76, 0xe9, 0xaf, 0xcf, 0xc1, 0x95, 0x34, 0x99, 0xa9, 0xb7, 0x4d,
	0x82, 0x94, 0x37, 0xb2, 0xb0, 0x3a, 0xaa, 0x4d, 0xec, 0xbc, 0x62, 0xe2, 0x58, 0xcb, 0xfc, 0xa7,
	0x03, 0x9d, 0xca, 0x6c, 0xef, 0xc1, 0x8f, 0x4b, 0xcd, 0x4d, 0x93, 0xd2, 0x23, 0x95, 0x17, 0x22,
	0x0b, 0xd7, 0x20, 0xf4, 0x3d, 0x74, 0x97, 0x0d, 0xc5, 0x4d, 0xd7, 0x8f, 0x4d, 0x52, 0xd3, 0x0c,
	0x91, 0x85, 0x37, 0xa0, 0x32, 0x35, 0x6f, 0xa8, 0x6e, 0xfe, 0xba, 0x32, 0xb5, 0x69, 0x08, 0x99,
	0xda, 0x84, 0xa2, 0x03, 0xe8, 0x65, 0x4d, 0x43, 0x98, 0x17, 0xe3, 0xcd, 0x26, 0x83, 0xfa, 0x2c,
	0xb2, 0xf0, 0x26, 0x58, 0x4e, 0x99, 0x97, 0xb2, 0xab, 0x67, 0xa3, 0x9e, 0xb2, 0xb2, 0x83, 0x9c,
	0xb2, 0x02, 0xa1, 0x0f, 0x00, 0x79, 0xa5, 0x5c, 0xe0, 0x6d, 0x3c, 0x4f, 0xb5, 0xa4, 0x91, 0x85,
	0x1b, 0xb0, 0x9a, 0xe3, 0xaf, 0x0f, 0xc0, 0xaf, 0x7e, 0x31, 0xe4, 0x81, 0x73, 0x3d, 0x1b, 0x58,
	0xa8, 0x03, 0xee, 0xc9, 0xe5, 0xcd, 0xcf, 0x03, 0x5b, 0xae, 0x2e, 0x4e, 0x7f, 0x9c, 0x0f, 0x1c,
	0xe4, 0x43, 0x0b, 0x9f, 0x9f, 0x45, 0xf3, 0xc1, 0x96, 0xdc, 0xbc, 0x9a, 0x5f, 0xce, 0x06, 0xee,
	0xb4, 0x00, 0xf7, 0x4c, 0x3e, 0x1c, 0xfb, 0xd0, 0x36, 0x8f, 0x03, 0xfa, 0xa4, 0x7a, 0x1a, 0x9b,
	0xcf, 0xcf, 0xf0, 0xd3, 0xe7, 0xdb, 0x5a, 0xd6, 0xd0, 0x42, 0x13, 0xf0, 0xae, 0x44, 0x4e, 0xe2,
	0x14, 0xf5, 0x2b, 0x7e, 0x75, 0xce, 0x4e, 0x15, 0x97, 0xe0, 0x91, 0xfd, 0xde, 0xbe, 0xf5, 0xd4,
	0xee, 0x87, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x54, 0x08, 0x41, 0xf7, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (Game_StreamClient, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/proto.Game/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Game_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Game_serviceDesc.Streams[0], "/proto.Game/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameStreamClient{stream}
	return x, nil
}

type Game_StreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type gameStreamClient struct {
	grpc.ClientStream
}

func (x *gameStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServer is the server API for Game service.
type GameServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Stream(Game_StreamServer) error
}

// UnimplementedGameServer can be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (*UnimplementedGameServer) Connect(ctx context.Context, req *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedGameServer) Stream(srv Game_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

func RegisterGameServer(s *grpc.Server, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Game/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServer).Stream(&gameStreamServer{stream})
}

type Game_StreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type gameStreamServer struct {
	grpc.ServerStream
}

func (x *gameStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Game_Connect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Game_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/main.proto",
}
