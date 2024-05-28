package web

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed all:dist
var resource embed.FS

func SPAAssets() (spa fs.FS) {
	var err error
	if spa, err = fs.Sub(resource, "dist"); err != nil {
		log.Fatalln("SPA_FS_ERR:", err.Error())
	}
	return spa
}
