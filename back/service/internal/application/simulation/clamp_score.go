package simulation

func clampScore(score int) int {
	if score < 0 {
		return 0
	}
	if score > 5 {
		return 5
	}
	return score
}
