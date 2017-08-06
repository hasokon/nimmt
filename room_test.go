package nimmt

import (
//	"fmt"
	"testing"
)

func TestRoomString(t *testing.T) {
	tr := NewRoom("TestRoom", 0)
	if tr.String() != "Room<Name: TestRoom, ID: 0> 0 users status:Not Playing" {
		t.Error(tr.String() + " is Wrong Output\n")
	}
}

func TestAddandDeleteUserInRoom(t *testing.T) {
	tr := NewRoom("TestRoom", 0)
	taro := NewUser("Taro", 1)
	hanako := NewUser("hanako", 2)
	err := tr.AddUser(taro)
	if err != nil {
		t.Error(err)
	}
	err = tr.AddUser(hanako)
	if err != nil {
		t.Error(err)
	}

	if tr.String() != "Room<Name: TestRoom, ID: 0> 2 users status:Not Playing" {
		t.Error(tr.String() + " is Wrong Output\n")
	}

	for i := 0; i < 8; i++ {
		err = tr.AddUser(NewUser("hoge", i+3))
		if err != nil {
			t.Error(err)
		}
	}

	if tr.String() != "Room<Name: TestRoom, ID: 0> 10 users status:Not Playing" {
		t.Error(tr.String() + " is Wrong Output\n")
	}

	err = tr.AddUser(NewUser("Test", 12))
	if err == nil {
		t.Errorf("Add User Error. A Number of Users in  Room is over the border 10\n")
	}

	tr.DeleteUserByID(1)
	if tr.String() != "Room<Name: TestRoom, ID: 0> 9 users status:Not Playing" {
		t.Error(tr.String() + " is Wrong Output\n")
	}

	tr.DeleteUserByID(13)
	if tr.String() != "Room<Name: TestRoom, ID: 0> 9 users status:Not Playing" {
		t.Error(tr.String() + " is Wrong Output\n")
	}
}

func TestIsPlayable(t *testing.T) {
	tr := NewRoom("TestRoom", 0)
	if tr.IsPlayable == true {
		t.Error("This Room is no Playable because there are no players\n")
	}
	tr.AddUser(NewUser("Taro",0)
	tr.AddUser(NewUser("Hanako", 1)
	if tr.IsPlayable == false {
		t.Error("This Room is Playable")
	}
}
