package spells

type MagicMissile struct {}

func (m MagicMissile) Cast(target string) bool {
    log.Println("The spell Magic Missile is not implemented yet")
}

func (m MagicMissile) PrettyPrint() string {
    return "Magic Missile"
}
