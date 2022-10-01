package spells

type CureWounds struct {}

func (c CureWounds) Cast(target string) bool {
    log.Println("The spell Cure Wounds is not implemented yet")
}

func (c CureWounds) PrettyPrint() string {
    return "Cure Wounds"
}
