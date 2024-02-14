// Package seven_segment_indicator
// Управление семисегментным индикатором, без использования драйвера
package seven_segment_indicator

import (
	"machine"
)

// Mode Тип индикатора, с общим катодом или общим анодом
type Mode uint8

const (
	// ModeCommonAnode С общим анодом +
	ModeCommonAnode Mode = iota
	// ModeCommonCathode С общим катодом -
	ModeCommonCathode
)

// CharacterSetter Интерфейс получения списка включаемых пинов, по руне
type CharacterSetter interface {
	Char(r rune) (aPin, bPin, cPin, dPin, ePin, fPin, gPin bool)
}

type segment struct {
	pin  machine.Pin
	char rune
	dot  bool
}

// SevenSegmentsIndicator Семисегментный индикатор
type SevenSegmentsIndicator struct {
	characterSet CharacterSetter
	mode         Mode

	aPin  machine.Pin
	bPin  machine.Pin
	cPin  machine.Pin
	dPin  machine.Pin
	ePin  machine.Pin
	fPin  machine.Pin
	gPin  machine.Pin
	dpPin machine.Pin

	digits []segment
}

// New Создание нового индикатора
func New(characterSet CharacterSetter, mode Mode, aPin, bPin, cPin, dPin, ePin, fPin, gPin, dpPin machine.Pin, commonPins ...machine.Pin) *SevenSegmentsIndicator {
	s := &SevenSegmentsIndicator{
		characterSet: characterSet,
		mode:         mode,

		aPin:  aPin,
		bPin:  bPin,
		cPin:  cPin,
		dPin:  dPin,
		ePin:  ePin,
		fPin:  fPin,
		gPin:  gPin,
		dpPin: dpPin,

		digits: make([]segment, len(commonPins)),
	}

	for i, pin := range commonPins {
		s.digits[i] = segment{
			pin:  pin,
			char: -1,
			dot:  false,
		}
	}

	return s
}

// SetValue Определяем строку, которую надо отобразить на индикаторе
func (ssi *SevenSegmentsIndicator) SetValue(str []rune) {
	digitIndex := 0
	for i := 0; i < len(str) && digitIndex < len(ssi.digits); i++ {
		r := str[i]
		if r == '.' {
			continue
		}
		ssi.digits[digitIndex].char = r
		if digitIndex+2 < len(str) && str[digitIndex+1] == '.' {
			ssi.digits[digitIndex].dot = true
		}
		digitIndex++
	}

	for ; digitIndex < len(ssi.digits); digitIndex++ {
		ssi.digits[digitIndex].char = -1
		ssi.digits[digitIndex].dot = false
	}
}

// Draw Перерисовка индикатора, необходимо вызывать её постоянно при каждой итерации
func (ssi *SevenSegmentsIndicator) Draw() {
	for _, d := range ssi.digits {
		for _, cp := range ssi.digits {
			ssi.setPin(cp.pin, true, false)
		}
		aPin, bPin, cPin, dPin, ePin, fPin, gPin := ssi.characterSet.Char(d.char)
		ssi.setPin(ssi.aPin, false, aPin)
		ssi.setPin(ssi.bPin, false, bPin)
		ssi.setPin(ssi.cPin, false, cPin)
		ssi.setPin(ssi.dPin, false, dPin)
		ssi.setPin(ssi.ePin, false, ePin)
		ssi.setPin(ssi.fPin, false, fPin)
		ssi.setPin(ssi.gPin, false, gPin)
		ssi.setPin(ssi.dpPin, false, d.dot)
		ssi.setPin(d.pin, true, true)
	}
}

func (ssi *SevenSegmentsIndicator) setPin(p machine.Pin, isCommonPin, on bool) {
	p.Configure(machine.PinConfig{Mode: machine.PinOutput})

	switch ssi.mode {
	case ModeCommonAnode:
		if !isCommonPin {
			on = !on
		}
	case ModeCommonCathode:
		if isCommonPin {
			on = !on
		}
	default:
		return
	}

	p.Set(on)
}
