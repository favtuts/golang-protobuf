// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

package main

import (
	fmt "fmt"
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

type Job struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Salary               float32  `protobuf:"fixed32,2,opt,name=salary,proto3" json:"salary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Job) GetSalary() float32 {
	if m != nil {
		return m.Salary
	}
	return 0
}

type Person struct {
	Firstname            string   `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{1}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Person) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func init() {
	proto.RegisterType((*Job)(nil), "main.Job")
	proto.RegisterType((*Person)(nil), "main.Person")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x32,
	0xe6, 0x62, 0xf6, 0xca, 0x4f, 0x12, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0x2c, 0xc9, 0x49, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0xc4, 0xb8, 0xd8, 0x8a, 0x13, 0x73, 0x12, 0x8b,
	0x2a, 0x25, 0x98, 0x14, 0x18, 0x35, 0x98, 0x82, 0xa0, 0x3c, 0x25, 0x27, 0x2e, 0xb6, 0x00, 0xb0,
	0x51, 0x42, 0x32, 0x5c, 0x9c, 0x69, 0x99, 0x45, 0xc5, 0x25, 0x79, 0x89, 0xb9, 0x30, 0xbd, 0x08,
	0x01, 0x21, 0x29, 0x2e, 0x8e, 0x9c, 0x44, 0xa8, 0x24, 0x13, 0x58, 0x12, 0xce, 0x4f, 0x62, 0x03,
	0xbb, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x92, 0x0d, 0x42, 0xab, 0x95, 0x00, 0x00, 0x00,
}
