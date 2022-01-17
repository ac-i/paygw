// go:build !slim

package data

import (
	"embed"
)

//go:embed *.json
var Files embed.FS
