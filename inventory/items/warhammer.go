package items

type Warhammer struct {
    Currslot string
}

func (w *Warhammer) GetCurrSlot() {
    return w.Currslot
}

func (w *Warhammer) SwitchSlots() {
    if w.Currslot == "BothHands" {
        w.Currslot = "RightHand"
    } else if w.Currslot == "RightHand" {
        w.Currslot = "BothHands"
    }
    return
}

func (w *Warhammer) Slot() string {
    return "Versatile"
}

func (w *Warhammer) Use() (string, []int) {
    return w.Action(), []int{1, 10}
}

func (w *Warhammer) Save() string {
    return "Warhammer"
}

func (w *Warhammer) PrettyPrint() string {
    return "Warhammer"
}

func (w *Warhammer) Function() string {
    return "melee"
}

func (w *Warhammer) Damage() (int, int, string) {
    return 1, 8, "bludgeoning"
}

func (w *Warhammer) Action() string {
    return "versatile"
}

func (w *Warhammer) GetQuantity() int {
    return 1
}

func (w *Warhammer) GetRange() []float64 {
    return []float64{0, 0}
}
