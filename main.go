package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/hillshum/gofib/fibonacci"
	"github.com/julienschmidt/httprouter"
)

func fib(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	n, err := strconv.Atoi(ps.ByName("n"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	enc := json.NewEncoder(w)
	enc.Encode(fibonacci.Fibonacci(n))
}

func main() {
	router := httprouter.New()
	router.GET("/fibonacci/:n", fib)

	log.Fatal(http.ListenAndServe(":8080", router))
}
