package goexpect_test

import (
	"github.com/drborges/goexpect"
	"testing"
)

func TestExpectStringEquality(t *testing.T) {
	expect := goexpect.New(t)

	expect("a").ToBe("a")
}

func TestExpectStructEquality(t *testing.T) {
	a := struct{ name string }{"diego"}
	b := struct{ name string }{"diego"}

	expect := goexpect.New(t)

	expect(a).ToBe(b)
}

func TestExpectDeepEquality(t *testing.T) {
	a := &struct{ name string }{"diego"}
	b := &struct{ name string }{"diego"}

	expect := goexpect.New(t)

	expect(a).ToDeepEqual(b)
}
