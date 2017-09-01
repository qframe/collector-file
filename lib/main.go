package main

import (
	"github.com/zpatrick/go-config"
	"github.com/qnib/qframe-types"
	"github.com/qframe/collector-file"

)

func Start(qChan qtypes.QChan, cfg *config.Config, name string) {
	p, _ := qcollector_file.New(qChan, cfg, name)
	p.Run()
}
