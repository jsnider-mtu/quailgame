package spells

type ThunderousSmite struct {}

func (t ThunderousSmite) Cast(target string) bool {
    log.Println("The spell Thunderous Smite is not implemented yet")
}

func (t ThunderousSmite) PrettyPrint() string {
    return "Thunderous Smite"
}
