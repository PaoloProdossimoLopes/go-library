package server

import (
	"net/http"
)

func Put(path string, handler HandlerBlock) {
	httpRouteHandler(http.MethodOptions, path, handler)
}
