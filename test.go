package main 

import (
    "fmt" 
    "time"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "github.com/astaxie/beedb"
)

func GetDblink() beedb.Model{
    db, err := sql.Open("mysql", "root:admin@/webchat_dev?charset=utf8")

    if err != nil {
        panic(err)
    }

    beedb.OnDebug = true
    orm := beedb.New(db)

    return orm
}

type User struct {
    Id int `pk`
    Name string
    Email string 
    Salt string 
    Encryptpasswd string
    Created time.Time
    Updated time.Time
}

func (user *User) Save() bool {
    db := GetDblink()
    fmt.Println(user)
    err := db.Save(user)

    if err != nil {
        return true
    }
    return false
}

func NewUser() (user *User){
    user = new(User)
    user.Name = "hello"
    user.Email = "hello@world.com"
    user.Salt = "asdf"
    user.Encryptpasswd = "asdfasdf"
    user.Created, user.Updated = time.Now(),time.Now()

    return user
}

func main() {
    db := GetDblink()
    user := NewUser()

    fmt.Println(user)

    db.Save(user)
}

