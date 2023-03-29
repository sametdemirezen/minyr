package yr

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/sametdemirezen/funtemps/conv"
)

func CelsiusToFahrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFahrenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

func CelsiusToFahrenheitLine(line string) (string, error) {
	elementsInLine := strings.Split(line, ";")
	var err error
	if len(elementsInLine) == 4 {
		elementsInLine[3], err = CelsiusToFahrenheitString(elementsInLine[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(elementsInLine, ";"), nil
}
