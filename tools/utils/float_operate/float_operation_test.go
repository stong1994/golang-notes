package float

import (
	"fmt"
	"testing"
)

func TestFloatOpFactory_CreateOpFactory(t *testing.T) {
	type checkFunc func(float64) error
	testFunc := func(want float64) checkFunc {
		return func(get float64) error {
			if want != get {
				return fmt.Errorf("wanted %f, but got %f", want, get)
			}
			return nil
		}
	}

	tests := []struct {
		name   string
		opName string
		a      float64
		b      float64
		check  checkFunc
	}{
		{
			"add",
			"+",
			12.12,
			0.2,
			testFunc(12.32),
		},
		{
			"add",
			"+",
			2.32,
			0.222,
			testFunc(2.542),
		},
		{
			"sub",
			"-",
			2.3,
			5.6,
			testFunc(-3.3),
		},
		{
			"sub",
			"-",
			2.33,
			1.111,
			testFunc(1.219),
		},
		{
			"div",
			"/",
			4.2,
			2.1,
			testFunc(2),
		},
		{
			"div",
			"/",
			5.0,
			4.0,
			testFunc(1.25),
		},
		{
			"mul",
			"*",
			4.2,
			2.1,
			testFunc(8.82),
		},
		{
			"mul",
			"*",
			1.9,
			4.0,
			testFunc(7.6),
		},
	}

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			op := CreateOpFactory(ts.opName)
			data := op.Calc(ts.a, ts.b)
			if err := ts.check(data); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestFloatOpFactory_MultiCalc(t *testing.T) {
	type checkFunc func(float64) error

	testFunc := func(want float64) checkFunc {
		return func(get float64) error {
			if get != want {
				return fmt.Errorf("expected %f, founded %f", want, get)
			}
			return nil
		}
	}
	tests := [...]struct {
		name   string
		params []interface{}
		checks checkFunc
	}{
		{
			"two add",
			[]interface{}{0.3, "+", 0.4},
			testFunc(0.7),
		},
		{
			"two sub",
			[]interface{}{0.3, "-", 0.4},
			testFunc(-0.1),
		},
		{
			"two mul",
			[]interface{}{0.3, "*", 0.4},
			testFunc(0.12),
		},
		{
			"two div",
			[]interface{}{0.3, "/", 0.4},
			testFunc(0.75),
		},
		{
			"three params",
			[]interface{}{0.3, "+", 0.4, "-", 0.12},
			testFunc(0.58),
		},
		{
			"four params",
			[]interface{}{0.3, "+", 0.4, "-", 0.2, "*", 0.6},
			testFunc(0.3),
		},
		{
			"five params",
			[]interface{}{0.3, "+", 0.4, "-", 0.2, "*", 0.6, "/", 1.2},
			testFunc(0.25),
		},
		{
			"six params",
			[]interface{}{0.3, "+", 0.4, "-", 0.2, "*", 0.6, "/", 1.2, "+", 3.5},
			testFunc(3.75),
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Factor of %s", tc.name), func(t *testing.T) {
			data := MultiCalc(tc.params...)
			if err := tc.checks(data); err != nil {
				t.Error(err)
			}
		})
	}
}
