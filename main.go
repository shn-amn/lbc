package main

import (
	"./fizzbuzz"

	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"regexp"
	"strconv"
	"errors"
)

var validPath = regexp.MustCompile("^/([1-9][0-9]*)/([1-9][0-9]*)/([1-9][0-9]*)/([a-zA-Z]+)/([a-zA-Z]+)$")

func getParams(w http.ResponseWriter, r *http.Request) (ints [3]int, strs [2]string, err error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		err = errors.New("Invalid parameters.")
		return
	}
	err = nil
	for i := range ints {
		ints[i], _ = strconv.Atoi(m[1 + i])
	}
	for i := range strs {
		strs[i] = m[4 + i]
	}
	return
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	if ints, strs, err := getParams(w, r); err == nil {
		fmt.Println(ints, strs)
		f := fizzbuzz.New(ints[0], ints[1], strs[0], strs[1])
		res := make([]string, ints[2])
		for i := range res {
			res[i] = f(i + 1)
		}
		json.NewEncoder(w).Encode(res)
	}
}

func main() {
	http.HandleFunc("/", fizzbuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

