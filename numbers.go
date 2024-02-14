package seven_segment_indicator

// Numbers Набор цифр от 0 до 9 и знак минус
type Numbers struct {
}

func (Numbers) Char(r rune) (aPin, bPin, cPin, dPin, ePin, fPin, gPin bool) {
	switch r {
	case '-':
		return false, false, false, false, false, false, true
	case '0':
		return true, true, true, true, true, true, false
	case '1':
		return false, true, true, false, false, false, false
	case '2':
		return true, true, false, true, true, false, true
	case '3':
		return true, true, true, true, false, false, true
	case '4':
		return false, true, true, false, false, true, true
	case '5':
		return true, false, true, true, false, true, true
	case '6':
		return true, false, true, true, true, true, true
	case '7':
		return true, true, true, false, false, false, false
	case '8':
		return true, true, true, true, true, true, true
	case '9':
		return true, true, true, true, false, true, true
	}

	return
}
