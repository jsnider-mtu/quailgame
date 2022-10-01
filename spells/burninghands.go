package spells

type BurningHands struct {}

func (b BurningHands) Cast(target string) bool {
    log.Println("The spell Burning Hands is not implemented yet")
}

func (b BurningHands) PrettyPrint() string {
    return "Burning Hands"
}
