package ads

import (
	"testing"
)

func TestNew(t *testing.T) {
	ad := New()

	// Проверяем, что время установлено
	if ad.CreatedAt.IsZero() {
		t.Errorf("CreatedAt должно быть инициализировано")
	}
	if ad.ModifiedAt.IsZero() {
		t.Errorf("ModifiedAt должно быть инициализировано")
	}

	// Проверяем, что остальные поля по умолчанию
	if ad.ID != 0 {
		t.Errorf("expected ID = 0, got %d", ad.ID)
	}
	if ad.Title != "" {
		t.Errorf("expected Title empty, got %q", ad.Title)
	}
	if ad.Text != "" {
		t.Errorf("expected Text empty, got %q", ad.Text)
	}
	if ad.AuthorID != 0 {
		t.Errorf("expected AuthorID = 0, got %d", ad.AuthorID)
	}
	if ad.Published {
		t.Errorf("Ожидали Published = false, получили true")
	}
}
