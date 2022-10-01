package items

type WarPick struct {
}

func (w WarPick) Slot() string {
    return "RightHand"
}

func (w WarPick) Use() {
    // must be equipped to use
}

func (w WarPick) Save() string {
    return "WarPick"
}

func (w WarPick) PrettyPrint() string {
    return "War Pick"
}
