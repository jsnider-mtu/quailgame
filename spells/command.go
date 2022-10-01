package spells

type Command struct {}

func (c Command) Cast(target string) bool {
    log.Println("The spell Command is not implemented yet")
}
