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
	r.HandleFunc("/admin", AdminHandler)
	http.ListenAndServe(":"+strconv.Itoa(conf.Port), r)
}

func AdminHandler(w http.ResponseWriter, req *http.Request) {
	adminPage, err := ioutil.ReadFile("admin/admin.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write(adminPage)
}
