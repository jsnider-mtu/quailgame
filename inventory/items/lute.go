package items

type Lute struct {
}

func (l Lute) Slot() string {
    return ""
}

func (l Lute) Use() {
}

func (l Lute) Save() string {
    return "Lute"
}

func (l Lute) PrettyPrint() string {
    return "Lute"
}
