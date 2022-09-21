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
