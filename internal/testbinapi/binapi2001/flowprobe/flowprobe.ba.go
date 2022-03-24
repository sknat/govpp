// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              20.01
// source: .vppapi/plugins/flowprobe.api.json

// Package flowprobe contains generated bindings for API file flowprobe.api.
//
// Contents:
//   1 alias
//   8 enums
//   4 messages
//
package flowprobe

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "flowprobe"
	APIVersion = "1.0.0"
	VersionCrc = 0xbb4dfc0d
)

// FlowprobeRecordFlags defines enum 'flowprobe_record_flags'.
type FlowprobeRecordFlags uint8

const (
	FLOWPROBE_RECORD_FLAG_L2 FlowprobeRecordFlags = 1
	FLOWPROBE_RECORD_FLAG_L3 FlowprobeRecordFlags = 2
	FLOWPROBE_RECORD_FLAG_L4 FlowprobeRecordFlags = 4
)

var (
	FlowprobeRecordFlags_name = map[uint8]string{
		1: "FLOWPROBE_RECORD_FLAG_L2",
		2: "FLOWPROBE_RECORD_FLAG_L3",
		4: "FLOWPROBE_RECORD_FLAG_L4",
	}
	FlowprobeRecordFlags_value = map[string]uint8{
		"FLOWPROBE_RECORD_FLAG_L2": 1,
		"FLOWPROBE_RECORD_FLAG_L3": 2,
		"FLOWPROBE_RECORD_FLAG_L4": 4,
	}
)

