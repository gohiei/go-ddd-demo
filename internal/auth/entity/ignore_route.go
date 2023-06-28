package auth

import "fmt"

var (
	IRoutes = map[string]bool{
		"POST /api/user": true,
	}
)

func IgnoreRoute(method, url string) bool {
	route := fmt.Sprintf("%s %s", method, url)

	return IRoutes[route]
}
