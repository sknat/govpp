// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package session

import (
	"context"
	"fmt"
	api "git.fd.io/govpp.git/api"
	vpe "git.fd.io/govpp.git/binapi/vpe"
	"io"
)

// RPCService defines RPC service  session.
type RPCService interface {
	AppAddCertKeyPair(ctx context.Context, in *AppAddCertKeyPair) (*AppAddCertKeyPairReply, error)
	AppAttach(ctx context.Context, in *AppAttach) (*AppAttachReply, error)
	AppDelCertKeyPair(ctx context.Context, in *AppDelCertKeyPair) (*AppDelCertKeyPairReply, error)
	AppNamespaceAddDel(ctx context.Context, in *AppNamespaceAddDel) (*AppNamespaceAddDelReply, error)
	AppWorkerAddDel(ctx context.Context, in *AppWorkerAddDel) (*AppWorkerAddDelReply, error)
	ApplicationDetach(ctx context.Context, in *ApplicationDetach) (*ApplicationDetachReply, error)
	ApplicationTLSCertAdd(ctx context.Context, in *ApplicationTLSCertAdd) (*ApplicationTLSCertAddReply, error)
	ApplicationTLSKeyAdd(ctx context.Context, in *ApplicationTLSKeyAdd) (*ApplicationTLSKeyAddReply, error)
	SessionEnableDisable(ctx context.Context, in *SessionEnableDisable) (*SessionEnableDisableReply, error)
	SessionRuleAddDel(ctx context.Context, in *SessionRuleAddDel) (*SessionRuleAddDelReply, error)
	SessionRulesDump(ctx context.Context, in *SessionRulesDump) (RPCService_SessionRulesDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) AppAddCertKeyPair(ctx context.Context, in *AppAddCertKeyPair) (*AppAddCertKeyPairReply, error) {
	out := new(AppAddCertKeyPairReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AppAttach(ctx context.Context, in *AppAttach) (*AppAttachReply, error) {
	out := new(AppAttachReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AppDelCertKeyPair(ctx context.Context, in *AppDelCertKeyPair) (*AppDelCertKeyPairReply, error) {
	out := new(AppDelCertKeyPairReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AppNamespaceAddDel(ctx context.Context, in *AppNamespaceAddDel) (*AppNamespaceAddDelReply, error) {
	out := new(AppNamespaceAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AppWorkerAddDel(ctx context.Context, in *AppWorkerAddDel) (*AppWorkerAddDelReply, error) {
	out := new(AppWorkerAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ApplicationDetach(ctx context.Context, in *ApplicationDetach) (*ApplicationDetachReply, error) {
	out := new(ApplicationDetachReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ApplicationTLSCertAdd(ctx context.Context, in *ApplicationTLSCertAdd) (*ApplicationTLSCertAddReply, error) {
	out := new(ApplicationTLSCertAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ApplicationTLSKeyAdd(ctx context.Context, in *ApplicationTLSKeyAdd) (*ApplicationTLSKeyAddReply, error) {
	out := new(ApplicationTLSKeyAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SessionEnableDisable(ctx context.Context, in *SessionEnableDisable) (*SessionEnableDisableReply, error) {
	out := new(SessionEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SessionRuleAddDel(ctx context.Context, in *SessionRuleAddDel) (*SessionRuleAddDelReply, error) {
	out := new(SessionRuleAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SessionRulesDump(ctx context.Context, in *SessionRulesDump) (RPCService_SessionRulesDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SessionRulesDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SessionRulesDumpClient interface {
	Recv() (*SessionRulesDetails, error)
	api.Stream
}

type serviceClient_SessionRulesDumpClient struct {
	api.Stream
}

func (c *serviceClient_SessionRulesDumpClient) Recv() (*SessionRulesDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SessionRulesDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
