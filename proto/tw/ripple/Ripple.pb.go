// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: Ripple.proto

package ripple

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	common "github.com/mamadeusia/WalletSrv/proto/tw/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// https://xrpl.org/currency-formats.html#token-amounts
type CurrencyAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Currency code
	// https://xrpl.org/currency-formats.html#currency-codes
	Currency string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	// String number
	// https://xrpl.org/currency-formats.html#string-numbers
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Account
	// https://xrpl.org/accounts.html
	Issuer string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
}

func (x *CurrencyAmount) Reset() {
	*x = CurrencyAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ripple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyAmount) ProtoMessage() {}

func (x *CurrencyAmount) ProtoReflect() protoreflect.Message {
	mi := &file_Ripple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyAmount.ProtoReflect.Descriptor instead.
func (*CurrencyAmount) Descriptor() ([]byte, []int) {
	return file_Ripple_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyAmount) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CurrencyAmount) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *CurrencyAmount) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

// https://xrpl.org/trustset.html
type OperationTrustSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LimitAmount *CurrencyAmount `protobuf:"bytes,1,opt,name=limit_amount,json=limitAmount,proto3" json:"limit_amount,omitempty"`
}

func (x *OperationTrustSet) Reset() {
	*x = OperationTrustSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ripple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationTrustSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationTrustSet) ProtoMessage() {}

func (x *OperationTrustSet) ProtoReflect() protoreflect.Message {
	mi := &file_Ripple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationTrustSet.ProtoReflect.Descriptor instead.
func (*OperationTrustSet) Descriptor() ([]byte, []int) {
	return file_Ripple_proto_rawDescGZIP(), []int{1}
}

func (x *OperationTrustSet) GetLimitAmount() *CurrencyAmount {
	if x != nil {
		return x.LimitAmount
	}
	return nil
}

// https://xrpl.org/payment.html
type OperationPayment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Transfer amount
	//
	// Types that are assignable to AmountOneof:
	//	*OperationPayment_Amount
	//	*OperationPayment_CurrencyAmount
	AmountOneof isOperationPayment_AmountOneof `protobuf_oneof:"amount_oneof"`
	// Target account
	Destination string `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	// A Destination Tag
	DestinationTag int64 `protobuf:"varint,4,opt,name=destination_tag,json=destinationTag,proto3" json:"destination_tag,omitempty"`
}

func (x *OperationPayment) Reset() {
	*x = OperationPayment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ripple_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationPayment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationPayment) ProtoMessage() {}

func (x *OperationPayment) ProtoReflect() protoreflect.Message {
	mi := &file_Ripple_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationPayment.ProtoReflect.Descriptor instead.
func (*OperationPayment) Descriptor() ([]byte, []int) {
	return file_Ripple_proto_rawDescGZIP(), []int{2}
}

func (m *OperationPayment) GetAmountOneof() isOperationPayment_AmountOneof {
	if m != nil {
		return m.AmountOneof
	}
	return nil
}

func (x *OperationPayment) GetAmount() int64 {
	if x, ok := x.GetAmountOneof().(*OperationPayment_Amount); ok {
		return x.Amount
	}
	return 0
}

func (x *OperationPayment) GetCurrencyAmount() *CurrencyAmount {
	if x, ok := x.GetAmountOneof().(*OperationPayment_CurrencyAmount); ok {
		return x.CurrencyAmount
	}
	return nil
}

func (x *OperationPayment) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *OperationPayment) GetDestinationTag() int64 {
	if x != nil {
		return x.DestinationTag
	}
	return 0
}

type isOperationPayment_AmountOneof interface {
	isOperationPayment_AmountOneof()
}

type OperationPayment_Amount struct {
	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3,oneof"`
}

type OperationPayment_CurrencyAmount struct {
	CurrencyAmount *CurrencyAmount `protobuf:"bytes,2,opt,name=currency_amount,json=currencyAmount,proto3,oneof"`
}

func (*OperationPayment_Amount) isOperationPayment_AmountOneof() {}

func (*OperationPayment_CurrencyAmount) isOperationPayment_AmountOneof() {}

// https://xrpl.org/nftokenburn.html
type OperationNFTokenBurn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hash256 NFTokenId
	NftokenId []byte `protobuf:"bytes,1,opt,name=nftoken_id,json=nftokenId,proto3" json:"nftoken_id,omitempty"`
}

