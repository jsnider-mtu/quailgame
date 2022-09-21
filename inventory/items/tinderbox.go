package items

type Tinderbox struct {
}

func (t Tinderbox) Slot() string {
    return ""
}

func (t Tinderbox) Use() {
}

func (t Tinderbox) Save() string {
    return "Tinderbox"
}
