package spells

type InflictWounds struct {}

func (i InflictWounds) Cast(target string) bool {
    log.Println("The spell Inflict Wounds is not implemented yet")
}

func (i InflictWounds) PrettyPrint() string {
    return "Inflict Wounds"
}
