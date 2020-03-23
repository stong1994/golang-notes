package mongo

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	empty := NewEmptySortCond()
	sort := empty.And(CondExpr("age", true)).And(CondExpr("name", false))
	fmt.Printf("%v\n%T\n", sort, sort)
}
