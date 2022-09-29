package spells

type Spell interface {
    Cast(string) bool
}
