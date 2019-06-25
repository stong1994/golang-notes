package tools

import "testing"

func TestFloatOpFactory_CreateOpFactory(t *testing.T) {
	addOperator := CreateOpFactory("+")
	result := addOperator.Calc(12.12, 0.2)
	if result != 12.32 {
		t.Fatalf("shoule be 12.32, but get %f", result)
	}

	result = addOperator.Calc(2.32, 0.222)
	if result != 2.542 {
		t.Fatalf("shoule be 2.542, but get %f", result)
	}

	subOperator := CreateOpFactory("-")
	result = subOperator.Calc(2.3, 5.6)
	if result != -3.3 {
		t.Fatalf("shoule be 2.542, but get %f", result)
	}
	result = subOperator.Calc(2.33, 1.111)
	if result != 1.219 {
		t.Fatalf("shoule be 1.219, but get %f", result)
	}

	divOperator := CreateOpFactory("/")
	result = divOperator.Calc(4.2, 2.1)
	if result != 2 {
		t.Fatalf("shoule be 2, but get %f", result)
	}
	result = divOperator.Calc(5.0, 4.0)
	if result != 1.25 {
		t.Fatalf("shoule be 1.25, but get %f", result)
	}

	mulOperator := CreateOpFactory("*")
	result = mulOperator.Calc(4.2, 2.1)
	if result != 8.82 {
		t.Fatalf("shoule be 8.82, but get %f", result)
	}
	result = mulOperator.Calc(1.9, 4.0)
	if result != 7.6 {
		t.Fatalf("shoule be 7.6, but get %f", result)
	}
	t.Log("success")
}
