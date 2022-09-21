package items

type Clothes struct {
    Quality string
}

func (c Clothes) Slot() string {
    return "Clothes"
}

func (c Clothes) Use() {
}

func (c Clothes) Save() string {
    return "Clothes," + c.Quality
}
