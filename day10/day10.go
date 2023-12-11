package day10

import (
	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/graph"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 10, solve2023Day10Part1, solve2023Day10Part2)
}

type pointNode struct {
	x          int
	y          int
	value      string
	neighbours []*pointNode
}

func (p *pointNode) GetNeighbours() []graph.Node {
	nodeNeighbours := make([]graph.Node, len(p.neighbours))
	for i, neighbour := range p.neighbours {
		nodeNeighbours[i] = neighbour
	}
	return nodeNeighbours
}
func (p *pointNode) GetCost(to graph.Node) float64 {
	return 1
}

func (p *pointNode) GetHeuristicCost(to graph.Node) float64 {
	return 1
}

func (p *pointNode) Equal(to graph.Node) bool {
	toPoint := to.(*pointNode)
	return p.x == toPoint.x && p.y == toPoint.y
}

func solve2023Day10Part1(lines []string) interface{} {
	lineMap := util.LinesToMap(lines)
	world := make(map[int]map[int]*pointNode)
	var start *pointNode
	for y := range lineMap {
		world[y] = make(map[int]*pointNode)
		for x := range lineMap[y] {
			world[y][x] = &pointNode{x, y, lineMap[y][x], make([]*pointNode, 0)}
			if lineMap[y][x] == "S" {
				start = world[y][x]
			}
		}
	}
	for y := range world {
		for x, node := range world[y] {
			switch world[y][x].value {
			// | is a vertical pipe connecting north and south.
			case "|":
				{
					if neighbour, ok := world[y-1][x]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y-1][x])
					}
					if neighbour, ok := world[y+1][x]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y+1][x])
					}
				}
			// - is a horizontal pipe connecting east and west.
			case "-":
				{
					if neighbour, ok := world[y][x-1]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y][x-1])
					}
					if neighbour, ok := world[y][x+1]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y][x+1])
					}
				}
			// L is a 90-degree bend connecting north and east.
			case "L":
				{
					{
						if neighbour, ok := world[y-1][x]; ok && neighbour.value != "." && neighbour.value != "S" {
							node.neighbours = append(node.neighbours, world[y-1][x])
						}
						if neighbour, ok := world[y][x+1]; ok && neighbour.value != "." && neighbour.value != "S" {
							node.neighbours = append(node.neighbours, world[y][x+1])
						}
					}
				}
			// J is a 90-degree bend connecting north and west.
			case "J":
				{
					if neighbour, ok := world[y-1][x]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y-1][x])
					}
					if neighbour, ok := world[y][x-1]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y][x-1])
					}
				}
			// 7 is a 90-degree bend connecting south and west.
			case "7":
				{
					if neighbour, ok := world[y+1][x]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y+1][x])
					}
					if neighbour, ok := world[y][x-1]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y][x-1])
					}
				}
			// F is a 90-degree bend connecting south and east.
			case "F":
				{
					if neighbour, ok := world[y+1][x]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y+1][x])
					}
					if neighbour, ok := world[y][x+1]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y][x+1])
					}
				}
			}
		}
	}

	if neighbour, ok := world[start.y-1][start.x]; ok && (neighbour.value == "|" || neighbour.value == "7" || neighbour.value == "F") {
		start.neighbours = append(start.neighbours, neighbour)
	}
	if neighbour, ok := world[start.y+1][start.x]; ok && (neighbour.value == "|" || neighbour.value == "L" || neighbour.value == "J") {
		start.neighbours = append(start.neighbours, neighbour)
	}
	if neighbour, ok := world[start.y][start.x-1]; ok && (neighbour.value == "-" || neighbour.value == "L" || neighbour.value == "F") {
		start.neighbours = append(start.neighbours, neighbour)
	}
	if neighbour, ok := world[start.y][start.x+1]; ok && (neighbour.value == "-" || neighbour.value == "7" || neighbour.value == "J") {
		start.neighbours = append(start.neighbours, neighbour)
	}

	_, steps, _ := graph.FindShortestPath(start.neighbours[0], start.neighbours[1])

	return (steps / 2) + 1
}

