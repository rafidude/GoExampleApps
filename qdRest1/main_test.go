package main

import (
	"testing"
)

// Test for deleteCollection
func TestDeleteCollection(t *testing.T) {
	err := deleteCollection("test3")
	if err != nil {
		t.Errorf("Failed to delete collection: %v", err)
	}
}

// Test for recreateCollection
func TestRecreateCollection(t *testing.T) {
	err := recreateCollection("test3")
	if err != nil {
		t.Errorf("Failed to recreate collection: %v", err)
	}
}

// Test for insertPoints
func TestInsertPoints(t *testing.T) {
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

	err := insertPoints("test3", points)
	if err != nil {
		t.Errorf("Failed to insert points: %v", err)
	}
}

// Test for searchPoints
func TestSearchPoints(t *testing.T) {
	vector := []float32{0.2, 0.1, 0.9, 0.7}
	limit := 3
	err := searchPoints("test3", vector, limit)
	if err != nil {
		t.Errorf("Failed to search points: %v", err)
	}
}

// Test for searchPointsFilter
func TestSearchPointsFilter(t *testing.T) {
	filter := map[string]interface{}{
		// the filter used in your main function...
	}
	vector := []float32{0.2, 0.1, 0.9, 0.7}
	limit := 3
	err := searchPointsFilter("test3", filter, vector, limit)
	if err != nil {
		t.Errorf("Failed to search points with filter: %v", err)
	}
}

// Test for deletePoints
func TestDeletePoints(t *testing.T) {
	pointIds := []int{0, 3, 100}
	err := deletePoints("test3", pointIds)
	if err != nil {
		t.Errorf("Failed to delete points: %v", err)
	}
}
