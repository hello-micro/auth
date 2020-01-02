// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package hello_micro_srv_auth

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

type MiniProgramLoginRequest struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MiniProgramLoginRequest) Reset()         { *m = MiniProgramLoginRequest{} }
func (m *MiniProgramLoginRequest) String() string { return proto.CompactTextString(m) }
func (*MiniProgramLoginRequest) ProtoMessage()    {}
func (*MiniProgramLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *MiniProgramLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MiniProgramLoginRequest.Unmarshal(m, b)
}
func (m *MiniProgramLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MiniProgramLoginRequest.Marshal(b, m, deterministic)
}
func (m *MiniProgramLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiniProgramLoginRequest.Merge(m, src)
}
func (m *MiniProgramLoginRequest) XXX_Size() int {
	return xxx_messageInfo_MiniProgramLoginRequest.Size(m)
}
func (m *MiniProgramLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MiniProgramLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MiniProgramLoginRequest proto.InternalMessageInfo

func (m *MiniProgramLoginRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *MiniProgramLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type PhoneCodeLoginRequest struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhoneCodeLoginRequest) Reset()         { *m = PhoneCodeLoginRequest{} }
func (m *PhoneCodeLoginRequest) String() string { return proto.CompactTextString(m) }
func (*PhoneCodeLoginRequest) ProtoMessage()    {}
func (*PhoneCodeLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *PhoneCodeLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneCodeLoginRequest.Unmarshal(m, b)
}
func (m *PhoneCodeLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneCodeLoginRequest.Marshal(b, m, deterministic)
}
func (m *PhoneCodeLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneCodeLoginRequest.Merge(m, src)
}
func (m *PhoneCodeLoginRequest) XXX_Size() int {
	return xxx_messageInfo_PhoneCodeLoginRequest.Size(m)
}
func (m *PhoneCodeLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneCodeLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneCodeLoginRequest proto.InternalMessageInfo

func (m *PhoneCodeLoginRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *PhoneCodeLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type EmailCodeLoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailCodeLoginRequest) Reset()         { *m = EmailCodeLoginRequest{} }
func (m *EmailCodeLoginRequest) String() string { return proto.CompactTextString(m) }
func (*EmailCodeLoginRequest) ProtoMessage()    {}
func (*EmailCodeLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *EmailCodeLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailCodeLoginRequest.Unmarshal(m, b)
}
func (m *EmailCodeLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailCodeLoginRequest.Marshal(b, m, deterministic)
}
func (m *EmailCodeLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailCodeLoginRequest.Merge(m, src)
}
func (m *EmailCodeLoginRequest) XXX_Size() int {
	return xxx_messageInfo_EmailCodeLoginRequest.Size(m)
}
func (m *EmailCodeLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailCodeLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailCodeLoginRequest proto.InternalMessageInfo

func (m *EmailCodeLoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EmailCodeLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type PasswordLoginRequest struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	LoginName            string   `protobuf:"bytes,2,opt,name=loginName,proto3" json:"loginName,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordLoginRequest) Reset()         { *m = PasswordLoginRequest{} }
func (m *PasswordLoginRequest) String() string { return proto.CompactTextString(m) }
func (*PasswordLoginRequest) ProtoMessage()    {}
func (*PasswordLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{3}
}

func (m *PasswordLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordLoginRequest.Unmarshal(m, b)
}
func (m *PasswordLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordLoginRequest.Marshal(b, m, deterministic)
}
func (m *PasswordLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordLoginRequest.Merge(m, src)
}
func (m *PasswordLoginRequest) XXX_Size() int {
	return xxx_messageInfo_PasswordLoginRequest.Size(m)
}
func (m *PasswordLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordLoginRequest proto.InternalMessageInfo

func (m *PasswordLoginRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *PasswordLoginRequest) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *PasswordLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{4}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTokenRequest) Reset()         { *m = VerifyTokenRequest{} }
func (m *VerifyTokenRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyTokenRequest) ProtoMessage()    {}
func (*VerifyTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{5}
}

func (m *VerifyTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTokenRequest.Unmarshal(m, b)
}
func (m *VerifyTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTokenRequest.Marshal(b, m, deterministic)
}
func (m *VerifyTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTokenRequest.Merge(m, src)
}
func (m *VerifyTokenRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyTokenRequest.Size(m)
}
func (m *VerifyTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTokenRequest proto.InternalMessageInfo

func (m *VerifyTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyTokenResponse struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTokenResponse) Reset()         { *m = VerifyTokenResponse{} }
func (m *VerifyTokenResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyTokenResponse) ProtoMessage()    {}
func (*VerifyTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{6}
}

func (m *VerifyTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTokenResponse.Unmarshal(m, b)
}
func (m *VerifyTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTokenResponse.Marshal(b, m, deterministic)
}
func (m *VerifyTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTokenResponse.Merge(m, src)
}
func (m *VerifyTokenResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyTokenResponse.Size(m)
}
func (m *VerifyTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTokenResponse proto.InternalMessageInfo

func (m *VerifyTokenResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type RefreshTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshTokenRequest) Reset()         { *m = RefreshTokenRequest{} }
func (m *RefreshTokenRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshTokenRequest) ProtoMessage()    {}
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{7}
}

func (m *RefreshTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshTokenRequest.Unmarshal(m, b)
}
func (m *RefreshTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshTokenRequest.Marshal(b, m, deterministic)
}
func (m *RefreshTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshTokenRequest.Merge(m, src)
}
func (m *RefreshTokenRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshTokenRequest.Size(m)
}
func (m *RefreshTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshTokenRequest proto.InternalMessageInfo

func (m *RefreshTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RefreshTokenResponse struct {
	// valid 原token是否有效
	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	// token new token
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshTokenResponse) Reset()         { *m = RefreshTokenResponse{} }
func (m *RefreshTokenResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshTokenResponse) ProtoMessage()    {}
func (*RefreshTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{8}
}

func (m *RefreshTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshTokenResponse.Unmarshal(m, b)
}
func (m *RefreshTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshTokenResponse.Marshal(b, m, deterministic)
}
func (m *RefreshTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshTokenResponse.Merge(m, src)
}
func (m *RefreshTokenResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshTokenResponse.Size(m)
}
func (m *RefreshTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshTokenResponse proto.InternalMessageInfo

func (m *RefreshTokenResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *RefreshTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*MiniProgramLoginRequest)(nil), "hello.micro.srv.auth.MiniProgramLoginRequest")
	proto.RegisterType((*PhoneCodeLoginRequest)(nil), "hello.micro.srv.auth.PhoneCodeLoginRequest")
	proto.RegisterType((*EmailCodeLoginRequest)(nil), "hello.micro.srv.auth.EmailCodeLoginRequest")
	proto.RegisterType((*PasswordLoginRequest)(nil), "hello.micro.srv.auth.PasswordLoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "hello.micro.srv.auth.LoginResponse")
	proto.RegisterType((*VerifyTokenRequest)(nil), "hello.micro.srv.auth.VerifyTokenRequest")
	proto.RegisterType((*VerifyTokenResponse)(nil), "hello.micro.srv.auth.VerifyTokenResponse")
	proto.RegisterType((*RefreshTokenRequest)(nil), "hello.micro.srv.auth.RefreshTokenRequest")
	proto.RegisterType((*RefreshTokenResponse)(nil), "hello.micro.srv.auth.RefreshTokenResponse")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x51, 0x4b, 0xeb, 0x30,
	0x14, 0xc7, 0xef, 0x76, 0x77, 0xef, 0xdd, 0xce, 0x75, 0x22, 0xb1, 0xc3, 0x52, 0x7c, 0x90, 0x8a,
	0x30, 0x37, 0xac, 0xa0, 0x9f, 0x40, 0x45, 0x7c, 0x51, 0x19, 0x45, 0x7c, 0xcf, 0xd6, 0x6c, 0x8d,
	0xb6, 0x49, 0x4d, 0xba, 0x89, 0x9f, 0xda, 0xaf, 0x20, 0xc9, 0xd6, 0xad, 0xd1, 0x54, 0xea, 0x4b,
	0xe9, 0x3f, 0xe7, 0x7f, 0xfe, 0x27, 0xa7, 0xfc, 0x28, 0xf4, 0x32, 0xc1, 0x73, 0x7e, 0x8a, 0xe7,
	0x79, 0xac, 0x1f, 0x81, 0xd6, 0xc8, 0x89, 0x49, 0x92, 0xf0, 0x20, 0xa5, 0x13, 0xc1, 0x03, 0x29,
	0x16, 0x81, 0xaa, 0xf9, 0x37, 0xb0, 0x77, 0x47, 0x19, 0x1d, 0x09, 0x3e, 0x13, 0x38, 0xbd, 0xe5,
	0x33, 0xca, 0x42, 0xf2, 0x32, 0x27, 0x32, 0x47, 0x2e, 0xfc, 0x9b, 0xc4, 0x98, 0x31, 0x92, 0xb8,
	0x8d, 0x83, 0x46, 0xbf, 0x13, 0x16, 0x12, 0x21, 0x68, 0x4d, 0x78, 0x44, 0xdc, 0xa6, 0x3e, 0xd6,
	0xef, 0xfe, 0x05, 0xf4, 0x46, 0x31, 0x67, 0xe4, 0x8a, 0x47, 0xc4, 0x88, 0x71, 0xe0, 0x4f, 0xa6,
	0x0a, 0xab, 0x90, 0xa5, 0xa8, 0x8a, 0xb8, 0x4e, 0x31, 0x4d, 0x6c, 0x11, 0x44, 0x15, 0x8a, 0x08,
	0x2d, 0xac, 0x11, 0x4f, 0xe0, 0x8c, 0xb0, 0x94, 0xaf, 0x5c, 0x44, 0x35, 0x77, 0xd9, 0x87, 0x4e,
	0xa2, 0x9c, 0xf7, 0x38, 0x2d, 0xa2, 0x36, 0x07, 0xc8, 0x83, 0x76, 0xb6, 0xca, 0x73, 0x7f, 0xeb,
	0xe2, 0x5a, 0xfb, 0x47, 0xd0, 0x5d, 0xcd, 0x90, 0x19, 0x67, 0x92, 0xa8, 0x6b, 0xe6, 0xfc, 0x99,
	0xb0, 0xe2, 0x9a, 0x5a, 0xf8, 0x03, 0x40, 0x8f, 0x44, 0xd0, 0xe9, 0xdb, 0x83, 0x92, 0xa5, 0x95,
	0x2c, 0xde, 0x21, 0xec, 0x1a, 0xde, 0x4d, 0xf0, 0x02, 0x27, 0x34, 0xd2, 0xe6, 0x76, 0xb8, 0x14,
	0xca, 0x1c, 0x92, 0xa9, 0x20, 0x32, 0xae, 0x91, 0x7c, 0x09, 0x8e, 0x69, 0xfe, 0x2e, 0x7a, 0x93,
	0xd1, 0x2c, 0x65, 0x9c, 0xbd, 0xb7, 0xa0, 0xa5, 0xa0, 0x41, 0x31, 0xec, 0x7c, 0x86, 0x06, 0x9d,
	0x04, 0x36, 0xbe, 0x82, 0x0a, 0xb8, 0xbc, 0x43, 0xbb, 0xdd, 0xf8, 0xa0, 0xfe, 0x2f, 0x14, 0xc1,
	0xb6, 0x49, 0x15, 0x1a, 0xda, 0x1b, 0xad, 0xec, 0xfd, 0x60, 0x8a, 0x09, 0x5e, 0xd5, 0x14, 0x2b,
	0x9e, 0x75, 0xa7, 0x8c, 0xa1, 0x6b, 0xb0, 0x89, 0x06, 0x15, 0xab, 0x58, 0x00, 0xae, 0xbf, 0xc9,
	0xff, 0x12, 0x40, 0xa8, 0x6f, 0xef, 0xfa, 0xca, 0xa3, 0x77, 0x5c, 0xc3, 0xb9, 0x9e, 0x32, 0x83,
	0xad, 0x32, 0x4c, 0xa8, 0xa2, 0xd9, 0x42, 0xa7, 0x37, 0xa8, 0x63, 0x2d, 0x06, 0x8d, 0xff, 0xea,
	0x5f, 0xd7, 0xf9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xf6, 0xae, 0xb7, 0xd3, 0x04, 0x00,
	0x00,
}
