package logs

import (
	"io"
	"log/syslog"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logs configuration & credentials of logs // `yaml:"logs" json:"logs"`
type Logs struct {
	App    string `yaml:"app" json:"app"`
	Enable EnableOutput
}

type EnableOutput struct {
	SysLog  bool `yaml:"syslog" json:"syslog"`
	StdErr  bool `yaml:"stderr" json:"stderr"`
	StdOut  bool `yaml:"stdout" json:"stdout"`
	Console bool `yaml:"console" json:"console"`
	None    bool `yaml:"none" json:"none"`
}

type Splunk struct {
	Hosts []string `yaml:"hosts" json:"hosts"`
	Token string   `yaml:"token" json:"token"`
}

// explicit confirmation to not log on any writer
// (in pair with None)
func (c *Logs) anyEnabled() bool {
	return c.Enable.SysLog ||
		c.Enable.StdErr ||
		c.Enable.StdOut ||
		c.Enable.Console
}

func (c *Logs) setSyslog() io.Writer {
	if !c.Enable.SysLog {
		return zerolog.Nop()
	}
	syslogWriter, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL6, c.App)
	if err != nil {
		Error().Err(err).Msg("Unable to setup syslog")
		return zerolog.Nop()
	}
	return zerolog.SyslogCEEWriter(syslogWriter)
}

func (c *Logs) setStdErr() io.Writer {
	if !c.Enable.StdErr {
		if c.anyEnabled() ||
			(!c.anyEnabled() && c.Enable.None) {
			return zerolog.Nop()
		}
	}
	return os.Stderr
}

func (c *Logs) setStdOut() io.Writer {
	if !c.Enable.StdOut {
		return zerolog.Nop()
	}
	return os.Stdout
}

func (c *Logs) setConsole() io.Writer {
	if !c.Enable.Console {
		return zerolog.Nop()
	}
	return zerolog.ConsoleWriter{Out: os.Stdout,
		// StampMicro = "Jan _2 15:04:05.000000"
		TimeFormat: time.StampMicro,
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			"app",
			"pid",
			zerolog.LevelFieldName,
			zerolog.CallerFieldName,
			zerolog.MessageFieldName,
		},
		PartsExclude: []string{
			"host",
		}}
}
