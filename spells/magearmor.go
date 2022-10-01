package spells

type MageArmor struct {}

func (m MageArmor) Cast(target string) bool {
    log.Println("The spell Mage Armor is not implemented yet")
}

func (m MageArmor) PrettyPrint() string {
    return "Mage Armor"
}
