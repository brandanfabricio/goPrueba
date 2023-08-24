package server

import (
	"fmt"
	"net/http"
)

func initRouter() {

	http.HandleFunc("/", index)
	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCountries(w, r)
		case http.MethodPost:
			addCountries(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Metodo no permitido")
			return
		}
	})

}
