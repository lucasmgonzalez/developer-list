package resources

import "testing"
import "net/http/httptest"

func TestNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	NotFound(w, httptest.NewRequest("GET", "/not-found", nil))
	if w.Body.String() != "Not could found this resource!" {
		t.Errorf("Not found test fail, result: %v, expected: Not could found this resource!", w.Body.String())
	}
}
