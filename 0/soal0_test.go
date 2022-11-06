package main

import (
	"bytes"
	"io"
	"testing"
)

func TestUserInputOne(t *testing.T) {
	input := "1 10"
	expect := 9
	var reader io.Reader = bytes.NewBufferString(input)
	err := countpalindrome(reader)
	t.Log("Input :", input)

	if err != expect {
		t.Errorf("Expected %d, but got %d", expect, err)
	}

}
func TestUserInputTwo(t *testing.T) {
	input := "99 100"
	expect := 1
	var reader io.Reader = bytes.NewBufferString(input)
	err := countpalindrome(reader)
	t.Log("Input :", input)

	if err != expect {
		t.Errorf("Expected %d, but got %d", expect, err)
	}

}
func TestUserInputThree(t *testing.T) {
	input := "21 31"
	expect := 1
	var reader io.Reader = bytes.NewBufferString(input)
	err := countpalindrome(reader)
	t.Log("Input :", input)

	if err != expect {
		t.Errorf("Expected %d, but got %d", expect, err)
	}

}
