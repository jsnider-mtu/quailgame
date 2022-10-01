package spells

type WrathfulSmite struct {}

func (w WrathfulSmite) Cast(target string) bool {
    log.Println("The spell Wrathful Smite is not implemented yet")
}

func (w WrathfulSmite) PrettyPrint() string {
    return "Wrathful Smite"
}
