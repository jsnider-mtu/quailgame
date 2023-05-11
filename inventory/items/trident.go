package items

type Trident struct {
    Currslot string
}

func (t *Trident) GetCurrSlot() string (
    return t.Currslot
}

func (t *Trident) SwitchSlots() {
    if t.Currslot == "BothHands" {
        t.Currslot = "RightHand"
    } else if t.Currslot == "RightHand" {
        t.Currslot = "BothHands"
    }
    return
}

func (t *Trident) Slot() string {
    return "Versatile"
}

func (t *Trident) Use() (string, []int) {
    return t.Action(), []int{96, 288}
}

func (t *Trident) Save() string {
    return "Trident"
}

func (t *Trident) PrettyPrint() string {
    return "Trident"
}

func (t *Trident) Function() string {
    return "melee-throw"
}

func (t *Trident) Damage() (int, int, string) {
    return 1, 6, "piercing"
}

func (t *Trident) Action() string {
    return "throw"
}

func (t *Trident) GetQuantity() int {
    return 1
}

func (t *Trident) GetRange() []float64 {
    return []float64{96.0, 288.0}
}
