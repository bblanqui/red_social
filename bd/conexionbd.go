package bd

import (
	//context es el contexto de la conexion hacia la base de datos
	"context"
	"log"

	////mongo driver es el manejador de la base de datos se usa para la coneccion el cual tiene el mogo client tiene una funcion llamada mogo connect lacual resibe 2 parametros el contexto y el clientoption
	"go.mongodb.org/mongo-driver/mongo"
	/// mongo options es el cual maneja  la url de la coneccion tiene una funcon llamada options.Cliente la cual resibe como parametro el ApplyUri resibe cimo paraetro la cadena de conexion
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC este es el objeto de conexion a la db*/
var MongoC = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://bblanqui:brian312@cluster0.8zeab.mongodb.net/twittor?retryWrites=true&w=majority")

/*ConnectarBD ES LA FUNCION QUE NOS PERMITE CONECTAR*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())

		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("conexion ok")
	return client
}

/*Chequeoconection es la prueba de la conexion*/
func ChequeoConection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {

		return 0
	}
	return 1
}
