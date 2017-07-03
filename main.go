package main

import (
	"log"
	"github.com/zpatrick/go-config"
	"github.com/qnib/qframe-types"
	"github.com/qframe/collector-file/lib"
)


func Run(qChan qtypes.QChan, cfg *config.Config, name string) {
	p, _ := collector_file.New(qChan, cfg, name)
	p.Run()
}


func main() {
	qChan := qtypes.NewQChan()
	qChan.Broadcast()
	cfgMap := map[string]string{
		"collector.file.path": "./resources/test.file",
	}

	cfg := config.NewConfig(
		[]config.Provider{
			config.NewStatic(cfgMap),
		},
	)

	p, err := collector_file.New(qChan, cfg, "file")
	if err != nil {
		log.Printf("[EE] Failed to create collector: %v", err)
		return
	}
	go p.Run()
	bg := p.QChan.Data.Join()
	for {
		select {
		case val := <- bg.Read:
			log.Printf("%v", val)
		}
	}
}
