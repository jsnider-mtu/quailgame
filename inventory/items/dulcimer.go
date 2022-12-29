package items

type Dulcimer struct {
}

func (d *Dulcimer) Slot() string {
    return "BothHands"
}

func (d *Dulcimer) Use() (string, []int) {
    return "", []int{}
}

func (d *Dulcimer) Save() string {
    return "Dulcimer"
}

func (d *Dulcimer) PrettyPrint() string {
    return "Dulcimer"
}

func (d *Dulcimer) Function() string {
    return "instrument"
}

func (d *Dulcimer) Damage() (int, int, string) {
    return 0, 0, ""
}

func (d *Dulcimer) Action() string {
    return ""
}

func (d *Dulcimer) GetQuantity() int {
    return 1
}
