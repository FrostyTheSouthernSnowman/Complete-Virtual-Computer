package logic_gates

func transistor(input1 bool, input2 bool) bool {
	if input1 && input2 {
		return true
	} else {
		return false
	}
}
