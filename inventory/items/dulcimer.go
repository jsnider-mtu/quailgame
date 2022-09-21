package items

type Dulcimer struct {
}

func (d Dulcimer) Slot() string {
    return ""
}

func (d Dulcimer) Use() {
}

func (d Dulcimer) Save() string {
    return "Dulcimer"
}
