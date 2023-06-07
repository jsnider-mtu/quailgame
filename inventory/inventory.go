package inventory

import (
    "errors"
    "fmt"
    "strconv"
)

type Item interface {
    Slot() string
    Use() (string, []int)
    Save() string
    PrettyPrint() string
    Function() string
    Damage() (int, int, string)
    Action() string
    GetQuantity() int
    GetRange() []float64
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

func (i *Inv) Clear() {
    i.items = make([]Item, 0)
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

func (i *Inv) Save() string {
    var result string
    for _, val := range i.items {
        result += val.Save() + ";"
    }
    result += strconv.Itoa(i.seeds) + ";"
    return result
}
