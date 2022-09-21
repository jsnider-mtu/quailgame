package items

type Clothes struct {
    quality string
}

func (c Clothes) Slot() string {
    return "Clothes"
}

func (c Clothes) Use() {
}

func (c Clothes) Save() string {
    return "Clothes," + c.quality
}
