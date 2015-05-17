package goexpect

import "reflect"

type subject interface{}
type expectFn func(interface{}) *expectation

type test interface {
	Errorf(string, ...interface{})
}

func New(t test) expectFn {
	return func (actual interface{}) *expectation {
		return &expectation{
			test: t,
			actual: actual,
		}
	}
}

type expectation struct {
	test test
	actual subject
}

func (this *expectation) ToBe(expected subject) {
	if this.actual != expected {
		this.test.Errorf("Expected %v, got %v", expected, this.actual)
	}
}

func (this *expectation) ToDeepEqual(expected subject) {
	if !reflect.DeepEqual(this.actual, expected) {
		this.test.Errorf("Expected %v, got %v", expected, this.actual)
	}
}
