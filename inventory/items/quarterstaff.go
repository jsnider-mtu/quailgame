package items

type Quarterstaff struct {
}

func (q Quarterstaff) Slot() string {
    return "RightHand"
}

func (q Quarterstaff) Use() {
    // must be equipped to use
}

func (q Quarterstaff) Save() string {
    return "Quarterstaff"
}
