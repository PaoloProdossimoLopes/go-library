package server

import (
	"net/http"
)

func Options(path string, handler HandlerBlock) {
	httpRouteHandler(http.MethodOptions, path, handler)
}
