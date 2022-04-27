// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              21.06-release
// source: /usr/share/vpp/api/plugins/graph.api.json

// Package graph contains generated bindings for API file graph.api.
//
// Contents:
//   1 enum
//   3 messages
//
package graph

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
	APIFile    = "graph"
	APIVersion = "1.0.0"
	VersionCrc = 0xa0b3fd1c
)

// NodeFlag defines enum 'node_flag'.
type NodeFlag uint32

const (
	NODE_FLAG_FRAME_NO_FREE_AFTER_DISPATCH          NodeFlag = 1
	NODE_FLAG_IS_OUTPUT                             NodeFlag = 2
	NODE_FLAG_IS_DROP                               NodeFlag = 4
	NODE_FLAG_IS_PUNT                               NodeFlag = 8
	NODE_FLAG_IS_HANDOFF                            NodeFlag = 16
	NODE_FLAG_TRACE                                 NodeFlag = 32
	NODE_FLAG_SWITCH_FROM_INTERRUPT_TO_POLLING_MODE NodeFlag = 64
	NODE_FLAG_SWITCH_FROM_POLLING_TO_INTERRUPT_MODE NodeFlag = 128
	NODE_FLAG_TRACE_SUPPORTED                       NodeFlag = 256
)

var (
	NodeFlag_name = map[uint32]string{
		1:   "NODE_FLAG_FRAME_NO_FREE_AFTER_DISPATCH",
		2:   "NODE_FLAG_IS_OUTPUT",
		4:   "NODE_FLAG_IS_DROP",
		8:   "NODE_FLAG_IS_PUNT",
		16:  "NODE_FLAG_IS_HANDOFF",
		32:  "NODE_FLAG_TRACE",
		64:  "NODE_FLAG_SWITCH_FROM_INTERRUPT_TO_POLLING_MODE",
		128: "NODE_FLAG_SWITCH_FROM_POLLING_TO_INTERRUPT_MODE",
		256: "NODE_FLAG_TRACE_SUPPORTED",
	}
	NodeFlag_value = map[string]uint32{
		"NODE_FLAG_FRAME_NO_FREE_AFTER_DISPATCH":          1,
		"NODE_FLAG_IS_OUTPUT":                             2,
		"NODE_FLAG_IS_DROP":                               4,
		"NODE_FLAG_IS_PUNT":                               8,
		"NODE_FLAG_IS_HANDOFF":                            16,
		"NODE_FLAG_TRACE":                                 32,
		"NODE_FLAG_SWITCH_FROM_INTERRUPT_TO_POLLING_MODE": 64,
		"NODE_FLAG_SWITCH_FROM_POLLING_TO_INTERRUPT_MODE": 128,
		"NODE_FLAG_TRACE_SUPPORTED":                       256,
	}
)

func (x NodeFlag) String() string {
	s, ok := NodeFlag_name[uint32(x)]
	if ok {
		return s
	}
	return "NodeFlag(" + strconv.Itoa(int(x)) + ")"
}

// GraphNodeDetails defines message 'graph_node_details'.
type GraphNodeDetails struct {
	Index   uint32   `binapi:"u32,name=index" json:"index,omitempty"`
	Name    string   `binapi:"string[64],name=name" json:"name,omitempty"`
	Flags   NodeFlag `binapi:"node_flag,name=flags" json:"flags,omitempty"`
	NArcs   uint32   `binapi:"u32,name=n_arcs" json:"-"`
	ArcsOut []uint32 `binapi:"u32[n_arcs],name=arcs_out" json:"arcs_out,omitempty"`
}

