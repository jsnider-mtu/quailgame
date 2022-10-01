package spells

type Identify struct {}

func (i Identify) Cast(target string) bool {
    log.Println("The spell Identify is not implemented yet")
}
