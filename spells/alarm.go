package spells

type Alarm struct {}

func (a Alarm) Cast(target string) bool {
    log.Println("The spell Alarm is not implemented yet")
}

func (a Alarm) PrettyPrint() string {
    return "Alarm"
}
