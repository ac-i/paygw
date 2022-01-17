// Package logs provides a global logger for zerolog.
// Enhancement of github.com/rs/zerolog/log of Olivier Poitrey
package logs

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	zerolog.TimestampFieldName = _time_µs
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.CallerMarshalFunc = formatCaller
	cfg := &Logs{
		App: "paygw",
		Enable: EnableOutput{
			StdOut:  false,
			Console: true,
		},
	}
	cfg.SetOnce()
}

var (
	Logger   zerolog.Logger
	LogC2    zerolog.Logger // LogC2 global logger with caller in place - CallerWithSkipFrameCount(2)
	LogC3    zerolog.Logger // LogC3 global logger with caller one before - CallerWithSkipFrameCount(3)
	setOK    bool
	basepath string
)

const (
	_time_µs = "µs"
)

// SetOnce() passes cfg at runtime
func (c *Logs) SetOnce() {
	if setOK || c.App == "" {
		return
	}
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	basepath, _ = os.Getwd()
	Logger = zerolog.New(zerolog.MultiLevelWriter(
		c.setSyslog(),
		c.setStdErr(),
		c.setStdOut(),
		c.setConsole(),
	)).
		With().
		Timestamp().
		Str("app", c.App).
		Int("pid", os.Getpid()).
		Logger()

	LogC2 = Logger.With().
		CallerWithSkipFrameCount(2).
		// Int("pkg-logs", 1).
		Logger()

	LogC3 = Logger.With().
		CallerWithSkipFrameCount(3).
		// Int("pkg-logs", 1).
		Logger()

	log.SetFlags(0)
	log.SetOutput(Logger.With().
		CallerWithSkipFrameCount(4).
		// Int("pkg-log", 1).
		Str("deprecated-log", "log").
		Logger())

	setOK = true
}

func formatCaller(file string, line int) string {
	if len(file) > 0 && len(basepath) > 0 {
		if rel, err := filepath.Rel(basepath, file); err == nil {
			file = filepath.Clean(rel)
			if strings.Contains(file, "src/") {
				file = file[strings.Index(file, "src/")+len("src/"):]
			}
			if strings.Contains(file, "../") {
				file = file[strings.Index(file, "../")+len("../"):]
			}
		}
	}
	return file + ":" + strconv.Itoa(line)
}
