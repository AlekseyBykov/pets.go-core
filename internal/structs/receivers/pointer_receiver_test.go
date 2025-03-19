package receivers

import "testing"

func TestUser_PrintUserInfo(t *testing.T) {
	u := User{Name: "Original", Age: 25}

	u.PrintUserInfo()

	if u.Name != "Renamed" {
		t.Errorf("expected Name to be 'Renamed', got %s", u.Name)
	}
}
