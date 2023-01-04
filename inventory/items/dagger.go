package items

type Dagger struct {
}

func (d *Dagger) Slot() string {
    return "RightHand"
}

func (d *Dagger) Use() (string, []int) {
    return d.Action(), []int{96, 288}
}

func (d *Dagger) Save() string {
    return "Dagger"
}

func (d *Dagger) PrettyPrint() string {
    return "Dagger"
}

func (d *Dagger) Function() string {
    return "melee-throw"
}

func (d *Dagger) Damage() (int, int, string) {
    return 1, 4, "piercing"
}

func (d *Dagger) Action() string {
    return "throw"
}

func (d *Dagger) GetQuantity() int {
    return 1
}

func (d *Dagger) GetRange() []float64 {
    return []float64{96.0, 288.0}
}
