package spells

type EnsnaringStrike struct {}

func (e EnsnaringStrike) Cast(target string) bool {
    log.Println("The spell Ensnaring Strike is not implemented yet")
}

func (e EnsnaringStrike) PrettyPrint() string {
    return "Ensnaring Strike"
}
