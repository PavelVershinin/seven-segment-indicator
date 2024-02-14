package main

import (
	"github.com/PavelVershinin/seven-segment-indicator"
	"machine"
	"math"
	"strconv"
)

func main() {
	// Создаём двухразрядный 7-сегментный индикатор с общим катодом
	ssi := seven_segment_indicator.New(
		// Набор символов
		seven_segment_indicator.Numbers{},
		// Тип индикатора, с общим катодом или общим анодом
		seven_segment_indicator.ModeCommonCathode,
		// Пины от A до G и точка Dp, для всех секций общие
		machine.D3,
		machine.D4,
		machine.D5,
		machine.D6,
		machine.D7,
		machine.D8,
		machine.D9,
		machine.D10,
		// Катоды для каждого разряда отдельно
		machine.D11,
		machine.D12,
	)

	// Закинем, 3.14 в индикатор, так как индикатор на два разряда, на нём будет выведено 3.1
	ssi.SetValue([]rune(strconv.FormatFloat(math.Pi, 'f', -1, 64)))

	for {
		// В бесконечном цикле, перерисовываем значения на индикаторе
		ssi.Draw()
	}
}
