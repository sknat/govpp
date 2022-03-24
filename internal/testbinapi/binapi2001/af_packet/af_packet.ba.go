// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              20.01
// source: .vppapi/core/af_packet.api.json

// Package af_packet contains generated bindings for API file af_packet.api.
//
// Contents:
//   2 aliases
//   6 enums
//   8 messages
//
package af_packet

import (
	"net"
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
	APIFile    = "af_packet"
	APIVersion = "2.0.0"
	VersionCrc = 0xba745e20
)

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

// MacAddress defines alias 'mac_address'.
type MacAddress [6]uint8

func ParseMacAddress(s string) (MacAddress, error) {
	var macaddr MacAddress
	mac, err := net.ParseMAC(s)
	if err != nil {
		return macaddr, err
	}
	copy(macaddr[:], mac[:])
	return macaddr, nil
}

func (x MacAddress) ToMAC() net.HardwareAddr {
	return net.HardwareAddr(x[:])
}

func (x MacAddress) String() string {
	return x.ToMAC().String()
}

func (x *MacAddress) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *MacAddress) UnmarshalText(text []byte) error {
	mac, err := ParseMacAddress(string(text))
	if err != nil {
		return err
	}
	*x = mac
	return nil
}

// AfPacketCreate defines message 'af_packet_create'.
type AfPacketCreate struct {
	HwAddr          MacAddress `binapi:"mac_address,name=hw_addr" json:"hw_addr,omitempty"`
	UseRandomHwAddr bool       `binapi:"bool,name=use_random_hw_addr" json:"use_random_hw_addr,omitempty"`
	HostIfName      string     `binapi:"string[64],name=host_if_name" json:"host_if_name,omitempty"`
}

func (m *AfPacketCreate) Reset()               { *m = AfPacketCreate{} }
func (*AfPacketCreate) GetMessageName() string { return "af_packet_create" }
func (*AfPacketCreate) GetCrcString() string   { return "a190415f" }
func (*AfPacketCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (m *AfPacketCreate) GetRetVal() error {
	return nil
}

func (m *AfPacketCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 6 // m.HwAddr
	size += 1     // m.UseRandomHwAddr
	size += 64    // m.HostIfName
	return size
}
func (m *AfPacketCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.HwAddr[:], 6)
	buf.EncodeBool(m.UseRandomHwAddr)
	buf.EncodeString(m.HostIfName, 64)
	return buf.Bytes(), nil
}
func (m *AfPacketCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.HwAddr[:], buf.DecodeBytes(6))
	m.UseRandomHwAddr = buf.DecodeBool()
	m.HostIfName = buf.DecodeString(64)
	return nil
}

// AfPacketCreateReply defines message 'af_packet_create_reply'.
type AfPacketCreateReply struct {
	Retval    int32          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *AfPacketCreateReply) Reset()               { *m = AfPacketCreateReply{} }
func (*AfPacketCreateReply) GetMessageName() string { return "af_packet_create_reply" }
func (*AfPacketCreateReply) GetCrcString() string   { return "5383d31f" }
func (*AfPacketCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *AfPacketCreateReply) GetRetVal() error {
	return api.RetvalToVPPApiError(int32(m.Retval))
}

func (m *AfPacketCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *AfPacketCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *AfPacketCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	return nil
}

// AfPacketDelete defines message 'af_packet_delete'.
type AfPacketDelete struct {
	HostIfName string `binapi:"string[64],name=host_if_name" json:"host_if_name,omitempty"`
}

func (m *AfPacketDelete) Reset()               { *m = AfPacketDelete{} }
func (*AfPacketDelete) GetMessageName() string { return "af_packet_delete" }
func (*AfPacketDelete) GetCrcString() string   { return "863fa648" }
func (*AfPacketDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (m *AfPacketDelete) GetRetVal() error {
	return nil
}

func (m *AfPacketDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.HostIfName
	return size
}
func (m *AfPacketDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.HostIfName, 64)
	return buf.Bytes(), nil
}
func (m *AfPacketDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.HostIfName = buf.DecodeString(64)
	return nil
}

// AfPacketDeleteReply defines message 'af_packet_delete_reply'.
type AfPacketDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AfPacketDeleteReply) Reset()               { *m = AfPacketDeleteReply{} }
func (*AfPacketDeleteReply) GetMessageName() string { return "af_packet_delete_reply" }
func (*AfPacketDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*AfPacketDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *AfPacketDeleteReply) GetRetVal() error {
	return api.RetvalToVPPApiError(int32(m.Retval))
}

