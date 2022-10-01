package spells

type Grease struct {}

func (g Grease) Cast(target string) bool {
    log.Println("The spell Grease is not implemented yet")
}

func (g Grease) PrettyPrint() string {
    return "Grease"
}
