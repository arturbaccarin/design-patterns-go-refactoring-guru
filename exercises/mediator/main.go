package main

import (
	"fmt"
)

type Mediator interface {
	Notify(sender Aircraft, action string)
}

type Aircraft interface {
	PerformAction(action string)

	GetName() string
}

type FlightControlTower struct {
	airplanes []Aircraft
}

func NewFlightControlTower() *FlightControlTower {
	return &FlightControlTower{}
}

func (tower *FlightControlTower) AddAircraft(aircraft Aircraft) {
	tower.airplanes = append(tower.airplanes, aircraft)
}

func (tower *FlightControlTower) Notify(sender Aircraft, action string) {
	for _, airplane := range tower.airplanes {

		if airplane != sender {
			airplane.PerformAction(fmt.Sprintf("%s (by %s)", action, sender.GetName()))
		}
	}
}

type Airplane struct {
	name     string
	tower    Mediator
	altitude int
}

func NewAirplane(name string, tower Mediator) *Airplane {
	return &Airplane{
		name:     name,
		tower:    tower,
		altitude: 0, // Start at ground level
	}
}

func (a *Airplane) PerformAction(action string) {
	fmt.Printf("Airplane %s is notified: %s\n", a.name, action)
}

func (a *Airplane) GetName() string {
	return a.name
}

func (a *Airplane) RequestAltitudeChange(newAltitude int) {
	if newAltitude != a.altitude {
		fmt.Printf("%s is requesting an altitude change to %d.\n", a.name, newAltitude)
		a.altitude = newAltitude
		a.tower.Notify(a, fmt.Sprintf("%s is changing altitude to %d", a.name, newAltitude))
	} else {
		fmt.Printf("%s is already at altitude %d.\n", a.name, newAltitude)
	}
}

func main() {
	tower := NewFlightControlTower()

	plane1 := NewAirplane("Plane1", tower)
	plane2 := NewAirplane("Plane2", tower)
	plane3 := NewAirplane("Plane3", tower)

	tower.AddAircraft(plane1)
	tower.AddAircraft(plane2)
	tower.AddAircraft(plane3)

	plane1.RequestAltitudeChange(10000)
	plane2.RequestAltitudeChange(15000)
	plane3.RequestAltitudeChange(12000)
}
