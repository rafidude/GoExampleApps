package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func recreateCollection(collectionName string) error {

	url := "http://localhost:6333/collections/" + collectionName

	// Create request body
	bodyJSON, err := json.Marshal(map[string]interface{}{
		"vectors": map[string]interface{}{
			"size":     4,
			"distance": "Dot",
		},
	})
	if err != nil {
		return err
	}
	fmt.Println(string(bodyJSON))

	// Make PUT request
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	return nil
}

func insertPoints(collectionName string, points []map[string]interface{}) error {

	url := "http://localhost:6333/collections/" + collectionName + "/points?wait=true"

	// Create request body
	bodyJSON, err := json.Marshal(map[string]interface{}{
		"points": points,
	})
	if err != nil {
		return err
	}

	// Create request
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Check response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	return nil
}

func searchPoints(collectionName string, vector []float32, limit int) error {

	url := "http://localhost:6333/collections/" + collectionName + "/points/search"

	// Create request body
	body, err := json.Marshal(map[string]interface{}{
		"vector": vector,
		"limit":  limit,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Check response
	var respData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return err
	}

	if respData["status"] != "ok" {
		return fmt.Errorf("search failed: %s", respData["status"])
	}

	results, ok := respData["result"].([]interface{})
	if !ok {
		return fmt.Errorf("invalid result format")
	}

	fmt.Println(results)

	return nil
}

func searchPointsFilter(collectionName string, filter map[string]interface{}, vector []float32, limit int) error {

	url := "http://localhost:6333/collections/" + collectionName + "/points/search"

	// Create request body
	body, err := json.Marshal(map[string]interface{}{
		"filter": filter,
		"vector": vector,
		"limit":  limit,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Handle response
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check response format
	var respData map[string]interface{}
	if err := json.Unmarshal(respBody, &respData); err != nil {
		return err
	}

	if respData["status"] != "ok" {
		return fmt.Errorf("search failed: %s", respData["status"])
	}

	results := respData["result"].([]interface{})

	// Validate results
	fmt.Println(results)

	return nil

}

func deletePoints(collectionName string, pointIds []int) error {

	url := "http://localhost:6333/collections/" + collectionName + "/points/delete"

	// Create request body
	body, err := json.Marshal(map[string]interface{}{
		"points": pointIds,
	})
	if err != nil {
		return err
	}

	// Create DELETE request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle response

	return nil
}

func deleteCollection(collectionName string) error {

	url := "http://localhost:6333/collections/" + collectionName

	// Create DELETE request
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error deleting collection, status: %d", resp.StatusCode)
	}

	return nil
}

func test_all_functions() {
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
	insertPoints("test3", points)

	vector := []float32{0.2, 0.1, 0.9, 0.7}
	limit := 3
	searchPoints("test3", vector, limit)

	filter := map[string]interface{}{
		"should": []map[string]interface{}{
			{
				"key":   "city",
				"match": map[string]interface{}{"value": "London"},
			},
		},
	}
	searchPointsFilter("test3", filter, []float32{0.2, 0.1, 0.9, 0.7}, 3)

	pointIds := []int{0, 3, 100}
	deletePoints("test3", pointIds)

	deleteCollection("test3")
}
