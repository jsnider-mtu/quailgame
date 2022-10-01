package spells

type FindFamiliar struct {}

func (f FindFamiliar) Cast(target string) bool {
    log.Println("The spell Find Familiar is not implemented yet")
}

func (f FindFamiliar) PrettyPrint() string {
    return "Find Familiar"
}
