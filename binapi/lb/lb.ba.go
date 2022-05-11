// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              21.06-release
// source: /usr/share/vpp/api/plugins/lb.api.json

// Package lb contains generated bindings for API file lb.api.
//
// Contents:
//  16 messages
//
package lb

import (
	api "git.fd.io/govpp.git/api"
	interface_types "git.fd.io/govpp.git/binapi/interface_types"
	ip_types "git.fd.io/govpp.git/binapi/ip_types"
	lb_types "git.fd.io/govpp.git/binapi/lb_types"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

// LbAddDelAs defines message 'lb_add_del_as'.
type LbAddDelAs struct {
	Pfx       ip_types.AddressWithPrefix `binapi:"address_with_prefix,name=pfx" json:"pfx,omitempty"`
	Protocol  uint8                      `binapi:"u8,name=protocol,default=255" json:"protocol,omitempty"`
	Port      uint16                     `binapi:"u16,name=port" json:"port,omitempty"`
	AsAddress ip_types.Address           `binapi:"address,name=as_address" json:"as_address,omitempty"`
	IsDel     bool                       `binapi:"bool,name=is_del" json:"is_del,omitempty"`
	IsFlush   bool                       `binapi:"bool,name=is_flush" json:"is_flush,omitempty"`
}

func (m *LbAddDelAs) Reset()               { *m = LbAddDelAs{} }
func (*LbAddDelAs) GetMessageName() string { return "lb_add_del_as" }
func (*LbAddDelAs) GetCrcString() string   { return "35d72500" }
func (*LbAddDelAs) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LbAddDelAs) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Pfx.Address.Af
	size += 1 * 16 // m.Pfx.Address.Un
	size += 1      // m.Pfx.Len
	size += 1      // m.Protocol
	size += 2      // m.Port
	size += 1      // m.AsAddress.Af
	size += 1 * 16 // m.AsAddress.Un
	size += 1      // m.IsDel
	size += 1      // m.IsFlush
	return size
}
func (m *LbAddDelAs) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Pfx.Address.Af))
	buf.EncodeBytes(m.Pfx.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Pfx.Len)
	buf.EncodeUint8(m.Protocol)
	buf.EncodeUint16(m.Port)
	buf.EncodeUint8(uint8(m.AsAddress.Af))
	buf.EncodeBytes(m.AsAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeBool(m.IsDel)
	buf.EncodeBool(m.IsFlush)
	return buf.Bytes(), nil
}
func (m *LbAddDelAs) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Pfx.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Pfx.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Pfx.Len = buf.DecodeUint8()
	m.Protocol = buf.DecodeUint8()
	m.Port = buf.DecodeUint16()
	m.AsAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.AsAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.IsDel = buf.DecodeBool()
	m.IsFlush = buf.DecodeBool()
	return nil
}

