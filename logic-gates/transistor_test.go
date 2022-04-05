package logic_gates

import (
	"testing"
)

func TestTransistorFalseAndFalse(t *testing.T) {
	var output bool = transistor(false, false)
	if output == true {
		t.Fatalf("transistor(false, false) should be false, but got true")
	}
}

func TestTransistorTrueAndFalse(t *testing.T) {
	var output bool = transistor(true, false)
	if output == true {
		t.Fatalf("transistor(true, false) should be false, but got true")
	}
}

func TestTransistorFalseAndTrue(t *testing.T) {
	var output bool = transistor(false, true)
	if output == true {
		t.Fatalf("transistor(false, true) should be false, but got true")
	}
}

func TestTransistorTrueAndTrue(t *testing.T) {
	var output bool = transistor(true, true)
	if output == false {
		t.Fatalf("transistor(true, true) should be true, but got false")
	}
}
