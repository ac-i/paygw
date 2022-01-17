package config

import (
	"paygw/proto/app"
)

var (
	def_GRPC_NETWORK   = "tcp"
	def_GRPC_ADDRESS   = ":8090"
	def_RESTGW_ADDRESS = ":8080"
	def_REST_ADDRESS   = ":8070"
)

func Server(e *app.Environment, active int32) *app.Server {
	return &app.Server{
		Active: active,

		ServGrpc:   ServGRPC(e, activeServ(e, active, serv_grpc)),
		ServRestgw: ServRESTGW(e, activeServ(e, active, serv_restgw)),
		ServRest:   ServREST(e, activeServ(e, active, serv_rest)), // unimplemented

		CliGrpc: CliGRPC(e, active),
	}
}

func ServGRPC(e *app.Environment, active int32) *app.ServGRPC {
	act := activeServ(e, active, serv_grpc)
	return &app.ServGRPC{
		Active:  act,
		Driver:  serv_grpc,
		Network: en("GRPC_NETWORK", def_GRPC_NETWORK),
		Address: en("GRPC_ADDRESS", def_GRPC_ADDRESS),
		Storage: Storage(e, act),
	}
}

func ServRESTGW(e *app.Environment, active int32) *app.ServRESTGW {
	return &app.ServRESTGW{
		Active:       activeServ(e, active, serv_restgw),
		Driver:       serv_restgw,
		GrpcEndpoint: en("GRPC_ADDRESS", def_GRPC_ADDRESS),
		RestgwAddr:   en("RESTGW_ADDRESS", def_RESTGW_ADDRESS),
	}
}

func ServREST(e *app.Environment, active int32) *app.ServREST {
	return &app.ServREST{
		Active:   active,
		Driver:   serv_rest,
		RestAddr: en("REST_ADDRESS", def_REST_ADDRESS),
	}
}

func CliGRPC(e *app.Environment, active int32) *app.CliGRPC {
	return &app.CliGRPC{
		Active: active,
		Driver: Cli_grpc,
		Target: en("GRPC_ADDRESS", def_GRPC_ADDRESS),
	}
}
