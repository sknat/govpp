// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              22.02-release
// source: /usr/share/vpp/api/core/memclnt.api.json

// Package memclnt contains generated bindings for API file memclnt.api.
//
// Contents:
//   2 structs
//  24 messages
//
package memclnt

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "memclnt"
	APIVersion = "2.1.0"
	VersionCrc = 0x230bb938
)

// MessageTableEntry defines type 'message_table_entry'.
type MessageTableEntry struct {
	Index uint16 `binapi:"u16,name=index" json:"index,omitempty"`
	Name  string `binapi:"string[64],name=name" json:"name,omitempty"`
}

// ModuleVersion defines type 'module_version'.
type ModuleVersion struct {
	Major uint32 `binapi:"u32,name=major" json:"major,omitempty"`
	Minor uint32 `binapi:"u32,name=minor" json:"minor,omitempty"`
	Patch uint32 `binapi:"u32,name=patch" json:"patch,omitempty"`
	Name  string `binapi:"string[64],name=name" json:"name,omitempty"`
}

// APIVersions defines message 'api_versions'.
type APIVersions struct{}

func (m *APIVersions) Reset()               { *m = APIVersions{} }
func (*APIVersions) GetMessageName() string { return "api_versions" }
func (*APIVersions) GetCrcString() string   { return "51077d14" }
func (*APIVersions) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *APIVersions) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *APIVersions) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *APIVersions) Unmarshal(b []byte) error {
	return nil
}

// APIVersionsReply defines message 'api_versions_reply'.
type APIVersionsReply struct {
	Retval      int32           `binapi:"i32,name=retval" json:"retval,omitempty"`
	Count       uint32          `binapi:"u32,name=count" json:"-"`
	APIVersions []ModuleVersion `binapi:"module_version[count],name=api_versions" json:"api_versions,omitempty"`
}

func (m *APIVersionsReply) Reset()               { *m = APIVersionsReply{} }
func (*APIVersionsReply) GetMessageName() string { return "api_versions_reply" }
func (*APIVersionsReply) GetCrcString() string   { return "5f0d99d6" }
func (*APIVersionsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *APIVersionsReply) GetRetVal() error {
	return api.RetvalToVPPApiError(int32(m.Retval))
}

func (m *APIVersionsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.Count
	for j1 := 0; j1 < len(m.APIVersions); j1++ {
		var s1 ModuleVersion
		_ = s1
		if j1 < len(m.APIVersions) {
			//lint:ignore SA4006 we might not use this variable
			s1 = m.APIVersions[j1]
		}
		size += 4  // s1.Major
		size += 4  // s1.Minor
		size += 4  // s1.Patch
		size += 64 // s1.Name
	}
	return size
}
func (m *APIVersionsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(len(m.APIVersions)))
	for j0 := 0; j0 < len(m.APIVersions); j0++ {
		var v0 ModuleVersion // APIVersions
		if j0 < len(m.APIVersions) {
			v0 = m.APIVersions[j0]
		}
		buf.EncodeUint32(v0.Major)
		buf.EncodeUint32(v0.Minor)
		buf.EncodeUint32(v0.Patch)
		buf.EncodeString(v0.Name, 64)
	}
	return buf.Bytes(), nil
}
func (m *APIVersionsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Count = buf.DecodeUint32()
	m.APIVersions = make([]ModuleVersion, m.Count)
	for j0 := 0; j0 < len(m.APIVersions); j0++ {
		m.APIVersions[j0].Major = buf.DecodeUint32()
		m.APIVersions[j0].Minor = buf.DecodeUint32()
		m.APIVersions[j0].Patch = buf.DecodeUint32()
		m.APIVersions[j0].Name = buf.DecodeString(64)
	}
	return nil
}

// ControlPing defines message 'control_ping'.
type ControlPing struct{}

func (m *ControlPing) Reset()               { *m = ControlPing{} }
func (*ControlPing) GetMessageName() string { return "control_ping" }
func (*ControlPing) GetCrcString() string   { return "51077d14" }
func (*ControlPing) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ControlPing) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *ControlPing) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *ControlPing) Unmarshal(b []byte) error {
	return nil
}

// ControlPingReply defines message 'control_ping_reply'.
type ControlPingReply struct {
	Retval      int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ClientIndex uint32 `binapi:"u32,name=client_index" json:"client_index,omitempty"`
	VpePID      uint32 `binapi:"u32,name=vpe_pid" json:"vpe_pid,omitempty"`
}

