package items

type Lute struct {
}

func (l *Lute) Slot() string {
    return "BothHands"
}

func (l *Lute) Use() (string, []int) {
    return l.Action(), []int{}
}

func (l *Lute) Save() string {
    return "Lute"
}

func (l *Lute) PrettyPrint() string {
    return "Lute"
}

func (l *Lute) Function() string {
    return "instrument"
}

func (l *Lute) Damage() (int, int, string) {
    return 0, 0, ""
}

func (l *Lute) Action() string {
    return "playmusic"
}

func (l *Lute) GetQuantity() int {
    return 1
}

func (l *Lute) GetRange() []float64 {
    return []float64{0, 0}
}
