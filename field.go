package nimmt

import (
	"errors"
	"fmt"
	"sort"
)

const (
	COLUMNS = 4
	CARDS   = 5
)

type Column struct {
	Cards [CARDS]Card
	Count int // between 1 to CARDS
}

func NewColumn(cards ...Card) (*Column, error) {
	clen := len(cards)
	if clen < 1 {
		return nil, errors.New("Column must have at least One Card")
	}

	if clen > COLUMNS {
		clen = COLUMNS
	}

	column := Column{Count: clen}
	for i := 0; i < clen; i++ {
		column.Cards[i] = cards[i]
	}

	return &column, nil
}

func (c *Column) String() string {
	str := []byte{}
	for i := 0; i < c.Count; i++ {
		str = append(str, []byte(fmt.Sprintf(" %d", c.Cards[i].Number))...)
	}
	return fmt.Sprintf("(%s )", string(str))
}

func (c *Column) MaxNumber() int {
	return c.Cards[c.Count-1].Number
}

func (c *Column) Add(newCard Card) (int, bool) {
	if c.Count < CARDS {
		c.Cards[c.Count] = newCard
		c.Count++
		return 0, true
	}

	cows := 0
	for _, card := range c.Cards {
		cows += card.Cows
	}

	c.Cards[0] = newCard
	c.Count = 1

	return cows, false
}

func (c *Column) AllChange(newCard Card) int {
	cows := 0
	for _, card := range c.Cards {
		cows += card.Cows
	}

	c.Cards[0] = newCard
	c.Count = 1

	return cows
}

type Field struct {
	Columns []*Column
}

func (f *Field) String() string {
	str := []byte{}
	for i := 0; i < COLUMNS-1; i++ {
		str = append(str, []byte(f.Columns[i].String()+",")...)
	}

	str = append(str, []byte(f.Columns[COLUMNS-1].String())...)

	return fmt.Sprintf("[%s]", string(str))
}

func NewField(initialCards []Card) (*Field, error) {
	if len(initialCards) < COLUMNS {
		return nil, errors.New("New need at least 4 Cards")
	}

	field := new(Field)
	field.Columns = make([]*Column, COLUMNS)
	for i, _ := range field.Columns {
		var err error
		field.Columns[i], err = NewColumn(initialCards[i])
		if err != nil {
			return nil, err
		}
	}

	return field, nil
}

// 大きい順
func (f *Field) Sort() {
	sort.Slice(f.Columns, func(i, j int) bool {
		return f.Columns[i].MaxNumber() > f.Columns[j].MaxNumber()
	})
}

func (f *Field) Add(newCard Card) (int, bool) {
	f.Sort()
	for _, column := range f.Columns {
		if column.MaxNumber() < newCard.Number {
			return column.Add(newCard)
		}
	}
	return f.Columns[len(f.Columns)-1].AllChange(newCard), true
}
