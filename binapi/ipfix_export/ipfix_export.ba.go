// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.6.0-dev
//  VPP:              22.02-release
// source: /usr/share/vpp/api/core/ipfix_export.api.json

// Package ipfix_export contains generated bindings for API file ipfix_export.api.
//
// Contents:
//  19 messages
//
package ipfix_export

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
	APIFile    = "ipfix_export"
	APIVersion = "2.0.3"
	VersionCrc = 0x63e0810a
)

// IpfixAllExporterDetails defines message 'ipfix_all_exporter_details'.
type IpfixAllExporterDetails struct {
	CollectorAddress ip_types.Address `binapi:"address,name=collector_address" json:"collector_address,omitempty"`
	CollectorPort    uint16           `binapi:"u16,name=collector_port" json:"collector_port,omitempty"`
	SrcAddress       ip_types.Address `binapi:"address,name=src_address" json:"src_address,omitempty"`
	VrfID            uint32           `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
	PathMtu          uint32           `binapi:"u32,name=path_mtu" json:"path_mtu,omitempty"`
	TemplateInterval uint32           `binapi:"u32,name=template_interval" json:"template_interval,omitempty"`
	UDPChecksum      bool             `binapi:"bool,name=udp_checksum" json:"udp_checksum,omitempty"`
}

func (m *IpfixAllExporterDetails) Reset()               { *m = IpfixAllExporterDetails{} }
func (*IpfixAllExporterDetails) GetMessageName() string { return "ipfix_all_exporter_details" }
func (*IpfixAllExporterDetails) GetCrcString() string   { return "0dedbfe4" }
func (*IpfixAllExporterDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IpfixAllExporterDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.CollectorAddress.Af
	size += 1 * 16 // m.CollectorAddress.Un
	size += 2      // m.CollectorPort
	size += 1      // m.SrcAddress.Af
	size += 1 * 16 // m.SrcAddress.Un
	size += 4      // m.VrfID
	size += 4      // m.PathMtu
	size += 4      // m.TemplateInterval
	size += 1      // m.UDPChecksum
	return size
}
func (m *IpfixAllExporterDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.CollectorAddress.Af))
	buf.EncodeBytes(m.CollectorAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint16(m.CollectorPort)
	buf.EncodeUint8(uint8(m.SrcAddress.Af))
	buf.EncodeBytes(m.SrcAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.VrfID)
	buf.EncodeUint32(m.PathMtu)
	buf.EncodeUint32(m.TemplateInterval)
	buf.EncodeBool(m.UDPChecksum)
	return buf.Bytes(), nil
}
func (m *IpfixAllExporterDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.CollectorAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.CollectorAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.CollectorPort = buf.DecodeUint16()
	m.SrcAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.SrcAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.VrfID = buf.DecodeUint32()
	m.PathMtu = buf.DecodeUint32()
	m.TemplateInterval = buf.DecodeUint32()
	m.UDPChecksum = buf.DecodeBool()
	return nil
}

// IpfixAllExporterGet defines message 'ipfix_all_exporter_get'.
type IpfixAllExporterGet struct {
	Cursor uint32 `binapi:"u32,name=cursor" json:"cursor,omitempty"`
}

func (m *IpfixAllExporterGet) Reset()               { *m = IpfixAllExporterGet{} }
func (*IpfixAllExporterGet) GetMessageName() string { return "ipfix_all_exporter_get" }
func (*IpfixAllExporterGet) GetCrcString() string   { return "f75ba505" }
func (*IpfixAllExporterGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IpfixAllExporterGet) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Cursor
	return size
}
func (m *IpfixAllExporterGet) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Cursor)
	return buf.Bytes(), nil
}
func (m *IpfixAllExporterGet) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Cursor = buf.DecodeUint32()
	return nil
}

// IpfixAllExporterGetReply defines message 'ipfix_all_exporter_get_reply'.
type IpfixAllExporterGetReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	Cursor uint32 `binapi:"u32,name=cursor" json:"cursor,omitempty"`
}

func (m *IpfixAllExporterGetReply) Reset()               { *m = IpfixAllExporterGetReply{} }
func (*IpfixAllExporterGetReply) GetMessageName() string { return "ipfix_all_exporter_get_reply" }
func (*IpfixAllExporterGetReply) GetCrcString() string   { return "53b48f5d" }
func (*IpfixAllExporterGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IpfixAllExporterGetReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.Cursor
	return size
}
func (m *IpfixAllExporterGetReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.Cursor)
	return buf.Bytes(), nil
}
func (m *IpfixAllExporterGetReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Cursor = buf.DecodeUint32()
	return nil
}

// IpfixClassifyStreamDetails defines message 'ipfix_classify_stream_details'.
type IpfixClassifyStreamDetails struct {
	DomainID uint32 `binapi:"u32,name=domain_id" json:"domain_id,omitempty"`
	SrcPort  uint16 `binapi:"u16,name=src_port" json:"src_port,omitempty"`
}

func (m *IpfixClassifyStreamDetails) Reset()               { *m = IpfixClassifyStreamDetails{} }
func (*IpfixClassifyStreamDetails) GetMessageName() string { return "ipfix_classify_stream_details" }
func (*IpfixClassifyStreamDetails) GetCrcString() string   { return "2903539d" }
func (*IpfixClassifyStreamDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IpfixClassifyStreamDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.DomainID
	size += 2 // m.SrcPort
	return size
}
func (m *IpfixClassifyStreamDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.DomainID)
	buf.EncodeUint16(m.SrcPort)
	return buf.Bytes(), nil
}
func (m *IpfixClassifyStreamDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.DomainID = buf.DecodeUint32()
	m.SrcPort = buf.DecodeUint16()
	return nil
}

// IpfixClassifyStreamDump defines message 'ipfix_classify_stream_dump'.
type IpfixClassifyStreamDump struct{}

func (m *IpfixClassifyStreamDump) Reset()               { *m = IpfixClassifyStreamDump{} }
func (*IpfixClassifyStreamDump) GetMessageName() string { return "ipfix_classify_stream_dump" }
func (*IpfixClassifyStreamDump) GetCrcString() string   { return "51077d14" }
func (*IpfixClassifyStreamDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IpfixClassifyStreamDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *IpfixClassifyStreamDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *IpfixClassifyStreamDump) Unmarshal(b []byte) error {
	return nil
}

// IpfixClassifyTableAddDel defines message 'ipfix_classify_table_add_del'.
type IpfixClassifyTableAddDel struct {
	TableID           uint32                 `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	IPVersion         ip_types.AddressFamily `binapi:"address_family,name=ip_version" json:"ip_version,omitempty"`
	TransportProtocol ip_types.IPProto       `binapi:"ip_proto,name=transport_protocol" json:"transport_protocol,omitempty"`
	IsAdd             bool                   `binapi:"bool,name=is_add" json:"is_add,omitempty"`
}

