package bus

var main_bus [16]bool
var connections []func([16]bool)

func Subscribe_to_bus(on_change func([16]bool)) {
	connections = append(connections, on_change)
}

func Broadcast_bus() {
	for _, connection := range connections {
		connection(main_bus) // The connected unit can use the bus data, or silently reject it
	}
}

func Load(data [16]bool) {
	main_bus = data
	Broadcast_bus()
}
