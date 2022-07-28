// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.6.0-dev
//  VPP:              22.02-release
// source: /usr/share/vpp/api/plugins/af_xdp.api.json

// Package af_xdp contains generated bindings for API file af_xdp.api.
//
// Contents:
//   2 enums
//   6 messages
//
package af_xdp

import (
	"strconv"

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
	APIFile    = "af_xdp"
	APIVersion = "1.0.0"
	VersionCrc = 0x346626ce
)

// AfXdpMode defines enum 'af_xdp_mode'.
type AfXdpMode uint32

const (
	AF_XDP_API_MODE_AUTO      AfXdpMode = 0
	AF_XDP_API_MODE_COPY      AfXdpMode = 1
	AF_XDP_API_MODE_ZERO_COPY AfXdpMode = 2
)

var (
	AfXdpMode_name = map[uint32]string{
		0: "AF_XDP_API_MODE_AUTO",
		1: "AF_XDP_API_MODE_COPY",
		2: "AF_XDP_API_MODE_ZERO_COPY",
	}
	AfXdpMode_value = map[string]uint32{
		"AF_XDP_API_MODE_AUTO":      0,
		"AF_XDP_API_MODE_COPY":      1,
		"AF_XDP_API_MODE_ZERO_COPY": 2,
	}
)

func (x AfXdpMode) String() string {
	s, ok := AfXdpMode_name[uint32(x)]
	if ok {
		return s
	}
	return "AfXdpMode(" + strconv.Itoa(int(x)) + ")"
}

// AfXdpFlag defines enum 'af_xdp_flag'.
type AfXdpFlag uint8

const (
	AF_XDP_API_FLAGS_NO_SYSCALL_LOCK AfXdpFlag = 1
)

var (
	AfXdpFlag_name = map[uint8]string{
		1: "AF_XDP_API_FLAGS_NO_SYSCALL_LOCK",
	}
	AfXdpFlag_value = map[string]uint8{
		"AF_XDP_API_FLAGS_NO_SYSCALL_LOCK": 1,
	}
)

func (x AfXdpFlag) String() string {
	s, ok := AfXdpFlag_name[uint8(x)]
	if ok {
		return s
	}
	str := func(n uint8) string {
		s, ok := AfXdpFlag_name[uint8(n)]
		if ok {
			return s
		}
		return "AfXdpFlag(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint8(0); i <= 8; i++ {
		val := uint8(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint8(x))
	}
	return s
}

// AfXdpCreate defines message 'af_xdp_create'.
type AfXdpCreate struct {
	HostIf  string    `binapi:"string[64],name=host_if" json:"host_if,omitempty"`
	Name    string    `binapi:"string[64],name=name" json:"name,omitempty"`
	RxqNum  uint16    `binapi:"u16,name=rxq_num,default=1" json:"rxq_num,omitempty"`
	RxqSize uint16    `binapi:"u16,name=rxq_size,default=0" json:"rxq_size,omitempty"`
	TxqSize uint16    `binapi:"u16,name=txq_size,default=0" json:"txq_size,omitempty"`
	Mode    AfXdpMode `binapi:"af_xdp_mode,name=mode,default=0" json:"mode,omitempty"`
	Flags   AfXdpFlag `binapi:"af_xdp_flag,name=flags,default=0" json:"flags,omitempty"`
	Prog    string    `binapi:"string[256],name=prog" json:"prog,omitempty"`
}

func (m *AfXdpCreate) Reset()               { *m = AfXdpCreate{} }
func (*AfXdpCreate) GetMessageName() string { return "af_xdp_create" }
func (*AfXdpCreate) GetCrcString() string   { return "21226c99" }
func (*AfXdpCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AfXdpCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64  // m.HostIf
	size += 64  // m.Name
	size += 2   // m.RxqNum
	size += 2   // m.RxqSize
	size += 2   // m.TxqSize
	size += 4   // m.Mode
	size += 1   // m.Flags
	size += 256 // m.Prog
	return size
}
func (m *AfXdpCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.HostIf, 64)
	buf.EncodeString(m.Name, 64)
	buf.EncodeUint16(m.RxqNum)
	buf.EncodeUint16(m.RxqSize)
	buf.EncodeUint16(m.TxqSize)
	buf.EncodeUint32(uint32(m.Mode))
	buf.EncodeUint8(uint8(m.Flags))
	buf.EncodeString(m.Prog, 256)
	return buf.Bytes(), nil
}
func (m *AfXdpCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.HostIf = buf.DecodeString(64)
	m.Name = buf.DecodeString(64)
	m.RxqNum = buf.DecodeUint16()
	m.RxqSize = buf.DecodeUint16()
	m.TxqSize = buf.DecodeUint16()
	m.Mode = AfXdpMode(buf.DecodeUint32())
	m.Flags = AfXdpFlag(buf.DecodeUint8())
	m.Prog = buf.DecodeString(256)
	return nil
}

