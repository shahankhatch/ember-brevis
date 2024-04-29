// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.23.4
// source: sdk/prover.proto

package sdkproto

import (
	commonproto "github.com/brevis-network/brevis-sdk/sdk/proto/commonproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ErrCode int32

const (
	ErrCode_ERROR_UNDEFINED            ErrCode = 0
	ErrCode_ERROR_DEFAULT              ErrCode = 1
	ErrCode_ERROR_INVALID_INPUT        ErrCode = 2
	ErrCode_ERROR_INVALID_CUSTOM_INPUT ErrCode = 3
	ErrCode_ERROR_FAILED_TO_PROVE      ErrCode = 4
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0: "ERROR_UNDEFINED",
		1: "ERROR_DEFAULT",
		2: "ERROR_INVALID_INPUT",
		3: "ERROR_INVALID_CUSTOM_INPUT",
		4: "ERROR_FAILED_TO_PROVE",
	}
	ErrCode_value = map[string]int32{
		"ERROR_UNDEFINED":            0,
		"ERROR_DEFAULT":              1,
		"ERROR_INVALID_INPUT":        2,
		"ERROR_INVALID_CUSTOM_INPUT": 3,
		"ERROR_FAILED_TO_PROVE":      4,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_sdk_prover_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_sdk_prover_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_sdk_prover_proto_rawDescGZIP(), []int{0}
}

type ProveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receipts     []*IndexedReceipt     `protobuf:"bytes,1,rep,name=receipts,proto3" json:"receipts,omitempty"`
	Storages     []*IndexedStorage     `protobuf:"bytes,2,rep,name=storages,proto3" json:"storages,omitempty"`
	Transactions []*IndexedTransaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
	CustomInput  *CustomInput          `protobuf:"bytes,4,opt,name=custom_input,json=customInput,proto3" json:"custom_input,omitempty"`
}

func (x *ProveRequest) Reset() {
	*x = ProveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_prover_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProveRequest) ProtoMessage() {}

func (x *ProveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_prover_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProveRequest.ProtoReflect.Descriptor instead.
func (*ProveRequest) Descriptor() ([]byte, []int) {
	return file_sdk_prover_proto_rawDescGZIP(), []int{0}
}

func (x *ProveRequest) GetReceipts() []*IndexedReceipt {
	if x != nil {
		return x.Receipts
	}
	return nil
}

func (x *ProveRequest) GetStorages() []*IndexedStorage {
	if x != nil {
		return x.Storages
	}
	return nil
}

func (x *ProveRequest) GetTransactions() []*IndexedTransaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *ProveRequest) GetCustomInput() *CustomInput {
	if x != nil {
		return x.CustomInput
	}
	return nil
}

type ProveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err         *Err                        `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	Proof       string                      `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
	CircuitInfo *commonproto.AppCircuitInfo `protobuf:"bytes,3,opt,name=circuit_info,json=circuitInfo,proto3" json:"circuit_info,omitempty"`
}

func (x *ProveResponse) Reset() {
	*x = ProveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_prover_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProveResponse) ProtoMessage() {}

