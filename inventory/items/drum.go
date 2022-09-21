package items

type Drum struct {
}

func (d Drum) Slot() string {
    return ""
}

func (d Drum) Use() {
}

func (d Drum) Save() string {
    return "Drum"
}
