package items

type Arcanefocus struct {
}

func (a Arcanefocus) Slot() string {
    return ""
}

func (a Arcanefocus) Use() {
}

func (a Arcanefocus) Save() string {
    return "Arcanefocus"
}