func (x *OperationNFTokenBurn) Reset() {
	*x = OperationNFTokenBurn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ripple_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationNFTokenBurn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationNFTokenBurn) ProtoMessage() {}

func (x *OperationNFTokenBurn) ProtoReflect() protoreflect.Message {
	mi := &file_Ripple_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationNFTokenBurn.ProtoReflect.Descriptor instead.
func (*OperationNFTokenBurn) Descriptor() ([]byte, []int) {
	return file_Ripple_proto_rawDescGZIP(), []int{3}
}

func (x *OperationNFTokenBurn) GetNftokenId() []byte {
	if x != nil {
		return x.NftokenId
	}
	return nil
}

// https://xrpl.org/nftokencreateoffer.html
type OperationNFTokenCreateOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hash256 NFTokenId
	NftokenId []byte `protobuf:"bytes,1,opt,name=nftoken_id,json=nftokenId,proto3" json:"nftoken_id,omitempty"`
	// Destination account
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *OperationNFTokenCreateOffer) Reset() {
	*x = OperationNFTokenCreateOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ripple_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationNFTokenCreateOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationNFTokenCreateOffer) ProtoMessage() {}

func (x *OperationNFTokenCreateOffer) ProtoReflect() protoreflect.Message {
	mi := &file_Ripple_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationNFTokenCreateOffer.ProtoReflect.Descriptor instead.
func (*OperationNFTokenCreateOffer) Descriptor() ([]byte, []int) {
	return file_Ripple_proto_rawDescGZIP(), []int{4}
}

func (x *OperationNFTokenCreateOffer) GetNftokenId() []byte {
	if x != nil {
		return x.NftokenId
	}
	return nil
}

func (x *OperationNFTokenCreateOffer) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

// https://xrpl.org/nftokenacceptoffer.html
type OperationNFTokenAcceptOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hash256 NFTokenOffer
	SellOffer []byte `protobuf:"bytes,1,opt,name=sell_offer,json=sellOffer,proto3" json:"sell_offer,omitempty"`
}

func (x *OperationNFTokenAcceptOffer) Reset() {
	*x = OperationNFTokenAcceptOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ripple_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationNFTokenAcceptOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationNFTokenAcceptOffer) ProtoMessage() {}

func (x *OperationNFTokenAcceptOffer) ProtoReflect() protoreflect.Message {
	mi := &file_Ripple_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationNFTokenAcceptOffer.ProtoReflect.Descriptor instead.
func (*OperationNFTokenAcceptOffer) Descriptor() ([]byte, []int) {
	return file_Ripple_proto_rawDescGZIP(), []int{5}
}

func (x *OperationNFTokenAcceptOffer) GetSellOffer() []byte {
	if x != nil {
		return x.SellOffer
	}
	return nil
}

// https://xrpl.org/nftokencanceloffer.html
type OperationNFTokenCancelOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Vector256 NFTokenOffers
	TokenOffers [][]byte `protobuf:"bytes,1,rep,name=token_offers,json=tokenOffers,proto3" json:"token_offers,omitempty"`
}

func (x *OperationNFTokenCancelOffer) Reset() {
	*x = OperationNFTokenCancelOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ripple_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationNFTokenCancelOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationNFTokenCancelOffer) ProtoMessage() {}

func (x *OperationNFTokenCancelOffer) ProtoReflect() protoreflect.Message {
	mi := &file_Ripple_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationNFTokenCancelOffer.ProtoReflect.Descriptor instead.
func (*OperationNFTokenCancelOffer) Descriptor() ([]byte, []int) {
	return file_Ripple_proto_rawDescGZIP(), []int{6}
}

func (x *OperationNFTokenCancelOffer) GetTokenOffers() [][]byte {
	if x != nil {
		return x.TokenOffers
	}
	return nil
}

// Input data necessary to create a signed transaction.
type SigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Transfer fee
	Fee int64 `protobuf:"varint,1,opt,name=fee,proto3" json:"fee,omitempty"`
	// Account sequence number
	Sequence int32 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Ledger sequence number
	LastLedgerSequence int32 `protobuf:"varint,3,opt,name=last_ledger_sequence,json=lastLedgerSequence,proto3" json:"last_ledger_sequence,omitempty"`
	// Source account
	Account string `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	// Transaction flags, optional
	Flags int64 `protobuf:"varint,5,opt,name=flags,proto3" json:"flags,omitempty"`
	// The secret private key used for signing (32 bytes).
	PrivateKey []byte `protobuf:"bytes,6,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Types that are assignable to OperationOneof:
	//	*SigningInput_OpTrustSet
	//	*SigningInput_OpPayment
	//	*SigningInput_OpNftokenBurn
	//	*SigningInput_OpNftokenCreateOffer
	//	*SigningInput_OpNftokenAcceptOffer
	//	*SigningInput_OpNftokenCancelOffer
	OperationOneof isSigningInput_OperationOneof `protobuf_oneof:"operation_oneof"`
}

