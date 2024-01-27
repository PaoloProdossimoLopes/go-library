package server

import (
	"net/http"
)

func Post(path string, handler HandlerBlock) {
	httpRouteHandler(http.MethodPost, path, handler)
}
