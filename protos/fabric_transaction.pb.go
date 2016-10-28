// Code generated by protoc-gen-go.
// source: fabric_transaction.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type InvalidTransaction_Cause int32

const (
	InvalidTransaction_TxIdAlreadyExists      InvalidTransaction_Cause = 0
	InvalidTransaction_RWConflictDuringCommit InvalidTransaction_Cause = 1
)

var InvalidTransaction_Cause_name = map[int32]string{
	0: "TxIdAlreadyExists",
	1: "RWConflictDuringCommit",
}
var InvalidTransaction_Cause_value = map[string]int32{
	"TxIdAlreadyExists":      0,
	"RWConflictDuringCommit": 1,
}

func (x InvalidTransaction_Cause) String() string {
	return proto.EnumName(InvalidTransaction_Cause_name, int32(x))
}
func (InvalidTransaction_Cause) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{1, 0} }

// This message is necessary to facilitate the verification of the signature
// (in the signature field) over the bytes of the transaction (in the
// transactionBytes field).
type SignedTransaction struct {
	// The bytes of the Transaction. NDD
	TransactionBytes []byte `protobuf:"bytes,1,opt,name=transactionBytes,proto3" json:"transactionBytes,omitempty"`
	// Signature of the transactionBytes The public key of the signature is in
	// the header field of TransactionAction There might be multiple
	// TransactionAction, so multiple headers, but there should be same
	// transactor identity (cert) in all headers
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignedTransaction) Reset()                    { *m = SignedTransaction{} }
func (m *SignedTransaction) String() string            { return proto.CompactTextString(m) }
func (*SignedTransaction) ProtoMessage()               {}
func (*SignedTransaction) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

