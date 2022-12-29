package items

type Javelin struct {
}

func (j *Javelin) Slot() string {
    return "RightHand"
}

func (j *Javelin) Use() (string, []int) {
    return "", []int{}
    // must be equipped to use
}

func (j *Javelin) Save() string {
    return "Javelin"
}

func (j *Javelin) PrettyPrint() string {
    return "Javelin"
}

func (j *Javelin) Function() string {
    return "melee"
}

func (j *Javelin) Damage() (int, int, string) {
    return 1, 6, "piercing"
}

func (j *Javelin) Action() string {
    return ""
}
