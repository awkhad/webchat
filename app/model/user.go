package model

import (
    //"github.com/robfig/revel"
    "time"
    "webchat/app/form"
    "crypto/sha1"
    "fmt"
    "math/rand"
    "errors"
)

type User struct {
    Id int `pk`
    Name string
    Email string 
    Salt string 
    Encryptpasswd string
    Created time.Time
    Updated time.Time
}

// generate a User form input form field
func NewUser(form *form.UserForm) (user *User){
    user = new(User)
    user.Id = 0
    user.Name = form.Name
    user.Email = form.Email
    user.Salt = generate_salt()
    user.Encryptpasswd = encryptPassword(form.Password, user.Salt)
    user.Created, user.Updated = time.Now(),time.Now()

    return user
}

func (user *User) ValidatesUniqueness() bool{
    db := GetDblink()
    var u User
    err := db.Where("name=? or email=?", user.Name, user.Email).Find(&u)
    // exist item 
    if err == nil {
        return false
    }

    return true
}

func (user *User) Save() error {
    db := GetDblink()

    if !user.ValidatesUniqueness()  {
        return errors.New("user name or email is exist")
    }

    err := db.Save(user)

    if err != nil {
        return err
    }
    return nil
}

// auth user
func Authenticate(name string, password string) bool {
    db := GetDblink()
    var user User
    err := db.Where("name=?", name).Find(&user)

    if err != nil {
        return false 
    }

    if (encryptPassword(password, user.Salt) == user.Encryptpasswd) {
        return true
    }
    return false
} 

// for generate rand salt 
func generate_salt() string {
    rand.Seed( time.Now().UnixNano())
    salt := fmt.Sprintf("%x", rand.Int31())
    return salt
}

// for enrypt password 
func encryptPassword(password, salt string) string {
    h := sha1.New()
    h.Write([]byte(password + salt))
    bs := fmt.Sprintf("%x",h.Sum(nil))
    return bs
}

