package day16

type Valve struct {
	name        string
	flowrate    int
	connections []*Valve
	open        bool
	release     int
	distances   map[string]int
}

func NewValve(name string, flowrate int) *Valve {
	valve := &Valve{
		name:        name,
		flowrate:    flowrate,
		open:        false,
		connections: make([]*Valve, 0),
		distances:   make(map[string]int, 0),
	}
	return valve
}

func (v *Valve) addConnection(valve *Valve) {
	v.connections = append(v.connections, valve)
}

func (v *Valve) openValve() {
	v.open = true
	time--
}

//func (v *Valve) getHighestPotentialReleaseValve() {
//	maxRelease := 0
//	potValves := v.connections
//	targetValves := make([]*Valve, 0)
//
//	for _, valve := range valves {
//		if valve.flowrate > 0 {
//			targetValves = append(targetValves, valve)
//		}
//	}
//	for _, targetValve := range targetValves {
//		v.getDistanceToTarget(targetValve)
//	}
//	return
//}

func (v *Valve) getDistances() {
	for _, connection := range valves {
		if connection.name != v.name {
			v.distances[connection.name] = v.getDistanceToTarget(connection)
		}
	}
}

func (v *Valve) getDistanceToTarget(target *Valve) int {
	steps := 0
	potentialTargets := v.connections
	match := false
	for {
		for _, connection := range potentialTargets {
			if connection == target {
				match = true
			}
		}
		if match {
			steps++
			break
		}
		steps++
		newTargets := make([]*Valve, 0)
		for _, potentialTarget := range potentialTargets {
			for _, valve := range potentialTarget.connections {
				newTargets = append(newTargets, valve)
			}
		}
		potentialTargets = newTargets
	}
	return steps
}
