package alu

import (
	logic_gates "Complete-Virtual-Computer/logic-gates"
)

func Sixteen_bit_adder(n1 [16]bool, n2 [16]bool, subtract bool) [16]bool {
	var addition bool
	var carry bool = subtract
	var result [16]bool

	for index, bit := range n2 {
		n2[index] = logic_gates.Xor_gate(bit, subtract)
	}

	for i := 1; i <= 16; i++ {
		// We allready used the first bit, so start on the second bit, or index 1
		var first_xor bool = logic_gates.Xor_gate(n1[16-i], n2[16-i])
		var first_and bool = logic_gates.And_gate(n1[16-i], n2[16-i])

		addition = logic_gates.Xor_gate(first_xor, carry)

		var second_and bool = logic_gates.And_gate(first_xor, carry)
		carry = logic_gates.Or_gate(first_and, second_and)

		result[16-i] = addition
	}

	return result
}
