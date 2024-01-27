package middleware

import (
	"fmt"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/logger"
	"github.com/PaoloProdossimoLopes/go-library/server"
)

type HandlerBlock func(hw http.ResponseWriter, hr *http.Request)

func HttpCoordinatorMethodMiddleware() HandlerBlock {
	return func(w http.ResponseWriter, r *http.Request) {
		requestMethod := r.Method
		requestPath := r.URL.Path

		routeIdentifier := buildRouteRegisterIdentifier(requestMethod, requestPath)

		if registeredRoutes[routeIdentifier] == nil {
			errorMessage := fmt.Sprintf(
				"Route '%v' not found for method %v.",
				requestPath,
				requestMethod,
			)
			logger.Error(errorMessage)
			w.WriteHeader(http.StatusNotFound)
			w.Write(server.ResponseError{
				Error:      "Not Found",
				Reason:     errorMessage,
				StatusCode: http.StatusNotFound,
			}.Marshal())
			return
		}

		registeredRoutes[routeIdentifier](w, r)
		return
	}
}
