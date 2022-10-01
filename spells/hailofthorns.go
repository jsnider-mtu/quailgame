package spells

type HailOfThorns struct {}

func (h HailOfThorns) Cast(target string) bool {
    log.Println("The spell Hail of Thorns is not implemented yet")
}

func (h HailOfThorns) PrettyPrint() string {
    return "Hail of Thorns"
}
