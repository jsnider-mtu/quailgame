package spells

type CharmPerson struct {}

func (c CharmPerson) Cast(target string) bool {
    log.Println("The spell Charm Person is not implemented yet")
}

func (c CharmPerson) PrettyPrint() string {
    return "Charm Person"
}
