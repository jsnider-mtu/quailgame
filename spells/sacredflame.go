package spells

type SacredFlame struct {}

func (s SacredFlame) Cast(target string) bool {
    log.Println("The spell Sacred Flame is not implemented yet")
}

func (s SacredFlame) PrettyPrint() string {
    return "Sacred Flame"
}
