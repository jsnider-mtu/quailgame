package items

type LightHammer struct {
}

func (l LightHammer) Slot() string {
    return "RightHand"
}

func (l LightHammer) Use() {
    // must be equipped to use
}

func (l LightHammer) Save() string {
    return "LightHammer"
}

func (l LightHammer) PrettyPrint() string {
    return "Light Hammer"
}
