package main

import (

	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"

	"os"
	_"time"
	_"fmt"
	"strconv"
	"strings"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	num_cliente int //identificador
}

func (c Client) testComm_emptyCount()  {

	//borrar y crear el archivo de datos
	f, err := os.Create( "./data/commEmpty/dataEmptyCount"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		commTest_emptyCount(size,f)
	}

}

func (c Client) testComm_emptyFind()  {

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/commEmpty/dataEmptyFindFind"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		commTest_emptyFind(size,f)
	}

}

func (c Client) testComm_emptyUpdate()  {

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/commEmpty/dataEmptyUpdate"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		commTest_emptyUpdate(size,f)
	}

}

func  (c Client) testComm_emptyDelete()  {

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/commEmpty/dataEmptyDelete"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		commTest_emptyDelete(size,f)
	}
}


/*------------------------------------------------*/

func (c Client) testComm_emptyInsert()  {

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/commEmpty/dataEmptyInsert"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		commTest_emptyInsert(size,f)
	}
}


func  (c Client) InsertOne(){

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/insert/dataInsert"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		insert(query_bson, f, size,coll)
	}

}

func  (c Client) FindOne(){

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/insert/dataEmptyInsert"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		commTest_emptyInsert(size,f)
	}

}

func (c Client) UpdateOne(){

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/insert/dataEmptyInsert"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		commTest_emptyInsert(size,f)
	}

}

func (c Client) DeleteOne(){

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/insert/dataEmptyInsert"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		commTest_emptyInsert(size,f)
	}

}




/*
func testComm_bulkFind()  {

	f, err := os.Create("./data/commEmptyCount/data")
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		commTest_emptyBulkFind()
	}

}
*/

