// @APIVersion 1.0.0
// @APITitle API
// @APIDescription API test
// @Contact test@test.com
// @BasePath
package main

import (
	"encoding/json"
	"net/http"

	"github.com/hmgle/swagger-demo/src/api"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(cors())
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Get("/", apiIndex)
	e.Get("/apis/miao", api.Miao)

	// api server: http://127.0.0.1:3333/docs/
	e.Static("/docs/", "./swagger-1")
	e.Get("/apis", func(c *echo.Context) error {
		return c.String(200, resourceListingJson)
	})
	for _, v := range apiDescriptionsJson {
		var apis ApiDeclaration
		if err := json.Unmarshal([]byte(v), &apis); err != nil {
			return
		}
		e.Get("/apis"+apis.ResourcePath, func(c *echo.Context) error {
			return c.JSON(200, &apis)
		})
	}

	e.Run(":3333")
}

// 允许跨域请求
func cors() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			c.Response().Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
			c.Response().Header().Set("Access-Control-Max-Age", "1728000")
			return h(c)
		}
	}
}

// Handler
func apiIndex(c *echo.Context) error {
	return c.String(http.StatusOK, "root!\n")
}
