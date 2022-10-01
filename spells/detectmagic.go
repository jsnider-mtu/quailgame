package spells

type DetectMagic struct {}

func (d DetectMagic) Cast(target string) bool {
    log.Println("The spell Detect Magic is not implemented yet")
}

func (d DetectMagic) PrettyPrint() string {
    return "Detect Magic"
}