func (m *IpfixClassifyTableAddDel) Reset()               { *m = IpfixClassifyTableAddDel{} }
func (*IpfixClassifyTableAddDel) GetMessageName() string { return "ipfix_classify_table_add_del" }
func (*IpfixClassifyTableAddDel) GetCrcString() string   { return "3e449bb9" }
func (*IpfixClassifyTableAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IpfixClassifyTableAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.TableID
	size += 1 // m.IPVersion
	size += 1 // m.TransportProtocol
	size += 1 // m.IsAdd
	return size
}
func (m *IpfixClassifyTableAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.TableID)
	buf.EncodeUint8(uint8(m.IPVersion))
	buf.EncodeUint8(uint8(m.TransportProtocol))
	buf.EncodeBool(m.IsAdd)
	return buf.Bytes(), nil
}
func (m *IpfixClassifyTableAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TableID = buf.DecodeUint32()
	m.IPVersion = ip_types.AddressFamily(buf.DecodeUint8())
	m.TransportProtocol = ip_types.IPProto(buf.DecodeUint8())
	m.IsAdd = buf.DecodeBool()
	return nil
}

// IpfixClassifyTableAddDelReply defines message 'ipfix_classify_table_add_del_reply'.
type IpfixClassifyTableAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IpfixClassifyTableAddDelReply) Reset() { *m = IpfixClassifyTableAddDelReply{} }
func (*IpfixClassifyTableAddDelReply) GetMessageName() string {
	return "ipfix_classify_table_add_del_reply"
}
func (*IpfixClassifyTableAddDelReply) GetCrcString() string { return "e8d4e804" }
func (*IpfixClassifyTableAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IpfixClassifyTableAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IpfixClassifyTableAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IpfixClassifyTableAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IpfixClassifyTableDetails defines message 'ipfix_classify_table_details'.
type IpfixClassifyTableDetails struct {
	TableID           uint32                 `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	IPVersion         ip_types.AddressFamily `binapi:"address_family,name=ip_version" json:"ip_version,omitempty"`
	TransportProtocol ip_types.IPProto       `binapi:"ip_proto,name=transport_protocol" json:"transport_protocol,omitempty"`
}

func (m *IpfixClassifyTableDetails) Reset()               { *m = IpfixClassifyTableDetails{} }
func (*IpfixClassifyTableDetails) GetMessageName() string { return "ipfix_classify_table_details" }
func (*IpfixClassifyTableDetails) GetCrcString() string   { return "1af8c28c" }
func (*IpfixClassifyTableDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IpfixClassifyTableDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.TableID
	size += 1 // m.IPVersion
	size += 1 // m.TransportProtocol
	return size
}
func (m *IpfixClassifyTableDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.TableID)
	buf.EncodeUint8(uint8(m.IPVersion))
	buf.EncodeUint8(uint8(m.TransportProtocol))
	return buf.Bytes(), nil
}
func (m *IpfixClassifyTableDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TableID = buf.DecodeUint32()
	m.IPVersion = ip_types.AddressFamily(buf.DecodeUint8())
	m.TransportProtocol = ip_types.IPProto(buf.DecodeUint8())
	return nil
}

// IpfixClassifyTableDump defines message 'ipfix_classify_table_dump'.
type IpfixClassifyTableDump struct{}

func (m *IpfixClassifyTableDump) Reset()               { *m = IpfixClassifyTableDump{} }
func (*IpfixClassifyTableDump) GetMessageName() string { return "ipfix_classify_table_dump" }
func (*IpfixClassifyTableDump) GetCrcString() string   { return "51077d14" }
func (*IpfixClassifyTableDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IpfixClassifyTableDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *IpfixClassifyTableDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *IpfixClassifyTableDump) Unmarshal(b []byte) error {
	return nil
}

// IpfixExporterCreateDelete defines message 'ipfix_exporter_create_delete'.
type IpfixExporterCreateDelete struct {
	IsCreate         bool             `binapi:"bool,name=is_create" json:"is_create,omitempty"`
	CollectorAddress ip_types.Address `binapi:"address,name=collector_address" json:"collector_address,omitempty"`
	CollectorPort    uint16           `binapi:"u16,name=collector_port" json:"collector_port,omitempty"`
	SrcAddress       ip_types.Address `binapi:"address,name=src_address" json:"src_address,omitempty"`
	VrfID            uint32           `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
	PathMtu          uint32           `binapi:"u32,name=path_mtu" json:"path_mtu,omitempty"`
	TemplateInterval uint32           `binapi:"u32,name=template_interval" json:"template_interval,omitempty"`
	UDPChecksum      bool             `binapi:"bool,name=udp_checksum" json:"udp_checksum,omitempty"`
}

func (m *IpfixExporterCreateDelete) Reset()               { *m = IpfixExporterCreateDelete{} }
func (*IpfixExporterCreateDelete) GetMessageName() string { return "ipfix_exporter_create_delete" }
func (*IpfixExporterCreateDelete) GetCrcString() string   { return "0753a768" }
func (*IpfixExporterCreateDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IpfixExporterCreateDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsCreate
	size += 1      // m.CollectorAddress.Af
	size += 1 * 16 // m.CollectorAddress.Un
	size += 2      // m.CollectorPort
	size += 1      // m.SrcAddress.Af
	size += 1 * 16 // m.SrcAddress.Un
	size += 4      // m.VrfID
	size += 4      // m.PathMtu
	size += 4      // m.TemplateInterval
	size += 1      // m.UDPChecksum
	return size
}
func (m *IpfixExporterCreateDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsCreate)
	buf.EncodeUint8(uint8(m.CollectorAddress.Af))
	buf.EncodeBytes(m.CollectorAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint16(m.CollectorPort)
	buf.EncodeUint8(uint8(m.SrcAddress.Af))
	buf.EncodeBytes(m.SrcAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.VrfID)
	buf.EncodeUint32(m.PathMtu)
	buf.EncodeUint32(m.TemplateInterval)
	buf.EncodeBool(m.UDPChecksum)
	return buf.Bytes(), nil
}
func (m *IpfixExporterCreateDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsCreate = buf.DecodeBool()
	m.CollectorAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.CollectorAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.CollectorPort = buf.DecodeUint16()
	m.SrcAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.SrcAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.VrfID = buf.DecodeUint32()
	m.PathMtu = buf.DecodeUint32()
	m.TemplateInterval = buf.DecodeUint32()
	m.UDPChecksum = buf.DecodeBool()
	return nil
}

// IpfixExporterCreateDeleteReply defines message 'ipfix_exporter_create_delete_reply'.
type IpfixExporterCreateDeleteReply struct {
	Retval    int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	StatIndex uint32 `binapi:"u32,name=stat_index" json:"stat_index,omitempty"`
}

func (m *IpfixExporterCreateDeleteReply) Reset() { *m = IpfixExporterCreateDeleteReply{} }
func (*IpfixExporterCreateDeleteReply) GetMessageName() string {
	return "ipfix_exporter_create_delete_reply"
}
func (*IpfixExporterCreateDeleteReply) GetCrcString() string { return "9ffac24b" }
func (*IpfixExporterCreateDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IpfixExporterCreateDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.StatIndex
	return size
}
func (m *IpfixExporterCreateDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.StatIndex)
	return buf.Bytes(), nil
}
func (m *IpfixExporterCreateDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.StatIndex = buf.DecodeUint32()
	return nil
}

