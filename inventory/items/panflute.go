package items

type PanFlute struct {
}

func (p *PanFlute) Slot() string {
    return "BothHands"
}

func (p *PanFlute) Use() (string, []int) {
    return p.Action(), []int{}
}

func (p *PanFlute) Save() string {
    return "PanFlute"
}

func (p *PanFlute) PrettyPrint() string {
    return "Pan Flute"
}

func (p *PanFlute) Function() string {
    return "instrument"
}

func (p *PanFlute) Damage() (int, int, string) {
    return 0, 0, ""
}

func (p *PanFlute) Action() string {
    return "playmusic"
}

func (p *PanFlute) GetQuantity() int {
    return 1
}

func (p *PanFlute) GetRange() []float64 {
    return []float64{0, 0}
}