// LbAddDelAsReply defines message 'lb_add_del_as_reply'.
type LbAddDelAsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *LbAddDelAsReply) Reset()               { *m = LbAddDelAsReply{} }
func (*LbAddDelAsReply) GetMessageName() string { return "lb_add_del_as_reply" }
func (*LbAddDelAsReply) GetCrcString() string   { return "e8d4e804" }
func (*LbAddDelAsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LbAddDelAsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *LbAddDelAsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *LbAddDelAsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// LbAddDelIntfNat4 defines message 'lb_add_del_intf_nat4'.
type LbAddDelIntfNat4 struct {
	IsAdd     bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *LbAddDelIntfNat4) Reset()               { *m = LbAddDelIntfNat4{} }
func (*LbAddDelIntfNat4) GetMessageName() string { return "lb_add_del_intf_nat4" }
func (*LbAddDelIntfNat4) GetCrcString() string   { return "47d6e753" }
func (*LbAddDelIntfNat4) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LbAddDelIntfNat4) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 4 // m.SwIfIndex
	return size
}
func (m *LbAddDelIntfNat4) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *LbAddDelIntfNat4) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// LbAddDelIntfNat4Reply defines message 'lb_add_del_intf_nat4_reply'.
type LbAddDelIntfNat4Reply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *LbAddDelIntfNat4Reply) Reset()               { *m = LbAddDelIntfNat4Reply{} }
func (*LbAddDelIntfNat4Reply) GetMessageName() string { return "lb_add_del_intf_nat4_reply" }
func (*LbAddDelIntfNat4Reply) GetCrcString() string   { return "e8d4e804" }
func (*LbAddDelIntfNat4Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LbAddDelIntfNat4Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *LbAddDelIntfNat4Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *LbAddDelIntfNat4Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// LbAddDelIntfNat6 defines message 'lb_add_del_intf_nat6'.
type LbAddDelIntfNat6 struct {
	IsAdd     bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *LbAddDelIntfNat6) Reset()               { *m = LbAddDelIntfNat6{} }
func (*LbAddDelIntfNat6) GetMessageName() string { return "lb_add_del_intf_nat6" }
func (*LbAddDelIntfNat6) GetCrcString() string   { return "47d6e753" }
func (*LbAddDelIntfNat6) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LbAddDelIntfNat6) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 4 // m.SwIfIndex
	return size
}
func (m *LbAddDelIntfNat6) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *LbAddDelIntfNat6) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// LbAddDelIntfNat6Reply defines message 'lb_add_del_intf_nat6_reply'.
type LbAddDelIntfNat6Reply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *LbAddDelIntfNat6Reply) Reset()               { *m = LbAddDelIntfNat6Reply{} }
func (*LbAddDelIntfNat6Reply) GetMessageName() string { return "lb_add_del_intf_nat6_reply" }
func (*LbAddDelIntfNat6Reply) GetCrcString() string   { return "e8d4e804" }
func (*LbAddDelIntfNat6Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LbAddDelIntfNat6Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *LbAddDelIntfNat6Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *LbAddDelIntfNat6Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// LbAddDelVip defines message 'lb_add_del_vip'.
type LbAddDelVip struct {
	Pfx                 ip_types.AddressWithPrefix `binapi:"address_with_prefix,name=pfx" json:"pfx,omitempty"`
	Protocol            uint8                      `binapi:"u8,name=protocol,default=255" json:"protocol,omitempty"`
	Port                uint16                     `binapi:"u16,name=port" json:"port,omitempty"`
	Encap               lb_types.LbEncapType       `binapi:"lb_encap_type,name=encap" json:"encap,omitempty"`
	Dscp                uint8                      `binapi:"u8,name=dscp" json:"dscp,omitempty"`
	Type                lb_types.LbSrvType         `binapi:"lb_srv_type,name=type" json:"type,omitempty"`
	TargetPort          uint16                     `binapi:"u16,name=target_port" json:"target_port,omitempty"`
	NodePort            uint16                     `binapi:"u16,name=node_port" json:"node_port,omitempty"`
	NewFlowsTableLength uint32                     `binapi:"u32,name=new_flows_table_length,default=1024" json:"new_flows_table_length,omitempty"`
	IsDel               bool                       `binapi:"bool,name=is_del" json:"is_del,omitempty"`
}

func (m *LbAddDelVip) Reset()               { *m = LbAddDelVip{} }
func (*LbAddDelVip) GetMessageName() string { return "lb_add_del_vip" }
func (*LbAddDelVip) GetCrcString() string   { return "6fa569c7" }
func (*LbAddDelVip) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LbAddDelVip) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Pfx.Address.Af
	size += 1 * 16 // m.Pfx.Address.Un
	size += 1      // m.Pfx.Len
	size += 1      // m.Protocol
	size += 2      // m.Port
	size += 4      // m.Encap
	size += 1      // m.Dscp
	size += 4      // m.Type
	size += 2      // m.TargetPort
	size += 2      // m.NodePort
	size += 4      // m.NewFlowsTableLength
	size += 1      // m.IsDel
	return size
}
func (m *LbAddDelVip) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Pfx.Address.Af))
	buf.EncodeBytes(m.Pfx.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Pfx.Len)
	buf.EncodeUint8(m.Protocol)
	buf.EncodeUint16(m.Port)
	buf.EncodeUint32(uint32(m.Encap))
	buf.EncodeUint8(m.Dscp)
	buf.EncodeUint32(uint32(m.Type))
	buf.EncodeUint16(m.TargetPort)
	buf.EncodeUint16(m.NodePort)
	buf.EncodeUint32(m.NewFlowsTableLength)
	buf.EncodeBool(m.IsDel)
	return buf.Bytes(), nil
}
func (m *LbAddDelVip) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Pfx.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Pfx.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Pfx.Len = buf.DecodeUint8()
	m.Protocol = buf.DecodeUint8()
	m.Port = buf.DecodeUint16()
	m.Encap = lb_types.LbEncapType(buf.DecodeUint32())
	m.Dscp = buf.DecodeUint8()
	m.Type = lb_types.LbSrvType(buf.DecodeUint32())
	m.TargetPort = buf.DecodeUint16()
	m.NodePort = buf.DecodeUint16()
	m.NewFlowsTableLength = buf.DecodeUint32()
	m.IsDel = buf.DecodeBool()
	return nil
}

