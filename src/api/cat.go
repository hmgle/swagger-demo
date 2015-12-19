// @SubApi 喵喵喵 [/Cat]
package api

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// @Title       api.Miao
// @Description 喵
// @Resource    Cat
// @Accept      json
// @Param       count query    int    false "喵几声?"
// @Success     200   {string} string "返回"
// @Router      /miao [get]
func Miao(c *echo.Context) error {
	miao := bytes.Buffer{}
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		return c.String(http.StatusOK, "喵...")
	}
	for i := 0; i < count; i++ {
		miao.WriteString("喵~")
	}
	return c.String(http.StatusOK, miao.String())
}