func solve2023Day10Part2(lines []string) interface{} {
	lineMap := util.LinesToMap(lines)
	world := make(map[int]map[int]*pointNode)
	var start *pointNode
	for y := range lineMap {
		world[y] = make(map[int]*pointNode)
		for x := range lineMap[y] {
			world[y][x] = &pointNode{x, y, lineMap[y][x], make([]*pointNode, 0)}
			if lineMap[y][x] == "S" {
				start = world[y][x]
			}
		}
	}
	for y := range world {
		for x, node := range world[y] {
			switch world[y][x].value {
			// | is a vertical pipe connecting north and south.
			case "|":
				{
					if neighbour, ok := world[y-1][x]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y-1][x])
					}
					if neighbour, ok := world[y+1][x]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y+1][x])
					}
				}
			// - is a horizontal pipe connecting east and west.
			case "-":
				{
					if neighbour, ok := world[y][x-1]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y][x-1])
					}
					if neighbour, ok := world[y][x+1]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y][x+1])
					}
				}
			// L is a 90-degree bend connecting north and east.
			case "L":
				{
					{
						if neighbour, ok := world[y-1][x]; ok && neighbour.value != "." && neighbour.value != "S" {
							node.neighbours = append(node.neighbours, world[y-1][x])
						}
						if neighbour, ok := world[y][x+1]; ok && neighbour.value != "." && neighbour.value != "S" {
							node.neighbours = append(node.neighbours, world[y][x+1])
						}
					}
				}
			// J is a 90-degree bend connecting north and west.
			case "J":
				{
					if neighbour, ok := world[y-1][x]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y-1][x])
					}
					if neighbour, ok := world[y][x-1]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y][x-1])
					}
				}
			// 7 is a 90-degree bend connecting south and west.
			case "7":
				{
					if neighbour, ok := world[y+1][x]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y+1][x])
					}
					if neighbour, ok := world[y][x-1]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y][x-1])
					}
				}
			// F is a 90-degree bend connecting south and east.
			case "F":
				{
					if neighbour, ok := world[y+1][x]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y+1][x])
					}
					if neighbour, ok := world[y][x+1]; ok && neighbour.value != "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y][x+1])
					}
				}
			case ".":
				{
					if neighbour, ok := world[y-1][x]; ok && neighbour.value == "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y-1][x])
					}
					if neighbour, ok := world[y+1][x]; ok && neighbour.value == "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y+1][x])
					}
					if neighbour, ok := world[y][x-1]; ok && neighbour.value == "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y][x-1])
					}
					if neighbour, ok := world[y][x+1]; ok && neighbour.value == "." && neighbour.value != "S" {
						node.neighbours = append(node.neighbours, world[y][x+1])
					}
				}
			}
		}
	}

	if neighbour, ok := world[start.y-1][start.x]; ok && (neighbour.value == "|" || neighbour.value == "7" || neighbour.value == "F") {
		start.neighbours = append(start.neighbours, neighbour)
	}
	if neighbour, ok := world[start.y+1][start.x]; ok && (neighbour.value == "|" || neighbour.value == "L" || neighbour.value == "J") {
		start.neighbours = append(start.neighbours, neighbour)
	}
	if neighbour, ok := world[start.y][start.x-1]; ok && (neighbour.value == "-" || neighbour.value == "L" || neighbour.value == "F") {
		start.neighbours = append(start.neighbours, neighbour)
	}
	if neighbour, ok := world[start.y][start.x+1]; ok && (neighbour.value == "-" || neighbour.value == "7" || neighbour.value == "J") {
		start.neighbours = append(start.neighbours, neighbour)
	}

	path, _, _ := graph.FindShortestPath(start.neighbours[0], start.neighbours[1])
	pathPoints := make(map[util.Point]string)
	for _, n := range path {
		pathPoints[util.Point{X: n.(*pointNode).x, Y: n.(*pointNode).y}] = n.(*pointNode).value
	}
	pathPoints[util.Point{X: start.x, Y: start.y}] = "S"
	keys := make([]util.Point, 0, len(pathPoints))
	for p := range pathPoints {
		keys = append(keys, p)
	}

	min, max := util.MinMax(keys)
	count := 0
	for y := min.Y; y <= max.Y; y++ {
		isInside := false
		for x := min.X; x <= max.X; x++ {
			p := util.Point{X: x, Y: y}
			var ok bool
			var c string
			if c, ok = pathPoints[p]; ok && (c == "|" || c == "7" || c == "F" || c == "S") {
				isInside = !isInside
			}
			if !ok && isInside {
				count++
			}
		}
	}
	return count
}
