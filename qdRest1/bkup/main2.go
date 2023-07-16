package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const baseURL = "http://localhost:6333/collections/"

// sendRequestAndHandleResponse is a helper function to send HTTP requests, handle the response,
// and unmarshal the response into the provided interface.
func sendRequestAndHandleResponse(method, url string, body interface{}, header map[string]string, result interface{}) error {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return err
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	if result != nil {
		err := json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			return err
		}
	}

	return nil
}

func recreateCollection(collectionName string) error {
	url := baseURL + collectionName
	body := map[string]interface{}{
		"vectors": map[string]interface{}{
			"size":     4,
			"distance": "Dot",
		},
	}
	return sendRequestAndHandleResponse("PUT", url, body, map[string]string{"Content-Type": "application/json"}, nil)
}

func insertPoints(collectionName string, points []map[string]interface{}) error {
	url := baseURL + collectionName + "/points?wait=true"
	body := map[string]interface{}{
		"points": points,
	}
	return sendRequestAndHandleResponse("PUT", url, body, map[string]string{"Content-Type": "application/json"}, nil)
}

func searchPoints(collectionName string, vector []float32, limit int) ([]interface{}, error) {
	url := baseURL + collectionName + "/points/search"
	body := map[string]interface{}{
		"vector": vector,
		"limit":  limit,
	}

	var respData map[string]interface{}
	err := sendRequestAndHandleResponse("POST", url, body, map[string]string{"Content-Type": "application/json"}, &respData)
	if err != nil {
		return nil, err
	}

	results, ok := respData["result"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid result format")
	}

	return results, nil
}

func searchPointsFilter(collectionName string, filter map[string]interface{}, vector []float32, limit int) ([]interface{}, error) {
	url := baseURL + collectionName + "/points/search"
	body := map[string]interface{}{
		"filter": filter,
		"vector": vector,
		"limit":  limit,
	}

	var respData map[string]interface{}
	err := sendRequestAndHandleResponse("POST", url, body, map[string]string{"Content-Type": "application/json"}, &respData)
	if err != nil {
		return nil, err
	}

	results, ok := respData["result"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid result format")
	}

	return results, nil
}

func deletePoints(collectionName string, pointIds []int) error {
	url := baseURL + collectionName + "/points/delete"
	body := map[string]interface{}{
		"points": pointIds,
	}
	return sendRequestAndHandleResponse("POST", url, body, map[string]string{"Content-Type": "application/json"}, nil)
}

func deleteCollection(collectionName string) error {
	url := baseURL + collectionName
	return sendRequestAndHandleResponse("DELETE", url, nil, nil, nil)
}

func main() {
	err := recreateCollection("test3")
	if err != nil {
		log.Fatalln(err)
	}

}
