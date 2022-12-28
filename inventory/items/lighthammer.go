package items

type LightHammer struct {
}

func (l LightHammer) Slot() string {
    return "RightHand"
}

func (l LightHammer) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (l LightHammer) Save() string {
    return "LightHammer"
}

func (l LightHammer) PrettyPrint() string {
    return "Light Hammer"
}

func (l LightHammer) Function() string {
    return "melee"
}

func (l LightHammer) Damage() (int, int, string) {
    return 1, 4, "bludgeoning"
}
