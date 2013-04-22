package model

import (
    //"github.com/robfig/revel"
    "time"
    "webchat/app/form"
    //"fmt"
    "errors"
)

type Room struct {
    Id      int `pk`
    UserId  int 
    RoomKey string
    Title   string 
    Desc    string 
    CreatedAt time.Time
    UpdatedAt time.Time
}

func NewRoom(form *form.RoomForm) (room *Room) {
    room = &Room{
        UserId: form.UserId,
        RoomKey: form.RoomKey,
        Title: form.Title,
        Desc: form.Desc,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    return room
}

func (room *Room) Save() ( *Room, error){
    db := GetDblink()

    if err := room.ValidatesUniqueness(); err != nil {
        return nil,err
    }

    if err := db.Save(room); err != nil {
        return nil,err
    }

    return room, nil
}

func (room *Room) ValidatesUniqueness() error {
    db := GetDblink()
    var r Room
    if err := db.Where("roomid=?", room.RoomKey).Find(&r); err == nil {
        return errors.New("input room id: " + room.RoomKey+ " has exist")
    }
    return nil
}
