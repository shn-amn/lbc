package main

import (
	"./fizzbuzz"
	"./input"

	"encoding/json"
	"log"
	"net/http"
)

func fizzbuzzList(r *http.Request) (int, interface{}) {
	params, err := input.ParseParams(r.URL.Query())
	if err == nil {
		input.RegisterHit(params)
		f := fizzbuzz.New(params.Int1, params.Int2, params.Str1, params.Str2)
		res := make([]string, params.Limit)
		for i := range res {
			res[i] = f(i + 1)
		}
		return http.StatusOK, res
	}
	return http.StatusBadRequest, err
}

func stats(r *http.Request) (int, interface{}) {
	return http.StatusOK, input.MostPopular()
}

func main() {
	http.HandleFunc("/fizzbuzz", jsonGetOnlyHandler(fizzbuzzList))
	http.HandleFunc("/statistics", jsonGetOnlyHandler(stats))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func jsonGetOnlyHandler(f func(*http.Request) (int, interface{})) (handler func(http.ResponseWriter, *http.Request)) {
	handler = func(w http.ResponseWriter, r *http.Request) {
		var status int
		var response interface{}
		if r.Method == http.MethodGet {
			status, response = f(r)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(response)
		} else {
			status = http.StatusNotFound
			http.NotFound(w, r)
		}
		log.Println(r.Method, r.URL, status)
	}
	return
}
