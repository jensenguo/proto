// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wxpay.proto

package wxpay

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TransactionsReq_PackageType int32

const (
	TransactionsReq_Unkown       TransactionsReq_PackageType = 0
	TransactionsReq_PackageOne   TransactionsReq_PackageType = 1
	TransactionsReq_PackageTwo   TransactionsReq_PackageType = 2
	TransactionsReq_PackageThree TransactionsReq_PackageType = 3
	TransactionsReq_PackageFour  TransactionsReq_PackageType = 4
)

var TransactionsReq_PackageType_name = map[int32]string{
	0: "Unkown",
	1: "PackageOne",
	2: "PackageTwo",
	3: "PackageThree",
	4: "PackageFour",
}

var TransactionsReq_PackageType_value = map[string]int32{
	"Unkown":       0,
	"PackageOne":   1,
	"PackageTwo":   2,
	"PackageThree": 3,
	"PackageFour":  4,
}

func (x TransactionsReq_PackageType) String() string {
	return proto.EnumName(TransactionsReq_PackageType_name, int32(x))
}

func (TransactionsReq_PackageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f807639548d25710, []int{0, 0}
}

// 微信交易接口
type TransactionsReq struct {
	UserOpenid           string                      `protobuf:"bytes,1,opt,name=user_openid,json=userOpenid,proto3" json:"user_openid,omitempty"`
	VipPackage           TransactionsReq_PackageType `protobuf:"varint,2,opt,name=vip_package,json=vipPackage,proto3,enum=wxpay.TransactionsReq_PackageType" json:"vip_package,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TransactionsReq) Reset()         { *m = TransactionsReq{} }
func (m *TransactionsReq) String() string { return proto.CompactTextString(m) }
func (*TransactionsReq) ProtoMessage()    {}
func (*TransactionsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f807639548d25710, []int{0}
}
func (m *TransactionsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionsReq.Unmarshal(m, b)
}
func (m *TransactionsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionsReq.Marshal(b, m, deterministic)
}
func (m *TransactionsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionsReq.Merge(m, src)
}
func (m *TransactionsReq) XXX_Size() int {
	return xxx_messageInfo_TransactionsReq.Size(m)
}
func (m *TransactionsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionsReq.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionsReq proto.InternalMessageInfo

func (m *TransactionsReq) GetUserOpenid() string {
	if m != nil {
		return m.UserOpenid
	}
	return ""
}

func (m *TransactionsReq) GetVipPackage() TransactionsReq_PackageType {
	if m != nil {
		return m.VipPackage
	}
	return TransactionsReq_Unkown
}

type TransactionsRsp struct {
	Appid                string   `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nonce                string   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Package              string   `protobuf:"bytes,4,opt,name=package,proto3" json:"package,omitempty"`
	SignType             string   `protobuf:"bytes,5,opt,name=sign_type,json=signType,proto3" json:"sign_type,omitempty"`
	PaySign              string   `protobuf:"bytes,6,opt,name=pay_sign,json=paySign,proto3" json:"pay_sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionsRsp) Reset()         { *m = TransactionsRsp{} }
func (m *TransactionsRsp) String() string { return proto.CompactTextString(m) }
func (*TransactionsRsp) ProtoMessage()    {}
func (*TransactionsRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f807639548d25710, []int{1}
}
func (m *TransactionsRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionsRsp.Unmarshal(m, b)
}
func (m *TransactionsRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionsRsp.Marshal(b, m, deterministic)
}
func (m *TransactionsRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionsRsp.Merge(m, src)
}
func (m *TransactionsRsp) XXX_Size() int {
	return xxx_messageInfo_TransactionsRsp.Size(m)
}
func (m *TransactionsRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionsRsp.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionsRsp proto.InternalMessageInfo

func (m *TransactionsRsp) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *TransactionsRsp) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *TransactionsRsp) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *TransactionsRsp) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *TransactionsRsp) GetSignType() string {
	if m != nil {
		return m.SignType
	}
	return ""
}

func (m *TransactionsRsp) GetPaySign() string {
	if m != nil {
		return m.PaySign
	}
	return ""
}

// 交易结果回调接口
type PayCallbackReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayCallbackReq) Reset()         { *m = PayCallbackReq{} }
func (m *PayCallbackReq) String() string { return proto.CompactTextString(m) }
func (*PayCallbackReq) ProtoMessage()    {}
func (*PayCallbackReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f807639548d25710, []int{2}
}
func (m *PayCallbackReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayCallbackReq.Unmarshal(m, b)
}
func (m *PayCallbackReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayCallbackReq.Marshal(b, m, deterministic)
}
func (m *PayCallbackReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayCallbackReq.Merge(m, src)
}
func (m *PayCallbackReq) XXX_Size() int {
	return xxx_messageInfo_PayCallbackReq.Size(m)
}
func (m *PayCallbackReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PayCallbackReq.DiscardUnknown(m)
}

var xxx_messageInfo_PayCallbackReq proto.InternalMessageInfo

type PayCallbackRsp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayCallbackRsp) Reset()         { *m = PayCallbackRsp{} }
func (m *PayCallbackRsp) String() string { return proto.CompactTextString(m) }
func (*PayCallbackRsp) ProtoMessage()    {}
func (*PayCallbackRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f807639548d25710, []int{3}
}
func (m *PayCallbackRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayCallbackRsp.Unmarshal(m, b)
}
func (m *PayCallbackRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayCallbackRsp.Marshal(b, m, deterministic)
}
func (m *PayCallbackRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayCallbackRsp.Merge(m, src)
}
func (m *PayCallbackRsp) XXX_Size() int {
	return xxx_messageInfo_PayCallbackRsp.Size(m)
}
func (m *PayCallbackRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PayCallbackRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PayCallbackRsp proto.InternalMessageInfo

func (m *PayCallbackRsp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *PayCallbackRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// 获取用户openid接口
type GetUserOpenidReq struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserOpenidReq) Reset()         { *m = GetUserOpenidReq{} }
func (m *GetUserOpenidReq) String() string { return proto.CompactTextString(m) }
func (*GetUserOpenidReq) ProtoMessage()    {}
func (*GetUserOpenidReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f807639548d25710, []int{4}
}
func (m *GetUserOpenidReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserOpenidReq.Unmarshal(m, b)
}
func (m *GetUserOpenidReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserOpenidReq.Marshal(b, m, deterministic)
}
func (m *GetUserOpenidReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserOpenidReq.Merge(m, src)
}
func (m *GetUserOpenidReq) XXX_Size() int {
	return xxx_messageInfo_GetUserOpenidReq.Size(m)
}
func (m *GetUserOpenidReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserOpenidReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserOpenidReq proto.InternalMessageInfo

func (m *GetUserOpenidReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type GetUserOpenidRsp struct {
	Openid               string   `protobuf:"bytes,1,opt,name=openid,proto3" json:"openid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserOpenidRsp) Reset()         { *m = GetUserOpenidRsp{} }
func (m *GetUserOpenidRsp) String() string { return proto.CompactTextString(m) }
func (*GetUserOpenidRsp) ProtoMessage()    {}
func (*GetUserOpenidRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f807639548d25710, []int{5}
}
func (m *GetUserOpenidRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserOpenidRsp.Unmarshal(m, b)
}
func (m *GetUserOpenidRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserOpenidRsp.Marshal(b, m, deterministic)
}
func (m *GetUserOpenidRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserOpenidRsp.Merge(m, src)
}
func (m *GetUserOpenidRsp) XXX_Size() int {
	return xxx_messageInfo_GetUserOpenidRsp.Size(m)
}
func (m *GetUserOpenidRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserOpenidRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserOpenidRsp proto.InternalMessageInfo

func (m *GetUserOpenidRsp) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

// 获取用户vip信息接口
type GetUserAmountReq struct {
	UserOpenid           string   `protobuf:"bytes,1,opt,name=user_openid,json=userOpenid,proto3" json:"user_openid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserAmountReq) Reset()         { *m = GetUserAmountReq{} }
func (m *GetUserAmountReq) String() string { return proto.CompactTextString(m) }
func (*GetUserAmountReq) ProtoMessage()    {}
func (*GetUserAmountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f807639548d25710, []int{6}
}
func (m *GetUserAmountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserAmountReq.Unmarshal(m, b)
}
func (m *GetUserAmountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserAmountReq.Marshal(b, m, deterministic)
}
func (m *GetUserAmountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserAmountReq.Merge(m, src)
}
func (m *GetUserAmountReq) XXX_Size() int {
	return xxx_messageInfo_GetUserAmountReq.Size(m)
}
func (m *GetUserAmountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserAmountReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserAmountReq proto.InternalMessageInfo

func (m *GetUserAmountReq) GetUserOpenid() string {
	if m != nil {
		return m.UserOpenid
	}
	return ""
}

type GetUserAmountRsp struct {
	Amount               int32    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserAmountRsp) Reset()         { *m = GetUserAmountRsp{} }
func (m *GetUserAmountRsp) String() string { return proto.CompactTextString(m) }
func (*GetUserAmountRsp) ProtoMessage()    {}
func (*GetUserAmountRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f807639548d25710, []int{7}
}
func (m *GetUserAmountRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserAmountRsp.Unmarshal(m, b)
}
func (m *GetUserAmountRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserAmountRsp.Marshal(b, m, deterministic)
}
func (m *GetUserAmountRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserAmountRsp.Merge(m, src)
}
func (m *GetUserAmountRsp) XXX_Size() int {
	return xxx_messageInfo_GetUserAmountRsp.Size(m)
}
func (m *GetUserAmountRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserAmountRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserAmountRsp proto.InternalMessageInfo

func (m *GetUserAmountRsp) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// *************************其他接口协议****************************
type GetAccessTokenRsp struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ExpiresIn            int64    `protobuf:"varint,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	RefreshToken         string   `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Openid               string   `protobuf:"bytes,4,opt,name=openid,proto3" json:"openid,omitempty"`
	Scope                string   `protobuf:"bytes,5,opt,name=scope,proto3" json:"scope,omitempty"`
	Unionid              string   `protobuf:"bytes,6,opt,name=unionid,proto3" json:"unionid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccessTokenRsp) Reset()         { *m = GetAccessTokenRsp{} }
func (m *GetAccessTokenRsp) String() string { return proto.CompactTextString(m) }
func (*GetAccessTokenRsp) ProtoMessage()    {}
func (*GetAccessTokenRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f807639548d25710, []int{8}
}
func (m *GetAccessTokenRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccessTokenRsp.Unmarshal(m, b)
}
func (m *GetAccessTokenRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccessTokenRsp.Marshal(b, m, deterministic)
}
func (m *GetAccessTokenRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccessTokenRsp.Merge(m, src)
}
func (m *GetAccessTokenRsp) XXX_Size() int {
	return xxx_messageInfo_GetAccessTokenRsp.Size(m)
}
func (m *GetAccessTokenRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccessTokenRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccessTokenRsp proto.InternalMessageInfo

func (m *GetAccessTokenRsp) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GetAccessTokenRsp) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func (m *GetAccessTokenRsp) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *GetAccessTokenRsp) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *GetAccessTokenRsp) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *GetAccessTokenRsp) GetUnionid() string {
	if m != nil {
		return m.Unionid
	}
	return ""
}

func init() {
	proto.RegisterEnum("wxpay.TransactionsReq_PackageType", TransactionsReq_PackageType_name, TransactionsReq_PackageType_value)
	proto.RegisterType((*TransactionsReq)(nil), "wxpay.TransactionsReq")
	proto.RegisterType((*TransactionsRsp)(nil), "wxpay.TransactionsRsp")
	proto.RegisterType((*PayCallbackReq)(nil), "wxpay.PayCallbackReq")
	proto.RegisterType((*PayCallbackRsp)(nil), "wxpay.PayCallbackRsp")
	proto.RegisterType((*GetUserOpenidReq)(nil), "wxpay.GetUserOpenidReq")
	proto.RegisterType((*GetUserOpenidRsp)(nil), "wxpay.GetUserOpenidRsp")
	proto.RegisterType((*GetUserAmountReq)(nil), "wxpay.GetUserAmountReq")
	proto.RegisterType((*GetUserAmountRsp)(nil), "wxpay.GetUserAmountRsp")
	proto.RegisterType((*GetAccessTokenRsp)(nil), "wxpay.GetAccessTokenRsp")
}

func init() { proto.RegisterFile("wxpay.proto", fileDescriptor_f807639548d25710) }

var fileDescriptor_f807639548d25710 = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6a, 0xdb, 0x4a,
	0x14, 0xbe, 0xf2, 0x5f, 0xe2, 0xb1, 0xe3, 0xe8, 0x0e, 0xb9, 0x89, 0xae, 0x93, 0x42, 0xaa, 0x6e,
	0x8c, 0xa1, 0x16, 0xa4, 0x14, 0x4a, 0x28, 0x85, 0xb8, 0xd0, 0xd0, 0x55, 0x82, 0xe2, 0x50, 0x28,
	0x05, 0x31, 0x91, 0x4f, 0xe5, 0xa9, 0xed, 0x99, 0x89, 0x46, 0xf2, 0x0f, 0xdd, 0xe5, 0x15, 0xfa,
	0x1c, 0x7d, 0x89, 0xbe, 0x42, 0x5f, 0xa1, 0xeb, 0x2e, 0xba, 0xcc, 0xaa, 0x68, 0x34, 0xb2, 0x65,
	0x63, 0x4a, 0xdb, 0x95, 0xe7, 0xfb, 0xce, 0xcf, 0x37, 0x47, 0xe7, 0x1b, 0xa3, 0xda, 0x74, 0x26,
	0xc8, 0xbc, 0x23, 0x42, 0x1e, 0x71, 0x5c, 0x56, 0xa0, 0x79, 0x14, 0x70, 0x1e, 0x8c, 0xc0, 0x21,
	0x82, 0x3a, 0x84, 0x31, 0x1e, 0x91, 0x88, 0x72, 0x26, 0xd3, 0xa4, 0xe6, 0xc1, 0x84, 0x8c, 0x68,
	0x9f, 0x44, 0xe0, 0x64, 0x07, 0x1d, 0xd8, 0x0b, 0x78, 0xc0, 0xd5, 0xd1, 0x49, 0x4e, 0x29, 0x6b,
	0x7f, 0x37, 0xd0, 0x6e, 0x2f, 0x24, 0x4c, 0x12, 0x5f, 0x75, 0x71, 0xe1, 0x16, 0xb7, 0x50, 0x2d,
	0x96, 0x10, 0x7a, 0x5c, 0x00, 0xa3, 0x7d, 0xcb, 0x38, 0x36, 0x5a, 0xd5, 0xee, 0xd6, 0x7d, 0xb7,
	0x14, 0x16, 0x4c, 0xc3, 0x45, 0x49, 0xec, 0x42, 0x85, 0xf0, 0x15, 0xaa, 0x4d, 0xa8, 0xf0, 0x04,
	0xf1, 0x87, 0x24, 0x00, 0xab, 0x70, 0x6c, 0xb4, 0x1a, 0x27, 0x76, 0x27, 0xbd, 0xf4, 0x5a, 0xdb,
	0xce, 0x65, 0x9a, 0xd5, 0x9b, 0x0b, 0xe8, 0x36, 0xee, 0xbb, 0xb5, 0x3b, 0x63, 0xdb, 0x32, 0xac,
	0x82, 0x55, 0xb4, 0x4a, 0x2e, 0x9a, 0x50, 0xa1, 0xe3, 0xf6, 0x3b, 0x54, 0xcb, 0xa5, 0x62, 0x84,
	0x2a, 0xd7, 0x6c, 0xc8, 0xa7, 0xcc, 0xfc, 0x07, 0x37, 0x10, 0xd2, 0xa1, 0x0b, 0x06, 0xa6, 0x91,
	0xc3, 0xbd, 0x29, 0x37, 0x0b, 0xd8, 0x44, 0xf5, 0x0c, 0x0f, 0x42, 0x00, 0xb3, 0x88, 0x77, 0x17,
	0xcd, 0x5e, 0xf1, 0x38, 0x34, 0x4b, 0xf6, 0xe7, 0xf5, 0x81, 0xa5, 0xc0, 0x7b, 0xa8, 0x4c, 0x84,
	0xc8, 0x46, 0x75, 0x53, 0x80, 0x8f, 0x50, 0x35, 0xa2, 0x63, 0x90, 0x11, 0x19, 0x0b, 0x35, 0x5a,
	0xd5, 0x5d, 0x12, 0x49, 0x0d, 0xe3, 0xcc, 0x07, 0xab, 0x98, 0xd6, 0x28, 0x80, 0x2d, 0xb4, 0x95,
	0x7d, 0x8c, 0x92, 0xe2, 0x33, 0x88, 0x0f, 0x51, 0x55, 0xd2, 0x80, 0x79, 0xd1, 0x5c, 0x80, 0x55,
	0x56, 0xb1, 0xed, 0x84, 0x50, 0x33, 0xfe, 0x8f, 0xb6, 0x05, 0x99, 0x7b, 0x09, 0xb6, 0x2a, 0x59,
	0xdd, 0xfc, 0x8a, 0x06, 0xcc, 0x36, 0x51, 0xe3, 0x92, 0xcc, 0x5f, 0x92, 0xd1, 0xe8, 0x86, 0xf8,
	0x43, 0x17, 0x6e, 0xed, 0x17, 0xab, 0x8c, 0x14, 0x18, 0xa3, 0x92, 0xcf, 0xfb, 0xa0, 0xaf, 0xaf,
	0xce, 0xc9, 0x4d, 0xc6, 0x20, 0x65, 0xb6, 0x96, 0xaa, 0x9b, 0x41, 0xdb, 0x41, 0xe6, 0x39, 0x44,
	0xd7, 0x8b, 0x2d, 0x26, 0x2b, 0x3f, 0xcc, 0x77, 0x58, 0xee, 0x5a, 0x91, 0x76, 0x7b, 0xbd, 0x40,
	0x0a, 0xbc, 0x8f, 0x2a, 0x79, 0x7b, 0xb8, 0x1a, 0xd9, 0xcf, 0x17, 0xb9, 0x67, 0x63, 0x1e, 0xb3,
	0xe8, 0x8f, 0xfc, 0x94, 0x53, 0xd2, 0xd5, 0xa9, 0x12, 0x51, 0x40, 0x15, 0x96, 0x5d, 0x8d, 0xec,
	0x2f, 0x06, 0xfa, 0xf7, 0x1c, 0xa2, 0x33, 0xdf, 0x07, 0x29, 0x7b, 0x7c, 0x08, 0x2c, 0xc9, 0x7e,
	0x88, 0xea, 0x44, 0x31, 0x5e, 0x94, 0x50, 0xfa, 0x76, 0x35, 0xb2, 0xcc, 0xc2, 0x0f, 0x10, 0x82,
	0x99, 0xa0, 0x21, 0x48, 0x8f, 0x32, 0xf5, 0x71, 0x8a, 0x6e, 0x55, 0x33, 0xaf, 0x19, 0x7e, 0x84,
	0x76, 0x42, 0x78, 0x1f, 0x82, 0x1c, 0xe8, 0x16, 0xe9, 0x82, 0xeb, 0x9a, 0x4c, 0x7b, 0x2c, 0xc7,
	0x2f, 0xe5, 0xc7, 0x4f, 0x5c, 0x21, 0x7d, 0xbe, 0xd8, 0x70, 0x0a, 0x92, 0x5d, 0xc4, 0x8c, 0xf2,
	0x24, 0x5d, 0x6f, 0x57, 0xc3, 0x93, 0x1f, 0x45, 0x54, 0x7b, 0x03, 0x33, 0xca, 0x2e, 0xc9, 0xfc,
	0x6a, 0x12, 0xe2, 0x18, 0xd5, 0xf3, 0xe6, 0xc4, 0xfb, 0x9b, 0xdf, 0x52, 0x73, 0x23, 0x2f, 0x85,
	0xfd, 0xf4, 0xee, 0xeb, 0xb7, 0x4f, 0x05, 0xc7, 0x6e, 0x3b, 0x72, 0x10, 0xf7, 0x39, 0x0b, 0x1c,
	0x95, 0xe7, 0x4c, 0x81, 0xce, 0x28, 0xf3, 0x94, 0xc9, 0x26, 0xa1, 0x13, 0xe5, 0xca, 0x4e, 0x8d,
	0x36, 0x0e, 0x93, 0x57, 0xb2, 0xb0, 0x14, 0xfe, 0x4f, 0x77, 0x5f, 0x35, 0x5e, 0x73, 0x13, 0xfd,
	0xfb, 0x9a, 0xc9, 0xaf, 0xaf, 0xcb, 0x12, 0xcd, 0x8f, 0x68, 0x67, 0xc5, 0x55, 0xf8, 0x40, 0xb7,
	0x5f, 0x37, 0x67, 0x73, 0x73, 0x40, 0x0a, 0xfb, 0x99, 0x52, 0x3e, 0xb1, 0x1f, 0xff, 0x5a, 0x39,
	0x80, 0xc8, 0xcb, 0x19, 0x70, 0x55, 0x3c, 0x35, 0xda, 0xba, 0xf8, 0xc2, 0xbc, 0xcd, 0xcd, 0x81,
	0xbf, 0x10, 0x4f, 0x6d, 0x7b, 0x6a, 0xb4, 0xbb, 0xed, 0xb7, 0xad, 0x80, 0x46, 0x83, 0xf8, 0xa6,
	0xe3, 0xf3, 0xb1, 0xf3, 0x01, 0x98, 0x04, 0x16, 0xc4, 0xdc, 0x49, 0xff, 0x9c, 0x57, 0xba, 0xdd,
	0x54, 0x14, 0xf9, 0xe4, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0xaa, 0x00, 0xf5, 0x09, 0x06,
	0x00, 0x00,
}
