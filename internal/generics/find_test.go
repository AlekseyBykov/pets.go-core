package generics

import "testing"

func TestFindInt(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	idx, ok := Find(arr, 3)
	if !ok || idx != 2 {
		t.Errorf("expected to find 3 at index 2, got idx=%d, ok=%v", idx, ok)
	}
}

func TestFindString(t *testing.T) {
	arr := []string{"go", "core", "pets"}

	idx, ok := Find(arr, "core")
	if !ok || idx != 1 {
		t.Errorf("expected to find 'core' at index 1, got idx=%d, ok=%v", idx, ok)
	}
}

func TestFindCustomStruct(t *testing.T) {
	type Item struct {
		ID int
	}

	arr := []Item{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}

	needle := Item{ID: 2}
	idx, ok := Find(arr, needle)
	if !ok || idx != 1 {
		t.Errorf("expected to find Item{ID:2} at index 1, got idx=%d, ok=%v", idx, ok)
	}
}
