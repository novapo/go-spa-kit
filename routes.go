package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/novapo/go-spa-kit/admin"
	"github.com/novapo/go-spa-kit/config"
	"github.com/novapo/go-spa-kit/plugins"
)

func createRouter() *mux.Router {
	conf := config.GetServerConfig()
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir(conf.Path)))
	r.HandleFunc("/admin", admin.AdminHandler).Methods("GET")
	r.HandleFunc("/login", admin.LoginHandler)
	plugins.AddPluginRoutes(r)
	return r
}
