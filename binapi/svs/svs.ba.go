// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.05-release
// source: /usr/share/vpp/api/plugins/svs.api.json

// Package svs contains generated bindings for API file svs.api.
//
// Contents:
//  10 messages
//
package svs

import (
	api "git.fd.io/govpp.git/api"
	interface_types "git.fd.io/govpp.git/binapi/interface_types"
	ip_types "git.fd.io/govpp.git/binapi/ip_types"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "svs"
	APIVersion = "1.0.0"
	VersionCrc = 0x1644b2d6
)

// SvsDetails defines message 'svs_details'.
type SvsDetails struct {
	TableID   uint32                         `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Af        ip_types.AddressFamily         `binapi:"address_family,name=af" json:"af,omitempty"`
}

func (m *SvsDetails) Reset()               { *m = SvsDetails{} }
func (*SvsDetails) GetMessageName() string { return "svs_details" }
func (*SvsDetails) GetCrcString() string   { return "b8523d64" }
func (*SvsDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SvsDetails) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.TableID
	size += 4 // m.SwIfIndex
	size += 1 // m.Af
	return size
}
func (m *SvsDetails) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.TableID))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(uint8(m.Af))
	return buf.Bytes(), nil
}
func (m *SvsDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TableID = buf.DecodeUint32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Af = ip_types.AddressFamily(buf.DecodeUint8())
	return nil
}

// SvsDump defines message 'svs_dump'.
type SvsDump struct{}

func (m *SvsDump) Reset()               { *m = SvsDump{} }
func (*SvsDump) GetMessageName() string { return "svs_dump" }
func (*SvsDump) GetCrcString() string   { return "51077d14" }
func (*SvsDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SvsDump) Size() int {
	if m == nil {
		return 0
	}
	var size int
	return size
}
func (m *SvsDump) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	return buf.Bytes(), nil
}
func (m *SvsDump) Unmarshal(b []byte) error {
	return nil
}

// SvsEnableDisable defines message 'svs_enable_disable'.
type SvsEnableDisable struct {
	IsEnable  bool                           `binapi:"bool,name=is_enable" json:"is_enable,omitempty"`
	Af        ip_types.AddressFamily         `binapi:"address_family,name=af" json:"af,omitempty"`
	TableID   uint32                         `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *SvsEnableDisable) Reset()               { *m = SvsEnableDisable{} }
func (*SvsEnableDisable) GetMessageName() string { return "svs_enable_disable" }
func (*SvsEnableDisable) GetCrcString() string   { return "634b89d2" }
func (*SvsEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SvsEnableDisable) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 1 // m.IsEnable
	size += 1 // m.Af
	size += 4 // m.TableID
	size += 4 // m.SwIfIndex
	return size
}
func (m *SvsEnableDisable) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeBool(m.IsEnable)
	buf.EncodeUint8(uint8(m.Af))
	buf.EncodeUint32(uint32(m.TableID))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *SvsEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsEnable = buf.DecodeBool()
	m.Af = ip_types.AddressFamily(buf.DecodeUint8())
	m.TableID = buf.DecodeUint32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// SvsEnableDisableReply defines message 'svs_enable_disable_reply'.
type SvsEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SvsEnableDisableReply) Reset()               { *m = SvsEnableDisableReply{} }
func (*SvsEnableDisableReply) GetMessageName() string { return "svs_enable_disable_reply" }
func (*SvsEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*SvsEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SvsEnableDisableReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Retval
	return size
}
func (m *SvsEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Retval))
	return buf.Bytes(), nil
}
func (m *SvsEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = int32(buf.DecodeUint32())
	return nil
}

// SvsPluginGetVersion defines message 'svs_plugin_get_version'.
type SvsPluginGetVersion struct{}

func (m *SvsPluginGetVersion) Reset()               { *m = SvsPluginGetVersion{} }
func (*SvsPluginGetVersion) GetMessageName() string { return "svs_plugin_get_version" }
func (*SvsPluginGetVersion) GetCrcString() string   { return "51077d14" }
func (*SvsPluginGetVersion) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SvsPluginGetVersion) Size() int {
	if m == nil {
		return 0
	}
	var size int
	return size
}
func (m *SvsPluginGetVersion) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	return buf.Bytes(), nil
}
func (m *SvsPluginGetVersion) Unmarshal(b []byte) error {
	return nil
}

// SvsPluginGetVersionReply defines message 'svs_plugin_get_version_reply'.
type SvsPluginGetVersionReply struct {
	Major uint32 `binapi:"u32,name=major" json:"major,omitempty"`
	Minor uint32 `binapi:"u32,name=minor" json:"minor,omitempty"`
}

func (m *SvsPluginGetVersionReply) Reset()               { *m = SvsPluginGetVersionReply{} }
func (*SvsPluginGetVersionReply) GetMessageName() string { return "svs_plugin_get_version_reply" }
func (*SvsPluginGetVersionReply) GetCrcString() string   { return "9b32cf86" }
func (*SvsPluginGetVersionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SvsPluginGetVersionReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Major
	size += 4 // m.Minor
	return size
}
func (m *SvsPluginGetVersionReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Major))
	buf.EncodeUint32(uint32(m.Minor))
	return buf.Bytes(), nil
}
func (m *SvsPluginGetVersionReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Major = buf.DecodeUint32()
	m.Minor = buf.DecodeUint32()
	return nil
}

