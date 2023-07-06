package main

import "fmt"

type Database interface {
	Get(key string) (string, error)
}

type MyDatabase struct{}

func (db MyDatabase) Get(key string) (string, error) {
	// implementation goes here
	return "", nil
}

func GetFromDatabase(db Database, key string) (string, error) {
	return db.Get(key)
}

type MockDatabase struct{}

func (db MockDatabase) Get(key string) (string, error) {
	if key == "foo" {
		return "bar", nil
	}
	return "", fmt.Errorf("key not found")
}
