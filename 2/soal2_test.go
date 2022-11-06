package main

import (
	"bytes"
	"io"
	"testing"
)

func TestUserInputOne(t *testing.T) {
	input := "23242526272830"
	expect := 29
	var reader io.Reader = bytes.NewBufferString(input)
	err := search(reader)
	t.Log("Input :", input)
	// fmt.Println(err)

	if err != expect {
		t.Errorf("Expected %d, but got %d", expect, err)
	}

}
func TestUserInputTwo(t *testing.T) {
	input := "101102103104105106107108109111112113"
	expect := 110
	var reader io.Reader = bytes.NewBufferString(input)
	err := search(reader)
	t.Log("Input :", input)
	// fmt.Println(err)

	if err != expect {
		t.Errorf("Expected %d, but got %d", expect, err)
	}
}

func TestUserInputThree(t *testing.T) {
	input := "12346789"
	expect := 5
	var reader io.Reader = bytes.NewBufferString(input)
	err := search(reader)
	t.Log("Input :", input)
	// fmt.Println(err)

	if err != expect {
		t.Errorf("Expected %d, but got %d", expect, err)
	}

}
