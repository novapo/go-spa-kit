package plugins

import "github.com/gorilla/mux"

var registeredPlugins []Plugin

// Plugin is the basic interface for all plugins
type Plugin interface {
	RegisterRoutes(r *mux.Router)
	GetName() string
}

// Register a plugin
func Register(p Plugin) {
	registeredPlugins = append(registeredPlugins, p)
}

// AddPluginRoutes will add all routes from the plugins to the given router
func AddPluginRoutes(r *mux.Router) {
	pluginRouter := r.PathPrefix("/plugins").Subrouter()
	for _, p := range registeredPlugins {
		r = pluginRouter.PathPrefix("/" + p.GetName()).Subrouter()
		p.RegisterRoutes(r)
	}
}
