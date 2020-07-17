// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package stn

import (
	"context"
	"fmt"
	api "git.fd.io/govpp.git/api"
	vpe "git.fd.io/govpp.git/binapi/vpe"
	"io"
)

// RPCService defines RPC service  stn.
type RPCService interface {
	StnAddDelRule(ctx context.Context, in *StnAddDelRule) (*StnAddDelRuleReply, error)
	StnRulesDump(ctx context.Context, in *StnRulesDump) (RPCService_StnRulesDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) StnAddDelRule(ctx context.Context, in *StnAddDelRule) (*StnAddDelRuleReply, error) {
	out := new(StnAddDelRuleReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) StnRulesDump(ctx context.Context, in *StnRulesDump) (RPCService_StnRulesDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_StnRulesDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_StnRulesDumpClient interface {
	Recv() (*StnRulesDetails, error)
	api.Stream
}

type serviceClient_StnRulesDumpClient struct {
	api.Stream
}

func (c *serviceClient_StnRulesDumpClient) Recv() (*StnRulesDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *StnRulesDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
