package nimmt

import (
	"fmt"
	"errors"
	)

type Room struct {
	Name string
	ID int
	Users     []User
	PlayField *Field
	Round     int
	IsPlaying bool
}

const (
	USER_LIM = 10
	)

func NewRoom() *Room {
	return &Room{
		Users : make([]User,0),
		PlayField : nil,
		Round : 0,
		IsPlaying : false,
	}
}

func (r *Room) String() string {
	status := ""
	switch(r.IsPlaying) {
		case true : status = "Playing"
		case false : status = "Not Playing"
	}
	return fmt.Sprintf("Room<Name: %s, ID: %d> %d users status:%s", r.Name, r.ID, len(r.Users), status)
}

func (r *Room) AddUser(user User) error {
	if r.IsPlaying {
		return errors.New("This Room has already started a game")
	}

	if len(r.Users) >= USER_LIM {
		return errors.New("This Room has already been full")
	}

	r.Users = append(r.Users, user)
	return nil
}

func (r *Room) DeleteUserByID(userId int) {
	origin := r.Users
	r.Users = make([]User, 0)

	for _, user := range origin {
		if user.ID != userId {
			r.Users = append(r.Users, user)
		}
	}
}

func (r *Room) IsPlayable() bool {
	if r.IsPlaying {
		return false
	}

	if len(r.Users) < 2 {
		return false
	}

	return true
}

func (r *Room) Start() error {
	if r.IsPlayable() == false {
		return errors.New("This room is not playable")
	}

	r.IsPlaying = true
	return nil
}
