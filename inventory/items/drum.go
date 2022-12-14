package items

type Drum struct {
}

func (d *Drum) Slot() string {
    return "BothHands"
}

func (d *Drum) Use() (string, []int) {
    return d.Action(), []int{}
}

func (d *Drum) Save() string {
    return "Drum"
}

func (d *Drum) PrettyPrint() string {
    return "Drum"
}

func (d *Drum) Function() string {
    return "instrument"
}

func (d *Drum) Damage() (int, int, string) {
    return 0, 0, ""
}

func (d *Drum) Action() string {
    return "playmusic"
}

func (d *Drum) GetQuantity() int {
    return 1
}

func (d *Drum) GetRange() []float64 {
    return []float64{0, 0}
}
