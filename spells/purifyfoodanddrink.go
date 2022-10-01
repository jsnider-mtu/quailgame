package spells

type PurifyFoodAndDrink struct {}

func (p PurifyFoodAndDrink) Cast(target string) bool {
    log.Println("The spell Purify Food and Drink is not implemented yet")
}

func (p PurifyFoodAndDrink) PrettyPrint() string {
    return "Purify Food and Drink"
}
