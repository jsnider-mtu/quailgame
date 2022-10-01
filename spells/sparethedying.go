package spells

type SpareTheDying struct {}

func (s SpareTheDying) Cast(target string) bool {
    log.Println("The spell Spare the Dying is not implemented yet")
}

func (s SpareTheDying) PrettyPrint() string {
    return "Spare the Dying"
}
