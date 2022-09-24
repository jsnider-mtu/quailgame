package items

type Shortsword struct {
}

func (s Shortsword) Slot() string {
    return "RightHand"
}

func (s Shortsword) Use() {
    // must be equipped to use
}

func (s Shortsword) Save() string {
    return "Shortsword"
}

func (s Shortsword) PrettyPrint() string {
    return "Shortsword"
}
