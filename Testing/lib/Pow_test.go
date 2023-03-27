package lib

import "testing"

func TestPow(t *testing.T) {
	v := Pow(3, 4)
	if v != 81 {
		t.Errorf("Expected %d, Got %d", 81, v)
	}
	v = Pow(3, 3)
	if v != 27 {
		t.Errorf("Expected %d, Got %d", 27, v)
	}
}