// IpfixExporterDetails defines message 'ipfix_exporter_details'.
type IpfixExporterDetails struct {
	CollectorAddress ip_types.Address `binapi:"address,name=collector_address" json:"collector_address,omitempty"`
	CollectorPort    uint16           `binapi:"u16,name=collector_port" json:"collector_port,omitempty"`
	SrcAddress       ip_types.Address `binapi:"address,name=src_address" json:"src_address,omitempty"`
	VrfID            uint32           `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
	PathMtu          uint32           `binapi:"u32,name=path_mtu" json:"path_mtu,omitempty"`
	TemplateInterval uint32           `binapi:"u32,name=template_interval" json:"template_interval,omitempty"`
	UDPChecksum      bool             `binapi:"bool,name=udp_checksum" json:"udp_checksum,omitempty"`
}

func (m *IpfixExporterDetails) Reset()               { *m = IpfixExporterDetails{} }
func (*IpfixExporterDetails) GetMessageName() string { return "ipfix_exporter_details" }
func (*IpfixExporterDetails) GetCrcString() string   { return "0dedbfe4" }
func (*IpfixExporterDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IpfixExporterDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.CollectorAddress.Af
	size += 1 * 16 // m.CollectorAddress.Un
	size += 2      // m.CollectorPort
	size += 1      // m.SrcAddress.Af
	size += 1 * 16 // m.SrcAddress.Un
	size += 4      // m.VrfID
	size += 4      // m.PathMtu
	size += 4      // m.TemplateInterval
	size += 1      // m.UDPChecksum
	return size
}
func (m *IpfixExporterDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.CollectorAddress.Af))
	buf.EncodeBytes(m.CollectorAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint16(m.CollectorPort)
	buf.EncodeUint8(uint8(m.SrcAddress.Af))
	buf.EncodeBytes(m.SrcAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.VrfID)
	buf.EncodeUint32(m.PathMtu)
	buf.EncodeUint32(m.TemplateInterval)
	buf.EncodeBool(m.UDPChecksum)
	return buf.Bytes(), nil
}
func (m *IpfixExporterDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.CollectorAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.CollectorAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.CollectorPort = buf.DecodeUint16()
	m.SrcAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.SrcAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.VrfID = buf.DecodeUint32()
	m.PathMtu = buf.DecodeUint32()
	m.TemplateInterval = buf.DecodeUint32()
	m.UDPChecksum = buf.DecodeBool()
	return nil
}

// IpfixExporterDump defines message 'ipfix_exporter_dump'.
type IpfixExporterDump struct{}

func (m *IpfixExporterDump) Reset()               { *m = IpfixExporterDump{} }
func (*IpfixExporterDump) GetMessageName() string { return "ipfix_exporter_dump" }
func (*IpfixExporterDump) GetCrcString() string   { return "51077d14" }
func (*IpfixExporterDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IpfixExporterDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *IpfixExporterDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *IpfixExporterDump) Unmarshal(b []byte) error {
	return nil
}

// IpfixFlush defines message 'ipfix_flush'.
type IpfixFlush struct{}

func (m *IpfixFlush) Reset()               { *m = IpfixFlush{} }
func (*IpfixFlush) GetMessageName() string { return "ipfix_flush" }
func (*IpfixFlush) GetCrcString() string   { return "51077d14" }
func (*IpfixFlush) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IpfixFlush) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *IpfixFlush) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *IpfixFlush) Unmarshal(b []byte) error {
	return nil
}

// IpfixFlushReply defines message 'ipfix_flush_reply'.
type IpfixFlushReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IpfixFlushReply) Reset()               { *m = IpfixFlushReply{} }
func (*IpfixFlushReply) GetMessageName() string { return "ipfix_flush_reply" }
func (*IpfixFlushReply) GetCrcString() string   { return "e8d4e804" }
func (*IpfixFlushReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IpfixFlushReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IpfixFlushReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IpfixFlushReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SetIpfixClassifyStream defines message 'set_ipfix_classify_stream'.
type SetIpfixClassifyStream struct {
	DomainID uint32 `binapi:"u32,name=domain_id" json:"domain_id,omitempty"`
	SrcPort  uint16 `binapi:"u16,name=src_port" json:"src_port,omitempty"`
}

func (m *SetIpfixClassifyStream) Reset()               { *m = SetIpfixClassifyStream{} }
func (*SetIpfixClassifyStream) GetMessageName() string { return "set_ipfix_classify_stream" }
func (*SetIpfixClassifyStream) GetCrcString() string   { return "c9cbe053" }
func (*SetIpfixClassifyStream) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SetIpfixClassifyStream) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.DomainID
	size += 2 // m.SrcPort
	return size
}
func (m *SetIpfixClassifyStream) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.DomainID)
	buf.EncodeUint16(m.SrcPort)
	return buf.Bytes(), nil
}
func (m *SetIpfixClassifyStream) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.DomainID = buf.DecodeUint32()
	m.SrcPort = buf.DecodeUint16()
	return nil
}

// SetIpfixClassifyStreamReply defines message 'set_ipfix_classify_stream_reply'.
type SetIpfixClassifyStreamReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SetIpfixClassifyStreamReply) Reset()               { *m = SetIpfixClassifyStreamReply{} }
func (*SetIpfixClassifyStreamReply) GetMessageName() string { return "set_ipfix_classify_stream_reply" }
func (*SetIpfixClassifyStreamReply) GetCrcString() string   { return "e8d4e804" }
func (*SetIpfixClassifyStreamReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SetIpfixClassifyStreamReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SetIpfixClassifyStreamReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SetIpfixClassifyStreamReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SetIpfixExporter defines message 'set_ipfix_exporter'.
type SetIpfixExporter struct {
	CollectorAddress ip_types.Address `binapi:"address,name=collector_address" json:"collector_address,omitempty"`
	CollectorPort    uint16           `binapi:"u16,name=collector_port" json:"collector_port,omitempty"`
	SrcAddress       ip_types.Address `binapi:"address,name=src_address" json:"src_address,omitempty"`
	VrfID            uint32           `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
	PathMtu          uint32           `binapi:"u32,name=path_mtu" json:"path_mtu,omitempty"`
	TemplateInterval uint32           `binapi:"u32,name=template_interval" json:"template_interval,omitempty"`
	UDPChecksum      bool             `binapi:"bool,name=udp_checksum" json:"udp_checksum,omitempty"`
}

