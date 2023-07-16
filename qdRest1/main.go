package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const endpoint = "http://localhost:6333/collections/"

// httpRequest is a utility function for making HTTP requests
func httpRequest(method, url string, body interface{}) (*http.Response, error) {
	bodyJSON, err := json.Marshal(body)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(req)
}

// handle.StatusCode is a utility function for handling the HTTP response
func handleStatusCode(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	return nil
}

func recreateCollection(collectionName string) error {
	url := endpoint + collectionName
	body := map[string]interface{}{
		"vectors": map[string]interface{}{
			"size":     4,
			"distance": "Dot",
		},
	}
	resp, err := httpRequest("PUT", url, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return handleStatusCode(resp)
}

func insertPoints(collectionName string, points []map[string]interface{}) error {
	url := endpoint + collectionName + "/points?wait=true"
	body := map[string]interface{}{
		"points": points,
	}
	resp, err := httpRequest("PUT", url, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return handleStatusCode(resp)
}

func deletePoints(collectionName string, pointIds []int) error {
	url := endpoint + collectionName + "/points/delete"
	body := map[string]interface{}{
		"points": pointIds,
	}
	resp, err := httpRequest("POST", url, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func deleteCollection(collectionName string) error {
	url := endpoint + collectionName
	resp, err := httpRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return handleStatusCode(resp)
}

// decodeResponse is a utility function for decoding HTTP responses
func decodeResponse(resp *http.Response) (map[string]interface{}, error) {
	var respData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	if respData["status"] != "ok" {
		return nil, fmt.Errorf("search failed: %s", respData["status"])
	}
	return respData, nil
}

func searchPoints(collectionName string, vector []float32, limit int, filter ...map[string]interface{}) error {
	url := endpoint + collectionName + "/points/search"
	reqBody := map[string]interface{}{
		"vector": vector,
		"limit":  limit,
	}
	// If there is a filter, add it to the request body
	if len(filter) > 0 {
		reqBody["filter"] = filter[0]
	}
	resp, err := httpRequest("POST", url, reqBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respData, err := decodeResponse(resp)
	if err != nil {
		return err
	}

	results, ok := respData["result"].([]interface{})
	if !ok {
		return fmt.Errorf("invalid result format")
	}

	fmt.Println(results)

	return nil
}

func searchPointsFilter(collectionName string, filter map[string]interface{}, vector []float32, limit int) error {
	return searchPoints(collectionName, vector, limit, filter)
}

func main() {
	err := recreateCollection("test3")
	if err != nil {
		panic(err)
	}

	points := []map[string]interface{}{
		{
			"id":      3,
			"vector":  []float32{0.36, 0.55, 0.47, 0.94},
			"payload": map[string]interface{}{"city": []string{"Berlin", "Moscow"}},
		},
		{
			"id":      4,
			"vector":  []float32{0.18, 0.01, 0.85, 0.80},
			"payload": map[string]interface{}{"city": []string{"London", "Moscow"}},
		},
		{
			"id":      5,
			"vector":  []float32{0.24, 0.18, 0.22, 0.44},
			"payload": map[string]interface{}{"count": []int{0}},
		},
		{
			"id":     6,
			"vector": []float32{0.35, 0.08, 0.11, 0.44},
		},
	}
	err = insertPoints("test3", points)
	if err != nil {
		panic(err)
	}

	vector := []float32{0.2, 0.1, 0.9, 0.7}
	limit := 3
	err = searchPoints("test3", vector, limit)
	if err != nil {
		panic(err)
	}

	filter := map[string]interface{}{
		"should": []map[string]interface{}{
			{
				"key":   "city",
				"match": map[string]interface{}{"value": "London"},
			},
		},
	}
	err = searchPointsFilter("test3", filter, vector, limit)
	if err != nil {
		panic(err)
	}

	pointIds := []int{0, 3, 100}
	err = deletePoints("test3", pointIds)
	if err != nil {
		panic(err)
	}

	err = deleteCollection("test3")
	if err != nil {
		panic(err)
	}
}
