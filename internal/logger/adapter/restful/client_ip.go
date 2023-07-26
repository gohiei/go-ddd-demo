package restful

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func GetClientIP(router *gin.Engine, trustedProxy []string, trustedHeader string) gin.HandlerFunc {
	router.TrustedPlatform = trustedHeader
	router.ForwardedByClientIP = true

	fullTrustedProxy := mergeTrustedProxy(trustedProxy)

	if err := router.SetTrustedProxies(fullTrustedProxy); err != nil {
		log.Fatal(err)
	}

	return func(ctx *gin.Context) {
		if ip, ok := parseFromAliCloud(); ok == nil {
			setClientIP(ctx, trustedHeader, ip)
			ctx.Next()

			return
		}

		ip, err := parseFromTunnel(ctx.ClientIP())

		if err != nil {
			ctx.Abort()
			return
		}

		setClientIP(ctx, trustedHeader, ip)
		ctx.Next()
	}
}

func setClientIP(ctx *gin.Context, header string, ip string) {
	ctx.Request.Header.Set(header, ip)
}

// @todo get trusted proxy from external service
func mergeTrustedProxy(proxy []string) []string {
	return proxy
}

// @todo
func parseFromAliCloud() (string, error) {
	return "", nil
}

// @todo
func parseFromTunnel(ip string) (string, error) {
	if !inTunnel(ip) {
		return ip, nil
	}

	if isApp() {
		return parseFromApp()
	}

	if isBrowser() {
		return parseFromBrowser()
	}

	return "", errors.New("unrecognized client ip")
}

// @todo
func inTunnel(ip string) bool {
	return false
}

// @todo
func isApp() bool {
	return false
}

// @todo
func parseFromApp() (string, error) {
	return "", nil
}

// @todo
func isBrowser() bool {
	return false
}

// @todo
func parseFromBrowser() (string, error) {
	return "", nil
}
