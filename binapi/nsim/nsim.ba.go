// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.05-release
// source: /usr/share/vpp/api/plugins/nsim.api.json

// Package nsim contains generated bindings for API file nsim.api.
//
// Contents:
//   6 messages
//
package nsim

import (
	api "git.fd.io/govpp.git/api"
	interface_types "git.fd.io/govpp.git/binapi/interface_types"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "nsim"
	APIVersion = "2.1.1"
	VersionCrc = 0xcd5b42e8
)

// NsimConfigure defines message 'nsim_configure'.
type NsimConfigure struct {
	DelayInUsec              uint32 `binapi:"u32,name=delay_in_usec" json:"delay_in_usec,omitempty"`
	AveragePacketSize        uint32 `binapi:"u32,name=average_packet_size" json:"average_packet_size,omitempty"`
	BandwidthInBitsPerSecond uint64 `binapi:"u64,name=bandwidth_in_bits_per_second" json:"bandwidth_in_bits_per_second,omitempty"`
	PacketsPerDrop           uint32 `binapi:"u32,name=packets_per_drop" json:"packets_per_drop,omitempty"`
}

func (m *NsimConfigure) Reset()               { *m = NsimConfigure{} }
func (*NsimConfigure) GetMessageName() string { return "nsim_configure" }
func (*NsimConfigure) GetCrcString() string   { return "16ed400f" }
func (*NsimConfigure) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *NsimConfigure) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.DelayInUsec
	size += 4 // m.AveragePacketSize
	size += 8 // m.BandwidthInBitsPerSecond
	size += 4 // m.PacketsPerDrop
	return size
}
func (m *NsimConfigure) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.DelayInUsec))
	buf.EncodeUint32(uint32(m.AveragePacketSize))
	buf.EncodeUint64(uint64(m.BandwidthInBitsPerSecond))
	buf.EncodeUint32(uint32(m.PacketsPerDrop))
	return buf.Bytes(), nil
}
func (m *NsimConfigure) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.DelayInUsec = buf.DecodeUint32()
	m.AveragePacketSize = buf.DecodeUint32()
	m.BandwidthInBitsPerSecond = buf.DecodeUint64()
	m.PacketsPerDrop = buf.DecodeUint32()
	return nil
}

// NsimConfigureReply defines message 'nsim_configure_reply'.
type NsimConfigureReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *NsimConfigureReply) Reset()               { *m = NsimConfigureReply{} }
func (*NsimConfigureReply) GetMessageName() string { return "nsim_configure_reply" }
func (*NsimConfigureReply) GetCrcString() string   { return "e8d4e804" }
func (*NsimConfigureReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *NsimConfigureReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Retval
	return size
}
func (m *NsimConfigureReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Retval))
	return buf.Bytes(), nil
}
func (m *NsimConfigureReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = int32(buf.DecodeUint32())
	return nil
}

// NsimCrossConnectEnableDisable defines message 'nsim_cross_connect_enable_disable'.
type NsimCrossConnectEnableDisable struct {
	EnableDisable bool                           `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
	SwIfIndex0    interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index0" json:"sw_if_index0,omitempty"`
	SwIfIndex1    interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index1" json:"sw_if_index1,omitempty"`
}

