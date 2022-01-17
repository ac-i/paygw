package server

import (
	"fmt"
	"net"

	"paygw/impl/app/instance"
	"paygw/proto/paycard_srv"

	"google.golang.org/grpc"
)

type ServGRPC struct {
	active     int32
	app        *instance.App
	netLis     net.Listener
	grpcServer *grpc.Server
}

func (s *Server) setServGRPC() (*ServGRPC, error) {
	if s.servGRPC != nil {
		return s.servGRPC, nil
	}
	if s == nil || s.cfg == nil || s.cfg.ServGrpc == nil {
		return nil, fmt.Errorf("invalid ServGrpc config")
	}
	s.servGRPC = &ServGRPC{
		app:    s.app,
		active: s.cfg.ServGrpc.Active,
	}

	// create a gRPC server
	// insecure,  no []grpc.ServerOption for TLS and Auth
	s.servGRPC.grpcServer = grpc.NewServer()
	paycard_srv.RegisterPayCardServiceServer(s.servGRPC.grpcServer, s.paycardService)

	return s.servGRPC, nil
}

func (s *Server) serveGRCP() {
	if s.servGRPC.active <= 0 {
		return
	}
	var err error
	// create a listener on TCP port (def 8090)
	s.servGRPC.netLis, err = net.Listen(s.cfg.ServGrpc.Network, s.cfg.ServGrpc.Address)
	if err != nil {
		s.app.ChErr <- err
		return
	}

	// create a gRPC server ^^^
	// start the server
	// s.app.ChMsg <- fmt.Sprintf("starting insecure gRCP server (no TLS & no auth) on %+v", s.servGRPC.netLis.Addr())
	s.app.Log.Info().Str("addr", fmt.Sprintf("%+v", s.servGRPC.netLis.Addr())).Msg("starting insecure gRCP server (no TLS & no auth)")
	err = s.servGRPC.grpcServer.Serve(s.servGRPC.netLis)
	if err != nil {
		s.app.Log.Error().Err(err).Send()
		s.app.ChErr <- err
		return
	}
	// return grpcServer.Serve(netLis)
	// Serve will return a non-nil error unless Stop or GracefulStop is called.

}

func (s *Server) closeGRCP() {
	if s.servGRPC.grpcServer == nil {
		return
	}
	s.servGRPC.grpcServer.GracefulStop()
}
