package generics

import "testing"

func TestSumInt(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result := Sum(nums)

	if result != 15 {
		t.Errorf("expected 15, got %d", result)
	}
}

func TestSumFloat(t *testing.T) {
	nums := []float64{1.5, 2.5, 3.0}
	result := Sum(nums)

	if result != 7.0 {
		t.Errorf("expected 7.0, got %f", result)
	}
}
