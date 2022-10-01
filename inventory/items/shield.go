package items

type Shield struct {
}

func (s Shield) Slot() string {
    return "LeftHand"
}

func (s Shield) Use() {
}

func (s Shield) Save() string {
    return "Shield"
}

func (s Shield) PrettyPrint() string {
    return "Shield"
}
