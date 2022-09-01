package inventory

import (
    "errors"
    "fmt"
)

type Item interface {
    Slot() string
    Use()
    Save() string
}

type Inv struct {
    items []Item
    seeds int
}

func (i *Inv) GetSeeds() int {
    return i.seeds
}

func (i *Inv) GetItems() []Item {
    return i.items
}

func (i *Inv) Add(item Item) error {
    i.items = append(i.items, item)
    return nil
}

func (i *Inv) Drop(item Item) {
    for a, b := range i.items {
        if b == item {
            i.items = append(i.items[:a], i.items[a + 1:]...)
            break
        }
    }
    return
}

func (i *Inv) AddSeeds(amount int) {
    i.seeds += amount
    return
}

func (i *Inv) SubtractSeeds(amount int) error {
    if newamount := i.seeds - amount; newamount < 0 {
        return errors.New(fmt.Sprintf("%d - %d = %d; Insufficient funds", i.seeds, amount, newamount))
    }
    i.seeds -= amount
    return nil
}
