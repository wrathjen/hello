package main

import "testing"

func Test_hello(t *testing.T) {
	result := hello("World")
	expected := "Hello World"

	if result != expected {
		t.Errorf("got '%s' want '%s'", result, expected)
	}
}
