package spells

type CompelledDuel struct {}

func (c CompelledDuel) Cast(target string) bool {
    log.Println("The spell Compelled Duel is not implemented yet")
}

func (c CompelledDuel) PrettyPrint() string {
    return "Compelled Duel"
}
