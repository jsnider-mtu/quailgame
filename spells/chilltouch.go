package spells

type ChillTouch struct {}

func (c ChillTouch) Cast(target string) bool {
    log.Println("The spell Chill Touch is not implemented yet")
}

func (c ChillTouch) PrettyPrint() string {
    return "Chill Touch"
}
