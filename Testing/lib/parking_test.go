package main

import "testing"

func TestParkingsystem(t *testing.T) {
	t.Run("testcase1 is ", func(t *testing.T) {
		ps := Constructor(1, 1, 0)
		//v1 := ps.AddCar(0)
		v := ps.AddCar(1)

		if v != true {
			t.Errorf("Expected %v ,  Got %v", true, v)
		}
	})
	t.Run("testcase2 is ", func(t *testing.T) {
		ps := Constructor(1, 1, 0)
		//v1 := ps.AddCar(0)
		v1 := ps.AddCar(0)

		if v1 != true {
			t.Errorf("Expected %v ,  Got %v", true, v1)
		}
	})

}
