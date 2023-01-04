package items

type Flute struct {
}

func (f *Flute) Slot() string {
    return "BothHands"
}

func (f *Flute) Use() (string, []int) {
    return f.Action(), []int{}
}

func (f *Flute) Save() string {
    return "Flute"
}

func (f *Flute) PrettyPrint() string {
    return "Flute"
}

func (f *Flute) Function() string {
    return "instrument"
}

func (f *Flute) Damage() (int, int, string) {
    return 0, 0, ""
}

func (f *Flute) Action() string {
    return "playmusic"
}

func (f *Flute) GetQuantity() int {
    return 1
}

func (f *Flute) GetRange() []float64 {
    return []float64{0, 0}
}
