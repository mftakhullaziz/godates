package common

func CheckInt64NotNilAndNotZero(val *int64) bool {
	return val != nil && *val != 0
}

func CheckIntNotNilAndNotZero(val *int) bool {
	return val != nil && *val != 0
}

func CheckInt64IsNilAndZero(val *int64) bool {
	return val == nil || *val == 0
}

func CheckIntIsNilAndZero(val *int) bool {
	return val == nil || *val == 0
}
