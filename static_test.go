package golem

import (
	"net/http"
	"testing"
)

func TestStatic(t *testing.T) {
	r := New()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic")
		}
	}()

	r.Static("/", http.Dir("."))
}