func (x *ProveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_prover_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProveResponse.ProtoReflect.Descriptor instead.
func (*ProveResponse) Descriptor() ([]byte, []int) {
	return file_sdk_prover_proto_rawDescGZIP(), []int{1}
}

func (x *ProveResponse) GetErr() *Err {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *ProveResponse) GetProof() string {
	if x != nil {
		return x.Proof
	}
	return ""
}

func (x *ProveResponse) GetCircuitInfo() *commonproto.AppCircuitInfo {
	if x != nil {
		return x.CircuitInfo
	}
	return nil
}

type ProveAsyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err         *Err                        `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	ProofId     string                      `protobuf:"bytes,2,opt,name=proof_id,json=proofId,proto3" json:"proof_id,omitempty"`
	CircuitInfo *commonproto.AppCircuitInfo `protobuf:"bytes,3,opt,name=circuit_info,json=circuitInfo,proto3" json:"circuit_info,omitempty"`
}

func (x *ProveAsyncResponse) Reset() {
	*x = ProveAsyncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_prover_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProveAsyncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProveAsyncResponse) ProtoMessage() {}

func (x *ProveAsyncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_prover_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProveAsyncResponse.ProtoReflect.Descriptor instead.
func (*ProveAsyncResponse) Descriptor() ([]byte, []int) {
	return file_sdk_prover_proto_rawDescGZIP(), []int{2}
}

func (x *ProveAsyncResponse) GetErr() *Err {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *ProveAsyncResponse) GetProofId() string {
	if x != nil {
		return x.ProofId
	}
	return ""
}

func (x *ProveAsyncResponse) GetCircuitInfo() *commonproto.AppCircuitInfo {
	if x != nil {
		return x.CircuitInfo
	}
	return nil
}

type GetProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProofId string `protobuf:"bytes,1,opt,name=proof_id,json=proofId,proto3" json:"proof_id,omitempty"`
}

func (x *GetProofRequest) Reset() {
	*x = GetProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_prover_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProofRequest) ProtoMessage() {}

func (x *GetProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_prover_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProofRequest.ProtoReflect.Descriptor instead.
func (*GetProofRequest) Descriptor() ([]byte, []int) {
	return file_sdk_prover_proto_rawDescGZIP(), []int{3}
}

func (x *GetProofRequest) GetProofId() string {
	if x != nil {
		return x.ProofId
	}
	return ""
}

type GetProofResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err *Err `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	// proof is an empty string until proving is finished
	Proof string `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *GetProofResponse) Reset() {
	*x = GetProofResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_prover_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProofResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProofResponse) ProtoMessage() {}

func (x *GetProofResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_prover_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProofResponse.ProtoReflect.Descriptor instead.
func (*GetProofResponse) Descriptor() ([]byte, []int) {
	return file_sdk_prover_proto_rawDescGZIP(), []int{4}
}

func (x *GetProofResponse) GetErr() *Err {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *GetProofResponse) GetProof() string {
	if x != nil {
		return x.Proof
	}
	return ""
}

type Err struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ErrCode `protobuf:"varint,1,opt,name=code,proto3,enum=sdk.ErrCode" json:"code,omitempty"`
	Msg  string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Err) Reset() {
	*x = Err{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_prover_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Err) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Err) ProtoMessage() {}

func (x *Err) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_prover_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Err.ProtoReflect.Descriptor instead.
func (*Err) Descriptor() ([]byte, []int) {
	return file_sdk_prover_proto_rawDescGZIP(), []int{5}
}

func (x *Err) GetCode() ErrCode {
	if x != nil {
		return x.Code
	}
	return ErrCode_ERROR_UNDEFINED
}

func (x *Err) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_sdk_prover_proto protoreflect.FileDescriptor

var file_sdk_prover_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x73, 0x64, 0x6b, 0x1a, 0x0f, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x65, 0x64, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x33, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x7c, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x45, 0x72, 0x72, 0x52, 0x03,
	0x65, 0x72, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x39, 0x0a, 0x0c, 0x63, 0x69, 0x72,
	0x63, 0x75, 0x69, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x69, 0x72, 0x63,
	0x75, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x41, 0x73,
	0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x65,
	0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x45,
	0x72, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x49, 0x64, 0x12, 0x39, 0x0a, 0x0c, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0b, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2c, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x45, 0x72, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x22, 0x39, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x45, 0x72, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0x85, 0x01, 0x0a,
	0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x01,
	0x12, 0x17, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f,
	0x4d, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x50, 0x52, 0x4f,
	0x56, 0x45, 0x10, 0x04, 0x32, 0xf7, 0x01, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x12,
	0x45, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x50,
	0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x41,
	0x73, 0x79, 0x6e, 0x63, 0x12, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x50, 0x72,
	0x6f, 0x76, 0x65, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x2d, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x4f, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x14, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_prover_proto_rawDescOnce sync.Once
	file_sdk_prover_proto_rawDescData = file_sdk_prover_proto_rawDesc
)

func file_sdk_prover_proto_rawDescGZIP() []byte {
	file_sdk_prover_proto_rawDescOnce.Do(func() {
		file_sdk_prover_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_prover_proto_rawDescData)
	})
	return file_sdk_prover_proto_rawDescData
}

var file_sdk_prover_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sdk_prover_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sdk_prover_proto_goTypes = []interface{}{
	(ErrCode)(0),                       // 0: sdk.ErrCode
	(*ProveRequest)(nil),               // 1: sdk.ProveRequest
	(*ProveResponse)(nil),              // 2: sdk.ProveResponse
	(*ProveAsyncResponse)(nil),         // 3: sdk.ProveAsyncResponse
	(*GetProofRequest)(nil),            // 4: sdk.GetProofRequest
	(*GetProofResponse)(nil),           // 5: sdk.GetProofResponse
	(*Err)(nil),                        // 6: sdk.Err
	(*IndexedReceipt)(nil),             // 7: sdk.IndexedReceipt
	(*IndexedStorage)(nil),             // 8: sdk.IndexedStorage
	(*IndexedTransaction)(nil),         // 9: sdk.IndexedTransaction
	(*CustomInput)(nil),                // 10: sdk.CustomInput
	(*commonproto.AppCircuitInfo)(nil), // 11: common.AppCircuitInfo
}
var file_sdk_prover_proto_depIdxs = []int32{
	7,  // 0: sdk.ProveRequest.receipts:type_name -> sdk.IndexedReceipt
	8,  // 1: sdk.ProveRequest.storages:type_name -> sdk.IndexedStorage
	9,  // 2: sdk.ProveRequest.transactions:type_name -> sdk.IndexedTransaction
	10, // 3: sdk.ProveRequest.custom_input:type_name -> sdk.CustomInput
	6,  // 4: sdk.ProveResponse.err:type_name -> sdk.Err
	11, // 5: sdk.ProveResponse.circuit_info:type_name -> common.AppCircuitInfo
	6,  // 6: sdk.ProveAsyncResponse.err:type_name -> sdk.Err
	11, // 7: sdk.ProveAsyncResponse.circuit_info:type_name -> common.AppCircuitInfo
	6,  // 8: sdk.GetProofResponse.err:type_name -> sdk.Err
	0,  // 9: sdk.Err.code:type_name -> sdk.ErrCode
	1,  // 10: sdk.Prover.Prove:input_type -> sdk.ProveRequest
	1,  // 11: sdk.Prover.ProveAsync:input_type -> sdk.ProveRequest
	4,  // 12: sdk.Prover.GetProof:input_type -> sdk.GetProofRequest
	2,  // 13: sdk.Prover.Prove:output_type -> sdk.ProveResponse
	3,  // 14: sdk.Prover.ProveAsync:output_type -> sdk.ProveAsyncResponse
	5,  // 15: sdk.Prover.GetProof:output_type -> sdk.GetProofResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_sdk_prover_proto_init() }
func file_sdk_prover_proto_init() {
	if File_sdk_prover_proto != nil {
		return
	}
	file_sdk_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sdk_prover_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProveRequest); i {
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
		file_sdk_prover_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProveResponse); i {
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
		file_sdk_prover_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProveAsyncResponse); i {
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
		file_sdk_prover_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProofRequest); i {
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
		file_sdk_prover_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProofResponse); i {
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
		file_sdk_prover_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Err); i {
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
			RawDescriptor: file_sdk_prover_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sdk_prover_proto_goTypes,
		DependencyIndexes: file_sdk_prover_proto_depIdxs,
		EnumInfos:         file_sdk_prover_proto_enumTypes,
		MessageInfos:      file_sdk_prover_proto_msgTypes,
	}.Build()
	File_sdk_prover_proto = out.File
	file_sdk_prover_proto_rawDesc = nil
	file_sdk_prover_proto_goTypes = nil
	file_sdk_prover_proto_depIdxs = nil
}
