package facade

import "testing"

func TestFacade(t *testing.T) {
	facade := NewFacade()
	if facade.method() != "subsystemA\nsubsystemB\n" {
		t.Error("error")
	}
}
