// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.6.0-dev
//  VPP:              22.02-release
// source: /usr/share/vpp/api/plugins/dslite.api.json

// Package dslite contains generated bindings for API file dslite.api.
//
// Contents:
//  12 messages
//
package dslite

import (
	api "git.fd.io/govpp.git/api"
	_ "git.fd.io/govpp.git/binapi/interface_types"
	ip_types "git.fd.io/govpp.git/binapi/ip_types"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "dslite"
	APIVersion = "1.0.0"
	VersionCrc = 0x1afa049b
)

// DsliteAddDelPoolAddrRange defines message 'dslite_add_del_pool_addr_range'.
type DsliteAddDelPoolAddrRange struct {
	StartAddr ip_types.IP4Address `binapi:"ip4_address,name=start_addr" json:"start_addr,omitempty"`
	EndAddr   ip_types.IP4Address `binapi:"ip4_address,name=end_addr" json:"end_addr,omitempty"`
	IsAdd     bool                `binapi:"bool,name=is_add" json:"is_add,omitempty"`
}

func (m *DsliteAddDelPoolAddrRange) Reset()               { *m = DsliteAddDelPoolAddrRange{} }
func (*DsliteAddDelPoolAddrRange) GetMessageName() string { return "dslite_add_del_pool_addr_range" }
func (*DsliteAddDelPoolAddrRange) GetCrcString() string   { return "de2a5b02" }
func (*DsliteAddDelPoolAddrRange) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DsliteAddDelPoolAddrRange) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 4 // m.StartAddr
	size += 1 * 4 // m.EndAddr
	size += 1     // m.IsAdd
	return size
}
func (m *DsliteAddDelPoolAddrRange) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.StartAddr[:], 4)
	buf.EncodeBytes(m.EndAddr[:], 4)
	buf.EncodeBool(m.IsAdd)
	return buf.Bytes(), nil
}
func (m *DsliteAddDelPoolAddrRange) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.StartAddr[:], buf.DecodeBytes(4))
	copy(m.EndAddr[:], buf.DecodeBytes(4))
	m.IsAdd = buf.DecodeBool()
	return nil
}

// DsliteAddDelPoolAddrRangeReply defines message 'dslite_add_del_pool_addr_range_reply'.
type DsliteAddDelPoolAddrRangeReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *DsliteAddDelPoolAddrRangeReply) Reset() { *m = DsliteAddDelPoolAddrRangeReply{} }
func (*DsliteAddDelPoolAddrRangeReply) GetMessageName() string {
	return "dslite_add_del_pool_addr_range_reply"
}
func (*DsliteAddDelPoolAddrRangeReply) GetCrcString() string { return "e8d4e804" }
func (*DsliteAddDelPoolAddrRangeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DsliteAddDelPoolAddrRangeReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *DsliteAddDelPoolAddrRangeReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *DsliteAddDelPoolAddrRangeReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// DsliteAddressDetails defines message 'dslite_address_details'.
type DsliteAddressDetails struct {
	IPAddress ip_types.IP4Address `binapi:"ip4_address,name=ip_address" json:"ip_address,omitempty"`
}

func (m *DsliteAddressDetails) Reset()               { *m = DsliteAddressDetails{} }
func (*DsliteAddressDetails) GetMessageName() string { return "dslite_address_details" }
func (*DsliteAddressDetails) GetCrcString() string   { return "ec26d648" }
func (*DsliteAddressDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DsliteAddressDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 4 // m.IPAddress
	return size
}
func (m *DsliteAddressDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.IPAddress[:], 4)
	return buf.Bytes(), nil
}
func (m *DsliteAddressDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.IPAddress[:], buf.DecodeBytes(4))
	return nil
}

// DsliteAddressDump defines message 'dslite_address_dump'.
type DsliteAddressDump struct{}

func (m *DsliteAddressDump) Reset()               { *m = DsliteAddressDump{} }
func (*DsliteAddressDump) GetMessageName() string { return "dslite_address_dump" }
func (*DsliteAddressDump) GetCrcString() string   { return "51077d14" }
func (*DsliteAddressDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DsliteAddressDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *DsliteAddressDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *DsliteAddressDump) Unmarshal(b []byte) error {
	return nil
}

// DsliteGetAftrAddr defines message 'dslite_get_aftr_addr'.
type DsliteGetAftrAddr struct{}

func (m *DsliteGetAftrAddr) Reset()               { *m = DsliteGetAftrAddr{} }
func (*DsliteGetAftrAddr) GetMessageName() string { return "dslite_get_aftr_addr" }
func (*DsliteGetAftrAddr) GetCrcString() string   { return "51077d14" }
func (*DsliteGetAftrAddr) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DsliteGetAftrAddr) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *DsliteGetAftrAddr) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *DsliteGetAftrAddr) Unmarshal(b []byte) error {
	return nil
}

