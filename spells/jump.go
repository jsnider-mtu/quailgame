package spells

type Jump struct {}

func (j Jump) Cast(target string) bool {
    log.Println("The spell Jump is not implemented yet")
}

func (j Jump) PrettyPrint() string {
    return "Jump"
}
