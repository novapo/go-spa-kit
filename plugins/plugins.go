package plugins

import "github.com/gorilla/mux"

var registeredPlugins []Plugin

type Plugin interface {
	RegisterRoutes(r *mux.Router)
	GetName() string
}

func Register(p Plugin) {
	registeredPlugins = append(registeredPlugins, p)
}

func AddPluginRoutes(r *mux.Router) {
	pluginRouter := r.PathPrefix("/plugins").Subrouter()
	for _, p := range registeredPlugins {
		r = pluginRouter.PathPrefix("/" + p.GetName()).Subrouter()
		p.RegisterRoutes(r)
	}
}
