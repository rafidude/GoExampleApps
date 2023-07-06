package main

import "testing"

func TestGetFromDatabase(t *testing.T) {
	db := MockDatabase{}
	want := "bar"
	got, err := GetFromDatabase(db, "foo")
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
