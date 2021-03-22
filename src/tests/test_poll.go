package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	//"net"
	//"strconv"
	//"github.com/silenceper/pool"
)

type add struct {
	getinfo string
}

func main() {

	db, err := sql.Open("mysql", "root:17dayup2017@(192.168.1.19:3306)/")
	fmt.Println(db)
	fmt.Println(err)
	fmt.Println(reflect.TypeOf(db))

	/*
		mChan := make(chan string)

		factory := func() (interface{}, error) {
			return net.Dial("tcp", "127.0.0.1:9000")
		}

		close := func(v interface{}) error { return v.(net.Conn).Close() }
		poolConfig := &pool.PoolConfig{}
	*/
	//a := &add{}
	//fmt.Println(a.getinfo == "")
	/*
		for i := 0; i < 10000; i++ {
			str = "This is " + strconv.Itoa(i)
			fmt.Println(str)
			mChan <- str
		}

		for {
			gstr := <-mChan
			fmt.Println("Get ", gstr)
		}
	*/
}