func (m *AfPacketDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AfPacketDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AfPacketDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// AfPacketDetails defines message 'af_packet_details'.
type AfPacketDetails struct {
	SwIfIndex  InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	HostIfName string         `binapi:"string[64],name=host_if_name" json:"host_if_name,omitempty"`
}

func (m *AfPacketDetails) Reset()               { *m = AfPacketDetails{} }
func (*AfPacketDetails) GetMessageName() string { return "af_packet_details" }
func (*AfPacketDetails) GetCrcString() string   { return "58c7c042" }
func (*AfPacketDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *AfPacketDetails) GetRetVal() error {
	return nil
}

func (m *AfPacketDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4  // m.SwIfIndex
	size += 64 // m.HostIfName
	return size
}
func (m *AfPacketDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeString(m.HostIfName, 64)
	return buf.Bytes(), nil
}
func (m *AfPacketDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	m.HostIfName = buf.DecodeString(64)
	return nil
}

// AfPacketDump defines message 'af_packet_dump'.
type AfPacketDump struct{}

func (m *AfPacketDump) Reset()               { *m = AfPacketDump{} }
func (*AfPacketDump) GetMessageName() string { return "af_packet_dump" }
func (*AfPacketDump) GetCrcString() string   { return "51077d14" }
func (*AfPacketDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (m *AfPacketDump) GetRetVal() error {
	return nil
}

func (m *AfPacketDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *AfPacketDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *AfPacketDump) Unmarshal(b []byte) error {
	return nil
}

// AfPacketSetL4CksumOffload defines message 'af_packet_set_l4_cksum_offload'.
type AfPacketSetL4CksumOffload struct {
	SwIfIndex InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Set       bool           `binapi:"bool,name=set" json:"set,omitempty"`
}

func (m *AfPacketSetL4CksumOffload) Reset()               { *m = AfPacketSetL4CksumOffload{} }
func (*AfPacketSetL4CksumOffload) GetMessageName() string { return "af_packet_set_l4_cksum_offload" }
func (*AfPacketSetL4CksumOffload) GetCrcString() string   { return "319cd5c8" }
func (*AfPacketSetL4CksumOffload) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (m *AfPacketSetL4CksumOffload) GetRetVal() error {
	return nil
}

func (m *AfPacketSetL4CksumOffload) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.Set
	return size
}
func (m *AfPacketSetL4CksumOffload) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.Set)
	return buf.Bytes(), nil
}
func (m *AfPacketSetL4CksumOffload) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	m.Set = buf.DecodeBool()
	return nil
}

// AfPacketSetL4CksumOffloadReply defines message 'af_packet_set_l4_cksum_offload_reply'.
type AfPacketSetL4CksumOffloadReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AfPacketSetL4CksumOffloadReply) Reset() { *m = AfPacketSetL4CksumOffloadReply{} }
func (*AfPacketSetL4CksumOffloadReply) GetMessageName() string {
	return "af_packet_set_l4_cksum_offload_reply"
}
func (*AfPacketSetL4CksumOffloadReply) GetCrcString() string { return "e8d4e804" }
func (*AfPacketSetL4CksumOffloadReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (m *AfPacketSetL4CksumOffloadReply) GetRetVal() error {
	return api.RetvalToVPPApiError(int32(m.Retval))
}

func (m *AfPacketSetL4CksumOffloadReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AfPacketSetL4CksumOffloadReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AfPacketSetL4CksumOffloadReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_af_packet_binapi_init() }
func file_af_packet_binapi_init() {
	api.RegisterMessage((*AfPacketCreate)(nil), "af_packet_create_a190415f")
	api.RegisterMessage((*AfPacketCreateReply)(nil), "af_packet_create_reply_5383d31f")
	api.RegisterMessage((*AfPacketDelete)(nil), "af_packet_delete_863fa648")
	api.RegisterMessage((*AfPacketDeleteReply)(nil), "af_packet_delete_reply_e8d4e804")
	api.RegisterMessage((*AfPacketDetails)(nil), "af_packet_details_58c7c042")
	api.RegisterMessage((*AfPacketDump)(nil), "af_packet_dump_51077d14")
	api.RegisterMessage((*AfPacketSetL4CksumOffload)(nil), "af_packet_set_l4_cksum_offload_319cd5c8")
	api.RegisterMessage((*AfPacketSetL4CksumOffloadReply)(nil), "af_packet_set_l4_cksum_offload_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AfPacketCreate)(nil),
		(*AfPacketCreateReply)(nil),
		(*AfPacketDelete)(nil),
		(*AfPacketDeleteReply)(nil),
		(*AfPacketDetails)(nil),
		(*AfPacketDump)(nil),
		(*AfPacketSetL4CksumOffload)(nil),
		(*AfPacketSetL4CksumOffloadReply)(nil),
	}
}
