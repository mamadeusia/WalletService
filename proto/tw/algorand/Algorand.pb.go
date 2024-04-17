// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: Algorand.proto

package algorand

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

// Simple transfer message, transfer an amount to an address
type Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Destination address (string)
	ToAddress string `protobuf:"bytes,1,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// Amount
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Transfer) Reset() {
	*x = Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Algorand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_Algorand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transfer.ProtoReflect.Descriptor instead.
func (*Transfer) Descriptor() ([]byte, []int) {
	return file_Algorand_proto_rawDescGZIP(), []int{0}
}

func (x *Transfer) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *Transfer) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// Asset Transfer message, with assetID
type AssetTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Destination address (string)
	ToAddress string `protobuf:"bytes,1,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// Amount
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// ID of the asset being transferred
	AssetId uint64 `protobuf:"varint,3,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
}

func (x *AssetTransfer) Reset() {
	*x = AssetTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Algorand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetTransfer) ProtoMessage() {}

func (x *AssetTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_Algorand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetTransfer.ProtoReflect.Descriptor instead.
func (*AssetTransfer) Descriptor() ([]byte, []int) {
	return file_Algorand_proto_rawDescGZIP(), []int{1}
}

func (x *AssetTransfer) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *AssetTransfer) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *AssetTransfer) GetAssetId() uint64 {
	if x != nil {
		return x.AssetId
	}
	return 0
}

// Opt-in message for an asset
type AssetOptIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the asset
	AssetId uint64 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
}

func (x *AssetOptIn) Reset() {
	*x = AssetOptIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Algorand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetOptIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetOptIn) ProtoMessage() {}

func (x *AssetOptIn) ProtoReflect() protoreflect.Message {
	mi := &file_Algorand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetOptIn.ProtoReflect.Descriptor instead.
func (*AssetOptIn) Descriptor() ([]byte, []int) {
	return file_Algorand_proto_rawDescGZIP(), []int{2}
}

func (x *AssetOptIn) GetAssetId() uint64 {
	if x != nil {
		return x.AssetId
	}
	return 0
}

// Input data necessary to create a signed transaction.
type SigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// network / chain id
	GenesisId string `protobuf:"bytes,1,opt,name=genesis_id,json=genesisId,proto3" json:"genesis_id,omitempty"`
	// network / chain hash
	GenesisHash []byte `protobuf:"bytes,2,opt,name=genesis_hash,json=genesisHash,proto3" json:"genesis_hash,omitempty"`
	// binary note data
	Note []byte `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
	// The secret private key used for signing (32 bytes).
	PrivateKey []byte `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// network / first round
	FirstRound uint64 `protobuf:"varint,5,opt,name=first_round,json=firstRound,proto3" json:"first_round,omitempty"`
	// network / last round
	LastRound uint64 `protobuf:"varint,6,opt,name=last_round,json=lastRound,proto3" json:"last_round,omitempty"`
	// fee amount
	Fee uint64 `protobuf:"varint,7,opt,name=fee,proto3" json:"fee,omitempty"`
	// message payload
	//
	// Types that are assignable to MessageOneof:
	//	*SigningInput_Transfer
	//	*SigningInput_AssetTransfer
	//	*SigningInput_AssetOptIn
	MessageOneof isSigningInput_MessageOneof `protobuf_oneof:"message_oneof"`
}

func (x *SigningInput) Reset() {
	*x = SigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Algorand_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningInput) ProtoMessage() {}

func (x *SigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_Algorand_proto_msgTypes[3]
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
	return file_Algorand_proto_rawDescGZIP(), []int{3}
}

func (x *SigningInput) GetGenesisId() string {
	if x != nil {
		return x.GenesisId
	}
	return ""
}

func (x *SigningInput) GetGenesisHash() []byte {
	if x != nil {
		return x.GenesisHash
	}
	return nil
}

func (x *SigningInput) GetNote() []byte {
	if x != nil {
		return x.Note
	}
	return nil
}

func (x *SigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *SigningInput) GetFirstRound() uint64 {
	if x != nil {
		return x.FirstRound
	}
	return 0
}

func (x *SigningInput) GetLastRound() uint64 {
	if x != nil {
		return x.LastRound
	}
	return 0
}

func (x *SigningInput) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (m *SigningInput) GetMessageOneof() isSigningInput_MessageOneof {
	if m != nil {
		return m.MessageOneof
	}
	return nil
}

func (x *SigningInput) GetTransfer() *Transfer {
	if x, ok := x.GetMessageOneof().(*SigningInput_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (x *SigningInput) GetAssetTransfer() *AssetTransfer {
	if x, ok := x.GetMessageOneof().(*SigningInput_AssetTransfer); ok {
		return x.AssetTransfer
	}
	return nil
}

func (x *SigningInput) GetAssetOptIn() *AssetOptIn {
	if x, ok := x.GetMessageOneof().(*SigningInput_AssetOptIn); ok {
		return x.AssetOptIn
	}
	return nil
}

type isSigningInput_MessageOneof interface {
	isSigningInput_MessageOneof()
}

type SigningInput_Transfer struct {
	Transfer *Transfer `protobuf:"bytes,10,opt,name=transfer,proto3,oneof"`
}

type SigningInput_AssetTransfer struct {
	AssetTransfer *AssetTransfer `protobuf:"bytes,11,opt,name=asset_transfer,json=assetTransfer,proto3,oneof"`
}

type SigningInput_AssetOptIn struct {
	AssetOptIn *AssetOptIn `protobuf:"bytes,12,opt,name=asset_opt_in,json=assetOptIn,proto3,oneof"`
}

func (*SigningInput_Transfer) isSigningInput_MessageOneof() {}

func (*SigningInput_AssetTransfer) isSigningInput_MessageOneof() {}

func (*SigningInput_AssetOptIn) isSigningInput_MessageOneof() {}

// Result containing the signed and encoded transaction.
type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Signed and encoded transaction bytes.
	Encoded []byte `protobuf:"bytes,1,opt,name=encoded,proto3" json:"encoded,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Algorand_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_Algorand_proto_msgTypes[4]
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
	return file_Algorand_proto_rawDescGZIP(), []int{4}
}

