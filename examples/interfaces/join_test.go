package main

import "testing"

func TestJoinStringers(t *testing.T) {
	strs := []Stringer{
		str("foo"),
		str("bar"),
		str("baz"),
	}
	exp := "foo, bar, baz"
	if act := JoinStringers(strs, ", "); act != exp {
		t.Errorf("JoinStringer() works incorrect: expected %q, found %q", exp, act)
	}
}

type str string

func (s str) String() string {
	return string(s)
}
