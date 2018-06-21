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
	"gopkg.in/mgo.v2"
)

//Benchmark de latencia de comunicación
//operacion + colecion (comm o bench) + size request + request
type Client struct {
	num_cliente int //identificador
}

func (c Client) testComm_emptyCount()  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)

	//borrar y crear el archivo de datos
	f, err := os.Create( "./data/Comm/count/dataCommCount"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		for i := 0; i < n_pruebas; i++ {
			total_time := count_op(query_bson , f, size, collCommm, conn)
			times = append(times, float64(total_time)/float64(1000000))
		}
		uso_tiempos(times) //agregar metodos aquí
	}
}

func (c Client) testComm_emptyFind()  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)
	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Comm/find/dataCommFind"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		for i := 0; i < n_pruebas; i++ {
			total_time := find_op(query_bson , f, size, collCommm, conn)
			times = append(times, float64(total_time)/float64(1000000))
		}
		uso_tiempos(times) //agregar metodos aquí
	}
}

func (c Client) testComm_emptyFindId()  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)
	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Comm/find/dataCommFind"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		for i := 0; i < n_pruebas; i++ {
			total_time := findId_op(query_bson , f, size, collCommm, conn)
			times = append(times, float64(total_time)/float64(1000000))
		}
		uso_tiempos(times) //agregar metodos aquí
	}
}

func (c Client) testComm_emptyUpdate()  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Comm/update/dataCommUpdate"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		update_bson := bson.M{"_id": "0"}
		for i := 0; i < n_pruebas; i++ {
			total_time := update_op(query_bson ,update_bson, f, size, collCommm, conn)
			times = append(times, float64(total_time)/float64(1000000))

		}
		uso_tiempos(times) //agregar metodos aquí
	}

}

func  (c Client) testComm_emptyDelete()  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Comm/delete/dataCommDelete"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		for i := 0; i < n_pruebas; i++ {
			total_time := delete_op(query_bson, f, size, collCommm, conn)
			times = append(times, float64(total_time)/float64(1000000))
		}
		uso_tiempos(times) //no hace nada aun
	}
}

func  (c Client) testComm_emptyInsert(){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Comm/insert/dataCommInsert"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		for i := 0; i < n_pruebas; i++ {
			total_time := insert_op(query_bson, f, size, collCommm, conn)
			times = append(times, float64(total_time)/float64(1000000))
		}
		uso_tiempos(times) //no hace nada aun
	}

}

// Benchmarks de latencias por operacion
/*------------------------------------------------*/

func  (c Client) InsertOne(){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)


	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Query/insert/dataInsert"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))

		f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
		for i := 0; i < n_pruebas; i++ {
			query_bson := bson.M{"data":data ,"cmp1": i, "cmp2": i, "cmp3": i}
			total_time := insert_op(query_bson, f, size,coll, conn)
			times = append(times,float64(total_time)/float64(1000000))
		}
		uso_tiempos(times)
	}
}


func (c Client) DeleteOne(){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Query/delete/dataDelete"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))

		f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
		for i := 0; i < n_pruebas; i++ {
			query_bson := bson.M{"data": data, "cmp1": i}
			total_time := delete_op(query_bson, f, size,coll, conn)
			times = append(times,float64(total_time)/float64(1000000))
		}
		uso_tiempos(times)
	}
}
func  (c Client) Coun(){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Query/find/dataFind"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))

		f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
		for i := 0; i < n_pruebas; i++ {
			query_bson := bson.M{"data": data, "cmp1": i}
			total_time := count_op(query_bson, f, size,coll, conn)
			times = append(times,float64(total_time)/float64(1000000))
		}
		uso_tiempos(times)
	}

}

func  (c Client) FindOne(){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Query/find/dataFind"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))

		f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
		for i := 0; i < n_pruebas; i++ {
			query_bson := bson.M{"data": data, "cmp1": i}
			total_time := find_op(query_bson, f, size,coll, conn)
			times = append(times,float64(total_time)/float64(1000000))
		}
		uso_tiempos(times)
	}

}

func  (c Client) FindIdOne(){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Query/find/dataFind"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))

		f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
		for i := 0; i < n_pruebas; i++ {
			query_bson := bson.M{"data": data, "cmp1": i}
			total_time := findId_op(query_bson, f, size,coll, conn)
			times = append(times,float64(total_time)/float64(1000000))
		}
		uso_tiempos(times)
	}

}

func (c Client) UpdateOne(){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)
	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Query/update/dataUpdate"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))

		f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
		for i := 0; i < n_pruebas; i++ {
			query_bson := bson.M{"data": data, "cmp1": i}
			update_bson := bson.M{"data": data, "cmp1": i, "cmp2": i, "cmp3": i}
			total_time := update_op(query_bson,update_bson, f, size,coll, conn)
			times = append(times,float64(total_time)/float64(1000000))
		}
		uso_tiempos(times)
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

