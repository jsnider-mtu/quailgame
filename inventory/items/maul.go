package items

type Maul struct {
}

func (m Maul) Slot() string {
    return "RightHand"
}

func (m Maul) Use() {
    // must be equipped to use
}

func (m Maul) Save() string {
    return "Maul"
}

func (m Maul) PrettyPrint() string {
    return "Maul"
}
