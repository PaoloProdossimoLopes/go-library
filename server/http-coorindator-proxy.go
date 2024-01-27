package server

import (
	"fmt"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/logger"
)

func HttpCoordinatorProxy() HandlerBlock {
	return func(w http.ResponseWriter, r *http.Request) {
		requestMethod := r.Method
		requestPath := r.URL.Path

		if route.getRouteHandler(requestMethod, requestPath) == nil {
			errorMessage := fmt.Sprintf(
				"Route '%v' not found for method %v.",
				requestPath,
				requestMethod,
			)
			logger.Error(errorMessage)
			w.WriteHeader(http.StatusNotFound)
			w.Write(ResponseError{
				Error:      "Not Found",
				Reason:     errorMessage,
				StatusCode: http.StatusNotFound,
			}.Marshal())
			return
		}

		route.getRouteHandler(requestMethod, requestPath)(w, r)
		return
	}
}
