package AddingSlave

import (
  "fmt"
  "time"
  "context"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type Slave struct{
  ip string
  senhaTelnet string
}


func connect_mongo(){
  fmt.Println("Mongodb")

  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))



}
