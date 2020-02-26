// Code generated by protoc-gen-go. DO NOT EDIT.
// source: x2ethereum.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EthBridgeStatus int32

const (
	EthBridgeStatus_PendingStatusText EthBridgeStatus = 0
	EthBridgeStatus_SuccessStatusText EthBridgeStatus = 1
	EthBridgeStatus_FailedStatusText  EthBridgeStatus = 2
)

var EthBridgeStatus_name = map[int32]string{
	0: "PendingStatusText",
	1: "SuccessStatusText",
	2: "FailedStatusText",
}
var EthBridgeStatus_value = map[string]int32{
	"PendingStatusText": 0,
	"SuccessStatusText": 1,
	"FailedStatusText":  2,
}

func (x EthBridgeStatus) String() string {
	return proto.EnumName(EthBridgeStatus_name, int32(x))
}
func (EthBridgeStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_x2ethereum_9a5e2abc8134bb49, []int{0}
}

type X2EthereumAction struct {
	// Types that are valid to be assigned to Value:
	//	*X2EthereumAction_EthBridgeClaim
	//	*X2EthereumAction_MsgBurn
	//	*X2EthereumAction_MsgLock
	//	*X2EthereumAction_MsgLogInValidator
	Value                isX2EthereumAction_Value `protobuf_oneof:"value"`
	Ty                   int32                    `protobuf:"varint,10,opt,name=ty" json:"ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *X2EthereumAction) Reset()         { *m = X2EthereumAction{} }
func (m *X2EthereumAction) String() string { return proto.CompactTextString(m) }
func (*X2EthereumAction) ProtoMessage()    {}
func (*X2EthereumAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_x2ethereum_9a5e2abc8134bb49, []int{0}
}
func (m *X2EthereumAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_X2EthereumAction.Unmarshal(m, b)
}
func (m *X2EthereumAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_X2EthereumAction.Marshal(b, m, deterministic)
}
func (dst *X2EthereumAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_X2EthereumAction.Merge(dst, src)
}
func (m *X2EthereumAction) XXX_Size() int {
	return xxx_messageInfo_X2EthereumAction.Size(m)
}
func (m *X2EthereumAction) XXX_DiscardUnknown() {
	xxx_messageInfo_X2EthereumAction.DiscardUnknown(m)
}

var xxx_messageInfo_X2EthereumAction proto.InternalMessageInfo

type isX2EthereumAction_Value interface {
	isX2EthereumAction_Value()
}

type X2EthereumAction_EthBridgeClaim struct {
	EthBridgeClaim *EthBridgeClaim `protobuf:"bytes,1,opt,name=ethBridgeClaim,oneof"`
}
type X2EthereumAction_MsgBurn struct {
	MsgBurn *MsgBurn `protobuf:"bytes,2,opt,name=msgBurn,oneof"`
}
type X2EthereumAction_MsgLock struct {
	MsgLock *MsgLock `protobuf:"bytes,3,opt,name=msgLock,oneof"`
}
type X2EthereumAction_MsgLogInValidator struct {
	MsgLogInValidator *MsgLogInValidator `protobuf:"bytes,4,opt,name=msgLogInValidator,oneof"`
}

func (*X2EthereumAction_EthBridgeClaim) isX2EthereumAction_Value()    {}
func (*X2EthereumAction_MsgBurn) isX2EthereumAction_Value()           {}
func (*X2EthereumAction_MsgLock) isX2EthereumAction_Value()           {}
func (*X2EthereumAction_MsgLogInValidator) isX2EthereumAction_Value() {}

func (m *X2EthereumAction) GetValue() isX2EthereumAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *X2EthereumAction) GetEthBridgeClaim() *EthBridgeClaim {
	if x, ok := m.GetValue().(*X2EthereumAction_EthBridgeClaim); ok {
		return x.EthBridgeClaim
	}
	return nil
}

func (m *X2EthereumAction) GetMsgBurn() *MsgBurn {
	if x, ok := m.GetValue().(*X2EthereumAction_MsgBurn); ok {
		return x.MsgBurn
	}
	return nil
}

func (m *X2EthereumAction) GetMsgLock() *MsgLock {
	if x, ok := m.GetValue().(*X2EthereumAction_MsgLock); ok {
		return x.MsgLock
	}
	return nil
}

func (m *X2EthereumAction) GetMsgLogInValidator() *MsgLogInValidator {
	if x, ok := m.GetValue().(*X2EthereumAction_MsgLogInValidator); ok {
		return x.MsgLogInValidator
	}
	return nil
}

func (m *X2EthereumAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*X2EthereumAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _X2EthereumAction_OneofMarshaler, _X2EthereumAction_OneofUnmarshaler, _X2EthereumAction_OneofSizer, []interface{}{
		(*X2EthereumAction_EthBridgeClaim)(nil),
		(*X2EthereumAction_MsgBurn)(nil),
		(*X2EthereumAction_MsgLock)(nil),
		(*X2EthereumAction_MsgLogInValidator)(nil),
	}
}

func _X2EthereumAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*X2EthereumAction)
	// value
	switch x := m.Value.(type) {
	case *X2EthereumAction_EthBridgeClaim:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EthBridgeClaim); err != nil {
			return err
		}
	case *X2EthereumAction_MsgBurn:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MsgBurn); err != nil {
			return err
		}
	case *X2EthereumAction_MsgLock:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MsgLock); err != nil {
			return err
		}
	case *X2EthereumAction_MsgLogInValidator:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MsgLogInValidator); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("X2EthereumAction.Value has unexpected type %T", x)
	}
	return nil
}

func _X2EthereumAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*X2EthereumAction)
	switch tag {
	case 1: // value.ethBridgeClaim
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EthBridgeClaim)
		err := b.DecodeMessage(msg)
		m.Value = &X2EthereumAction_EthBridgeClaim{msg}
		return true, err
	case 2: // value.msgBurn
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MsgBurn)
		err := b.DecodeMessage(msg)
		m.Value = &X2EthereumAction_MsgBurn{msg}
		return true, err
	case 3: // value.msgLock
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MsgLock)
		err := b.DecodeMessage(msg)
		m.Value = &X2EthereumAction_MsgLock{msg}
		return true, err
	case 4: // value.msgLogInValidator
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MsgLogInValidator)
		err := b.DecodeMessage(msg)
		m.Value = &X2EthereumAction_MsgLogInValidator{msg}
		return true, err
	default:
		return false, nil
	}
}

func _X2EthereumAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*X2EthereumAction)
	// value
	switch x := m.Value.(type) {
	case *X2EthereumAction_EthBridgeClaim:
		s := proto.Size(x.EthBridgeClaim)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *X2EthereumAction_MsgBurn:
		s := proto.Size(x.MsgBurn)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *X2EthereumAction_MsgLock:
		s := proto.Size(x.MsgLock)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *X2EthereumAction_MsgLogInValidator:
		s := proto.Size(x.MsgLogInValidator)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type MsgLogInValidator struct {
	Address              string   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Power                float64  `protobuf:"fixed64,2,opt,name=power" json:"power,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgLogInValidator) Reset()         { *m = MsgLogInValidator{} }
func (m *MsgLogInValidator) String() string { return proto.CompactTextString(m) }
func (*MsgLogInValidator) ProtoMessage()    {}
func (*MsgLogInValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_x2ethereum_9a5e2abc8134bb49, []int{1}
}
func (m *MsgLogInValidator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgLogInValidator.Unmarshal(m, b)
}
func (m *MsgLogInValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgLogInValidator.Marshal(b, m, deterministic)
}
func (dst *MsgLogInValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLogInValidator.Merge(dst, src)
}
func (m *MsgLogInValidator) XXX_Size() int {
	return xxx_messageInfo_MsgLogInValidator.Size(m)
}
func (m *MsgLogInValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLogInValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLogInValidator proto.InternalMessageInfo

func (m *MsgLogInValidator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgLogInValidator) GetPower() float64 {
	if m != nil {
		return m.Power
	}
	return 0
}

// EthBridgeClaim is a structure that contains all the data for a particular bridge claim
type EthBridgeClaim struct {
	EthereumChainID       int64    `protobuf:"varint,1,opt,name=EthereumChainID" json:"EthereumChainID,omitempty"`
	BridgeContractAddress string   `protobuf:"bytes,2,opt,name=BridgeContractAddress" json:"BridgeContractAddress,omitempty"`
	Nonce                 int64    `protobuf:"varint,3,opt,name=Nonce" json:"Nonce,omitempty"`
	Symbol                string   `protobuf:"bytes,4,opt,name=Symbol" json:"Symbol,omitempty"`
	TokenContractAddress  string   `protobuf:"bytes,5,opt,name=TokenContractAddress" json:"TokenContractAddress,omitempty"`
	EthereumSender        string   `protobuf:"bytes,6,opt,name=EthereumSender" json:"EthereumSender,omitempty"`
	Chain33Receiver       string   `protobuf:"bytes,7,opt,name=Chain33Receiver" json:"Chain33Receiver,omitempty"`
	ValidatorAddress      string   `protobuf:"bytes,8,opt,name=ValidatorAddress" json:"ValidatorAddress,omitempty"`
	Amount                uint64   `protobuf:"varint,9,opt,name=Amount" json:"Amount,omitempty"`
	ClaimType             int64    `protobuf:"varint,10,opt,name=ClaimType" json:"ClaimType,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *EthBridgeClaim) Reset()         { *m = EthBridgeClaim{} }
func (m *EthBridgeClaim) String() string { return proto.CompactTextString(m) }
func (*EthBridgeClaim) ProtoMessage()    {}
func (*EthBridgeClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_x2ethereum_9a5e2abc8134bb49, []int{2}
}
func (m *EthBridgeClaim) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthBridgeClaim.Unmarshal(m, b)
}
func (m *EthBridgeClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthBridgeClaim.Marshal(b, m, deterministic)
}
func (dst *EthBridgeClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthBridgeClaim.Merge(dst, src)
}
func (m *EthBridgeClaim) XXX_Size() int {
	return xxx_messageInfo_EthBridgeClaim.Size(m)
}
func (m *EthBridgeClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_EthBridgeClaim.DiscardUnknown(m)
}

var xxx_messageInfo_EthBridgeClaim proto.InternalMessageInfo

func (m *EthBridgeClaim) GetEthereumChainID() int64 {
	if m != nil {
		return m.EthereumChainID
	}
	return 0
}

func (m *EthBridgeClaim) GetBridgeContractAddress() string {
	if m != nil {
		return m.BridgeContractAddress
	}
	return ""
}

func (m *EthBridgeClaim) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *EthBridgeClaim) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *EthBridgeClaim) GetTokenContractAddress() string {
	if m != nil {
		return m.TokenContractAddress
	}
	return ""
}

func (m *EthBridgeClaim) GetEthereumSender() string {
	if m != nil {
		return m.EthereumSender
	}
	return ""
}

func (m *EthBridgeClaim) GetChain33Receiver() string {
	if m != nil {
		return m.Chain33Receiver
	}
	return ""
}

func (m *EthBridgeClaim) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *EthBridgeClaim) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *EthBridgeClaim) GetClaimType() int64 {
	if m != nil {
		return m.ClaimType
	}
	return 0
}

// OracleClaimContent is the details of how the content of the claim for each validator will be stored in the oracle
type OracleClaimContent struct {
	Chain33Receiver      string   `protobuf:"bytes,1,opt,name=Chain33Receiver" json:"Chain33Receiver,omitempty"`
	Amount               uint64   `protobuf:"varint,2,opt,name=Amount" json:"Amount,omitempty"`
	ClaimType            int64    `protobuf:"varint,3,opt,name=ClaimType" json:"ClaimType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OracleClaimContent) Reset()         { *m = OracleClaimContent{} }
func (m *OracleClaimContent) String() string { return proto.CompactTextString(m) }
func (*OracleClaimContent) ProtoMessage()    {}
func (*OracleClaimContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_x2ethereum_9a5e2abc8134bb49, []int{3}
}
func (m *OracleClaimContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleClaimContent.Unmarshal(m, b)
}
func (m *OracleClaimContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleClaimContent.Marshal(b, m, deterministic)
}
func (dst *OracleClaimContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleClaimContent.Merge(dst, src)
}
func (m *OracleClaimContent) XXX_Size() int {
	return xxx_messageInfo_OracleClaimContent.Size(m)
}
func (m *OracleClaimContent) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleClaimContent.DiscardUnknown(m)
}

var xxx_messageInfo_OracleClaimContent proto.InternalMessageInfo

func (m *OracleClaimContent) GetChain33Receiver() string {
	if m != nil {
		return m.Chain33Receiver
	}
	return ""
}

func (m *OracleClaimContent) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *OracleClaimContent) GetClaimType() int64 {
	if m != nil {
		return m.ClaimType
	}
	return 0
}

// MsgBurn defines a message for burning coins and triggering a related event
type MsgBurn struct {
	EthereumChainID      int64    `protobuf:"varint,1,opt,name=EthereumChainID" json:"EthereumChainID,omitempty"`
	TokenContract        string   `protobuf:"bytes,2,opt,name=TokenContract" json:"TokenContract,omitempty"`
	Chain33Sender        string   `protobuf:"bytes,3,opt,name=Chain33Sender" json:"Chain33Sender,omitempty"`
	EthereumReceiver     string   `protobuf:"bytes,4,opt,name=EthereumReceiver" json:"EthereumReceiver,omitempty"`
	Amount               uint64   `protobuf:"varint,5,opt,name=Amount" json:"Amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgBurn) Reset()         { *m = MsgBurn{} }
func (m *MsgBurn) String() string { return proto.CompactTextString(m) }
func (*MsgBurn) ProtoMessage()    {}
func (*MsgBurn) Descriptor() ([]byte, []int) {
	return fileDescriptor_x2ethereum_9a5e2abc8134bb49, []int{4}
}
func (m *MsgBurn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBurn.Unmarshal(m, b)
}
func (m *MsgBurn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBurn.Marshal(b, m, deterministic)
}
func (dst *MsgBurn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurn.Merge(dst, src)
}
func (m *MsgBurn) XXX_Size() int {
	return xxx_messageInfo_MsgBurn.Size(m)
}
func (m *MsgBurn) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurn.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurn proto.InternalMessageInfo

func (m *MsgBurn) GetEthereumChainID() int64 {
	if m != nil {
		return m.EthereumChainID
	}
	return 0
}

func (m *MsgBurn) GetTokenContract() string {
	if m != nil {
		return m.TokenContract
	}
	return ""
}

func (m *MsgBurn) GetChain33Sender() string {
	if m != nil {
		return m.Chain33Sender
	}
	return ""
}

func (m *MsgBurn) GetEthereumReceiver() string {
	if m != nil {
		return m.EthereumReceiver
	}
	return ""
}

func (m *MsgBurn) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// MsgLock defines a message for locking coins and triggering a related event
type MsgLock struct {
	EthereumChainID      int64    `protobuf:"varint,1,opt,name=EthereumChainID" json:"EthereumChainID,omitempty"`
	TokenContract        string   `protobuf:"bytes,2,opt,name=TokenContract" json:"TokenContract,omitempty"`
	Chain33Sender        string   `protobuf:"bytes,3,opt,name=Chain33Sender" json:"Chain33Sender,omitempty"`
	EthereumReceiver     string   `protobuf:"bytes,4,opt,name=EthereumReceiver" json:"EthereumReceiver,omitempty"`
	Amount               uint64   `protobuf:"varint,5,opt,name=Amount" json:"Amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgLock) Reset()         { *m = MsgLock{} }
func (m *MsgLock) String() string { return proto.CompactTextString(m) }
func (*MsgLock) ProtoMessage()    {}
func (*MsgLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_x2ethereum_9a5e2abc8134bb49, []int{5}
}
func (m *MsgLock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgLock.Unmarshal(m, b)
}
func (m *MsgLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgLock.Marshal(b, m, deterministic)
}
func (dst *MsgLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLock.Merge(dst, src)
}
func (m *MsgLock) XXX_Size() int {
	return xxx_messageInfo_MsgLock.Size(m)
}
func (m *MsgLock) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLock.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLock proto.InternalMessageInfo

func (m *MsgLock) GetEthereumChainID() int64 {
	if m != nil {
		return m.EthereumChainID
	}
	return 0
}

func (m *MsgLock) GetTokenContract() string {
	if m != nil {
		return m.TokenContract
	}
	return ""
}

func (m *MsgLock) GetChain33Sender() string {
	if m != nil {
		return m.Chain33Sender
	}
	return ""
}

func (m *MsgLock) GetEthereumReceiver() string {
	if m != nil {
		return m.EthereumReceiver
	}
	return ""
}

func (m *MsgLock) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type QueryEthProphecyParams struct {
	EthereumChainID       int64    `protobuf:"varint,1,opt,name=EthereumChainID" json:"EthereumChainID,omitempty"`
	BridgeContractAddress string   `protobuf:"bytes,2,opt,name=BridgeContractAddress" json:"BridgeContractAddress,omitempty"`
	Nonce                 int64    `protobuf:"varint,3,opt,name=Nonce" json:"Nonce,omitempty"`
	Symbol                string   `protobuf:"bytes,4,opt,name=Symbol" json:"Symbol,omitempty"`
	TokenContractAddress  string   `protobuf:"bytes,5,opt,name=TokenContractAddress" json:"TokenContractAddress,omitempty"`
	EthereumSender        string   `protobuf:"bytes,6,opt,name=EthereumSender" json:"EthereumSender,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *QueryEthProphecyParams) Reset()         { *m = QueryEthProphecyParams{} }
func (m *QueryEthProphecyParams) String() string { return proto.CompactTextString(m) }
func (*QueryEthProphecyParams) ProtoMessage()    {}
func (*QueryEthProphecyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_x2ethereum_9a5e2abc8134bb49, []int{6}
}
func (m *QueryEthProphecyParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryEthProphecyParams.Unmarshal(m, b)
}
func (m *QueryEthProphecyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryEthProphecyParams.Marshal(b, m, deterministic)
}
func (dst *QueryEthProphecyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEthProphecyParams.Merge(dst, src)
}
func (m *QueryEthProphecyParams) XXX_Size() int {
	return xxx_messageInfo_QueryEthProphecyParams.Size(m)
}
func (m *QueryEthProphecyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEthProphecyParams.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEthProphecyParams proto.InternalMessageInfo

func (m *QueryEthProphecyParams) GetEthereumChainID() int64 {
	if m != nil {
		return m.EthereumChainID
	}
	return 0
}

func (m *QueryEthProphecyParams) GetBridgeContractAddress() string {
	if m != nil {
		return m.BridgeContractAddress
	}
	return ""
}

func (m *QueryEthProphecyParams) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *QueryEthProphecyParams) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *QueryEthProphecyParams) GetTokenContractAddress() string {
	if m != nil {
		return m.TokenContractAddress
	}
	return ""
}

func (m *QueryEthProphecyParams) GetEthereumSender() string {
	if m != nil {
		return m.EthereumSender
	}
	return ""
}

type QueryEthProphecyResponse struct {
	ID                   string            `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Status               *ProphecyStatus   `protobuf:"bytes,2,opt,name=Status" json:"Status,omitempty"`
	Claims               []*EthBridgeClaim `protobuf:"bytes,3,rep,name=Claims" json:"Claims,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *QueryEthProphecyResponse) Reset()         { *m = QueryEthProphecyResponse{} }
func (m *QueryEthProphecyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEthProphecyResponse) ProtoMessage()    {}
func (*QueryEthProphecyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_x2ethereum_9a5e2abc8134bb49, []int{7}
}
func (m *QueryEthProphecyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryEthProphecyResponse.Unmarshal(m, b)
}
func (m *QueryEthProphecyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryEthProphecyResponse.Marshal(b, m, deterministic)
}
func (dst *QueryEthProphecyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEthProphecyResponse.Merge(dst, src)
}
func (m *QueryEthProphecyResponse) XXX_Size() int {
	return xxx_messageInfo_QueryEthProphecyResponse.Size(m)
}
func (m *QueryEthProphecyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEthProphecyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEthProphecyResponse proto.InternalMessageInfo

func (m *QueryEthProphecyResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *QueryEthProphecyResponse) GetStatus() *ProphecyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *QueryEthProphecyResponse) GetClaims() []*EthBridgeClaim {
	if m != nil {
		return m.Claims
	}
	return nil
}

type ProphecyStatus struct {
	Text                 EthBridgeStatus `protobuf:"varint,1,opt,name=Text,enum=types.EthBridgeStatus" json:"Text,omitempty"`
	FinalClaim           string          `protobuf:"bytes,2,opt,name=FinalClaim" json:"FinalClaim,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProphecyStatus) Reset()         { *m = ProphecyStatus{} }
func (m *ProphecyStatus) String() string { return proto.CompactTextString(m) }
func (*ProphecyStatus) ProtoMessage()    {}
func (*ProphecyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_x2ethereum_9a5e2abc8134bb49, []int{8}
}
func (m *ProphecyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProphecyStatus.Unmarshal(m, b)
}
func (m *ProphecyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProphecyStatus.Marshal(b, m, deterministic)
}
func (dst *ProphecyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProphecyStatus.Merge(dst, src)
}
func (m *ProphecyStatus) XXX_Size() int {
	return xxx_messageInfo_ProphecyStatus.Size(m)
}
func (m *ProphecyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ProphecyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ProphecyStatus proto.InternalMessageInfo

func (m *ProphecyStatus) GetText() EthBridgeStatus {
	if m != nil {
		return m.Text
	}
	return EthBridgeStatus_PendingStatusText
}

func (m *ProphecyStatus) GetFinalClaim() string {
	if m != nil {
		return m.FinalClaim
	}
	return ""
}

func init() {
	proto.RegisterType((*X2EthereumAction)(nil), "types.X2ethereumAction")
	proto.RegisterType((*MsgLogInValidator)(nil), "types.MsgLogInValidator")
	proto.RegisterType((*EthBridgeClaim)(nil), "types.EthBridgeClaim")
	proto.RegisterType((*OracleClaimContent)(nil), "types.OracleClaimContent")
	proto.RegisterType((*MsgBurn)(nil), "types.MsgBurn")
	proto.RegisterType((*MsgLock)(nil), "types.MsgLock")
	proto.RegisterType((*QueryEthProphecyParams)(nil), "types.QueryEthProphecyParams")
	proto.RegisterType((*QueryEthProphecyResponse)(nil), "types.QueryEthProphecyResponse")
	proto.RegisterType((*ProphecyStatus)(nil), "types.ProphecyStatus")
	proto.RegisterEnum("types.EthBridgeStatus", EthBridgeStatus_name, EthBridgeStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for X2Ethereum service

type X2EthereumClient interface {
}

type x2EthereumClient struct {
	cc *grpc.ClientConn
}

func NewX2EthereumClient(cc *grpc.ClientConn) X2EthereumClient {
	return &x2EthereumClient{cc}
}

// Server API for X2Ethereum service

type X2EthereumServer interface {
}

func RegisterX2EthereumServer(s *grpc.Server, srv X2EthereumServer) {
	s.RegisterService(&_X2Ethereum_serviceDesc, srv)
}

var _X2Ethereum_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.x2ethereum",
	HandlerType: (*X2EthereumServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "x2ethereum.proto",
}

func init() { proto.RegisterFile("x2ethereum.proto", fileDescriptor_x2ethereum_9a5e2abc8134bb49) }

var fileDescriptor_x2ethereum_9a5e2abc8134bb49 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x55, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x8d, 0xed, 0x26, 0xf9, 0x65, 0x7e, 0x10, 0xd2, 0x55, 0x5b, 0xf9, 0x80, 0x50, 0x14, 0x21,
	0x14, 0x45, 0xa2, 0x87, 0x94, 0x3b, 0xea, 0x5f, 0xa5, 0x12, 0x7f, 0xca, 0xa6, 0x20, 0x0e, 0x5c,
	0xb6, 0xf6, 0x28, 0xb1, 0xea, 0xec, 0x46, 0xeb, 0x75, 0xa9, 0xaf, 0x1c, 0x10, 0x27, 0xbe, 0x14,
	0x5f, 0x0c, 0x79, 0xbc, 0x6e, 0x1d, 0x27, 0x3d, 0x70, 0x44, 0xdc, 0xbc, 0x6f, 0xde, 0xec, 0xcc,
	0x7b, 0xb3, 0xde, 0x85, 0xde, 0xed, 0x18, 0xcd, 0x1c, 0x35, 0xa6, 0x8b, 0xfd, 0xa5, 0x56, 0x46,
	0xb1, 0xa6, 0xc9, 0x96, 0x98, 0x0c, 0x7e, 0xba, 0xd0, 0xfb, 0x7c, 0x17, 0x3b, 0x0c, 0x4c, 0xa4,
	0x24, 0x7b, 0x0d, 0x5d, 0x34, 0xf3, 0x23, 0x1d, 0x85, 0x33, 0x3c, 0x8e, 0x45, 0xb4, 0xf0, 0x9d,
	0xbe, 0x33, 0xfc, 0x7f, 0xbc, 0xbb, 0x4f, 0x49, 0xfb, 0xa7, 0x2b, 0xc1, 0x49, 0x83, 0xd7, 0xe8,
	0x6c, 0x04, 0xed, 0x45, 0x32, 0x3b, 0x4a, 0xb5, 0xf4, 0x5d, 0xca, 0xec, 0xda, 0xcc, 0xb7, 0x05,
	0x3a, 0x69, 0xf0, 0x92, 0x60, 0xb9, 0x6f, 0x54, 0x70, 0xed, 0x7b, 0x75, 0x6e, 0x8e, 0x5a, 0x6e,
	0xfe, 0xc9, 0x26, 0xb0, 0x4d, 0x9f, 0xb3, 0x73, 0xf9, 0x49, 0xc4, 0x51, 0x28, 0x8c, 0xd2, 0xfe,
	0x16, 0x65, 0xf9, 0xd5, 0xac, 0x6a, 0x7c, 0xd2, 0xe0, 0xeb, 0x49, 0xac, 0x0b, 0xae, 0xc9, 0x7c,
	0xe8, 0x3b, 0xc3, 0x26, 0x77, 0x4d, 0x76, 0xd4, 0x86, 0xe6, 0x8d, 0x88, 0x53, 0x1c, 0x1c, 0xc3,
	0xf6, 0xda, 0x16, 0xcc, 0x87, 0xb6, 0x08, 0x43, 0x8d, 0x49, 0x42, 0x4e, 0x74, 0x78, 0xb9, 0x64,
	0x3b, 0xd0, 0x5c, 0xaa, 0xaf, 0xa8, 0x49, 0xa7, 0xc3, 0x8b, 0xc5, 0xe0, 0xbb, 0x07, 0xdd, 0x55,
	0x93, 0xd8, 0x10, 0x9e, 0x9c, 0x5a, 0x97, 0x8f, 0xe7, 0x22, 0x92, 0xe7, 0x27, 0xb4, 0x95, 0xc7,
	0xeb, 0x30, 0x7b, 0x05, 0xbb, 0x36, 0x51, 0x49, 0xa3, 0x45, 0x60, 0x0e, 0x6d, 0x69, 0x97, 0x4a,
	0x6f, 0x0e, 0xe6, 0x8d, 0xbc, 0x53, 0x32, 0x40, 0x32, 0xd1, 0xe3, 0xc5, 0x82, 0xed, 0x41, 0x6b,
	0x9a, 0x2d, 0xae, 0x54, 0x4c, 0x2e, 0x75, 0xb8, 0x5d, 0xb1, 0x31, 0xec, 0x5c, 0xaa, 0x6b, 0x94,
	0xf5, 0x12, 0x4d, 0x62, 0x6d, 0x8c, 0xb1, 0x17, 0xa4, 0x89, 0x5a, 0x9d, 0xa2, 0x0c, 0x51, 0xfb,
	0x2d, 0x62, 0xd7, 0xd0, 0x5c, 0x29, 0x49, 0x39, 0x38, 0xe0, 0x18, 0x60, 0x74, 0x83, 0xda, 0x6f,
	0x13, 0xb1, 0x0e, 0xb3, 0x11, 0xf4, 0xee, 0x3c, 0x2e, 0x3b, 0xf8, 0x8f, 0xa8, 0x6b, 0x78, 0xae,
	0xe4, 0x70, 0xa1, 0x52, 0x69, 0xfc, 0x4e, 0xdf, 0x19, 0x6e, 0x71, 0xbb, 0x62, 0x4f, 0xa1, 0x43,
	0x06, 0x5f, 0x66, 0x4b, 0xa4, 0x79, 0x7a, 0xfc, 0x1e, 0x18, 0x18, 0x60, 0xef, 0xb5, 0x08, 0xe2,
	0x62, 0x08, 0xb9, 0x22, 0x94, 0x66, 0x53, 0x87, 0xce, 0xe6, 0x0e, 0xef, 0xab, 0xba, 0x0f, 0x57,
	0xf5, 0xea, 0x55, 0x7f, 0x39, 0xd0, 0xb6, 0x27, 0xfd, 0x0f, 0xe6, 0xfe, 0x1c, 0x1e, 0xaf, 0xf8,
	0x6e, 0xe7, 0xbd, 0x0a, 0xe6, 0x2c, 0xdb, 0xa4, 0x1d, 0x82, 0x57, 0xb0, 0x56, 0xc0, 0xdc, 0xd9,
	0x72, 0xfb, 0x3b, 0x89, 0xc5, 0x09, 0x58, 0xc3, 0x2b, 0x1a, 0x9b, 0x55, 0x8d, 0xa5, 0x0a, 0xfa,
	0xf1, 0xfe, 0x5e, 0x15, 0xdf, 0x5c, 0xd8, 0xfb, 0x90, 0xa2, 0xce, 0x4e, 0xcd, 0xfc, 0x42, 0xab,
	0xe5, 0x1c, 0x83, 0xec, 0x42, 0x68, 0xb1, 0x48, 0xfe, 0x9d, 0x5f, 0x72, 0xf0, 0xc3, 0x01, 0xbf,
	0x6e, 0x02, 0xc7, 0x64, 0xa9, 0x64, 0x82, 0xf9, 0x55, 0x68, 0x95, 0x77, 0xb8, 0x7b, 0x7e, 0xc2,
	0x5e, 0x42, 0x6b, 0x6a, 0x84, 0x49, 0x13, 0x7b, 0x77, 0x97, 0xb7, 0x7e, 0x99, 0x58, 0x04, 0xb9,
	0x25, 0xe5, 0x74, 0x3a, 0xf9, 0x89, 0xef, 0xf5, 0xbd, 0x07, 0x1f, 0x09, 0x6e, 0x49, 0x83, 0x2f,
	0xd0, 0x5d, 0xdd, 0x88, 0x8d, 0x60, 0xeb, 0x12, 0x6f, 0x0d, 0x75, 0xd0, 0x1d, 0xef, 0xd5, 0xd3,
	0x6d, 0x39, 0xe2, 0xb0, 0x67, 0x00, 0x67, 0x91, 0x14, 0x71, 0xf1, 0x2a, 0x15, 0xee, 0x57, 0x90,
	0xd1, 0x47, 0x1a, 0x69, 0x35, 0x91, 0xed, 0xc2, 0xf6, 0x05, 0xca, 0x30, 0x92, 0xb3, 0x02, 0xc8,
	0xf7, 0xe9, 0x35, 0x72, 0x78, 0x9a, 0x06, 0x01, 0x26, 0x49, 0x05, 0x76, 0xd8, 0x0e, 0xf4, 0xce,
	0x44, 0x14, 0x63, 0x58, 0x41, 0xdd, 0xf1, 0x23, 0x80, 0xfb, 0x07, 0xf4, 0xaa, 0x45, 0x2f, 0xe8,
	0xc1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0xaf, 0xb9, 0x7d, 0x55, 0x07, 0x00, 0x00,
}
