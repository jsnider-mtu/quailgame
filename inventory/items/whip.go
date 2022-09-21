package items

type Whip struct {
}

func (w Whip) Slot() string {
    return "RightHand"
}

func (w Whip) Use() {
    // must be equipped to use
}

func (w Whip) Save() string {
    return "Whip"
}
