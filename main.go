package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type File struct {
	Value []byte
}

type Table struct {
	Files map[string]File
}

type Database struct {
	Name   string
	Tables map[string]*Table
}

func Create(name string) Database {
	var Db Database
	Db.Name = name
	Db.Tables = make(map[string]*Table)
	fmt.Println("Database " + name + " created")
	return Db
}

func Connect(name string) Database {
	file, _ := os.Open(name)
	defer file.Close()
	var db Database
	dec := gob.NewDecoder(file)
	file.Seek(0, 0)
	dec.Decode(&db)

	fmt.Println("Connected to database: " + name)
	return db
}

func (D *Database) Close() {
	file, _ := os.Create(D.Name)
	defer file.Close()
	enc := gob.NewEncoder(file)
	enc.Encode(&D)
	fmt.Println("db " + D.Name + " closed and saved to disk")
}

func (D *Database) CreateTable(table string) {
	var T Table
	T.Files = make(map[string]File)
	D.Tables[table] = &T
	fmt.Println("Table " + table + " created")
}

func (D *Database) Insert(table string, key string, value []byte) {
	var F File = File{value}
	D.Tables[table].Files[key] = F
	fmt.Println("File " + key + " inserted")
}

func (D *Database) Delete(table string, key string) {
	delete(D.Tables[table].Files, key)
	fmt.Println("File " + key + " deleted from table " + table)
}

func (D *Database) Update(table string, key string, newvalue []byte) {
	var F File = File{newvalue}
	D.Tables[table].Files[key] = F
	fmt.Println("File " + key + " updated")
}

func main() {
	/*
		mydb := Create("testDB")
		mydb.CreateTable("Table")
		mydb.Insert("Table", "HELLO", []byte("valuehello"))
		mydb.Insert("Table", "WORLD", []byte("valueworld"))
		fmt.Println(mydb.Tables["Table"])
		mydb.Delete("Table", "HELLO")
		mydb.Close()
		mydb2 := Connect("testDB")
		fmt.Println(mydb2.Tables["Table"])
	*/
	testTolstoy()
}

func testTolstoy() {
	file, _ := ioutil.ReadFile("warandpeace")
	db := Create("db")
	db.CreateTable("tbl")

	for i := 0; i < 20; i++ {
		start := time.Now()
		db.Insert("tbl", "copy"+strconv.Itoa(i), file)
		fmt.Println(time.Since(start))
	}
	//fmt.Println(db.Tables["tbl"].Files)
	s := time.Now()
	db.Close()
	fmt.Println(time.Since(s))
}
