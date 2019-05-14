package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	//vendor packages
	"github.com/gorilla/mux"

	"github.com/spaceCh1mp/lil-url/handler"
)

const yamlFile string = "./url.yaml"

func main() {
	//mux using gorilla/mux package
	mux := gorillaMux()
	//fetch yaml file and pass content to var yaml
	yaml, err := ioutil.ReadFile(yamlFile)
	handler.Check(err)
	//call yaml handler to parse file, register paths:urls
	yamlHandler, err := handler.HandleYAML(yaml, mux)
	handler.Check(err)
	fmt.Println("Starting the server on :3000")
	//start server
	http.ListenAndServe(":3000", yamlHandler)
}

func gorillaMux() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	short, _ := ioutil.ReadFile(yamlFile)
	msg := "Example: http://localhost:3000/twit.ly  will take redirect you to twitter.com\n\n" + string(short)
	fmt.Fprintln(w, msg)
}
