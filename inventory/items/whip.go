package items

type Whip struct {
}

func (w *Whip) Slot() string {
    return "RightHand"
}

func (w *Whip) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (w *Whip) Save() string {
    return "Whip"
}

func (w *Whip) PrettyPrint() string {
    return "Whip"
}

func (w *Whip) Function() string {
    return "melee"
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