func (m *ControlPingReply) Reset()               { *m = ControlPingReply{} }
func (*ControlPingReply) GetMessageName() string { return "control_ping_reply" }
func (*ControlPingReply) GetCrcString() string   { return "f6b0b8ca" }
func (*ControlPingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *ControlPingReply) GetRetVal() error {
	return api.RetvalToVPPApiError(int32(m.Retval))
}

func (m *ControlPingReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.ClientIndex
	size += 4 // m.VpePID
	return size
}
func (m *ControlPingReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.ClientIndex)
	buf.EncodeUint32(m.VpePID)
	return buf.Bytes(), nil
}
func (m *ControlPingReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ClientIndex = buf.DecodeUint32()
	m.VpePID = buf.DecodeUint32()
	return nil
}

// GetFirstMsgID defines message 'get_first_msg_id'.
type GetFirstMsgID struct {
	Name string `binapi:"string[64],name=name" json:"name,omitempty"`
}

func (m *GetFirstMsgID) Reset()               { *m = GetFirstMsgID{} }
func (*GetFirstMsgID) GetMessageName() string { return "get_first_msg_id" }
func (*GetFirstMsgID) GetCrcString() string   { return "ebf79a66" }
func (*GetFirstMsgID) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GetFirstMsgID) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.Name
	return size
}
func (m *GetFirstMsgID) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.Name, 64)
	return buf.Bytes(), nil
}
func (m *GetFirstMsgID) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Name = buf.DecodeString(64)
	return nil
}

// GetFirstMsgIDReply defines message 'get_first_msg_id_reply'.
type GetFirstMsgIDReply struct {
	Retval     int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	FirstMsgID uint16 `binapi:"u16,name=first_msg_id" json:"first_msg_id,omitempty"`
}

func (m *GetFirstMsgIDReply) Reset()               { *m = GetFirstMsgIDReply{} }
func (*GetFirstMsgIDReply) GetMessageName() string { return "get_first_msg_id_reply" }
func (*GetFirstMsgIDReply) GetCrcString() string   { return "7d337472" }
func (*GetFirstMsgIDReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *GetFirstMsgIDReply) GetRetVal() error {
	return api.RetvalToVPPApiError(int32(m.Retval))
}

func (m *GetFirstMsgIDReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 2 // m.FirstMsgID
	return size
}
func (m *GetFirstMsgIDReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint16(m.FirstMsgID)
	return buf.Bytes(), nil
}
func (m *GetFirstMsgIDReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.FirstMsgID = buf.DecodeUint16()
	return nil
}

// MemclntCreate defines message 'memclnt_create'.
type MemclntCreate struct {
	CtxQuota    int32    `binapi:"i32,name=ctx_quota" json:"ctx_quota,omitempty"`
	InputQueue  uint64   `binapi:"u64,name=input_queue" json:"input_queue,omitempty"`
	Name        string   `binapi:"string[64],name=name" json:"name,omitempty"`
	APIVersions []uint32 `binapi:"u32[8],name=api_versions" json:"api_versions,omitempty"`
}

func (m *MemclntCreate) Reset()               { *m = MemclntCreate{} }
func (*MemclntCreate) GetMessageName() string { return "memclnt_create" }
func (*MemclntCreate) GetCrcString() string   { return "9c5e1c2f" }
func (*MemclntCreate) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *MemclntCreate) GetRetVal() error {
	return nil
}

func (m *MemclntCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.CtxQuota
	size += 8     // m.InputQueue
	size += 64    // m.Name
	size += 4 * 8 // m.APIVersions
	return size
}
func (m *MemclntCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.CtxQuota)
	buf.EncodeUint64(m.InputQueue)
	buf.EncodeString(m.Name, 64)
	for i := 0; i < 8; i++ {
		var x uint32
		if i < len(m.APIVersions) {
			x = uint32(m.APIVersions[i])
		}
		buf.EncodeUint32(x)
	}
	return buf.Bytes(), nil
}
func (m *MemclntCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.CtxQuota = buf.DecodeInt32()
	m.InputQueue = buf.DecodeUint64()
	m.Name = buf.DecodeString(64)
	m.APIVersions = make([]uint32, 8)
	for i := 0; i < len(m.APIVersions); i++ {
		m.APIVersions[i] = buf.DecodeUint32()
	}
	return nil
}

