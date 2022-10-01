package items

type Sickle struct {
}

func (s Sickle) Slot() string {
    return "RightHand"
}

func (s Sickle) Use() {
    // must be equipped to use
}

func (s Sickle) Save() string {
    return "Sickle"
}

func (s Sickle) PrettyPrint() string {
    return "Sickle"
}
