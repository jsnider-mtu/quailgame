package items

type HandCrossbow struct {
}

func (h HandCrossbow) Slot() string {
    return "RightHand"
}

func (h HandCrossbow) Use() {
    // must be equipped to use
}

func (h HandCrossbow) Save() string {
    return "HandCrossbow"
}

func (h HandCrossbow) PrettyPrint() string {
    return "Hand Crossbow"
}
