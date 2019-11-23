package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	//what is actually printed onto the webpage
	fmt.Fprintf(w, "Welcome to the HomePage!")
	//what is printed onto the console
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	//gorilla mux allows use to specify what verbs we 
	//can use and in what fashion we can use them
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	// add our articles route and map it to our
	// returnAllArticles function like so
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {

	//does the job of encoding our articles array into a
	//JSON string and then writing as part of our response.
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {

	//does the job of encoding our articles array into a
	//JSON string and then writing as part of our response.
	
	fmt.Fprintf(w, "Test POST endpoint worked")
}

func main() {

	handleRequests()
}
