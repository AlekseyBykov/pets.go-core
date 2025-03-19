package custom_type

import "testing"

func TestAge_IsAdult(t *testing.T) {
	fixtures := []struct {
		age      Age
		expected bool
	}{
		{17, false},
		{18, true},
		{25, true},
	}

	for _, fixture := range fixtures {
		if res := fixture.age.IsAdult(); res != fixture.expected {
			t.Errorf("Age %d: expected %v, got %v",
				fixture.age, fixture.expected, res)
		}
	}
}
