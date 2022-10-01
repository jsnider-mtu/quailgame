package items

type LightCrossbow struct {
}

func (l LightCrossbow) Slot() string {
    return "BothHands"
}

func (l LightCrossbow) Use() {
    // must be equipped to use
}

func (l LightCrossbow) Save() string {
    return "LightCrossbow"
}

func (l LightCrossbow) PrettyPrint() string {
    return "Light Crossbow"
}
