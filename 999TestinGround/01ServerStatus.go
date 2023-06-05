package main

import "fmt"

const (
	Offline uint8 = iota
	Online
	Maintenance
	Retired
)

type serverDetails struct {
	Name   string
	Info   string
	States uint8
}

func main() {
	var AllServers = make(map[uint16]serverDetails, 10)
	AllServers[0] = serverDetails{"ProductionDB", "Serving in US", Offline}
	AllServers[1] = serverDetails{"ProductionDB", "Serving in UK", Offline}
	AllServers[2] = serverDetails{"ProductionDB", "Serving in IN", Offline}
	AllServers[3] = serverDetails{"ProductionDB", "Serving in IR", Offline}
	AllServers[4] = serverDetails{"ProductionDB", "Serving in EU", Offline}
	AllServers[5] = serverDetails{"ProductionDB", "Serving in CN", Offline}
	AllServers[6] = serverDetails{"ProductionDB", "Serving in JP", Offline}
	AllServers[7] = serverDetails{"ProductionDB", "Serving in RU", Offline}
	AllServers[8] = serverDetails{"ProductionDB", "Serving in CA", Offline}
	AllServers[9] = serverDetails{"ProductionDB", "Serving in AU", Offline}

	for key, val := range AllServers {
		fmt.Printf("%v: %v\n", key, val)
	}
}
