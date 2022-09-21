package items

type Lamp struct {
}

func (l Lamp) Slot() string {
    return ""
}

func (l Lamp) Use() {
}

func (l Lamp) Save() string {
    return "Lamp"
}
