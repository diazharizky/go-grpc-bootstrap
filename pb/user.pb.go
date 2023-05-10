// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

package pb

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

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Age                  int32    `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type UserList struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{1}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*UserList)(nil), "pb.UserList")
}

func init() {
	proto.RegisterFile("proto/user.proto", fileDescriptor_d570e3e37e5899c5)
}

var fileDescriptor_d570e3e37e5899c5 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x03, 0x33, 0x85, 0x98, 0x0a, 0x92, 0x94, 0xda, 0x18,
	0xb9, 0x58, 0x42, 0x8b, 0x53, 0x8b, 0x84, 0xa4, 0xb8, 0x38, 0x40, 0x52, 0x79, 0x89, 0xb9, 0xa9,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x2c, 0x17, 0x57, 0x5a, 0x66, 0x51,
	0x71, 0x49, 0x3c, 0x58, 0x96, 0x09, 0x2c, 0xcb, 0x09, 0x16, 0xf1, 0x03, 0x49, 0x4b, 0x73, 0x71,
	0xe6, 0x24, 0xc2, 0x64, 0x99, 0x21, 0x7a, 0x41, 0x02, 0x60, 0x49, 0x11, 0x2e, 0xd6, 0xd4, 0xdc,
	0xc4, 0xcc, 0x1c, 0x09, 0x16, 0xb0, 0x04, 0x84, 0x23, 0x24, 0xc0, 0xc5, 0x9c, 0x98, 0x9e, 0x2a,
	0xc1, 0xaa, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0x62, 0x2a, 0x69, 0x71, 0x71, 0x80, 0xdc, 0xe1, 0x93,
	0x59, 0x5c, 0x22, 0x24, 0xc7, 0xc5, 0x0a, 0xb2, 0xbb, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb,
	0x88, 0x43, 0xaf, 0x20, 0x49, 0x0f, 0x24, 0x19, 0x04, 0x11, 0x76, 0x62, 0x8b, 0x62, 0xd1, 0xd3,
	0x2f, 0x48, 0x4a, 0x62, 0x03, 0xfb, 0xc3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xde, 0xb9, 0x11,
	0x94, 0xdb, 0x00, 0x00, 0x00,
}
