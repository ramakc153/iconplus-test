package practicaltest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetSummary(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running")
}

func GetSummaryApi(w http.ResponseWriter, r *http.Request) {
	apiData := ConstructSummaryApi()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apiData)
	if err != nil {
		log.Println("error when encoding data: ", err.Error())
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}

}

func Server() {
	http.HandleFunc("/", GetSummary)
	http.HandleFunc("/api", GetSummaryApi)
	addr := ":8080"
	log.Printf("server running on %s", addr)
	log.Println(http.ListenAndServe(addr, nil))
}
