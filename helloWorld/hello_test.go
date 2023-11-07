package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, Test name"
	recieved := Hello("Test name")

	if recieved != expected {
		t.Errorf("recieved: %q\n expected: %q", recieved, expected)
	}
}