package render

import (
	"net/http"

	renderer "github.com/unrolled/render"
)

var (
	Renderer *renderer.Render
)

func init() {
	Renderer = renderer.New()
}

func Render(w http.ResponseWriter, e renderer.Engine, data interface{}) error {
	return Renderer.Render(w, e, data)
}

func Data(w http.ResponseWriter, status int, v []byte) error {
	return Renderer.Data(w, status, v)
}

func HTML(w http.ResponseWriter, status int, name string, binding interface{}, htmlOpt ...renderer.HTMLOptions) error {
	return Renderer.HTML(w, status, name, binding, htmlOpt...)
}

func JSON(w http.ResponseWriter, status int, v interface{}) error {
	return Renderer.JSON(w, status, v)
}

func JSONP(w http.ResponseWriter, status int, callback string, v interface{}) error {
	return Renderer.JSONP(w, status, callback, v)
}

func Text(w http.ResponseWriter, status int, v string) error {
	return Renderer.Text(w, status, v)
}

func XML(w http.ResponseWriter, status int, v interface{}) error {
	return Renderer.XML(w, status, v)
}
