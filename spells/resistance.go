package spells

type Resistance struct {}

func (r Resistance) Cast(target string) bool {
    log.Println("The spell Resistance is not implemented yet")
}

func (r Resistance) PrettyPrint() string {
    return "Resistance"
}
