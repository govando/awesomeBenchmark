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
	"sync"
	
)

//  Benchmark de latencia de comunicación
//  operacion + colecion (comm o bench) + size request + request
//
//  query_bson := bson.M{"data": data, "cmp1": i}

type Client struct {
	num_cliente int //identificador
	timestamp 	uint64

	totalOp		clientData
}

func (c *Client) testComm_emptyCount(wg *sync.WaitGroup)  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collCommm)

	//borrar y crear el archivo de datos
	f, err := os.Create( "./data/Comm/count/dataCommCount"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		for i := 0; i < n_pruebas; i++ {
			total_time := count_op(query_bson , f, size, collCommm, conn)
			times = append(times, float64(total_time)/float64(1000000))
			c.totalOp.CountComm+=1
		}
		uso_tiempos(times) //agregar metodos aquí
	}

	wg.Done()
}

func (c *Client) testComm_emptyFind(wg *sync.WaitGroup)  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collCommm)
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
			c.totalOp.FindComm+=1
		}
		uso_tiempos(times) //agregar metodos aquí
	}
	wg.Done()
}

func (c *Client) testComm_emptyFindId(wg *sync.WaitGroup)  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collCommm)
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
			c.totalOp.FindIdComm+=1
		}
		uso_tiempos(times) //agregar metodos aquí
	}
	wg.Done()
}

func (c *Client) testComm_emptyUpdate(wg *sync.WaitGroup)  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collCommm)

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
			c.totalOp.UpdateComm+=1
		}
		uso_tiempos(times) //agregar metodos aquí
	}
	wg.Done()
}

func  (c *Client) testComm_emptyDelete(wg *sync.WaitGroup)  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collCommm)

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
			c.totalOp.DeleteComm+=1
		}
		uso_tiempos(times) //no hace nada aun
	}
	wg.Done()
}

func  (c *Client) testComm_emptyInsert(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collCommm)

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
			c.totalOp.InsertComm+=1
		}
		uso_tiempos(times) //no hace nada aun
	}
	wg.Done()
}

// Benchmarks de latencias por operacion
// Modificar aquí los query_bson
/*------------------------------------------------*/

func  (c *Client) InsertOne(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collBench)

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Query/insert/dataInsert"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))
		f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
		for i := 0; i < n_pruebas; i++ {
			query_bson := bson.M{"data":data ,"cliente":c.num_cliente,"cmp1": i, "cmp2": i, "cmp3": i}
			total_time := insert_op(query_bson, f, size,collBench, conn)
			times = append(times,float64(total_time)/float64(1000000))
			c.totalOp.Insert+=1.
		}
		uso_tiempos(times)
	}
	wg.Done()
}

func (c *Client) DeleteOne(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collBench)

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Query/delete/dataDelete"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))

		f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
		for i := 0; i < n_pruebas; i++ {
			query_bson := bson.M{"data": data, "cmp1": i}
			total_time := delete_op(query_bson, f, size,collBench, conn)
			times = append(times,float64(total_time)/float64(1000000))
			c.totalOp.Delete+=1
		}
		uso_tiempos(times)
	}
	wg.Done()
}

func  (c *Client) Count(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collBench)

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Query/count/dataCount"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))

		f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
		for i := 0; i < n_pruebas; i++ {
			query_bson := bson.M{"data": data,"cliente":c.num_cliente, "cmp1": i}
			total_time := count_op(query_bson, f, size,collBench, conn)
			times = append(times,float64(total_time)/float64(1000000))
			c.totalOp.Count+=1
		}
		uso_tiempos(times)
	}
	wg.Done()
}

func  (c *Client) FindOne(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collBench)

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Query/find/dataFind"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))

		f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
		for i := 0; i < n_pruebas; i++ {
			query_bson := bson.M{"data": data,"cliente":c.num_cliente, "cmp1": i}
			total_time := find_op(query_bson, f, size,collBench, conn)
			times = append(times,float64(total_time)/float64(1000000))
			c.totalOp.Find+=1
		}
		uso_tiempos(times)
	}
	wg.Done()
}

func  (c *Client) FindIdOne(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collBench)

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Query/findId/dataFindId"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		data := strings.Repeat("a", int(size))

		f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
		for i := 0; i < n_pruebas; i++ {
			query_bson := bson.M{"data": data, "cmp1": i}
			total_time := findId_op(query_bson, f, size,collBench, conn)
			times = append(times,float64(total_time)/float64(1000000))
			c.totalOp.FindId+=1
		}
		uso_tiempos(times)
	}
	wg.Done()
}

func (c *Client) UpdateOne(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collBench)
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
			total_time := update_op(query_bson,update_bson, f, size,collBench, conn)
			times = append(times,float64(total_time)/float64(1000000))
			c.totalOp.Update+=1
		}
		uso_tiempos(times)
	}
	wg.Done()
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

