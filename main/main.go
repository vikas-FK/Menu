package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"menu/model"
	"encoding/json"
)

/*
	Approach 1
 */
//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprint(w, "Hello World")
//	})
//
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}

/*
	Approach 2
 */
//func main() {
//	router := mux.NewRouter()
//	router.HandleFunc("/", Index)
//
//	log.Fatal(http.ListenAndServe(":8080", router))
//}
//
//func Index(w http.ResponseWriter, r *http.Request){
//	fmt.Fprint(w, "Hello World")
//}

/*
	Approach 3
 */
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/{name}", GreetingInPerson)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World")
}

func GreetingInPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	helloWorld := model.HelloWorld{
		Name: name,
		Message: "Welcome to the GoLang",
	}

	bytes, err := json.Marshal(helloWorld)
	if err != nil {
		fmt.Fprintln(w, "Unexpected Error")
	}

	w.Write(bytes)
}