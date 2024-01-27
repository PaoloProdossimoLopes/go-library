package server

import (
	"net/http"
)

func Get(path string, handler HandlerBlock) {
	httpRouteHandler(http.MethodGet, path, handler)
}
