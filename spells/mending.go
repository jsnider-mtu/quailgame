package spells

type Mending struct {}

func (m Mending) Cast(target string) bool {
    log.Println("The spell Mending is not implemented yet")
}

func (m Mending) PrettyPrint() string {
    return "Mending"
}
