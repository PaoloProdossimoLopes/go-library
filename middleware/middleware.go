package middleware

import (
	"fmt"
	"net/http"
)

var registeredRoutes = map[string]HandlerBlock{}

func Get(path string, handler HandlerBlock) {
	const httpMethod = http.MethodGet
	routeSpec := buildRouteRegisterIdentifier(httpMethod, path)

	if registeredRoutes[routeSpec] != nil {
		routeAlreadyRegisteredError := fmt.Errorf("Route '%v' already registered for method %v.", path, httpMethod)
		panic(routeAlreadyRegisteredError)
	}

	registeredRoutes[buildRouteRegisterIdentifier(httpMethod, path)] = handler
	http.HandleFunc(path, HttpCoordinatorMethodMiddleware())
}

func buildRouteRegisterIdentifier(httpMethod, pathRoute string) string {
	return fmt.Sprintf("%v:%v", httpMethod, pathRoute)
}
