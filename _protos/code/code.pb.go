// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: code/code.proto

/*
Package codes is a generated protocol buffer package.

It is generated from these files:
	code/code.proto

It has these top-level messages:
*/
package codes

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The canonical error codes.
//
//
// Sometimes multiple error codes may apply.  Services should return
// the most specific error code that applies.  For example, prefer
// `OUT_OF_RANGE` over `FAILED_PRECONDITION` if both codes apply.
// Similarly prefer `NOT_FOUND` or `ALREADY_EXISTS` over `FAILED_PRECONDITION`.
type Code int32

const (
	// Not an error; returned on success
	//
	// HTTP Mapping: 200 OK
	Code_OK Code = 0
	// The operation was cancelled, typically by the caller.
	//
	// HTTP Mapping: 499 Client Closed Request
	Code_CANCELLED Code = 1
	// Unknown error.  For example, this error may be returned when
	// a `Status` value received from another address space belongs to
	// an error space that is not known in this address space.  Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_UNKNOWN Code = 2
	// The client specified an invalid argument.  Note that this differs
	// from `FAILED_PRECONDITION`.  `INVALID_ARGUMENT` indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	//
	// HTTP Mapping: 400 Bad Request
	Code_INVALID_ARGUMENT Code = 3
	// The deadline expired before the operation could complete. For operations
	// that change the state of the system, this error may be returned
	// even if the operation has completed successfully.  For example, a
	// successful response from a server could have been delayed long
	// enough for the deadline to expire.
	//
	// HTTP Mapping: 504 Gateway Timeout
	Code_DEADLINE_EXCEEDED Code = 4
	// Some requested entity (e.g., file or directory) was not found.
	//
	// Note to server developers: if a request is denied for an entire class
	// of users, such as gradual feature rollout or undocumented whitelist,
	// `NOT_FOUND` may be used. If a request is denied for some users within
	// a class of users, such as user-based access control, `PERMISSION_DENIED`
	// must be used.
	//
	// HTTP Mapping: 404 Not Found
	Code_NOT_FOUND Code = 5
	// The entity that a client attempted to create (e.g., file or directory)
	// already exists.
	//
	// HTTP Mapping: 409 Conflict
	Code_ALREADY_EXISTS Code = 6
	// The caller does not have permission to execute the specified
	// operation. `PERMISSION_DENIED` must not be used for rejections
	// caused by exhausting some resource (use `RESOURCE_EXHAUSTED`
	// instead for those errors). `PERMISSION_DENIED` must not be
	// used if the caller can not be identified (use `UNAUTHENTICATED`
	// instead for those errors). This error code does not imply the
	// request is valid or the requested entity exists or satisfies
	// other pre-conditions.
	//
	// HTTP Mapping: 403 Forbidden
	Code_PERMISSION_DENIED Code = 7
	// The request does not have valid authentication credentials for the
	// operation.
	//
	// HTTP Mapping: 401 Unauthorized
	Code_UNAUTHENTICATED Code = 16
	// Some resource has been exhausted, perhaps a per-user quota, or
	// perhaps the entire file system is out of space.
	//
	// HTTP Mapping: 429 Too Many Requests
	Code_RESOURCE_EXHAUSTED Code = 8
	// The operation was rejected because the system is not in a state
	// required for the operation's execution.  For example, the directory
	// to be deleted is non-empty, an rmdir operation is applied to
	// a non-directory, etc.
	//
	// Service implementors can use the following guidelines to decide
	// between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`:
	//  (a) Use `UNAVAILABLE` if the client can retry just the failing call.
	//  (b) Use `ABORTED` if the client should retry at a higher level
	//      (e.g., when a client-specified test-and-set fails, indicating the
	//      client should restart a read-modify-write sequence).
	//  (c) Use `FAILED_PRECONDITION` if the client should not retry until
	//      the system state has been explicitly fixed.  E.g., if an "rmdir"
	//      fails because the directory is non-empty, `FAILED_PRECONDITION`
	//      should be returned since the client should not retry unless
	//      the files are deleted from the directory.
	//
	// HTTP Mapping: 400 Bad Request
	Code_FAILED_PRECONDITION Code = 9
	// The operation was aborted, typically due to a concurrency issue such as
	// a sequencer check failure or transaction abort.
	//
	// See the guidelines above for deciding between `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 409 Conflict
	Code_ABORTED Code = 10
	// The operation was attempted past the valid range.  E.g., seeking or
	// reading past end-of-file.
	//
	// Unlike `INVALID_ARGUMENT`, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate `INVALID_ARGUMENT` if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// `OUT_OF_RANGE` if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between `FAILED_PRECONDITION` and
	// `OUT_OF_RANGE`.  We recommend using `OUT_OF_RANGE` (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an `OUT_OF_RANGE` error to detect when
	// they are done.
	//
	// HTTP Mapping: 400 Bad Request
	Code_OUT_OF_RANGE Code = 11
	// The operation is not implemented or is not supported/enabled in this
	// service.
	//
	// HTTP Mapping: 501 Not Implemented
	Code_UNIMPLEMENTED Code = 12
	// Internal errors.  This means that some invariants expected by the
	// underlying system have been broken.  This error code is reserved
	// for serious errors.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_INTERNAL Code = 13
	// The service is currently unavailable.  This is most likely a
	// transient condition, which can be corrected by retrying with
	// a backoff.
	//
	// See the guidelines above for deciding between `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 503 Service Unavailable
	Code_UNAVAILABLE Code = 14
	// Unrecoverable data loss or corruption.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_DATA_LOSS Code = 15
)

