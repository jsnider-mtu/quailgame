package items

type Scimitar struct {
}

func (s Scimitar) Slot() string {
    return "RightHand"
}

func (s Scimitar) Use() {
    // must be equipped to use
}

func (s Scimitar) Save() string {
    return "Scimitar"
}
