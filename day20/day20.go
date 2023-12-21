package day20

import (
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 20, solve2023Day20Part1, solve2023Day20Part2)
}

type moduleType struct {
	name         string
	typ          string
	destinations []string
	currentState bool
	lastSignals  map[string]bool
}

func (m *moduleType) receivePulse(pulse bool, sender string) []pulseType {
	if m == nil {
		return []pulseType{}
	}
	newPulse := pulse
	newReceivers := make([]pulseType, 0)
	switch m.typ {
	case "%":
		if !pulse {
			m.currentState = !m.currentState
			newPulse = m.currentState
		} else {
			return []pulseType{}
		}
	case "&":
		m.lastSignals[sender] = pulse
		for _, dest := range m.lastSignals {
			newPulse = newPulse && dest
		}
		newPulse = !newPulse
	}
	for _, destination := range m.destinations {
		newReceivers = append(newReceivers, pulseType{destination, m.name, newPulse})
	}
	return newReceivers
}

type pulseType struct {
	destination string
	sender      string
	signal      bool
}

var modules = map[string]*moduleType{}

func solve2023Day20Part1(lines []string) interface{} {
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		typ := ""
		name := parts[0]
		if parts[0][0] == '%' || parts[0][0] == '&' {
			typ = string(parts[0][0])
			name = name[1:]
		}
		destinations := strings.Split(parts[1], ", ")
		modules[name] = &moduleType{name, typ, destinations, false, map[string]bool{}}
	}

	for _, module := range modules {
		for _, destination := range module.destinations {
			if d, ok := modules[destination]; ok && d.typ == "&" {
				modules[destination].lastSignals[module.name] = false
			}
		}
	}

	highCount, lowCount := 0, 0

	for i := 0; i < 1000; i++ {
		pulses := []pulseType{{"broadcaster", "button", false}}
		for len(pulses) > 0 {
			currentPulse := pulses[0]
			if currentPulse.signal {
				highCount++
			} else {
				lowCount++
			}
			pulses = pulses[1:]
			newPulses := modules[currentPulse.destination].receivePulse(currentPulse.signal, currentPulse.sender)
			pulses = append(pulses, newPulses...)
		}
	}

	return highCount * lowCount
}

func solve2023Day20Part2(lines []string) interface{} {
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		typ := ""
		name := parts[0]
		if parts[0][0] == '%' || parts[0][0] == '&' {
			typ = string(parts[0][0])
			name = name[1:]
		}
		destinations := strings.Split(parts[1], ", ")
		modules[name] = &moduleType{name, typ, destinations, false, map[string]bool{}}
	}

	rxSource := ""
	for _, module := range modules {
		for _, destination := range module.destinations {
			if d, ok := modules[destination]; ok && d.typ == "&" {
				modules[destination].lastSignals[module.name] = false
			}
			if destination == "rx" {
				rxSource = module.name
			}
		}
	}

	highCount, lowCount := 0, 0
	buttonPresses := 0

	rxSourcesMap := map[string]int{}

	for len(rxSourcesMap) < 4 {
		pulses := []pulseType{{"broadcaster", "button", false}}
		buttonPresses++
		for len(pulses) > 0 {
			currentPulse := pulses[0]
			if currentPulse.signal {
				highCount++
			} else {
				lowCount++
			}
			pulses = pulses[1:]
			newPulses := modules[currentPulse.destination].receivePulse(currentPulse.signal, currentPulse.sender)
			pulses = append(pulses, newPulses...)
			for name, lastSignal := range modules[rxSource].lastSignals {
				if _, ok := rxSourcesMap[name]; !ok && lastSignal {
					rxSourcesMap[name] = buttonPresses
				}
			}
		}
	}

	counts := []int{}
	for _, count := range rxSourcesMap {
		counts = append(counts, count)
	}
	return util.LeastCommonMultiple(counts[0], counts[1], counts...)
}
