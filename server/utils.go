package server

import (
	"fmt"
	"net/http"
)

func buildRouteRegisterIdentifier(httpMethod, pathRoute string) string {
	return fmt.Sprintf("%v:%v", httpMethod, pathRoute)
}

func httpRouteHandler(httpMethod, path string, handler HandlerBlock) {
	if route.getRouteHandler(httpMethod, path) != nil {
		routeAlreadyRegisteredError := fmt.Errorf("Route '%v' already registered for method %v.", path, httpMethod)
		panic(routeAlreadyRegisteredError)
	}

	if !route.hasRouteForPath(path) {
		http.HandleFunc(path, HttpCoordinatorProxy())
	}

	route.register(httpMethod, path, handler)
}
