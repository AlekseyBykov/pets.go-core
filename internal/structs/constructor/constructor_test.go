package constructor

import "testing"

func TestNewUser(t *testing.T) {
	u := NewUser("SomeName", 25)

	if u.Name != "SomeName" {
		t.Errorf("expected Name 'SomeName', got '%s'", u.Name)
	}
	if u.Age != 25 {
		t.Errorf("expected Age 25, got %d", u.Age)
	}
}

func TestNewUserPtr(t *testing.T) {
	up := NewUserPtr("SomeName", 30)

	if up == nil {
		t.Fatal("expected non-nil pointer")
	}
	if up.Name != "SomeName" {
		t.Errorf("expected Name 'SomeName', got '%s'", up.Name)
	}
	if up.Age != 30 {
		t.Errorf("expected Age 30, got %d", up.Age)
	}
}
