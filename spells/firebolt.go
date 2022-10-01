package spells

type FireBolt struct {}

func (f FireBolt) Cast(target string) bool {
    log.Println("The spell Fire Bolt is not implemented yet")
}

func (f FireBolt) PrettyPrint() string {
    return "Fire Bolt"
}