// AfXdpCreateReply defines message 'af_xdp_create_reply'.
type AfXdpCreateReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *AfXdpCreateReply) Reset()               { *m = AfXdpCreateReply{} }
func (*AfXdpCreateReply) GetMessageName() string { return "af_xdp_create_reply" }
func (*AfXdpCreateReply) GetCrcString() string   { return "5383d31f" }
func (*AfXdpCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AfXdpCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *AfXdpCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *AfXdpCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// AfXdpCreateV2 defines message 'af_xdp_create_v2'.
type AfXdpCreateV2 struct {
	HostIf    string    `binapi:"string[64],name=host_if" json:"host_if,omitempty"`
	Name      string    `binapi:"string[64],name=name" json:"name,omitempty"`
	RxqNum    uint16    `binapi:"u16,name=rxq_num,default=1" json:"rxq_num,omitempty"`
	RxqSize   uint16    `binapi:"u16,name=rxq_size,default=0" json:"rxq_size,omitempty"`
	TxqSize   uint16    `binapi:"u16,name=txq_size,default=0" json:"txq_size,omitempty"`
	Mode      AfXdpMode `binapi:"af_xdp_mode,name=mode,default=0" json:"mode,omitempty"`
	Flags     AfXdpFlag `binapi:"af_xdp_flag,name=flags,default=0" json:"flags,omitempty"`
	Prog      string    `binapi:"string[256],name=prog" json:"prog,omitempty"`
	Namespace string    `binapi:"string[64],name=namespace" json:"namespace,omitempty"`
}

func (m *AfXdpCreateV2) Reset()               { *m = AfXdpCreateV2{} }
func (*AfXdpCreateV2) GetMessageName() string { return "af_xdp_create_v2" }
func (*AfXdpCreateV2) GetCrcString() string   { return "e17ec2eb" }
func (*AfXdpCreateV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AfXdpCreateV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64  // m.HostIf
	size += 64  // m.Name
	size += 2   // m.RxqNum
	size += 2   // m.RxqSize
	size += 2   // m.TxqSize
	size += 4   // m.Mode
	size += 1   // m.Flags
	size += 256 // m.Prog
	size += 64  // m.Namespace
	return size
}
func (m *AfXdpCreateV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.HostIf, 64)
	buf.EncodeString(m.Name, 64)
	buf.EncodeUint16(m.RxqNum)
	buf.EncodeUint16(m.RxqSize)
	buf.EncodeUint16(m.TxqSize)
	buf.EncodeUint32(uint32(m.Mode))
	buf.EncodeUint8(uint8(m.Flags))
	buf.EncodeString(m.Prog, 256)
	buf.EncodeString(m.Namespace, 64)
	return buf.Bytes(), nil
}
func (m *AfXdpCreateV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.HostIf = buf.DecodeString(64)
	m.Name = buf.DecodeString(64)
	m.RxqNum = buf.DecodeUint16()
	m.RxqSize = buf.DecodeUint16()
	m.TxqSize = buf.DecodeUint16()
	m.Mode = AfXdpMode(buf.DecodeUint32())
	m.Flags = AfXdpFlag(buf.DecodeUint8())
	m.Prog = buf.DecodeString(256)
	m.Namespace = buf.DecodeString(64)
	return nil
}

// AfXdpCreateV2Reply defines message 'af_xdp_create_v2_reply'.
type AfXdpCreateV2Reply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *AfXdpCreateV2Reply) Reset()               { *m = AfXdpCreateV2Reply{} }
func (*AfXdpCreateV2Reply) GetMessageName() string { return "af_xdp_create_v2_reply" }
func (*AfXdpCreateV2Reply) GetCrcString() string   { return "5383d31f" }
func (*AfXdpCreateV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AfXdpCreateV2Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *AfXdpCreateV2Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *AfXdpCreateV2Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// AfXdpDelete defines message 'af_xdp_delete'.
type AfXdpDelete struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *AfXdpDelete) Reset()               { *m = AfXdpDelete{} }
func (*AfXdpDelete) GetMessageName() string { return "af_xdp_delete" }
func (*AfXdpDelete) GetCrcString() string   { return "f9e6675e" }
func (*AfXdpDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AfXdpDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *AfXdpDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *AfXdpDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// AfXdpDeleteReply defines message 'af_xdp_delete_reply'.
type AfXdpDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AfXdpDeleteReply) Reset()               { *m = AfXdpDeleteReply{} }
func (*AfXdpDeleteReply) GetMessageName() string { return "af_xdp_delete_reply" }
func (*AfXdpDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*AfXdpDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AfXdpDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AfXdpDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AfXdpDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_af_xdp_binapi_init() }
func file_af_xdp_binapi_init() {
	api.RegisterMessage((*AfXdpCreate)(nil), "af_xdp_create_21226c99")
	api.RegisterMessage((*AfXdpCreateReply)(nil), "af_xdp_create_reply_5383d31f")
	api.RegisterMessage((*AfXdpCreateV2)(nil), "af_xdp_create_v2_e17ec2eb")
	api.RegisterMessage((*AfXdpCreateV2Reply)(nil), "af_xdp_create_v2_reply_5383d31f")
	api.RegisterMessage((*AfXdpDelete)(nil), "af_xdp_delete_f9e6675e")
	api.RegisterMessage((*AfXdpDeleteReply)(nil), "af_xdp_delete_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AfXdpCreate)(nil),
		(*AfXdpCreateReply)(nil),
		(*AfXdpCreateV2)(nil),
		(*AfXdpCreateV2Reply)(nil),
		(*AfXdpDelete)(nil),
		(*AfXdpDeleteReply)(nil),
	}
}