func (x *SigningInput) Reset() {
	*x = SigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ripple_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningInput) ProtoMessage() {}

func (x *SigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_Ripple_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigningInput.ProtoReflect.Descriptor instead.
func (*SigningInput) Descriptor() ([]byte, []int) {
	return file_Ripple_proto_rawDescGZIP(), []int{7}
}

func (x *SigningInput) GetFee() int64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *SigningInput) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *SigningInput) GetLastLedgerSequence() int32 {
	if x != nil {
		return x.LastLedgerSequence
	}
	return 0
}

func (x *SigningInput) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SigningInput) GetFlags() int64 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *SigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (m *SigningInput) GetOperationOneof() isSigningInput_OperationOneof {
	if m != nil {
		return m.OperationOneof
	}
	return nil
}

func (x *SigningInput) GetOpTrustSet() *OperationTrustSet {
	if x, ok := x.GetOperationOneof().(*SigningInput_OpTrustSet); ok {
		return x.OpTrustSet
	}
	return nil
}

func (x *SigningInput) GetOpPayment() *OperationPayment {
	if x, ok := x.GetOperationOneof().(*SigningInput_OpPayment); ok {
		return x.OpPayment
	}
	return nil
}

func (x *SigningInput) GetOpNftokenBurn() *OperationNFTokenBurn {
	if x, ok := x.GetOperationOneof().(*SigningInput_OpNftokenBurn); ok {
		return x.OpNftokenBurn
	}
	return nil
}

func (x *SigningInput) GetOpNftokenCreateOffer() *OperationNFTokenCreateOffer {
	if x, ok := x.GetOperationOneof().(*SigningInput_OpNftokenCreateOffer); ok {
		return x.OpNftokenCreateOffer
	}
	return nil
}

func (x *SigningInput) GetOpNftokenAcceptOffer() *OperationNFTokenAcceptOffer {
	if x, ok := x.GetOperationOneof().(*SigningInput_OpNftokenAcceptOffer); ok {
		return x.OpNftokenAcceptOffer
	}
	return nil
}

func (x *SigningInput) GetOpNftokenCancelOffer() *OperationNFTokenCancelOffer {
	if x, ok := x.GetOperationOneof().(*SigningInput_OpNftokenCancelOffer); ok {
		return x.OpNftokenCancelOffer
	}
	return nil
}

type isSigningInput_OperationOneof interface {
	isSigningInput_OperationOneof()
}

type SigningInput_OpTrustSet struct {
	OpTrustSet *OperationTrustSet `protobuf:"bytes,7,opt,name=op_trust_set,json=opTrustSet,proto3,oneof"`
}

type SigningInput_OpPayment struct {
	OpPayment *OperationPayment `protobuf:"bytes,8,opt,name=op_payment,json=opPayment,proto3,oneof"`
}

type SigningInput_OpNftokenBurn struct {
	OpNftokenBurn *OperationNFTokenBurn `protobuf:"bytes,9,opt,name=op_nftoken_burn,json=opNftokenBurn,proto3,oneof"`
}

type SigningInput_OpNftokenCreateOffer struct {
	OpNftokenCreateOffer *OperationNFTokenCreateOffer `protobuf:"bytes,10,opt,name=op_nftoken_create_offer,json=opNftokenCreateOffer,proto3,oneof"`
}

type SigningInput_OpNftokenAcceptOffer struct {
	OpNftokenAcceptOffer *OperationNFTokenAcceptOffer `protobuf:"bytes,11,opt,name=op_nftoken_accept_offer,json=opNftokenAcceptOffer,proto3,oneof"`
}

type SigningInput_OpNftokenCancelOffer struct {
	OpNftokenCancelOffer *OperationNFTokenCancelOffer `protobuf:"bytes,12,opt,name=op_nftoken_cancel_offer,json=opNftokenCancelOffer,proto3,oneof"`
}

