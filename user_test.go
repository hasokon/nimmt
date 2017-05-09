package nimmt

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	user := NewUser("test")

	if user.String() != "test [0 cows] ( )" {
		t.Error(user.String() + " is wrong Output")
	}
}

func TestUserBehavior(t *testing.T) {
	user := NewUser("test")

	user.TakeCows(3)
	if user.Cows != 3 {
		t.Error("TakeCows Error")
	}

	err := user.Draw(Card{55,5})
	if err != nil {
		t.Error(err.Error())
	}
	if user.String() != "test [3 cows] ( 55 )" {
		t.Error("Draw Error. " + user.String() + " is wrong output")
	}

	user.Draw(Card{33,5})
	user.Draw(Card{12,2})

	card, err := user.Put(55)
	if err != nil {
		t.Error(err.Error())
	}
	if card.Number != 55 {
		t.Error("Put Error")
	}

	if user.String() != "test [3 cows] ( 33 12 )" {
		t.Error("Put Error")
	}

	card, err = user.Put(56)
	if err == nil {
		t.Error("Put Error")
	}
}
