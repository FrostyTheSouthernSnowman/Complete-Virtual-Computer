package register

import (
	"Complete-Virtual-Computer/bus"
)

func Register(main_bus [16]bool, register_data [16]bool, load bool, output bool) [16]bool {
	if load {
		register_data = main_bus
	}
	if output {
		bus.Load(register_data)
	}

	return register_data
}
