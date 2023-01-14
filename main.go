package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("https://catfact.ninja/#")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	})
	r.HandleFunc("/Nachin", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("General Kenobi!\n"))
	})
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))

	// srv := &http.Server{
	// 	Handler:      r,
	// 	Addr:         "127.0.0.1:8080",
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// log.Fatal(srv.ListenAndServe())

}
