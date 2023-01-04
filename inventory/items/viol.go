package items

type Viol struct {
}

func (v *Viol) Slot() string {
    return "BothHands"
}

func (v *Viol) Use() (string, []int) {
    return "", []int{}
}

func (v *Viol) Save() string {
    return "Viol"
}

func (v *Viol) PrettyPrint() string {
    return "Viol"
}

func (v *Viol) Function() string {
    return "instrument"
}

func (v *Viol) Damage() (int, int, string) {
    return 0, 0, ""
}

func (v *Viol) Action() string {
    return ""
}

func (v *Viol) GetQuantity() int {
    return 1
}

func (v *Viol) GetRange() []float64 {
    return []float64{0, 0}
}
