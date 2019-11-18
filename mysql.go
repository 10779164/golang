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
    PASSWD = "flasker0115" //"9ztrkRLc"
    NETWORK = "tcp"
    SERVER = "localhost" //"mysql-2272-0.tripanels.com"
    PORT = 3306 //10009
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
	for {
        dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWD,NETWORK,SERVER,PORT,DATABASE)
        db, err := sql.Open("mysql",dsn)
        if err != nil{
            fmt.Printf("Open mysql failed,err:%v\n",err)
            return
        }
	text := "Test ceph rbd backup. Test application mysql~!"
	ctable := table
	//sql := "insert into "+ctable+" value('',"+ctime+","+text+")"
	DB,_ := db.Begin()
        //for {
	    for i := 0; i < 5000; i++ {
	        //DB.Exec("insert into ? value('',?,?)",ctable,ctime,text)
	        ctime := time.Now().Format(time.UnixDate)
	        sql := fmt.Sprintf("insert into %s value('','%s','%s')",ctable,ctime,text)
	        DB.Exec(sql)
	    }
            DB.Commit()
        }
}

func main() {
    go insert("test1")
    //go insert("test2")
    //go insert("test3")
    for i := 0; i < 10; i++ {
	time.Sleep(time.Second)
}
    
}
