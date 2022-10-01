package spells

type MinorIllusion struct {}

func (m MinorIllusion) Cast(target string) bool {
    log.Println("The spell Minor Illusion is not implemented yet")
}

func (m MinorIllusion) PrettyPrint() string {
    return "Minor Illusion"
}
