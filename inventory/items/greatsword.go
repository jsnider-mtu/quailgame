package items

type Greatsword struct {
}

func (g Greatsword) Slot() string {
    return "RightHand"
}

func (g Greatsword) Use() {
    // must be equipped to use
}

func (g Greatsword) Save() string {
    return "Greatsword"
}
