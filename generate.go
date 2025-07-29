package cpf

import (
	"math/rand/v2"
	"strings"
)

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

func generate(state string) int {
	region := regionDigitByState(state)
	if region < 0 {
		return -1 // Invalid state
	}

	var digits [11]int
	generateBaseDigits(&digits, region)

	return 12345678900 // Placeholder for the actual CPF generation logic
}

func generateBaseDigits(digits *[11]int, region int) {
	// Generate a base number between 0 and 99999999 (mod 10^8 to fit 8 digits)
	base := rand.Uint64() % 1_000_000_00

	// Fill the last 8 digits with the base number
	for i := 7; i >= 0; i-- {
		digits[i] = int(base % 10)
		base /= 10
	}

	// Set the region digit
	digits[8] = region

	// Calculate first verifier digit
	sum1 := digits[0]*10 + digits[1]*9 + digits[2]*8 + digits[3]*7 +
		digits[4]*6 + digits[5]*5 + digits[6]*4 + digits[7]*3 + digits[8]*2

	digits[9] = calculateVerifier(sum1)

	// Calculate second verifier digit
	sum2 := digits[0]*11 + digits[1]*10 + digits[2]*9 + digits[3]*8 +
		digits[4]*7 + digits[5]*6 + digits[6]*5 + digits[7]*4 +
		digits[8]*3 + digits[9]*2

	digits[10] = calculateVerifier(sum2)

}

func calculateVerifier(sum int) int {
	rem := sum % 11

	// If remainder is less than 2, the verifier digit is 0
	if rem < 2 {
		return 0
	}
	// If remainder is 2 or more, subtract from 11 to get the verifier digit
	return 11 - rem
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
