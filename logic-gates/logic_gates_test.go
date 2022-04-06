package logic_gates

import (
	"fmt"
	"testing"
)

func TestAndGate(t *testing.T) {
	// construct a truth table to loop through
	truth_table := [4][3]bool{{false, false, false}, {true, false, false}, {false, true, false}, {true, true, true}}

	for _, values := range truth_table {
		var result bool = And_gate(values[0], values[1])

		if result != values[2] {
			t.Fatalf(fmt.Sprintf("and_gate(%t, %t) should be %t, but got %t", values[0], values[1], values[2], result))
		}
	}
}

func TestNotGate(t *testing.T) {
	// construct a truth table to loop through
	truth_table := [2][2]bool{{false, true}, {true, false}}

	for _, values := range truth_table {
		var result bool = Not_gate(values[0])

		if result != values[1] {
			t.Fatalf(fmt.Sprintf("not_gate(%t) should be %t, but got %t", values[0], values[1], result))
		}
	}
}

func TestOrGate(t *testing.T) {
	// construct a truth table to loop through
	truth_table := [4][3]bool{{false, false, false}, {true, false, true}, {false, true, true}, {true, true, true}}

	for _, values := range truth_table {
		var result bool = Or_gate(values[0], values[1])

		if result != values[2] {
			t.Fatalf(fmt.Sprintf("or_gate(%t, %t) should be %t, but got %t", values[0], values[1], values[2], result))
		}
	}
}

func TestXorGate(t *testing.T) {
	// construct a truth table to loop through
	truth_table := [4][3]bool{{false, false, false}, {true, false, true}, {false, true, true}, {true, true, false}}

	for _, values := range truth_table {
		var result bool = Xor_gate(values[0], values[1])

		if result != values[2] {
			t.Fatalf(fmt.Sprintf("xor_gate(%t, %t) should be %t, but got %t", values[0], values[1], values[2], result))
		}
	}
}

func TestNandGate(t *testing.T) {
	// construct a truth table to loop through
	truth_table := [4][3]bool{{false, false, true}, {true, false, true}, {false, true, true}, {true, true, false}}

	for _, values := range truth_table {
		var result bool = Nand_gate(values[0], values[1])

		if result != values[2] {
			t.Fatalf(fmt.Sprintf("nand_gate(%t, %t) should be %t, but got %t", values[0], values[1], values[2], result))
		}
	}
}

func TestNorGate(t *testing.T) {
	// construct a truth table to loop through
	truth_table := [4][3]bool{{false, false, true}, {true, false, false}, {false, true, false}, {true, true, false}}

	for _, values := range truth_table {
		var result bool = nor_gate(values[0], values[1])

		if result != values[2] {
			t.Fatalf(fmt.Sprintf("nor_gate(%t, %t) should be %t, but got %t", values[0], values[1], values[2], result))
		}
	}
}

func TestXnorGate(t *testing.T) {
	// construct a truth table to loop through
	truth_table := [4][3]bool{{false, false, true}, {true, false, false}, {false, true, false}, {true, true, true}}

	for _, values := range truth_table {
		var result bool = Xnor_gate(values[0], values[1])

		if result != values[2] {
			t.Fatalf(fmt.Sprintf("xnor_gate(%t, %t) should be %t, but got %t", values[0], values[1], values[2], result))
		}
	}
}
