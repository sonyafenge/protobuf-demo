package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/sonyafenge/protobuf-demo/proto/echo"
	"github.com/sonyafenge/protobuf-demo/proto/echojson"
)

func Echo(resp http.ResponseWriter, req *http.Request) {
	contentLength := req.ContentLength
	fmt.Printf("Content echo Length Received : %v\n", contentLength)
	request := &echo.EchoRequest{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}
	proto.Unmarshal(data, request)
	name := request.GetName()
	result := &echo.EchoResponse{Message: "Hello " + name}
	response, err := proto.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	resp.Write(response)

}

func EchoJson(resp http.ResponseWriter, req *http.Request) {
	contentLength := req.ContentLength
	fmt.Printf("Content json Length Received : %v\n", contentLength)
	request := &echojson.EchoJsonRequest{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}
	json.Unmarshal(data, request)
	name := request.Name
	result := &echojson.EchoJsonResponse{Message: "Hello " + name}
	response, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	resp.Write(response)
}

func main() {
	fmt.Println("Starting the API server...")
	r := mux.NewRouter()
	r.HandleFunc("/echo", Echo).Methods("POST")
	r.HandleFunc("/echo_json", EchoJson).Methods("POST")

	server := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
