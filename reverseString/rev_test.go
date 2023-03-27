package main

import (
	"bytes"
	"testing"
)

func TestReverseString(t *testing.T) {
	t.Run("revesecase1", func(t *testing.T) {
		input := []byte("hello")
		want := []byte("olleh")

		reverseString(input)

		if !bytes.Equal(input, want) {
			t.Errorf("Expected %q  , Got %q", input, want)
		}
	})
	t.Run("reversecase2", func(t *testing.T) {
		input := []byte("hello")
		want := []byte("olleh")

		reverseString(input)

		if bytes.Equal(input, want) {
			t.Errorf("Expected %q, Got %q", input, want)
		}
	})
}
