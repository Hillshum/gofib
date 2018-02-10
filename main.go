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
	if n < 0 {
		http.Error(w, "Input must be non-negative", 400)
		return
	}
	enc := json.NewEncoder(w)
	fibs, err := fibonacci.Fibonacci(n)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	enc.Encode(fibs)
}

func setupRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/fibonacci/:n", fib)
	return router
}

func main() {
	router := setupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
