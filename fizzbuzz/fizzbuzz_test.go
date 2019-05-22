package fizzbuzz

import (
	"../spell"

	"fmt"
	"testing"
)

var n, m = 5, 7
var fizz, buzz = "hello", "world"

func TestConvertor(t *testing.T) {
	f := Convertor(n, m, fizz, buzz)
	e := func(x int, s string) string {
		return fmt.Sprintf(
			"Convertor(%d, %d, %s, %s)(%d) is %s; want %s.",
			n, m, fizz, buzz, x, f(x), s,
		)
	}
	if f(2*n) != fizz {
		t.Error(e(3*n, fizz))
	}
	if f(3*m) != buzz {
		t.Error(e(2*n, buzz))
	}
	if f(4*n*m) != fizz+buzz {
		t.Error(e(4*n*m, fizz+buzz))
	}
	if n != 1 && m != 1 && f(n*m+1) != spell.NumberInLetters(n*m+1) {
		t.Error(e(n*m+1, spell.NumberInLetters(n*m+1)))
	}
}

func TestList(t *testing.T) {
	var limit = n * m
	var list = List(limit, n, m, fizz, buzz)
	if len(list) != limit {
		t.Errorf(
			"List(%d, %d, %d, %s, %s) is on length %d; want %d.",
			limit, n, m, fizz, buzz, len(list), limit,
		)
	}
	f := Convertor(n, m, fizz, buzz)
	for i := 0; i < limit; i++ {
		if list[i] != f(i+1) {
			t.Errorf(
				"List(%d, %d, %d, %s, %s)[%d] is %s; want Convertor(%d, %d, %s, %s)(%d) = %s.",
				limit, n, m, fizz, buzz, i, list[i], n, m, fizz, buzz, i, f(i),
			)
		}
	}
}
