package main

import (
	"gopkg.in/mgo.v2/bson"
)

var (
 	host			= "127.0.0.1:27017"
	db				= "test"
	collCommm		= "emptyColl"
	collBench		= "awsColl"
	nClientes 		= 1
	Operacion		= ""
	bytesReserved	= 55
)

var (
	tamannos	=  []uint32{100, 400, 700, 1000, 4000, 7000, 10000, 40000, 70000, 100000, 400000, 700000, 800000, 900000, 1000000, 2000000, 3000000, 4000000, 5000000, 6000000, 7000000, 8000000, 9000000, 10000000, 11000000, 12000000, 13000000, 14000000, 15000000, 16000000} //bytes
	n_pruebas	= 404
	idColl		= 0
	cantDatosIniciales = 400
	//CantPruebasRandom = 400
	nombrePruebas= []string{"CommCount","InsertCreateColl", "InsertCreateDoc","InsertErrorPK","Find","FindId","Update","UpdateId","Delete","DeleteId","DeleteStatic","DeleteIdStatic"}

)


type Doc struct {
	//ID		bson.ObjectId `bson:"_id,omitempty"`
	ID		int `bson:"_id,omitempty"`
	Data	string
	Cmp1	int
	Cmp2	int
	Cmp3	int
}

type Header struct {
	Tamannos []uint32 `json:"tamannos"`
	N_pruebas int `json:"n_pruebas"`
}

//se usa como Validador de la coleccion de Test Communication
//permite rechazar todas las operaicones update/insert del tipo Doc
var (
	invalidDoc	=  bson.M{"b": bson.M{"$exists": true}}
)

