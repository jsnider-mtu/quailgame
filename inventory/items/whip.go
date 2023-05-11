package items

type Whip struct {
}

func (w *Whip) Slot() string {
    return "RightHand"
}

func (w *Whip) Use() (string, []int) {
    return w.Action(), []int{}
}

func (w *Whip) Save() string {
    return "Whip"
}

func (w *Whip) PrettyPrint() string {
    return "Whip"
}

func (w *Whip) Function() string {
    return "melee-finesse-reach"
}

func (w *Whip) Damage() (int, int, string) {
    return 1, 4, "slashing"
}

func (w *Whip) Action() string {
    return ""
}

func (w *Whip) GetQuantity() int {
    return 1
}

func (w *Whip) GetRange() []float64 {
    return []float64{0, 0}
}
