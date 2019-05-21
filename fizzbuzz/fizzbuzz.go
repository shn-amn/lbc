package fizzbuzz

import "../spell"

func New(n, m int, fizz, buzz string) func(int) string {
	f := func(i int) string {
		switch rn, rm := i%n, i%m; {
		case rn == 0 && rm == 0:
			return fizz + buzz
		case rn == 0:
			return fizz
		case rm == 0:
			return buzz
		default:
			return spell.NumberInLetters(i)
		}
	}
	return f
}
