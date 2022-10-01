package spells

type Thaumaturgy struct {}

func (t Thaumaturgy) Cast(target string) bool {
    log.Println("The spell Thaumaturgy is not implemented yet")
}

func (t Thaumaturgy) PrettyPrint() string {
    return "Thaumaturgy"
}
