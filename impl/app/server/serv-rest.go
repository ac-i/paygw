package server

import (
	"fmt"
	"paygw/impl/app/instance"
)

type ServREST struct {
	app    *instance.App
	active int32
}

func (s *Server) setServREST() (*ServREST, error) {
	if s.servREST != nil {
		return s.servREST, nil
	}
	if s == nil || s.cfg == nil || s.cfg.ServRestgw == nil {
		return &ServREST{}, fmt.Errorf("invalid ServREST config")
	}
	s.servREST = &ServREST{
		app:    s.app,
		active: s.cfg.ServRest.Active,
	}

	return s.servREST, nil
}
