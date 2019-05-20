package input

import (
	"strings"
)

type ParsingError []string

func (pe ParsingError) Error() string {
	return strings.Join([]string(pe), " ")
}

func AggErrors(pe ParsingError, s string) ParsingError {
	return ParsingError(append([]string(pe), s))
}
