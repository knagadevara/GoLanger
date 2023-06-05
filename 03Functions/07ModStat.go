package main

import "fmt"

type setStat interface {
	UpdateStat()
}

type Health struct {
	Max, current uint16
}

type Energy struct {
	Max, current uint16
}

type PlayerV struct {
	id     uint16
	name   string
	health *Health
	energy *Energy
}

const MAX uint16 = 100
const HE_P20 uint16 = 20
const HE_P50 uint16 = 50
const HE_P75 uint16 = 75

func (h *Health) UpdateStat(hPortion uint16) {
	println("Updating Health")
	totalNeededHealth := (h.Max - h.current)
	if (h.current != 0) && (h.current < h.Max) {
		if totalNeededHealth <= hPortion {
			h.current = h.current + (hPortion - (hPortion - totalNeededHealth))
		} else if totalNeededHealth >= hPortion {
			h.current = h.current + hPortion
		} else {
			fmt.Printf("Update Not Required")
		}
	}
}

func (e *Energy) UpdateStat(ePortion uint16) {
	println("Updating Energy")
	updatedEnergy := (e.current + ePortion)
	if (e.current != 0) && (e.current < e.Max) {
		if updatedEnergy < e.Max {
			e.current = updatedEnergy
		} else if updatedEnergy >= e.Max {
			e.current = e.Max
		} else {
			fmt.Printf("Update Not Required")
		}
	}
}

func (p *PlayerV) UpdateStat(pId uint16, pName string) {
	p.id = pId
	p.name = pName
	fmt.Printf("Updated Player Name!\n")
}

func (p *PlayerV) printStat() {
	fmt.Printf("Name:\t%v\nHealth:\t%v\nEnergy:\t%v\n",
		p.name,
		p.health.current,
		p.energy.current)
}

func UpdateStat(sStat setStat) {
	sStat.UpdateStat()
}

func main() {

	var AdhiHealth Health
	AdhiHealth.current = 50
	AdhiHealth.Max = MAX
	var AdhiEnergy Energy
	AdhiEnergy.current = 50
	AdhiEnergy.Max = MAX
	var Adhi PlayerV
	Adhi.id, Adhi.name, Adhi.health, Adhi.energy = 1111, "Adhi", &AdhiHealth, &AdhiEnergy
	AdhiEnergy.UpdateStat(120)
	AdhiHealth.UpdateStat(HE_P75)
	Adhi.printStat()
	Adhi.UpdateStat(007, "Gaurav")
	Adhi.printStat()

}
