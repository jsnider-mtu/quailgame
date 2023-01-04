package items

type Flail struct {
}

func (f *Flail) Slot() string {
    return "RightHand"
}

func (f *Flail) Use() (string, []int) {
    return f.Action(), []int{}
}

func (f *Flail) Save() string {
    return "Flail"
}

func (f *Flail) PrettyPrint() string {
    return "Flail"
}

func (f *Flail) Function() string {
    return "melee"
}

func (f *Flail) Damage() (int, int, string) {
    return 1, 8, "bludgeoning"
}

func (f *Flail) Action() string {
    return ""
}

func (f *Flail) GetQuantity() int {
    return 1
}

func (f *Flail) GetRange() []float64 {
    return []float64{0, 0}
}
