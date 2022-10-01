package spells

type SpeakWithAnimals struct {}

func (s SpeakWithAnimals) Cast(target string) bool {
    log.Println("The spell Speak with Animals is not implemented yet")
}

func (s SpeakWithAnimals) PrettyPrint() string {
    return "Speak with Animals"
}
