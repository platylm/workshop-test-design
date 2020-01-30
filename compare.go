package testing

func maxAndMean(numberA, numberB, numberC int) int {
	if numberA > numberB {
		if numberA > numberC {
			return numberA
		}
	}
	if numberB > numberC {
		return numberB
	}
	return numberC
}
