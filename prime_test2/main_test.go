package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_prompt(t *testing.T) {
	var buf bytes.Buffer
	prompt(&buf)

	want := "Enter a number: "
	if buf.String() != want {
		t.Errorf("prompt() = %q, want %q", buf.String(), want)
	}
}

func Test_intro(t *testing.T) {
	var buf bytes.Buffer
	intro(&buf)

	want := "Is it a Prime?\nEnter a whole number to check if it is a prime number.\nEnter q to exit.\nEnter a number: "
	if buf.String() != want {
		t.Errorf("intro() = %q, want %q", buf.String(), want)
	}
}

// func Test_checkNumbers(t *testing.T) {
// 	testCases := []struct {
// 		input    string
// 		expected string
// 		exit     bool
// 	}{
// 		{"q", "", true},
// 		{"10", "10 is not a prime number!", false},
// 		{"5", "5 is a prime number!", false},
// 		{"abc", "Please enter a whole number!", false},
// 	}

// 	for _, tc := range testCases {
// 		scanner := bufio.NewScanner(strings.NewReader(tc.input))
// 		got, exit := checkNumbers(scanner)

// 		if got != tc.expected || exit != tc.exit {
// 			t.Errorf("checkNumbers(%q) = %q, %t, want %q, %t", tc.input, got, exit, tc.expected, tc.exit)
// 		}
// 	}
// }

func Test_readUserInput(t *testing.T) {
	// to test this function, we need a channel, and an instance of an io.Reader
	doneChan := make(chan bool)

	// create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not prime, by definition!"},
		{name: "one", input: "1", expected: "1 is not prime, by definition!"},
		{name: "two", input: "2", expected: "2 is a prime number!"},
		{name: "three", input: "3", expected: "3 is a prime number!"},
		{name: "negative", input: "-1", expected: "Negative numbers are not prime, by definition!"},
		{name: "typed", input: "three", expected: "Please enter a whole number!"},
		{name: "decimal", input: "1.1", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "QUIT", input: "Q", expected: ""},
		{name: "greek", input: "επτά", expected: "Please enter a whole number!"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}
}
