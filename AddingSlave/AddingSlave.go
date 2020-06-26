package main 

import (
  "fmt"
  "time"
  "context"

  "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type Slave struct{
  ID primitive.ObjectID 'bson:"_id,omitempty"'
  ip string 'bson:"ip,omitempty"'
  senha string 'bson:"senha,omitempty"'
}


func Connect_mongo() (*Client,*Database,*Context){
  fmt.Println("Conectando ao Mongodb")

  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
  if err != nil { panic(err) }
  database := client.Database("EDP-Botnet").Collection("Slave")
  return client,database,ctx
}

func Add_slave(slave *Slave, bd *Database,ctx *Context) (){
  res, err := bd.InsertOne(ctx, slave)
  if err != nil { panic(err) }
  fmt.Printf("O dispositivo foi inserido com o ID: %v",res.InsertedID)
}

func main(){
  exemplo:= Slave{
    ip: "127.0.0.1",
    senha: "12345"
  }
  client, bd, ctx := Connect_mongo()
  Add_slave(exemplo,bd,ctx)
  defer client.Disconnect(ctx)
}
