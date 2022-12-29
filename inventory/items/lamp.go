package items

type Lamp struct {
}

func (l Lamp) Slot() string {
    return "LeftHand"
}

func (l Lamp) Use() (string, []int) {
    return "illuminate", []int{15, 30, 3600}
}

func (l Lamp) Save() string {
    return "Lamp"
}

func (l Lamp) PrettyPrint() string {
    return "Lamp"
}

func (l Lamp) Function() string {
    return "light"
}

func (l Lamp) Damage() (int, int, string) {
    return 0, 0, ""
}