// LbAddDelVipReply defines message 'lb_add_del_vip_reply'.
type LbAddDelVipReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *LbAddDelVipReply) Reset()               { *m = LbAddDelVipReply{} }
func (*LbAddDelVipReply) GetMessageName() string { return "lb_add_del_vip_reply" }
func (*LbAddDelVipReply) GetCrcString() string   { return "e8d4e804" }
func (*LbAddDelVipReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LbAddDelVipReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *LbAddDelVipReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *LbAddDelVipReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// LbAsDetails defines message 'lb_as_details'.
type LbAsDetails struct {
	Vip        lb_types.LbVip   `binapi:"lb_vip,name=vip" json:"vip,omitempty"`
	AppSrv     ip_types.Address `binapi:"address,name=app_srv" json:"app_srv,omitempty"`
	Flags      uint8            `binapi:"u8,name=flags" json:"flags,omitempty"`
	InUseSince uint32           `binapi:"u32,name=in_use_since" json:"in_use_since,omitempty"`
}

func (m *LbAsDetails) Reset()               { *m = LbAsDetails{} }
func (*LbAsDetails) GetMessageName() string { return "lb_as_details" }
func (*LbAsDetails) GetCrcString() string   { return "8d24c29e" }
func (*LbAsDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LbAsDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Vip.Pfx.Address.Af
	size += 1 * 16 // m.Vip.Pfx.Address.Un
	size += 1      // m.Vip.Pfx.Len
	size += 1      // m.Vip.Protocol
	size += 2      // m.Vip.Port
	size += 1      // m.AppSrv.Af
	size += 1 * 16 // m.AppSrv.Un
	size += 1      // m.Flags
	size += 4      // m.InUseSince
	return size
}
func (m *LbAsDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Vip.Pfx.Address.Af))
	buf.EncodeBytes(m.Vip.Pfx.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Vip.Pfx.Len)
	buf.EncodeUint8(uint8(m.Vip.Protocol))
	buf.EncodeUint16(m.Vip.Port)
	buf.EncodeUint8(uint8(m.AppSrv.Af))
	buf.EncodeBytes(m.AppSrv.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Flags)
	buf.EncodeUint32(m.InUseSince)
	return buf.Bytes(), nil
}
func (m *LbAsDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Vip.Pfx.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Vip.Pfx.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Vip.Pfx.Len = buf.DecodeUint8()
	m.Vip.Protocol = ip_types.IPProto(buf.DecodeUint8())
	m.Vip.Port = buf.DecodeUint16()
	m.AppSrv.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.AppSrv.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Flags = buf.DecodeUint8()
	m.InUseSince = buf.DecodeUint32()
	return nil
}

// LbAsDump defines message 'lb_as_dump'.
type LbAsDump struct {
	Pfx      ip_types.AddressWithPrefix `binapi:"address_with_prefix,name=pfx" json:"pfx,omitempty"`
	Protocol uint8                      `binapi:"u8,name=protocol" json:"protocol,omitempty"`
	Port     uint16                     `binapi:"u16,name=port" json:"port,omitempty"`
}

