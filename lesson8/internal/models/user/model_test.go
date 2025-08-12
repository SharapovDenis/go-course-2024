package user

import "testing"

func TestNew(t *testing.T) {
	u := New()

	if u.ID != 0 {
		t.Errorf("expected ID = 0, got %d", u.ID)
	}
	if u.Name != "" {
		t.Errorf("expected Name empty, got %q", u.Name)
	}
	if u.Email != "" {
		t.Errorf("expected Email empty, got %q", u.Email)
	}
}
