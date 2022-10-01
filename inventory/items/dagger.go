package items

type Dagger struct {
}

func (d Dagger) Slot() string {
    return "RightHand"
}

func (d Dagger) Use() {
    // must be equipped to use
}

func (d Dagger) Save() string {
    return "Dagger"
}

func (d Dagger) PrettyPrint() string {
    return "Dagger"
}
