package items

type Candles struct {
    quantity int
}

func (c Candles) Slot() string {
    return ""
}

func (c Candles) Use() {
}

func (c Candles) Save() string {
    return "Candles"
}
