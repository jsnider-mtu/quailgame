package spells

type Entangle struct {}

func (e Entangle) Cast(target string) bool {
    log.Println("The spell Entangle is not implemented yet")
}

func (e Entangle) PrettyPrint() string {
    return "Entangle"
}