// MemclntCreateReply defines message 'memclnt_create_reply'.
type MemclntCreateReply struct {
	Response     int32  `binapi:"i32,name=response" json:"response,omitempty"`
	Handle       uint64 `binapi:"u64,name=handle" json:"handle,omitempty"`
	Index        uint32 `binapi:"u32,name=index" json:"index,omitempty"`
	MessageTable uint64 `binapi:"u64,name=message_table" json:"message_table,omitempty"`
}

func (m *MemclntCreateReply) Reset()               { *m = MemclntCreateReply{} }
func (*MemclntCreateReply) GetMessageName() string { return "memclnt_create_reply" }
func (*MemclntCreateReply) GetCrcString() string   { return "42ec4560" }
func (*MemclntCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *MemclntCreateReply) GetRetVal() error {
	return nil
}

func (m *MemclntCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Response
	size += 8 // m.Handle
	size += 4 // m.Index
	size += 8 // m.MessageTable
	return size
}
func (m *MemclntCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Response)
	buf.EncodeUint64(m.Handle)
	buf.EncodeUint32(m.Index)
	buf.EncodeUint64(m.MessageTable)
	return buf.Bytes(), nil
}
func (m *MemclntCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Response = buf.DecodeInt32()
	m.Handle = buf.DecodeUint64()
	m.Index = buf.DecodeUint32()
	m.MessageTable = buf.DecodeUint64()
	return nil
}

// MemclntDelete defines message 'memclnt_delete'.
type MemclntDelete struct {
	Index     uint32 `binapi:"u32,name=index" json:"index,omitempty"`
	Handle    uint64 `binapi:"u64,name=handle" json:"handle,omitempty"`
	DoCleanup bool   `binapi:"bool,name=do_cleanup" json:"do_cleanup,omitempty"`
}

func (m *MemclntDelete) Reset()               { *m = MemclntDelete{} }
func (*MemclntDelete) GetMessageName() string { return "memclnt_delete" }
func (*MemclntDelete) GetCrcString() string   { return "7e1c04e3" }
func (*MemclntDelete) GetMessageType() api.MessageType {
	return api.OtherMessage
}

func (m *MemclntDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Index
	size += 8 // m.Handle
	size += 1 // m.DoCleanup
	return size
}
func (m *MemclntDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Index)
	buf.EncodeUint64(m.Handle)
	buf.EncodeBool(m.DoCleanup)
	return buf.Bytes(), nil
}
func (m *MemclntDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Index = buf.DecodeUint32()
	m.Handle = buf.DecodeUint64()
	m.DoCleanup = buf.DecodeBool()
	return nil
}

// MemclntDeleteReply defines message 'memclnt_delete_reply'.
type MemclntDeleteReply struct {
	Response int32  `binapi:"i32,name=response" json:"response,omitempty"`
	Handle   uint64 `binapi:"u64,name=handle" json:"handle,omitempty"`
}

func (m *MemclntDeleteReply) Reset()               { *m = MemclntDeleteReply{} }
func (*MemclntDeleteReply) GetMessageName() string { return "memclnt_delete_reply" }
func (*MemclntDeleteReply) GetCrcString() string   { return "3d3b6312" }
func (*MemclntDeleteReply) GetMessageType() api.MessageType {
	return api.OtherMessage
}

func (m *MemclntDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Response
	size += 8 // m.Handle
	return size
}
func (m *MemclntDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Response)
	buf.EncodeUint64(m.Handle)
	return buf.Bytes(), nil
}
func (m *MemclntDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Response = buf.DecodeInt32()
	m.Handle = buf.DecodeUint64()
	return nil
}

// MemclntKeepalive defines message 'memclnt_keepalive'.
type MemclntKeepalive struct{}

func (m *MemclntKeepalive) Reset()               { *m = MemclntKeepalive{} }
func (*MemclntKeepalive) GetMessageName() string { return "memclnt_keepalive" }
func (*MemclntKeepalive) GetCrcString() string   { return "51077d14" }
func (*MemclntKeepalive) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemclntKeepalive) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *MemclntKeepalive) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *MemclntKeepalive) Unmarshal(b []byte) error {
	return nil
}

