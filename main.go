package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/xd", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("https://catfact.ninja/fact")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(responseData)
		fmt.Println(string(responseData))
	})
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
