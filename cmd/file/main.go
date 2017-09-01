package main

import (
	"log"
	"github.com/zpatrick/go-config"
	"github.com/qframe/collector-file"
	"os"
	"github.com/qframe/types/qchannel"
)

func main() {
	qChan := qtypes_qchannel.NewQChan()
	qChan.Broadcast()
	if len(os.Args) != 2 {
		log.Fatal("usage: ./file <path>")

	}
	fPath := os.Args[1]
	cfgMap := map[string]string{
		"collector.file.path": fPath,
		"collector.file.reopen": "false",
	}
	cfg := config.NewConfig([]config.Provider{config.NewStatic(cfgMap)})

	p, err := qcollector_file.New(qChan, cfg, "file")
	if err != nil {
		log.Fatalf("[EE] Failed to create collector: %v", err)
	}
	go p.Run()
	bg,_,_ := p.JoinChannels()
	for {
		val := <- bg.Read
		log.Printf("%v", val)
	}
}