// This is used to wrap an invalid Transaction with the cause
type InvalidTransaction struct {
	Transaction *Transaction2            `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
	Cause       InvalidTransaction_Cause `protobuf:"varint,2,opt,name=cause,enum=protos.InvalidTransaction_Cause" json:"cause,omitempty"`
}

func (m *InvalidTransaction) Reset()                    { *m = InvalidTransaction{} }
func (m *InvalidTransaction) String() string            { return proto.CompactTextString(m) }
func (*InvalidTransaction) ProtoMessage()               {}
func (*InvalidTransaction) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *InvalidTransaction) GetTransaction() *Transaction2 {
	if m != nil {
		return m.Transaction
	}
	return nil
}

// The transaction to be sent to the ordering service. A transaction contains
// one or more TransactionAction. Each TransactionAction binds a proposal to
// potentially multiple actions. The transaction is atomic meaning that either
// all actions in the transaction will be committed or none will.  Note that
// while a Transaction might include more than one Header, the Header.creator
// field must be the same in each.
// A single client is free to issue a number of independent Proposal, each with
// their header (Header) and request payload (ChaincodeProposalPayload).  Each
// proposal is independently endorsed generating an action
// (ProposalResponsePayload) with one signature per Endorser. Any number of
// independent proposals (and their action) might be included in a transaction
// to ensure that they are treated atomically.
type Transaction2 struct {
	// Version indicates message protocol version.
	Version int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Timestamp is the local time that the
	// message was created by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// The payload is an array of TransactionAction. An array is necessary to
	// accommodate multiple actions per transaction
	Actions []*TransactionAction `protobuf:"bytes,3,rep,name=actions" json:"actions,omitempty"`
}

func (m *Transaction2) Reset()                    { *m = Transaction2{} }
func (m *Transaction2) String() string            { return proto.CompactTextString(m) }
func (*Transaction2) ProtoMessage()               {}
func (*Transaction2) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *Transaction2) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Transaction2) GetActions() []*TransactionAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

// TransactionAction binds a proposal to its action.  The type field in the
// header dictates the type of action to be applied to the ledger.
type TransactionAction struct {
	// The header of the proposal action, which is the proposal header
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The payload of the action as defined by the type in the header For
	// chaincode, it's the bytes of ChaincodeActionPayload
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *TransactionAction) Reset()                    { *m = TransactionAction{} }
func (m *TransactionAction) String() string            { return proto.CompactTextString(m) }
func (*TransactionAction) ProtoMessage()               {}
func (*TransactionAction) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func init() {
	proto.RegisterType((*SignedTransaction)(nil), "protos.SignedTransaction")
	proto.RegisterType((*InvalidTransaction)(nil), "protos.InvalidTransaction")
	proto.RegisterType((*Transaction2)(nil), "protos.Transaction2")
	proto.RegisterType((*TransactionAction)(nil), "protos.TransactionAction")
	proto.RegisterEnum("protos.InvalidTransaction_Cause", InvalidTransaction_Cause_name, InvalidTransaction_Cause_value)
}

func init() { proto.RegisterFile("fabric_transaction.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x6b, 0xc2, 0x30,
	0x1c, 0xc5, 0xd7, 0x89, 0x8a, 0xff, 0xca, 0xd0, 0xb0, 0x49, 0x27, 0x83, 0x49, 0x4f, 0xb2, 0x43,
	0x85, 0x0a, 0x32, 0x76, 0x73, 0xce, 0x83, 0xd7, 0x4c, 0xd8, 0x69, 0x8c, 0xd8, 0xc6, 0x2e, 0xd0,
	0x26, 0x92, 0xa4, 0xa2, 0x9f, 0x64, 0x5f, 0x67, 0x1f, 0x6d, 0x98, 0x98, 0xd9, 0xe1, 0x2e, 0x2d,
	0x2f, 0xfd, 0xf5, 0xfd, 0x1f, 0xff, 0x17, 0x08, 0xd6, 0x64, 0x25, 0x59, 0xf2, 0xa1, 0x25, 0xe1,
	0x8a, 0x24, 0x9a, 0x09, 0x1e, 0x6d, 0xa4, 0xd0, 0x02, 0x35, 0xcc, 0x4b, 0xf5, 0xef, 0x33, 0x21,
	0xb2, 0x9c, 0x8e, 0x8c, 0x5c, 0x95, 0xeb, 0x91, 0x66, 0x05, 0x55, 0x9a, 0x14, 0x1b, 0x0b, 0x86,
	0xef, 0xd0, 0x7d, 0x65, 0x19, 0xa7, 0xe9, 0xf2, 0xe4, 0x81, 0x1e, 0xa0, 0x53, 0xb1, 0x7c, 0xde,
	0x6b, 0xaa, 0x02, 0x6f, 0xe0, 0x0d, 0xdb, 0xf8, 0xec, 0x1c, 0xdd, 0x41, 0x4b, 0xb1, 0x8c, 0x13,
	0x5d, 0x4a, 0x1a, 0x5c, 0x1a, 0xe8, 0x74, 0x10, 0x7e, 0x7b, 0x80, 0x16, 0x7c, 0x4b, 0x72, 0xf6,
	0x67, 0xc0, 0x04, 0xfc, 0x8a, 0x91, 0xf1, 0xf6, 0xe3, 0x6b, 0x1b, 0x49, 0x45, 0x15, 0x32, 0xc6,
	0x55, 0x10, 0x4d, 0xa0, 0x9e, 0x90, 0x52, 0xd9, 0x41, 0x57, 0xf1, 0xc0, 0xfd, 0x71, 0x3e, 0x22,
	0x9a, 0x1d, 0x38, 0x6c, 0xf1, 0xf0, 0x09, 0xea, 0x46, 0xa3, 0x1b, 0xe8, 0x2e, 0x77, 0x8b, 0x74,
	0x9a, 0x4b, 0x4a, 0xd2, 0xfd, 0x7c, 0xc7, 0x94, 0x56, 0x9d, 0x0b, 0xd4, 0x87, 0x1e, 0x7e, 0x9b,
	0x09, 0xbe, 0xce, 0x59, 0xa2, 0x5f, 0x4a, 0xc9, 0x78, 0x36, 0x13, 0x45, 0xc1, 0x74, 0xc7, 0x0b,
	0xbf, 0x3c, 0x68, 0x57, 0x13, 0xa1, 0x00, 0x9a, 0x5b, 0x2a, 0x95, 0x0b, 0x5e, 0xc7, 0x4e, 0xa2,
	0x47, 0x68, 0xfd, 0xee, 0xd7, 0x44, 0xf4, 0xe3, 0x7e, 0x64, 0x1b, 0x88, 0x5c, 0x03, 0xd1, 0xd2,
	0x11, 0xf8, 0x04, 0xa3, 0x31, 0x34, 0xad, 0xbd, 0x0a, 0x6a, 0x83, 0xda, 0xd0, 0x8f, 0x6f, 0xff,
	0x59, 0xc6, 0xd4, 0x3c, 0xb1, 0x23, 0xc3, 0x39, 0x74, 0xcf, 0xbe, 0xa2, 0x1e, 0x34, 0x3e, 0x29,
	0x49, 0xa9, 0x3c, 0x36, 0x76, 0x54, 0x87, 0xd4, 0x1b, 0xb2, 0xcf, 0x05, 0x49, 0x8f, 0x2d, 0x39,
	0xb9, 0xb2, 0x77, 0x65, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x40, 0xd2, 0x16, 0x76, 0x4e, 0x02,
	0x00, 0x00,
}
