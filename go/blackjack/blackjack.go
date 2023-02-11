package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace": return 11
	case "two": return 2
	case "three": return 3
	case "four": return 4
	case "five": return 5
	case "six": return 6
	case "seven": return 7
	case "eight": return 8
	case "nine": return 9
	case "ten": return 10
	case "jack": return 10
	case "queen": return 10
	case "king": return 10
	default: return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardsFaceValue := ParseCard(card1) + ParseCard(card2)

	dealerCardFaceValue := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case cardsFaceValue == 21 && (dealerCardFaceValue == 10 || dealerCardFaceValue == 11):
		return "S"
	case cardsFaceValue == 21:
		return "W"
	case cardsFaceValue >= 17 && cardsFaceValue <= 20:
		return "S"
	case (cardsFaceValue >= 12 && cardsFaceValue <= 16) && dealerCardFaceValue >= 7:
		return "H"
	case (cardsFaceValue >= 12 && cardsFaceValue <= 16):
		return "S"
	case cardsFaceValue <= 11:
		return "H"
	default: return "S"
	}
}
