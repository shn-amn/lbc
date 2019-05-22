package fizzbuzz

import "strconv"

func Convertor(n, m int, fizz, buzz string) func(int) string {
	f := func(i int) string {
		switch rn, rm := i%n, i%m; {
		case rn == 0 && rm == 0:
			return fizz + buzz
		case rn == 0:
			return fizz
		case rm == 0:
			return buzz
		default:
			return strconv.Itoa(i)
		}
	}
	return f
}

func List(limit, n, m int, fizz, buzz string) []string {
	f := Convertor(n, m, fizz, buzz)
	list := make([]string, limit)
	for i := range list {
		list[i] = f(i + 1)
	}
	return list
}
