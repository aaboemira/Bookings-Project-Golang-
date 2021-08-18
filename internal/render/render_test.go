package render

import (
	"github.com/aaboemira/bookings/internal/models"
	"net/http"
	"testing"
)

func TestAddTempData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Fatal(err)
	}
	session.Put(r.Context(), "warning", "123")

	result := AddTempData(&td, r)
	if result.Warning != "123" {
		t.Error("warning value of 123 not found in session")
	}
}
func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}
