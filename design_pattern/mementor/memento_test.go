package memento

import "testing"

func TestMemento(t *testing.T) {
	g := newGameState()
	g.update(12, 34)
	if g.hp != 12 || g.mp != 34 {
		t.Fatal("update error")
	}

	caretaker := newMementoManagers()
	m := g.Save()
	caretaker.setMemento(m.name, m)
	g.update(34, 45)
	if g.hp != 34 || g.mp != 45 {
		t.Fatal("update error")
	}

	g.Load(caretaker.getMemento(m.name))
	if g.hp != 12 || g.mp != 34 {
		t.Fatal("memento error")
	}
}