func (m *SetIpfixExporter) Reset()               { *m = SetIpfixExporter{} }
func (*SetIpfixExporter) GetMessageName() string { return "set_ipfix_exporter" }
func (*SetIpfixExporter) GetCrcString() string   { return "5530c8a0" }
func (*SetIpfixExporter) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SetIpfixExporter) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.CollectorAddress.Af
	size += 1 * 16 // m.CollectorAddress.Un
	size += 2      // m.CollectorPort
	size += 1      // m.SrcAddress.Af
	size += 1 * 16 // m.SrcAddress.Un
	size += 4      // m.VrfID
	size += 4      // m.PathMtu
	size += 4      // m.TemplateInterval
	size += 1      // m.UDPChecksum
	return size
}
func (m *SetIpfixExporter) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.CollectorAddress.Af))
	buf.EncodeBytes(m.CollectorAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint16(m.CollectorPort)
	buf.EncodeUint8(uint8(m.SrcAddress.Af))
	buf.EncodeBytes(m.SrcAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.VrfID)
	buf.EncodeUint32(m.PathMtu)
	buf.EncodeUint32(m.TemplateInterval)
	buf.EncodeBool(m.UDPChecksum)
	return buf.Bytes(), nil
}
func (m *SetIpfixExporter) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.CollectorAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.CollectorAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.CollectorPort = buf.DecodeUint16()
	m.SrcAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.SrcAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.VrfID = buf.DecodeUint32()
	m.PathMtu = buf.DecodeUint32()
	m.TemplateInterval = buf.DecodeUint32()
	m.UDPChecksum = buf.DecodeBool()
	return nil
}