// DsliteGetAftrAddrReply defines message 'dslite_get_aftr_addr_reply'.
type DsliteGetAftrAddrReply struct {
	Retval  int32               `binapi:"i32,name=retval" json:"retval,omitempty"`
	IP4Addr ip_types.IP4Address `binapi:"ip4_address,name=ip4_addr" json:"ip4_addr,omitempty"`
	IP6Addr ip_types.IP6Address `binapi:"ip6_address,name=ip6_addr" json:"ip6_addr,omitempty"`
}

func (m *DsliteGetAftrAddrReply) Reset()               { *m = DsliteGetAftrAddrReply{} }
func (*DsliteGetAftrAddrReply) GetMessageName() string { return "dslite_get_aftr_addr_reply" }
func (*DsliteGetAftrAddrReply) GetCrcString() string   { return "8e23608e" }
func (*DsliteGetAftrAddrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DsliteGetAftrAddrReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Retval
	size += 1 * 4  // m.IP4Addr
	size += 1 * 16 // m.IP6Addr
	return size
}
func (m *DsliteGetAftrAddrReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeBytes(m.IP4Addr[:], 4)
	buf.EncodeBytes(m.IP6Addr[:], 16)
	return buf.Bytes(), nil
}
func (m *DsliteGetAftrAddrReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	copy(m.IP4Addr[:], buf.DecodeBytes(4))
	copy(m.IP6Addr[:], buf.DecodeBytes(16))
	return nil
}

// DsliteGetB4Addr defines message 'dslite_get_b4_addr'.
type DsliteGetB4Addr struct{}

func (m *DsliteGetB4Addr) Reset()               { *m = DsliteGetB4Addr{} }
func (*DsliteGetB4Addr) GetMessageName() string { return "dslite_get_b4_addr" }
func (*DsliteGetB4Addr) GetCrcString() string   { return "51077d14" }
func (*DsliteGetB4Addr) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DsliteGetB4Addr) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *DsliteGetB4Addr) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *DsliteGetB4Addr) Unmarshal(b []byte) error {
	return nil
}

// DsliteGetB4AddrReply defines message 'dslite_get_b4_addr_reply'.
type DsliteGetB4AddrReply struct {
	Retval  int32               `binapi:"i32,name=retval" json:"retval,omitempty"`
	IP4Addr ip_types.IP4Address `binapi:"ip4_address,name=ip4_addr" json:"ip4_addr,omitempty"`
	IP6Addr ip_types.IP6Address `binapi:"ip6_address,name=ip6_addr" json:"ip6_addr,omitempty"`
}

func (m *DsliteGetB4AddrReply) Reset()               { *m = DsliteGetB4AddrReply{} }
func (*DsliteGetB4AddrReply) GetMessageName() string { return "dslite_get_b4_addr_reply" }
func (*DsliteGetB4AddrReply) GetCrcString() string   { return "8e23608e" }
func (*DsliteGetB4AddrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DsliteGetB4AddrReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Retval
	size += 1 * 4  // m.IP4Addr
	size += 1 * 16 // m.IP6Addr
	return size
}
func (m *DsliteGetB4AddrReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeBytes(m.IP4Addr[:], 4)
	buf.EncodeBytes(m.IP6Addr[:], 16)
	return buf.Bytes(), nil
}
func (m *DsliteGetB4AddrReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	copy(m.IP4Addr[:], buf.DecodeBytes(4))
	copy(m.IP6Addr[:], buf.DecodeBytes(16))
	return nil
}

// DsliteSetAftrAddr defines message 'dslite_set_aftr_addr'.
type DsliteSetAftrAddr struct {
	IP4Addr ip_types.IP4Address `binapi:"ip4_address,name=ip4_addr" json:"ip4_addr,omitempty"`
	IP6Addr ip_types.IP6Address `binapi:"ip6_address,name=ip6_addr" json:"ip6_addr,omitempty"`
}

func (m *DsliteSetAftrAddr) Reset()               { *m = DsliteSetAftrAddr{} }
func (*DsliteSetAftrAddr) GetMessageName() string { return "dslite_set_aftr_addr" }
func (*DsliteSetAftrAddr) GetCrcString() string   { return "78b50fdf" }
func (*DsliteSetAftrAddr) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DsliteSetAftrAddr) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 4  // m.IP4Addr
	size += 1 * 16 // m.IP6Addr
	return size
}
func (m *DsliteSetAftrAddr) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.IP4Addr[:], 4)
	buf.EncodeBytes(m.IP6Addr[:], 16)
	return buf.Bytes(), nil
}
func (m *DsliteSetAftrAddr) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.IP4Addr[:], buf.DecodeBytes(4))
	copy(m.IP6Addr[:], buf.DecodeBytes(16))
	return nil
}

