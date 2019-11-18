package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

//create table test1(id int primary key auto_increment,time varchar(40),text varchar(120))auto_increment=1;

const (
    USERNAME = "root"
    PASSWD = "9ztrkRLc"
    NETWORK = "tcp"
    SERVER = "mysql-2272-0.tripanels.com"
    PORT = 10009
    DATABASE = "test"
)

//var db = &sql.DB{}

/*
func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWD,NETWORK,SERVER,PORT,DATABASE)
	db, err := sql.Open("mysql",dsn)
	if err != nil {
            fmt.Printf("Open mysql failed,err:%v\n",err)
	    return 
	}
	db.SetMaxOpenConns(600)
}
*/


func insert(table string) {
        dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWD,NETWORK,SERVER,PORT,DATABASE)
        db, err := sql.Open("mysql",dsn)
        if err != nil{
            fmt.Printf("Open mysql failed,err:%v\n",err)
            return
        }
	ctext := "Test ceph rbd backup."
	ctable := table
	//sql := "insert into "+ctable+" value('',"+ctime+","+text+")"
	DB,_ := db.Begin()
	        ctime := time.Now().Format(time.UnixDate)
	        sql := fmt.Sprintf("insert into %s(time,text) value('%s','%s')",ctable,ctime,ctext)
		fmt.Println(sql)
	        DB.Exec(sql)
            DB.Commit()
}

func main() {
    insert("test1")
}
