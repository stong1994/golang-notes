package rand_num

import "testing"

func TestRank(t *testing.T) {
	n := RandNum(10)
	if n >= 0 && n < 10 {
		t.Log("right")
	} else {
		t.Errorf("expect n >= 0 && n < 10, but get %d", n)
	}
}
