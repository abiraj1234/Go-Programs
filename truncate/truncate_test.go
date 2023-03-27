package main

import "testing"

func TestTrunccateSentence(t *testing.T) {

	t.Run("testcase1", func(t *testing.T) {

		input := "Hello how are you Contestant"
		want := "Hello how are you"
		k := 4
		res := trunccateSentence(input, k)
		if res != want {
			t.Errorf("TruncateSentence(%q, %d) = %q want %q", input, k, res, want)
		}
	})
	t.Run("testcase2", func(t *testing.T) {

		inputs := "Hello how are you Contestant"
		Expected := "Hello how are"
		k1 := 3
		result := trunccateSentence(inputs, k1)
		if result != Expected {
			t.Errorf("Expexted (%q, %d)= %q want %q ", inputs, k1, result, Expected)
		}

	})
}
