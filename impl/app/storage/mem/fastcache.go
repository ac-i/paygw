package mem

import (
	"github.com/VictoriaMetrics/fastcache"
)

type Storage struct {
	Conn *fastcache.Cache
}

func NewStorageMEM() *Storage {
	return &Storage{
		Conn: fastcache.New(1024),
	}
}
