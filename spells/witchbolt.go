package spells

type WitchBolt struct {}

func (w WitchBolt) Cast(target string) bool {
    log.Println("The spell Witch Bolt is not implemented yet")
}

func (w WitchBolt) PrettyPrint() string {
    return "Witch Bolt"
}
