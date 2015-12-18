// @SubApi 喵喵喵 [/Cat]
package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// @Title       api.Miao
// @Description 喵
// @Resource    Cat
// @Accept      json
// @Success     200   {string} string "返回"
// @Router      /miao [get]
func Miao(c *echo.Context) error {
	return c.String(http.StatusOK, "喵~\n")
}
