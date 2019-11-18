package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "time"
    "os"
)

//create table test1(id int primary key auto_increment,time varchar(40),text varchar(120))auto_increment=1;

const (
    USERNAME = "root"
    NETWORK = "tcp"
    SERVER = "localhost"
    PORT = 3306
    DATABASE = "test"
)


//var db = &sql.DB{}


func init() {
	var PASSWD string
        PASSWD = os.Getenv("ROOT_PASS")
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWD,NETWORK,SERVER,PORT,DATABASE)
	db, err := sql.Open("mysql",dsn)
	if err != nil {
            fmt.Printf("Open mysql failed,err:%v\n",err)
	    return 
	}
	for i :=1; i <= 3; i++ {
	sql := fmt.Sprintf("create table if not exists test%d(id int primary key auto_increment,time varchar(40),text varchar(120))auto_increment=1",i)
	db.Exec(sql)
	}
}



func insert(table string) {
	var PASSWD string
        PASSWD = os.Getenv("ROOT_PASS")
        dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWD,NETWORK,SERVER,PORT,DATABASE)
        db, err := sql.Open("mysql",dsn)
        if err != nil{
            fmt.Printf("Open mysql failed,err:%v\n",err)
            return
        }
	for {
	ctext := "Test ceph rbd backup."
	ctable := table
	DB,_ := db.Begin()
	for i := 0; i < 3000; i++ {
	    ctime := time.Now().Format(time.UnixDate)
	    sql := fmt.Sprintf("insert into %s(time,text) value('%s','%s')",ctable,ctime,ctext)
	    DB.Exec(sql)
	    }
            DB.Commit()
	    time.Sleep(time.Second)
        }
}

func main() {
    go insert("test1")
    go insert("test2")
    go insert("test3")
    for i :=0; i < 600; i++{
	time.Sleep(time.Second)
}
}
