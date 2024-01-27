package server

import (
	"net/http"
)

type HandlerBlock func(hw http.ResponseWriter, hr *http.Request)

var route = &Rout{
	registrations: map[string]HandlerBlock{},
	paths:         map[string]bool{},
}

type Rout struct {
	registrations map[string]HandlerBlock
	paths         map[string]bool
}

func (r *Rout) register(method string, path string, handler HandlerBlock) {
	identifier := buildRouteRegisterIdentifier(method, path)
	r.registrations[identifier] = handler
	r.paths[path] = true
}

func (r *Rout) getRouteHandler(method string, path string) HandlerBlock {
	identifier := buildRouteRegisterIdentifier(method, path)
	return r.registrations[identifier]
}

func (r *Rout) hasRouteForPath(path string) bool {
	return r.paths[path]
}
