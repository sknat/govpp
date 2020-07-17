// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.05-release
// source: /usr/share/vpp/api/plugins/udp_ping.api.json

// Package udp_ping contains generated bindings for API file udp_ping.api.
//
// Contents:
//   4 messages
//
package udp_ping

import (
	api "git.fd.io/govpp.git/api"
	ip_types "git.fd.io/govpp.git/binapi/ip_types"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "udp_ping"
	APIVersion = "3.0.0"
	VersionCrc = 0x3b2c67de
)

// UDPPingAddDel defines message 'udp_ping_add_del'.
type UDPPingAddDel struct {
	SrcIPAddress ip_types.Address `binapi:"address,name=src_ip_address" json:"src_ip_address,omitempty"`
	DstIPAddress ip_types.Address `binapi:"address,name=dst_ip_address" json:"dst_ip_address,omitempty"`
	StartSrcPort uint16           `binapi:"u16,name=start_src_port" json:"start_src_port,omitempty"`
	EndSrcPort   uint16           `binapi:"u16,name=end_src_port" json:"end_src_port,omitempty"`
	StartDstPort uint16           `binapi:"u16,name=start_dst_port" json:"start_dst_port,omitempty"`
	EndDstPort   uint16           `binapi:"u16,name=end_dst_port" json:"end_dst_port,omitempty"`
	Interval     uint16           `binapi:"u16,name=interval" json:"interval,omitempty"`
	Dis          uint8            `binapi:"u8,name=dis" json:"dis,omitempty"`
	FaultDet     uint8            `binapi:"u8,name=fault_det" json:"fault_det,omitempty"`
	Reserve      []byte           `binapi:"u8[3],name=reserve" json:"reserve,omitempty"`
}

func (m *UDPPingAddDel) Reset()               { *m = UDPPingAddDel{} }
func (*UDPPingAddDel) GetMessageName() string { return "udp_ping_add_del" }
func (*UDPPingAddDel) GetCrcString() string   { return "c692b188" }
func (*UDPPingAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *UDPPingAddDel) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 1      // m.SrcIPAddress.Af
	size += 1 * 16 // m.SrcIPAddress.Un
	size += 1      // m.DstIPAddress.Af
	size += 1 * 16 // m.DstIPAddress.Un
	size += 2      // m.StartSrcPort
	size += 2      // m.EndSrcPort
	size += 2      // m.StartDstPort
	size += 2      // m.EndDstPort
	size += 2      // m.Interval
	size += 1      // m.Dis
	size += 1      // m.FaultDet
	size += 1 * 3  // m.Reserve
	return size
}
func (m *UDPPingAddDel) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint8(uint8(m.SrcIPAddress.Af))
	buf.EncodeBytes(m.SrcIPAddress.Un.XXX_UnionData[:], 0)
	buf.EncodeUint8(uint8(m.DstIPAddress.Af))
	buf.EncodeBytes(m.DstIPAddress.Un.XXX_UnionData[:], 0)
	buf.EncodeUint16(uint16(m.StartSrcPort))
	buf.EncodeUint16(uint16(m.EndSrcPort))
	buf.EncodeUint16(uint16(m.StartDstPort))
	buf.EncodeUint16(uint16(m.EndDstPort))
	buf.EncodeUint16(uint16(m.Interval))
	buf.EncodeUint8(uint8(m.Dis))
	buf.EncodeUint8(uint8(m.FaultDet))
	buf.EncodeBytes(m.Reserve[:], 3)
	return buf.Bytes(), nil
}
func (m *UDPPingAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SrcIPAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.SrcIPAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.DstIPAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.DstIPAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.StartSrcPort = buf.DecodeUint16()
	m.EndSrcPort = buf.DecodeUint16()
	m.StartDstPort = buf.DecodeUint16()
	m.EndDstPort = buf.DecodeUint16()
	m.Interval = buf.DecodeUint16()
	m.Dis = buf.DecodeUint8()
	m.FaultDet = buf.DecodeUint8()
	copy(m.Reserve[:], buf.DecodeBytes(3))
	return nil
}

// UDPPingAddDelReply defines message 'udp_ping_add_del_reply'.
type UDPPingAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *UDPPingAddDelReply) Reset()               { *m = UDPPingAddDelReply{} }
func (*UDPPingAddDelReply) GetMessageName() string { return "udp_ping_add_del_reply" }
func (*UDPPingAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*UDPPingAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *UDPPingAddDelReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Retval
	return size
}
func (m *UDPPingAddDelReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Retval))
	return buf.Bytes(), nil
}
func (m *UDPPingAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = int32(buf.DecodeUint32())
	return nil
}

// UDPPingExport defines message 'udp_ping_export'.
type UDPPingExport struct {
	Enable bool `binapi:"bool,name=enable" json:"enable,omitempty"`
}

func (m *UDPPingExport) Reset()               { *m = UDPPingExport{} }
func (*UDPPingExport) GetMessageName() string { return "udp_ping_export" }
func (*UDPPingExport) GetCrcString() string   { return "b3e225d2" }
func (*UDPPingExport) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *UDPPingExport) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 1 // m.Enable
	return size
}
func (m *UDPPingExport) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeBool(m.Enable)
	return buf.Bytes(), nil
}
func (m *UDPPingExport) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Enable = buf.DecodeBool()
	return nil
}

// UDPPingExportReply defines message 'udp_ping_export_reply'.
type UDPPingExportReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *UDPPingExportReply) Reset()               { *m = UDPPingExportReply{} }
func (*UDPPingExportReply) GetMessageName() string { return "udp_ping_export_reply" }
func (*UDPPingExportReply) GetCrcString() string   { return "e8d4e804" }
func (*UDPPingExportReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *UDPPingExportReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Retval
	return size
}
func (m *UDPPingExportReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Retval))
	return buf.Bytes(), nil
}
func (m *UDPPingExportReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = int32(buf.DecodeUint32())
	return nil
}

func init() { file_udp_ping_binapi_init() }
func file_udp_ping_binapi_init() {
	api.RegisterMessage((*UDPPingAddDel)(nil), "udp_ping_add_del_c692b188")
	api.RegisterMessage((*UDPPingAddDelReply)(nil), "udp_ping_add_del_reply_e8d4e804")
	api.RegisterMessage((*UDPPingExport)(nil), "udp_ping_export_b3e225d2")
	api.RegisterMessage((*UDPPingExportReply)(nil), "udp_ping_export_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*UDPPingAddDel)(nil),
		(*UDPPingAddDelReply)(nil),
		(*UDPPingExport)(nil),
		(*UDPPingExportReply)(nil),
	}
}
