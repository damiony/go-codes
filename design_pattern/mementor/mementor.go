package memento

// gameState 发起人
type gameState struct {
	hp, mp int
}

func (g *gameState) update(h, m int) {
	g.hp = h
	g.mp = m
}

func (g *gameState) Save() *gameStateMemento {
	return &gameStateMemento{
		name: "gameState",
		hp:   g.hp,
		mp:   g.mp,
	}
}

func (g *gameState) Load(i interface{}) {
	gm, ok := i.(*gameStateMemento)
	if ok {
		g.hp = gm.hp
		g.mp = gm.mp
	}
}

func newGameState() *gameState {
	return &gameState{}
}

// gameStateMemento 备忘录
type gameStateMemento struct {
	name   string
	hp, mp int
}

// mementoManagers 备忘录管理者
type mementoManagers struct {
	m map[string]interface{}
}

func (mm *mementoManagers) setMemento(name string, i interface{}) {
	mm.m[name] = i
}

func (mm *mementoManagers) getMemento(name string) interface{} {
	return mm.m[name]
}

func newMementoManagers() *mementoManagers {
	return &mementoManagers{m: make(map[string]interface{})}
}
