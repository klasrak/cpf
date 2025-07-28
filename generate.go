package cpf

import "strings"

var digits = map[string]int{
	"RS": 0,
	"DF": 1, "GO": 1, "MS": 1, "MT": 1, "TO": 1,
	"AC": 2, "AM": 2, "AP": 2, "PA": 2, "RO": 2, "RR": 2,
	"CE": 3, "MA": 3, "PI": 3,
	"AL": 4, "PB": 4, "PE": 4, "RN": 4,
	"BA": 5, "SE": 5,
	"MG": 6,
	"ES": 7, "RJ": 7,
	"SP": 8,
	"PR": 9, "SC": 9,
}

func generate(state string) string {
	return "123.456.789-00" // Placeholder for actual CPF generation logic
}

func regionDigitByState(state string) int {
	state = validateStateInput(state)
	if state == "" {
		return -1 // Invalid state
	}

	if digit, ok := digits[state]; ok {
		return digit
	}

	return -1 // Invalid state
}

func validateStateInput(state string) string {
	if len(state) != 2 {
		return ""
	}
	for _, char := range state {
		if char < 'A' || char > 'Z' {
			return ""
		}
	}
	return strings.ToUpper(state)
}