func (m *LbAsDump) Reset()               { *m = LbAsDump{} }
func (*LbAsDump) GetMessageName() string { return "lb_as_dump" }
func (*LbAsDump) GetCrcString() string   { return "1063f819" }
func (*LbAsDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LbAsDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Pfx.Address.Af
	size += 1 * 16 // m.Pfx.Address.Un
	size += 1      // m.Pfx.Len
	size += 1      // m.Protocol
	size += 2      // m.Port
	return size
}
func (m *LbAsDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Pfx.Address.Af))
	buf.EncodeBytes(m.Pfx.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Pfx.Len)
	buf.EncodeUint8(m.Protocol)
	buf.EncodeUint16(m.Port)
	return buf.Bytes(), nil
}
func (m *LbAsDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Pfx.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Pfx.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Pfx.Len = buf.DecodeUint8()
	m.Protocol = buf.DecodeUint8()
	m.Port = buf.DecodeUint16()
	return nil
}

// LbConf defines message 'lb_conf'.
type LbConf struct {
	IP4SrcAddress        ip_types.IP4Address `binapi:"ip4_address,name=ip4_src_address" json:"ip4_src_address,omitempty"`
	IP6SrcAddress        ip_types.IP6Address `binapi:"ip6_address,name=ip6_src_address" json:"ip6_src_address,omitempty"`
	StickyBucketsPerCore uint32              `binapi:"u32,name=sticky_buckets_per_core,default=4294967295" json:"sticky_buckets_per_core,omitempty"`
	FlowTimeout          uint32              `binapi:"u32,name=flow_timeout,default=4294967295" json:"flow_timeout,omitempty"`
}

func (m *LbConf) Reset()               { *m = LbConf{} }
func (*LbConf) GetMessageName() string { return "lb_conf" }
func (*LbConf) GetCrcString() string   { return "56cd3261" }
func (*LbConf) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LbConf) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 4  // m.IP4SrcAddress
	size += 1 * 16 // m.IP6SrcAddress
	size += 4      // m.StickyBucketsPerCore
	size += 4      // m.FlowTimeout
	return size
}
func (m *LbConf) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.IP4SrcAddress[:], 4)
	buf.EncodeBytes(m.IP6SrcAddress[:], 16)
	buf.EncodeUint32(m.StickyBucketsPerCore)
	buf.EncodeUint32(m.FlowTimeout)
	return buf.Bytes(), nil
}
func (m *LbConf) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.IP4SrcAddress[:], buf.DecodeBytes(4))
	copy(m.IP6SrcAddress[:], buf.DecodeBytes(16))
	m.StickyBucketsPerCore = buf.DecodeUint32()
	m.FlowTimeout = buf.DecodeUint32()
	return nil
}

// LbConfReply defines message 'lb_conf_reply'.
type LbConfReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *LbConfReply) Reset()               { *m = LbConfReply{} }
func (*LbConfReply) GetMessageName() string { return "lb_conf_reply" }
func (*LbConfReply) GetCrcString() string   { return "e8d4e804" }
func (*LbConfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LbConfReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *LbConfReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *LbConfReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// LbFlushVip defines message 'lb_flush_vip'.
type LbFlushVip struct {
	Pfx      ip_types.AddressWithPrefix `binapi:"address_with_prefix,name=pfx" json:"pfx,omitempty"`
	Protocol uint8                      `binapi:"u8,name=protocol" json:"protocol,omitempty"`
	Port     uint16                     `binapi:"u16,name=port" json:"port,omitempty"`
}

func (m *LbFlushVip) Reset()               { *m = LbFlushVip{} }
func (*LbFlushVip) GetMessageName() string { return "lb_flush_vip" }
func (*LbFlushVip) GetCrcString() string   { return "1063f819" }
func (*LbFlushVip) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LbFlushVip) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Pfx.Address.Af
	size += 1 * 16 // m.Pfx.Address.Un
	size += 1      // m.Pfx.Len
	size += 1      // m.Protocol
	size += 2      // m.Port
	return size
}
func (m *LbFlushVip) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Pfx.Address.Af))
	buf.EncodeBytes(m.Pfx.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Pfx.Len)
	buf.EncodeUint8(m.Protocol)
	buf.EncodeUint16(m.Port)
	return buf.Bytes(), nil
}
func (m *LbFlushVip) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Pfx.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Pfx.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Pfx.Len = buf.DecodeUint8()
	m.Protocol = buf.DecodeUint8()
	m.Port = buf.DecodeUint16()
	return nil
}

