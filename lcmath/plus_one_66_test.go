package lcmath_test

import (
	"fmt"
	lcmath "golc/lcmath"
	"reflect"
	"testing"
)

func TestBestDemo(t *testing.T) {

	if !reflect.DeepEqual(lcmath.PlusOne([]int{8, 9, 9, 9}), []int{9, 0, 0, 0}) {
		fmt.Println(lcmath.PlusOne([]int{8, 9, 9, 9}))
		t.Fatal("0.Arrays not equal")
	}

	if !reflect.DeepEqual(lcmath.PlusOne([]int{1, 2, 3}), []int{1, 2, 4}) {
		t.Fatal("1.Arrays not equal")
	}

	if !reflect.DeepEqual(lcmath.PlusOne([]int{9, 9, 9}), []int{1, 0, 0, 0}) {
		t.Fatal("2.Arrays not equal")
	}

	if !reflect.DeepEqual(lcmath.PlusOne([]int{9}), []int{1, 0}) {
		t.Fatal("3.Arrays not equal")
	}
}
