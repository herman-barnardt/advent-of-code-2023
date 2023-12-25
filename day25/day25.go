package day25

import (
	"fmt"
	"math"
	"slices"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2023, 25, solve2023Day25Part1, solve2023Day25Part2)
}

func solve2023Day25Part1(lines []string) interface{} {
	edges := make(map[string]map[string]bool)
	for _, line := range lines {
		parts := strings.Split(strings.ReplaceAll(line, ":", ""), " ")
		if _, ok := edges[parts[0]]; !ok {
			edges[parts[0]] = make(map[string]bool)
		}
		for _, dest := range parts[1:] {
			edges[parts[0]][dest] = true
			if _, ok := edges[dest]; !ok {
				edges[dest] = make(map[string]bool)
			}
			edges[dest][parts[0]] = true
		}
	}

	edgeCount := make(map[string]int)
	for n := range edges {
		reachAll(edges, n, edgeCount)
	}

	max, max2, max3 := math.MinInt, math.MinInt, math.MinInt
	var e, e2, e3 string
	for k, v := range edgeCount {
		if v > max {
			max3 = max2
			max2 = max
			max = v
			e3 = e2
			e2 = e
			e = k
		} else if v > max2 {
			max3 = max2
			max2 = v
			e3 = e2
			e2 = k
		} else if v > max3 {
			max3 = v
			e3 = k
		}
	}

	for n := range edges {
		d1, d2 := e[:strings.Index(e, ".")], e[strings.Index(e, ".")+1:]
		if n == d1 {
			delete(edges[n], d2)
			delete(edges[d2], n)
		}
		if n == d2 {
			delete(edges[n], d1)
			delete(edges[d1], n)
		}
		d1, d2 = e2[:strings.Index(e2, ".")], e2[strings.Index(e2, ".")+1:]
		if n == d1 {
			delete(edges[n], d2)
			delete(edges[d2], n)
		}
		if n == d2 {
			delete(edges[n], d1)
			delete(edges[d1], n)
		}
		d1, d2 = e3[:strings.Index(e3, ".")], e3[strings.Index(e3, ".")+1:]
		if n == d1 {
			delete(edges[n], d2)
			delete(edges[d2], n)
		}
		if n == d2 {
			delete(edges[n], d1)
			delete(edges[d1], n)
		}
	}

	var first string
	for v := range edges {
		first = v
		break
	}
	toVisit := []string{first}
	visited := make(map[string]bool)
	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]
		visited[current] = true

		for n := range edges[current] {
			if _, ok := visited[n]; !ok {
				toVisit = append(toVisit, n)
			}
		}
	}

	fmt.Println(len(visited))

	return len(visited) * (len(edges) - len(visited))
}

func solve2023Day25Part2(lines []string) interface{} {
	return "Push the Big Red Button"
}

func reachAll(graph map[string]map[string]bool, start string, counter map[string]int) int {
	seen := make(map[string]bool)
	pending := []string{start}
	seen[start] = true
	for len(pending) > 0 {
		current := pending[0]
		pending = pending[1:]
		edges := graph[current]
		for e := range edges {
			if _, ok := seen[e]; ok {
				continue
			}
			seen[e] = true
			names := []string{current, e}
			slices.Sort(names)
			key := strings.Join(names, ".")
			counter[key] = counter[key] + 1
			pending = append(pending, e)
		}
	}
	return len(seen)
}