// MemclntKeepaliveReply defines message 'memclnt_keepalive_reply'.
type MemclntKeepaliveReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *MemclntKeepaliveReply) Reset()               { *m = MemclntKeepaliveReply{} }
func (*MemclntKeepaliveReply) GetMessageName() string { return "memclnt_keepalive_reply" }
func (*MemclntKeepaliveReply) GetCrcString() string   { return "e8d4e804" }
func (*MemclntKeepaliveReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *MemclntKeepaliveReply) GetRetVal() error {
	return api.RetvalToVPPApiError(int32(m.Retval))
}

func (m *MemclntKeepaliveReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *MemclntKeepaliveReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *MemclntKeepaliveReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// MemclntReadTimeout defines message 'memclnt_read_timeout'.
type MemclntReadTimeout struct {
	Dummy uint8 `binapi:"u8,name=dummy" json:"dummy,omitempty"`
}

func (m *MemclntReadTimeout) Reset()               { *m = MemclntReadTimeout{} }
func (*MemclntReadTimeout) GetMessageName() string { return "memclnt_read_timeout" }
func (*MemclntReadTimeout) GetCrcString() string   { return "c3a3a452" }
func (*MemclntReadTimeout) GetMessageType() api.MessageType {
	return api.OtherMessage
}

func (m *MemclntReadTimeout) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Dummy
	return size
}
func (m *MemclntReadTimeout) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.Dummy)
	return buf.Bytes(), nil
}
func (m *MemclntReadTimeout) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Dummy = buf.DecodeUint8()
	return nil
}

// MemclntRxThreadSuspend defines message 'memclnt_rx_thread_suspend'.
type MemclntRxThreadSuspend struct {
	Dummy uint8 `binapi:"u8,name=dummy" json:"dummy,omitempty"`
}

func (m *MemclntRxThreadSuspend) Reset()               { *m = MemclntRxThreadSuspend{} }
func (*MemclntRxThreadSuspend) GetMessageName() string { return "memclnt_rx_thread_suspend" }
func (*MemclntRxThreadSuspend) GetCrcString() string   { return "c3a3a452" }
func (*MemclntRxThreadSuspend) GetMessageType() api.MessageType {
	return api.OtherMessage
}

func (m *MemclntRxThreadSuspend) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Dummy
	return size
}
func (m *MemclntRxThreadSuspend) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.Dummy)
	return buf.Bytes(), nil
}
func (m *MemclntRxThreadSuspend) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Dummy = buf.DecodeUint8()
	return nil
}

// RPCCall defines message 'rpc_call'.
type RPCCall struct {
	Function        uint64 `binapi:"u64,name=function" json:"function,omitempty"`
	Multicast       uint8  `binapi:"u8,name=multicast" json:"multicast,omitempty"`
	NeedBarrierSync uint8  `binapi:"u8,name=need_barrier_sync" json:"need_barrier_sync,omitempty"`
	SendReply       uint8  `binapi:"u8,name=send_reply" json:"send_reply,omitempty"`
	DataLen         uint32 `binapi:"u32,name=data_len" json:"-"`
	Data            []byte `binapi:"u8[data_len],name=data" json:"data,omitempty"`
}

func (m *RPCCall) Reset()               { *m = RPCCall{} }
func (*RPCCall) GetMessageName() string { return "rpc_call" }
func (*RPCCall) GetCrcString() string   { return "7e8a2c95" }
func (*RPCCall) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *RPCCall) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8               // m.Function
	size += 1               // m.Multicast
	size += 1               // m.NeedBarrierSync
	size += 1               // m.SendReply
	size += 4               // m.DataLen
	size += 1 * len(m.Data) // m.Data
	return size
}
func (m *RPCCall) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint64(m.Function)
	buf.EncodeUint8(m.Multicast)
	buf.EncodeUint8(m.NeedBarrierSync)
	buf.EncodeUint8(m.SendReply)
	buf.EncodeUint32(uint32(len(m.Data)))
	buf.EncodeBytes(m.Data, 0)
	return buf.Bytes(), nil
}
func (m *RPCCall) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Function = buf.DecodeUint64()
	m.Multicast = buf.DecodeUint8()
	m.NeedBarrierSync = buf.DecodeUint8()
	m.SendReply = buf.DecodeUint8()
	m.DataLen = buf.DecodeUint32()
	m.Data = make([]byte, m.DataLen)
	copy(m.Data, buf.DecodeBytes(len(m.Data)))
	return nil
}

