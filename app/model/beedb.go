package model

import (
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
