package spells

type EldritchBlast struct {}

func (e EldritchBlast) Cast(target string) bool {
    log.Println("The spell Eldritch Blast is not implemented yet")
}

func (e EldritchBlast) PrettyPrint() string {
    return "Eldritch Blast"
}
