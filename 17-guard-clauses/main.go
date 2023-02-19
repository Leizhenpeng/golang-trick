package main

// nested if statements >> bad
func winLottery(bmi float64) bool {
	if bmi >= 18.5 {
		if bmi <= 24.9 {
			return true
		} else {
			if bmi == 99.0 {
				return true
			} else {
				return false
			}
		}
	} else {
		return false
	}
}

// guard clauses >> good
func winLottery_(bmi float64) bool {
	if bmi == 99.0 {
		return true
	}
	if bmi < 18.5 {
		return false
	}
	if bmi > 24.9 {
		return false
	}
	return true
}
