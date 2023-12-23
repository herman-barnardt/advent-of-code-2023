package day22

import (
	"fmt"
	"slices"
	"sort"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 22, solve2023Day22Part1, solve2023Day22Part2)
}

type brick struct {
	name   string
	startX int
	startY int
	startZ int
	endX   int
	endY   int
	endZ   int
}

func solve2023Day22Part1(lines []string) interface{} {
	layout := map[int]map[int]map[int]string{}
	bricks := []*brick{}
	for i, line := range lines {
		var startX, startY, startZ, endX, endY, endZ int
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &startX, &startY, &startZ, &endX, &endY, &endZ)
		name := string(rune(i + 65))
		bricks = append(bricks, &brick{name, startX, startY, startZ, endX, endY, endZ})
	}
	sort.Slice(bricks, func(i, j int) bool {
		return util.IntMin(bricks[i].startZ, bricks[i].endZ) < util.IntMin(bricks[j].startZ, bricks[j].endZ)
	})
	for _, b := range bricks {
		canMove := true
		for canMove {
			for x := b.startX; x <= b.endX; x++ {
				for y := b.startY; y <= b.endY; y++ {
					for z := b.startZ; z <= b.endZ; z++ {
						if c, ok := layout[x][y][z-1]; z-1 <= 0 || (ok && c != b.name) {
							canMove = false
						}
					}
				}
			}
			if canMove {
				b.startZ = b.startZ - 1
				b.endZ = b.endZ - 1
			}
		}
		for x := b.startX; x <= b.endX; x++ {
			if _, ok := layout[x]; !ok {
				layout[x] = map[int]map[int]string{}
			}
			for y := b.startY; y <= b.endY; y++ {
				if _, ok := layout[x][y]; !ok {
					layout[x][y] = make(map[int]string)
				}
				for z := b.startZ; z <= b.endZ; z++ {
					layout[x][y][z] = b.name
				}
			}
		}
	}

	supportedBy := make(map[string]map[string]bool)
	supports := make(map[string]map[string]bool)
	for _, b := range bricks {
		if _, ok := supportedBy[b.name]; !ok {
			supportedBy[b.name] = make(map[string]bool)
		}
		if _, ok := supports[b.name]; !ok {
			supports[b.name] = make(map[string]bool)
		}
		for x := b.startX; x <= b.endX; x++ {
			for y := b.startY; y <= b.endY; y++ {
				for z := b.startZ; z <= b.endZ; z++ {
					if c, ok := layout[x][y][z-1]; ok && c != b.name {
						if len(c) > 0 {
							supportedBy[b.name][c] = true
						}
					}
					if c, ok := layout[x][y][z+1]; ok && c != b.name {
						if len(c) > 0 {
							supports[b.name][c] = true
						}
					}
				}
			}
		}
	}
	count := 0
	for _, b := range bricks {
		supportedHasExtraSupport := true
		for supported := range supports[b.name] {
			supportedHasExtraSupport = supportedHasExtraSupport && len(supportedBy[supported]) > 1
		}
		if len(supports[b.name]) == 0 || supportedHasExtraSupport {
			count++
		}
	}
	return count
}

func solve2023Day22Part2(lines []string) interface{} {
	layout := map[int]map[int]map[int]string{}
	bricks := []*brick{}
	for i, line := range lines {
		var startX, startY, startZ, endX, endY, endZ int
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &startX, &startY, &startZ, &endX, &endY, &endZ)
		name := string(rune(i + 65))
		bricks = append(bricks, &brick{name, startX, startY, startZ, endX, endY, endZ})
	}
	sort.Slice(bricks, func(i, j int) bool {
		return util.IntMin(bricks[i].startZ, bricks[i].endZ) < util.IntMin(bricks[j].startZ, bricks[j].endZ)
	})
	for _, b := range bricks {
		canMove := true
		for canMove {
			for x := b.startX; x <= b.endX; x++ {
				for y := b.startY; y <= b.endY; y++ {
					for z := b.startZ; z <= b.endZ; z++ {
						if c, ok := layout[x][y][z-1]; z-1 <= 0 || (ok && c != b.name) {
							canMove = false
						}
					}
				}
			}
			if canMove {
				b.startZ = b.startZ - 1
				b.endZ = b.endZ - 1
			}
		}
		for x := b.startX; x <= b.endX; x++ {
			if _, ok := layout[x]; !ok {
				layout[x] = map[int]map[int]string{}
			}
			for y := b.startY; y <= b.endY; y++ {
				if _, ok := layout[x][y]; !ok {
					layout[x][y] = make(map[int]string)
				}
				for z := b.startZ; z <= b.endZ; z++ {
					layout[x][y][z] = b.name
				}
			}
		}
	}

	supportedBy := make(map[string]map[string]bool)
	supports := make(map[string]map[string]bool)
	for _, b := range bricks {
		if _, ok := supportedBy[b.name]; !ok {
			supportedBy[b.name] = make(map[string]bool)
		}
		if _, ok := supports[b.name]; !ok {
			supports[b.name] = make(map[string]bool)
		}
		for x := b.startX; x <= b.endX; x++ {
			for y := b.startY; y <= b.endY; y++ {
				for z := b.startZ; z <= b.endZ; z++ {
					if c, ok := layout[x][y][z-1]; ok && c != b.name {
						if len(c) > 0 {
							supportedBy[b.name][c] = true
						}
					}
					if c, ok := layout[x][y][z+1]; ok && c != b.name {
						if len(c) > 0 {
							supports[b.name][c] = true
						}
					}
				}
			}
		}
	}

	count := 0
	for _, b := range bricks {
		supportedHasExtraSupport := true
		for supported := range supports[b.name] {
			supportedHasExtraSupport = supportedHasExtraSupport && len(supportedBy[supported]) > 1
		}
		if len(supports[b.name]) == 0 || supportedHasExtraSupport {
			continue
		}
		toDisintegrate := []string{b.name}
		fallen := map[string]bool{}

		for len(toDisintegrate) > 0 {
			current := toDisintegrate[0]
			toDisintegrate = toDisintegrate[1:]

			willAlsoFallList := []string{}
			for supportedByCurrent := range supports[current] {
				willAlsoFall := true
				for x := range supportedBy[supportedByCurrent] {
					_, inFallen := fallen[x]
					willAlsoFall = willAlsoFall && (inFallen || x == current || slices.Contains(toDisintegrate, x))
				}
				if willAlsoFall {
					willAlsoFallList = append(willAlsoFallList, supportedByCurrent)
				}
			}
			fallen[current] = true
			for _, x := range willAlsoFallList {
				if _, ok := fallen[x]; !ok {
					toDisintegrate = append(toDisintegrate, x)
				}
			}
		}
		count += len(fallen) - 1
	}
	return count
}
