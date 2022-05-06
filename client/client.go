package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/sonyafenge/protobuf-demo/proto/echo"
	"github.com/sonyafenge/protobuf-demo/proto/echojson"
)

func makeRequest(request *echo.EchoRequest) *echo.EchoResponse {

	req, err := proto.Marshal(request)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	resp, err := http.Post("http://0.0.0.0:8080/echo", "application/x-binary", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &echo.EchoResponse{}
	proto.Unmarshal(respBytes, respObj)
	return respObj

}

func makeJsonRequest(request *echojson.EchoJsonRequest) *echojson.EchoJsonResponse {

	req, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	resp, err := http.Post("http://0.0.0.0:8080/echo_json", "application/json", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &echojson.EchoJsonResponse{}
	json.Unmarshal(respBytes, respObj)
	return respObj

}

func main() {

	var totalPBTime, totalJSONTime int64
	requestPb := &echo.EchoRequest{Name: "Sushil"}
	for i := 1; i <= 10; i++ {
		fmt.Printf("Sending echo request %v\n", i)
		startTime := time.Now()
		makeRequest(requestPb)
		elapsed := time.Since(startTime)
		totalPBTime += elapsed.Nanoseconds()
	}

	requestJson := &echojson.EchoJsonRequest{Name: "Sushil"}

	for i := 1; i <= 10; i++ {
		fmt.Printf("Sending json request %v\n", i)
		startTime := time.Now()
		makeJsonRequest(requestJson)
		elapsed := time.Since(startTime)
		totalJSONTime += elapsed.Nanoseconds()
	}

	fmt.Printf("Average Protobuf Response time : %v nano-seconds\n", totalPBTime/1000)
	fmt.Printf("Average JSON Response time : %v nano-seconds\n", totalJSONTime/1000)
}
