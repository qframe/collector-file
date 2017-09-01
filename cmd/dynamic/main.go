package main


import (
	"log"
	"os"
	"plugin"
	"github.com/qnib/qframe-types"
	"github.com/zpatrick/go-config"
)


func main() {
	qChan := qtypes.NewQChan()
	qChan.Broadcast()
	cfgMap := map[string]string{"collector.file.path": os.Args[1]}
	cfg := config.NewConfig([]config.Provider{config.NewStatic(cfgMap)})
	plug, _ := plugin.Open("./file.so")
	start, _ := plug.Lookup("Start")
	go start.(func(qtypes.QChan, *config.Config, string))(qChan, cfg, "file")
	bg := qChan.Data.Join()
	for {
		val := <- bg.Read
		log.Printf("%v", val)
	}
}