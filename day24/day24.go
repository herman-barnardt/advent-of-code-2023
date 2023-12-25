package day24

import (
	"fmt"
	"math"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2023, 24, solve2023Day24Part1, solve2023Day24Part2)
}

type Point3D struct {
	X, Y, Z float64
}

func (p Point3D) Add(q Point3D) Point3D {
	return Point3D{X: p.X + q.X, Y: p.Y + q.Y, Z: p.Z + q.Z}
}

func (p Point3D) Minus(q Point3D) Point3D {
	return Point3D{X: p.X - q.X, Y: p.Y - q.Y, Z: p.Z - q.Z}
}

type hailstone struct {
	index     int
	initial   Point3D
	direction Point3D
	slope     float64
	offset    float64
}

func solve2023Day24Part1(lines []string) interface{} {
	hailstones := make([]hailstone, 0)
	for i, line := range lines {
		newHailstone := hailstone{index: i, slope: math.MaxInt}
		fmt.Sscanf(line, "%f, %f, %f @ %f, %f, %f", &newHailstone.initial.X, &newHailstone.initial.Y, &newHailstone.initial.Z, &newHailstone.direction.X, &newHailstone.direction.Y, &newHailstone.direction.Z)
		if newHailstone.direction.X != 0 {
			newHailstone.slope = float64(newHailstone.direction.Y) / float64(newHailstone.direction.X)
		}
		newHailstone.offset = float64(newHailstone.initial.Y) - newHailstone.slope*float64(newHailstone.initial.X)
		hailstones = append(hailstones, newHailstone)
	}

	// min := 7.0
	// max := 27.0
	min := 200000000000000.0
	max := 400000000000000.0
	count := 0

	for i, h1 := range hailstones {
		for _, h2 := range hailstones[i+1:] {
			x := (h1.offset - h2.offset) / (h2.slope - h1.slope)
			y := h1.slope*x + h1.offset
			if (h1.direction.X < 0 && x > h1.initial.X) ||
				(h1.direction.X > 0 && x < h1.initial.X) ||
				(h1.direction.Y < 0 && y > h1.initial.Y) ||
				(h1.direction.Y > 0 && y < h1.initial.Y) ||
				(h2.direction.X < 0 && x > h2.initial.X) ||
				(h2.direction.X > 0 && x < h2.initial.X) ||
				(h2.direction.Y < 0 && y > h2.initial.Y) ||
				(h2.direction.Y > 0 && y < h2.initial.Y) {
				continue
			}
			if x >= min && x <= max && y >= min && y <= max {
				count++
			}
		}
	}

	return count
}

func solve2023Day24Part2(lines []string) interface{} {
	hailstones := make([]hailstone, 0)
	for i, line := range lines {
		newHailstone := hailstone{index: i, slope: math.MaxInt}
		fmt.Sscanf(line, "%f, %f, %f @ %f, %f, %f", &newHailstone.initial.X, &newHailstone.initial.Y, &newHailstone.initial.Z, &newHailstone.direction.X, &newHailstone.direction.Y, &newHailstone.direction.Z)
		if newHailstone.direction.X != 0 {
			newHailstone.slope = float64(newHailstone.direction.Y) / float64(newHailstone.direction.X)
		}
		newHailstone.offset = float64(newHailstone.initial.Y) - newHailstone.slope*float64(newHailstone.initial.X)
		hailstones = append(hailstones, newHailstone)
	}

	potentialX := make(map[float64]bool)
	potentialY := make(map[float64]bool)
	potentialZ := make(map[float64]bool)
	for i, h1 := range hailstones {
		for _, h2 := range hailstones[i+1:] {
			if h1.direction.X == h2.direction.X {
				newX := make([]float64, 0)
				deltaX := int(h2.initial.X - h1.initial.X)
				for v := -1000; v <= 1000; v++ {
					if float64(v) == h1.direction.X {
						continue
					}
					if deltaX%(v-int(h1.direction.X)) == 0 {
						newX = append(newX, float64(v))
					}
				}
				if len(potentialX) == 0 {
					potentialX = make(map[float64]bool)
					for _, x := range newX {
						potentialX[x] = true
					}
				} else {
					intersection := make([]float64, 0)
					for _, x := range newX {
						if _, ok := potentialX[x]; ok {
							intersection = append(intersection, x)
						}
					}
					potentialX = make(map[float64]bool)
					for _, x := range intersection {
						potentialX[x] = true
					}
				}
			}
			if h1.direction.Y == h2.direction.Y {
				newY := make([]float64, 0)
				deltaY := int(h2.initial.Y - h1.initial.Y)
				for v := -1000; v <= 1000; v++ {
					if float64(v) == h1.direction.Y {
						continue
					}
					if deltaY%(v-int(h1.direction.Y)) == 0 {
						newY = append(newY, float64(v))
					}
				}
				if len(potentialY) == 0 {
					potentialY = make(map[float64]bool)
					for _, y := range newY {
						potentialY[y] = true
					}
				} else {
					intersection := make([]float64, 0)
					for _, y := range newY {
						if _, ok := potentialY[y]; ok {
							intersection = append(intersection, y)
						}
					}
					potentialY = make(map[float64]bool)
					for _, y := range intersection {
						potentialY[y] = true
					}
				}
			}
			if h1.direction.Z == h2.direction.Z {
				newZ := make([]float64, 0)
				deltaZ := int(h2.initial.Z - h1.initial.Z)
				for v := -1000; v <= 1000; v++ {
					if float64(v) == h1.direction.Z {
						continue
					}
					if deltaZ%(v-int(h1.direction.Z)) == 0 {
						newZ = append(newZ, float64(v))
					}
				}
				if len(potentialZ) == 0 {
					potentialZ = make(map[float64]bool)
					for _, z := range newZ {
						potentialZ[z] = true
					}
				} else {
					intersection := make([]float64, 0)
					for _, z := range newZ {
						if _, ok := potentialZ[z]; ok {
							intersection = append(intersection, z)
						}
					}
					potentialZ = make(map[float64]bool)
					for _, z := range intersection {
						potentialZ[z] = true
					}
				}
			}
		}
	}

	rockX := 0.0
	for k := range potentialX {
		rockX = k
		break
	}
	rockY := 0.0
	for k := range potentialY {
		rockY = k
		break
	}
	rockZ := 0.0
	for k := range potentialZ {
		rockZ = k
		break
	}

	h1 := hailstones[0]
	h2 := hailstones[1]
	slope1 := (h1.direction.Y - rockY) / (h1.direction.X - rockX)
	slope2 := (h2.direction.Y - rockY) / (h2.direction.X - rockX)
	offset1 := h1.initial.Y - (slope1 * h1.initial.X)
	offset2 := h2.initial.Y - (slope2 * h2.initial.X)
	x := (offset2 - offset1) / (slope1 - slope2)
	y := (slope1*x + offset1)
	time := (x - h1.initial.X) / (h1.direction.X - rockX)
	z := h1.initial.Z + (h1.direction.Z-rockZ)*time
	return int(x + y + z)
}