// SetIpfixExporterReply defines message 'set_ipfix_exporter_reply'.
type SetIpfixExporterReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SetIpfixExporterReply) Reset()               { *m = SetIpfixExporterReply{} }
func (*SetIpfixExporterReply) GetMessageName() string { return "set_ipfix_exporter_reply" }
func (*SetIpfixExporterReply) GetCrcString() string   { return "e8d4e804" }
func (*SetIpfixExporterReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SetIpfixExporterReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SetIpfixExporterReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SetIpfixExporterReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_ipfix_export_binapi_init() }
func file_ipfix_export_binapi_init() {
	api.RegisterMessage((*IpfixAllExporterDetails)(nil), "ipfix_all_exporter_details_0dedbfe4")
	api.RegisterMessage((*IpfixAllExporterGet)(nil), "ipfix_all_exporter_get_f75ba505")
	api.RegisterMessage((*IpfixAllExporterGetReply)(nil), "ipfix_all_exporter_get_reply_53b48f5d")
	api.RegisterMessage((*IpfixClassifyStreamDetails)(nil), "ipfix_classify_stream_details_2903539d")
	api.RegisterMessage((*IpfixClassifyStreamDump)(nil), "ipfix_classify_stream_dump_51077d14")
	api.RegisterMessage((*IpfixClassifyTableAddDel)(nil), "ipfix_classify_table_add_del_3e449bb9")
	api.RegisterMessage((*IpfixClassifyTableAddDelReply)(nil), "ipfix_classify_table_add_del_reply_e8d4e804")
	api.RegisterMessage((*IpfixClassifyTableDetails)(nil), "ipfix_classify_table_details_1af8c28c")
	api.RegisterMessage((*IpfixClassifyTableDump)(nil), "ipfix_classify_table_dump_51077d14")
	api.RegisterMessage((*IpfixExporterCreateDelete)(nil), "ipfix_exporter_create_delete_0753a768")
	api.RegisterMessage((*IpfixExporterCreateDeleteReply)(nil), "ipfix_exporter_create_delete_reply_9ffac24b")
	api.RegisterMessage((*IpfixExporterDetails)(nil), "ipfix_exporter_details_0dedbfe4")
	api.RegisterMessage((*IpfixExporterDump)(nil), "ipfix_exporter_dump_51077d14")
	api.RegisterMessage((*IpfixFlush)(nil), "ipfix_flush_51077d14")
	api.RegisterMessage((*IpfixFlushReply)(nil), "ipfix_flush_reply_e8d4e804")
	api.RegisterMessage((*SetIpfixClassifyStream)(nil), "set_ipfix_classify_stream_c9cbe053")
	api.RegisterMessage((*SetIpfixClassifyStreamReply)(nil), "set_ipfix_classify_stream_reply_e8d4e804")
	api.RegisterMessage((*SetIpfixExporter)(nil), "set_ipfix_exporter_5530c8a0")
	api.RegisterMessage((*SetIpfixExporterReply)(nil), "set_ipfix_exporter_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IpfixAllExporterDetails)(nil),
		(*IpfixAllExporterGet)(nil),
		(*IpfixAllExporterGetReply)(nil),
		(*IpfixClassifyStreamDetails)(nil),
		(*IpfixClassifyStreamDump)(nil),
		(*IpfixClassifyTableAddDel)(nil),
		(*IpfixClassifyTableAddDelReply)(nil),
		(*IpfixClassifyTableDetails)(nil),
		(*IpfixClassifyTableDump)(nil),
		(*IpfixExporterCreateDelete)(nil),
		(*IpfixExporterCreateDeleteReply)(nil),
		(*IpfixExporterDetails)(nil),
		(*IpfixExporterDump)(nil),
		(*IpfixFlush)(nil),
		(*IpfixFlushReply)(nil),
		(*SetIpfixClassifyStream)(nil),
		(*SetIpfixClassifyStreamReply)(nil),
		(*SetIpfixExporter)(nil),
		(*SetIpfixExporterReply)(nil),
	}
}
