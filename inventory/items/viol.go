package items

type Viol struct {
}

func (v Viol) Slot() string {
    return "BothHands"
}

func (v Viol) Use() {
}

func (v Viol) Save() string {
    return "Viol"
}

func (v Viol) PrettyPrint() string {
    return "Viol"
}
