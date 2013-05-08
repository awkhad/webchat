package model

import (
	//"github.com/robfig/revel"
	"errors"
	"fmt"
	"time"
	"webchat/app/form"
)

const (
	PageSize int = 12
)

type Room struct {
	Id          int `pk`
	UserId      int
	RoomKey     string
	Title       string
	Private     bool
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func AllRoom() []Room {
	var rooms []Room
	db := GetDblink()
	db.FindAll(&rooms)
	return rooms
}

func RoomCount() int {
	db := GetDblink()
	var itemCount int

	err := db.Db.QueryRow("select count(*) as count from room").Scan(&itemCount)

	if err != nil {
		panic(err)
	}

	fmt.Println("itemCount is :", itemCount)
	return itemCount
}

func FindRoomByUserId(user_id int) []Room {
	db := GetDblink()
	var rooms []Room
	fmt.Println("userid is--:", user_id)

	db.Where("user_id=?", user_id).FindAll(&rooms)
	fmt.Println("myrooms is --:", rooms)
	return rooms
}

func FindOnePage(p int) []Room {
	var rooms []Room
	db := GetDblink()
	var offset int
	if p == 0 {
		offset = 0
	} else {
		offset = (p - 1) * PageSize
	}
	//fmt.Println(p)
	db.Limit(PageSize, offset).FindAll(&rooms)
	return rooms
}

func FindRoomByRoomKey(rk string) *Room {
	db := GetDblink()
	var room Room

	if err := db.Where("room_key=?", rk).Find(&room); err != nil {
		panic(err)
	}

	return &room
}

func NewRoom(form *form.RoomForm) (room *Room) {
	room = &Room{
		UserId:      form.UserId,
		RoomKey:     form.RoomKey,
		Title:       form.Title,
		Private:     form.Private,
		Description: form.Desc,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return room
}

func (room *Room) Save() (*Room, error) {
	db := GetDblink()

	if err := room.ValidatesUniqueness(); err != nil {
		return nil, err
	}

	if err := db.Save(room); err != nil {
		return nil, err
	}

	return room, nil
}

func (room *Room) ValidatesUniqueness() error {
	db := GetDblink()
	var r Room
	db.Where("room_key=?", room.RoomKey).Find(&r)
	if r.Id != 0 {
		return errors.New("input room id: " + room.RoomKey + " has exist")
	}
	return nil
}
