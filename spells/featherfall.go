package spells

type FeatherFall struct {}

func (f FeatherFall) Cast(target string) bool {
    log.Println("The spell Feather Fall is not implemented yet")
}

func (f FeatherFall) PrettyPrint() string {
    return "Feather Fall"
}
