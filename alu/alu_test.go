package alu

import (
	"testing"
)

func TestSixteenBitAdder(t *testing.T) {
	number1 := [16]bool{true, true, true, false, false, true, false, false, true, true, false, false, true, true, false, true}
	number2 := [16]bool{true, false, true, false, true, true, true, false, false, true, true, true, false, true, false, false}
	expected_addition_result := [16]bool{true, false, false, true, false, false, true, true, false, true, false, false, false, false, false, true}
	expected_subtraction_result := [16]bool{false, false, true, true, false, true, true, false, false, true, false, true, true, false, false, true}

	result := Sixteen_bit_adder(number1, number2, false)

	if result != expected_addition_result {
		t.Fatalf("16-bit adder result mismatch for addition. results in result, expected result:\n%v\n%v\n", result, expected_addition_result)
	}

	result = Sixteen_bit_adder(number1, number2, true)

	if result != expected_subtraction_result {
		t.Fatalf("16-bit adder result mismatch for subtraction. results in result, expected result:\n%v\n%v", result, expected_subtraction_result)
	}
}
