package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sumc := ParseCard(card1) + ParseCard(card2)
	dv := ParseCard(dealerCard)
	switch {
	case (sumc == 22):
		return "P"
	case (sumc == 21 && dv == 11):
		return "S"
	case (sumc == 21 && dv == 10):
		return "S"
	case (sumc == 21 && dv != 10):
		return "W"
	case (sumc >= 17 && sumc <= 20):
		return "S"
	case ((sumc >= 12 || sumc <= 16) && dv >= 7):
		return "H"
	case (sumc >= 12 && sumc <= 16):
		return "S"
	case (sumc <= 11):
		return "H"
	default:
		return ""
	}
}
