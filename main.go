package main

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/novapo/go-spa-kit/config"
)

func main() {
	conf := config.GetServerConfig()
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir(conf.Path)))
	r.HandleFunc("/admin", AdminHandler).Methods("GET")
	http.ListenAndServe(":"+strconv.Itoa(conf.Port), r)
}

func AdminHandler(w http.ResponseWriter, req *http.Request) {
	adminPage, err := ioutil.ReadFile("admin/admin.html")
	username, password, ok := req.BasicAuth()
	if !ok || username != config.AdnConf.Login || password != config.AdnConf.Password {
		w.Header().Add("WWW-Authenticate", "Basic realm='Login alter'")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write(adminPage)
}
