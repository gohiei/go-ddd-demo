package entity

import (
	"fmt"
	"strings"
)

// IRoutes is a map that represents the routes to be ignored.
var (
	IRoutes = map[string]bool{
		"POST /api/user": true,
	}
)

// IgnoreRoute checks if a route should be ignored based on the method and URL.
func IgnoreRoute(method, url string) bool {
	route := fmt.Sprintf("%s %s", method, url)

	if !strings.HasPrefix(url, "/api") {
		return true
	}

	return IRoutes[route]
}
