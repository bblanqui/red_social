package bd

import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

)
/*MongoC este es el objeto de conexion a la db*/
var MongoC=ConectarBD()
var clientOptions=options.Client().ApplyURI("mongodb+srv://bblanqui:brian312@cluster0.8zeab.mongodb.net/twittor?retryWrites=true&w=majority")

/*ConnectarBD ES LA FUNCION QUE NOS PERMITE CONECTAR*/
func ConectarBD() *mongo.Client {
  client, err := mongo.Connect(context.TODO(), clientOptions)
  if err != nil {
      log.Fatal(err.Error())
      return client
  }
    
  err = client.Ping(context.TODO(), nil)
  
  if err != nil{
    log.Fatal(err.Error())
    return client
  }
  log.Println("conexion ok")

  return client
}
/*Chequeoconection es la prueba de la conexion*/
func ChequeoConection ()int {
    err := MongoC.Ping(context.TODO(), nil)
    if err != nil {
       
        return 0
    }
    return 1
}