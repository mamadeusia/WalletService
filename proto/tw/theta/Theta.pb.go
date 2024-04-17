// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: Theta.proto

package theta

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

/// Input data necessary to create a signed transaction
type SigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// Chain ID string, mainnet, testnet and privatenet
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	/// Recipient address
	ToAddress string `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	/// Theta token amount to send in wei (uint256, serialized little endian)
	ThetaAmount []byte `protobuf:"bytes,3,opt,name=theta_amount,json=thetaAmount,proto3" json:"theta_amount,omitempty"`
	/// TFuel token amount to send in wei (uint256, serialized little endian)
	TfuelAmount []byte `protobuf:"bytes,4,opt,name=tfuel_amount,json=tfuelAmount,proto3" json:"tfuel_amount,omitempty"`
	/// Sequence number of the transaction for the sender address
	Sequence uint64 `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
	/// Fee amount in TFuel wei for the transaction (uint256, serialized little endian)
	Fee []byte `protobuf:"bytes,6,opt,name=fee,proto3" json:"fee,omitempty"`
	/// The secret private key used for signing (32 bytes).
	PrivateKey []byte `protobuf:"bytes,7,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
}

func (x *SigningInput) Reset() {
	*x = SigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Theta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningInput) ProtoMessage() {}

func (x *SigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_Theta_proto_msgTypes[0]
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
	return file_Theta_proto_rawDescGZIP(), []int{0}
}

func (x *SigningInput) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *SigningInput) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *SigningInput) GetThetaAmount() []byte {
	if x != nil {
		return x.ThetaAmount
	}
	return nil
}

func (x *SigningInput) GetTfuelAmount() []byte {
	if x != nil {
		return x.TfuelAmount
	}
	return nil
}

func (x *SigningInput) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *SigningInput) GetFee() []byte {
	if x != nil {
		return x.Fee
	}
	return nil
}

func (x *SigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

// Result containing the signed and encoded transaction.
type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// Signed and encoded transaction bytes
	Encoded []byte `protobuf:"bytes,1,opt,name=encoded,proto3" json:"encoded,omitempty"`
	/// Signature
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Theta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_Theta_proto_msgTypes[1]
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
	return file_Theta_proto_rawDescGZIP(), []int{1}
}

func (x *SigningOutput) GetEncoded() []byte {
	if x != nil {
		return x.Encoded
	}
	return nil
}

func (x *SigningOutput) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_Theta_proto protoreflect.FileDescriptor

var file_Theta_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x54, 0x68, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x54,
	0x57, 0x2e, 0x54, 0x68, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01,
	0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74,
	0x66, 0x75, 0x65, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x74, 0x66, 0x75, 0x65, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x47, 0x0a,
	0x0d, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x17, 0x0a, 0x15, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6a, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Theta_proto_rawDescOnce sync.Once
	file_Theta_proto_rawDescData = file_Theta_proto_rawDesc
)

func file_Theta_proto_rawDescGZIP() []byte {
	file_Theta_proto_rawDescOnce.Do(func() {
		file_Theta_proto_rawDescData = protoimpl.X.CompressGZIP(file_Theta_proto_rawDescData)
	})
	return file_Theta_proto_rawDescData
}

var file_Theta_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Theta_proto_goTypes = []interface{}{
	(*SigningInput)(nil),  // 0: TW.Theta.Proto.SigningInput
	(*SigningOutput)(nil), // 1: TW.Theta.Proto.SigningOutput
}
var file_Theta_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Theta_proto_init() }
func file_Theta_proto_init() {
	if File_Theta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Theta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Theta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Theta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Theta_proto_goTypes,
		DependencyIndexes: file_Theta_proto_depIdxs,
		MessageInfos:      file_Theta_proto_msgTypes,
	}.Build()
	File_Theta_proto = out.File
	file_Theta_proto_rawDesc = nil
	file_Theta_proto_goTypes = nil
	file_Theta_proto_depIdxs = nil
}