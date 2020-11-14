// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages-EvolutionStellarToken.proto

package trezor

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

//*
// Request: Ask device for public key corresponding to address_n path
// @start
// @next EvolutionStellarTokenPublicKey
// @next Failure
type EvolutionStellarTokenGetPublicKey struct {
	AddressN             []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	ShowDisplay          *bool    `protobuf:"varint,2,opt,name=show_display,json=showDisplay" json:"show_display,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvolutionStellarTokenGetPublicKey) Reset()         { *m = EvolutionStellarTokenGetPublicKey{} }
func (m *EvolutionStellarTokenGetPublicKey) String() string { return proto.CompactTextString(m) }
func (*EvolutionStellarTokenGetPublicKey) ProtoMessage()    {}
func (*EvolutionStellarTokenGetPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{0}
}

func (m *EvolutionStellarTokenGetPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvolutionStellarTokenGetPublicKey.Unmarshal(m, b)
}
func (m *EvolutionStellarTokenGetPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvolutionStellarTokenGetPublicKey.Marshal(b, m, deterministic)
}
func (m *EvolutionStellarTokenGetPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvolutionStellarTokenGetPublicKey.Merge(m, src)
}
func (m *EvolutionStellarTokenGetPublicKey) XXX_Size() int {
	return xxx_messageInfo_EvolutionStellarTokenGetPublicKey.Size(m)
}
func (m *EvolutionStellarTokenGetPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_EvolutionStellarTokenGetPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_EvolutionStellarTokenGetPublicKey proto.InternalMessageInfo

func (m *EvolutionStellarTokenGetPublicKey) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *EvolutionStellarTokenGetPublicKey) GetShowDisplay() bool {
	if m != nil && m.ShowDisplay != nil {
		return *m.ShowDisplay
	}
	return false
}

//*
// Response: Contains public key derived from device private seed
// @end
type EvolutionStellarTokenPublicKey struct {
	Node                 *HDNodeType `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	Xpub                 *string     `protobuf:"bytes,2,opt,name=xpub" json:"xpub,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EvolutionStellarTokenPublicKey) Reset()         { *m = EvolutionStellarTokenPublicKey{} }
func (m *EvolutionStellarTokenPublicKey) String() string { return proto.CompactTextString(m) }
func (*EvolutionStellarTokenPublicKey) ProtoMessage()    {}
func (*EvolutionStellarTokenPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{1}
}

func (m *EvolutionStellarTokenPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvolutionStellarTokenPublicKey.Unmarshal(m, b)
}
func (m *EvolutionStellarTokenPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvolutionStellarTokenPublicKey.Marshal(b, m, deterministic)
}
func (m *EvolutionStellarTokenPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvolutionStellarTokenPublicKey.Merge(m, src)
}
func (m *EvolutionStellarTokenPublicKey) XXX_Size() int {
	return xxx_messageInfo_EvolutionStellarTokenPublicKey.Size(m)
}
func (m *EvolutionStellarTokenPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_EvolutionStellarTokenPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_EvolutionStellarTokenPublicKey proto.InternalMessageInfo

func (m *EvolutionStellarTokenPublicKey) GetNode() *HDNodeType {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *EvolutionStellarTokenPublicKey) GetXpub() string {
	if m != nil && m.Xpub != nil {
		return *m.Xpub
	}
	return ""
}

//*
// Request: Ask device for EvolutionStellarToken address corresponding to address_n path
// @start
// @next EvolutionStellarTokenAddress
// @next Failure
type EvolutionStellarTokenGetAddress struct {
	AddressN             []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	ShowDisplay          *bool    `protobuf:"varint,2,opt,name=show_display,json=showDisplay" json:"show_display,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvolutionStellarTokenGetAddress) Reset()         { *m = EvolutionStellarTokenGetAddress{} }
func (m *EvolutionStellarTokenGetAddress) String() string { return proto.CompactTextString(m) }
func (*EvolutionStellarTokenGetAddress) ProtoMessage()    {}
func (*EvolutionStellarTokenGetAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{2}
}

func (m *EvolutionStellarTokenGetAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvolutionStellarTokenGetAddress.Unmarshal(m, b)
}
func (m *EvolutionStellarTokenGetAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvolutionStellarTokenGetAddress.Marshal(b, m, deterministic)
}
func (m *EvolutionStellarTokenGetAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvolutionStellarTokenGetAddress.Merge(m, src)
}
func (m *EvolutionStellarTokenGetAddress) XXX_Size() int {
	return xxx_messageInfo_EvolutionStellarTokenGetAddress.Size(m)
}
func (m *EvolutionStellarTokenGetAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_EvolutionStellarTokenGetAddress.DiscardUnknown(m)
}

var xxx_messageInfo_EvolutionStellarTokenGetAddress proto.InternalMessageInfo

func (m *EvolutionStellarTokenGetAddress) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *EvolutionStellarTokenGetAddress) GetShowDisplay() bool {
	if m != nil && m.ShowDisplay != nil {
		return *m.ShowDisplay
	}
	return false
}

