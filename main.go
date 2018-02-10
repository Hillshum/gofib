package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/hillshum/gofib/fibonacci"
	"github.com/julienschmidt/httprouter"
)

func fib(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	n, _ := strconv.Atoi(ps.ByName("n"))
	fibs := fibonacci.Fibonacci(n)
	fmt.Print(fibs)
	fmt.Fprintf(w, "The fib series is %v\n", fibs)

}

func main() {
	router := httprouter.New()
	router.GET("/fibonacci/:n", fib)

	log.Fatal(http.ListenAndServe(":8080", router))
}
