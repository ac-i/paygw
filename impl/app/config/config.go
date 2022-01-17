package config

import (
	"os"
	"strings"
	"time"

	"paygw/proto/app"
)

const (
	AppName        = "paygw"
	appEnvModeFld  = "PAYGW_MODE"
	appServersFld  = "PAYGW_SERVERS"
	appStoragesFld = "PAYGW_STORAGES"
	Driver_mem     = "mem"
	Driver_grpc    = "grpc"
	Driver_aql     = "aql"
	Driver_cql     = "cql"
	Driver_unset   = "unset"
	serv_grpc      = "grpc"
	serv_restgw    = "restgw"
	serv_rest      = "rest"
	Cli_grpc       = "grpc"
)

var (
	appEnvModeValDef  = app.AppMode(app.AppMode_APP_MODE_DEV).String()
	appServersValDef  = strings.Join([]string{serv_grpc, serv_restgw}, " ")
	appStoragesValDef = strings.Join([]string{Driver_mem}, " ")
)

func Config() *app.Config {
	env := Environment()
	return &app.Config{
		Environment: env,
		Active:      env.Active,
		Server:      Server(env, 1),
	}
}

func Environment() *app.Environment {
	appMod, ok := app.AppMode_value[en(appEnvModeFld, appEnvModeValDef)]
	var ac int32
	if ok && appMod > 0 {
		ac = 1
	}
	h, _ := os.Hostname()
	return &app.Environment{
		Active:   ac,
		AppName:  AppName,
		AppMode:  app.AppMode(appMod),
		AppHost:  h,
		AppPid:   int32(os.Getpid()),
		AppStart: time.Now().UnixNano(),
		Servers:  es(appServersFld, appServersValDef),
		Storages: es(appStoragesFld, appStoragesValDef),
	}
}

func en(env, def string) (s string) {
	s = os.Getenv(env)
	if s != "" {
		return
	}
	return def
}

func es(env, def string) (s []string) {
	if os.Getenv(env) != "" {
		return strings.Split(os.Getenv(env), " ")
	}
	return strings.Split(def, " ")
}

func activeServ(e *app.Environment, active int32, serv string) int32 {
	if e.Active <= 0 || active <= 0 {
		return -1
	}
	for _, s := range e.GetServers() {
		if serv == s {
			return active
		}
	}
	return -1
}

func storageDrivers(e *app.Environment, active int32) map[string]*app.StorageDriver {
	grpcActive := activeServ(e, active, serv_grpc)
	drivers := make(map[string]*app.StorageDriver, len(e.Storages))
	for _, dr := range e.Storages {
		if dr == Driver_mem {
			drivers[dr] = StorageMEM(grpcActive)
		} else if dr == Driver_grpc {
			drivers[dr] = StorageGRPC(grpcActive)
		} else if dr == Driver_aql {
			drivers[dr] = StorageAQL(grpcActive)
		} else if dr == Driver_cql {
			drivers[dr] = StorageCQL(grpcActive)
		} else {
			drivers[Driver_unset] = StorageUNSET(-1)
		}
	}
	return drivers
}
