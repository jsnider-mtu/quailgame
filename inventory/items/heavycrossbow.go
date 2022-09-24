package items

type Heavycrossbow struct {
}

func (h Heavycrossbow) Slot() string {
    return "RightHand"
}

func (h Heavycrossbow) Use() {
    // must be equipped to use
}

func (h Heavycrossbow) Save() string {
    return "Heavycrossbow"
}

func (h Heavycrossbow) PrettyPrint() string {
    return "Heavycrossbow"
}
