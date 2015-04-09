package main

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/novapo/go-spa-kit/config"
	"github.com/novapo/go-spa-kit/db"
)

func main() {
	go start()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	<-sigc
}

func start() {
	conf := config.GetServerConfig()
	db.Init()
	http.ListenAndServe(":"+strconv.Itoa(conf.Port), createRouter())
}
