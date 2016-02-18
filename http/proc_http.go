package http

import (
	"github.com/niean/opsadmin/proc"
	"net/http"
)

func configProcHttpRoutes() {
	// counter
	http.HandleFunc("/counter/all", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, proc.GetAll())
	})

	// trace
	http.HandleFunc("/trace/", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, "not supported right now")
	})
}
