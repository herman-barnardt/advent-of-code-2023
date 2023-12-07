package day7

import (
	"fmt"
	"sort"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2023, 7, solve2023Day7Part1, solve2023Day7Part2)
}

type camelCardHand struct {
	handType int
	cards    string
	bid      int
}

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR        = 3
	PAIR            = 2
	HIGH_CARD       = 1
)

func solve2023Day7Part1(lines []string) interface{} {
	var cardValue = map[string]int{
		"2": 1,
		"3": 2,
		"4": 3,
		"5": 4,
		"6": 5,
		"7": 6,
		"8": 7,
		"9": 8,
		"T": 9,
		"J": 10,
		"Q": 11,
		"K": 12,
		"A": 13,
	}
	hands := make([]camelCardHand, 0)
	for _, line := range lines {
		var handString string
		var bid int
		fmt.Sscanf(line, "%s %d", &handString, &bid)
		hand := make(map[string]int)
		for _, c := range strings.Split(handString, "") {
			hand[c]++
		}

		countMap := make(map[int]int)
		for _, v := range hand {
			countMap[v]++
		}
		score := HIGH_CARD
		if countMap[5] > 0 {
			score = FIVE_OF_A_KIND
		} else if countMap[4] > 0 {
			score = FOUR_OF_A_KIND
		} else if countMap[3] > 0 && countMap[2] > 0 {
			score = FULL_HOUSE
		} else if countMap[3] > 0 {
			score = THREE_OF_A_KIND
		} else if countMap[2] > 1 {
			score = TWO_PAIR
		} else if countMap[2] > 0 {
			score = PAIR
		}
		hands = append(hands, camelCardHand{score, handString, bid})
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			c := 0
			for cardValue[string(hands[i].cards[c])] == cardValue[string(hands[j].cards[c])] {
				c++
			}
			return cardValue[string(hands[i].cards[c])] < cardValue[string(hands[j].cards[c])]
		}

		return hands[i].handType < hands[j].handType
	})
	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}
	return sum
}

func solve2023Day7Part2(lines []string) interface{} {
	var cardValue = map[string]int{
		"J": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"Q": 11,
		"K": 12,
		"A": 13,
	}
	hands := make([]camelCardHand, 0)
	for _, line := range lines {
		var handString string
		var bid int
		fmt.Sscanf(line, "%s %d", &handString, &bid)
		hand := make(map[string]int)
		for _, c := range strings.Split(handString, "") {
			hand[c]++
		}

		countMap := make(map[int]int)
		for k, v := range hand {
			if k != "J" {
				countMap[v]++
			}
		}
		jCount := hand["J"]
		score := HIGH_CARD
		if jCount == 0 {
			if countMap[5] > 0 {
				score = FIVE_OF_A_KIND
			} else if countMap[4] > 0 {
				score = FOUR_OF_A_KIND
			} else if countMap[3] > 0 && countMap[2] > 0 {
				score = FULL_HOUSE
			} else if countMap[3] > 0 {
				score = THREE_OF_A_KIND
			} else if countMap[2] > 1 {
				score = TWO_PAIR
			} else if countMap[2] > 0 {
				score = PAIR
			}
		} else {
			if jCount == 1 {
				score = PAIR
				if countMap[4] > 0 {
					score = FIVE_OF_A_KIND
				} else if countMap[3] > 0 {
					score = FOUR_OF_A_KIND
				} else if countMap[2] > 1 {
					score = FULL_HOUSE
				} else if countMap[2] > 0 {
					score = THREE_OF_A_KIND
				}
			}
			if jCount == 2 {
				score = THREE_OF_A_KIND
				if countMap[3] > 0 {
					score = FIVE_OF_A_KIND
				} else if countMap[2] > 0 {
					score = FOUR_OF_A_KIND
				}
			}
			if jCount == 3 {
				score = FOUR_OF_A_KIND
				if countMap[2] > 0 {
					score = FIVE_OF_A_KIND
				}
			}
			if jCount == 4 || jCount == 5 {
				score = FIVE_OF_A_KIND
			}
		}
		hands = append(hands, camelCardHand{score, handString, bid})
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			c := 0
			for cardValue[string(hands[i].cards[c])] == cardValue[string(hands[j].cards[c])] {
				c++
			}
			return cardValue[string(hands[i].cards[c])] < cardValue[string(hands[j].cards[c])]
		}

		return hands[i].handType < hands[j].handType
	})
	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}
	return sum
}
