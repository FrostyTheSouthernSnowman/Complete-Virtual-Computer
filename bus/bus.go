package bus

var bus [16]bool
var connections []func([16]bool)

func subscribe_to_bus(on_change func([16]bool)) {
	connections = append(connections, on_change)
}

func broadcast_bus() {
	for _, connection := range connections {
		connection(bus) // The connected unit can use the bus data, or silently reject it
	}
}

func load(data [16]bool) {
	bus = data
	broadcast_bus()
}