func (x FlowprobeRecordFlags) String() string {
	s, ok := FlowprobeRecordFlags_name[uint8(x)]
	if ok {
		return s
	}
	str := func(n uint8) string {
		s, ok := FlowprobeRecordFlags_name[uint8(n)]
		if ok {
			return s
		}
		return "FlowprobeRecordFlags(" + strconv.Itoa(int(n)) + ")"
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

// FlowprobeWhichFlags defines enum 'flowprobe_which_flags'.
type FlowprobeWhichFlags uint8

const (
	FLOWPROBE_WHICH_FLAG_IP4 FlowprobeWhichFlags = 1
	FLOWPROBE_WHICH_FLAG_L2  FlowprobeWhichFlags = 2
	FLOWPROBE_WHICH_FLAG_IP6 FlowprobeWhichFlags = 4
)

var (
	FlowprobeWhichFlags_name = map[uint8]string{
		1: "FLOWPROBE_WHICH_FLAG_IP4",
		2: "FLOWPROBE_WHICH_FLAG_L2",
		4: "FLOWPROBE_WHICH_FLAG_IP6",
	}
	FlowprobeWhichFlags_value = map[string]uint8{
		"FLOWPROBE_WHICH_FLAG_IP4": 1,
		"FLOWPROBE_WHICH_FLAG_L2":  2,
		"FLOWPROBE_WHICH_FLAG_IP6": 4,
	}
)

func (x FlowprobeWhichFlags) String() string {
	s, ok := FlowprobeWhichFlags_name[uint8(x)]
	if ok {
		return s
	}
	str := func(n uint8) string {
		s, ok := FlowprobeWhichFlags_name[uint8(n)]
		if ok {
			return s
		}
		return "FlowprobeWhichFlags(" + strconv.Itoa(int(n)) + ")"
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

// IfStatusFlags defines enum 'if_status_flags'.
type IfStatusFlags uint32

const (
	IF_STATUS_API_FLAG_ADMIN_UP IfStatusFlags = 1
	IF_STATUS_API_FLAG_LINK_UP  IfStatusFlags = 2
)

var (
	IfStatusFlags_name = map[uint32]string{
		1: "IF_STATUS_API_FLAG_ADMIN_UP",
		2: "IF_STATUS_API_FLAG_LINK_UP",
	}
	IfStatusFlags_value = map[string]uint32{
		"IF_STATUS_API_FLAG_ADMIN_UP": 1,
		"IF_STATUS_API_FLAG_LINK_UP":  2,
	}
)

func (x IfStatusFlags) String() string {
	s, ok := IfStatusFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := IfStatusFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "IfStatusFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// IfType defines enum 'if_type'.
type IfType uint32

const (
	IF_API_TYPE_HARDWARE IfType = 1
	IF_API_TYPE_SUB      IfType = 2
	IF_API_TYPE_P2P      IfType = 3
	IF_API_TYPE_PIPE     IfType = 4
)

var (
	IfType_name = map[uint32]string{
		1: "IF_API_TYPE_HARDWARE",
		2: "IF_API_TYPE_SUB",
		3: "IF_API_TYPE_P2P",
		4: "IF_API_TYPE_PIPE",
	}
	IfType_value = map[string]uint32{
		"IF_API_TYPE_HARDWARE": 1,
		"IF_API_TYPE_SUB":      2,
		"IF_API_TYPE_P2P":      3,
		"IF_API_TYPE_PIPE":     4,
	}
)

func (x IfType) String() string {
	s, ok := IfType_name[uint32(x)]
	if ok {
		return s
	}
	return "IfType(" + strconv.Itoa(int(x)) + ")"
}

// LinkDuplex defines enum 'link_duplex'.
type LinkDuplex uint32

const (
	LINK_DUPLEX_API_UNKNOWN LinkDuplex = 0
	LINK_DUPLEX_API_HALF    LinkDuplex = 1
	LINK_DUPLEX_API_FULL    LinkDuplex = 2
)

var (
	LinkDuplex_name = map[uint32]string{
		0: "LINK_DUPLEX_API_UNKNOWN",
		1: "LINK_DUPLEX_API_HALF",
		2: "LINK_DUPLEX_API_FULL",
	}
	LinkDuplex_value = map[string]uint32{
		"LINK_DUPLEX_API_UNKNOWN": 0,
		"LINK_DUPLEX_API_HALF":    1,
		"LINK_DUPLEX_API_FULL":    2,
	}
)

func (x LinkDuplex) String() string {
	s, ok := LinkDuplex_name[uint32(x)]
	if ok {
		return s
	}
	return "LinkDuplex(" + strconv.Itoa(int(x)) + ")"
}

// MtuProto defines enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 1
	MTU_PROTO_API_IP4  MtuProto = 2
	MTU_PROTO_API_IP6  MtuProto = 3
	MTU_PROTO_API_MPLS MtuProto = 4
	MTU_PROTO_API_N    MtuProto = 5
)

var (
	MtuProto_name = map[uint32]string{
		1: "MTU_PROTO_API_L3",
		2: "MTU_PROTO_API_IP4",
		3: "MTU_PROTO_API_IP6",
		4: "MTU_PROTO_API_MPLS",
		5: "MTU_PROTO_API_N",
	}
	MtuProto_value = map[string]uint32{
		"MTU_PROTO_API_L3":   1,
		"MTU_PROTO_API_IP4":  2,
		"MTU_PROTO_API_IP6":  3,
		"MTU_PROTO_API_MPLS": 4,
		"MTU_PROTO_API_N":    5,
	}
)

func (x MtuProto) String() string {
	s, ok := MtuProto_name[uint32(x)]
	if ok {
		return s
	}
	return "MtuProto(" + strconv.Itoa(int(x)) + ")"
}

// RxMode defines enum 'rx_mode'.
type RxMode uint32

const (
	RX_MODE_API_UNKNOWN   RxMode = 0
	RX_MODE_API_POLLING   RxMode = 1
	RX_MODE_API_INTERRUPT RxMode = 2
	RX_MODE_API_ADAPTIVE  RxMode = 3
	RX_MODE_API_DEFAULT   RxMode = 4
)

var (
	RxMode_name = map[uint32]string{
		0: "RX_MODE_API_UNKNOWN",
		1: "RX_MODE_API_POLLING",
		2: "RX_MODE_API_INTERRUPT",
		3: "RX_MODE_API_ADAPTIVE",
		4: "RX_MODE_API_DEFAULT",
	}
	RxMode_value = map[string]uint32{
		"RX_MODE_API_UNKNOWN":   0,
		"RX_MODE_API_POLLING":   1,
		"RX_MODE_API_INTERRUPT": 2,
		"RX_MODE_API_ADAPTIVE":  3,
		"RX_MODE_API_DEFAULT":   4,
	}
)

func (x RxMode) String() string {
	s, ok := RxMode_name[uint32(x)]
	if ok {
		return s
	}
	return "RxMode(" + strconv.Itoa(int(x)) + ")"
}

// SubIfFlags defines enum 'sub_if_flags'.
type SubIfFlags uint32

const (
	SUB_IF_API_FLAG_NO_TAGS           SubIfFlags = 1
	SUB_IF_API_FLAG_ONE_TAG           SubIfFlags = 2
	SUB_IF_API_FLAG_TWO_TAGS          SubIfFlags = 4
	SUB_IF_API_FLAG_DOT1AD            SubIfFlags = 8
	SUB_IF_API_FLAG_EXACT_MATCH       SubIfFlags = 16
	SUB_IF_API_FLAG_DEFAULT           SubIfFlags = 32
	SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY SubIfFlags = 64
	SUB_IF_API_FLAG_INNER_VLAN_ID_ANY SubIfFlags = 128
	SUB_IF_API_FLAG_MASK_VNET         SubIfFlags = 254
	SUB_IF_API_FLAG_DOT1AH            SubIfFlags = 256
)

var (
	SubIfFlags_name = map[uint32]string{
		1:   "SUB_IF_API_FLAG_NO_TAGS",
		2:   "SUB_IF_API_FLAG_ONE_TAG",
		4:   "SUB_IF_API_FLAG_TWO_TAGS",
		8:   "SUB_IF_API_FLAG_DOT1AD",
		16:  "SUB_IF_API_FLAG_EXACT_MATCH",
		32:  "SUB_IF_API_FLAG_DEFAULT",
		64:  "SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY",
		128: "SUB_IF_API_FLAG_INNER_VLAN_ID_ANY",
		254: "SUB_IF_API_FLAG_MASK_VNET",
		256: "SUB_IF_API_FLAG_DOT1AH",
	}
	SubIfFlags_value = map[string]uint32{
		"SUB_IF_API_FLAG_NO_TAGS":           1,
		"SUB_IF_API_FLAG_ONE_TAG":           2,
		"SUB_IF_API_FLAG_TWO_TAGS":          4,
		"SUB_IF_API_FLAG_DOT1AD":            8,
		"SUB_IF_API_FLAG_EXACT_MATCH":       16,
		"SUB_IF_API_FLAG_DEFAULT":           32,
		"SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY": 64,
		"SUB_IF_API_FLAG_INNER_VLAN_ID_ANY": 128,
		"SUB_IF_API_FLAG_MASK_VNET":         254,
		"SUB_IF_API_FLAG_DOT1AH":            256,
	}
)

func (x SubIfFlags) String() string {
	s, ok := SubIfFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := SubIfFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "SubIfFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// InterfaceIndex defines alias 'interface_index'.
type InterfaceIndex uint32

// FlowprobeParams defines message 'flowprobe_params'.
type FlowprobeParams struct {
	RecordFlags  FlowprobeRecordFlags `binapi:"flowprobe_record_flags,name=record_flags" json:"record_flags,omitempty"`
	ActiveTimer  uint32               `binapi:"u32,name=active_timer" json:"active_timer,omitempty"`
	PassiveTimer uint32               `binapi:"u32,name=passive_timer" json:"passive_timer,omitempty"`
}

func (m *FlowprobeParams) Reset()               { *m = FlowprobeParams{} }
func (*FlowprobeParams) GetMessageName() string { return "flowprobe_params" }
func (*FlowprobeParams) GetCrcString() string   { return "baa46c09" }
func (*FlowprobeParams) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (m *FlowprobeParams) GetRetVal() error {
	return nil
}

func (m *FlowprobeParams) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.RecordFlags
	size += 4 // m.ActiveTimer
	size += 4 // m.PassiveTimer
	return size
}
func (m *FlowprobeParams) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.RecordFlags))
	buf.EncodeUint32(m.ActiveTimer)
	buf.EncodeUint32(m.PassiveTimer)
	return buf.Bytes(), nil
}
func (m *FlowprobeParams) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.RecordFlags = FlowprobeRecordFlags(buf.DecodeUint8())
	m.ActiveTimer = buf.DecodeUint32()
	m.PassiveTimer = buf.DecodeUint32()
	return nil
}

