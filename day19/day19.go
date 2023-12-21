package day19

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2023, 19, solve2023Day19Part1, solve2023Day19Part2)
}

type ruleType struct {
	field       string
	operator    string
	value       int
	destination string
}

func solve2023Day19Part1(lines []string) interface{} {
	sections := util.SplitLines(lines, "")
	ruleMap := make(map[string][]ruleType)
	for _, rule := range sections[0] {
		var ruleName, conditions string
		ruleName = rule[:strings.Index(rule, "{")]
		conditions = rule[strings.Index(rule, "{")+1 : len(rule)-1]
		ruleMap[ruleName] = make([]ruleType, 0)
		rules := strings.Split(conditions, ",")
		for _, condition := range rules[:len(rules)-1] {
			field := string(condition[0])
			operator := string(condition[1])
			destination := condition[strings.Index(condition, ":")+1:]
			value, _ := strconv.Atoi(condition[2:strings.Index(condition, ":")])
			ruleMap[ruleName] = append(ruleMap[ruleName], ruleType{field, operator, value, destination})
		}
		ruleMap[ruleName] = append(ruleMap[ruleName], ruleType{"", "", 0, rules[len(rules)-1]})
	}
	sum := 0
	for _, part := range sections[1] {
		var x, m, a, s int
		fmt.Sscanf(part, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
		currentPart := map[string]int{"x": x, "m": m, "a": a, "s": s, "T": x + m + a + s}
		currentWorkflow := "in"
		finished := false
		for !finished {
		ruleLoop:
			for _, rule := range ruleMap[currentWorkflow] {
				switch rule.operator {
				case "<":
					if currentPart[rule.field] < rule.value {
						currentWorkflow = rule.destination
						break ruleLoop
					}
				case ">":
					if currentPart[rule.field] > rule.value {
						currentWorkflow = rule.destination
						break ruleLoop
					}
				default:
					{
						currentWorkflow = rule.destination
					}
				}
			}
			if currentWorkflow == "A" {
				sum += currentPart["T"]
				finished = true
			} else if currentWorkflow == "R" {
				finished = true
			}
		}
	}
	return sum
}

func solve2023Day19Part2(lines []string) interface{} {
	sections := util.SplitLines(lines, "")
	ruleMap := make(map[string][]ruleType)
	for _, rule := range sections[0] {
		var ruleName, conditions string
		ruleName = rule[:strings.Index(rule, "{")]
		conditions = rule[strings.Index(rule, "{")+1 : len(rule)-1]
		ruleMap[ruleName] = make([]ruleType, 0)
		rules := strings.Split(conditions, ",")
		for _, condition := range rules[:len(rules)-1] {
			field := string(condition[0])
			operator := string(condition[1])
			destination := condition[strings.Index(condition, ":")+1:]
			value, _ := strconv.Atoi(condition[2:strings.Index(condition, ":")])
			ruleMap[ruleName] = append(ruleMap[ruleName], ruleType{field, operator, value, destination})
		}
		ruleMap[ruleName] = append(ruleMap[ruleName], ruleType{"", "", 0, rules[len(rules)-1]})
	}
	maxPart := map[string]int{"x": 4000, "m": 4000, "a": 4000, "s": 4000}
	minPart := map[string]int{"x": 1, "m": 1, "a": 1, "s": 1}
	return calculateCombinations("in", minPart, maxPart, ruleMap)
}

func calculateCombinations(currentWorkflow string, min, max map[string]int, ruleMap map[string][]ruleType) int {
	if currentWorkflow == "R" {
		return 0
	}
	if currentWorkflow == "A" {
		return (max["x"] - min["x"] + 1) * (max["m"] - min["m"] + 1) * (max["a"] - min["a"] + 1) * (max["s"] - min["s"] + 1)
	}

	total := 0
	rules := ruleMap[currentWorkflow]
	for _, rule := range rules {
		if rule.operator == "<" {
			if min[rule.field] < rule.value {
				newMin := map[string]int{"x": min["x"], "m": min["m"], "a": min["a"], "s": min["s"]}
				newMax := map[string]int{"x": max["x"], "m": max["m"], "a": max["a"], "s": max["s"]}
				newMax[rule.field] = rule.value - 1
				total += calculateCombinations(rule.destination, newMin, newMax, ruleMap)
			}
			if max[rule.field] >= rule.value {
				min[rule.field] = rule.value
			} else {
				break
			}
		} else {
			if max[rule.field] > rule.value {
				newMin := map[string]int{"x": min["x"], "m": min["m"], "a": min["a"], "s": min["s"]}
				newMax := map[string]int{"x": max["x"], "m": max["m"], "a": max["a"], "s": max["s"]}
				newMin[rule.field] = rule.value + 1
				total += calculateCombinations(rule.destination, newMin, newMax, ruleMap)
			}
			if min[rule.field] <= rule.value {
				max[rule.field] = rule.value
			} else {
				break
			}
		}
	}
	total += calculateCombinations(rules[len(rules)-1].destination, min, max, ruleMap)
	return total
}
