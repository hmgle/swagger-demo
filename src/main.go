// @APIVersion 1.0.0
// @APITitle API
// @APIDescription API test
// @Contact test@test.com
// @BasePath /apis
package main

import (
	"encoding/json"
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"

	"github.com/hmgle/swagger-demo/src/api"
	ren "github.com/hmgle/swagger-demo/src/render"
)

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "HEAD, POST, OPTIONS, GET, PUT, DELETE")
		w.Header().Set("Access-Control-Max-Age", "1728000")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := chi.NewRouter()
	r.Use(EnableCors)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", apiIndex)
	r.Get("/apis/miao", api.Miao)

	// api server: http://127.0.0.1:3333/docs/
	r.Static("/docs", "./swagger-ui")
	r.Get("/docs", http.RedirectHandler("/docs/", 301))
	r.Get("/apis", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resourceListingJson))
	})
	for _, v := range apiDescriptionsJson {
		var apis ApiDeclaration
		if err := json.Unmarshal([]byte(v), &apis); err != nil {
			return
		}
		r.Get("/apis"+apis.ResourcePath, func(w http.ResponseWriter, r *http.Request) {
			ren.JSON(w, http.StatusOK, &apis)
		})
	}
	http.ListenAndServe(":3333", r)
}

// Handler
func apiIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("root!\n"))
}
