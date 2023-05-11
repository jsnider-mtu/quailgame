package items

type Quarterstaff struct {
    Currslot string
}

func (q *Quarterstaff) GetCurrSlot() string {
    return q.Currslot
}

func (q *Quarterstaff) SwitchSlots() {
    if q.Currslot == "BothHands" {
        q.Currslot = "RightHand"
    } else if q.Currslot == "RightHand" {
        q.Currslot = "BothHands"
    }
    return
}

func (q *Quarterstaff) Slot() string {
    return "Versatile"
}

func (q *Quarterstaff) Use() (string, []int) {
    return q.Action(), []int{1, 8}
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
    return "versatile"
}

func (q *Quarterstaff) GetQuantity() int {
    return 1
}

func (q *Quarterstaff) GetRange() []float64 {
    return []float64{0, 0}
}