func (x *SigningOutput) GetEncoded() []byte {
	if x != nil {
		return x.Encoded
	}
	return nil
}

var File_Algorand_proto protoreflect.FileDescriptor

var file_Algorand_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x54, 0x57, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x61, 0x0a, 0x0d, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x0a, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x4f, 0x70, 0x74, 0x49, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x49, 0x64, 0x22, 0xb1, 0x03, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x39, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x54, 0x57, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x54, 0x57, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x12, 0x41, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x5f,
	0x69, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x54, 0x57, 0x2e, 0x41, 0x6c,
	0x67, 0x6f, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x4f, 0x70, 0x74, 0x49, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x4f, 0x70, 0x74, 0x49, 0x6e, 0x42, 0x0f, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e,
	0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x42, 0x17, 0x0a, 0x15, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x6a, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Algorand_proto_rawDescOnce sync.Once
	file_Algorand_proto_rawDescData = file_Algorand_proto_rawDesc
)

func file_Algorand_proto_rawDescGZIP() []byte {
	file_Algorand_proto_rawDescOnce.Do(func() {
		file_Algorand_proto_rawDescData = protoimpl.X.CompressGZIP(file_Algorand_proto_rawDescData)
	})
	return file_Algorand_proto_rawDescData
}

var file_Algorand_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Algorand_proto_goTypes = []interface{}{
	(*Transfer)(nil),      // 0: TW.Algorand.Proto.Transfer
	(*AssetTransfer)(nil), // 1: TW.Algorand.Proto.AssetTransfer
	(*AssetOptIn)(nil),    // 2: TW.Algorand.Proto.AssetOptIn
	(*SigningInput)(nil),  // 3: TW.Algorand.Proto.SigningInput
	(*SigningOutput)(nil), // 4: TW.Algorand.Proto.SigningOutput
}
var file_Algorand_proto_depIdxs = []int32{
	0, // 0: TW.Algorand.Proto.SigningInput.transfer:type_name -> TW.Algorand.Proto.Transfer
	1, // 1: TW.Algorand.Proto.SigningInput.asset_transfer:type_name -> TW.Algorand.Proto.AssetTransfer
	2, // 2: TW.Algorand.Proto.SigningInput.asset_opt_in:type_name -> TW.Algorand.Proto.AssetOptIn
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Algorand_proto_init() }
func file_Algorand_proto_init() {
	if File_Algorand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Algorand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transfer); i {
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
		file_Algorand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetTransfer); i {
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
		file_Algorand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetOptIn); i {
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
		file_Algorand_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_Algorand_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	file_Algorand_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SigningInput_Transfer)(nil),
		(*SigningInput_AssetTransfer)(nil),
		(*SigningInput_AssetOptIn)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Algorand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Algorand_proto_goTypes,
		DependencyIndexes: file_Algorand_proto_depIdxs,
		MessageInfos:      file_Algorand_proto_msgTypes,
	}.Build()
	File_Algorand_proto = out.File
	file_Algorand_proto_rawDesc = nil
	file_Algorand_proto_goTypes = nil
	file_Algorand_proto_depIdxs = nil
}
