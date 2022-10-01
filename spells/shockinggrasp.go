package spells

type ShockingGrasp struct {}

func (s ShockingGrasp) Cast(target string) bool {
    log.Println("The spell Shocking Grasp is not implemented yet")
}

func (s ShockingGrasp) PrettyPrint() string {
    return "Shocking Grasp"
}
