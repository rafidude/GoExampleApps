package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Shared client
var client = &http.Client{}

func printResults(respData map[string]interface{}, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(respData)
	results := respData["result"].(map[string]interface{})

	// Validate results
	fmt.Println(results)
}

// Helper to make request
func makeRequest(req *http.Request) (*http.Response, error) {
	return client.Do(req)
}

// Helper to create JSON body
func jsonBody(v interface{}) (io.Reader, error) {
	body, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}

// Helper to handle response
func handleResponse(resp *http.Response) (map[string]interface{}, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, err
	}
	return respData, nil
}

// Refactored functions

func recreateCollection(name string) error {
	req, _ := http.NewRequest("PUT", "http://localhost:6333/collections/"+name, nil)
	resp, err := makeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// check response
	respData, err := handleResponse(resp)
	if err != nil {
		fmt.Println(err)
		return err
	}
	printResults(respData, nil)
	return err
}

func insertPoints(name string, points []map[string]interface{}) error {
	body, _ := jsonBody(map[string]interface{}{"points": points})
	req, _ := http.NewRequest("PUT", "/collections/"+name+"/points", body)
	resp, err := makeRequest(req)
	// check response
	log.Println(handleResponse(resp))
	return err
}

func searchPoints(name string, vector []float32, limit int) error {
	body, _ := jsonBody(map[string]interface{}{"vector": vector, "limit": limit})
	req, _ := http.NewRequest("POST", "/collections/"+name+"/search", body)
	resp, err := makeRequest(req)
	// handle response
	log.Println(handleResponse(resp))
	return err
}

// Shared helpers...

func searchPointsFilter(name string, filter map[string]interface{}, vector []float32, limit int) error {
	body, _ := jsonBody(map[string]interface{}{"filter": filter, "vector": vector, "limit": limit})
	req, _ := http.NewRequest("POST", "/collections/"+name+"/search", body)
	resp, err := makeRequest(req)
	// handle response
	log.Println(handleResponse(resp))
	return err
}

func deletePoints(name string, pointIds []int) error {
	body, _ := jsonBody(map[string]interface{}{"points": pointIds})
	req, _ := http.NewRequest("POST", "/collections/"+name+"/delete", body)
	resp, err := makeRequest(req)
	// check response
	log.Println(handleResponse(resp))
	return err
}

func deleteCollection(name string) error {
	req, _ := http.NewRequest("DELETE", "/collections/"+name, nil)
	resp, err := makeRequest(req)
	// check response
	log.Println(handleResponse(resp))
	return err
}

// Main function

func main() {

	// Create collection
	err := recreateCollection("test3")
	if err != nil {
		log.Fatal(err)
	}

	//   // Insert points
	//   points := []map[string]interface{}{
	//     {"id": 1, "vector": [0.1, 0.2], "payload": {"foo": "bar"}},
	//   }
	//   insertPoints("test", points)

	//   // Search
	//   vector := []float32{0.1, 0.2}
	//   searchPoints("test", vector, 10)

	//   // etc...

	//   // Delete collection
	//   deleteCollection("test")

}
