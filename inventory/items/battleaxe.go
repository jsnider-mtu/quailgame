package items

type Battleaxe struct {
}

func (b Battleaxe) Slot() string {
    return "RightHand"
}

func (b Battleaxe) Use() {
    // must be equipped to use
}

func (b Battleaxe) Save() string {
    return "Battleaxe"
}

func (b Battleaxe) PrettyPrint() string {
    return "Battleaxe"
}
