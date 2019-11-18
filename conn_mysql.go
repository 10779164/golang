package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

const (
    USERNAME = "root"
    PASSWD = "flasker0115"
    NETWORK = "tcp"
    SERVER = "localhost"
    PORT = 3306
    DATABASE = "test"
)

var db = &sql.DB{}

/*
func insertDb(DB) {
    time := time.Now()
    fmt.Println(time)
    text := "Test ceph rbd backup."
    fmt.Println(text)
    result, _ := DB.Exec("insert into test(id,time,text) value(2,?,?)",time,text) 
    n, _ := result.RowsAffected()
    fmt.Println("受影响的记录数是", n)
}
*/

func main() {
    //fmt.Println("test")
    dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWD,NETWORK,SERVER,PORT,DATABASE)
    db, err := sql.Open("mysql",dsn)
    if err != nil{
        fmt.Printf("Open mysql failed,err:%v\n",err)
        return
    }
    //db.SetMaxOpenConns(600)
    time := time.Now()
    text := "Test ceph rbd backup."
    tx,_ := db.Begin()
    tx.Exec("insert into test value('',?,?)",time,text)
    tx.Commit()
}
