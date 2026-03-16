package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var errInput = errors.New("input value error")

func DefinitionAndConversion (input string) (string, error) {
	if input == "" {
		return "", errInput
	}
	
	symbols := "-. "
	if strings.Trim(input, symbols) != "" {
		output := morse.ToMorse(input)
		return output, nil
	}

	output := morse.ToText(input)
	return output, nil
}