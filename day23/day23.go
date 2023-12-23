package day23

import (
	"math"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 23, solve2023Day23Part1, solve2023Day23Part2)
}

var directions = map[string]util.Point{">": {X: 1, Y: 0}, "<": {X: -1, Y: 0}, "v": {X: 0, Y: 1}, "^": {X: 0, Y: -1}}
var oppositeDirections = map[string]string{">": "<", "<": ">", "^": "v", "v": "^"}

type pathNode struct {
	point      util.Point
	value      string
	neighbours map[*pathNode]int
}

func (n *pathNode) GetNeighbours() []*pathNode {
	neighbours := make([]*pathNode, 0)
	if n.value == "." {
		for indicator, direction := range directions {
			point := n.point.Add(direction)
			if node, ok := nodeMap[point]; ok && node.value != "#" && node.value != oppositeDirections[indicator] {
				neighbours = append(neighbours, node)
			}
		}
	} else {
		point := n.point.Add(directions[n.value])
		if node, ok := nodeMap[point]; ok && node.value != "#" {
			neighbours = append(neighbours, node)
		}
	}

	for _, neighbour := range neighbours {
		n.neighbours[neighbour] = 1
	}

	return neighbours
}

var nodeMap = map[util.Point]*pathNode{}

func solve2023Day23Part1(lines []string) interface{} {
	grid := util.LinesToPointMap(lines)
	var startNode, endNode *pathNode
	for point, value := range grid {
		if value != "#" {
			node := &pathNode{point, value, make(map[*pathNode]int)}
			nodeMap[point] = node
			if point.Y == 0 {
				startNode = node
			}
			if point.Y == len(lines)-1 {
				endNode = node
			}
		}
	}

	for _, node := range nodeMap {
		node.GetNeighbours()
	}

	return getLongestPath(startNode, endNode, make(map[*pathNode]bool))
}

func solve2023Day23Part2(lines []string) interface{} {
	nodeMap = map[util.Point]*pathNode{}
	grid := util.LinesToPointMap(lines)
	var startNode, endNode *pathNode
	for point, value := range grid {
		if value != "#" {
			node := &pathNode{point, ".", make(map[*pathNode]int)}
			nodeMap[point] = node
			if point.Y == 0 {
				startNode = node
			}
			if point.Y == len(lines)-1 {
				endNode = node
			}
		}
	}
	for _, node := range nodeMap {
		node.GetNeighbours()
	}
	for {
		count := 0
		nodesToDelete := make([]*pathNode, 0)
		for _, node := range nodeMap {
			if len(node.neighbours) == 2 {
				count++
				keys := make([]*pathNode, 0, len(node.neighbours))
				sum := 0
				for k, v := range node.neighbours {
					keys = append(keys, k)
					sum += v
				}
				delete(keys[0].neighbours, node)
				delete(keys[1].neighbours, node)
				keys[1].neighbours[keys[0]] = sum
				keys[0].neighbours[keys[1]] = sum
				node.neighbours = make(map[*pathNode]int)
				nodesToDelete = append(nodesToDelete, node)
			}
		}
		for _, n := range nodesToDelete {
			delete(nodeMap, n.point)
		}
		if count == 0 {
			break
		}
	}

	return getLongestPath(startNode, endNode, make(map[*pathNode]bool))
}

func getLongestPath(current *pathNode, end *pathNode, seen map[*pathNode]bool) int {
	seen[current] = true
	max := math.MinInt
	if current == end {
		max = 0
	} else {
		for next, cost := range current.neighbours {
			if _, ok := seen[next]; !ok {
				max = util.IntMax(max, getLongestPath(next, end, seen)+cost)
			}
		}
	}
	delete(seen, current)
	return max
}
