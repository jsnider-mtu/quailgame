package spells

type FaerieFire struct {}

func (f FaerieFire) Cast(target string) bool {
    log.Println("The spell Faerie Fire is not implemented yet")
}

func (f FaerieFire) PrettyPrint() string {
    return "Faerie Fire"
}
