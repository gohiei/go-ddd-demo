// Package restful provides functionality for integrating Swagger API documentation generation.
package restful

// nolint: gci
import (
	_ "cypt/docs/swagger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewSwaggerRestful(router *gin.Engine) {
	// @todo low performance when using swaggo/gen
	// err := gen.New().Build(&gen.Config{
	// 	SearchDir:       "./",
	// 	OutputDir:       swaggerDir,
	// 	OutputTypes:     []string{"go", "json", "yaml"},
	// 	MainAPIFile:     "main.go",
	// 	ParseDependency: true,
	// 	ParseDepth:      100,
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	router.GET("/docs/*any", ginSwagger.WrapHandler((swaggerFiles.Handler)))
}
