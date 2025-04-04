package main

import (
	"encoding/json"
	"featureGen/controller"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/karate", func(writer http.ResponseWriter, request *http.Request) {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		defer request.Body.Close()
		var req map[string]interface{}
		if err = json.Unmarshal(body, &req); err != nil {
			panic(err)
		}
		controller.Orchestrate(req)
	}).Methods("POST")
	fmt.Println("serving on port 9090")

	http.ListenAndServe(":9090", router)
}
