package register

import (
	"Complete-Virtual-Computer/bus"
	"math/rand"
	"testing"
)

var random_register_data [16]bool

var test_bus [16]bool

func subscriber(current_bus [16]bool) {
	test_bus = current_bus
}

func TestRegister(t *testing.T) {
	for i := 1; i < len(random_register_data); i++ {
		random_register_data[i-1] = rand.Float32() < 0.5
	}

	bus.Subscribe_to_bus(subscriber)
	bus.Load(random_register_data)

	register_data := Register(test_bus, random_register_data, false, false)

	if register_data != random_register_data {
		t.Fatalf("Data mismatch between register_data and random_register_data")
	}

	test_bus[10] = !test_bus[10]
	test_bus[3] = !test_bus[3]
	test_bus[13] = !test_bus[13]

	bus.Load(test_bus)

	register_data = Register(test_bus, register_data, true, false)

	if register_data != test_bus {
		t.Fatalf("Register_time_step is not loading bus data correctly.")
	}

	register_data[10] = !register_data[10]
	register_data[3] = !register_data[3]
	register_data[13] = !register_data[13]

	register_data = Register(test_bus, register_data, false, true)

	if register_data != test_bus {
		t.Fatalf("Register_time_step is not outputing data to the bus correctly.")
	}
}
