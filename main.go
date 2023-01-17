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
		//w.Write(responseData)
		//fmt.Println(string(responseData))
		var metadata BreedMetadata
		err = json.Unmarshal([]byte(responseData), &metadata)
		if err != nil {
			fmt.Println("error:", err)
		}

		catBreedList := make([]string, 0, len(metadata.Data))
		for _, element := range metadata.Data {
			//catBreedList[i] = element.Breed
			catBreedList = append(catBreedList, element.Breed)
		}
		//res, err := json.Marshal(catBreedList)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }
		sort.Sort(sort.Reverse(sort.StringSlice(catBreedList)))
		res, err := json.Marshal(catBreedList)
		if err != nil {
			fmt.Println("error:", err)
		}
		w.Write(res)
	})
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))

}
