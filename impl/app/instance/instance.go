package instance

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"

	"paygw/impl/app/config"
	"paygw/impl/app/logs"
	"paygw/proto/app"

	"github.com/rs/zerolog"
)

type App struct {
	Cfg    *app.Config
	Ctx    context.Context
	Cancel context.CancelFunc
	ChErr  chan error
	ChMsg  chan string
	Log    zerolog.Logger
}

func NewApp() *App {
	app := new(App)

	app.Ctx, app.Cancel = context.WithCancel(context.Background())
	// a.ctx, a.cancel = context.WithTimeout(context.Background(), 5*time.Second)
	// defer app.Cancel()

	app.ChErr = make(chan error, 2)
	app.ChMsg = make(chan string, 8)
	app.Log = logs.LogC2
	go app.appInterrupt()

	app.Cfg = config.Config()
	jb, _ := json.Marshal(app.Cfg)

	app.Log.Info().RawJSON("config", jb).Send()

	return app
}

func (a *App) appInterrupt() {
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, os.Interrupt)
	for {
		select {
		case sig := <-chSig:
			if sig == os.Interrupt {
				a.Log.Info().Msgf("exit os.Interrupt: (%v)", sig)
				a.Cancel()
				return
			}
		case <-a.Ctx.Done():
			a.Log.Info().Msg("ctx done")
			return
		case err := <-a.ChErr:
			a.Log.Error().Err(err).Msg("chErr")
			return
		case msg := <-a.ChMsg:
			a.Log.Info().Msg(msg)
		}
	}
}