// DsliteSetAftrAddrReply defines message 'dslite_set_aftr_addr_reply'.
type DsliteSetAftrAddrReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *DsliteSetAftrAddrReply) Reset()               { *m = DsliteSetAftrAddrReply{} }
func (*DsliteSetAftrAddrReply) GetMessageName() string { return "dslite_set_aftr_addr_reply" }
func (*DsliteSetAftrAddrReply) GetCrcString() string   { return "e8d4e804" }
func (*DsliteSetAftrAddrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DsliteSetAftrAddrReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *DsliteSetAftrAddrReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *DsliteSetAftrAddrReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// DsliteSetB4Addr defines message 'dslite_set_b4_addr'.
type DsliteSetB4Addr struct {
	IP4Addr ip_types.IP4Address `binapi:"ip4_address,name=ip4_addr" json:"ip4_addr,omitempty"`
	IP6Addr ip_types.IP6Address `binapi:"ip6_address,name=ip6_addr" json:"ip6_addr,omitempty"`
}

func (m *DsliteSetB4Addr) Reset()               { *m = DsliteSetB4Addr{} }
func (*DsliteSetB4Addr) GetMessageName() string { return "dslite_set_b4_addr" }
func (*DsliteSetB4Addr) GetCrcString() string   { return "78b50fdf" }
func (*DsliteSetB4Addr) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DsliteSetB4Addr) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 4  // m.IP4Addr
	size += 1 * 16 // m.IP6Addr
	return size
}
func (m *DsliteSetB4Addr) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.IP4Addr[:], 4)
	buf.EncodeBytes(m.IP6Addr[:], 16)
	return buf.Bytes(), nil
}
func (m *DsliteSetB4Addr) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.IP4Addr[:], buf.DecodeBytes(4))
	copy(m.IP6Addr[:], buf.DecodeBytes(16))
	return nil
}

// DsliteSetB4AddrReply defines message 'dslite_set_b4_addr_reply'.
type DsliteSetB4AddrReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *DsliteSetB4AddrReply) Reset()               { *m = DsliteSetB4AddrReply{} }
func (*DsliteSetB4AddrReply) GetMessageName() string { return "dslite_set_b4_addr_reply" }
func (*DsliteSetB4AddrReply) GetCrcString() string   { return "e8d4e804" }
func (*DsliteSetB4AddrReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DsliteSetB4AddrReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *DsliteSetB4AddrReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *DsliteSetB4AddrReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_dslite_binapi_init() }
func file_dslite_binapi_init() {
	api.RegisterMessage((*DsliteAddDelPoolAddrRange)(nil), "dslite_add_del_pool_addr_range_de2a5b02")
	api.RegisterMessage((*DsliteAddDelPoolAddrRangeReply)(nil), "dslite_add_del_pool_addr_range_reply_e8d4e804")
	api.RegisterMessage((*DsliteAddressDetails)(nil), "dslite_address_details_ec26d648")
	api.RegisterMessage((*DsliteAddressDump)(nil), "dslite_address_dump_51077d14")
	api.RegisterMessage((*DsliteGetAftrAddr)(nil), "dslite_get_aftr_addr_51077d14")
	api.RegisterMessage((*DsliteGetAftrAddrReply)(nil), "dslite_get_aftr_addr_reply_8e23608e")
	api.RegisterMessage((*DsliteGetB4Addr)(nil), "dslite_get_b4_addr_51077d14")
	api.RegisterMessage((*DsliteGetB4AddrReply)(nil), "dslite_get_b4_addr_reply_8e23608e")
	api.RegisterMessage((*DsliteSetAftrAddr)(nil), "dslite_set_aftr_addr_78b50fdf")
	api.RegisterMessage((*DsliteSetAftrAddrReply)(nil), "dslite_set_aftr_addr_reply_e8d4e804")
	api.RegisterMessage((*DsliteSetB4Addr)(nil), "dslite_set_b4_addr_78b50fdf")
	api.RegisterMessage((*DsliteSetB4AddrReply)(nil), "dslite_set_b4_addr_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*DsliteAddDelPoolAddrRange)(nil),
		(*DsliteAddDelPoolAddrRangeReply)(nil),
		(*DsliteAddressDetails)(nil),
		(*DsliteAddressDump)(nil),
		(*DsliteGetAftrAddr)(nil),
		(*DsliteGetAftrAddrReply)(nil),
		(*DsliteGetB4Addr)(nil),
		(*DsliteGetB4AddrReply)(nil),
		(*DsliteSetAftrAddr)(nil),
		(*DsliteSetAftrAddrReply)(nil),
		(*DsliteSetB4Addr)(nil),
		(*DsliteSetB4AddrReply)(nil),
	}
}
