package spells

type GuidingBolt struct {}

func (g GuidingBolt) Cast(target string) bool {
    log.Println("The spell Guiding Bolt is not implemented yet")
}

func (g GuidingBolt) PrettyPrint() string {
    return "Guiding Bolt"
}
