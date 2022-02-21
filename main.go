package main

import (
	"strconv"
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type MyData struct {
	ID int
	Name string
	Mail string
	Age int
}

func (m * MyData) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\"" + m.Mail + "," + strconv.Itoa(m.Age) + ">"
}

func main() {
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}

	defer con.Close()

	q := "select * from mydata"
	rs, er := con.Query(q)
	if er != nil {
		panic(er)
	}

	for rs.Next() {
		var md MyData
		er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		fmt.Println(md.Str())
	}

	
}
