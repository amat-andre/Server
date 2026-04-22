package service

import (
	"errors"
	"strings"

	"github.com/amat-andre/Server/pkg/morse"
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
