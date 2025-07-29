package cpf

func isValidString(cpf string) bool {
	var digits [11]int
	ok := extractDigits(cpf, &digits)

	if !ok {
		return false
	}

	if isAllDigitsEqual(&digits) {
		return false
	}

	if !validateFirstVerifier(&digits) {
		return false
	}

	return validateSecondVerifier(&digits)
}

func isValidInt(cpf int) bool {
	if cpf <= 0 {
		return false
	}

	var digits [11]int
	for i := 10; i >= 0; i-- {
		digits[i] = cpf % 10
		cpf /= 10
	}

	if isAllDigitsEqual(&digits) {
		return false
	}

	if !validateFirstVerifier(&digits) {
		return false
	}

	return validateSecondVerifier(&digits)
}

// validateFirstVerifier checks the 10th digit (first verifier).
func validateFirstVerifier(digits *[11]int) bool {
	sum := digits[0]*10 + digits[1]*9 + digits[2]*8 + digits[3]*7 +
		digits[4]*6 + digits[5]*5 + digits[6]*4 + digits[7]*3 + digits[8]*2
	check := calculateVerifier(sum)
	return check == digits[9]
}

// validateSecondVerifier checks the 11th digit (second verifier).
func validateSecondVerifier(digits *[11]int) bool {
	sum := digits[0]*11 + digits[1]*10 + digits[2]*9 + digits[3]*8 +
		digits[4]*7 + digits[5]*6 + digits[6]*5 + digits[7]*4 +
		digits[8]*3 + digits[9]*2
	check := calculateVerifier(sum)
	return check == digits[10]
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

func isAllDigitsEqual(digits *[11]int) bool {
	first := digits[0]
	for i := 1; i < 11; i++ {
		if digits[i]^first != 0 { // Bitwise XOR for comparison.
			return false
		}
	}
	return true
}
