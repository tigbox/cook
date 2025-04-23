package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	list := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	if !reflect.DeepEqual(TwoSum(list, target), expected) {
		t.Errorf("Expected %v, got %v", expected, TwoSum(list, target))
	}
}

func TestTwoSumHash(t *testing.T) {
	list := []int{2, 11, 7, 15}
	target := 9
	expected := []int{0, 2}
	result := TwoSumHash(list, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestTwoSumHashSort(t *testing.T) {
	list := []int{2, 11, 7, 15}
	target := 9
	expected := []int{0, 2}
	result := TwoSumHash1(list, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
