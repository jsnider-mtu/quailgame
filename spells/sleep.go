package spells

type Sleep struct {}

func (s Sleep) Cast(target string) bool {
    log.Println("The spell Sleep is not implemented yet")
}

func (s Sleep) PrettyPrint() string {
    return "Sleep"
}
