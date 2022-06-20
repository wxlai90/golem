package golem

import "testing"

func TestNew(t *testing.T) {
	router := New()

	if router == nil {
		t.Errorf("New() router is nil")
	}
}

func TestNewSubRouter(t *testing.T) {
	router := NewSubRouter()

	if router == nil {
		t.Errorf("New() router is nil")
	}
}
