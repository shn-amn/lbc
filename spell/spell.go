package spell

import (
	"strings"
	"errors"
	"fmt"
)

func NumberInLetters(n int) string {
	if n == 0 {
		return zero
	}
	var parts = []string{subthousandInLetters(n % 1000, n > 1000)}
	biggie := 1000
	for n /= 1000; n > 0; n /= 1000 {
		r := n % 1000
		switch name, ok := biggies[biggie]; {
		case !ok:
			panic(errors.New("I can't yet count this hight!"))
		case r == 0:
		default:
			parts = append(parts, fmt.Sprintf("%s %s", subthousandInLetters(r, false), name))
		}
		biggie *= 1000
	}
	return strings.Join(reverse(parts), " ")
}

func subthousandInLetters(n int, keepAnd bool) string{
	switch {
	case n == 0:
		panic(errors.New("0 should be treated separately."))
	case n < 0:
		panic(errors.New(fmt.Sprintf("%d is a negative number.", n)))
	case n > 999:
		panic(errors.New(fmt.Sprintf("%d is greater than a thousand.", n)))
	case n < 100 && keepAnd:
		return fmt.Sprintf("and %s", subhundredInLetters(n))
	case n < 100:
		return subhundredInLetters(n)
	case n % 100 == 0:
		return fmt.Sprintf("%s hundred", fingers[n / 100])
	default:
		return fmt.Sprintf("%s hundred and %s", fingers[n / 100], subhundredInLetters(n % 100))
	}
}

func subhundredInLetters(n int) string {
	switch {
	case n == 0:
		panic(errors.New("0 should be treated separately."))
	case n < 0:
		panic(errors.New(fmt.Sprintf("%d is a negative number.", n)))
	case n > 99:
		panic(errors.New(fmt.Sprintf("%d is greater than a hundred.", n)))
	case n % 10 == 0:
		return tens[n / 10]
	case n < 10:
		return fingers[n]
	case n < 20:
		return toes[n]
	default:
		return fmt.Sprintf("%s-%s", tens[n / 10], fingers[n % 10])
	}
}

func reverse(strs []string) []string {
	n := len(strs)
	rvrs := make([]string, n)
	for i, s := range strs {
		rvrs[n - i - 1] = s
	}
	return rvrs
}
