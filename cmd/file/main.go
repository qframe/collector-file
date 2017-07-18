package main

import (
	"log"
	"github.com/zpatrick/go-config"
	"github.com/qnib/qframe-types"
	"github.com/qframe/collector-file"
	"os"
)

func main() {
	qChan := qtypes.NewQChan()
	qChan.Broadcast()
	if len(os.Args) != 2 {
		log.Fatal("usage: ./file <path>")

	}
	fPath := os.Args[1]
	cfgMap := map[string]string{
		"collector.file.path": fPath,
	}
	cfg := config.NewConfig([]config.Provider{config.NewStatic(cfgMap)})

	p, err := collector_file.New(qChan, cfg, "file")
	if err != nil {
		log.Fatalf("[EE] Failed to create collector: %v", err)
	}
	go p.Run()
	bg := p.QChan.Data.Join()
	for {
		val := <- bg.Read
		log.Printf("%v", val)
	}
}
