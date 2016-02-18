package main

import (
	"flag"
	"fmt"
	"github.com/niean/opsadmin/g"
	"github.com/niean/opsadmin/http"
	"github.com/niean/opsadmin/sender"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := flag.String("c", "cfg.json", "config file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.LoadConfig(*cfg)
	sender.Start()

	http.Start()

	wait_signal()
}

func wait_signal() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-sc:
		os.Exit(0)
	}
}