var Code_name = map[int32]string{
	0:  "OK",
	1:  "CANCELLED",
	2:  "UNKNOWN",
	3:  "INVALID_ARGUMENT",
	4:  "DEADLINE_EXCEEDED",
	5:  "NOT_FOUND",
	6:  "ALREADY_EXISTS",
	7:  "PERMISSION_DENIED",
	16: "UNAUTHENTICATED",
	8:  "RESOURCE_EXHAUSTED",
	9:  "FAILED_PRECONDITION",
	10: "ABORTED",
	11: "OUT_OF_RANGE",
	12: "UNIMPLEMENTED",
	13: "INTERNAL",
	14: "UNAVAILABLE",
	15: "DATA_LOSS",
}
var Code_value = map[string]int32{
	"OK":                  0,
	"CANCELLED":           1,
	"UNKNOWN":             2,
	"INVALID_ARGUMENT":    3,
	"DEADLINE_EXCEEDED":   4,
	"NOT_FOUND":           5,
	"ALREADY_EXISTS":      6,
	"PERMISSION_DENIED":   7,
	"UNAUTHENTICATED":     16,
	"RESOURCE_EXHAUSTED":  8,
	"FAILED_PRECONDITION": 9,
	"ABORTED":             10,
	"OUT_OF_RANGE":        11,
	"UNIMPLEMENTED":       12,
	"INTERNAL":            13,
	"UNAVAILABLE":         14,
	"DATA_LOSS":           15,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}
func (Code) EnumDescriptor() ([]byte, []int) { return fileDescriptorCode, []int{0} }

func init() {
	proto.RegisterEnum("core.codes.Code", Code_name, Code_value)
}

func init() { proto.RegisterFile("code/code.proto", fileDescriptorCode) }

var fileDescriptorCode = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x1c, 0x90, 0x49, 0x92, 0x5a, 0x31,
	0x0c, 0x86, 0x13, 0xc2, 0x28, 0x86, 0x27, 0x44, 0x86, 0x3b, 0x64, 0x41, 0x16, 0x39, 0x81, 0x78,
	0x16, 0xe0, 0xc2, 0xc8, 0x94, 0x07, 0x42, 0xb2, 0x71, 0x55, 0x80, 0xf5, 0xeb, 0xa2, 0xfb, 0x6e,
	0x7d, 0xbd, 0x2e, 0xb3, 0xd1, 0xe2, 0x2f, 0x4d, 0xdf, 0x07, 0xcd, 0xb5, 0xbb, 0xdd, 0x7f, 0xd5,
	0xb2, 0x7e, 0x79, 0x74, 0x6f, 0x1d, 0xc1, 0xb5, 0x7b, 0xdc, 0xd7, 0x35, 0x78, 0xfd, 0xf9, 0xde,
	0x83, 0x7e, 0xdb, 0xdd, 0xee, 0x34, 0x84, 0x9e, 0x3f, 0xe0, 0x27, 0x9a, 0xc3, 0xa4, 0x65, 0x6d,
	0xc5, 0x39, 0x31, 0xf8, 0x99, 0xa6, 0x30, 0xca, 0x7a, 0x50, 0xff, 0x47, 0xb1, 0x47, 0x5f, 0x01,
	0xad, 0x9e, 0xd9, 0x59, 0x53, 0x38, 0xec, 0xf2, 0x51, 0x34, 0xe1, 0x17, 0xfa, 0x06, 0x4b, 0x23,
	0x6c, 0x9c, 0x55, 0x29, 0x72, 0x69, 0x45, 0x8c, 0x18, 0xec, 0xd7, 0x45, 0xea, 0x53, 0xd9, 0xfa,
	0xac, 0x06, 0x07, 0x44, 0xb0, 0x60, 0x17, 0x84, 0xcd, 0xdf, 0x22, 0x17, 0x1b, 0x53, 0xc4, 0x61,
	0x9d, 0x3c, 0x49, 0x38, 0xda, 0x18, 0xad, 0xd7, 0x62, 0x44, 0xad, 0x18, 0x1c, 0xd1, 0x0a, 0x9a,
	0xac, 0x9c, 0xd3, 0x5e, 0x34, 0xd9, 0x96, 0x93, 0x18, 0x44, 0xfa, 0x0e, 0x14, 0x24, 0xfa, 0x1c,
	0xda, 0x7a, 0x65, 0xcf, 0x39, 0xd6, 0x7c, 0x4c, 0x3f, 0x60, 0xb5, 0x65, 0xeb, 0xc4, 0x94, 0x53,
	0x90, 0xd6, 0xab, 0xb1, 0xc9, 0x7a, 0xc5, 0x49, 0xfd, 0x9c, 0x37, 0x3e, 0xd4, 0x2e, 0x20, 0x84,
	0x99, 0xcf, 0xa9, 0xf8, 0x6d, 0x09, 0xac, 0x3b, 0xc1, 0x29, 0x2d, 0x61, 0x9e, 0xd5, 0x1e, 0x4f,
	0x4e, 0x2a, 0x86, 0x18, 0x9c, 0xd1, 0x0c, 0xc6, 0x56, 0x93, 0x04, 0x65, 0x87, 0x73, 0x6a, 0x60,
	0x9a, 0x95, 0xcf, 0x6c, 0x1d, 0x6f, 0x9c, 0xe0, 0xa2, 0x02, 0x19, 0x4e, 0x5c, 0x9c, 0x8f, 0x11,
	0x9b, 0xcd, 0xe8, 0xdf, 0xe0, 0xa9, 0xf0, 0xff, 0xf0, 0x69, 0xf5, 0xf7, 0x47, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x39, 0x75, 0xbc, 0xf1, 0x68, 0x01, 0x00, 0x00,
}
