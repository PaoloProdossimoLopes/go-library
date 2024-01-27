package server

import (
	"net/http"
)

func Delete(path string, handler HandlerBlock) {
	httpRouteHandler(http.MethodDelete, path, handler)
}