// LbFlushVipReply defines message 'lb_flush_vip_reply'.
type LbFlushVipReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *LbFlushVipReply) Reset()               { *m = LbFlushVipReply{} }
func (*LbFlushVipReply) GetMessageName() string { return "lb_flush_vip_reply" }
func (*LbFlushVipReply) GetCrcString() string   { return "e8d4e804" }
func (*LbFlushVipReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LbFlushVipReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *LbFlushVipReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *LbFlushVipReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// LbVipDetails defines message 'lb_vip_details'.
type LbVipDetails struct {
	Vip             lb_types.LbVip       `binapi:"lb_vip,name=vip" json:"vip,omitempty"`
	Encap           lb_types.LbEncapType `binapi:"lb_encap_type,name=encap" json:"encap,omitempty"`
	Dscp            ip_types.IPDscp      `binapi:"ip_dscp,name=dscp" json:"dscp,omitempty"`
	SrvType         lb_types.LbSrvType   `binapi:"lb_srv_type,name=srv_type" json:"srv_type,omitempty"`
	TargetPort      uint16               `binapi:"u16,name=target_port" json:"target_port,omitempty"`
	FlowTableLength uint16               `binapi:"u16,name=flow_table_length" json:"flow_table_length,omitempty"`
}

func (m *LbVipDetails) Reset()               { *m = LbVipDetails{} }
func (*LbVipDetails) GetMessageName() string { return "lb_vip_details" }
func (*LbVipDetails) GetCrcString() string   { return "1329ec9b" }
func (*LbVipDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LbVipDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Vip.Pfx.Address.Af
	size += 1 * 16 // m.Vip.Pfx.Address.Un
	size += 1      // m.Vip.Pfx.Len
	size += 1      // m.Vip.Protocol
	size += 2      // m.Vip.Port
	size += 4      // m.Encap
	size += 1      // m.Dscp
	size += 4      // m.SrvType
	size += 2      // m.TargetPort
	size += 2      // m.FlowTableLength
	return size
}
func (m *LbVipDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Vip.Pfx.Address.Af))
	buf.EncodeBytes(m.Vip.Pfx.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Vip.Pfx.Len)
	buf.EncodeUint8(uint8(m.Vip.Protocol))
	buf.EncodeUint16(m.Vip.Port)
	buf.EncodeUint32(uint32(m.Encap))
	buf.EncodeUint8(uint8(m.Dscp))
	buf.EncodeUint32(uint32(m.SrvType))
	buf.EncodeUint16(m.TargetPort)
	buf.EncodeUint16(m.FlowTableLength)
	return buf.Bytes(), nil
}
func (m *LbVipDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Vip.Pfx.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Vip.Pfx.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Vip.Pfx.Len = buf.DecodeUint8()
	m.Vip.Protocol = ip_types.IPProto(buf.DecodeUint8())
	m.Vip.Port = buf.DecodeUint16()
	m.Encap = lb_types.LbEncapType(buf.DecodeUint32())
	m.Dscp = ip_types.IPDscp(buf.DecodeUint8())
	m.SrvType = lb_types.LbSrvType(buf.DecodeUint32())
	m.TargetPort = buf.DecodeUint16()
	m.FlowTableLength = buf.DecodeUint16()
	return nil
}

// LbVipDump defines message 'lb_vip_dump'.
type LbVipDump struct {
	Pfx        ip_types.AddressWithPrefix `binapi:"address_with_prefix,name=pfx" json:"pfx,omitempty"`
	PfxMatcher ip_types.PrefixMatcher     `binapi:"prefix_matcher,name=pfx_matcher" json:"pfx_matcher,omitempty"`
	Protocol   uint8                      `binapi:"u8,name=protocol,default=255" json:"protocol,omitempty"`
	Port       uint16                     `binapi:"u16,name=port" json:"port,omitempty"`
}

