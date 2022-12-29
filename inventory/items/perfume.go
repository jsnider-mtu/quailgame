package items

type Perfume struct {
}

func (p *Perfume) Slot() string {
    return ""
}

func (p *Perfume) Use() (string, []int) {
    return "", []int{}
}

func (p *Perfume) Save() string {
    return "Perfume"
}

func (p *Perfume) PrettyPrint() string {
    return "Perfume"
}

func (p *Perfume) Function() string {
    return "disguise"
}

func (p *Perfume) Damage() (int, int, string) {
    return 0, 0, ""
}

func (p *Perfume) Action() string {
    return ""
}
