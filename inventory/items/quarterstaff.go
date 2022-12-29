package items

type Quarterstaff struct {
}

func (q *Quarterstaff) Slot() string {
    return "BothHands"
}

func (q *Quarterstaff) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (q *Quarterstaff) Save() string {
    return "Quarterstaff"
}

func (q *Quarterstaff) PrettyPrint() string {
    return "Quarterstaff"
}

func (q *Quarterstaff) Function() string {
    return "melee"
}

func (q *Quarterstaff) Damage() (int, int, string) {
    return 1, 6, "bludgeoning"
}

func (q *Quarterstaff) Action() string {
    return ""
}

func (q *Quarterstaff) GetQuantity() int {
    return 1
}
