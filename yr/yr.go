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
	elements := strings.Split(line, ";")
	var err error
	if len(elements) == 4 {
		if strings.HasSuffix(elements[3], "tur") {
			return line, nil
		} else if strings.HasPrefix(elements[0], "Data") {
			return "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Samet Demirezen", nil
		} else {
			elements[3], err = CelsiusToFahrenheitString(elements[3])
			if err != nil {
				return "", err
			}
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(elements, ";"), nil
}
