package spells

type DissonantWhispers struct {}

func (d DissonantWhispers) Cast(target string) bool {
    log.Println("The spell Dissonant Whispers is not implemented yet")
}

func (d DissonantWhispers) PrettyPrint() string {
    return "Dissonant Whispers"
}
