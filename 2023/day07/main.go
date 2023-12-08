package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	var ans int = part1(input)
	fmt.Println(ans)
}

type HandType int

type Hand struct {
	Cards string
	Type  HandType
}

const (
	HighCard     HandType = 1
	OnePair      HandType = 2
	TwoPair      HandType = 3
	ThreeOfAKind HandType = 4
	FullHouse    HandType = 5
	FourOfAKind  HandType = 6
	FiveOfAKind  HandType = 7
)

func part1(input []string) int {
	handToBet := make(map[Hand]int)
	hands := []Hand{}

	for _, line := range input {
		game := strings.Fields(line)
		cards := game[0]
		bet, _ := strconv.Atoi(game[1])

		hand := Hand{
			Cards: cards,
			Type:  getHandType(cards),
		}

		handToBet[hand] = bet
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHandType(hands[i], hands[j], false)
	})

	ans := 0
	for i, hand := range hands {
		ans += handToBet[hand] * (len(hands) - i)
	}

	return ans
}

func part2(input []string) int {
	handToBet := make(map[Hand]int)
	hands := []Hand{}

	for _, line := range input {
		game := strings.Fields(line)
		cards := game[0]
		bet, _ := strconv.Atoi(game[1])

		hand := Hand{
			Cards: cards,
			Type:  getHandType(cards),
		}

		if strings.Contains(cards, "J") {
			hand.Type = findBestHand(cards)
		}

		handToBet[hand] = bet
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHandType(hands[i], hands[j], true)
	})

	ans := 0
	for i, hand := range hands {
		ans += handToBet[hand] * (len(hands) - i)
	}

	return ans
}

func findBestHand(hand string) HandType {
	rankingMap := map[byte]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}

	bestScore := 0
	for card := range rankingMap {
		newHand := strings.ReplaceAll(hand, "J", string(card))
		score := int(getHandType(newHand))

		if score > bestScore {
			bestScore = score
		}
	}

	return HandType(bestScore)
}

func getHandType(hand string) HandType {
	cardCount := make(map[rune]int)
	maxRepeats := 0
	for _, card := range hand {
		cardCount[card]++
		maxRepeats = max(maxRepeats, cardCount[card])
	}

	switch {
	case maxRepeats == 5 && len(cardCount) == 1: // Five of a kind
		return FiveOfAKind
	case maxRepeats == 4 && len(cardCount) == 2: // Four of a kind
		return FourOfAKind
	case maxRepeats == 3 && len(cardCount) == 2: // Full House
		return FullHouse
	case maxRepeats == 3 && len(cardCount) == 3: // Three of a kind
		return ThreeOfAKind
	case maxRepeats == 2 && len(cardCount) == 3: // Two Pair
		return TwoPair
	case maxRepeats == 2 && len(cardCount) == 4: // One Pair
		return OnePair
	default: // High Card
		return HighCard
	}
}

func compareHandType(hand1, hand2 Hand, joker bool) bool {
	type1 := hand1.Type
	type2 := hand2.Type

	if type1 > type2 {
		return true
	}
	if type1 < type2 {
		return false
	}
	return compareCards(hand1.Cards, hand2.Cards, joker)
}

func compareCards(hand1, hand2 string, joker bool) bool {
	rankingMap := map[byte]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	if joker {
		rankingMap['J'] = 1
	}

	for i := 0; i < 5; i++ {
		card1 := rankingMap[byte(hand1[i])]
		card2 := rankingMap[byte(hand2[i])]

		if card1 > card2 {
			return true
		}

		if card1 < card2 {
			return false
		}
	}

	return false
}