// FlowprobeParamsReply defines message 'flowprobe_params_reply'.
type FlowprobeParamsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *FlowprobeParamsReply) Reset()               { *m = FlowprobeParamsReply{} }
func (*FlowprobeParamsReply) GetMessageName() string { return "flowprobe_params_reply" }
func (*FlowprobeParamsReply) GetCrcString() string   { return "e8d4e804" }
func (*FlowprobeParamsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *FlowprobeParamsReply) GetRetVal() error {
	return api.RetvalToVPPApiError(int32(m.Retval))
}

func (m *FlowprobeParamsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *FlowprobeParamsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *FlowprobeParamsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// FlowprobeTxInterfaceAddDel defines message 'flowprobe_tx_interface_add_del'.
type FlowprobeTxInterfaceAddDel struct {
	IsAdd     bool                `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Which     FlowprobeWhichFlags `binapi:"flowprobe_which_flags,name=which" json:"which,omitempty"`
	SwIfIndex InterfaceIndex      `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *FlowprobeTxInterfaceAddDel) Reset()               { *m = FlowprobeTxInterfaceAddDel{} }
func (*FlowprobeTxInterfaceAddDel) GetMessageName() string { return "flowprobe_tx_interface_add_del" }
func (*FlowprobeTxInterfaceAddDel) GetCrcString() string   { return "b782c976" }
func (*FlowprobeTxInterfaceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (m *FlowprobeTxInterfaceAddDel) GetRetVal() error {
	return nil
}

func (m *FlowprobeTxInterfaceAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 1 // m.Which
	size += 4 // m.SwIfIndex
	return size
}
func (m *FlowprobeTxInterfaceAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint8(uint8(m.Which))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *FlowprobeTxInterfaceAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Which = FlowprobeWhichFlags(buf.DecodeUint8())
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	return nil
}

// FlowprobeTxInterfaceAddDelReply defines message 'flowprobe_tx_interface_add_del_reply'.
type FlowprobeTxInterfaceAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *FlowprobeTxInterfaceAddDelReply) Reset() { *m = FlowprobeTxInterfaceAddDelReply{} }
func (*FlowprobeTxInterfaceAddDelReply) GetMessageName() string {
	return "flowprobe_tx_interface_add_del_reply"
}
func (*FlowprobeTxInterfaceAddDelReply) GetCrcString() string { return "e8d4e804" }
func (*FlowprobeTxInterfaceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *FlowprobeTxInterfaceAddDelReply) GetRetVal() error {
	return api.RetvalToVPPApiError(int32(m.Retval))
}

func (m *FlowprobeTxInterfaceAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *FlowprobeTxInterfaceAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *FlowprobeTxInterfaceAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_flowprobe_binapi_init() }
func file_flowprobe_binapi_init() {
	api.RegisterMessage((*FlowprobeParams)(nil), "flowprobe_params_baa46c09")
	api.RegisterMessage((*FlowprobeParamsReply)(nil), "flowprobe_params_reply_e8d4e804")
	api.RegisterMessage((*FlowprobeTxInterfaceAddDel)(nil), "flowprobe_tx_interface_add_del_b782c976")
	api.RegisterMessage((*FlowprobeTxInterfaceAddDelReply)(nil), "flowprobe_tx_interface_add_del_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*FlowprobeParams)(nil),
		(*FlowprobeParamsReply)(nil),
		(*FlowprobeTxInterfaceAddDel)(nil),
		(*FlowprobeTxInterfaceAddDelReply)(nil),
	}
}
