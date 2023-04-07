package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	rtr  *mux.Router
	tmpl *template.Template
)

func main() {

	rtr = mux.NewRouter()
	rtr.Use(middleware)
	tmpl = template.Must(template.ParseGlob("app/dist/*.html"))
	rtr.HandleFunc("/", index)
	rtr.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("app/dist/assets"))))
	rtr.PathPrefix("/img").Handler(http.StripPrefix("/img", http.FileServer(http.Dir("app/dist/img"))))

	log.Println("Server running at http://localhost:8080")

	//****************** MUST BE END OF MAIN ******************\\
	s := &http.Server{
		Addr:           ":8080",
		Handler:        handlers.LoggingHandler(os.Stdout, rtr),
		MaxHeaderBytes: 1 << 62,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("server failed", err)
	}
	//****************** MUST BE END OF MAIN ******************\\
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cfg.Active.Environment == "Local" {
			w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Host, Content-Type, X-Amz-Date, Authorization, X-Api-Key, X-Amz-Security-Token, X-XSRF-TOKEN, Origin, Access-Control-Request-Origin, Access-Control-Request-Method, Access-Control-Request-Headers, Access-Control-Allow-Origin, access-control-allow-origin, Access-Control-Allow-Credentials, access-control-allow-credentials, Access-Control-Allow-Headers, access-control-allow-headers, Access-Control-Allow-Methods, access-control-allow-methods")
		}
		next.ServeHTTP(w, r)
	})
}
