package logic_gates

func and_gate(input1 bool, input2 bool) bool {
	var transistor_1 bool = transistor(input1, true)
	var transistor_2 bool = transistor(input2, transistor_1)
	return transistor_2
}

func not_gate(input bool) bool {
	var transistor_output bool = transistor(input, true)
	if transistor_output {
		return false
	} else {
		return true
	}
}

func or_gate(input1 bool, intput2 bool) bool {
	var transistor1 bool = transistor(input1, true)
	var transistor2 bool = transistor(intput2, true)
	if transistor1 {
		return true
	} else if transistor2 {
		return true
	} else {
		return false
	}
}

func xor_gate(input1 bool, input2 bool) bool {
	var after_or = or_gate(input1, input2)
	var after_and = and_gate(input1, input2)
	var after_not_and = not_gate(after_and)
	var after_second_and = and_gate(after_not_and, after_or)
	return after_second_and
}

func nand_gate(input1 bool, input2 bool) bool {
	return not_gate(and_gate(input1, input2))
}

func nor_gate(input1 bool, input2 bool) bool {
	return not_gate(or_gate(input1, input2))
}

func xnor_gate(input1 bool, input2 bool) bool {
	return not_gate(xor_gate(input1, input2))
}
