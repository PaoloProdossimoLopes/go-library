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

		fmt.Println("|->", requestMethod, requestPath)

		if route.getRouteHandler(requestMethod, requestPath) == nil {
			errorMessage := fmt.Sprintf(
				"Route '%v' not found for method %v.",
				requestPath,
				requestMethod,
			)
			logger.Error(errorMessage)
			SendErrorResponse(w, ResponseError{
				Error:      "Not Found",
				Reason:     errorMessage,
				StatusCode: http.StatusNotFound,
			})
			return
		}

		route.getRouteHandler(requestMethod, requestPath)(w, r)
		return
	}
}
