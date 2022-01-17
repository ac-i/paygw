package server

import (
	"context"
	"fmt"

	"paygw/impl/app/instance"
	"paygw/impl/paycard"
	"paygw/proto/app"
)

type Server struct {
	active     int32
	app        *instance.App
	cfg        *app.Server
	servGRPC   *ServGRPC
	servRESTGW *ServRESTGW
	servREST   *ServREST
	cliGRPC    *CliGRPC
	//
	paycardService *paycard.PayCardService
}

func NewServer(app *instance.App) *Server {
	if app == nil {
		app.ChErr <- fmt.Errorf("invalid app")
		return nil
	}

	s := &Server{
		app:            app,
		cfg:            app.Cfg.Server,
		active:         app.Cfg.Server.Active,
		paycardService: paycard.NewPayCardService(app),
	}

	var err error
	s.servGRPC, err = s.setServGRPC()
	if err != nil {
		app.ChErr <- err
		return nil
	}
	s.servRESTGW, err = s.setServRESTGW()
	if err != nil {
		app.ChErr <- err
		return nil
	}
	s.servREST, err = s.setServREST()
	if err != nil {
		app.ChErr <- err
		return nil
	}
	if err = s.setCliGRPC(); err != nil {
		if err != nil {
			app.ChErr <- fmt.Errorf("invalid newCliGRPC, err: %v", err)
			return nil
		}
	}
	return s
}

func (s *Server) Serve() {
	if s.servGRPC.active > 0 {
		go s.serveGRCP()
	}
	if s.servRESTGW.active > 0 {
		go s.serveRESTGW()
	}

	<-s.app.Ctx.Done()
	s.Close()
}

func (s *Server) Close() {
	if s.servRESTGW.active > 0 {
		s.closeRESTGW()
	}
	if s.cliGRPC.active > 0 {
		s.cliGRPC.closeGRCP()
	}
	if s.servGRPC.active > 0 {
		s.closeGRCP()
	}

	s.app.Log.Info().Msg("Server closed")
	s.app.Cancel()
}

func (s *Server) LoadTests(ctx context.Context, fname string, storageCase string) {
	// s.LoadTestData(ctx, fname, storageCase)
}
