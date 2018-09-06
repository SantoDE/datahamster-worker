// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datakube.proto

package datakube

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SaveDumpFileRequest struct {
	Targetname           string   `protobuf:"bytes,1,opt,name=targetname,proto3" json:"targetname,omitempty"`
	Filename             string   `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveDumpFileRequest) Reset()         { *m = SaveDumpFileRequest{} }
func (m *SaveDumpFileRequest) String() string { return proto.CompactTextString(m) }
func (*SaveDumpFileRequest) ProtoMessage()    {}
func (*SaveDumpFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_datakube_f38b21e522d757a4, []int{0}
}
func (m *SaveDumpFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveDumpFileRequest.Unmarshal(m, b)
}
func (m *SaveDumpFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveDumpFileRequest.Marshal(b, m, deterministic)
}
func (dst *SaveDumpFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveDumpFileRequest.Merge(dst, src)
}
func (m *SaveDumpFileRequest) XXX_Size() int {
	return xxx_messageInfo_SaveDumpFileRequest.Size(m)
}
func (m *SaveDumpFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveDumpFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveDumpFileRequest proto.InternalMessageInfo

func (m *SaveDumpFileRequest) GetTargetname() string {
	if m != nil {
		return m.Targetname
	}
	return ""
}

func (m *SaveDumpFileRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *SaveDumpFileRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SaveDumpFileResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveDumpFileResponse) Reset()         { *m = SaveDumpFileResponse{} }
func (m *SaveDumpFileResponse) String() string { return proto.CompactTextString(m) }
func (*SaveDumpFileResponse) ProtoMessage()    {}
func (*SaveDumpFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_datakube_f38b21e522d757a4, []int{1}
}
func (m *SaveDumpFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveDumpFileResponse.Unmarshal(m, b)
}
func (m *SaveDumpFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveDumpFileResponse.Marshal(b, m, deterministic)
}
func (dst *SaveDumpFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveDumpFileResponse.Merge(dst, src)
}
func (m *SaveDumpFileResponse) XXX_Size() int {
	return xxx_messageInfo_SaveDumpFileResponse.Size(m)
}
func (m *SaveDumpFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveDumpFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveDumpFileResponse proto.InternalMessageInfo

func (m *SaveDumpFileResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListJobsRequest struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobsRequest) Reset()         { *m = ListJobsRequest{} }
func (m *ListJobsRequest) String() string { return proto.CompactTextString(m) }
func (*ListJobsRequest) ProtoMessage()    {}
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_datakube_f38b21e522d757a4, []int{2}
}
func (m *ListJobsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobsRequest.Unmarshal(m, b)
}
func (m *ListJobsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobsRequest.Marshal(b, m, deterministic)
}
func (dst *ListJobsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobsRequest.Merge(dst, src)
}
func (m *ListJobsRequest) XXX_Size() int {
	return xxx_messageInfo_ListJobsRequest.Size(m)
}
func (m *ListJobsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobsRequest proto.InternalMessageInfo

func (m *ListJobsRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Target struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string       `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Credentials          *Credentials `protobuf:"bytes,3,opt,name=credentials,proto3" json:"credentials,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_datakube_f38b21e522d757a4, []int{3}
}
func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (dst *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(dst, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Target) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Target) GetCredentials() *Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

type Credentials struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 string   `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	User                 string   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Database             string   `protobuf:"bytes,5,opt,name=database,proto3" json:"database,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}
func (*Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_datakube_f38b21e522d757a4, []int{4}
}
func (m *Credentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credentials.Unmarshal(m, b)
}
func (m *Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credentials.Marshal(b, m, deterministic)
}
func (dst *Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credentials.Merge(dst, src)
}
func (m *Credentials) XXX_Size() int {
	return xxx_messageInfo_Credentials.Size(m)
}
func (m *Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Credentials proto.InternalMessageInfo

func (m *Credentials) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Credentials) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Credentials) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Credentials) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type Job struct {
	Target               *Target  `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_datakube_f38b21e522d757a4, []int{5}
}
func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (dst *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(dst, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Job) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type ListJobsResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Jobs                 []*Job   `protobuf:"bytes,2,rep,name=jobs,proto3" json:"jobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobsResponse) Reset()         { *m = ListJobsResponse{} }
func (m *ListJobsResponse) String() string { return proto.CompactTextString(m) }
func (*ListJobsResponse) ProtoMessage()    {}
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_datakube_f38b21e522d757a4, []int{6}
}
func (m *ListJobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobsResponse.Unmarshal(m, b)
}
func (m *ListJobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobsResponse.Marshal(b, m, deterministic)
}
func (dst *ListJobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobsResponse.Merge(dst, src)
}
func (m *ListJobsResponse) XXX_Size() int {
	return xxx_messageInfo_ListJobsResponse.Size(m)
}
func (m *ListJobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobsResponse proto.InternalMessageInfo

func (m *ListJobsResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ListJobsResponse) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func init() {
	proto.RegisterType((*SaveDumpFileRequest)(nil), "datakube.SaveDumpFileRequest")
	proto.RegisterType((*SaveDumpFileResponse)(nil), "datakube.SaveDumpFileResponse")
	proto.RegisterType((*ListJobsRequest)(nil), "datakube.ListJobsRequest")
	proto.RegisterType((*Target)(nil), "datakube.Target")
	proto.RegisterType((*Credentials)(nil), "datakube.Credentials")
	proto.RegisterType((*Job)(nil), "datakube.Job")
	proto.RegisterType((*ListJobsResponse)(nil), "datakube.ListJobsResponse")
}

func init() { proto.RegisterFile("datakube.proto", fileDescriptor_datakube_f38b21e522d757a4) }

var fileDescriptor_datakube_f38b21e522d757a4 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4e, 0xe3, 0x30,
	0x10, 0xdd, 0xb4, 0x69, 0x36, 0x9d, 0x74, 0x77, 0x2b, 0x6f, 0x41, 0x21, 0x12, 0x55, 0xf1, 0x29,
	0x5c, 0x2a, 0x14, 0x0e, 0x7c, 0x40, 0x81, 0x43, 0x85, 0x54, 0xc9, 0xf0, 0x03, 0x4e, 0x6b, 0x20,
	0xd0, 0xd6, 0x21, 0xe3, 0x80, 0x38, 0xf2, 0x27, 0x7c, 0x2a, 0xb2, 0x53, 0x37, 0x01, 0x15, 0x71,
	0x9b, 0x37, 0xcf, 0x9e, 0xf7, 0xf4, 0x66, 0xe0, 0xef, 0x82, 0x2b, 0xfe, 0x58, 0xa6, 0x62, 0x9c,
	0x17, 0x52, 0x49, 0xe2, 0x5b, 0x4c, 0x05, 0xfc, 0xbf, 0xe6, 0xcf, 0xe2, 0xbc, 0x5c, 0xe5, 0x97,
	0xd9, 0x52, 0x30, 0xf1, 0x54, 0x0a, 0x54, 0x64, 0x08, 0xa0, 0x78, 0x71, 0x27, 0xd4, 0x9a, 0xaf,
	0x44, 0xe8, 0x8c, 0x9c, 0xb8, 0xcb, 0x1a, 0x1d, 0x12, 0x81, 0x7f, 0x9b, 0x2d, 0x85, 0x61, 0x5b,
	0x86, 0xdd, 0x62, 0x42, 0xc0, 0xd5, 0xe3, 0xc3, 0xf6, 0xc8, 0x89, 0x7b, 0xcc, 0xd4, 0xf4, 0x04,
	0x06, 0x9f, 0x65, 0x30, 0x97, 0x6b, 0x14, 0x24, 0x84, 0xdf, 0x58, 0xce, 0xe7, 0x02, 0xd1, 0x88,
	0xf8, 0xcc, 0x42, 0x7a, 0x0c, 0xff, 0xae, 0x32, 0x54, 0x53, 0x99, 0xa2, 0x35, 0xb5, 0x0f, 0x1e,
	0x2a, 0xae, 0x4a, 0xdc, 0x18, 0xda, 0x20, 0x9a, 0x81, 0x77, 0x63, 0xac, 0x69, 0xe9, 0x86, 0x61,
	0xd7, 0xda, 0x51, 0xaf, 0xb9, 0xb5, 0x69, 0x6a, 0x72, 0x06, 0xc1, 0xbc, 0x10, 0x0b, 0xb1, 0x56,
	0x19, 0x5f, 0xa2, 0x71, 0x1a, 0x24, 0x7b, 0xe3, 0x6d, 0x4a, 0x93, 0x9a, 0x64, 0xcd, 0x97, 0xf4,
	0xcd, 0x81, 0xa0, 0x41, 0xea, 0xe1, 0xf7, 0x12, 0x95, 0x15, 0xd4, 0xb5, 0xee, 0xe5, 0xb2, 0x50,
	0x56, 0x50, 0xd7, 0xba, 0x57, 0xa2, 0x28, 0x8c, 0x52, 0x97, 0x99, 0x5a, 0x67, 0x98, 0x73, 0xc4,
	0x17, 0x59, 0x2c, 0x42, 0xb7, 0xca, 0xd0, 0x62, 0xcd, 0x69, 0x33, 0x29, 0x47, 0x11, 0x76, 0x2a,
	0xce, 0x62, 0x7a, 0x01, 0xed, 0xa9, 0x4c, 0x49, 0x0c, 0x5e, 0xb5, 0x10, 0x23, 0x1e, 0x24, 0xfd,
	0xda, 0x7e, 0x95, 0x06, 0xdb, 0xf0, 0x64, 0x00, 0x1d, 0x9d, 0x94, 0x8d, 0xa0, 0x02, 0x74, 0x06,
	0xfd, 0x3a, 0xe0, 0x9f, 0xd6, 0x41, 0x8e, 0xc0, 0x7d, 0x90, 0x29, 0x86, 0xad, 0x51, 0x3b, 0x0e,
	0x92, 0x3f, 0xb5, 0xd6, 0x54, 0xa6, 0xcc, 0x50, 0xc9, 0xbb, 0x03, 0xdb, 0xbb, 0x22, 0x33, 0xe8,
	0x35, 0x17, 0x4e, 0x0e, 0xeb, 0x1f, 0x3b, 0xee, 0x2d, 0x1a, 0x7e, 0x47, 0x57, 0xc6, 0xe8, 0x2f,
	0x32, 0x01, 0xdf, 0xda, 0x25, 0x07, 0xf5, 0xeb, 0x2f, 0x37, 0x12, 0x45, 0xbb, 0x28, 0x3b, 0x24,
	0xf5, 0xcc, 0xf9, 0x9f, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0xab, 0xe0, 0x60, 0x04, 0x10, 0x03,
	0x00, 0x00,
}
