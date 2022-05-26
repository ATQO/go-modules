package mutils

import "testing"

func TestSliceContains(t *testing.T) {

	// preparing slices of different types
	stringSlices := []string{"hello", "world"}
	intSlices := []int{0, 1, 2, 3}
	floatSlices := []float64{1.0, 2.0, 3.5}

	// positive contains string
	if res := SliceContains(stringSlices, "world"); res != true {
		t.Fatalf("the result of test should be true, instead of false")
	}

	// negative contains string
	if res := SliceContains(stringSlices, "sea"); res != false {
		t.Fatalf("the result of test should be false, instead of true")
	}

	// positive contains int
	if res := SliceContains(intSlices, 1); res != true {
		t.Fatalf("the result of test should be true, instead of false")
	}

	// negative contains int
	if res := SliceContains(intSlices, 5); res != false {
		t.Fatalf("the result of test should be false, instead of true")
	}

	// positive contains float64
	if res := SliceContains(floatSlices, 1.0); res != true {
		t.Fatalf("the result of test should be true, instead of false")
	}

	// negative contains float64
	if res := SliceContains(floatSlices, 5.0); res != false {
		t.Fatalf("the result of test should be false, instead of true")
	}

}
