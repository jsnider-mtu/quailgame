package items

type Dulcimer struct {
}

func (d Dulcimer) Slot() string {
    return "BothHands"
}

func (d Dulcimer) Use() {
}

func (d Dulcimer) Save() string {
    return "Dulcimer"
}

func (d Dulcimer) PrettyPrint() string {
    return "Dulcimer"
}
