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

func (c Clothes) PrettyPrint() string {
    if c.Quality == "Costume" {
        return c.Quality
    } else {
        return c.Quality + " Clothes"
    }
}
