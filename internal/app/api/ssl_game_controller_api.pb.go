// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_game_controller_api.proto

package api

import (
	fmt "fmt"
	state "github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
	statemachine "github.com/RoboCup-SSL/ssl-game-controller/internal/app/statemachine"
	proto "github.com/golang/protobuf/proto"
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

type Output struct {
	MatchState           *state.State         `protobuf:"bytes,1,opt,name=match_state,json=matchState" json:"match_state,omitempty"`
	GcState              *GameControllerState `protobuf:"bytes,2,opt,name=gc_state,json=gcState" json:"gc_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_03cda82ae2a01c2e, []int{0}
}

func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetMatchState() *state.State {
	if m != nil {
		return m.MatchState
	}
	return nil
}

func (m *Output) GetGcState() *GameControllerState {
	if m != nil {
		return m.GcState
	}
	return nil
}

type GameControllerState struct {
	AutoRefsConnected      []string        `protobuf:"bytes,1,rep,name=auto_refs_connected,json=autoRefsConnected" json:"auto_refs_connected,omitempty"`
	TeamConnected          map[string]bool `protobuf:"bytes,2,rep,name=team_connected,json=teamConnected" json:"team_connected,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	TeamConnectionVerified map[string]bool `protobuf:"bytes,3,rep,name=team_connection_verified,json=teamConnectionVerified" json:"team_connection_verified,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_NoUnkeyedLiteral   struct{}        `json:"-"`
	XXX_unrecognized       []byte          `json:"-"`
	XXX_sizecache          int32           `json:"-"`
}

func (m *GameControllerState) Reset()         { *m = GameControllerState{} }
func (m *GameControllerState) String() string { return proto.CompactTextString(m) }
func (*GameControllerState) ProtoMessage()    {}
func (*GameControllerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_03cda82ae2a01c2e, []int{1}
}

func (m *GameControllerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameControllerState.Unmarshal(m, b)
}
func (m *GameControllerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameControllerState.Marshal(b, m, deterministic)
}
func (m *GameControllerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameControllerState.Merge(m, src)
}
func (m *GameControllerState) XXX_Size() int {
	return xxx_messageInfo_GameControllerState.Size(m)
}
func (m *GameControllerState) XXX_DiscardUnknown() {
	xxx_messageInfo_GameControllerState.DiscardUnknown(m)
}

var xxx_messageInfo_GameControllerState proto.InternalMessageInfo

func (m *GameControllerState) GetAutoRefsConnected() []string {
	if m != nil {
		return m.AutoRefsConnected
	}
	return nil
}

func (m *GameControllerState) GetTeamConnected() map[string]bool {
	if m != nil {
		return m.TeamConnected
	}
	return nil
}

func (m *GameControllerState) GetTeamConnectionVerified() map[string]bool {
	if m != nil {
		return m.TeamConnectionVerified
	}
	return nil
}

type Input struct {
	Change               *statemachine.Change `protobuf:"bytes,1,opt,name=change" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_03cda82ae2a01c2e, []int{2}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetChange() *statemachine.Change {
	if m != nil {
		return m.Change
	}
	return nil
}

func init() {
	proto.RegisterType((*Output)(nil), "Output")
	proto.RegisterType((*GameControllerState)(nil), "GameControllerState")
	proto.RegisterMapType((map[string]bool)(nil), "GameControllerState.TeamConnectedEntry")
	proto.RegisterMapType((map[string]bool)(nil), "GameControllerState.TeamConnectionVerifiedEntry")
	proto.RegisterType((*Input)(nil), "Input")
}

func init() {
	proto.RegisterFile("ssl_game_controller_api.proto", fileDescriptor_03cda82ae2a01c2e)
}

var fileDescriptor_03cda82ae2a01c2e = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x5d, 0x4b, 0xeb, 0x40,
	0x10, 0x25, 0x0d, 0xfd, 0x9a, 0x72, 0x2f, 0xf7, 0x6e, 0x8b, 0x84, 0x8a, 0x34, 0xf4, 0xa5, 0x79,
	0x69, 0x22, 0x7d, 0x12, 0x45, 0x10, 0x83, 0x48, 0x41, 0x14, 0xb6, 0xe2, 0x83, 0x2f, 0x61, 0x9b,
	0x6e, 0xd3, 0xd5, 0x64, 0x37, 0x24, 0x93, 0x42, 0xff, 0x84, 0xbf, 0x59, 0xf2, 0x61, 0x2d, 0x18,
	0x14, 0xdf, 0x66, 0xe6, 0x9c, 0x39, 0x67, 0x98, 0x19, 0x38, 0x49, 0xd3, 0xd0, 0x0b, 0x58, 0xc4,
	0x3d, 0x5f, 0x49, 0x4c, 0x54, 0x18, 0xf2, 0xc4, 0x63, 0xb1, 0xb0, 0xe3, 0x44, 0xa1, 0x1a, 0x8e,
	0xea, 0xe0, 0x14, 0x19, 0xf2, 0x8a, 0x60, 0xd6, 0x11, 0xfc, 0x0d, 0x93, 0x41, 0xc5, 0x18, 0x2f,
	0xa1, 0xf5, 0x90, 0x61, 0x9c, 0x21, 0x99, 0x40, 0x2f, 0x62, 0xe8, 0x6f, 0x4a, 0x01, 0x43, 0x33,
	0x35, 0xab, 0x37, 0x6b, 0xd9, 0x8b, 0x3c, 0xa3, 0x50, 0x40, 0x45, 0x4c, 0x1c, 0xe8, 0x04, 0x7e,
	0xc5, 0x6a, 0x14, 0xac, 0x81, 0x7d, 0xcb, 0x22, 0xee, 0xee, 0x2d, 0xca, 0x9e, 0x76, 0xe0, 0x17,
	0xc1, 0xf8, 0x4d, 0x87, 0x7e, 0x0d, 0x81, 0xd8, 0xd0, 0x67, 0x19, 0x2a, 0x2f, 0xe1, 0xeb, 0x34,
	0x1f, 0x50, 0x72, 0x1f, 0xf9, 0xca, 0xd0, 0x4c, 0xdd, 0xea, 0xd2, 0xff, 0x39, 0x44, 0xf9, 0x3a,
	0x75, 0x3f, 0x00, 0x72, 0x0f, 0x7f, 0x91, 0xb3, 0xe8, 0x80, 0xda, 0x30, 0x75, 0xab, 0x37, 0x9b,
	0xd4, 0xd9, 0xdb, 0x8f, 0x9c, 0x45, 0xfb, 0xde, 0x1b, 0x89, 0xc9, 0x8e, 0xfe, 0xc1, 0xc3, 0x1a,
	0x79, 0x01, 0xe3, 0x50, 0x4f, 0x28, 0xe9, 0x6d, 0x79, 0x22, 0xd6, 0x82, 0xaf, 0x0c, 0xbd, 0x50,
	0x3e, 0xfd, 0x49, 0x59, 0x28, 0xf9, 0x54, 0xb5, 0x94, 0x16, 0x47, 0x58, 0x0b, 0x0e, 0xaf, 0x80,
	0x7c, 0x1d, 0x88, 0xfc, 0x03, 0xfd, 0x95, 0xef, 0x8a, 0x5d, 0x77, 0x69, 0x1e, 0x92, 0x01, 0x34,
	0xb7, 0x2c, 0xcc, 0xca, 0xcd, 0x76, 0x68, 0x99, 0x9c, 0x37, 0xce, 0xb4, 0xe1, 0x1c, 0x8e, 0xbf,
	0x31, 0xfe, 0x8d, 0xd4, 0xd8, 0x82, 0xe6, 0x5c, 0xe6, 0x37, 0x1f, 0x41, 0xab, 0xfc, 0x86, 0xea,
	0xdc, 0x6d, 0xdb, 0x2d, 0x52, 0x5a, 0x95, 0xaf, 0x2f, 0x9f, 0x2f, 0x02, 0x81, 0x9b, 0x6c, 0x69,
	0xfb, 0x2a, 0x72, 0xa8, 0x5a, 0x2a, 0x37, 0x8b, 0xa7, 0x8b, 0xc5, 0x9d, 0x93, 0xa6, 0xe1, 0x34,
	0xff, 0xac, 0xe9, 0xe7, 0x67, 0x39, 0x42, 0x22, 0x4f, 0x24, 0x0b, 0x1d, 0x16, 0xc7, 0x0e, 0x8b,
	0xc5, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0xb8, 0x80, 0x0e, 0xc0, 0x02, 0x00, 0x00,
}
