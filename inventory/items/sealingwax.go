package items

type Sealingwax struct {
}

func (s Sealingwax) Slot() string {
    return ""
}

func (s Sealingwax) Use() {
}

func (s Sealingwax) Save() string {
    return "Sealingwax"
}

func (s Sealingwax) PrettyPrint() string {
    return "Sealingwax"
}