//*
// Response: Contains an EvolutionStellarToken address derived from device private seed
// @end
type EvolutionStellarTokenAddress struct {
	AddressBin           []byte   `protobuf:"bytes,1,opt,name=addressBin" json:"addressBin,omitempty"`
	AddressHex           *string  `protobuf:"bytes,2,opt,name=addressHex" json:"addressHex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvolutionStellarTokenAddress) Reset()         { *m = EvolutionStellarTokenAddress{} }
func (m *EvolutionStellarTokenAddress) String() string { return proto.CompactTextString(m) }
func (*EvolutionStellarTokenAddress) ProtoMessage()    {}
func (*EvolutionStellarTokenAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{3}
}

func (m *EvolutionStellarTokenAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvolutionStellarTokenAddress.Unmarshal(m, b)
}
func (m *EvolutionStellarTokenAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvolutionStellarTokenAddress.Marshal(b, m, deterministic)
}
func (m *EvolutionStellarTokenAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvolutionStellarTokenAddress.Merge(m, src)
}
func (m *EvolutionStellarTokenAddress) XXX_Size() int {
	return xxx_messageInfo_EvolutionStellarTokenAddress.Size(m)
}
func (m *EvolutionStellarTokenAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_EvolutionStellarTokenAddress.DiscardUnknown(m)
}

var xxx_messageInfo_EvolutionStellarTokenAddress proto.InternalMessageInfo

func (m *EvolutionStellarTokenAddress) GetAddressBin() []byte {
	if m != nil {
		return m.AddressBin
	}
	return nil
}

func (m *EvolutionStellarTokenAddress) GetAddressHex() string {
	if m != nil && m.AddressHex != nil {
		return *m.AddressHex
	}
	return ""
}

//*
// Request: Ask device to sign transaction
// All fields are optional from the protocol's point of view. Each field defaults to value `0` if missing.
// Note: the first at most 1024 bytes of data MUST be transmitted as part of this message.
// @start
// @next EvolutionStellarTokenTxRequest
// @next Failure
type EvolutionStellarTokenSignTx struct {
	AddressN             []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	Nonce                []byte   `protobuf:"bytes,2,opt,name=nonce" json:"nonce,omitempty"`
	GasPrice             []byte   `protobuf:"bytes,3,opt,name=gas_price,json=gasPrice" json:"gas_price,omitempty"`
	GasLimit             []byte   `protobuf:"bytes,4,opt,name=gas_limit,json=gasLimit" json:"gas_limit,omitempty"`
	ToBin                []byte   `protobuf:"bytes,5,opt,name=toBin" json:"toBin,omitempty"`
	ToHex                *string  `protobuf:"bytes,11,opt,name=toHex" json:"toHex,omitempty"`
	Value                []byte   `protobuf:"bytes,6,opt,name=value" json:"value,omitempty"`
	DataInitialChunk     []byte   `protobuf:"bytes,7,opt,name=data_initial_chunk,json=dataInitialChunk" json:"data_initial_chunk,omitempty"`
	DataLength           *uint32  `protobuf:"varint,8,opt,name=data_length,json=dataLength" json:"data_length,omitempty"`
	ChainId              *uint32  `protobuf:"varint,9,opt,name=chain_id,json=chainId" json:"chain_id,omitempty"`
	TxType               *uint32  `protobuf:"varint,10,opt,name=tx_type,json=txType" json:"tx_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvolutionStellarTokenSignTx) Reset()         { *m = EvolutionStellarTokenSignTx{} }
func (m *EvolutionStellarTokenSignTx) String() string { return proto.CompactTextString(m) }
func (*EvolutionStellarTokenSignTx) ProtoMessage()    {}
func (*EvolutionStellarTokenSignTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{4}
}

func (m *EvolutionStellarTokenSignTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvolutionStellarTokenSignTx.Unmarshal(m, b)
}
func (m *EvolutionStellarTokenSignTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvolutionStellarTokenSignTx.Marshal(b, m, deterministic)
}
func (m *EvolutionStellarTokenSignTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvolutionStellarTokenSignTx.Merge(m, src)
}
func (m *EvolutionStellarTokenSignTx) XXX_Size() int {
	return xxx_messageInfo_EvolutionStellarTokenSignTx.Size(m)
}
func (m *EvolutionStellarTokenSignTx) XXX_DiscardUnknown() {
	xxx_messageInfo_EvolutionStellarTokenSignTx.DiscardUnknown(m)
}

var xxx_messageInfo_EvolutionStellarTokenSignTx proto.InternalMessageInfo

func (m *EvolutionStellarTokenSignTx) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *EvolutionStellarTokenSignTx) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *EvolutionStellarTokenSignTx) GetGasPrice() []byte {
	if m != nil {
		return m.GasPrice
	}
	return nil
}