// RPCCallReply defines message 'rpc_call_reply'.
type RPCCallReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *RPCCallReply) Reset()               { *m = RPCCallReply{} }
func (*RPCCallReply) GetMessageName() string { return "rpc_call_reply" }
func (*RPCCallReply) GetCrcString() string   { return "e8d4e804" }
func (*RPCCallReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *RPCCallReply) GetRetVal() error {
	return api.RetvalToVPPApiError(int32(m.Retval))
}

func (m *RPCCallReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *RPCCallReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *RPCCallReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// RxThreadExit defines message 'rx_thread_exit'.
type RxThreadExit struct {
	Dummy uint8 `binapi:"u8,name=dummy" json:"dummy,omitempty"`
}

func (m *RxThreadExit) Reset()               { *m = RxThreadExit{} }
func (*RxThreadExit) GetMessageName() string { return "rx_thread_exit" }
func (*RxThreadExit) GetCrcString() string   { return "c3a3a452" }
func (*RxThreadExit) GetMessageType() api.MessageType {
	return api.OtherMessage
}

func (m *RxThreadExit) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Dummy
	return size
}
func (m *RxThreadExit) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.Dummy)
	return buf.Bytes(), nil
}
func (m *RxThreadExit) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Dummy = buf.DecodeUint8()
	return nil
}

// SockInitShm defines message 'sock_init_shm'.
type SockInitShm struct {
	RequestedSize uint32   `binapi:"u32,name=requested_size" json:"requested_size,omitempty"`
	Nitems        uint8    `binapi:"u8,name=nitems" json:"-"`
	Configs       []uint64 `binapi:"u64[nitems],name=configs" json:"configs,omitempty"`
}

func (m *SockInitShm) Reset()               { *m = SockInitShm{} }
func (*SockInitShm) GetMessageName() string { return "sock_init_shm" }
func (*SockInitShm) GetCrcString() string   { return "51646d92" }
func (*SockInitShm) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SockInitShm) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                  // m.RequestedSize
	size += 1                  // m.Nitems
	size += 8 * len(m.Configs) // m.Configs
	return size
}
func (m *SockInitShm) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.RequestedSize)
	buf.EncodeUint8(uint8(len(m.Configs)))
	for i := 0; i < len(m.Configs); i++ {
		var x uint64
		if i < len(m.Configs) {
			x = uint64(m.Configs[i])
		}
		buf.EncodeUint64(x)
	}
	return buf.Bytes(), nil
}
func (m *SockInitShm) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.RequestedSize = buf.DecodeUint32()
	m.Nitems = buf.DecodeUint8()
	m.Configs = make([]uint64, m.Nitems)
	for i := 0; i < len(m.Configs); i++ {
		m.Configs[i] = buf.DecodeUint64()
	}
	return nil
}

// SockInitShmReply defines message 'sock_init_shm_reply'.
type SockInitShmReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SockInitShmReply) Reset()               { *m = SockInitShmReply{} }
func (*SockInitShmReply) GetMessageName() string { return "sock_init_shm_reply" }
func (*SockInitShmReply) GetCrcString() string   { return "e8d4e804" }
func (*SockInitShmReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *SockInitShmReply) GetRetVal() error {
	return api.RetvalToVPPApiError(int32(m.Retval))
}

func (m *SockInitShmReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SockInitShmReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SockInitShmReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SockclntCreate defines message 'sockclnt_create'.
type SockclntCreate struct {
	Name string `binapi:"string[64],name=name" json:"name,omitempty"`
}

func (m *SockclntCreate) Reset()               { *m = SockclntCreate{} }
func (*SockclntCreate) GetMessageName() string { return "sockclnt_create" }
func (*SockclntCreate) GetCrcString() string   { return "455fb9c4" }
func (*SockclntCreate) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *SockclntCreate) GetRetVal() error {
	return nil
}

func (m *SockclntCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.Name
	return size
}
func (m *SockclntCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.Name, 64)
	return buf.Bytes(), nil
}
func (m *SockclntCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Name = buf.DecodeString(64)
	return nil
}

