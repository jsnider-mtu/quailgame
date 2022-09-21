package items

type Componentpouch struct {
}

func (c Componentpouch) Slot() string {
    return ""
}

func (c Componentpouch) Use() {
}

func (c Componentpouch) Save() string {
    return "Componentpouch"
}
