package template

import "testing"

func TestTemplate(t *testing.T) {
	phone := NewPhone()
	computer := NewComputer()

	if phone.Print() != "phone games:phone music" {
		t.Error("wrong phone implementor")
	}

	if computer.Print() != "computer games:computer music" {
		t.Error("wrong computer implementor")
	}
}
