package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	if err := runScheduler(); err != nil {
		glog.Fatal(err)
	}
}
