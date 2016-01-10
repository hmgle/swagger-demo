// @SubApi 喵喵喵 [/Cat]
package api

import (
	"bytes"
	"net/http"
	"strconv"

	ren "github.com/hmgle/swagger-demo/src/render"
)

// @Title       api.Miao
// @Description 喵
// @Resource    Cat
// @Accept      json
// @Param       count query    int    false "喵几声?"
// @Success     200   {string} string "返回"
// @Router      /miao [get]
func Miao(w http.ResponseWriter, r *http.Request) {
	miao := bytes.Buffer{}
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		ren.Text(w, http.StatusOK, "喵...")
		return
	}
	for i := 0; i < count; i++ {
		miao.WriteString("喵~")
	}
	ren.Text(w, http.StatusOK, miao.String())
}
