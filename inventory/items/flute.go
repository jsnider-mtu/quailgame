package items

type Flute struct {
}

func (f Flute) Slot() string {
    return ""
}

func (f Flute) Use() {
}

func (f Flute) Save() string {
    return "Flute"
}

func (f Flute) PrettyPrint() string {
    return "Flute"
}
