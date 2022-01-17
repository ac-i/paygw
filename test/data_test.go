package data_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"paygw/impl/app/config"
	"paygw/impl/app/instance"
	"paygw/impl/app/server"
)

var tests = []struct {
	name        string
	fname       string
	storageCase string
}{
	{name: "[0-1] cliGRPC",
		fname:       "data-1.json",
		storageCase: config.Cli_grpc},
	{name: "[0-1] drvMEM",
		fname:       "data-1.json",
		storageCase: config.Driver_mem},
	// {name: "[1-1] cliGRPC",
	// 	fname:       "data-1.json",
	// 	storageCase: config.Cli_grpc},
	// {name: "[2-1] drvMEM",
	// 	fname:       "data-1.json",
	// 	storageCase: config.Driver_mem},
	// {name: "[1-10] cliGRPC",
	// 	fname:       "data-10.json",
	// 	storageCase: config.Cli_grpc},
	// {name: "[2-10] drvMEM",
	// 	fname:       "data-10.json",
	// 	storageCase: config.Driver_mem},
	// {name: "[1-100] cliGRPC",
	// 	fname:       "data-100.json",
	// 	storageCase: config.Cli_grpc},
	// {name: "[2-100] drvMEM",
	// 	fname:       "data-100.json",
	// 	storageCase: config.Driver_mem},
}

func TestServer_LoadTests(t *testing.T) {
	fmt.Println("TestServ_LoadTests")
	app := instance.NewApp()
	defer app.Cancel()

	srv := server.NewServer(app)
	go srv.Serve()
	defer srv.Close()

	// const timeout = 10 * time.Second
	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
	time.Sleep(1 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv.LoadTests(ctx, tt.fname, tt.storageCase)
		})
	}
}

func BenchmarkServer_LoadTests(b *testing.B) {
	fmt.Println("BenchmarkServ_LoadTests")
	app := instance.NewApp()
	defer app.Cancel()

	srv := server.NewServer(app)
	go srv.Serve()

	// const timeout = 10 * time.Second
	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
	time.Sleep(1 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer srv.Close()

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				srv.LoadTests(ctx, tt.fname, tt.storageCase)
			}
		})
	}
}
