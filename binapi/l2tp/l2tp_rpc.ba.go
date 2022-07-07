// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package l2tp

import (
	"context"
	"fmt"
	"io"

	api "go.fd.io/govpp/api"
	memclnt "go.fd.io/govpp/binapi/memclnt"
)

// RPCService defines RPC service l2tp.
type RPCService interface {
	L2tpv3CreateTunnel(ctx context.Context, in *L2tpv3CreateTunnel) (*L2tpv3CreateTunnelReply, error)
	L2tpv3InterfaceEnableDisable(ctx context.Context, in *L2tpv3InterfaceEnableDisable) (*L2tpv3InterfaceEnableDisableReply, error)
	L2tpv3SetLookupKey(ctx context.Context, in *L2tpv3SetLookupKey) (*L2tpv3SetLookupKeyReply, error)
	L2tpv3SetTunnelCookies(ctx context.Context, in *L2tpv3SetTunnelCookies) (*L2tpv3SetTunnelCookiesReply, error)
	SwIfL2tpv3TunnelDump(ctx context.Context, in *SwIfL2tpv3TunnelDump) (RPCService_SwIfL2tpv3TunnelDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) L2tpv3CreateTunnel(ctx context.Context, in *L2tpv3CreateTunnel) (*L2tpv3CreateTunnelReply, error) {
	out := new(L2tpv3CreateTunnelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2tpv3InterfaceEnableDisable(ctx context.Context, in *L2tpv3InterfaceEnableDisable) (*L2tpv3InterfaceEnableDisableReply, error) {
	out := new(L2tpv3InterfaceEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2tpv3SetLookupKey(ctx context.Context, in *L2tpv3SetLookupKey) (*L2tpv3SetLookupKeyReply, error) {
	out := new(L2tpv3SetLookupKeyReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2tpv3SetTunnelCookies(ctx context.Context, in *L2tpv3SetTunnelCookies) (*L2tpv3SetTunnelCookiesReply, error) {
	out := new(L2tpv3SetTunnelCookiesReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwIfL2tpv3TunnelDump(ctx context.Context, in *SwIfL2tpv3TunnelDump) (RPCService_SwIfL2tpv3TunnelDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SwIfL2tpv3TunnelDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SwIfL2tpv3TunnelDumpClient interface {
	Recv() (*SwIfL2tpv3TunnelDetails, error)
	api.Stream
}

type serviceClient_SwIfL2tpv3TunnelDumpClient struct {
	api.Stream
}

func (c *serviceClient_SwIfL2tpv3TunnelDumpClient) Recv() (*SwIfL2tpv3TunnelDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SwIfL2tpv3TunnelDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