func (m *EvolutionStellarTokenSignTx) GetGasLimit() []byte {
	if m != nil {
		return m.GasLimit
	}
	return nil
}

func (m *EvolutionStellarTokenSignTx) GetToBin() []byte {
	if m != nil {
		return m.ToBin
	}
	return nil
}

func (m *EvolutionStellarTokenSignTx) GetToHex() string {
	if m != nil && m.ToHex != nil {
		return *m.ToHex
	}
	return ""
}

func (m *EvolutionStellarTokenSignTx) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *EvolutionStellarTokenSignTx) GetDataInitialChunk() []byte {
	if m != nil {
		return m.DataInitialChunk
	}
	return nil
}

func (m *EvolutionStellarTokenSignTx) GetDataLength() uint32 {
	if m != nil && m.DataLength != nil {
		return *m.DataLength
	}
	return 0
}

func (m *EvolutionStellarTokenSignTx) GetChainId() uint32 {
	if m != nil && m.ChainId != nil {
		return *m.ChainId
	}
	return 0
}

func (m *EvolutionStellarTokenSignTx) GetTxType() uint32 {
	if m != nil && m.TxType != nil {
		return *m.TxType
	}
	return 0
}

//*
// Response: Device asks for more data from transaction payload, or returns the signature.
// If data_length is set, device awaits that many more bytes of payload.
// Otherwise, the signature_* fields contain the computed transaction signature. All three fields will be present.
// @end
// @next EvolutionStellarTokenTxAck
type EvolutionStellarTokenTxRequest struct {
	DataLength           *uint32  `protobuf:"varint,1,opt,name=data_length,json=dataLength" json:"data_length,omitempty"`
	SignatureV           *uint32  `protobuf:"varint,2,opt,name=signature_v,json=signatureV" json:"signature_v,omitempty"`
	SignatureR           []byte   `protobuf:"bytes,3,opt,name=signature_r,json=signatureR" json:"signature_r,omitempty"`
	SignatureS           []byte   `protobuf:"bytes,4,opt,name=signature_s,json=signatureS" json:"signature_s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvolutionStellarTokenTxRequest) Reset()         { *m = EvolutionStellarTokenTxRequest{} }
func (m *EvolutionStellarTokenTxRequest) String() string { return proto.CompactTextString(m) }
func (*EvolutionStellarTokenTxRequest) ProtoMessage()    {}
func (*EvolutionStellarTokenTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{5}
}

func (m *EvolutionStellarTokenTxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvolutionStellarTokenTxRequest.Unmarshal(m, b)
}
func (m *EvolutionStellarTokenTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvolutionStellarTokenTxRequest.Marshal(b, m, deterministic)
}
func (m *EvolutionStellarTokenTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvolutionStellarTokenTxRequest.Merge(m, src)
}
func (m *EvolutionStellarTokenTxRequest) XXX_Size() int {
	return xxx_messageInfo_EvolutionStellarTokenTxRequest.Size(m)
}
func (m *EvolutionStellarTokenTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EvolutionStellarTokenTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EvolutionStellarTokenTxRequest proto.InternalMessageInfo

func (m *EvolutionStellarTokenTxRequest) GetDataLength() uint32 {
	if m != nil && m.DataLength != nil {
		return *m.DataLength
	}
	return 0
}

func (m *EvolutionStellarTokenTxRequest) GetSignatureV() uint32 {
	if m != nil && m.SignatureV != nil {
		return *m.SignatureV
	}
	return 0
}

func (m *EvolutionStellarTokenTxRequest) GetSignatureR() []byte {
	if m != nil {
		return m.SignatureR
	}
	return nil
}

func (m *EvolutionStellarTokenTxRequest) GetSignatureS() []byte {
	if m != nil {
		return m.SignatureS
	}
	return nil
}

//*
// Request: Transaction payload data.
// @next EvolutionStellarTokenTxRequest
type EvolutionStellarTokenTxAck struct {
	DataChunk            []byte   `protobuf:"bytes,1,opt,name=data_chunk,json=dataChunk" json:"data_chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvolutionStellarTokenTxAck) Reset()         { *m = EvolutionStellarTokenTxAck{} }
func (m *EvolutionStellarTokenTxAck) String() string { return proto.CompactTextString(m) }
func (*EvolutionStellarTokenTxAck) ProtoMessage()    {}
func (*EvolutionStellarTokenTxAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{6}
}

func (m *EvolutionStellarTokenTxAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvolutionStellarTokenTxAck.Unmarshal(m, b)
}
func (m *EvolutionStellarTokenTxAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvolutionStellarTokenTxAck.Marshal(b, m, deterministic)
}
func (m *EvolutionStellarTokenTxAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvolutionStellarTokenTxAck.Merge(m, src)
}
func (m *EvolutionStellarTokenTxAck) XXX_Size() int {
	return xxx_messageInfo_EvolutionStellarTokenTxAck.Size(m)
}
func (m *EvolutionStellarTokenTxAck) XXX_DiscardUnknown() {
	xxx_messageInfo_EvolutionStellarTokenTxAck.DiscardUnknown(m)
}

var xxx_messageInfo_EvolutionStellarTokenTxAck proto.InternalMessageInfo

func (m *EvolutionStellarTokenTxAck) GetDataChunk() []byte {
	if m != nil {
		return m.DataChunk
	}
	return nil
}

//*
// Request: Ask device to sign message
// @start
// @next EvolutionStellarTokenMessageSignature
// @next Failure
type EvolutionStellarTokenSignMessage struct {
	AddressN             []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	Message              []byte   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvolutionStellarTokenSignMessage) Reset()         { *m = EvolutionStellarTokenSignMessage{} }
func (m *EvolutionStellarTokenSignMessage) String() string { return proto.CompactTextString(m) }
func (*EvolutionStellarTokenSignMessage) ProtoMessage()    {}
func (*EvolutionStellarTokenSignMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{7}
}

func (m *EvolutionStellarTokenSignMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvolutionStellarTokenSignMessage.Unmarshal(m, b)
}
func (m *EvolutionStellarTokenSignMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvolutionStellarTokenSignMessage.Marshal(b, m, deterministic)
}
func (m *EvolutionStellarTokenSignMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvolutionStellarTokenSignMessage.Merge(m, src)
}
func (m *EvolutionStellarTokenSignMessage) XXX_Size() int {
	return xxx_messageInfo_EvolutionStellarTokenSignMessage.Size(m)
}
func (m *EvolutionStellarTokenSignMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EvolutionStellarTokenSignMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EvolutionStellarTokenSignMessage proto.InternalMessageInfo

func (m *EvolutionStellarTokenSignMessage) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *EvolutionStellarTokenSignMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

//*
// Response: Signed message
// @end
type EvolutionStellarTokenMessageSignature struct {
	AddressBin           []byte   `protobuf:"bytes,1,opt,name=addressBin" json:"addressBin,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	AddressHex           *string  `protobuf:"bytes,3,opt,name=addressHex" json:"addressHex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvolutionStellarTokenMessageSignature) Reset()         { *m = EvolutionStellarTokenMessageSignature{} }
func (m *EvolutionStellarTokenMessageSignature) String() string { return proto.CompactTextString(m) }
func (*EvolutionStellarTokenMessageSignature) ProtoMessage()    {}
func (*EvolutionStellarTokenMessageSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{8}
}

func (m *EvolutionStellarTokenMessageSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvolutionStellarTokenMessageSignature.Unmarshal(m, b)
}
func (m *EvolutionStellarTokenMessageSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvolutionStellarTokenMessageSignature.Marshal(b, m, deterministic)
}
func (m *EvolutionStellarTokenMessageSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvolutionStellarTokenMessageSignature.Merge(m, src)
}
func (m *EvolutionStellarTokenMessageSignature) XXX_Size() int {
	return xxx_messageInfo_EvolutionStellarTokenMessageSignature.Size(m)
}
func (m *EvolutionStellarTokenMessageSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_EvolutionStellarTokenMessageSignature.DiscardUnknown(m)
}

var xxx_messageInfo_EvolutionStellarTokenMessageSignature proto.InternalMessageInfo

func (m *EvolutionStellarTokenMessageSignature) GetAddressBin() []byte {
	if m != nil {
		return m.AddressBin
	}
	return nil
}

func (m *EvolutionStellarTokenMessageSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *EvolutionStellarTokenMessageSignature) GetAddressHex() string {
	if m != nil && m.AddressHex != nil {
		return *m.AddressHex
	}
	return ""
}

//*
// Request: Ask device to verify message
// @start
// @next Success
// @next Failure
type EvolutionStellarTokenVerifyMessage struct {
	AddressBin           []byte   `protobuf:"bytes,1,opt,name=addressBin" json:"addressBin,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	Message              []byte   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	AddressHex           *string  `protobuf:"bytes,4,opt,name=addressHex" json:"addressHex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvolutionStellarTokenVerifyMessage) Reset()         { *m = EvolutionStellarTokenVerifyMessage{} }
func (m *EvolutionStellarTokenVerifyMessage) String() string { return proto.CompactTextString(m) }
func (*EvolutionStellarTokenVerifyMessage) ProtoMessage()    {}
func (*EvolutionStellarTokenVerifyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{9}
}

func (m *EvolutionStellarTokenVerifyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvolutionStellarTokenVerifyMessage.Unmarshal(m, b)
}
func (m *EvolutionStellarTokenVerifyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvolutionStellarTokenVerifyMessage.Marshal(b, m, deterministic)
}
func (m *EvolutionStellarTokenVerifyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvolutionStellarTokenVerifyMessage.Merge(m, src)
}
func (m *EvolutionStellarTokenVerifyMessage) XXX_Size() int {
	return xxx_messageInfo_EvolutionStellarTokenVerifyMessage.Size(m)
}
func (m *EvolutionStellarTokenVerifyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EvolutionStellarTokenVerifyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EvolutionStellarTokenVerifyMessage proto.InternalMessageInfo

func (m *EvolutionStellarTokenVerifyMessage) GetAddressBin() []byte {
	if m != nil {
		return m.AddressBin
	}
	return nil
}

func (m *EvolutionStellarTokenVerifyMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *EvolutionStellarTokenVerifyMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *EvolutionStellarTokenVerifyMessage) GetAddressHex() string {
	if m != nil && m.AddressHex != nil {
		return *m.AddressHex
	}
	return ""
}

func init() {
	proto.RegisterType((*EvolutionStellarTokenGetPublicKey)(nil), "hw.trezor.messages.EvolutionStellarToken.EvolutionStellarTokenGetPublicKey")
	proto.RegisterType((*EvolutionStellarTokenPublicKey)(nil), "hw.trezor.messages.EvolutionStellarToken.EvolutionStellarTokenPublicKey")
	proto.RegisterType((*EvolutionStellarTokenGetAddress)(nil), "hw.trezor.messages.EvolutionStellarToken.EvolutionStellarTokenGetAddress")
	proto.RegisterType((*EvolutionStellarTokenAddress)(nil), "hw.trezor.messages.EvolutionStellarToken.EvolutionStellarTokenAddress")
	proto.RegisterType((*EvolutionStellarTokenSignTx)(nil), "hw.trezor.messages.EvolutionStellarToken.EvolutionStellarTokenSignTx")
	proto.RegisterType((*EvolutionStellarTokenTxRequest)(nil), "hw.trezor.messages.EvolutionStellarToken.EvolutionStellarTokenTxRequest")
	proto.RegisterType((*EvolutionStellarTokenTxAck)(nil), "hw.trezor.messages.EvolutionStellarToken.EvolutionStellarTokenTxAck")
	proto.RegisterType((*EvolutionStellarTokenSignMessage)(nil), "hw.trezor.messages.EvolutionStellarToken.EvolutionStellarTokenSignMessage")
	proto.RegisterType((*EvolutionStellarTokenMessageSignature)(nil), "hw.trezor.messages.EvolutionStellarToken.EvolutionStellarTokenMessageSignature")
	proto.RegisterType((*EvolutionStellarTokenVerifyMessage)(nil), "hw.trezor.messages.EvolutionStellarToken.EvolutionStellarTokenVerifyMessage")
}

func init() { proto.RegisterFile("messages-EvolutionStellarToken.proto", fileDescriptor_cb33f46ba915f15c) }

var fileDescriptor_cb33f46ba915f15c = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0xb4, 0x49, 0x26, 0x0d, 0x1f, 0xa6, 0x55, 0x17, 0x0a, 0x34, 0x18, 0x21, 0xe5,
	0x00, 0x3e, 0x70, 0x43, 0xe2, 0xd2, 0x52, 0x44, 0x2b, 0x4a, 0x55, 0xdc, 0xa8, 0x57, 0x6b, 0x63,
	0x6f, 0xe3, 0x55, 0x9d, 0xdd, 0xe0, 0x5d, 0xb7, 0x0e, 0x7f, 0x82, 0x23, 0xff, 0x87, 0x5f, 0x86,
	0xf6, 0x2b, 0x71, 0x52, 0x54, 0x0e, 0xbd, 0x65, 0xde, 0xbc, 0x7d, 0xf3, 0x66, 0xf4, 0x62, 0xd8,
	0x99, 0x10, 0x21, 0xf0, 0x98, 0x88, 0x77, 0x44, 0x66, 0xa4, 0x20, 0xe5, 0x24, 0x9c, 0x16, 0x5c,
	0x72, 0x7f, 0x37, 0xbb, 0x09, 0x65, 0x41, 0x7e, 0xf2, 0x22, 0x74, 0x94, 0xd0, 0x51, 0x9e, 0x6d,
	0xcf, 0x5f, 0x25, 0x7c, 0x32, 0xe1, 0xcc, 0xbc, 0x09, 0x2e, 0x60, 0xeb, 0xb3, 0xa5, 0x7c, 0x21,
	0xf2, 0xac, 0x1c, 0xe5, 0x34, 0xf9, 0x4a, 0x66, 0xfe, 0x2e, 0x74, 0x70, 0x9a, 0x16, 0x44, 0x88,
	0x98, 0x21, 0xaf, 0xdf, 0x18, 0xf4, 0xa2, 0xb6, 0x05, 0x4e, 0xfd, 0x57, 0xb0, 0x29, 0x32, 0x7e,
	0x13, 0xa7, 0x54, 0x4c, 0x73, 0x3c, 0x43, 0x6b, 0x7d, 0x6f, 0xd0, 0x8e, 0xba, 0x0a, 0x3b, 0x34,
	0x50, 0x30, 0x82, 0xc7, 0x4e, 0x77, 0x21, 0xfa, 0x01, 0x9a, 0x8c, 0xa7, 0x04, 0x79, 0x7d, 0x6f,
	0xd0, 0x7d, 0xff, 0x26, 0xfc, 0x87, 0x5f, 0x6b, 0xee, 0xe8, 0xf0, 0x94, 0xa7, 0x64, 0x38, 0x9b,
	0x92, 0x48, 0x3f, 0xf1, 0x7d, 0x68, 0x56, 0xd3, 0x72, 0xa4, 0x47, 0x75, 0x22, 0xfd, 0x3b, 0x18,
	0x82, 0x5f, 0xf3, 0xbe, 0x6f, 0xdc, 0xdd, 0xdb, 0xf9, 0x77, 0x78, 0xe8, 0x54, 0x9d, 0xe4, 0x4b,
	0x00, 0xab, 0x70, 0x40, 0x99, 0x76, 0xbf, 0x19, 0xd5, 0x90, 0x5a, 0xff, 0x88, 0x54, 0xd6, 0x62,
	0x0d, 0x09, 0xfe, 0xac, 0xc1, 0x03, 0xa7, 0x79, 0x4e, 0xc7, 0x6c, 0x58, 0xdd, 0xed, 0x72, 0x0b,
	0xd6, 0x19, 0x67, 0x09, 0xd1, 0x52, 0x9b, 0x91, 0x29, 0xd4, 0x93, 0x31, 0x16, 0xf1, 0xb4, 0xa0,
	0x09, 0x41, 0x0d, 0xdd, 0x69, 0x8f, 0xb1, 0x38, 0x53, 0xb5, 0x6b, 0xe6, 0x74, 0x42, 0x25, 0x6a,
	0xce, 0x9b, 0x27, 0xaa, 0x56, 0x7a, 0x92, 0x2b, 0xeb, 0xeb, 0x46, 0x4f, 0x17, 0x06, 0x55, 0x86,
	0xbb, 0xda, 0xb0, 0x29, 0x14, 0x7a, 0x8d, 0xf3, 0x92, 0xa0, 0x0d, 0xc3, 0xd5, 0x85, 0xff, 0x16,
	0xfc, 0x14, 0x4b, 0x1c, 0x53, 0x46, 0x25, 0xc5, 0x79, 0x9c, 0x64, 0x25, 0xbb, 0x42, 0x2d, 0x4d,
	0x79, 0xa4, 0x3a, 0xc7, 0xa6, 0xf1, 0x49, 0xe1, 0xfe, 0x1e, 0x74, 0x35, 0x3b, 0x27, 0x6c, 0x2c,
	0x33, 0xd4, 0xee, 0x7b, 0x83, 0x5e, 0x04, 0x0a, 0x3a, 0xd1, 0x88, 0xff, 0x14, 0xda, 0x49, 0x86,
	0x29, 0x8b, 0x69, 0x8a, 0x3a, 0xba, 0xdb, 0xd2, 0xf5, 0x71, 0xea, 0xef, 0x40, 0x4b, 0x56, 0xb1,
	0x9c, 0x4d, 0x09, 0x02, 0xdd, 0xd9, 0x90, 0x95, 0xca, 0x41, 0xf0, 0xdb, 0x5b, 0x44, 0x6a, 0x58,
	0x45, 0xe4, 0x47, 0x49, 0x84, 0x5c, 0x1d, 0xe5, 0xdd, 0x1a, 0xb5, 0x07, 0x5d, 0x41, 0xc7, 0x0c,
	0xcb, 0xb2, 0x20, 0xf1, 0xb5, 0xbe, 0x68, 0x2f, 0x82, 0x39, 0x74, 0xb1, 0x4c, 0x28, 0xec, 0x61,
	0x17, 0x84, 0x68, 0x99, 0x20, 0xec, 0x71, 0x17, 0x84, 0xf3, 0x20, 0x84, 0xde, 0xc2, 0xd8, 0x7e,
	0x72, 0xe5, 0xbf, 0x00, 0xed, 0xc0, 0x5e, 0xc9, 0xe4, 0xa5, 0xa3, 0x10, 0x7d, 0x9e, 0xe0, 0x04,
	0x9e, 0xd4, 0xd3, 0xf0, 0xcd, 0x64, 0xff, 0xee, 0x48, 0x20, 0x68, 0xd9, 0xff, 0x88, 0x0d, 0x85,
	0x2b, 0x83, 0x0a, 0x90, 0x53, 0xb3, 0x4a, 0xe7, 0xce, 0xda, 0x7f, 0x83, 0xfb, 0x1c, 0x3a, 0xf3,
	0x3d, 0xac, 0xee, 0x02, 0x58, 0x89, 0x75, 0xe3, 0x56, 0xac, 0x7f, 0x79, 0xb0, 0xed, 0x46, 0x5f,
	0x90, 0x82, 0x5e, 0xce, 0xdc, 0x2a, 0xf7, 0x9b, 0x5b, 0xdb, 0xb5, 0xb1, 0xb4, 0xeb, 0x8a, 0xa3,
	0xe6, 0xaa, 0xa3, 0x83, 0x8f, 0xf0, 0x3a, 0xe1, 0x93, 0x50, 0x60, 0xc9, 0x45, 0x46, 0x73, 0x3c,
	0x12, 0xee, 0x03, 0x93, 0xd3, 0x91, 0xf9, 0xe2, 0x8d, 0xca, 0xcb, 0x83, 0xed, 0xa1, 0x06, 0xad,
	0x5b, 0xb7, 0xc2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xce, 0x81, 0xc8, 0x59, 0x05, 0x00,
	0x00,
}