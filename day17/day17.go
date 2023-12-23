package day17

import (
	"math"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/graph"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 17, solve2023Day17Part1, solve2023Day17Part2)
}

var grid = make(map[util.Point]int)

var directions = map[string]util.Point{
	">": {X: 1, Y: 0},
	"<": {X: -1, Y: 0},
	"v": {X: 0, Y: 1},
	"^": {X: 0, Y: -1},
}

var minMoves, maxMoves int = 0, 3

type crucibleNode struct {
	point         util.Point
	value         int
	streak        int
	lastDirection string
}

func (n crucibleNode) GetNeighbours() []graph.Node {
	retVal := make([]graph.Node, 0)
	for direction, dirPoint := range directions {
		newPoint := n.point.Add(dirPoint)
		streak := 1
		value := 0
		if n.lastDirection == direction {
			streak = n.streak + 1
		}

		if v, ok := grid[newPoint]; ok {
			value = v
		} else {
			continue
		}
		if streak > maxMoves {
			continue
		}
		if (n.lastDirection == ">" && direction == "<") || (n.lastDirection == "<" && direction == ">") ||
			(n.lastDirection == "v" && direction == "^") || (n.lastDirection == "^" && direction == "v") {
			continue
		}
		if direction != n.lastDirection && n.streak < minMoves {
			continue
		}
		retVal = append(retVal, crucibleNode{newPoint, value, streak, direction})
	}
	return retVal
}
func (n crucibleNode) GetCost(to graph.Node) float64 {
	return float64(to.(crucibleNode).value)
}
func (n crucibleNode) GetHeuristicCost(to graph.Node) float64 {
	toPoint := to.(*crucibleNode).point
	return float64(util.DistanceBetween(&n.point, &toPoint))
}
func (n crucibleNode) Equal(to graph.Node) bool {
	toPoint := to.(*crucibleNode)
	return n.point.X == toPoint.point.X && n.point.Y == toPoint.point.Y
}

func solve2023Day17Part1(lines []string) interface{} {
	for y, row := range util.LinesToMapofInts(lines) {
		for x, v := range row {
			grid[util.Point{X: x, Y: y}] = v
		}
	}

	end := &crucibleNode{util.Point{X: len(lines[0]) - 1, Y: len(lines) - 1}, 0, 0, ""}

	_, costDown, _ := graph.FindShortestPath(&crucibleNode{util.Point{X: 0, Y: 0}, 0, 1, "v"}, end)
	_, costRight, _ := graph.FindShortestPath(&crucibleNode{util.Point{X: 0, Y: 0}, 0, 1, ">"}, end)
	return math.Min(costDown, costRight)
}

//THERE IS A BUG SOMEWHERE THAT IS CAUSING THE STARTING DOWN ON PART 2 TO HAVE A LOWER VALUE THAT IS NOT CORRECT

func solve2023Day17Part2(lines []string) interface{} {
	minMoves = 4
	maxMoves = 10
	for y, row := range util.LinesToMapofInts(lines) {
		for x, v := range row {
			grid[util.Point{X: x, Y: y}] = v
		}
	}
	end := &crucibleNode{util.Point{X: len(lines[0]) - 1, Y: len(lines) - 1}, 0, 0, ""}

	_, costDown, _ := graph.FindShortestPath(&crucibleNode{util.Point{X: 0, Y: 0}, 0, 1, "v"}, end)
	_, costRight, _ := graph.FindShortestPath(&crucibleNode{util.Point{X: 0, Y: 0}, 0, 1, ">"}, end)
	return math.Min(costDown, costRight)
}