// SockclntCreateReply defines message 'sockclnt_create_reply'.
type SockclntCreateReply struct {
	Response     int32               `binapi:"i32,name=response" json:"response,omitempty"`
	Index        uint32              `binapi:"u32,name=index" json:"index,omitempty"`
	Count        uint16              `binapi:"u16,name=count" json:"-"`
	MessageTable []MessageTableEntry `binapi:"message_table_entry[count],name=message_table" json:"message_table,omitempty"`
}

func (m *SockclntCreateReply) Reset()               { *m = SockclntCreateReply{} }
func (*SockclntCreateReply) GetMessageName() string { return "sockclnt_create_reply" }
func (*SockclntCreateReply) GetCrcString() string   { return "35166268" }
func (*SockclntCreateReply) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SockclntCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Response
	size += 4 // m.Index
	size += 2 // m.Count
	for j1 := 0; j1 < len(m.MessageTable); j1++ {
		var s1 MessageTableEntry
		_ = s1
		if j1 < len(m.MessageTable) {
			//lint:ignore SA4006 we might not use this variable
			s1 = m.MessageTable[j1]
		}
		size += 2  // s1.Index
		size += 64 // s1.Name
	}
	return size
}
func (m *SockclntCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Response)
	buf.EncodeUint32(m.Index)
	buf.EncodeUint16(uint16(len(m.MessageTable)))
	for j0 := 0; j0 < len(m.MessageTable); j0++ {
		var v0 MessageTableEntry // MessageTable
		if j0 < len(m.MessageTable) {
			v0 = m.MessageTable[j0]
		}
		buf.EncodeUint16(v0.Index)
		buf.EncodeString(v0.Name, 64)
	}
	return buf.Bytes(), nil
}
func (m *SockclntCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Response = buf.DecodeInt32()
	m.Index = buf.DecodeUint32()
	m.Count = buf.DecodeUint16()
	m.MessageTable = make([]MessageTableEntry, m.Count)
	for j0 := 0; j0 < len(m.MessageTable); j0++ {
		m.MessageTable[j0].Index = buf.DecodeUint16()
		m.MessageTable[j0].Name = buf.DecodeString(64)
	}
	return nil
}

// SockclntDelete defines message 'sockclnt_delete'.
type SockclntDelete struct {
	Index uint32 `binapi:"u32,name=index" json:"index,omitempty"`
}

func (m *SockclntDelete) Reset()               { *m = SockclntDelete{} }
func (*SockclntDelete) GetMessageName() string { return "sockclnt_delete" }
func (*SockclntDelete) GetCrcString() string   { return "8ac76db6" }
func (*SockclntDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SockclntDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Index
	return size
}
func (m *SockclntDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Index)
	return buf.Bytes(), nil
}
func (m *SockclntDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Index = buf.DecodeUint32()
	return nil
}

// SockclntDeleteReply defines message 'sockclnt_delete_reply'.
type SockclntDeleteReply struct {
	Response int32 `binapi:"i32,name=response" json:"response,omitempty"`
}

func (m *SockclntDeleteReply) Reset()               { *m = SockclntDeleteReply{} }
func (*SockclntDeleteReply) GetMessageName() string { return "sockclnt_delete_reply" }
func (*SockclntDeleteReply) GetCrcString() string   { return "8f38b1ee" }
func (*SockclntDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *SockclntDeleteReply) GetRetVal() error {
	return nil
}

func (m *SockclntDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Response
	return size
}
func (m *SockclntDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Response)
	return buf.Bytes(), nil
}
func (m *SockclntDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Response = buf.DecodeInt32()
	return nil
}

// TracePluginMsgIds defines message 'trace_plugin_msg_ids'.
type TracePluginMsgIds struct {
	PluginName string `binapi:"string[128],name=plugin_name" json:"plugin_name,omitempty"`
	FirstMsgID uint16 `binapi:"u16,name=first_msg_id" json:"first_msg_id,omitempty"`
	LastMsgID  uint16 `binapi:"u16,name=last_msg_id" json:"last_msg_id,omitempty"`
}

