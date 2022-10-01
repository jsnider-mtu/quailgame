package spells

type DetectEvilAndGood struct {}

func (d DetectEvilAndGood) Cast(target string) bool {
    log.Println("The spell Detect Evil and Good is not implemented yet")
}

func (d DetectEvilAndGood) PrettyPrint() string {
    return "Detect Evil and Good"
}
