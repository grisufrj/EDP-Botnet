package AddingSlave 

import (
  "fmt"
  "time"
  "context"

  "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

/*type Slave struct{
  ID primitive.ObjectID 'bson:"_id,omitempty"'
  ip string 'bson:"ip,omitempty"'
  senha string 'bson:"senha,omitempty"'
}*/


func Connect_mongo() (*mongo.Client,*mongo.Collection,context.Context){
  fmt.Println("Conectando ao Mongodb")

  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
  if err != nil { panic(err) }
  database := client.Database("EDP-Botnet").Collection("Slave")
  return client,database,ctx
}

func Add_slave(slave primitive.D, bd *mongo.Collection,ctx context.Context) (){
  res, err := bd.InsertOne(ctx, slave)
  if err != nil { panic(err) }
  fmt.Printf("O dispositivo foi inserido com o ID: %v\n",res.InsertedID)
}

func Get_slavebd(bd *mongo.Collection, ctx context.Context)([]bson.M){
  var result []bson.M
  cursor, _ :=bd.Find(ctx,bson.M{})
  if err:= cursor.All(ctx,&result); err!= nil{panic(err)}
  return result

}


