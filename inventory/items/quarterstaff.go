package items

type Quarterstaff struct {
}

func (q Quarterstaff) Slot() string {
    return "BothHands"
}

func (q Quarterstaff) Use() {
    // must be equipped to use
}

func (q Quarterstaff) Save() string {
    return "Quarterstaff"
}

func (q Quarterstaff) PrettyPrint() string {
    return "Quarterstaff"
}