func (m *NsimCrossConnectEnableDisable) Reset() { *m = NsimCrossConnectEnableDisable{} }
func (*NsimCrossConnectEnableDisable) GetMessageName() string {
	return "nsim_cross_connect_enable_disable"
}
func (*NsimCrossConnectEnableDisable) GetCrcString() string { return "16f70bdf" }
func (*NsimCrossConnectEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *NsimCrossConnectEnableDisable) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 1 // m.EnableDisable
	size += 4 // m.SwIfIndex0
	size += 4 // m.SwIfIndex1
	return size
}
func (m *NsimCrossConnectEnableDisable) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeBool(m.EnableDisable)
	buf.EncodeUint32(uint32(m.SwIfIndex0))
	buf.EncodeUint32(uint32(m.SwIfIndex1))
	return buf.Bytes(), nil
}
func (m *NsimCrossConnectEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EnableDisable = buf.DecodeBool()
	m.SwIfIndex0 = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.SwIfIndex1 = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// NsimCrossConnectEnableDisableReply defines message 'nsim_cross_connect_enable_disable_reply'.
type NsimCrossConnectEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *NsimCrossConnectEnableDisableReply) Reset() { *m = NsimCrossConnectEnableDisableReply{} }
func (*NsimCrossConnectEnableDisableReply) GetMessageName() string {
	return "nsim_cross_connect_enable_disable_reply"
}
func (*NsimCrossConnectEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*NsimCrossConnectEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *NsimCrossConnectEnableDisableReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Retval
	return size
}
func (m *NsimCrossConnectEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Retval))
	return buf.Bytes(), nil
}
func (m *NsimCrossConnectEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = int32(buf.DecodeUint32())
	return nil
}

// NsimOutputFeatureEnableDisable defines message 'nsim_output_feature_enable_disable'.
type NsimOutputFeatureEnableDisable struct {
	EnableDisable bool                           `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *NsimOutputFeatureEnableDisable) Reset() { *m = NsimOutputFeatureEnableDisable{} }
func (*NsimOutputFeatureEnableDisable) GetMessageName() string {
	return "nsim_output_feature_enable_disable"
}
func (*NsimOutputFeatureEnableDisable) GetCrcString() string { return "3865946c" }
func (*NsimOutputFeatureEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *NsimOutputFeatureEnableDisable) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 1 // m.EnableDisable
	size += 4 // m.SwIfIndex
	return size
}
func (m *NsimOutputFeatureEnableDisable) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeBool(m.EnableDisable)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *NsimOutputFeatureEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EnableDisable = buf.DecodeBool()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// NsimOutputFeatureEnableDisableReply defines message 'nsim_output_feature_enable_disable_reply'.
type NsimOutputFeatureEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *NsimOutputFeatureEnableDisableReply) Reset() { *m = NsimOutputFeatureEnableDisableReply{} }
func (*NsimOutputFeatureEnableDisableReply) GetMessageName() string {
	return "nsim_output_feature_enable_disable_reply"
}
func (*NsimOutputFeatureEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*NsimOutputFeatureEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *NsimOutputFeatureEnableDisableReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	size += 4 // m.Retval
	return size
}
func (m *NsimOutputFeatureEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	var buf *codec.Buffer
	if b == nil {
		buf = codec.NewBuffer(make([]byte, m.Size()))
	} else {
		buf = codec.NewBuffer(b)
	}
	buf.EncodeUint32(uint32(m.Retval))
	return buf.Bytes(), nil
}
func (m *NsimOutputFeatureEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = int32(buf.DecodeUint32())
	return nil
}

func init() { file_nsim_binapi_init() }
func file_nsim_binapi_init() {
	api.RegisterMessage((*NsimConfigure)(nil), "nsim_configure_16ed400f")
	api.RegisterMessage((*NsimConfigureReply)(nil), "nsim_configure_reply_e8d4e804")
	api.RegisterMessage((*NsimCrossConnectEnableDisable)(nil), "nsim_cross_connect_enable_disable_16f70bdf")
	api.RegisterMessage((*NsimCrossConnectEnableDisableReply)(nil), "nsim_cross_connect_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*NsimOutputFeatureEnableDisable)(nil), "nsim_output_feature_enable_disable_3865946c")
	api.RegisterMessage((*NsimOutputFeatureEnableDisableReply)(nil), "nsim_output_feature_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*NsimConfigure)(nil),
		(*NsimConfigureReply)(nil),
		(*NsimCrossConnectEnableDisable)(nil),
		(*NsimCrossConnectEnableDisableReply)(nil),
		(*NsimOutputFeatureEnableDisable)(nil),
		(*NsimOutputFeatureEnableDisableReply)(nil),
	}
}
