package spells

type ThornWhip struct {}

func (t ThornWhip) Cast(target string) bool {
    log.Println("The spell Thorn Whip is not implemented yet")
}

func (t ThornWhip) PrettyPrint() string {
    return "Thorn Whip"
}
