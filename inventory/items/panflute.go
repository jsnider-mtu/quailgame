package items

type Panflute struct {
}

func (p Panflute) Slot() string {
    return ""
}

func (p Panflute) Use() {
}

func (p Panflute) Save() string {
    return "Panflute"
}

func (p Panflute) PrettyPrint() string {
    return "Panflute"
}
