// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package ioam_cache

import (
	"context"
	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service  ioam_cache.
type RPCService interface {
	IoamCacheIP6EnableDisable(ctx context.Context, in *IoamCacheIP6EnableDisable) (*IoamCacheIP6EnableDisableReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) IoamCacheIP6EnableDisable(ctx context.Context, in *IoamCacheIP6EnableDisable) (*IoamCacheIP6EnableDisableReply, error) {
	out := new(IoamCacheIP6EnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
