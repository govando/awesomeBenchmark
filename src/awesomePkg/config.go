package main

import (
	"gopkg.in/mgo.v2/bson"
)

var (
 	host		= "127.0.0.1:27017"
	db			= "test"
	collCommm	= "emptyColl"
	collBench	= "awsColl"
	nClientes 	= 10
)

var (
	tamannos	= []uint32{64, 100, 400, 700, 1000} //bytes
	n_pruebas	= 30
)


type Doc struct {
	ID		bson.ObjectId `bson:"_id,omitempty"`
	Data	string
	Cmp1	int
	Cmp2	int
	Cmp3	int
}

//se usa como Validador de la coleccion de Test Communication
//permite rechazar todas las operaicones update/insert del tipo Doc
var (
	invalidDoc	=  bson.M{"b": bson.M{"$exists": true}}
)