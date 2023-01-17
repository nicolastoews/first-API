package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type CurrentPage struct {
	Current_page int        `json:"current_page"`
	Data         []CatBreed `json:"data"`
}

type CatBreed struct {
	Breed string `json:"breed"`
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/breeds", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("https://catfact.ninja/breeds")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(responseData)
		//fmt.Println(string(responseData))
		var currentPage CurrentPage
		errM := json.Unmarshal([]byte(responseData), &currentPage)
		if errM != nil {
			fmt.Println("error:", err)
		}
		fmt.Println((currentPage))
		for _, element := range currentPage.Data {
			w.Write([]byte(element.Breed))
		}
	})
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))

}
