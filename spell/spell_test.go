package spell

import "testing"

func TestNumberInLetters(t *testing.T) {
	test := func(n int, s string) {
		if got := NumberInLetters(n); got != s {
			t.Errorf("NumberInLetters(%d) is %s; want %s.", n, got, s)
		}
	}

	test(0, "zero")
	test(1, "one")

	test(12, "twelve")
	test(23, "twenty-three")
	test(34, "thirty-four")
	test(45, "fourty-five")
	test(56, "fifty-six")
	test(67, "sixty-seven")
	test(78, "seventy-eight")
	test(89, "eighty-nine")
	test(90, "ninety")

	test(101, "one hundred and one")
	test(128, "one hundred and twenty-eight")
	test(256, "two hundred and fifty-six")

	test(1001, "one thousand and one")
	test(2010, "two thousand and ten")
	test(2048, "two thousand and fourty-eight")

	test(523288, "five hundred and twenty-three thousand two hundred and eighty-eight")

	test(1098304, "one million ninety-eight thousand three hundred and four")
}
