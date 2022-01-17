package config

import (
	"paygw/proto/app"
)

var (
	def_AQL_HOSTS    = "localhost 127.0.0.1"
	def_AQL_USERNAME = "arango"
	def_AQL_PASSWORD = "arango"
	def_CQL_HOSTS    = "localhost 127.0.0.1"
	def_CQL_USERNAME = "scylla"
	def_CQL_PASSWORD = "scylla"
)

func Storage(e *app.Environment, active int32) *app.Storage {
	return &app.Storage{
		Active:  activeServ(e, active, serv_grpc),
		Drivers: storageDrivers(e, active),
	}
}

func StorageMEM(active int32) *app.StorageDriver {
	return &app.StorageDriver{
		Active: active,
		Driver: Driver_mem,
	}
}

func StorageGRPC(active int32) *app.StorageDriver {
	return &app.StorageDriver{
		Active: active,
		Driver: Driver_grpc,
	}
}

func StorageAQL(active int32) *app.StorageDriver {
	return &app.StorageDriver{
		Active: active,
		Driver: Driver_aql,
		Hosts:  es("AQL_HOSTS", def_AQL_HOSTS),
		Auth: &app.BasicAuth{
			Username: en("AQL_USERNAME", def_AQL_USERNAME),
			Password: en("AQL_PASSWORD", def_AQL_PASSWORD),
		},
	}
}

func StorageCQL(active int32) *app.StorageDriver {
	return &app.StorageDriver{
		Active: active,
		Driver: Driver_cql,
		Hosts:  es("CQL_HOSTS", def_CQL_HOSTS),
		Auth: &app.BasicAuth{
			Username: en("CQL_USERNAME", def_CQL_USERNAME),
			Password: en("CQL_PASSWORD", def_CQL_PASSWORD),
		},
	}
}

func StorageUNSET(active int32) *app.StorageDriver {
	return &app.StorageDriver{
		Active: -1,
		Driver: Driver_unset,
	}
}