func (*SigningInput_OpTrustSet) isSigningInput_OperationOneof() {}

func (*SigningInput_OpPayment) isSigningInput_OperationOneof() {}

func (*SigningInput_OpNftokenBurn) isSigningInput_OperationOneof() {}

func (*SigningInput_OpNftokenCreateOffer) isSigningInput_OperationOneof() {}

func (*SigningInput_OpNftokenAcceptOffer) isSigningInput_OperationOneof() {}

func (*SigningInput_OpNftokenCancelOffer) isSigningInput_OperationOneof() {}

// Result containing the signed and encoded transaction.
type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Encoded transaction
	Encoded []byte `protobuf:"bytes,1,opt,name=encoded,proto3" json:"encoded,omitempty"`
	// Optional error
	Error common.SigningError `protobuf:"varint,2,opt,name=error,proto3,enum=TW.Common.Proto.SigningError" json:"error,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ripple_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_Ripple_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigningOutput.ProtoReflect.Descriptor instead.
func (*SigningOutput) Descriptor() ([]byte, []int) {
	return file_Ripple_proto_rawDescGZIP(), []int{8}
}

func (x *SigningOutput) GetEncoded() []byte {
	if x != nil {
		return x.Encoded
	}
	return nil
}

func (x *SigningOutput) GetError() common.SigningError {
	if x != nil {
		return x.Error
	}
	return common.SigningError(0)
}

var File_Ripple_proto protoreflect.FileDescriptor

var file_Ripple_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x52, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x54, 0x57, 0x2e, 0x52, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a,
	0x0e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x22, 0x57, 0x0a, 0x11, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x53, 0x65, 0x74, 0x12, 0x42,
	0x0a, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x54, 0x57, 0x2e, 0x52, 0x69, 0x70, 0x70, 0x6c, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0xd3, 0x01, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x4a, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x54, 0x57, 0x2e,
	0x52, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x42, 0x0e, 0x0a, 0x0c, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0x35, 0x0a, 0x14, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x46, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x72, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x66, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6e, 0x66, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x22,
	0x5e, 0x0a, 0x1b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x46, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x6e, 0x66, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x6e, 0x66, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x3c, 0x0a, 0x1b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x46, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x22, 0x40, 0x0a,
	0x1b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x46, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x22,
	0xe4, 0x05, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x66,
	0x65, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x30,
	0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6c, 0x61,
	0x73, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x46, 0x0a, 0x0c, 0x6f, 0x70, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x73, 0x65,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x54, 0x57, 0x2e, 0x52, 0x69, 0x70,
	0x70, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x6f,
	0x70, 0x54, 0x72, 0x75, 0x73, 0x74, 0x53, 0x65, 0x74, 0x12, 0x42, 0x0a, 0x0a, 0x6f, 0x70, 0x5f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x54, 0x57, 0x2e, 0x52, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x09, 0x6f, 0x70, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4f, 0x0a,
	0x0f, 0x6f, 0x70, 0x5f, 0x6e, 0x66, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x75, 0x72, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x54, 0x57, 0x2e, 0x52, 0x69, 0x70, 0x70,
	0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x46, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x72, 0x6e, 0x48, 0x00, 0x52,
	0x0d, 0x6f, 0x70, 0x4e, 0x66, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x72, 0x6e, 0x12, 0x65,
	0x0a, 0x17, 0x6f, 0x70, 0x5f, 0x6e, 0x66, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x54, 0x57, 0x2e, 0x52, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x46, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x14, 0x6f, 0x70, 0x4e, 0x66, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x65, 0x0a, 0x17, 0x6f, 0x70, 0x5f, 0x6e, 0x66, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x54, 0x57, 0x2e, 0x52, 0x69, 0x70, 0x70,
	0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x46, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52, 0x14, 0x6f, 0x70, 0x4e, 0x66, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x65, 0x0a, 0x17,
	0x6f, 0x70, 0x5f, 0x6e, 0x66, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x54, 0x57, 0x2e, 0x52, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x46, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52, 0x14, 0x6f,
	0x70, 0x4e, 0x66, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x42, 0x11, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0x5e, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e,
	0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x54, 0x57, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x17, 0x0a, 0x15, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6a, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Ripple_proto_rawDescOnce sync.Once
	file_Ripple_proto_rawDescData = file_Ripple_proto_rawDesc
)

func file_Ripple_proto_rawDescGZIP() []byte {
	file_Ripple_proto_rawDescOnce.Do(func() {
		file_Ripple_proto_rawDescData = protoimpl.X.CompressGZIP(file_Ripple_proto_rawDescData)
	})
	return file_Ripple_proto_rawDescData
}

var file_Ripple_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_Ripple_proto_goTypes = []interface{}{
	(*CurrencyAmount)(nil),              // 0: TW.Ripple.Proto.CurrencyAmount
	(*OperationTrustSet)(nil),           // 1: TW.Ripple.Proto.OperationTrustSet
	(*OperationPayment)(nil),            // 2: TW.Ripple.Proto.OperationPayment
	(*OperationNFTokenBurn)(nil),        // 3: TW.Ripple.Proto.OperationNFTokenBurn
	(*OperationNFTokenCreateOffer)(nil), // 4: TW.Ripple.Proto.OperationNFTokenCreateOffer
	(*OperationNFTokenAcceptOffer)(nil), // 5: TW.Ripple.Proto.OperationNFTokenAcceptOffer
	(*OperationNFTokenCancelOffer)(nil), // 6: TW.Ripple.Proto.OperationNFTokenCancelOffer
	(*SigningInput)(nil),                // 7: TW.Ripple.Proto.SigningInput
	(*SigningOutput)(nil),               // 8: TW.Ripple.Proto.SigningOutput
	(common.SigningError)(0),            // 9: TW.Common.Proto.SigningError
}
var file_Ripple_proto_depIdxs = []int32{
	0, // 0: TW.Ripple.Proto.OperationTrustSet.limit_amount:type_name -> TW.Ripple.Proto.CurrencyAmount
	0, // 1: TW.Ripple.Proto.OperationPayment.currency_amount:type_name -> TW.Ripple.Proto.CurrencyAmount
	1, // 2: TW.Ripple.Proto.SigningInput.op_trust_set:type_name -> TW.Ripple.Proto.OperationTrustSet
	2, // 3: TW.Ripple.Proto.SigningInput.op_payment:type_name -> TW.Ripple.Proto.OperationPayment
	3, // 4: TW.Ripple.Proto.SigningInput.op_nftoken_burn:type_name -> TW.Ripple.Proto.OperationNFTokenBurn
	4, // 5: TW.Ripple.Proto.SigningInput.op_nftoken_create_offer:type_name -> TW.Ripple.Proto.OperationNFTokenCreateOffer
	5, // 6: TW.Ripple.Proto.SigningInput.op_nftoken_accept_offer:type_name -> TW.Ripple.Proto.OperationNFTokenAcceptOffer
	6, // 7: TW.Ripple.Proto.SigningInput.op_nftoken_cancel_offer:type_name -> TW.Ripple.Proto.OperationNFTokenCancelOffer
	9, // 8: TW.Ripple.Proto.SigningOutput.error:type_name -> TW.Common.Proto.SigningError
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_Ripple_proto_init() }
func file_Ripple_proto_init() {
	if File_Ripple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Ripple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyAmount); i {
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
		file_Ripple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationTrustSet); i {
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
		file_Ripple_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationPayment); i {
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
		file_Ripple_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationNFTokenBurn); i {
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
		file_Ripple_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationNFTokenCreateOffer); i {
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
		file_Ripple_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationNFTokenAcceptOffer); i {
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
		file_Ripple_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationNFTokenCancelOffer); i {
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
		file_Ripple_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigningInput); i {
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
		file_Ripple_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigningOutput); i {
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
	file_Ripple_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*OperationPayment_Amount)(nil),
		(*OperationPayment_CurrencyAmount)(nil),
	}
	file_Ripple_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*SigningInput_OpTrustSet)(nil),
		(*SigningInput_OpPayment)(nil),
		(*SigningInput_OpNftokenBurn)(nil),
		(*SigningInput_OpNftokenCreateOffer)(nil),
		(*SigningInput_OpNftokenAcceptOffer)(nil),
		(*SigningInput_OpNftokenCancelOffer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Ripple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Ripple_proto_goTypes,
		DependencyIndexes: file_Ripple_proto_depIdxs,
		MessageInfos:      file_Ripple_proto_msgTypes,
	}.Build()
	File_Ripple_proto = out.File
	file_Ripple_proto_rawDesc = nil
	file_Ripple_proto_goTypes = nil
	file_Ripple_proto_depIdxs = nil
}
