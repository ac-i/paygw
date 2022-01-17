package server

import (
	"fmt"
	"net/http"
	"time"

	"paygw/impl/app/instance"
	"paygw/proto/paycard_srv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServRESTGW struct {
	app        *instance.App
	active     int32
	httpServer *http.Server
}

func (s *Server) setServRESTGW() (*ServRESTGW, error) {
	if s.servRESTGW != nil {
		return s.servRESTGW, nil
	}
	if s == nil || s.cfg == nil || s.cfg.ServRestgw == nil {
		return nil, fmt.Errorf("invalid ServRESTGW config")
	}
	s.servRESTGW = &ServRESTGW{
		app:    s.app,
		active: s.cfg.ServRestgw.Active,
	}

	var err error
	gwMux := runtime.NewServeMux()
	// err = gwMux.HandlePath("POST", "/api/file", s.servRESTGW...HandleBinaryFileUploadPortMaps)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = paycard_srv.RegisterPayCardServiceHandlerFromEndpoint(s.servRESTGW.app.Ctx, gwMux, s.cfg.ServRestgw.GrpcEndpoint, opts)
	if err != nil {
		s.app.ChErr <- err
		return nil, err
	}

	s.servRESTGW.httpServer = &http.Server{
		Addr:              s.cfg.ServRestgw.RestgwAddr,
		Handler:           gwMux,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       10 * time.Minute, // synch max ctx timeout
		ReadTimeout:       10 * time.Second,
	}

	return s.servRESTGW, nil
}

func (s *Server) serveRESTGW() {
	if s.servRESTGW.active <= 0 {
		return
	}
	// Note: Make sure the gRPC server is running properly and accessible
	if !s.cliGRPC.connOK() {
		s.app.ChErr <- fmt.Errorf("invalid ServRESTGW cliGRPC connection")
		return
	}

	var err error
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	// s.app.ChMsg <- fmt.Sprintf("starting insecure HTTP REST GATEWAY (no TLS & no auth) on %s (gw to grpc: %s)", s.cfg.ServRestgw.RestgwAddr, s.cfg.ServRestgw.GrpcEndpoint)
	s.app.Log.Info().
		Str("addr-gw", s.cfg.ServRestgw.RestgwAddr).Str("addr-grpc", s.cfg.ServRestgw.GrpcEndpoint).
		Msg("starting insecure HTTP REST GATEWAY (no TLS & no auth)")
	err = s.servRESTGW.httpServer.ListenAndServe()
	if err != nil {
		s.app.ChErr <- err
		return
	}
	// return http.ListenAndServe("srv.GetRest().GetAddress()", gwMux)
	// After Shutdown or Close, the returned error is ErrServerClosed.
}

func (s *Server) closeRESTGW() {
	if s.servRESTGW.httpServer == nil {
		return
	}
	_ = s.servRESTGW.httpServer.Shutdown(s.app.Ctx)
}
