package spells

type TrueStrike struct {}

func (t TrueStrike) Cast(target string) bool {
    log.Println("The spell True Strike is not implemented yet")
}

func (t TrueStrike) PrettyPrint() string {
    return "True Strike"
}
