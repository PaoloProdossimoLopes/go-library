package server

import (
	"net/http"
)

func Patch(path string, handler HandlerBlock) {
	httpRouteHandler(http.MethodPatch, path, handler)
}
