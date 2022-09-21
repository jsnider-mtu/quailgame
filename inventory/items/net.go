package items

type Net struct {
}

func (n Net) Slot() string {
    return "RightHand"
}

func (n Net) Use() {
    // must be equipped to use
}

func (n Net) Save() string {
    return "Net"
}
