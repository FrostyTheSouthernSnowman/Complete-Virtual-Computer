package bus

var bus [64]bool
var connections []func([64]bool)

func subscribe_to_bus(on_change func([64]bool)) {
	connections = append(connections, on_change)
}

func broadcast_bus() {
	for _, connection := range connections {
		connection(bus) // The connected unit can use the bus data, or silently reject it
	}
}

func load(data [64]bool) {
	bus = data
	broadcast_bus()
}
