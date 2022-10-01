package spells

type Druidcraft struct {}

func (d Druidcraft) Cast(target string) bool {
    log.Println("The spell Druidcraft is not implemented yet")
}

func (d Druidcraft) PrettyPrint() string {
    return "Druidcraft"
}
