package nimmt

import (
	"fmt"
	"errors"
)

type User struct {
	Name  string
	Cows  int
	Hands []Card
}

const (
	HANDS = 10
)

func (u *User) String() string {
	cards := []byte{}
	for _, v := range u.Hands {
		cards = append(cards, []byte(" " + v.Number)...)
	}
	return fmt.Sprintf("%s [ %d cows] (%s )",u.Name, u.Cows, cards)
}

func NewUser(name string) *User {
	return &User{
		Name : name,
		Cows : 0,
		Hands : make([]Card, 0),
	}
}

func (u *User) TakeCows(cows int) {
	u.Cows += cows
}

func (u *User) Draw(card Card) error {
	if len(u.Hands) >= HANDS {
		return errors.New(fmt.Sprintf("This User <%s> has already %d cards",u.Name, HANDS))
	}

	u.Hands = append(u.Hands, card)
	return nil
}

func (u *User) Put(number int) (Card ,error) {
	for _, card := range u.Hands {
		if card.Number == number {
			u.Hands = deleteCard(u.Hands, number)
			return card, nil
		}
	}

	return Card{}, errors.New("The number " + number + " is not in this hands")
}

func deleteCard(hands []Card, card int) []Card {
	newCards := make([]Card, 0)
	for _, v := range hands {
		if v.Number != card {
			newCards = append(newCards, v)
		}
	}

	return newCards
}
