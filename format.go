package cpf

var (
	pow11 = [11]int{10000000000, 1000000000, 100000000, 10000000, 1000000, 100000, 10000, 1000, 100, 10, 1}
)

func maskInt(cpf int) string {
	if cpf <= 0 {
		return ""
	}

	digits := toDigits(int64(cpf))

	result := bufferToStringMasked(&digits)

	if len(result) != 14 {
		return ""
	}

	return result
}

func maskString(cpf string) string {
	var digits [11]int
	if ok := extractDigits(cpf, &digits); !ok {
		return ""
	}

	result := bufferToStringMasked(&digits)

	if len(result) != 14 {
		return ""
	}

	return result
}

func unmask(cpf string) string {
	var digits [11]int
	if ok := extractDigits(cpf, &digits); !ok {
		return ""
	}

	return bufferToStringUnmasked(&digits)
}

func unmaskToInt(cpf string) int {
	var digits [11]int
	if ok := extractDigits(cpf, &digits); !ok {
		return -1
	}

	return toNumber(&digits)
}

func bufferToStringMasked(digits *[11]int) string {
	var buf [14]byte
	buf[3], buf[7], buf[11] = '.', '.', '-'

	idx := [11]int{0, 1, 2, 4, 5, 6, 8, 9, 10, 12, 13}
	for i, d := range digits {
		buf[idx[i]] = byte(d + '0')
	}

	return string(buf[:])
}

func bufferToStringUnmasked(digits *[11]int) string {
	var buf [11]byte

	for i, d := range digits {
		buf[i] = byte(d + '0')
	}

	return string(buf[:])
}

func extractDigits(cpf string, digits *[11]int) bool {
	var count int
	for i := 0; i < len(cpf); i++ {
		c := cpf[i]
		if c >= '0' && c <= '9' {
			if count >= 11 {
				return false
			}
			digits[count] = int(c - '0')
			count++
		}
	}
	return count == 11
}

func toDigits(num int64) [11]int {
	var digits [11]int
	for i := 10; i >= 0; i-- {
		digits[i] = int(num % 10)
		num /= 10
	}
	return digits
}

func toNumber(digits *[11]int) int {
	result := 0
	for i, d := range digits {
		result += d * pow11[i]
	}
	return result
}