// SvsRouteAddDel defines message 'svs_route_add_del'.
type SvsRouteAddDel struct {
	IsAdd         bool            `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Prefix        ip_types.Prefix `binapi:"prefix,name=prefix" json:"prefix,omitempty"`
	TableID       uint32          `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	SourceTableID uint32          `binapi:"u32,name=source_table_id" json:"source_table_id,omitempty"`
}

func (m *SvsRouteAddDel) Reset()               { *m = SvsRouteAddDel{} }
func (*SvsRouteAddDel) GetMessageName() string { return "svs_route_add_del" }
func (*SvsRouteAddDel) GetCrcString() string   { return "d39e31fc" }
func (*SvsRouteAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SvsRouteAddDel) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 1      // m.IsAdd
	size += 1      // m.Prefix.Address.Af
	size += 1 * 16 // m.Prefix.Address.Un
	size += 1      // m.Prefix.Len
	size += 4      // m.TableID
	size += 4      // m.SourceTableID
	return size
}
func (m *SvsRouteAddDel) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint8(uint8(m.Prefix.Address.Af))
	buf.EncodeBytes(m.Prefix.Address.Un.XXX_UnionData[:], 0)
	buf.EncodeUint8(uint8(m.Prefix.Len))
	buf.EncodeUint32(uint32(m.TableID))
	buf.EncodeUint32(uint32(m.SourceTableID))
	return buf.Bytes(), nil
}
func (m *SvsRouteAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Prefix.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Prefix.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Prefix.Len = buf.DecodeUint8()
	m.TableID = buf.DecodeUint32()
	m.SourceTableID = buf.DecodeUint32()
	return nil
}

// SvsRouteAddDelReply defines message 'svs_route_add_del_reply'.
type SvsRouteAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SvsRouteAddDelReply) Reset()               { *m = SvsRouteAddDelReply{} }
func (*SvsRouteAddDelReply) GetMessageName() string { return "svs_route_add_del_reply" }
func (*SvsRouteAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*SvsRouteAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SvsRouteAddDelReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Retval
	return size
}
func (m *SvsRouteAddDelReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Retval))
	return buf.Bytes(), nil
}
func (m *SvsRouteAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = int32(buf.DecodeUint32())
	return nil
}

// SvsTableAddDel defines message 'svs_table_add_del'.
type SvsTableAddDel struct {
	IsAdd   bool                   `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Af      ip_types.AddressFamily `binapi:"address_family,name=af" json:"af,omitempty"`
	TableID uint32                 `binapi:"u32,name=table_id" json:"table_id,omitempty"`
}

func (m *SvsTableAddDel) Reset()               { *m = SvsTableAddDel{} }
func (*SvsTableAddDel) GetMessageName() string { return "svs_table_add_del" }
func (*SvsTableAddDel) GetCrcString() string   { return "7d21cb2a" }
func (*SvsTableAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SvsTableAddDel) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 1 // m.IsAdd
	size += 1 // m.Af
	size += 4 // m.TableID
	return size
}
func (m *SvsTableAddDel) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint8(uint8(m.Af))
	buf.EncodeUint32(uint32(m.TableID))
	return buf.Bytes(), nil
}
func (m *SvsTableAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Af = ip_types.AddressFamily(buf.DecodeUint8())
	m.TableID = buf.DecodeUint32()
	return nil
}

// SvsTableAddDelReply defines message 'svs_table_add_del_reply'.
type SvsTableAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SvsTableAddDelReply) Reset()               { *m = SvsTableAddDelReply{} }
func (*SvsTableAddDelReply) GetMessageName() string { return "svs_table_add_del_reply" }
func (*SvsTableAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*SvsTableAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SvsTableAddDelReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Retval
	return size
}
func (m *SvsTableAddDelReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Retval))
	return buf.Bytes(), nil
}
func (m *SvsTableAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = int32(buf.DecodeUint32())
	return nil
}

func init() { file_svs_binapi_init() }
func file_svs_binapi_init() {
	api.RegisterMessage((*SvsDetails)(nil), "svs_details_b8523d64")
	api.RegisterMessage((*SvsDump)(nil), "svs_dump_51077d14")
	api.RegisterMessage((*SvsEnableDisable)(nil), "svs_enable_disable_634b89d2")
	api.RegisterMessage((*SvsEnableDisableReply)(nil), "svs_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*SvsPluginGetVersion)(nil), "svs_plugin_get_version_51077d14")
	api.RegisterMessage((*SvsPluginGetVersionReply)(nil), "svs_plugin_get_version_reply_9b32cf86")
	api.RegisterMessage((*SvsRouteAddDel)(nil), "svs_route_add_del_d39e31fc")
	api.RegisterMessage((*SvsRouteAddDelReply)(nil), "svs_route_add_del_reply_e8d4e804")
	api.RegisterMessage((*SvsTableAddDel)(nil), "svs_table_add_del_7d21cb2a")
	api.RegisterMessage((*SvsTableAddDelReply)(nil), "svs_table_add_del_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SvsDetails)(nil),
		(*SvsDump)(nil),
		(*SvsEnableDisable)(nil),
		(*SvsEnableDisableReply)(nil),
		(*SvsPluginGetVersion)(nil),
		(*SvsPluginGetVersionReply)(nil),
		(*SvsRouteAddDel)(nil),
		(*SvsRouteAddDelReply)(nil),
		(*SvsTableAddDel)(nil),
		(*SvsTableAddDelReply)(nil),
	}
}
