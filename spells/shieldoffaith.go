package spells

type ShieldOfFaith struct {}

func (s ShieldOfFaith) Cast(target string) bool {
    log.Println("The spell Shield of Faith is not implemented yet")
}

func (s ShieldOfFaith) PrettyPrint() string {
    return "Shield of Faith"
}