func (m *TracePluginMsgIds) Reset()               { *m = TracePluginMsgIds{} }
func (*TracePluginMsgIds) GetMessageName() string { return "trace_plugin_msg_ids" }
func (*TracePluginMsgIds) GetCrcString() string   { return "f476d3ce" }
func (*TracePluginMsgIds) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TracePluginMsgIds) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 128 // m.PluginName
	size += 2   // m.FirstMsgID
	size += 2   // m.LastMsgID
	return size
}
func (m *TracePluginMsgIds) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.PluginName, 128)
	buf.EncodeUint16(m.FirstMsgID)
	buf.EncodeUint16(m.LastMsgID)
	return buf.Bytes(), nil
}
func (m *TracePluginMsgIds) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.PluginName = buf.DecodeString(128)
	m.FirstMsgID = buf.DecodeUint16()
	m.LastMsgID = buf.DecodeUint16()
	return nil
}

func init() { file_memclnt_binapi_init() }
func file_memclnt_binapi_init() {
	api.RegisterMessage((*APIVersions)(nil), "api_versions_51077d14")
	api.RegisterMessage((*APIVersionsReply)(nil), "api_versions_reply_5f0d99d6")
	api.RegisterMessage((*ControlPing)(nil), "control_ping_51077d14")
	api.RegisterMessage((*ControlPingReply)(nil), "control_ping_reply_f6b0b8ca")
	api.RegisterMessage((*GetFirstMsgID)(nil), "get_first_msg_id_ebf79a66")
	api.RegisterMessage((*GetFirstMsgIDReply)(nil), "get_first_msg_id_reply_7d337472")
	api.RegisterMessage((*MemclntCreate)(nil), "memclnt_create_9c5e1c2f")
	api.RegisterMessage((*MemclntCreateReply)(nil), "memclnt_create_reply_42ec4560")
	api.RegisterMessage((*MemclntDelete)(nil), "memclnt_delete_7e1c04e3")
	api.RegisterMessage((*MemclntDeleteReply)(nil), "memclnt_delete_reply_3d3b6312")
	api.RegisterMessage((*MemclntKeepalive)(nil), "memclnt_keepalive_51077d14")
	api.RegisterMessage((*MemclntKeepaliveReply)(nil), "memclnt_keepalive_reply_e8d4e804")
	api.RegisterMessage((*MemclntReadTimeout)(nil), "memclnt_read_timeout_c3a3a452")
	api.RegisterMessage((*MemclntRxThreadSuspend)(nil), "memclnt_rx_thread_suspend_c3a3a452")
	api.RegisterMessage((*RPCCall)(nil), "rpc_call_7e8a2c95")
	api.RegisterMessage((*RPCCallReply)(nil), "rpc_call_reply_e8d4e804")
	api.RegisterMessage((*RxThreadExit)(nil), "rx_thread_exit_c3a3a452")
	api.RegisterMessage((*SockInitShm)(nil), "sock_init_shm_51646d92")
	api.RegisterMessage((*SockInitShmReply)(nil), "sock_init_shm_reply_e8d4e804")
	api.RegisterMessage((*SockclntCreate)(nil), "sockclnt_create_455fb9c4")
	api.RegisterMessage((*SockclntCreateReply)(nil), "sockclnt_create_reply_35166268")
	api.RegisterMessage((*SockclntDelete)(nil), "sockclnt_delete_8ac76db6")
	api.RegisterMessage((*SockclntDeleteReply)(nil), "sockclnt_delete_reply_8f38b1ee")
	api.RegisterMessage((*TracePluginMsgIds)(nil), "trace_plugin_msg_ids_f476d3ce")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*APIVersions)(nil),
		(*APIVersionsReply)(nil),
		(*ControlPing)(nil),
		(*ControlPingReply)(nil),
		(*GetFirstMsgID)(nil),
		(*GetFirstMsgIDReply)(nil),
		(*MemclntCreate)(nil),
		(*MemclntCreateReply)(nil),
		(*MemclntDelete)(nil),
		(*MemclntDeleteReply)(nil),
		(*MemclntKeepalive)(nil),
		(*MemclntKeepaliveReply)(nil),
		(*MemclntReadTimeout)(nil),
		(*MemclntRxThreadSuspend)(nil),
		(*RPCCall)(nil),
		(*RPCCallReply)(nil),
		(*RxThreadExit)(nil),
		(*SockInitShm)(nil),
		(*SockInitShmReply)(nil),
		(*SockclntCreate)(nil),
		(*SockclntCreateReply)(nil),
		(*SockclntDelete)(nil),
		(*SockclntDeleteReply)(nil),
		(*TracePluginMsgIds)(nil),
	}
}
