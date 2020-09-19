package transfer

const bonusPercent = 0.0050

func Bonus(amount int) int {
	bonus := int(float64(amount) * bonusPercent)
	return bonus
}

func Total(amount int) int {
	total := amount + Bonus(amount)
	return total
}