func (m *GraphNodeDetails) Reset()               { *m = GraphNodeDetails{} }
func (*GraphNodeDetails) GetMessageName() string { return "graph_node_details" }
func (*GraphNodeDetails) GetCrcString() string   { return "ac762018" }
func (*GraphNodeDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GraphNodeDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                  // m.Index
	size += 64                 // m.Name
	size += 4                  // m.Flags
	size += 4                  // m.NArcs
	size += 4 * len(m.ArcsOut) // m.ArcsOut
	return size
}
func (m *GraphNodeDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Index)
	buf.EncodeString(m.Name, 64)
	buf.EncodeUint32(uint32(m.Flags))
	buf.EncodeUint32(uint32(len(m.ArcsOut)))
	for i := 0; i < len(m.ArcsOut); i++ {
		var x uint32
		if i < len(m.ArcsOut) {
			x = uint32(m.ArcsOut[i])
		}
		buf.EncodeUint32(x)
	}
	return buf.Bytes(), nil
}
func (m *GraphNodeDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Index = buf.DecodeUint32()
	m.Name = buf.DecodeString(64)
	m.Flags = NodeFlag(buf.DecodeUint32())
	m.NArcs = buf.DecodeUint32()
	m.ArcsOut = make([]uint32, m.NArcs)
	for i := 0; i < len(m.ArcsOut); i++ {
		m.ArcsOut[i] = buf.DecodeUint32()
	}
	return nil
}

// GraphNodeGet defines message 'graph_node_get'.
type GraphNodeGet struct {
	Cursor   uint32   `binapi:"u32,name=cursor" json:"cursor,omitempty"`
	Index    uint32   `binapi:"u32,name=index" json:"index,omitempty"`
	Name     string   `binapi:"string[64],name=name" json:"name,omitempty"`
	Flags    NodeFlag `binapi:"node_flag,name=flags" json:"flags,omitempty"`
	WantArcs bool     `binapi:"bool,name=want_arcs" json:"want_arcs,omitempty"`
}

func (m *GraphNodeGet) Reset()               { *m = GraphNodeGet{} }
func (*GraphNodeGet) GetMessageName() string { return "graph_node_get" }
func (*GraphNodeGet) GetCrcString() string   { return "39c8792e" }
func (*GraphNodeGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GraphNodeGet) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4  // m.Cursor
	size += 4  // m.Index
	size += 64 // m.Name
	size += 4  // m.Flags
	size += 1  // m.WantArcs
	return size
}
func (m *GraphNodeGet) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Cursor)
	buf.EncodeUint32(m.Index)
	buf.EncodeString(m.Name, 64)
	buf.EncodeUint32(uint32(m.Flags))
	buf.EncodeBool(m.WantArcs)
	return buf.Bytes(), nil
}
func (m *GraphNodeGet) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Cursor = buf.DecodeUint32()
	m.Index = buf.DecodeUint32()
	m.Name = buf.DecodeString(64)
	m.Flags = NodeFlag(buf.DecodeUint32())
	m.WantArcs = buf.DecodeBool()
	return nil
}

// GraphNodeGetReply defines message 'graph_node_get_reply'.
type GraphNodeGetReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	Cursor uint32 `binapi:"u32,name=cursor" json:"cursor,omitempty"`
}

func (m *GraphNodeGetReply) Reset()               { *m = GraphNodeGetReply{} }
func (*GraphNodeGetReply) GetMessageName() string { return "graph_node_get_reply" }
func (*GraphNodeGetReply) GetCrcString() string   { return "53b48f5d" }
func (*GraphNodeGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GraphNodeGetReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.Cursor
	return size
}
func (m *GraphNodeGetReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.Cursor)
	return buf.Bytes(), nil
}
func (m *GraphNodeGetReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Cursor = buf.DecodeUint32()
	return nil
}

func init() { file_graph_binapi_init() }
func file_graph_binapi_init() {
	api.RegisterMessage((*GraphNodeDetails)(nil), "graph_node_details_ac762018")
	api.RegisterMessage((*GraphNodeGet)(nil), "graph_node_get_39c8792e")
	api.RegisterMessage((*GraphNodeGetReply)(nil), "graph_node_get_reply_53b48f5d")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*GraphNodeDetails)(nil),
		(*GraphNodeGet)(nil),
		(*GraphNodeGetReply)(nil),
	}
}
