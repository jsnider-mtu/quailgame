package items

type Inkbottle struct {
}

func (i Inkbottle) Slot() string {
    return ""
}

func (i Inkbottle) Use() {
}

func (i Inkbottle) Save() string {
    return "Inkbottle"
}

func (i Inkbottle) PrettyPrint() string {
    return "Inkbottle"
}
