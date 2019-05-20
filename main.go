package main

import (
	"./fizzbuzz"
	"./input"

	"encoding/json"
	"log"
	"net/http"
)

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	params, err := input.ParseParams(r.URL.Query())
	if err == nil {
		log.Println(params, input.RegisterStats(params))
		f := fizzbuzz.New(params.Int1, params.Int2, params.Str1, params.Str2)
		res := make([]string, params.Limit)
		for i := range res {
			res[i] = f(i + 1)
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		log.Println(r.URL, 200)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode([]string(err))
	log.Println(r.URL, 400)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	mostPopular := input.MostPopular()
	log.Println(mostPopular)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mostPopular)
}

func main() {
	http.HandleFunc("/fizzbuzz", fizzbuzzHandler)
	http.HandleFunc("/statistics", statsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