func (m *LbVipDump) Reset()               { *m = LbVipDump{} }
func (*LbVipDump) GetMessageName() string { return "lb_vip_dump" }
func (*LbVipDump) GetCrcString() string   { return "56110cb7" }
func (*LbVipDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LbVipDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Pfx.Address.Af
	size += 1 * 16 // m.Pfx.Address.Un
	size += 1      // m.Pfx.Len
	size += 1      // m.PfxMatcher.Le
	size += 1      // m.PfxMatcher.Ge
	size += 1      // m.Protocol
	size += 2      // m.Port
	return size
}
func (m *LbVipDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Pfx.Address.Af))
	buf.EncodeBytes(m.Pfx.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Pfx.Len)
	buf.EncodeUint8(m.PfxMatcher.Le)
	buf.EncodeUint8(m.PfxMatcher.Ge)
	buf.EncodeUint8(m.Protocol)
	buf.EncodeUint16(m.Port)
	return buf.Bytes(), nil
}
func (m *LbVipDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Pfx.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Pfx.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Pfx.Len = buf.DecodeUint8()
	m.PfxMatcher.Le = buf.DecodeUint8()
	m.PfxMatcher.Ge = buf.DecodeUint8()
	m.Protocol = buf.DecodeUint8()
	m.Port = buf.DecodeUint16()
	return nil
}

func init() { file_lb_binapi_init() }
func file_lb_binapi_init() {
	api.RegisterMessage((*LbAddDelAs)(nil), "lb_add_del_as_35d72500")
	api.RegisterMessage((*LbAddDelAsReply)(nil), "lb_add_del_as_reply_e8d4e804")
	api.RegisterMessage((*LbAddDelIntfNat4)(nil), "lb_add_del_intf_nat4_47d6e753")
	api.RegisterMessage((*LbAddDelIntfNat4Reply)(nil), "lb_add_del_intf_nat4_reply_e8d4e804")
	api.RegisterMessage((*LbAddDelIntfNat6)(nil), "lb_add_del_intf_nat6_47d6e753")
	api.RegisterMessage((*LbAddDelIntfNat6Reply)(nil), "lb_add_del_intf_nat6_reply_e8d4e804")
	api.RegisterMessage((*LbAddDelVip)(nil), "lb_add_del_vip_6fa569c7")
	api.RegisterMessage((*LbAddDelVipReply)(nil), "lb_add_del_vip_reply_e8d4e804")
	api.RegisterMessage((*LbAsDetails)(nil), "lb_as_details_8d24c29e")
	api.RegisterMessage((*LbAsDump)(nil), "lb_as_dump_1063f819")
	api.RegisterMessage((*LbConf)(nil), "lb_conf_56cd3261")
	api.RegisterMessage((*LbConfReply)(nil), "lb_conf_reply_e8d4e804")
	api.RegisterMessage((*LbFlushVip)(nil), "lb_flush_vip_1063f819")
	api.RegisterMessage((*LbFlushVipReply)(nil), "lb_flush_vip_reply_e8d4e804")
	api.RegisterMessage((*LbVipDetails)(nil), "lb_vip_details_1329ec9b")
	api.RegisterMessage((*LbVipDump)(nil), "lb_vip_dump_56110cb7")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*LbAddDelAs)(nil),
		(*LbAddDelAsReply)(nil),
		(*LbAddDelIntfNat4)(nil),
		(*LbAddDelIntfNat4Reply)(nil),
		(*LbAddDelIntfNat6)(nil),
		(*LbAddDelIntfNat6Reply)(nil),
		(*LbAddDelVip)(nil),
		(*LbAddDelVipReply)(nil),
		(*LbAsDetails)(nil),
		(*LbAsDump)(nil),
		(*LbConf)(nil),
		(*LbConfReply)(nil),
		(*LbFlushVip)(nil),
		(*LbFlushVipReply)(nil),
		(*LbVipDetails)(nil),
		(*LbVipDump)(nil),
	}
}
