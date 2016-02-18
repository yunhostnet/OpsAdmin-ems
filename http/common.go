package http

import (
	"fmt"
	"github.com/niean/opsadmin/g"
	"net/http"
)

func configCommonRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%s\n", g.VERSION)))
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Access ip:%s\n", r.RemoteAddr)))
	})

}
