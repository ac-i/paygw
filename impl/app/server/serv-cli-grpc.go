package server

import (
	"context"
	"fmt"
	"paygw/impl/app/instance"
	"paygw/proto/app"
	"paygw/proto/paycard"
	"paygw/proto/paycard_srv"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CliGRPC struct {
	app        *instance.App
	cfg        *app.CliGRPC
	active     int32
	mu         sync.Mutex
	targetAddr string
	connGRPC   *grpc.ClientConn
	connCancel context.CancelFunc
}

func (s *Server) setCliGRPC() error {
	if s.cliGRPC != nil {
		return nil
	}
	if s == nil || s.cfg == nil || s.cfg.GetCliGrpc() == nil {
		return fmt.Errorf("invalid Client gRPC config")
	}
	s.cliGRPC = &CliGRPC{
		app:    s.app,
		cfg:    s.cfg.CliGrpc,
		active: s.cfg.CliGrpc.Active,
	}
	return nil
}

func (c *CliGRPC) connOK() bool {
	if c.connGRPC != nil {
		return true
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()), // grpc.WithInsecure(),
	}

	var (
		err     error
		ctxDial context.Context
	)
	ctxDial, c.connCancel = context.WithTimeout(c.app.Ctx, 10*time.Second)

	c.connGRPC, err = grpc.DialContext(ctxDial, c.cfg.Target, opts...)
	if err != nil {
		c.app.Log.Error().Err(err).Str("target", c.cfg.Target).Msg("unsuccessful CliGRPC diall")
		c.app.ChErr <- err
		return false
	}

	c.targetAddr = c.cfg.Target

	return true
}

func (c *CliGRPC) closeGRCP() {
	if c.connCancel != nil {
		c.connCancel()
	}
	if c.connGRPC != nil {
		c.connGRPC.Close()
	}
}

func (c *CliGRPC) connRPC(ctx context.Context) (rpcCancel context.CancelFunc, ctxRPC context.Context, ok bool) {
	if !c.connOK() {
		return nil, nil, false
	}
	deadline, ok := ctx.Deadline()
	if ok {
		ctxRPC, rpcCancel = context.WithDeadline(c.app.Ctx, deadline)
	} else {
		ctxRPC, rpcCancel = context.WithTimeout(c.app.Ctx, 10*time.Minute)
	}
	// opts = []grpc.CallOption{
	paycard_srv.NewPayCardServiceClient(c.connGRPC).Process(ctxRPC, &paycard.Transaction{})
	return rpcCancel, ctxRPC, true
}

func (c *CliGRPC) PayCardAuthorize(ctx context.Context, in *paycard.Authorization) (*paycard.AuthorizationState, error) {
	rpcCancel, ctxRPC, ok := c.connRPC(ctx)
	if !ok {
		return nil, fmt.Errorf("no valid CliGRPC RPC options")
	}
	defer rpcCancel()

	res, err := paycard_srv.NewPayCardServiceClient(c.connGRPC).Authorize(ctxRPC, in) //in &paycard.Authorization{}
	if err != nil {
		return nil, fmt.Errorf("no valid CliGRPC RPC options")
	}

	return res, nil
}

func (c *CliGRPC) PayCardProcess(ctx context.Context, in *paycard.Transaction) (*paycard.TransactionState, error) {
	rpcCancel, ctxRPC, ok := c.connRPC(ctx)
	if !ok {
		return nil, fmt.Errorf("no valid CliGRPC RPC options")
	}
	defer rpcCancel()

	res, err := paycard_srv.NewPayCardServiceClient(c.connGRPC).Process(ctxRPC, in) // in &paycard.Transaction{}
	if err != nil {
		return nil, fmt.Errorf("no valid CliGRPC RPC options")
	}

	return res, nil
}
