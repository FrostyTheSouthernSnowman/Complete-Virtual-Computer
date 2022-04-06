package bus

import (
	"math/rand"
	"testing"
)

var random_bus [64]bool

var test_bus [64]bool

func subscriber(current_bus [64]bool) {
	test_bus = current_bus
}

func TestBus(t *testing.T) {
	for i := 1; i < len(random_bus); i++ {
		random_bus[i-1] = rand.Float32() < 0.5
	}

	subscribe_to_bus(subscriber)
	load(random_bus)
	if test_bus != random_bus || test_bus != bus || bus != random_bus {
		t.Fatalf("Bus mismatch\ntest_bus: %v\nrandom_bus: %v\nmain bus: %v", test_bus, random_bus, bus)
	}
}
