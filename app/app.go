package app

import (
	"context"

	"paygw/impl/app/instance"
	"paygw/impl/app/server"
)

func Main() {
	// os.Setenv("PAYGW_SERVERS", "grpc restgw -rest")
	// os.Setenv("PAYGW_STORAGES", "mem grpc aql cql")

	app := instance.NewApp()
	defer app.Cancel()

	server.NewServer(app).Serve()

	<-app.Ctx.Done()
	app.Log.Info().Msg("app exit done")
}

func LoadTests(fname string, storageCase string) {

	app := instance.NewApp()
	defer app.Cancel()

	serv := server.NewServer(app)
	go serv.Serve()
	// const timeout = 10 * time.Second
	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	serv.LoadTests(ctx, fname, storageCase)
	defer serv.Close()

	// <-app.Ctx.Done()
	app.Log.Info().Msg("app  LoadTests exit done")

}
