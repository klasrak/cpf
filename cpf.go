package cpf

// Generate generates a valid brazilian CPF number for the given state or returns an error if the state is invalid.
//
//	Parameters:
//	   state string: A two-letter string representing the Brazilian state (e.g., "SP" for São Paulo).
//	Returns:
//	   string: A valid CPF number in the format "XXX.XXX.XXX-XX".
//	   error: An error if the state is invalid or if CPF generation fails.
//
// To generate a valid Brazilian CPF, start with a 9-digit number (e.g., 111.444.77<state digit>-XX) and calculate two check digits using the modulo 11 algorithm.
//
//	First Check Digit:
//
// Assign weights 10 to 2 (left to right) to the 9 digits.
// Multiply each digit by its weight and sum the results (e.g., 10+9+8+28+24+20+28+21+14 = 162).
// Divide the sum by 11, take the remainder (162 ÷ 11 = remainder 8).
// If remainder < 2, the digit is 0; otherwise, subtract from 11 (11 - 8 = 3). CPF becomes 111.444.777-3X.
//
//	Second Check Digit:
//
// Include the first check digit, assign weights 11 to 2 to the 10 digits.
// Multiply and sum (e.g., 11+10+9+32+28+24+35+28+21+6 = 204).
// Divide by 11, take the remainder (204 ÷ 11 = remainder 6).
// If remainder < 2, the digit is 0; otherwise, subtract from 11 (11 - 6 = 5).
// Final CPF: 111.444.777-35. The generator creates valid CPFs by randomly selecting 9 digits, calculating the first check digit, appending it, then calculating the second.
func Generate(state string) string {
	return generate(state)
}
