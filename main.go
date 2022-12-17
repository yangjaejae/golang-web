package main

import (
	"net/http"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello World")
	// })

	// http.HandleFunc("/bar", barHandler)

	// http.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
