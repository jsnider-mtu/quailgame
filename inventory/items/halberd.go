package items

type Halberd struct {
}

func (h Halberd) Slot() string {
    return "RightHand"
}

func (h Halberd) Use() {
    // must be equipped to use
}

func (h Halberd) Save() string {
    return "Halberd"
}

func (h Halberd) PrettyPrint() string {
    return "Halberd"
}
