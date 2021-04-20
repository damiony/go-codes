package flyweight

import "testing"

func TestFlyWeight(t *testing.T) {
	u := NewUnsharableFlyweight()
	if u.Display() != "unsharable:sharable:a" {
		t.Error("error")
	}
}
