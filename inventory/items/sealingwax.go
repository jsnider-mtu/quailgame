package items

type SealingWax struct {
}

func (s SealingWax) Slot() string {
    return ""
}

func (s SealingWax) Use() {
}

func (s SealingWax) Save() string {
    return "SealingWax"
}

func (s SealingWax) PrettyPrint() string {
    return "Sealing Wax"
}
