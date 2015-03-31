package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/novapo/go-spa-kit/config"
)

func main() {
	conf := config.GetServerConfig()
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(conf.Port), http.FileServer(http.Dir(conf.Path))))
}
