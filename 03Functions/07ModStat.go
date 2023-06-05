package main

import "fmt"

const MAX uint16 = 100
const h20 uint16 = 20
const h50 uint16 = 50
const h75 uint16 = 75

type Health struct {
	current uint16
	Max     uint16
}

type Energy struct {
	current uint16
	Max     uint16
}

type PlayerV struct {
	id     uint16
	name   string
	health *Health
	energy *Energy
}

func (h *Health) setHealth(hPortion uint16) {
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

func (e *Energy) setEnergy(ePortion uint16) {
	totalNeededEnergy := (e.Max - e.current)
	if (e.current != 0) && (e.current < e.Max) {
		if totalNeededEnergy <= ePortion {
			e.current = e.current + (ePortion - (ePortion - totalNeededEnergy))
		} else if totalNeededEnergy >= ePortion {
			e.current = e.current + ePortion
		} else {
			fmt.Printf("Update Not Required")
		}
	}
}

func (p *PlayerV) changePlayerDetails(pId uint16, pName string) {
	p.id = pId
	p.name = pName
	fmt.Printf("\tGame ON!\nName:\t%v", p.name)
}

func (p *PlayerV) printStats() {
	fmt.Printf("Name:\t%v\nHealth:\t%v\nEnergy:\t%v\n",
		p.name,
		p.health.current,
		p.energy.current)
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
	Adhi.health.setHealth(120)
	Adhi.printStats()
}
