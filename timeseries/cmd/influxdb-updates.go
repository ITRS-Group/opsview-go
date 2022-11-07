package main

import (
	"flag"
	"github.com/ITRS-Group/opsview-go/timeseries"
)

func main() {
	conf_dir := flag.String("c", "./etc", "default configuration directory")
	flag.Parse()

	server := &timeseries.TimeseriesServer{}
	server.ReadConfig(*conf_dir)
	server.Launch("updates")
}
