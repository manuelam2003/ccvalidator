package luhn

func CheckLuhn(cardNumber string) bool {
	n := len(cardNumber)
	sum := 0
	double := true

	for i := n - 2; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	checkDigit := 10 - sum%10
	return int(cardNumber[n-1]-'0') == checkDigit
}
