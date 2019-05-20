package input

import (
	"errors"
	"fmt"
	"strconv"
)

type Parameters struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

func ParseParams(queries map[string][]string) (params Parameters, pe ParsingError) {
	var e error
	if params.Int1, e = validateInt(queries, "int1"); e != nil {
		pe = AggErrors(pe, e.Error())
	}
	if params.Int2, e = validateInt(queries, "int2"); e != nil {
		pe = AggErrors(pe, e.Error())
	}
	if params.Limit, e = validateInt(queries, "limit"); e != nil {
		pe = AggErrors(pe, e.Error())
	}
	if params.Str1, e = validateStr(queries, "str1"); e != nil {
		pe = AggErrors(pe, e.Error())
	}
	if params.Str2, e = validateStr(queries, "str2"); e != nil {
		pe = AggErrors(pe, e.Error())
	}
	return
}

func validateStr(q map[string][]string, s string) (string, error) {
	arr, ok := q[s]
	if !ok || len(arr) == 0 {
		return "", errors.New(fmt.Sprintf("Parameter '%s' is not specified.", s))
	}
	return arr[0], nil
}

func validateInt(q map[string][]string, s string) (int, error) {
	str, err := validateStr(q, s)
	if err != nil {
		return 0, err
	}
	i, e := strconv.Atoi(str)
	if e != nil || i < 1 {
		return i, errors.New(fmt.Sprintf("Parameter '%s' must be a positive integer but provided value '%s' is not.", s, str))
	}
	return i, nil
}
