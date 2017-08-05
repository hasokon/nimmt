package nimmt

import (
	"testing"
)

func TestNewColumnAndString(t *testing.T) {
	column, err := NewColumn(Card{1, 4}, Card{34, 5})
	if err != nil {
		t.Error("NewColumn Error : " + err.Error())
	}

	if column.String() != "( 1 34 )" {
		t.Error(column.String() + " is Wrong Output")
	}
}

func TestMaxNumber(t *testing.T) {

}

func TestColmunAdd(t *testing.T) {
	column, err := NewColumn(Card{23, 3}, Card{25, 2}, Card{27, 3}, Card{30, 3})
	if err != nil {
		t.Error("NewColumn Err :" + err.Error())
	}

	num, ok := column.Add(Card{34, 1})
	if !ok {
		t.Errorf("Column Add Error. It has to be not over. Return number is %d\n", num)
	}

	num, ok = column.Add(Card{45, 2})
	if ok {
		t.Errorf("Column Add Error. It has to be over. Return number is %d\n", num)
	}
}

func TestNewFieldAndString(t *testing.T) {
	cards := []Card{
		{1, 4},
		{24, 5},
		{23, 5},
		{78, 5},
	}
	field, err := NewField(cards)
	if err != nil {
		t.Error("NewField Error : " + err.Error())
	}

	if field.String() != "[( 1 ),( 24 ),( 23 ),( 78 )]" {
		t.Error(field.String() + " is Wrong Output")
	}
}

func TestSort(t *testing.T) {
	cards := []Card{
		{12, 4},
		{79, 5},
		{45, 5},
		{91, 5},
	}
	field, err := NewField(cards)
	if err != nil {
		t.Error("NewField Error : " + err.Error())
	}

	field.Sort()

	if field.String() != "[( 91 ),( 79 ),( 45 ),( 12 )]" {
		t.Error(field.String() + " is Wrong Output")
	}
}

func TestFieldAdd(t *testing.T) {
	cards := []Card{
		{12, 4},
		{79, 5},
		{45, 5},
		{91, 5},
	}
	field, err := NewField(cards)
	if err != nil {
		t.Error("NewField Error : " + err.Error())
	}

	card := Card{56, 5}
	num, ok := field.Add(card)
	if !ok {
		t.Errorf("Field Add Error. Add Card is %v. Number : %d\n", card, num)
	}
	if field.String() != "[( 91 ),( 79 ),( 45 56 ),( 12 )]" {
		t.Error(field.String() + " is Wrong Output")
	}

	card = Card{85, 5}
	num, ok = field.Add(card)
	if !ok {
		t.Errorf("Field Add Error. Add Card is %v. Number : %d\n", card, num)
	}
	if field.String() != "[( 91 ),( 79 85 ),( 45 56 ),( 12 )]" {
		t.Error(field.String() + " is Wrong Output")
	}

	card = Card{83, 5}
	num, ok = field.Add(card)
	if !ok {
		t.Errorf("Field Add Error. Add Card is %v. Number : %d\n", card, num)
	}
	if field.String() != "[( 91 ),( 79 85 ),( 45 56 83 ),( 12 )]" {
		t.Error(field.String() + " is Wrong Output")
	}

	card = Card{84, 5}
	num, ok = field.Add(card)
	if !ok {
		t.Errorf("Field Add Error. Add Card is %v. Number : %d\n", card, num)
	}
	if field.String() != "[( 91 ),( 79 85 ),( 45 56 83 84 ),( 12 )]" {
		t.Error(field.String() + " is Wrong Output")
	}

	card = Card{90, 5}
	num, ok = field.Add(card)
	if !ok {
		t.Errorf("Field Add Error. Add Card is %v. Number : %d\n", card, num)
	}
	if field.String() != "[( 91 ),( 79 85 90 ),( 45 56 83 84 ),( 12 )]" {
		t.Error(field.String() + " is Wrong Output")
	}

	card = Card{87, 5}
	num, ok = field.Add(card)
	if !ok {
		t.Errorf("Field Add Error. Add Card is %v. Number : %d\n", card, num)
	}
	if field.String() != "[( 91 ),( 79 85 90 ),( 45 56 83 84 87 ),( 12 )]" {
		t.Error(field.String() + " is Wrong Output")
	}

	card = Card{89, 5}
	num, ok = field.Add(card)
	if ok {
		t.Errorf("Field Add Error. Add Card is %v. Number : %d\n", card, num)
	}
	if field.String() != "[( 91 ),( 79 85 90 ),( 89 ),( 12 )]" {
		t.Error(field.String() + " is Wrong Output")
	}

	card = Card{11, 3}
	num, ok = field.Add(card)
	if !ok {
		t.Errorf("Field Add Error. Add Card is %v. Number : %d\n", card, num)
	}
	if field.String() != "[( 91 ),( 79 85 90 ),( 89 ),( 11 )]" {
		t.Error(field.String() + " is Wrong Output\n")
	}
}
