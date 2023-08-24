package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		return
	}
	fmt.Fprintf(w, "Hola %s", "fabricio")

}

func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(countries)
}
func addCountries(w http.ResponseWriter, r *http.Request) {

	contry := &Country{}

 err := json.NewDecoder(r.Body).Decode(contry)
	if err != nil {
	
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"%v",err)
		return
	}
	countries = append(countries, contry)
	fmt.Fprintf(w,"country was addedd")
	
}
