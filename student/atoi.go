package student

func Atoi(s string) int {
	result := 0
	sign := 1

	for i, char := range s {
		if char == '+' && i == 0 {
			continue
		} else if char == '-' && i == 0 {
			sign = -1
			continue
		}
		if char < '0' || char > '9' {
			return 0
		}
		digit := int(char - '0')
		result = result*10 + digit
	}

	return result * sign
}
