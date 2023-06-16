package main

import (
	"reflect"
	"testing"
)

func TestBetween(t *testing.T) {
	got := Between(1, 5)
	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}
