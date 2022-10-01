package spells

type Hex struct {}

func (h Hex) Cast(target string) bool {
    log.Println("The spell Hex is not implemented yet")
}

func (h Hex) PrettyPrint() string {
    return "Hex"
}
