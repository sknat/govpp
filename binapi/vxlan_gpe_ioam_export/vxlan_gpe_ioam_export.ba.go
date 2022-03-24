// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              22.02-release
// source: /usr/share/vpp/api/plugins/vxlan_gpe_ioam_export.api.json

// Package vxlan_gpe_ioam_export contains generated bindings for API file vxlan_gpe_ioam_export.api.
//
// Contents:
//   2 messages
//
package vxlan_gpe_ioam_export

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
	APIFile    = "vxlan_gpe_ioam_export"
	APIVersion = "1.0.0"
	VersionCrc = 0x26bebf64
)

// VxlanGpeIoamExportEnableDisable defines message 'vxlan_gpe_ioam_export_enable_disable'.
type VxlanGpeIoamExportEnableDisable struct {
	IsDisable        bool                `binapi:"bool,name=is_disable" json:"is_disable,omitempty"`
	CollectorAddress ip_types.IP4Address `binapi:"ip4_address,name=collector_address" json:"collector_address,omitempty"`
	SrcAddress       ip_types.IP4Address `binapi:"ip4_address,name=src_address" json:"src_address,omitempty"`
}

func (m *VxlanGpeIoamExportEnableDisable) Reset() { *m = VxlanGpeIoamExportEnableDisable{} }
func (*VxlanGpeIoamExportEnableDisable) GetMessageName() string {
	return "vxlan_gpe_ioam_export_enable_disable"
}
func (*VxlanGpeIoamExportEnableDisable) GetCrcString() string { return "d4c76d3a" }
func (*VxlanGpeIoamExportEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (m *VxlanGpeIoamExportEnableDisable) GetRetVal() error {
	return nil
}

func (m *VxlanGpeIoamExportEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1     // m.IsDisable
	size += 1 * 4 // m.CollectorAddress
	size += 1 * 4 // m.SrcAddress
	return size
}
func (m *VxlanGpeIoamExportEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsDisable)
	buf.EncodeBytes(m.CollectorAddress[:], 4)
	buf.EncodeBytes(m.SrcAddress[:], 4)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamExportEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsDisable = buf.DecodeBool()
	copy(m.CollectorAddress[:], buf.DecodeBytes(4))
	copy(m.SrcAddress[:], buf.DecodeBytes(4))
	return nil
}

// VxlanGpeIoamExportEnableDisableReply defines message 'vxlan_gpe_ioam_export_enable_disable_reply'.
type VxlanGpeIoamExportEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *VxlanGpeIoamExportEnableDisableReply) Reset() { *m = VxlanGpeIoamExportEnableDisableReply{} }
func (*VxlanGpeIoamExportEnableDisableReply) GetMessageName() string {
	return "vxlan_gpe_ioam_export_enable_disable_reply"
}
func (*VxlanGpeIoamExportEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*VxlanGpeIoamExportEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *VxlanGpeIoamExportEnableDisableReply) GetRetVal() error {
	return api.RetvalToVPPApiError(int32(m.Retval))
}

func (m *VxlanGpeIoamExportEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *VxlanGpeIoamExportEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *VxlanGpeIoamExportEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_vxlan_gpe_ioam_export_binapi_init() }
func file_vxlan_gpe_ioam_export_binapi_init() {
	api.RegisterMessage((*VxlanGpeIoamExportEnableDisable)(nil), "vxlan_gpe_ioam_export_enable_disable_d4c76d3a")
	api.RegisterMessage((*VxlanGpeIoamExportEnableDisableReply)(nil), "vxlan_gpe_ioam_export_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*VxlanGpeIoamExportEnableDisable)(nil),
		(*VxlanGpeIoamExportEnableDisableReply)(nil),
	}
}
