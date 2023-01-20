package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/gorilla/mux"
)

type BreedMetadata struct {
	Currentpage int        `json:"current_page"`
	Data        []CatBreed `json:"data"`
}

type CatBreed struct {
	Breed string `json:"breed"`
}

type FactMetadata struct {
	Currentpage int       `json:"current_page"`
	Data        []CatFact `json:"data"`
}

type CatFact struct {
	Fact string `json:"fact"`
}

type ByLen []string

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Less(i, j int) bool { return len(a[i]) < len(a[j]) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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
		var metadata BreedMetadata
		err = json.Unmarshal([]byte(responseData), &metadata)
		if err != nil {
			fmt.Println("error:", err)
		}

		catBreedList := make([]string, 0, len(metadata.Data))
		for _, element := range metadata.Data {
			catBreedList = append(catBreedList, element.Breed)
		}
		sort.Sort(sort.Reverse(sort.StringSlice(catBreedList)))
		res, err := json.Marshal(catBreedList)
		if err != nil {
			fmt.Println("error:", err)
		}
		w.Write(res)
	})

	r.HandleFunc("/fact", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("https://catfact.ninja/fact")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		var metadata FactMetadata
		err = json.Unmarshal([]byte(responseData), &metadata)
		if err != nil {
			fmt.Println("error:", err)
		}
		w.Write(responseData)
	})

	r.HandleFunc("/facts", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("https://catfact.ninja/facts")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		var metadata BreedMetadata
		err = json.Unmarshal([]byte(responseData), &metadata)
		if err != nil {
			fmt.Println("error:", err)
		}

		FactsList := make([]string, 0, len(metadata.Data))
		for _, element := range metadata.Data {
			FactsList = append(FactsList, element.Breed)
		}
		sort.Strings(FactsList)

		sort.Sort(ByLen(FactsList))

		res, err := json.Marshal(FactsList)
		if err != nil {
			fmt.Println("error:", err)
		}
		w.Write(res)
	})

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))

}
