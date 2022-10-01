package spells

type Heroism struct {}

func (h Heroism) Cast(target string) bool {
    log.Println("The spell Heroism is not implemented yet")
}

func (h Heroism) PrettyPrint() string {
    return "Heroism"
}
